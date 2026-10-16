package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = v1
    var v3 uint16 = 1
    var v4 uint64 = 42
    var v5 int8 = 1
    v1 = v2
    fmt.Println(v5, v1)
    v6 := []int{32767, -919}
    for fi7 := 0; fi7 < 16; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, (fi7 | 127))
        _ = fi7
    }
    switch "\n\t" {
    case "":
        var v8 string = "\n\t"
        _ = v8
    case "\n\t":
        var v9 int8 = v5
        _ = v9
    }
    var v10 string = ("a" + (("foo bar" + "") + ("" + (("foo bar" + "a") + ("\n\t" + (("" + "foo bar") + ("a" + "hello")))))))
    fmt.Println(v10)
    var v11 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
