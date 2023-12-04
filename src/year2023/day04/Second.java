package year2023.day04;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;

public class Second extends AOCDay {

    private static String[] getMyNumbers(String card) {
        return card.substring(card.indexOf(":") + 1).trim().split(" \\|")[0].split("\\s+");
    }

    private static String[] getWinningNumbers(String card) {
        return card.substring(card.indexOf("|")+1).trim().split("\\s+");
    }

    private static int getWinningScratchcards(String scratchcard) {
        String[] myNumbers = getMyNumbers(scratchcard);
        String[] winningNumbers = getWinningNumbers(scratchcard);
        int point = 0;
        for(String myNumber : myNumbers) {
            if(Arrays.asList(winningNumbers).contains(myNumber)) {
                point++;
            }
        }
        return point;
    }

    private static int getResult(List<String> scratchcards) {
        Integer[] numbersScratchcardsWon = new Integer[scratchcards.size()];
        Arrays.fill(numbersScratchcardsWon, 1);
        for(int nScratchcard = 0; nScratchcard < scratchcards.size(); nScratchcard++) {
            for(int nScratchcardWon = 1; nScratchcardWon <= getWinningScratchcards(scratchcards.get(nScratchcard)); nScratchcardWon++) {
                int position = nScratchcard+nScratchcardWon;
                numbersScratchcardsWon[position] += numbersScratchcardsWon[nScratchcard];
            }
        }
        return Arrays.stream(numbersScratchcardsWon).reduce(0, Integer::sum);
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: "+getResult(getInputToList("2023", "04")));
    }
}
