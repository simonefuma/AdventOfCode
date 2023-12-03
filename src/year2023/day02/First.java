package year2023.day02;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class First extends AOCDay {

    private static final Map<String, Integer> maxCubes = Map.of("red", 12, "green", 13, "blue", 14);

    private static boolean isPossibleSet(Map<String, Integer> set) {
        for(String color : set.keySet()) {
            if(set.getOrDefault(color, 0) > maxCubes.get(color)) return false;
        }
        return true;
    }

    private static boolean isPossibleGame(List<Map<String, Integer>> game) {
        for(Map<String,Integer> set : game) {
            if(!isPossibleSet(set)) return false;
        }
        return true;
    }

    private static Map<String, Integer> setOf(String setString) {
        Map<String, Integer> set = new HashMap<>();
        for(String cubeString : setString.split(",")) {
            String[] cubeArray = cubeString.trim().split(" ");
            set.put(cubeArray[1], Integer.parseInt(cubeArray[0]));
        }
        return set;
    }

    private static List<Map<String, Integer>> gameOf(String gameString) {
        List<Map<String, Integer>> game = new ArrayList<>();
        for(String setString : gameString.substring(gameString.indexOf(":")+1).split(";")) {
            game.add(setOf(setString));
        }
        return game;
    }

    private static int gameNumber(String gameString) {
        return Integer.parseInt(gameString.substring(5, gameString.indexOf(":")));
    }

    private static int getResult(List<String> games) {
        int sum = 0;
        for(String gameString : games) {
            if(isPossibleGame(gameOf(gameString))) {
                sum += gameNumber(gameString);
            }
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "02")));
    }

}
