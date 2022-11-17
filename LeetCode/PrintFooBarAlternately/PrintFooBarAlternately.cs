using System.Threading;

//doesn't have golang version... sad
//description : https://leetcode.com/problems/print-foobar-alternately/description/
public class FooBar {
    private int n;

    private bool _lock;

    public FooBar(int n) {
        this.n = n;
        _lock = false;
    }

    public void Foo(Action printFoo) {
        for (int i = 0; i < n; i++) {
            if (_lock) {
                i--;
                continue;
            }
        	// printFoo() outputs "foo". Do not change or remove this line.
        	printFoo();
            _lock = true;
        }
    }

    public void Bar(Action printBar) {
        for (int i = 0; i < n; i++) {
           if (!_lock) {
                i--;
                continue;
            }
            // printBar() outputs "bar". Do not change or remove this line.
        	printBar();
            _lock = false;
        }
    }
}