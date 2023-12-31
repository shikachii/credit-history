const SPREADSHEET_ID =
  PropertiesService.getScriptProperties().getProperty("SPREADSHEET_ID");
const API_KEY = PropertiesService.getScriptProperties().getProperty("API_KEY");
const EP_URL = PropertiesService.getScriptProperties().getProperty("EP_URL");

const recordCreditHistory = () => {
  const query =
    "is:unread from:(statement@vpass.ne.jp) subject:(ご利用のお知らせ) after:2023-6-10";
  const threads = GmailApp.search(query);

  threads.forEach((thread) => {
    const messages = thread.getMessages();
    messages.forEach((message) => {
      // メールが未読のときのみ発火
      if (message.isUnread()) {
        const body = message.getBody();
        
        console.log(body);

        const data = {
          email: body.replace(/\r/g, "").replace(/\n/g, "\n "),
        };
        const options = {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "x-api-key": API_KEY,
          },
          payload: JSON.stringify(data),
        };

        const result = UrlFetchApp.fetch(EP_URL, options);

        if (result.getResponseCode() === 200) {
          const resultText = result.getContentText();
          const resultJson = JSON.parse(resultText);

          const spshe = SpreadsheetApp.openById(SPREADSHEET_ID);
          const sheet = spshe.getSheetByName("main");

          sheet.appendRow(Object.values(resultJson));
        }

        // message.markRead(); // 既読に設定
      }
    });
  });
};

const reportCreditHistory = () => {
  const spshe = SpreadsheetApp.openById(SPREADSHEET_ID);
  const sheet = spshe.getSheetByName("main");

  const filter = sheet.getFilter();
  console.log(filter.getRange().getColumn());
};
