package year2023.day09;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class First extends AOCDay {

    private static List<int[]> parseLines(List<String> lines) {
        List<int[]> historys = new ArrayList<>();
        for(String line : lines) {
            historys.add(
                    Arrays.stream(line.split(" "))
                    .mapToInt(Integer::valueOf)
                    .toArray()
            );
        }
        return historys;
    }

    private static List<Integer> getLastElements(int[] history) {
        List<Integer> lasts = new ArrayList<>();

        while(Arrays.stream(history).anyMatch((element) -> element != 0)) {
            int[] newHistory = new int[history.length-1];
            for(int i = 1; i < history.length; i++) {
                newHistory[i-1] = history[i] - history[i-1];
            }
            lasts.add(0, history[history.length - 1]);
            history = newHistory;
        }
        return lasts;
    }

    private static int getNextValue(int[] history) {
        List<Integer> lasts = getLastElements(history);

        int next = 0;
        for(Integer value : lasts) {
           next += value;
        }
        return next;
    }

    private static int getResult(List<String> lines) {
        List<int[]> historys = parseLines(lines);
        int result = 0;

        for(int[] history : historys) {
           result += getNextValue(history);
        }
        return result;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "09")));
    }

}