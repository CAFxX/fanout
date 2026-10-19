package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 int8, p1 int, p2 uint32) uint32 {
    var v7 int8 = 1
    p2 -= (p2 * p2)
    v7 = (v7 & p0)
    fmt.Println(p1, p1)
    var v8 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 << ((1) & 63))
}
func genFn1(p0 uint16, p1 uint8, p2 uint8) uint16 {
    fmt.Println(p1)
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 563
    v3 := []uint32{490, v1, (v1 ^ v1), (v1 * (v1 & v1))}
    var v4 string = "foo bar"
    v5 := []string{}
    var v6 float64 = 1e10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: 0, F1: (!(-471 > 32768)), F2: -108}
    var v10 float32 = (float32(3.14159) * ((float32(float64(0.1)) / float32((v6 + v6))) + (((-2.25 / (float32((-2.25 * v6)) / (1.5 / 1e10))) + (((1.5 * 1e10) * 1e10) / 1.0)) / 0.1)))
    v11 := GenS2{F0: (1 == 532), F1: (v2 ^ 2)}
    for fi12 := 0; fi12 < 11; fi12++ {
        v1 -= v1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v13 := GenS2{F0: false, F1: v2}
        _ = fi12
        _ = v13
    }
    v4 = "hello"
    v6 = (1.5 * (v6 - -2.25))
    v3 = append(v3, genFn0(-118, -422, v1))
    if (1 <= 9223372036854775807) {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 uint8 = uint8(v2)
            _ = fi14
            _ = v15
        }
    } else {
        if (true && false) {
            var v16 int64 = 9223372036854775807
            _ = v16
        }
        v1 = v1
        for fi17 := 0; fi17 < 12; fi17++ {
            if (false || (-1 == -1)) {
                v2 += v2
            }
            v10 += (v10 + (3.14159 - (v10 + float32(v6))))
            v18 := GenS2{F0: ((0 > 201) && (0 > 0)), F1: v2}
            _ = fi17
            _ = v18
        }
    }
    switch -566 {
    case 42:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                v4 = (v4 + ("" + v4))
                v10 -= (float32(v6) - (-2.25 - v10))
                var v21 int32 = 1
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: 255, F1: true, F2: -128}
            _ = fi19
            _ = v22
        }
    case 127:
        var v23 int32 = 2147483647
        _ = v23
    case 9223372036854775807:
        switch 128 {
        case 32768:
            v24 := GenS1{F0: 175, F1: false, F2: 127}
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
