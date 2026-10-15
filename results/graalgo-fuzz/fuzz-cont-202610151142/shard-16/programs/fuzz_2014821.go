package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
func genFn0(p0 int64, p1 uint64, p2 uint16) int {
    v5 := []int{-200, 436, -179}
    v6 := GenS1{F0: (p0 - (p0 & 417))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -793
}
func genFn1(p0 int32, p1 int8) uint16 {
    if false {
        var v7 uint8 = 127
        if ((1.5 + float32(1e10)) != -2.25) {
            var v8 int16 = 1000
            _ = v8
        }
        var v9 uint16 = 624
        _ = v7
        _ = v9
    } else {
        fmt.Println(p0)
        var v10 uint32 = 186
        _ = v10
    }
    v11 := GenS1{F0: -631}
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 7}
        _ = fi12
        _ = v13
    }
    v14 := []int{32768, 0, -44}
    p0 = (p0 | (p0 | p0))
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return 720
}

func main() {
    var v1 int = 297
    var v2 float64 = (((float64(0.1) / -2.25) / 1.0) * 0.0)
    var v3 float64 = (0.1 - ((float64(-2.25) * v2) / v2))
    var v4 float32 = 0.1
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint8 = 42
        _ = fi15
        _ = v16
    }
    switch ("" + "") {
    case "\n\t":
        for fi17 := 0; fi17 < 4; fi17++ {
            if (fi17 <= 4294967295) {
                var v18 uint32 = 4294967295
                var v19 float64 = v3
                _ = v18
                _ = v19
            } else {
                if (true || (false && false)) {
                    switch genFn0(-1, 18446744073709551615, 0) {
                    case -1:
                        switch (v1 >> ((18446744073709551615) & 63)) {
                        case 963:
                            v2 -= float64(float32(v3))
                        }
                    default:
                        v2 -= (0.0 / (((v3 + -2.25) + ((v3 / (1e10 / (((0.1 / v2) / v3) * 3.14159))) - (((v3 + v3) / v3) - (v3 * v2)))) + ((v2 - 1e10) * v2)))
                    }
                }
                v20 := []int{fi17, genFn0(-774, 931, 255), (fi17 % (((fi17 >> ((18446744073709551615) & 63))) | 1)), (v1 + fi17)}
                _ = v20
            }
            switch (("foo bar" + "a") + "hello") {
            case "a":
                v2 += float64(float32(((v3 - 3.14159) + v3)))
            default:
                v21 := []int{}
                _ = v21
            }
            _ = fi17
        }
    }
    v22 := []int{genFn0(9223372036854775807, 2, 255), v1}
    v1 += (v1 << ((2) & 63))
    v23 := GenS3{F0: (^(v1)), F1: -32768}
    v24 := GenS4{F0: 42, F1: 0, F2: 42}
    fmt.Println(v2, v2)
    var v25 bool = (!(v1 >= -1))
    _ = len(v22)
    fmt.Println(v25, v25, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
