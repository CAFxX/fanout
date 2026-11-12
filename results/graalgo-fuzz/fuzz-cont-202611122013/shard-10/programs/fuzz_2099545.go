package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 int64, p1 int8, p2 uint8) uint16 {
    var v7 uint64 = 18446744073709551615
    p2 = p2
    if ((!false) || (214 >= 1)) {
        fmt.Println(p0)
    }
    p0 += -9223372036854775808
    v8 := []int{520, 128, 42, int(v7)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1() bool {
    v9 := []int{-9223372036854775808, 42}
    for fi10 := 0; fi10 < 1; fi10++ {
        if (false || (646 != 1)) {
            var v11 bool = (32767 <= 32767)
            _ = v11
        }
        if len(v9) > 0 {
            v9[0] = fi10
        }
        _ = len(v9)
        _ = fi10
    }
    var v12 float32 = 0.0
    if false {
        v12 = v12
        if (128 <= -465) {
            switch (("\n\t" + ("foo bar" + (("" + "hello") + ""))) + ("a" + ("foo bar" + "\n\t"))) {
            case "hello":
                v9 = append(v9, 4294967295)
            case "\n\t":
                _ = len(v9)
            }
            var v13 uint64 = 0
            v12 = -2.25
            _ = v13
        }
    }
    if (85 < 0) {
        var v14 bool = (3.14159 >= 0.0)
        _ = v14
    }
    _ = v12
    _ = v9
    return true
}

func main() {
    var v1 bool = true
    var v2 float64 = 1e10
    var v3 int32 = -2147483648
    var v4 float64 = (v2 - v2)
    v5 := make(map[int]int)
    var v6 float32 = ((1e10 + (float32(v2) * 0.1)) / 0.1)
    var v15 int8 = -99
    for fi16 := 0; fi16 < 18; fi16++ {
        switch (("" + "\n\t") + "foo bar") {
        case "":
            var v17 float64 = (v2 + (v2 - (0.1 * v4)))
            _ = v17
        case "\n\t":
            fmt.Println(v2, v4, v2)
        case "hello":
            v18 := GenS1{F0: 32767, F1: 0}
            _ = v18
        default:
            var v19 uint32 = 593
            _ = v19
        }
        var v20 bool = (v3 >= (v3 ^ (v3 ^ v3)))
        _ = fi16
        _ = v20
    }
    if ((!v1) && true) {
        fmt.Println(v2)
    } else {
        v4 = float64(v6)
        var v21 int = 222
        fmt.Println(v21)
        _ = v21
    }
    if true {
        v2 += (((v2 + (v2 * (1.5 + v2))) * v2) - 3.14159)
        v4 -= (((1e10 * v4) / (v2 - ((v4 + (0.0 * float64(1.5))) * (-2.25 * v2)))) * v2)
        v4 = float64((v6 - 1e10))
    }
    v4 += (1.5 - v4)
    v15 += (v15 % (((v15 << ((18446744073709551615) & 63))) | 1))
    var v22 int64 = -1
    v23 := GenS2{F0: genFn0((v22 << ((867) & 63)), 127, 210), F1: 255}
    fmt.Println(v3)
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := []int{-738, (fi24 << ((837) & 63)), fi24}
        var v26 float64 = (v4 + 0.1)
        v22 = v22
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
