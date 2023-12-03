package year2023.day03;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;

public class Second extends AOCDay {

    private static List<Integer> getAllIndex(String line) {
        List<Integer> indexes = new ArrayList<>();
        for(int fromIndex = 0; line.indexOf("*", fromIndex) != -1; fromIndex++) {
            fromIndex = line.indexOf("*", fromIndex);
            indexes.add(fromIndex);
        }
        return indexes;
    }

    private static Integer getNumber(String line, int index) {
        StringBuilder number = new StringBuilder();
        while(index > 0 && String.valueOf(line.charAt(index-1)).matches("[0-9]")) { index--; }
        while(index < line.length() && String.valueOf(line.charAt(index)).matches("[0-9]")) {
            number.append(line.charAt(index));
            index++;
        }
        return Integer.valueOf(number.toString());
    }

    private static List<Integer> getNumbers(String line, int index) {
        List<Integer> numbers = new ArrayList<>();
        if(String.valueOf(line.charAt(index)).matches("[0-9]")) { numbers.add(getNumber(line, index)); }
        else {
            if(String.valueOf(line.charAt(index-1)).matches("[0-9]")) { numbers.add(getNumber(line, index-1)); }
            if(String.valueOf(line.charAt(index+1)).matches("[0-9]")) { numbers.add(getNumber(line, index+1)); }
        }
        return numbers;
    }

    private static int getGearRatio(List<String> scheme, int nLine, int index) {
        List<Integer> numbers = new ArrayList<>();
        for(int nLineControl = nLine-1; nLineControl <= nLine+1; nLineControl++) {
            try {numbers.addAll(getNumbers(scheme.get(nLineControl), index)); }
            catch(IndexOutOfBoundsException ignored) {}
        }
        if(numbers.size() == 2) return numbers.get(0) * numbers.get(1);
        return 0;
    }

    private static int getResult(List<String> scheme) {
        int sum = 0;
        for(int nLine = 0; nLine < scheme.size(); nLine++) {
            for(Integer index : getAllIndex(scheme.get(nLine))) {
                sum += getGearRatio(scheme, nLine, index);
            }
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "03")));
    }
}
