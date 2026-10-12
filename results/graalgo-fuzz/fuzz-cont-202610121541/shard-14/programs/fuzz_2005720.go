package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 int64, p1 int16, p2 int8) string {
    var v5 uint = 65535
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}

func main() {
    var v1 float64 = (0.1 / ((3.14159 * 1.5) / 1.0))
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint8 = 239
    if false {
        switch 2147483647 {
        case -245:
            fmt.Println(v1, v4)
        case 1:
            delete(v2, 690)
        case -9223372036854775808:
            v1 = v1
        }
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 int16 = 42
            _ = fi6
            _ = v7
        }
        if (1000 != 311) {
            fmt.Println(v1, v4)
        } else {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := []int{(^(fi8))}
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 uint16 = 2
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            v12 := []int{-590, -1}
            v4 += 128
            _ = v12
        }
    } else {
        v1 += (v1 / ((((v1 + (((((float64(1e10) - (v1 * 3.14159)) * float64(1.5)) / 1e10) / ((-2.25 * v1) - (v1 * float64(3.14159)))) / 1.0)) - (float64(0.1) + v1)) - v1) - 1.5))
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v13 int = -9223372036854775808
            _ = v13
        default:
            var v14 float32 = (3.14159 * (0.0 * float32(3.14159)))
            _ = v14
        }
        v1 = 0.0
    }
    var v15 int32 = -820
    var v16 int64 = -1
    var v17 bool = false
    v15 -= (^(v15))
    var v18 int32 = 1
    if v17 {
        if ((float32(v1) - 1.5) >= (((3.14159 * 3.14159) / float32(v1)) / 1.5)) {
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := []int{(fi19 / ((fi19) | 1)), (fi19 >> ((931) & 63)), fi19}
                var v21 int64 = (v16 - (v16 & v16))
                _ = fi19
                _ = v20
                _ = v21
            }
            v15 = -2147483648
        }
        v22 := []int{-2147483648, -2147483648, 2147483647}
        _ = v22
    } else {
        v16 = (v16 | v16)
    }
    fmt.Println(v16)
    v16 += 146
    v23 := GenS1{F0: v17, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
