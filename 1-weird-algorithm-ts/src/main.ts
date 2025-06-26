process.stdin.on("data", (line) => {
  let number = parseInt(line.toString());

  process.stdout.write(`${number} `);

  while (number != 1) {
    number = number % 2 === 0 ? number / 2 : number * 3 + 1;
    process.stdout.write(`${number} `);
  }

  process.stdout.write("\n");
  process.exit();
});
