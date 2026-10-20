package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int64, p1 uint32) uint32 {
    if false {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 bool = (336 == 598)
            fmt.Println(p0)
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint8 = 1
            p0 = p0
            var v11 float64 = (-2.25 * (-2.25 * float64(1e10)))
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    if ((p1 - (^(p1))) <= 116) {
        var v12 uint32 = (p1 - p1)
        var v13 bool = (1000 > 120)
        p0 -= p0
        _ = v12
        _ = v13
    } else {
        fmt.Println(p0, p1)
        p1 = (^(p1))
    }
    p0 -= 42
    var v14 uint16 = 1
    for fi15 := 0; fi15 < 14; fi15++ {
        if (true && false) {
            v16 := []int{fi15, (fi15 % ((fi15) | 1)), 42, fi15}
            v17 := GenS1{F0: (1 > 102)}
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{763, 446, fi15}
                p0 += (p0 * (p0 + p0))
                _ = fi18
                _ = v19
            }
            var v20 int8 = 0
            _ = v20
        }
        v14 = 65535
        var v21 float32 = (0.1 + 3.14159)
        _ = fi15
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v14
    return 928
}
func genFn1(p0 string) int {
    var v22 uint = 376
    p0 = (p0 + (p0 + p0))
    _ = p0
    _ = v22
    return -2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    var v3 int64 = 68
    var v4 int = -112
    var v5 uint = 870
    var v6 bool = false
    var v23 uint8 = 171
    v24 := []int{v4, v4, v4, 9223372036854775807}
    _ = len(v1)
    v5 += v5
    if v6 {
        v5 -= v5
        _ = len(v1)
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 string = "a"
            _ = fi25
            _ = v26
        }
    }
    if true {
        v27 := GenS2{F0: (v23 & v23)}
        var v28 int8 = 42
        _ = v27
        _ = v28
    } else {
        var v29 int8 = -1
        if v6 {
            var v30 float32 = ((v2 * (((v2 / ((1.5 - (v2 + v2)) - (v2 / v2))) * 1e10) * (float32((((1e10 * -2.25) - 1.5) - (float64((-2.25 + (1.5 / 3.14159))) / ((0.1 * 1.5) / 3.14159)))) - (((float32(1e10) / (v2 * -2.25)) * (v2 / v2)) - ((v2 * v2) / (v2 + (v2 - 1.5))))))) * (v2 * (v2 + (((((v2 * v2) + (float32((1e10 + 0.0)) / v2)) / v2) * ((0.1 - v2) * v2)) + (v2 / -2.25)))))
            switch "" {
            case "a":
                v5 += 65535
            }
            _ = v30
        }
        _ = v29
    }
    if v6 {
        var v31 int32 = -2147483648
        _ = v31
    } else {
        var v32 float64 = 0.1
        _ = v32
    }
    v33 := GenS3{F0: 0, F1: 1, F2: 836}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
