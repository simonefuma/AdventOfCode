package year2023.day01;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.List;

public class First extends AOCDay {

    private static int getLineValue(String line) {
        String numbers = line.replaceAll("[^0-9]", "");
        return Integer.parseInt(new String(new char[]{numbers.charAt(0), numbers.charAt(numbers.length() - 1)}));
    }

    private static int getResult(List<String> lines) {
        int sum = 0;
        for(String line : lines) {
            sum += getLineValue(line);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "01")));
    }
}