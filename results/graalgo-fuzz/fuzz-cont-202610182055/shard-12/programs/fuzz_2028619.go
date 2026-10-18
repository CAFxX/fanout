package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{("" + "a"), "foo bar", ((("" + "\n\t") + "foo bar") + ("a" + "a")), "hello"}
    v3 := make(map[int]string)
    var v4 string = "a"
    v5 := make(map[uint32]string)
    var v6 int32 = 375
    v4 = ("" + "foo bar")
    v7 := GenS1{F0: -1, F1: 4294967295}
    var v8 int16 = 15
    v8 = v8
    v6 += int32(v8)
    _ = len(v2)
    v1 += 1
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            v3[551] = v4
            v4 = v4
        }
        _ = fi9
    }
    v4 = v4
    v6 = -2147483648
    if (2 == 42) {
        fmt.Println(v4, v8)
        var v10 int = 32767
        _ = v10
    }
    delete(v5, 154)
    var v11 string = ""
    v12 := GenS1{F0: 421, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
