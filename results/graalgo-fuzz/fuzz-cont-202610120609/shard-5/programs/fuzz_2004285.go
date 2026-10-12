package main

import (
    "fmt"
)


func main() {
    v1 := []string{("hello" + ((("a" + ("foo bar" + "foo bar")) + ("\n\t" + ("hello" + "a"))) + "foo bar"))}
    v2 := make(map[int]int)
    v3 := []string{"\n\t", "", "foo bar"}
    var v4 uint = 243
    _ = len(v2)
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v4, v4)
        _ = fi5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := v2[(fi6 ^ (fi6 & (fi6 << ((v4) & 63))))]
        var v8 int64 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 42
    var v10 float64 = -2.25
    if ((1.5 > v10) || true) {
        switch 128 {
        case 42:
            var v11 int8 = -128
            _ = v11
        case 218:
            var v12 uint64 = 462
            _ = v12
        default:
            v10 = ((0.1 * v10) / v10)
        }
        fmt.Println(v10, v10, v9)
    }
    delete(v2, 32767)
    var v13 uint16 = 832
    fmt.Println(v9)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
