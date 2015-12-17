import java.math.BigInteger;
import java.lang.*;
import java.util.*;

public class fib{
  public static BigInteger mySum(int n){
    BigInteger[] myArray = nth(n);
    BigInteger ret = new BigInteger("0");
    for(BigInteger i : myArray) ret = ret.add(i);
    return ret;
  }
  public static BigInteger[] nth(int n){
    BigInteger[] ret = new BigInteger[n+1];
    BigInteger t = new BigInteger("1");
    ret[0] = t ;
    ret[1] = t ;
    ret[2] = t ;
    ret[3] = t ;

    for(int i = 4; i <= n; i++) {
      ret[i] = (ret[i-1].add(ret[i-2])).multiply(ret[i-3]).divide(ret[i-4]);
    }
    return ret;
  }
  public static BigInteger recersive(int n){
    BigInteger t = new BigInteger("1");
    if (n < 4) return t;

    return (recersive(n-1).add(recersive(n-2))).multiply(recersive(n-3)).divide(recersive(n-4));
  }
  public static void main(String[] args){
    int n = 209;
//    for(BigInteger i : nth(n))
//      System.out.println(i);
//    System.out.println(mySum(40));
    BigInteger[] ret = nth(40);
//    if (ret[40] == recersive(40) ) System.out.println("PASS");
    System.out.println(recersive(40));
    System.out.println(ret[40]);
  }
}
