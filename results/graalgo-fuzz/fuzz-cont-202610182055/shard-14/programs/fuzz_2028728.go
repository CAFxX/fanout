package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 int8, p1 int64, p2 int) uint32 {
    v5 := []int{127}
    var v6 uint8 = 172
    var v7 int16 = 202
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 uint16 = 42
    var v4 uint32 = 2147483648
    if false {
        v8 := GenS1{F0: 42, F1: genFn0(42, 771, 42), F2: 4294967295}
        var v9 float64 = ((0.1 / 1.5) + (0.1 + ((1e10 / -2.25) - (3.14159 * (1.5 - 0.0)))))
        _ = v8
        _ = v9
    } else {
        if (false || ((!(true && (2 < 42))) || (false && (false && false)))) {
            v10 := GenS2{F0: true, F1: 42}
            var v11 float64 = 0.1
            _ = v10
            _ = v11
        }
        v12 := GenS3{F0: 42, F1: 100, F2: 48}
        _ = v12
    }
    var v13 uint16 = (v3 ^ v3)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        if (2147483647 < 312) {
            for fi14 := 0; fi14 < 12; fi14++ {
                if false {
                    var v15 int16 = 0
                    if (-2147483648 <= 1) {
                        fmt.Println(v13, v3, v13)
                        if ((!false) || false) {
                            v16 := GenS3{F0: -128, F1: 1, F2: (fi14 * -2147483648)}
                            _ = v16
                        } else {
                            var v17 bool = (255 >= 255)
                            _ = v17
                        }
                    }
                    _ = len(v2)
                    _ = v15
                }
                _ = fi14
            }
            v3 -= v13
        } else {
            if false {
                v18 := GenS3{F0: 42, F1: 1, F2: 9223372036854775807}
                _ = v18
            }
            v4 += (v4 ^ (v4 - 42))
            v19 := []int{-9223372036854775808, 32767}
            _ = v19
        }
    case "a":
        v13 -= uint16(v4)
    }
    v20 := GenS1{F0: int8(v13), F1: v4, F2: 0}
    if (!((65535 < 551) || (4294967295 >= genFn0(75, int64(v4), -527)))) {
        var v21 int = -2147483648
        _ = v21
    }
    var v22 uint16 = (v13 >> ((916) & 63))
    v13 += (v3 << ((0) & 63))
    v23 := GenS2{F0: true, F1: -9223372036854775808}
    var v24 int32 = 197
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
