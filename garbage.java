import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.NoSuchElementException;
import java.util.Scanner;

public class Zadanie {
    private static class InputStringAsNumber {
        public int textLength;
        public int textValue;
        public String textSource;
    }

    static Integer convertArrayToNumber(Integer[] array) {
        Integer result = 0;
        for (Integer i = 0; i < array.length; i++) {
            result = (result * 10) + array[i];
        }
        return result;
    }

    public static void main(String[] args) {
        // System.out.println(zadanie1());
        // System.out.println(zadanie2());
        // System.out.println(zadanie34());
    }

    static InputStringAsNumber takeStringAndCheckIsNumber() {
        Scanner scanner = new Scanner(System.in);
        String inputText = scanner.nextLine();
        scanner.close();
        InputStringAsNumber number = new InputStringAsNumber();
        number.textSource = inputText;
        try {
            number.textValue = Integer.parseInt(inputText);
        } catch (NumberFormatException e) {
            return null;
        } catch (NullPointerException e) {
            return null;
        }
        number.textLength = inputText.length();
        return number;
    }

    static ArrayList<String> zadanie1() {
        ArrayList<String> secondnames = new ArrayList<String>();
        Scanner scanner = new Scanner(System.in);
        String secondname;
        outer: while (true) {
            try {
                secondname = scanner.nextLine();
            } catch (NoSuchElementException e) {
                break outer;
            }
            if (secondname.length() == 0 || secondname == null) {
                break outer;
            }
            secondnames.add(secondname);
        }
        scanner.close();
        Collections.reverse(secondnames);
        return secondnames;
    }

    static boolean zadanie2() {
        InputStringAsNumber numberFromInput = takeStringAndCheckIsNumber();
        StringBuilder inputNumberAsStrBuilder = new StringBuilder(numberFromInput.textSource);
        inputNumberAsStrBuilder.reverse();
        return inputNumberAsStrBuilder.toString().equals(numberFromInput.textSource);
    }

    static Integer zadanie34() {
        InputStringAsNumber numberFromInput = takeStringAndCheckIsNumber();
        Integer[] numbers = new Integer[numberFromInput.textLength];
        for (int i = 0; i < numberFromInput.textLength; i++) {
            numbers[i] = numberFromInput.textValue % 10;
            numberFromInput.textValue = numberFromInput.textValue / 10;
        }
        Arrays.sort(numbers, Collections.reverseOrder());
        Integer largestNumber = convertArrayToNumber(numbers);
        Arrays.sort(numbers);
        Integer smallestNumber = convertArrayToNumber(numbers);
        return largestNumber - smallestNumber;
    }
}
