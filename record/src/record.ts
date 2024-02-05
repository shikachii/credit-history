import { parse } from "node-html-parser";

const SPREADSHEET_ID = PropertiesService.getScriptProperties().getProperty("SPREADSHEET_ID") ?? "";

type CreditHistory = {
  date: string;
  shop: string;
  amount: number;
  transaction: string;
  card: string;
};

const fetchUnreadMessages = (): GoogleAppsScript.Gmail.GmailMessage[] => {
  const query = "is:unread from:(statement@vpass.ne.jp) subject:(ご利用のお知らせ) after:2024-01-04";
  const threads = GmailApp.search(query);

  const unreadEmails = threads.flatMap((thread, i) => {
    const messages = thread.getMessages();
    return messages.filter((message) => message.isUnread());
  });

  return unreadEmails;
};

const parseCreditHistory = (body: string): CreditHistory => {
  const creditHistory: CreditHistory = {
    date: "",
    shop: "",
    amount: 0,
    transaction: "",
    card: "",
  };

  const root = parse(body);

  const preProcess = (text: string) => text.trim().replaceAll("\n", "").replaceAll("\r", "");

  const query =
    "html body table tr td table tr td table tbody tr td table tbody tr td table tbody tr td table tbody tr td";
  root.querySelectorAll(query).forEach(({ text }, index) => {
    const preProcessedText = preProcess(text);

    switch (index) {
      case 0: {
        // get date
        const date = preProcessedText.replace("ご利用日時：", "");
        creditHistory.date = date;
        break;
      }
      case 1: {
        //get shop & transaction
        const [shop] = preProcessedText.split("（");
        creditHistory.shop = shop;

        const regexp = /（(.*)）/;
        const match = preProcessedText.match(regexp);
        const transaction = match ? match[match.length - 1] : "";
        creditHistory.transaction = transaction;
        break;
      }
      case 2: {
        // get amount
        const amount = parseInt(preProcessedText.split("円")[0].replaceAll(",", ""));
        creditHistory.amount = amount;
        break;
      }
    }
  });

  const cardQuery = "html body table tr td table tr td";
  root.querySelectorAll(cardQuery).forEach(({ text }, index) => {
    if (index === 6) {
      const preProcessedText = preProcess(text);
      const regexp = /ございます。(.*)について/;
      const match = preProcessedText.match(regexp);

      if (match && match[1]) {
        creditHistory.card = match[1];
      }
    }
  });

  return creditHistory;
};

const recordToSpreadSheet = (creditHistory: CreditHistory) => {
  const spreadSheet = SpreadsheetApp.openById(SPREADSHEET_ID);
  const mainSheet = spreadSheet.getSheetByName("main");

  mainSheet?.appendRow(Object.values(creditHistory));
};

export const record = () => {
  const unreadMessages = fetchUnreadMessages();
  unreadMessages.forEach((message) => {
    const body = message.getBody();
    const creditHistory: CreditHistory = parseCreditHistory(body);
    recordToSpreadSheet(creditHistory);

    message.markRead();
  });
};
