package year2023.day08;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.List;

public class First extends AOCDay {

    private static HashMap<String, String[]> getNodes(List<String> lines) {
        HashMap<String, String[]> nodes = new HashMap<>(lines.size());
        for(String line : lines) {
            String[] split = line.split(" = ");
            nodes.put(split[0], split[1].replaceAll("[()]", "").split(", "));
        }
        return nodes;
    }

    private static int getResult(List<String> lines) {
        String instructions = lines.get(0);
        HashMap<String, String[]> nodes = getNodes(lines.subList(2, lines.size()));
        String position = "AAA";
        int step = 0;

        while(!position.equals("ZZZ")) {
            if(instructions.charAt(step % instructions.length()) == 'L') position = nodes.get(position)[0];
            else position = nodes.get(position)[1];
            step++;
        }
        return step;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "08")));
    }

}