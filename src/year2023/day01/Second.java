package year2023.day01;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.List;

public class Second extends AOCDay {

    private static final HashMap<String, String> numbers = new HashMap<>();

    private static void setNumbers() {
        numbers.put("one", "1");
        numbers.put("two", "2");
        numbers.put("three", "3");
        numbers.put("four", "4");
        numbers.put("five", "5");
        numbers.put("six", "6");
        numbers.put("seven", "7");
        numbers.put("eight", "8");
        numbers.put("nine", "9");
        numbers.put("1", "1");
        numbers.put("2", "2");
        numbers.put("3", "3");
        numbers.put("4", "4");
        numbers.put("5", "5");
        numbers.put("6", "6");
        numbers.put("7", "7");
        numbers.put("8", "8");
        numbers.put("9", "9");
    }

    private static int getLineValue(String line) {
        String firstStr = "", lastStr = "";
        int minPos = Integer.MAX_VALUE, maxPos = Integer.MIN_VALUE;
        for(String key : numbers.keySet()) {
            int firstIndex = line.indexOf(key);
            int lastIndex = line.lastIndexOf(key);
            if(firstIndex != -1 && firstIndex < minPos) {
                firstStr = key;
                minPos = firstIndex;
            }
            if(lastIndex != -1 && lastIndex > maxPos) {
                lastStr = key;
                maxPos = lastIndex;
            }
        }
        return Integer.parseInt(numbers.get(firstStr) + numbers.get(lastStr));
    }

    private static int getResult(List<String> lines) {
        int sum = 0;
        for(String line : lines) {
            sum += getLineValue(line);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        setNumbers();
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "01")));
    }

}
