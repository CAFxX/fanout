package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{-9223372036854775808, 513}
    v3 := []int{128, 1}
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: -963, F1: 2, F2: 127}
    var v6 int16 = -967
    switch "" {
    case "hello":
        v7 := GenS2{F0: true, F1: 127}
        _ = v7
    }
    v6 = -1
    delete(v1, "hello")
    switch 42 {
    case -2147483648:
        var v8 float64 = (1.5 / 1e10)
        _ = v8
    case 32768:
        v6 -= (v6 * (v6 + -1))
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    var v9 float32 = 1.5
    var v10 float64 = float64((float32(1e10) - v9))
    v9 += float32(3.14159)
    v10 = 1.5
    v11 := GenS1{F0: -275, F1: 0, F2: 42}
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            v9 += v9
        } else {
            var v13 float32 = ((v9 + v9) / v9)
            v9 -= v13
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 uint16 = 32767
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        if false {
            if ((!true) && false) {
                fmt.Println(v9)
                v16 := GenS2{F0: ((v9 - ((v9 + 0.0) / (((v9 - 0.1) + v9) + float32((v10 - float64(-2.25)))))) == 0.1), F1: -128}
                _ = v16
            }
        }
        _ = len(v3)
        _ = fi12
    }
    var v17 string = ("hello" + ("foo bar" + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
