package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint64) int64 {
    switch "" {
    case "a":
        for fi4 := 0; fi4 < 10; fi4++ {
            fmt.Println(p0, p0)
            p0 = p0
            p0 = (p0 * (p0 * 2))
            _ = fi4
        }
    }
    p0 = p0
    p0 -= p0
    if (!((!((3.14159 - (1e10 * 1.5)) < (3.14159 * (-2.25 / (1.5 - (0.1 + float64(0.0))))))) || false)) {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 int8 = 100
            if true {
                v7 := GenS1{F0: (fi5 + (fi5 + fi5))}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        v8 := GenS1{F0: 42}
        _ = v8
    } else {
        var v9 uint16 = 1
        _ = v9
    }
    _ = p0
    return 915
}

func main() {
    var v1 float64 = (0.1 / ((-2.25 - 3.14159) / 3.14159))
    var v2 uint64 = 0
    v3 := []int{}
    v10 := GenS2{F0: 0, F1: 42}
    _ = len(v3)
    if (1.5 <= (1e10 - 1e10)) {
        switch 787 {
        case -1:
            if len(v3) > 0 {
                v3[0] = 1
            }
        case 104:
            v11 := GenS3{F0: 501, F1: 900}
            _ = v11
        case 42:
            if (-2147483648 > -2147483648) {
                if len(v3) > 0 {
                    v3[0] = 64
                }
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 bool = true
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v1 -= 1.5
                        fmt.Println(fi12)
                        _ = fi14
                    }
                    fmt.Println(v2, v1)
                    _ = fi12
                    _ = v13
                }
                v2 = v2
            }
        }
    }
    v2 -= (v2 | 1)
    v1 = v1
    if (-2147483648 < 0) {
        v15 := GenS4{F0: -1, F1: v2, F2: (v2 - v2)}
        v16 := GenS5{F0: 2147483648}
        v1 -= (v1 - (v1 + ((v1 * ((1.5 * (((v1 * (3.14159 * (1e10 - ((3.14159 + (float64(0.0) / v1)) + 1e10)))) + v1) + (v1 + v1))) - (v1 - (v1 + v1)))) * ((0.0 * (float64(float32(((v1 - (v1 * v1)) + 0.0))) - (((v1 / float64((((((1e10 + (0.1 * (3.14159 * 3.14159))) * ((float32(v1) * 0.0) / ((0.1 - (1.5 - (0.1 * -2.25))) / 3.14159))) / -2.25) + float32(0.0)) - 1e10))) - (3.14159 - (0.0 + float64(1.5)))) - v1))) + (v1 - (((0.1 - float64(1e10)) + 0.1) - (v1 - (v1 / float64(3.14159)))))))))
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 = 2
            _ = fi18
        }
        _ = fi17
    }
    var v19 uint16 = 235
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int32 = 42
        var v22 int16 = int16(v21)
        _ = fi20
        _ = v21
        _ = v22
    }
    v19 = (v19 / (((v19 << ((0) & 63))) | 1))
    v1 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
}
