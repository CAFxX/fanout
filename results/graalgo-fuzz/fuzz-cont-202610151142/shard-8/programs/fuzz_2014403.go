package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 int32; F1 int }

func main() {
    var v1 bool = true
    v2 := []uint32{65535}
    v3 := make(map[string]string)
    v4 := v3["a"]
    v4 = (v4 + ("a" + (v4 + (((v4 + "a") + "hello") + "hello"))))
    _ = len(v3)
    var v5 bool = (-128 >= -1)
    v6 := GenS1{F0: 18446744073709551615, F1: -97, F2: ("a" + ("a" + ("a" + ((v4 + "\n\t") + ""))))}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 string = (v4 + "\n\t")
            var v10 float32 = 3.14159
            _ = len(v3)
            v9 = v9
            _ = fi8
            _ = v10
            _ = v9
        }
        switch fi7 {
        case 4294967295:
            fmt.Println(v5, v5)
        case 42:
            v11 := GenS2{F0: 0, F1: (-(fi7))}
            _ = v11
        case 32768:
            switch 2147483647 {
            case 777:
                v4 = ("" + v4)
            case 32767:
                v5 = v1
            }
        }
        _ = fi7
    }
    if ((151 > 0) || ("" > "hello")) {
        fmt.Println(v1, v1)
    } else {
        v1 = (-2.25 == float64(float32((3.14159 / (0.1 / ((((3.14159 * ((1e10 * 3.14159) / 0.1)) / (0.0 - 1e10)) * 1.5) + -2.25))))))
        v12 := GenS2{F0: 42, F1: -2147483648}
        _ = v12
    }
    var v13 int8 = 42
    if false {
        v14 := GenS2{F0: -810, F1: 450}
        if true {
            v5 = v5
            v15 := GenS1{F0: 4294967295, F1: 100, F2: ((v4 + v4) + v4)}
            _ = v15
        } else {
            v2 = append(v2, 309)
        }
        _ = v14
    }
    v16 := GenS1{F0: 38, F1: 56, F2: "\n\t"}
    var v17 int = 32767
    v4 = (v4 + v4)
    v4 = "\n\t"
    v1 = false
    if ((1 == -691) && (v17 >= v17)) {
        switch v17 {
        case 0:
            v13 -= v13
        case 2147483647:
            v1 = true
        }
    } else {
        v18 := GenS1{F0: 65535, F1: (v13 << ((474) & 63)), F2: v4}
        _ = v18
    }
    v1 = (true || (0 == 2))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
