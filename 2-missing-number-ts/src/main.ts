let rawLines = "";

process.stdin.on("data", (line) => {
  rawLines += line;
});

process.stdin.on("end", () => {
  const lines = rawLines.split("\n");

  solution(lines);
});

function sortNumbers(a: number, b: number) {
  return a < b ? -1 : a > b ? 1 : 0;
}

function solution(lines: string[]) {
  const n = parseInt(lines[0]);
  const numbers = lines[1]
    .split(" ")
    .map((value) => parseInt(value.trim()))
    .sort(sortNumbers);

  for (let idx = 0; idx < n; idx++) {
    if (idx + 1 === numbers[idx]) continue;

    process.stdout.write(`${idx + 1}\n`);
    break;
  }
}
