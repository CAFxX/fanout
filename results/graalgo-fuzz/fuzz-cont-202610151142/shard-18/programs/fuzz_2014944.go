package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 4294967295
    v3 := []string{}
    var v4 float64 = ((1.5 + 3.14159) - 1.5)
    v2 -= 0
    var v5 uint = 927
    v4 += v4
    if (!true) {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(v5, v1, v4)
            _ = fi6
        }
    }
    if (true && true) {
        switch 1 {
        case 1:
            fmt.Println(v2)
        case -9223372036854775808:
            v2 += v2
        }
    }
    switch "hello" {
    case "a":
        var v7 int = 42
        _ = v7
    case "foo bar":
        fmt.Println(v1)
    }
    _ = len(v3)
    var v8 string = (("hello" + (("hello" + "foo bar") + ("a" + (("a" + ("hello" + "\n\t")) + ("hello" + ((("\n\t" + ("\n\t" + "foo bar")) + ((("" + (("foo bar" + "foo bar") + "\n\t")) + ("foo bar" + "foo bar")) + "")) + "foo bar")))))) + "hello")
    fmt.Println(v1, v4, v5)
    fmt.Println(v4, v1)
    _ = len(v3)
    fmt.Println(v8, v5, v2)
    var v9 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
