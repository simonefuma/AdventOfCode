package AOC;

import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.List;
import java.io.File;

public abstract class AOCDay {

    protected static List<String> getInputToList(String year, String day) throws FileNotFoundException {
        List<String> input = new ArrayList<>();
        Scanner scanner = new Scanner(new File(System.getProperty("user.dir"), String.format("src/year%s/day%s/input/input.txt", year, day)));
        while(scanner.hasNextLine()) input.add(scanner.nextLine());
        return input;
    }

}