package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]string)
    var v3 uint16 = v1
    v4 := []string{((("" + ("\n\t" + (("hello" + (("hello" + "hello") + "foo bar")) + ("\n\t" + "")))) + "hello") + (("" + "\n\t") + "foo bar")), "", ""}
    v5 := make(map[uint32]bool)
    var v6 int = -40
    fmt.Println(v3, v6)
    if (true || false) {
        v7 := GenS1{F0: (("\n\t" + "foo bar") + "a"), F1: 919, F2: 4294967295}
        v1 += v1
        fmt.Println(v6, v6)
        _ = v7
    }
    _ = len(v5)
    if true {
        if false {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            fmt.Println(v1, v1, v1)
        } else {
            var v8 uint32 = 718
            _ = v8
        }
        v3 -= (v1 & 0)
    }
    if false {
        v1 += (v3 & v1)
    }
    v3 -= 42
    if false {
        var v9 int8 = 111
        var v10 bool = false
        _ = v10
        _ = v9
    }
    var v11 uint16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
