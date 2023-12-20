# aoc2023
Advent of Code 2023

https://adventofcode.com/2023

Using Go 1.21 Std Library.

I don't really have many goals for this except to use as efficient of an algorithm as I can come up with
in a reasonable time with only the standard library. I am trying to keep each solution self-contained, but
if I am able to use a previous solution as a black-box with input/output modifications, I will do that.

Will there be a lot of code duplication? potentially, but that is just how the cookie crumbles today. I actually prefer
to keep solutions as separate as possible (unless I can use an entire solution as a black-box)

Every day's input is just an input path. Some days may take optional positional arguments, but that is just to be used 
for black-box modification of certain variables. The overall solution will still only require an input path.

## How to Run a Day's Solution
To run a specific day's solution, follow these steps:

### Clone the Repository

```bash
git clone https://github.com/ryandem1/aoc2023.git
cd aoc2023
```

### Run a day with the Makefile

```bash
make dayX
```
Replace X with the day number you want to run.

By default, the Makefile runs Part 1 with the full dataset. You can customize this by modifying the part and dataset variables in the Makefile.

part: Specifies whether to run Part 1 or Part 2 of the day's puzzle.
dataset: Refers to the dataset to use. Use 'example' for a small example provided in the problem statement, and 'full' for the full puzzle input.
Example:

```bash
make day2 part=2 dataset=example
```

The output will display the day, part, and the solution.
