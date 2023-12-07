package year2023.day07;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.List;

public class First extends AOCDay {

    private static int getType(String hand) {
        HashMap<Character, Integer> count = new HashMap<>();
        for(int index = 0; index < 5; index++)
            count.put(hand.charAt(index), count.getOrDefault(hand.charAt(index), 0) + 1);

        if(count.containsValue(5)) return 6;
        if(count.containsValue(4)) return 5;
        if(count.size() == 2) return 4;
        if(count.containsValue(3)) return 3;
        if(count.size() == 3) return 2;
        if(count.containsValue(2)) return 1;
        return 0;
    }

    private static char getCardValue(char first) {
        switch(first) {
            case 'A' -> { return 12; }
            case 'K' -> { return 11; }
            case 'Q' -> { return 10; }
            case 'J' -> { return 9; }
            case 'T' -> { return 8; }
            case '9' -> { return 7; }
            case '8' -> { return 6; }
            case '7' -> { return 5; }
            case '6' -> { return 4; }
            case '5' -> { return 3; }
            case '4' -> { return 2; }
            case '3' -> { return 1; }
            case '2' -> { return 0; }
            default -> throw new IllegalArgumentException();
        }
    }

    private static int compareCard(char first, char second) {
        first = getCardValue(first);
        second = getCardValue(second);
        if(first > second) return 1;
        else if(first < second) return -1;
        return 0;
    }

    private static int compareHand(String first, String second) {
        int firstType = getType(first);
        int secondType = getType(second);

        if(firstType > secondType) return 1;
        else if (firstType < secondType) return -1;

        for(int index = 0; index < 5; index++) {
            int resultCompare = compareCard(first.charAt(index), second.charAt(index));
            if(resultCompare > 0) return 1;
            else if(resultCompare < 0) return -1;
        }
        return 0;
    }

    private static int getResult(List<String> lines) {
        lines.sort((first, second) -> {
            String firstHand = first.split(" ")[0];
            String secondHand = second.split(" ")[0];
            return compareHand(firstHand, secondHand);
        });

        int sum = 0;
        for(int nLine = 0; nLine < lines.size(); nLine++) {
            sum += Integer.parseInt(lines.get(nLine).split(" ")[1]) * (nLine+1);
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "07")));
    }

}