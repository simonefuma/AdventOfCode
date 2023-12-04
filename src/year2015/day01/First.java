package year2015.day01;

import AOC.AOCDay;

import java.io.FileNotFoundException;

public class First extends AOCDay {

    private static int getResult(String line) {
        int floor = 0;
        for(int nChar = 0; nChar < line.length(); nChar++) {
            if(line.charAt(nChar) == '(') floor++;
            else floor--;
        }
        return floor;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToString("2015", "01")));
    }

}