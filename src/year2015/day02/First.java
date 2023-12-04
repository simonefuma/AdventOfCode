package year2015.day02;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;

public class First extends AOCDay {

    private static int getBoxPaperDimension(String line) {
        int[] dimensions = Arrays.stream(line.split("x"))
                .mapToInt(Integer::parseInt)
                .toArray();
        int[] squareDimensions = new int[]{
                dimensions[0] * dimensions[1],
                dimensions[1] * dimensions[2],
                dimensions[2] * dimensions[0]
        };
        return Arrays.stream(squareDimensions).map((element) -> 2 * element).sum() + Arrays.stream(squareDimensions).min().getAsInt();
    }

    private static int getResult(List<String> lines) {
        int sum = 0;
        for(String line : lines) {
           sum += getBoxPaperDimension(line);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2015", "02")));
    }

}