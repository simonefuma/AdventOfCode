package year2023.day05;

import AOC.AOCDay;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Second extends AOCDay {

    private static List<long[]> getSeeds(String line) {
        String[] seedsRange = line.substring(line.indexOf(":") + 2).split(" ");
        List<long[]> seeds = new ArrayList<>();
        for(int i = 0; i < seedsRange.length; i += 2) {
            long start = Long.parseLong(seedsRange[i]);
            long finish = Long.parseLong(seedsRange[i+1]);
            seeds.add(new long[]{start, start+finish});
        }
        return seeds;
    }

    private static List<List<long[]>> getCategoryMaps(List<String> lines) {
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

    private static long[] getMappedRange(long[] range, long[] map) {
        if(range[0] > map[1] && range[0] < map[1] + map[2])
            return new long[]{range[0] - map[1] + map[0], Math.min(range[1], map[2] - (range[0] - map[1]))};
        return null;
    }

    private static List<long[]> rangeSubtractions(long[] range, long[] map) {
        List<long[]> ranges = new ArrayList<>();
        //new long[]{}
        return null;
    }

    // range 45 30
    // map 30 50 14
    // mappedRange 50 - 63 -> 30 43
    // newRange 45 5 - 64 12

    private static List<long[]> getMappedRanges(List<long[]> ranges, List<long[]> category) {
        List<long[]> mappedRanges = new ArrayList<>();

        for(long[] range : ranges) {
           List<long[]> newRanges = new ArrayList<>();
           newRanges.add(range);
            for(long[] map : category) {
                for(long[] newRange : newRanges) {
                   long[] mappedRange =  getMappedRange(newRange, map);
                   if(mappedRange != null) {
                       mappedRanges.add(mappedRange);
                       newRanges.remove(newRange);
                       // aggiungi a newRanges i range ottenuti facendo newRange - map
                   }
                }
            }
        }

        return mappedRanges;
    }

    private static long getResult(List<String> lines) {
        List<long[]> ranges = getSeeds(lines.get(0));
        List<List<long[]>> categoryMaps = getCategoryMaps(lines);
        long minLocation = Integer.MAX_VALUE;

        for(List<long[]> category : categoryMaps) {
           ranges = getMappedRanges(ranges, category);
        }

        return minLocation;
    }

    public static void execute() throws FileNotFoundException {
        System.out.println("Second Puzzle Result: " + getResult(getInputToList("2023", "05")));
    }

}