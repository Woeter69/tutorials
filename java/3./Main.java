import java.util.Scanner;

public class Main {
	public static void main() {
		// TO get input from user we have to make a scanner object which takes user input

		Scanner scanner = new Scanner(System.in); // new Scanner is capital

		System.out.print("Enter your name"); // print does not give a new line 
		String name = scanner.nextLine(); // Line is capital always fuck this up btw

		System.out.println("Hello" + name);
		
		System.out.print("Enter your age");

		int age = scanner.nextInt();

		scanner.nextLine(); // If I don't do this when i press enter after entering the digit the new like character stays in buffer and the \n becomes the input for the String which is inputted after the digit in our case color so we just add a newline after the scanning of a number to avoid these issues this basically clears the input buffer.

		System.out.print("Enter your color");

		String color = scanner.nextLine();

		System.out.print("Hello " + name + " Your age is " + age);

		
		scanner.close(); // Same thing always close your object after program is done
	}
}
