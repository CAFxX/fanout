package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 bool }
func genFn0(p0 uint8, p1 int16, p2 uint8) uint8 {
    v7 := []int{405, 32767}
    v7 = append(v7, 127)
    if (42 <= 896) {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 2, F1: 0, F2: -32768}
            _ = fi8
            _ = v9
        }
        p1 += 32767
        v7 = append(v7, 1)
    } else {
        var v10 float32 = (1e10 - (3.14159 + 1.5))
        for fi11 := 0; fi11 < 14; fi11++ {
            v7 = append(v7, -9223372036854775808)
            v12 := GenS1{F0: 4294967295, F1: 76, F2: int16(p0)}
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if (!((42 == 516) || ((0.1 > (0.1 * float32((0.0 * 1.5)))) || (4294967295 > 42)))) {
            p2 += (p2 & p0)
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                if false {
                    p1 -= p1
                }
                var v15 int32 = 0
                _ = fi14
                _ = v15
            }
            v16 := GenS2{F0: p2}
            _ = len(v7)
            _ = v16
        }
        fmt.Println(p2, p2)
        _ = fi13
    }
    v17 := GenS1{F0: 42, F1: 127, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v7
    return p2
}

func main() {
    var v1 int = 32768
    var v2 uint64 = 4294967295
    var v3 uint = 65535
    var v4 int32 = 161
    v5 := make(map[string]string)
    var v6 uint32 = 247
    switch "\n\t" {
    case "foo bar":
        _ = len(v5)
    default:
        var v18 int = 32767
        _ = v18
    }
    v19 := GenS3{F0: "foo bar", F1: true, F2: false}
    v20 := v5["foo bar"]
    if false {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: (v2 % ((uint64(v3)) | 1)), F1: 100, F2: int16(v2)}
            if (563 <= (v2 ^ (v2 << (((v3 / (((v3 & v3)) | 1))) & 63)))) {
                if false {
                    fmt.Println(v1, v2, fi21)
                }
            }
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            switch 32767 {
            case -9223372036854775808:
                if ((v4 | v4) >= v4) {
                    v3 -= v3
                    var v24 float64 = 1.5
                    _ = v24
                }
            case -853:
                fmt.Println(v4, v6, v6)
            case 1:
                v25 := GenS1{F0: v2, F1: -122, F2: 649}
                _ = v25
            default:
                var v26 string = v20
                _ = v26
            }
            var v27 int64 = -9223372036854775808
            _ = fi23
            _ = v27
        }
    }
    v6 -= 65535
    v6 += v6
    fmt.Println(v2, v6, v2)
    var v28 float32 = -2.25
    var v29 float64 = (1.5 * (0.1 + (1e10 * float64(v28))))
    fmt.Println(v29, v20, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
