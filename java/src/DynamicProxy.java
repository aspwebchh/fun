import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class DynamicProxy {
    interface HelloWorldInterface{
        void hello();
        void world();
    }

    static class HelloWorld implements HelloWorldInterface{
        public void hello() {
            System.out.println("hello");
        }

        public void world() {
            System.out.println("world");
        }
    }

    static class ProxyHandler implements InvocationHandler {
        private Object proxied;
        public ProxyHandler( Object proxied) {
            this.proxied = proxied;
        }

        @Override
        public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
            System.out.println("aspwebchh");
            Object result = method.invoke(proxied,args);
            return result;
        }
    }

    public static void run() {
        HelloWorld helloWorld = new HelloWorld();
        InvocationHandler handler = new ProxyHandler(helloWorld);
        HelloWorldInterface proxy = (HelloWorldInterface) Proxy.newProxyInstance(handler.getClass().getClassLoader(),helloWorld.getClass().getInterfaces(), handler);
        proxy.hello();
        proxy.world();
    }
}
