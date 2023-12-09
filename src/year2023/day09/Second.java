package year2023.day09;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Second extends AOCDay {

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

    private static List<Integer> getFirstElements(int[] history) {
        List<Integer> firsts = new ArrayList<>();

        while(Arrays.stream(history).anyMatch((element) -> element != 0)) {
            int[] newHistory = new int[history.length-1];
            for(int i = 1; i < history.length; i++) {
                newHistory[i-1] = history[i] - history[i-1];
            }
            firsts.add(0, history[0]);
            history = newHistory;
        }
        return firsts;
    }

    private static int getNextValue(int[] history) {
        List<Integer> firsts = getFirstElements(history);

        int next = 0;
        for(Integer value : firsts) {
            next = value - next;
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
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "09")));
    }

}