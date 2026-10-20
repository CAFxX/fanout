package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int = 127
    v2 := []string{"a", "", "a", "hello"}
    var v3 float64 = 1e10
    v4 := make(map[string]bool)
    var v5 int64 = -1
    delete(v4, "\n\t")
    v3 += v3
    v3 = ((float64((((0.0 / 3.14159) - ((1.5 - 1.5) / 1.0)) - 0.0)) * ((v3 / (v3 * 0.0)) - float64(1e10))) + ((v3 / v3) - v3))
    v1 += (v1 % (((v1 | (v1 / ((v1) | 1)))) | 1))
    if (-1 > (v5 & v5)) {
        v3 -= (v3 + 1e10)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(fi6, v3)
            _ = len(v4)
            _ = fi7
        }
        _ = fi6
    }
    v8 := GenS1{F0: 0, F1: "hello"}
    for fi9 := 0; fi9 < 6; fi9++ {
        _ = len(v2)
        v10 := GenS1{F0: 2, F1: "\n\t"}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 1}
    var v12 bool = (false && false)
    v3 += ((v3 + (0.0 / v3)) + 1.5)
    if v12 {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint16 = 65535
            _ = fi13
            _ = v14
        }
    } else {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    v12 = (!(381 == 742))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
