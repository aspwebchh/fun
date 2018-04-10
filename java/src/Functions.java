import java.security.MessageDigest;
import java.util.regex.Pattern;

public class Functions {
    public static boolean isNullOrEmpty(String str ) {
        if( str == null) {
            return true;
        }
        if( str.trim().length() == 0 ) {
            return true;
        }
        return false;
    }


    public static boolean isInteger( String val ) {
        if( val == null || val.trim().length() == 0 ) {
            return false;
        }
        Pattern pattern = Pattern.compile("^\\d+$");
        if( pattern.matcher(val.trim()).find()) {
            return true;
        }
        return false;
    }

    public  static String MD5(String s) {
        char hexDigits[]={'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'};
        try {
            byte[] btInput = s.getBytes();
            // 获得MD5摘要算法的 MessageDigest 对象
            MessageDigest mdInst = MessageDigest.getInstance("MD5");
            // 使用指定的字节更新摘要
            mdInst.update(btInput);
            // 获得密文
            byte[] md = mdInst.digest();
            // 把密文转换成十六进制的字符串形式
            int j = md.length;
            char str[] = new char[j * 2];
            int k = 0;
            for (int i = 0; i < j; i++) {
                byte byte0 = md[i];
                str[k++] = hexDigits[byte0 >>> 4 & 0xf];
                str[k++] = hexDigits[byte0 & 0xf];
            }
            return new String(str);
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }
}
