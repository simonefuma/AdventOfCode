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

    private static HashMap<String, List<Integer>> initializePositionPatternMap(List<String> positions) {
        HashMap<String, List<Integer>> positionPattern = new HashMap<>(positions.size());
        for(String position : positions) positionPattern.put(position, new ArrayList<>());
        return positionPattern;
    }

    private static int mcm(int... numbers) {
        int mcm = Arrays.stream(numbers).max().getAsInt();
        while(Arrays.stream(numbers).allMatch((number) -> mcm % number == 0)) mcm++;
        return mcm;
    }

    private static int getResult(List<String> lines) {
        String instructions = lines.get(0);
        HashMap<String, String[]> nodes = getNodes(lines.subList(2, lines.size()));
        List<String> positions = getPositions(nodes.keySet());
        HashMap<String, List<Integer>> positionsPattern = initializePositionPatternMap(positions);

        for(String position : positions) {
            String stepPosition = position;
            List<String> startPositions = new ArrayList<>();
            int step = 0;
            while(true) {
                if(step % instructions.length() == 0) {
                    if(startPositions.contains(stepPosition)) break;
                    startPositions.add(stepPosition);
                }
                if(stepPosition.charAt(2) == 'Z') positionsPattern.get(position).add(step);
                stepPosition = nodes.get(stepPosition)[(instructions.charAt(step % instructions.length()) == 'L') ? 0 : 1];
                step++;
            }
            System.out.println(positionsPattern.get(position));
        }


        return 0;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "08")));
    }

}