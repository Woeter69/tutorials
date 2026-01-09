import java.util.Scanner;

public class Ex {
	public static void main() {
		Scanner scanner = new Scanner(System.in);

		double length;
		double width;

		System.out.print("Enter length of the rectangle ");
		length = scanner.nextDouble();
		System.out.print("Enter width of the rectangle ");
		width = scanner.nextDouble();

		double area = length * width;

		System.out.println("Area of the rectangle is " + area);
	}
}
