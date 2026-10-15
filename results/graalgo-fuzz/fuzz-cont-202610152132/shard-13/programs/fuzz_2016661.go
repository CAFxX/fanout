package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 int16) uint8 {
    var v4 int16 = (p0 | p0)
    p0 = p0
    for fi5 := 0; fi5 < 4; fi5++ {
        switch 32767 {
        case -9223372036854775808:
            v4 = v4
        case 32768:
            for fi6 := 0; fi6 < 9; fi6++ {
                if false {
                    v4 = p0
                    v7 := GenS1{F0: true, F1: (v4 << ((4294967295) & 63))}
                    _ = v7
                }
                if (-2147483648 < -233) {
                    if (v4 <= p0) {
                        fmt.Println(v4, fi6)
                        fmt.Println(fi6)
                        var v8 uint8 = 39
                        _ = v8
                    }
                    v9 := GenS2{F0: 32767, F1: 42, F2: 42}
                    v10 := GenS3{F0: p0, F1: 18446744073709551615}
                    _ = v10
                    _ = v9
                } else {
                    var v11 string = ((((("foo bar" + "\n\t") + "hello") + "") + "") + "a")
                    var v12 int = fi6
                    var v13 float64 = 0.0
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi6
            }
        }
        v14 := GenS2{F0: 65535, F1: 720, F2: 100}
        _ = fi5
        _ = v14
    }
    v4 += (v4 ^ 1)
    v15 := []int{1, 2147483647, 9223372036854775807, 2147483647}
    _ = p0
    _ = v15
    _ = v4
    return 255
}
func genFn1(p0 uint8, p1 uint, p2 uint64) uint8 {
    p1 += uint(p0)
    v16 := GenS4{F0: 42, F1: 2147483647}
    var v17 float64 = 0.1
    fmt.Println(p0)
    var v18 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (p0 << ((255) & 63))
}

func main() {
    var v1 int8 = -124
    v2 := []string{}
    var v3 uint32 = 364
    _ = len(v2)
    v3 += v3
    v3 += (v3 >> ((2) & 63))
    var v19 float32 = 1e10
    fmt.Println(v19, v19, v19)
    v19 = float32(1e10)
    if true {
        var v20 int64 = -988
        fmt.Println(v3, v20, v3)
        var v21 uint32 = (v3 / ((v3) | 1))
        _ = v20
        _ = v21
    }
    var v22 int16 = 0
    v22 -= (^(v22))
    fmt.Println(v1)
    var v23 uint32 = 2
    var v24 uint8 = 7
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
