package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int }
func genFn0() int64 {
    if ("a" < (("\n\t" + "") + ("hello" + "foo bar"))) {
        v5 := []int{1, 128}
        _ = v5
    }
    if (3.14159 >= (1.5 / 3.14159)) {
        if (32767 == 255) {
            var v6 int64 = 144
            v6 = v6
            v6 = (v6 * v6)
            _ = v6
        }
        v7 := []int{}
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if (-32768 == -407) {
            v10 := GenS1{F0: 0, F1: 1000, F2: 1000}
            v11 := []int{(-(fi9)), (fi9 / ((-385) | 1))}
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    return 604
}

func main() {
    var v1 uint16 = 147
    var v2 float64 = 0.0
    var v3 float32 = ((float32(v2) * 1e10) + 1e10)
    var v4 float32 = 1.5
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := []int{fi12, fi12, (fi12 >> ((2) & 63))}
        var v14 uint = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 -= v4
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 127, F1: -32768, F2: -32768}
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v2)
    v3 = float32(v2)
    v17 := GenS1{F0: 0, F1: 124, F2: -32768}
    if false {
        var v18 uint = 22
        var v19 bool = true
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: 242, F1: -146}
        v2 += 0.0
        if (true && (531 > 4294967295)) {
            if false {
                var v21 bool = true
                v21 = (((-1 <= 0) || v21) || ((1.5 != (float32(v2) * ((v3 / v4) + (v4 / v3)))) || (18446744073709551615 >= 18446744073709551615)))
                _ = v21
            } else {
                v2 -= v2
                v22 := []int{878, 1, 117, 127}
                switch "hello" {
                case "foo bar":
                    v23 := GenS3{F0: -1}
                    _ = v23
                case "a":
                    if len(v22) > 0 {
                        v22[0] = 32768
                    }
                case "\n\t":
                    var v24 int = -2147483648
                    _ = v24
                default:
                    if len(v22) > 0 {
                        v22[0] = 1
                    }
                }
                _ = v22
            }
        } else {
            var v25 int16 = 1
            var v26 int16 = (v25 * v25)
            _ = v25
            _ = v26
        }
        _ = v20
    }
    var v27 int8 = -42
    v3 += 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
