package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    var v2 int = -2147483648
    v3 := make(map[string]string)
    v4 := []string{("\n\t" + (("" + ("a" + "foo bar")) + "\n\t")), "hello"}
    v5 := make(map[int]string)
    var v6 int = (v2 - (v2 & v2))
    var v7 string = (("\n\t" + "\n\t") + "foo bar")
    v6 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6, v6)
    fmt.Println(v7, v6)
    var v8 int32 = -639
    switch (v6 / ((v2) | 1)) {
    case 0:
        var v9 bool = true
        _ = v9
    }
    var v10 int64 = 0
    v6 = (-(v6))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
