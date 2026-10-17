package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0(p0 string, p1 uint) uint64 {
    p1 += p1
    fmt.Println(p0)
    var v6 float64 = (0.1 + (0.1 / 0.1))
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 1, F1: (p0 + p0)}
        var v9 uint64 = 940
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 0.0
    v3 := make(map[int]string)
    v4 := []string{"\n\t", ("foo bar" + "")}
    v5 := make(map[int]uint32)
    if true {
        var v10 int64 = 42
        fmt.Println(v2, v10, v1)
        var v11 int8 = -128
        _ = v10
        _ = v11
    }
    var v12 int = 2147483647
    v12 = (-(v12))
    v2 = 3.14159
    var v13 uint32 = 540
    var v14 float32 = 0.1
    if ((((v2 - 0.0) < float64(0.1)) || false) && true) {
        switch (v12 | (v12 & (v12 + 0))) {
        case 32768:
            if (false && (!((255 <= 255) || false))) {
                _ = len(v4)
            } else {
                fmt.Println(v12, v12)
                v15 := GenS1{F0: 641, F1: "hello"}
                _ = v15
            }
        case -654:
            v16 := v5[32768]
            _ = v16
        case -289:
            var v17 int8 = -99
            _ = v17
        }
        var v18 int16 = 42
        _ = v18
    } else {
        var v19 uint8 = 128
        v12 -= -770
        if ((v13 & v13) == (v13 - v13)) {
            v2 = (v2 / -2.25)
            v13 += 1
        } else {
            var v20 int32 = 888
            var v21 uint64 = 583
            _ = v20
            _ = v21
        }
        _ = v19
    }
    v1 = 2
    v22 := GenS1{F0: 4294967295, F1: "\n\t"}
    fmt.Println(v2)
    if false {
        v23 := GenS1{F0: 0, F1: ("a" + "\n\t")}
        _ = v23
    } else {
        v12 = (^(v12))
        for fi24 := 0; fi24 < 10; fi24++ {
            switch (fi24 % (((-(fi24))) | 1)) {
            case 419:
                var v25 int = (fi24 >> ((1) & 63))
                _ = v25
            default:
                var v26 int8 = -11
                _ = v26
            }
            v13 = 2147483648
            _ = fi24
        }
    }
    var v27 float64 = ((1e10 - v2) / (((v2 + (0.1 - float64(1e10))) / 1.5) / v2))
    v2 += (float64(v14) + ((v27 * v2) * (v27 - v27)))
    v14 = v14
    v12 = v12
    v28 := GenS2{F0: 966, F1: (v1 + v1), F2: -21}
    v27 = v27
    if (!(188 <= 1)) {
        v29 := GenS2{F0: 2, F1: 32767, F2: 42}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
