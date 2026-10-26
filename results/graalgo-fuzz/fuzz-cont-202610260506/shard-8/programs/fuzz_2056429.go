package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{42, 829, 795}
    v3 := []string{"a", ("" + "")}
    v4 := []int{0, 2147483647}
    v5 := make(map[string]int)
    var v6 uint16 = 1
    v1 = ((v1 + v1) / 0.1)
    v1 += 0.1
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        switch 4294967295 {
        case 1:
            v7 := GenS1{F0: 2, F1: 2147483647}
            _ = v7
        }
    case "\n\t":
        var v8 uint64 = 4294967295
        _ = v8
    case "hello":
        v9 := GenS1{F0: 2, F1: 745}
        _ = v9
    default:
        delete(v5, ("\n\t" + ("hello" + "")))
    }
    var v10 int64 = 797
    if (127 <= 9223372036854775807) {
        if false {
            switch (("foo bar" + "a") + "hello") {
            case "a":
                v11 := GenS1{F0: 4294967295, F1: 470}
                _ = v11
            default:
                v10 = 0
            }
        } else {
            var v12 string = "a"
            fmt.Println(v10)
            var v13 int64 = (-(v10))
            _ = v12
            _ = v13
        }
    }
    var v14 uint = 879
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint32 = 1
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 1
    if ((v14 << ((uint(v17)) & 63)) >= (v14 / ((v14) | 1))) {
        if false {
            v18 := v5[("hello" + (("" + ("hello" + "a")) + "hello"))]
            _ = v18
        }
        _ = len(v5)
        v5[(("foo bar" + "") + "hello")] = -271
    }
    if false {
        v14 -= v14
        switch "a" {
        case "a":
            delete(v5, "\n\t")
        }
    }
    v19 := GenS1{F0: 4294967295, F1: -2147483648}
    var v20 int = -1
    v21 := GenS2{F0: false, F1: 128, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
