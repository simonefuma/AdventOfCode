package year2023.day06;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;

public class First extends AOCDay {

    private static int[] getArrayLine(String line) {
        return Arrays.stream(line.substring(line.indexOf(":") + 1).trim().split("\\s+"))
                .mapToInt(Integer::parseInt)
                .toArray();
    }

    private static int getMarginOfError(int time, int distance) {
        int min = 0;
        while((time - min) * min <= distance) min++;

        int max = time;
        while((time - max) * max <= distance) max--;

        return (max - min) + 1;
    }

    private static int getResult(List<String> lines) {
        int result = 1;
        int[] times = getArrayLine(lines.get(0));
        int[] distances = getArrayLine(lines.get(1));

        for(int nRace = 0; nRace < times.length; nRace++) {
            result *= getMarginOfError(times[nRace], distances[nRace]);
        }
        return result;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "06")));
    }

}