const SPREADSHEET_ID = PropertiesService.getScriptProperties().getProperty("SPREADSHEET_ID");
const SLACK_WEBHOOK_URL = PropertiesService.getScriptProperties().getProperty("SLACK_WEBHOOK_URL");

type CreditHistory = {
  date: string;
  shop: string;
  amount: number;
  transaction: string;
  card: string;
};

export const notify = () => {
  if (!SPREADSHEET_ID) {
    throw new Error("SPREADSHEET_ID not found");
  }
  const sheet = SpreadsheetApp.openById(SPREADSHEET_ID);

  // day-aggregateシートの最後尾の行を取得
  const dayAggregateSheet = sheet.getSheetByName("day-aggregate");
  if (!dayAggregateSheet) {
    throw new Error("Sheet not found");
  }
  const dayAggregateSheetLastRow = dayAggregateSheet.getLastRow();
  const dayAggregateAmount = dayAggregateSheet.getRange(dayAggregateSheetLastRow, 2).getValue() as number;

  // aggregateシートの最後尾の行を取得
  const aggregateSheet = sheet.getSheetByName("aggregate");
  if (!aggregateSheet) {
    throw new Error("Sheet not found");
  }
  const aggregateSheetLastRow = aggregateSheet.getLastRow();
  const aggregateAmount = aggregateSheet.getRange(aggregateSheetLastRow, 2).getValue() as number;

  // 日・月の使用金額をslackに通知
  if (!SLACK_WEBHOOK_URL) {
    throw new Error("SLACK_WEBHOOK_URL not found");
  }
  UrlFetchApp.fetch(SLACK_WEBHOOK_URL, {
    method: "post",
    headers: {
      "Content-Type": "application/json",
    },
    payload: JSON.stringify({
      username: "通知くん",
      icon_emoji: ":moneybag:",
      text: `昨日の使用金額: ${dayAggregateAmount}円\n今月の使用金額: ${aggregateAmount}円`,
    }),
  });

  // day-aggregateシートの最後尾に今日の日付を追加
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  dayAggregateSheet.appendRow([
    today,
    `=SUMIFS(main!$D$2:$D$996, main!$B$2:$B$996, ">="&A${dayAggregateSheetLastRow + 1}, main!$B$2:$B$996, "<"&A${dayAggregateSheetLastRow + 1}+1)`,
  ]);

  // 今日が月初の場合、aggregateシートに今日の日付を追加
  if (today.getDate() === 1) {
    aggregateSheet.appendRow([
      today,
      `=SUMIFS(main!$D$2:$D$996, main!$B$2:$B$996, ">="&A${aggregateSheetLastRow + 1}, main!$B$2:$B$996, "<"&EDATE(A${aggregateSheetLastRow + 1}, 1))`,
    ]);
  }
};
