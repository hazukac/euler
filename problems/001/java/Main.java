
public class Main {
    private static int sumMultiplesOf(int[] factors, int limit) {
        int res = 0;
        for (int i = 0; i < limit; i++) {
            res += summable(i, factors);
        }
        return res;
    }
    private static int summable(int subject, int[] factors) {
        for (int factor: factors) {
            if (subject % factor == 0) {
                return subject;
            }
        }
        return 0;
    }
    public static void main(String[] args) {
        int[] factors = {3, 5};
        System.out.println(sumMultiplesOf(factors, 1000));
    }
}
