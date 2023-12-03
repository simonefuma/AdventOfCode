package year2023.day02;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


public class Second extends AOCDay {

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

    private static int getGamePower(List<Map<String, Integer>> game) {
        Map<String, Integer> maxCube = new HashMap<>();
        for(Map<String,Integer> set : game) {
            for(String color : set.keySet()) {
                if(set.getOrDefault(color, 0) > maxCube.getOrDefault(color, 0)) maxCube.put(color, set.get(color));
            }
        }
        return maxCube.values().stream().reduce(1, (x, y) -> x * y);
    }

    private static int getResult(List<String> games) {
        int sum = 0;
        for(String gameString : games) {
            sum += getGamePower(gameOf(gameString));
        }
        return sum;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "02")));
    }

}
