import { readLines } from "https://deno.land/std@0.79.0/io/mod.ts";
import * as path from "https://deno.land/std@0.79.0/path/mod.ts";

async function getTargetValuesFromFile(targetValue: number, file: Deno.File): Promise<[number, number] | undefined> {
  let current;
  let targets: [number, number] | undefined;

  for await (const line of readLines(fileReader)) {
    if (targets) {
      break;
    }
    current = line;
    for await (const line2 of readLines(fileReader)) {
      console.log('current:', current)

      if (line2 !== current) {
        const parsed1 = parseInt(current);
        const parsed2 = parseInt(line2);

        if (isNaN(parsed1) || isNaN(parsed2)) {
          continue;
        }

        const total = parsed1 + parsed2
        console.log('total:', total)
        if (total == target) {
          targets = [parsed1, parsed2]
          break
        }
      }
    }
  }

  return targets;
}

const target = 2020;

const filename = path.join(Deno.cwd(), "data.txt");
const fileReader = await Deno.open(filename);

const result = await getTargetValuesFromFile(target, fileReader);

console.log("value 1:", result?.[0]);
console.log("value 2:", result?.[1]);
