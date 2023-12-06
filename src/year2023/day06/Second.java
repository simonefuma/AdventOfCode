package year2023.day06;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.List;

public class Second extends AOCDay {

    private static long getLineNumber(String line) {
        return Long.parseLong(line.substring(line.indexOf(":") + 1).replaceAll("\\s+", ""));
    }

    private static long getMarginOfError(long time, long distance) {
        long min = 0;
        while((time - min) * min <= distance) min++;

        long max = time;
        while((time - max) * max <= distance) max--;

        return (max - min) + 1;
    }

    private static long getResult(List<String> lines) {
        long time = getLineNumber(lines.get(0));
        long distance = getLineNumber(lines.get(1));
        return getMarginOfError(time, distance);
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "06")));
    }

}