const SPREADSHEET_ID = PropertiesService.getScriptProperties().getProperty("SPREADSHEET_ID") ?? "";

type CreditHistory = {
  date: string;
  shop: string;
  amount: number;
  transaction: string;
  card: string;
};

export const notify = () => {
  // const today = new Date();
  // const yesterday = new Date();
  // yesterday.setDate(today.getDate() - 1);
  // const spreadSheet = SpreadsheetApp.openById(SPREADSHEET_ID);
  // const sheet = spreadSheet.getSheetByName("main");
  // try {
  //   if (!sheet) {
  //     throw new Error("Sheet not found");
  //   }
  // } catch (error) {
  //   console.error(error);
  //   return;
  // }
  // const lastRow = sheet.getLastRow();
  // const lastColumn = sheet.getLastColumn();
  // const allRange = sheet.getRange(2, 1, lastRow, lastColumn).getValues();
  // const creditHistories: CreditHistory[] = allRange.map((row) => {
  //   return {
  //     date: row[1],
  //     shop: row[2],
  //     amount: parseInt(row[3]),
  //     transaction: row[4],
  //     card: row[5],
  //   };
  // });
  // // 前日の日付の合計額を取得
  // const yesterdayCreditHistories = creditHistories.filter((creditHistory) => {
  //   return new Date(creditHistory.date).getDate() === yesterday.getDate();
  // });
  // const yesterdayTotalAmount = yesterdayCreditHistories.reduce((acc, creditHistory) => {
  //   return acc + creditHistory.amount;
  // }, 0);
  // // 前日までの月間合計額を取得
  // const thisMonthCreditHistories = creditHistories.filter((creditHistory) => {
  //   return new Date(creditHistory.date).getMonth() === today.getMonth();
  // });
  // const thisMonthTotalAmount = thisMonthCreditHistories.reduce((acc, creditHistory) => {
  //   return acc + creditHistory.amount;
  // }, 0);
  // // slackに通知
  // // const slackApp = SlackApp.create("SLACK_TOKEN");
  // // slackApp.postMessageToChannel("channel", `昨日の合計額: ${yesterdayTotalAmount}`);
  // // slackApp.postMessageToChannel("channel", `今月の合計額: ${thisMonthTotalAmount}`);
};
