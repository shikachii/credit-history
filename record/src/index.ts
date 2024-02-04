import { parse } from "node-html-parser";

const SPREADSHEET_ID =
  PropertiesService.getScriptProperties().getProperty('SPREADSHEET_ID') ?? "";
const API_KEY = PropertiesService.getScriptProperties().getProperty('API_KEY') ?? "";
const EP_URL = PropertiesService.getScriptProperties().getProperty('EP_URL') ?? "";

type CreditHistory = {
  date: Date;
  shop: string;
  amount: number;
  transaction: string;
  card: string;
}

const fetchUnreadMessages = (): GoogleAppsScript.Gmail.GmailMessage[] => {
  const query = 'is:unread from:(statement@vpass.ne.jp) subject:(ご利用のお知らせ) after:2023-12-06';
  const threads = GmailApp.search(query);

  const unreadEmails = threads.flatMap((thread) => {
    const messages = thread.getMessages()
    return messages.filter((message) => message.isUnread())
  });

  return unreadEmails;
}

const parseCreditHistory = (body: string): CreditHistory => {
  const root = parse(body);
  root.querySelector("tr");

  return {
    date: new Date(),
    shop: "",
    amount: 0,
    transaction: "",
    card: "",
  }
};

const recordToSpreadSheet = (creditHistory: CreditHistory) => {
  const spreadSheet = SpreadsheetApp.openById(SPREADSHEET_ID);
  const mainSheet = spreadSheet.getSheetByName("main");

  mainSheet?.appendRow(Object.values(creditHistory));
}

const record = () => {
  const unreadMessages = fetchUnreadMessages();
  unreadMessages.forEach((message) => {
    const body = message.getBody();
    const creditHistory: CreditHistory = parseCreditHistory(body);
    recordToSpreadSheet(creditHistory);

    message.markRead();
  });
}

const notify = () => {

}

const recordCreditHistory = () => {
  const query =
    'is:unread from:(statement@vpass.ne.jp) subject:(ご利用のお知らせ) after:2023-12-06'
  const threads = GmailApp.search(query)

  threads.forEach((thread) => {
    const messages = thread.getMessages()
    messages.forEach((message) => {
      // メールが未読のときのみ発火
      if (message.isUnread()) {
        const body = message.getBody()

        const data = {
          email: body.replace(/\r/g, '').replace(/\n/g, '\n '),
        }
        const options: GoogleAppsScript.URL_Fetch.URLFetchRequestOptions = {
          method: 'post',
          headers: {
            'Content-Type': 'application/json',
            'x-api-key': API_KEY,
          },
          payload: JSON.stringify(data),
        }

        const result = UrlFetchApp.fetch(EP_URL, options)

        if (result.getResponseCode() === 200) {
          const resultText = result.getContentText()
          const resultJson = JSON.parse(resultText)

          const spshe = SpreadsheetApp.openById(SPREADSHEET_ID)
          const sheet = spshe.getSheetByName('main')

          sheet?.appendRow(Object.values(resultJson))
        }

        // message.markRead(); // 既読に設定
      }
    })
  })
}

const reportCreditHistory = () => {
  const spshe = SpreadsheetApp.openById(SPREADSHEET_ID)
  const sheet = spshe.getSheetByName('main')

  const filter = sheet?.getFilter()
  console.log(filter?.getRange().getColumn())
}
