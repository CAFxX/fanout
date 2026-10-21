package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    var v3 float64 = 0.0
    v4 := []int{17}
    var v5 uint16 = 42
    if true {
        _ = len(v1)
        v3 += v3
    } else {
        var v6 uint8 = 42
        v7 := GenS1{F0: ("a" + "\n\t"), F1: -2147483648, F2: 650}
        var v8 uint8 = (v6 | v6)
        _ = v6
        _ = v7
        _ = v8
    }
    v5 = 1
    _ = len(v4)
    if ((!true) || false) {
        var v9 uint8 = 172
        var v10 uint = 1
        var v11 uint = 4294967295
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 int = -409
        _ = v12
    }
    v13 := GenS1{F0: "", F1: 1, F2: 0}
    v14 := GenS1{F0: "", F1: -1, F2: 0}
    v2 += 1.5
    v2 += (1.5 / (3.14159 - ((-2.25 - float64(3.14159)) + (1.5 - ((v3 / v2) - v2)))))
    if (0 == 79) {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint32 = 65535
            var v17 float64 = v3
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v5 = (v5 ^ v5)
    for fi18 := 0; fi18 < 9; fi18++ {
        if len(v4) > 0 {
            v4[0] = fi18
        }
        fmt.Println(fi18)
        _ = fi18
    }
    switch 4294967295 {
    case -118:
        switch "foo bar" {
        case "":
            v4 = append(v4, 4294967295)
        default:
            fmt.Println(v3, v3, v2)
        }
    default:
        var v19 bool = (-1 > -185)
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
