package year2023.day07;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.List;

public class Second extends AOCDay {

    private static final List<Character> cards = List.of('A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2');

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
            case 'T' -> { return 9; }
            case '9' -> { return 8; }
            case '8' -> { return 7; }
            case '7' -> { return 6; }
            case '6' -> { return 5; }
            case '5' -> { return 4; }
            case '4' -> { return 3; }
            case '3' -> { return 2; }
            case '2' -> { return 1; }
            case 'J' -> { return 0; }
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

    private static String replaceJoker(String hand) {
        if(!hand.contains("J")) return hand;
        int point = 0;
        String newHand = "";
        for(Character card : cards) {
            String tempHand = replaceJoker(hand.replaceFirst("J", String.valueOf(card)));
            int newPoint = getType(tempHand);
            if(newPoint > point) {
                point = newPoint;
                newHand = tempHand;
            }
        }
        return newHand;
    }

    private static int compareHand(String first, String second) {
        int firstType = getType(replaceJoker(first));
        int secondType = getType(replaceJoker(second));

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
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "07")));
    }

}