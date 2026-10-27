package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (-2.25 / (3.14159 + 0.1))
    var v3 float64 = float64(v2)
    v4 := []uint32{235, 2147483647}
    v5 := []int{127}
    var v6 uint32 = 0
    var v7 uint64 = 42
    v8 := GenS1{F0: -2147483648}
    v3 -= v3
    if (!true) {
        v9 := GenS2{F0: 127, F1: -128, F2: (v6 | v6)}
        _ = v9
    } else {
        v6 = 2147483648
    }
    var v10 int32 = 0
    switch 4294967295 {
    case 127:
        if (false && (("\n\t" + "a") != "")) {
            if (v6 >= 65535) {
                v11 := GenS1{F0: 42}
                var v12 uint16 = 65535
                var v13 int = 42
                _ = v11
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v1["foo bar"] = ("foo bar" <= "a")
                    v15 := GenS1{F0: (v10 ^ (^(v10)))}
                    _ = fi14
                    _ = v15
                }
                v4 = append(v4, (^(v6)))
                v16 := GenS3{F0: -20, F1: 750, F2: uint8(v6)}
                _ = v16
            }
            var v17 int8 = 0
            _ = v17
        } else {
            v6 -= v6
        }
    case -9223372036854775808:
        v18 := GenS3{F0: 55, F1: 1000, F2: 127}
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v2 -= v2
        fmt.Println(v7, v7, v10)
        _ = fi19
    }
    v20 := v1["foo bar"]
    var v21 uint8 = 0
    var v22 int32 = (-(v10))
    var v23 uint64 = v7
    fmt.Println(v2)
    var v24 bool = (int8(v21) <= 25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
