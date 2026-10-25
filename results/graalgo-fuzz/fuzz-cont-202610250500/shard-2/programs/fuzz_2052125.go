package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8, p1 bool, p2 int8) uint {
    v5 := GenS1{F0: 392, F1: 0, F2: 1}
    var v6 float32 = 0.1
    var v7 int64 = 239
    p0 += (^(p0))
    v8 := []int{-9223372036854775808, 2147483647, 128, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 180
}
func genFn1(p0 bool, p1 uint16) uint {
    v9 := GenS1{F0: 0, F1: 100, F2: 1}
    var v10 int64 = 1
    var v11 uint32 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = (1.5 - -2.25)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 901
    fmt.Println(v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int16 = -1
        _ = fi12
        _ = v13
    }
    var v14 float32 = float32(3.14159)
    v14 = ((float32(((0.1 + (float64((3.14159 / 0.1)) / 1.5)) - ((1e10 + (0.1 / 1e10)) * ((((((1e10 - 0.1) + -2.25) + (3.14159 / 1.0)) - ((-2.25 / float64(((-2.25 - (v2 / 1e10)) / -2.25))) + 1.5)) + -2.25) + (0.0 * (float64(v2) * (0.0 - (0.0 * ((0.0 + ((((0.0 + 1.5) - 0.1) - 0.1) * (0.1 - 1e10))) + -2.25))))))))) * (0.0 * 0.1)) + float32(0.1))
    if false {
        var v15 string = "a"
        _ = v15
    } else {
        if (true || (false || false)) {
            fmt.Println(v1, v2)
        }
        v14 = v2
    }
    var v16 int = 649
    v17 := GenS2{F0: 11, F1: 39}
    var v18 bool = (true && false)
    delete(v3, 4294967295)
    switch "a" {
    case "\n\t":
        var v19 int8 = -49
        _ = v19
    }
    switch v16 {
    case -2147483648:
        var v20 uint8 = 2
        _ = v20
    default:
        fmt.Println(v16)
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 int8 = -1
        if (!v18) {
            if (-2147483648 >= v4) {
                var v23 uint16 = (v1 + 333)
                for fi24 := 0; fi24 < 11; fi24++ {
                    v14 = 0.0
                    v25 := []int{fi24, v16, 128}
                    v22 += 0
                    if (839 < -1) {
                        v16 = v16
                    }
                    v1 = 652
                    _ = fi24
                    _ = v25
                }
                v26 := GenS2{F0: 18446744073709551615, F1: (-(v22))}
                _ = v23
                _ = v26
            }
        }
        _ = fi21
        _ = v22
    }
    var v27 bool = v18
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
