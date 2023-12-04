package AOC.Setup;

import java.util.StringJoiner;
import java.util.Scanner;
import java.io.*;

public class Setup {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        System.out.print("Please enter the year: ");
        String year = br.readLine();
        System.out.print("Please enter the day: ");
        String day = br.readLine();

        createYearDirectory(year);

        if(!createDayDirectory(year, day)) {
            System.out.println("Error: This day's directory already exists.");
            return;
        }

        createSetup(year, day);
    }

    private static void createSetup(String year, String day) throws IOException {
        createInput(year, day);
        createFile("First", year, day);
        createFile("Second", year, day);
        createFile("Main", year, day);
    }

    private static void createYearDirectory(String year) {
        File yearDirectory = new File(System.getProperty("user.dir"), String.format("src/year%s", year));
        if(!yearDirectory.exists()) { yearDirectory.mkdir(); }
    }

    private static boolean createDayDirectory(String year, String day) {
        File dayDirectory = new File(System.getProperty("user.dir"), String.format("src/year%s/day%s", year, day));
        if(!dayDirectory.exists()) {
            dayDirectory.mkdir();
            return true;
        }
        return false;
    }

    private static void createInput(String year, String day) throws IOException {
        File inputDirectory = new File(System.getProperty("user.dir"), String.format("src/year%s/day%s/input", year, day));
        inputDirectory.mkdir();
        File inputFile = new File(inputDirectory, "input.txt");
        inputFile.createNewFile();
        gitAdd(inputFile);
    }

    private static void createFile(String fileName, String year, String day) throws IOException {
        StringJoiner sj = new StringJoiner("\n");
        Scanner scanner = new Scanner(new File(System.getProperty("user.dir"), String.format("src/AOC/Setup/%s.txt", fileName)));
        while(scanner.hasNextLine()) sj.add(scanner.nextLine());

        File file = new File(System.getProperty("user.dir"), String.format("src/year%s/day%s/%s.java", year, day, fileName));
        file.createNewFile();
        FileWriter writer = new FileWriter(file);
        writer.write(sj.toString().replaceAll("\\{year}", year).replaceAll("\\{day}", day));
        writer.close();
        gitAdd(file);
    }

    private static void gitAdd(File file) throws IOException {
        ProcessBuilder processBuilder = new ProcessBuilder("git", "add", file.getAbsolutePath());
        processBuilder.start();
    }

}
