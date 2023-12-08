package year2023.day05;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class First extends AOCDay {

    private static long[] getSeeds(String line) {
        return Arrays.stream(line.substring(line.indexOf(":") + 2).split(" "))
                .mapToLong(Long::parseLong)
                .toArray();
    }

    private static List<List<long[]>> getMaps(List<String> lines) {
        List<List<long[]>> maps = new ArrayList<>();
        for(int nLine = 2; nLine < lines.size(); nLine++) {
            if(lines.get(nLine).contains(":")) {
                List<long[]> map = new ArrayList<>();
                nLine++;
                while(nLine < lines.size() && !lines.get(nLine).isBlank()) {
                    map.add(Arrays.stream(lines.get(nLine).split(" ")).mapToLong(Long::parseLong).toArray());
                    nLine++;
                }
                maps.add(map);
            }
        }
        return maps;
    }

    private static long getLocation(long seed, List<List<long[]>> maps) {
        for(List<long[]> Category : maps) {
            for(long[] map : Category) {
                long difference = seed - map[1];
                if(difference >= 0 && difference < map[2]) {
                    seed = map[0] + difference;
                    break;
                }
            }
        }
        return seed;
    }

    private static long getResult(List<String> lines) {
        long[] seeds = getSeeds(lines.get(0));
        List<List<long[]>> maps = getMaps(lines);

        long minLocation = Integer.MAX_VALUE;
        for(long seed : seeds) {
            long location = getLocation(seed, maps);
            if(location < minLocation) minLocation = location;
        }

        return minLocation;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("First Puzzle Result: " + getResult(getInputToList("2023", "05")));
    }

}