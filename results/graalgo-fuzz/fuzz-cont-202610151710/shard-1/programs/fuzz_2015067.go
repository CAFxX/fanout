package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint64 = 1
        _ = fi5
        _ = v6
    }
    p0 = true
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "hello"
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]string)
    v2 = ("" + ("\n\t" + "a"))
    v3 -= 464
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42, F1: (v3 ^ 57)}
        _ = fi7
        _ = v8
    }
    delete(v4, -740)
    switch v2 {
    case "hello":
        var v9 float64 = 0.0
        _ = v9
    case "a":
        if false {
            v3 -= 9223372036854775807
            var v10 uint8 = 2
            _ = v10
        }
    case "foo bar":
        fmt.Println(v3, v3)
    }
    switch 32767 {
    case -2147483648:
        v11 := GenS1{F0: 32767, F1: v3}
        _ = v11
    }
    fmt.Println(v2)
    v12 := []int{127}
    v3 = (v3 % ((829) | 1))
    var v13 int = 369
    v14 := v4[v13]
    v15 := GenS1{F0: 32767, F1: v3}
    var v16 int64 = 0
    if (false && true) {
        _ = len(v12)
    } else {
        var v17 float32 = 0.0
        _ = v17
    }
    fmt.Println(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
