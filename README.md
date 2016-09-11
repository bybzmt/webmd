#Web Markdown 查看工具

这是一个在Web上查看Markdown的工具

* 支持语法高亮
* 支持graphviz绘图

如果需要使用graphviz需要单独安装它

语法高亮跟github是同一格式

    ```go
    func main() {
        log.Pinttln("test");
    }
    ```

graphviz按下面这样写就可以了 (`<dot>`必需顶格写)

    <dot>
    digraph G {
        size="4,4";
        main [shape=box]; /*注释*/
        main -> parse [weight=8];
        parse -> execute;
        main -> init [style=dotted];
        main -> cleanup;
        execute -> {make_string; printf}
        init -> make_string;
        edge [color=red];
        main -> printf [style=bold,label="100 times"];
        make_string [label="make a\nstring"];
        node [shape=box, style=filled, color=".7 .3 1.0"];
        execute -> compare;
    }
    </dot>

上面的脚本会显成这张图：

![dot绘图](https://raw.githubusercontent.com/bybzmt/webmd/master/dot_example.png)


