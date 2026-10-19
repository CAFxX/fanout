package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }

func main() {
    var v1 uint16 = 271
    var v2 string = ("hello" + "a")
    var v3 bool = true
    v4 := []int{32768, 174, 0, 127}
    fmt.Println(v3, v1)
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (((289 > 2) || v3) && v3) {
        fmt.Println(v2)
        if len(v4) > 0 {
            v4[0] = -1
        }
        v2 = v2
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v5 float64 = (0.1 / ((0.1 / 3.14159) * 1.5))
        var v6 uint32 = 0
        _ = v5
        _ = v6
    }
    v1 = 2
    fmt.Println(v2, v3)
    v2 = v2
    v1 -= v1
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 = ((v2 + v2) + v2)
        v8 := GenS1{F0: (fi7 * (fi7 * (fi7 + fi7))), F1: (v2 + v2), F2: -2147483648}
        v1 = (v1 + (v1 + v1))
        _ = fi7
        _ = v8
    }
    v1 += (v1 ^ (v1 - (v1 & (v1 << ((2) & 63)))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
