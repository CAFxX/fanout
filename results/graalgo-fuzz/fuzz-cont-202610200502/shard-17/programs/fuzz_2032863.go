package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int64; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32) int16 {
    v6 := []int{32768, 0, -86}
    p0 += (p0 >> ((65) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 := GenS1{F0: 4294967295, F1: false, F2: 451}
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    return 38
}
func genFn1(p0 int32, p1 int32) int64 {
    v8 := []int{955, 32767, 9223372036854775807}
    v9 := GenS2{F0: 32767, F1: 9223372036854775807, F2: p1}
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "\n\t"
    var v3 float32 = (0.1 + (0.1 - ((3.14159 * 1.5) * (3.14159 + 0.1))))
    v4 := []int{860, 1, 42}
    var v5 uint64 = 42
    v3 = v3
    var v10 uint8 = 171
    v11 := GenS1{F0: uint(v5), F1: (0.0 != -2.25), F2: -695}
    switch 32768 {
    case 32768:
        var v12 float64 = ((float64(v3) - 1e10) - (0.1 + 0.0))
        _ = v12
    }
    v2 = v2
    switch v2 {
    case "\n\t":
        var v13 int = 128
        _ = v13
    }
    v3 += v3
    fmt.Println(v10)
    if len(v4) > 0 {
        v4[0] = 127
    }
    fmt.Println(v3)
    v1 += v1
    switch "\n\t" {
    case "":
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint = 255
            if true {
                var v16 int = (fi14 << (((v15 + v15)) & 63))
                v17 := GenS3{F0: 483, F1: v1}
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: 32767, F1: (^(v1)), F2: -1}
            _ = fi14
            _ = v15
            _ = v18
        }
    case "hello":
        v2 = (v2 + v2)
    default:
        v2 = "a"
    }
    v19 := GenS2{F0: 1, F1: v1, F2: -2147483648}
    fmt.Println(v1)
    v3 = 0.1
    v4 = append(v4, -2147483648)
    v20 := GenS2{F0: 128, F1: v1, F2: 598}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
