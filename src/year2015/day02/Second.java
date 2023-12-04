package year2015.day02;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;

public class Second extends AOCDay {

    private static int getRibbonLength(String line) {
        int[] dimensions = Arrays.stream(line.split("x"))
                .mapToInt(Integer::parseInt)
                .toArray();
        return Arrays.stream(Arrays.stream(dimensions)
                .sorted()
                .limit(2)
                .toArray())
                .map((element) -> 2 * element)
                .sum() + Arrays.stream(dimensions).reduce(1, (x, y) -> x * y);
    }

    private static int getResult(List<String> lines) {
        int sum = 0;
        for(String line : lines) {
            sum += getRibbonLength(line);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2015", "02")));
    }

}