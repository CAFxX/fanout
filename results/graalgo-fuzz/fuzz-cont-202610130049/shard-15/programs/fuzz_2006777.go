package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 int16, p2 string) bool {
    v6 := GenS1{F0: "\n\t", F1: -250}
    v7 := GenS2{F0: p0, F1: 864}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 float64 = 1.5
        v10 := GenS3{F0: 2, F1: (p1 >> (((^(p0))) & 63)), F2: 1}
        v11 := GenS3{F0: uint16(p0), F1: 1000, F2: 2}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        p1 -= (-(p1))
        for fi13 := 0; fi13 < 9; fi13++ {
            p2 = "a"
            var v14 uint8 = 255
            var v15 string = (p2 + p2)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    if (!(42 > 0)) {
        p0 = (p0 + p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ((!(p2 < "\n\t")) && false)
}

func main() {
    var v1 uint = 743
    var v2 int = 896
    var v3 uint16 = 0
    var v4 uint16 = v3
    var v5 string = ""
    v1 += 65535
    v16 := GenS1{F0: v5, F1: (v2 + 32767)}
    var v17 float32 = ((0.0 + (0.0 - float32(-2.25))) + ((-2.25 * 0.1) * (float32(float64(3.14159)) - 3.14159)))
    v2 -= 32767
    v3 -= v4
    v1 += (v1 << ((v1) & 63))
    fmt.Println(v2, v17, v5)
    if (v2 < v2) {
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(v5)
                v2 = fi18
                _ = fi19
            }
            var v20 int16 = 1
            _ = fi18
            _ = v20
        }
    } else {
        v21 := GenS4{F0: -48}
        var v22 int64 = 1
        _ = v21
        _ = v22
    }
    v23 := []int{int(v4), (v2 >> ((1) & 63)), (v2 % ((1) | 1)), -2147483648}
    switch (v2 | (v2 ^ (v2 << (((v1 + 883)) & 63)))) {
    case 734:
        fmt.Println(v2)
    case -2147483648:
        fmt.Println(v4)
    case 128:
        for fi24 := 0; fi24 < 19; fi24++ {
            if true {
                for fi25 := 0; fi25 < 17; fi25++ {
                    v2 = fi24
                    v17 -= (float32(0.0) - 1.5)
                    _ = fi25
                }
            }
            _ = fi24
        }
    default:
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 float32 = -2.25
            v28 := GenS1{F0: v5, F1: 127}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v29 int16 = -1
    if (-2147483648 == -266) {
        v1 = 42
    } else {
        for fi30 := 0; fi30 < 4; fi30++ {
            v1 -= (v1 & v1)
            var v31 uint = 18446744073709551615
            _ = fi30
            _ = v31
        }
        fmt.Println(v17, v3)
    }
    v32 := GenS3{F0: (^(v4)), F1: -880, F2: 128}
    var v33 int32 = 962
    v34 := GenS2{F0: v1, F1: 9223372036854775807}
    var v35 bool = (v29 < v29)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
