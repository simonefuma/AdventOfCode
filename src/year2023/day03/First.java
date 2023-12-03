package year2023.day03;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.List;

public class First extends AOCDay {

    private static boolean isValid(List<String> scheme, int nLine, int nChar) {
        for(int nLineControl = nLine-1; nLineControl <= nLine+1; nLineControl++) {
            for(int nCharControl = nChar-1; nCharControl <= nChar+1; nCharControl++) {
                try {
                    if(String.valueOf(scheme.get(nLineControl).charAt(nCharControl)).matches("[^0-9.]"))
                        return true;
                } catch(IndexOutOfBoundsException ignored) {}
            }
        }
        return false;
    }

    private static int getResult(List<String> scheme) {
        int sum = 0;
        StringBuilder number = new StringBuilder();
        boolean isValid = false;
        for(int nLine = 0; nLine < scheme.size(); nLine++) {
            String line = scheme.get(nLine) + ".";
            for(int nChar = 0; nChar < line.length(); nChar++) {
                String token = String.valueOf(line.charAt(nChar));
                if(token.matches("[0-9]")) {
                    if(isValid(scheme, nLine, nChar)) isValid = true;
                    number.append(token);
                } else {
                    if(isValid) {
                        sum += Integer.parseInt(number.toString());
                        isValid = false;
                    }
                    if(number.length() > 0) number = new StringBuilder();
                }
            }
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "03")));
    }
}
