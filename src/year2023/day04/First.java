package year2023.day04;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;

public class First extends AOCDay {

    private static String[] getMyNumbers(String scratchcard) {
        return scratchcard.substring(scratchcard.indexOf(":") + 1).trim().split(" \\|")[0].split("\\s+");
    }

    private static String[] getWinningNumbers(String scratchcard) {
        return scratchcard.substring(scratchcard.indexOf("|")+1).trim().split("\\s+");
    }

    private static int getPoint(String scratchcard) {
        String[] myNumbers = getMyNumbers(scratchcard);
        String[] winningNumbers = getWinningNumbers(scratchcard);
        int point = 0;
        for(String myNumber  : myNumbers) {
           if(Arrays.asList(winningNumbers).contains(myNumber)) {
               if(point == 0) point++;
               else point *= 2;
           }
        }
        return point;
    }

    private static int getResult(List<String> scratchcards) {
        int sum = 0;
        for(String scratchcard : scratchcards) {
            sum += getPoint(scratchcard);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "04")));
    }
}
