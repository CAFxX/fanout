package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 450
    var v3 int16 = 420
    v4 := make(map[uint32]bool)
    fmt.Println(v3)
    v3 += 63
    var v5 int8 = 100
    v5 = 127
    v5 += v5
    v6 := GenS1{F0: (v3 < v3), F1: (v5 << (((^(v2))) & 63)), F2: 128}
    if true {
        v4[695] = ((!false) || true)
    }
    v7 := []int{4294967295, 4294967295, 127, 128}
    v5 -= (v5 * int8(v3))
    var v8 uint = 255
    v9 := GenS2{F0: ("foo bar" + ("" + "hello")), F1: 255, F2: 2}
    v8 += (v8 - (^(v2)))
    v8 += 42
    v8 = (v8 & v8)
    var v10 float64 = (1.5 * (((1.5 * 3.14159) / 0.1) + 3.14159))
    if false {
        if (128 >= -9223372036854775808) {
            fmt.Println(v2)
            fmt.Println(v3, v8, v3)
            if true {
                fmt.Println(v5)
            } else {
                v11 := GenS1{F0: (v8 <= 255), F1: (v5 % ((127) | 1)), F2: 128}
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 bool = true
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 uint32 = 65535
                    var v16 float64 = (-2.25 * v10)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
        }
        switch 23 {
        case 0:
            fmt.Println(v2)
        case -1:
            v8 += v2
        case 243:
            var v17 int32 = 42
            _ = v17
        }
        var v18 bool = true
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
