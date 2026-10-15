package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 42
    var v2 uint64 = 42
    v3 := []string{"hello", ((("" + "") + "\n\t") + ("" + ("a" + "\n\t"))), (("" + "hello") + "hello"), "\n\t"}
    v4 := make(map[int]int)
    _ = len(v4)
    var v5 int64 = 9223372036854775807
    fmt.Println(v5, v1, v1)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int8 = (v1 % ((v1) | 1))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 += v2
        _ = fi8
    }
    v5 = v5
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v5, v1)
        if (0 < 2) {
            v5 = v5
        } else {
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            var v10 string = "foo bar"
            _ = v10
        }
        _ = fi9
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
