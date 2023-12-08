package year2023.day08;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.*;

public class Second extends AOCDay {

    private static HashMap<String, String[]> getNodes(List<String> lines) {
        HashMap<String, String[]> nodes = new HashMap<>(lines.size());
        for(String line : lines) {
            String[] split = line.split(" = ");
            nodes.put(split[0], split[1].replaceAll("[()]", "").split(", "));
        }
        return nodes;
    }

    private static List<String> getPositions(Set<String> nodes) {
        List<String> startNodes = new ArrayList<>();
        for(String node : nodes) {
           if(node.charAt(2) == 'A') startNodes.add(node);
        }
        return startNodes;
    }

    private static long mcd(long a, long b) {
        while (b != 0) {
            long temp = b;
            b = a % b;
            a = temp;
        }
        return a;
    }

    private static long mcm(long a, long b) {
        return a * (b / mcd(a, b));
    }

    private static long mcm(HashMap<String, Integer> numbers) {
        long mcm = 1;
        for(String position : numbers.keySet()) {
            mcm = mcm(mcm, numbers.get(position));
        }
        return mcm;
    }

    private static long getResult(List<String> lines) {
        String instructions = lines.get(0);
        HashMap<String, String[]> nodes = getNodes(lines.subList(2, lines.size()));
        List<String> positions = getPositions(nodes.keySet());
        HashMap<String, Integer> positionsPattern = new HashMap<>(positions.size());

        for(String position : positions) {
            String stepPosition = position;
            List<String> startPositions = new ArrayList<>();
            int step = 0;
            while(true) {
                if(step % instructions.length() == 0) {
                    if(startPositions.contains(stepPosition)) break;
                    startPositions.add(stepPosition);
                }
                if(stepPosition.charAt(2) == 'Z') positionsPattern.put(position, step);
                stepPosition = nodes.get(stepPosition)[(instructions.charAt(step % instructions.length()) == 'L') ? 0 : 1];
                step++;
            }
        }
        return mcm(positionsPattern);
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "08")));
    }

}