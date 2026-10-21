package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int8 = 100
    fmt.Println(v3)
    var v4 int = -125
    v5 := GenS1{F0: -1, F1: 9223372036854775807, F2: false}
    if true {
        v6 := GenS2{F0: (("foo bar" + "foo bar") + "a")}
        for fi7 := 0; fi7 < 8; fi7++ {
            v4 -= (fi7 | (-(fi7)))
            v8 := GenS2{F0: "hello"}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    switch v4 {
    case -272:
        fmt.Println(v3, v3)
    case 959:
        fmt.Println(v4, v3)
    case -2147483648:
        v4 = -920
    }
    v9 := []int{}
    fmt.Println(v3, v3)
    v3 = -74
    if ((!(!false)) && (false || (971 == 2))) {
        var v10 uint32 = 2147483647
        v1[(((("hello" + "\n\t") + ("" + ("a" + "a"))) + ("" + "foo bar")) + ("a" + "hello"))] = v4
        _ = v10
    }
    v11 := GenS1{F0: 545, F1: v4, F2: false}
    var v12 int8 = 70
    var v13 bool = (true && (-878 < -9223372036854775808))
    var v14 bool = ((!v13) || true)
    if (2 == 65535) {
        for fi15 := 0; fi15 < 12; fi15++ {
            v2[(v4 - v4)] = 42
            if true {
                v12 -= (v3 % ((v3) | 1))
                var v16 uint16 = 65535
                var v17 string = ""
                _ = v16
                _ = v17
            }
            switch "\n\t" {
            case "a":
                v14 = (float64((3.14159 * (1e10 * (0.0 + 1e10)))) != 1e10)
            case "foo bar":
                v4 = (fi15 & (fi15 - fi15))
            }
            _ = fi15
        }
        var v18 string = "hello"
        fmt.Println(v4, v18)
        _ = v18
    } else {
        var v19 uint32 = 1
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: -325, F1: -9223372036854775808, F2: ((^(v19)) >= v19)}
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
