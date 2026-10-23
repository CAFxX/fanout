package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16) uint64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 128
        _ = fi6
        _ = v7
    }
    p0 -= (p0 | p0)
    var v8 string = "foo bar"
    _ = p0
    _ = v8
    return 467
}

func main() {
    var v1 uint16 = 28
    v2 := make(map[int]int)
    v3 := []uint32{0, 1, 2147483647}
    v4 := make(map[string]bool)
    var v5 uint8 = 185
    v1 = v1
    if false {
        v9 := GenS1{F0: 2147483647, F1: -873, F2: (true && false)}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 1
            v1 = v1
            v12 := GenS1{F0: 0, F1: fi10, F2: (!true)}
            _ = fi10
            _ = v11
            _ = v12
        }
        v5 += 247
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if (("foo bar" < "foo bar") || (42 <= 973)) {
            var v14 uint = 255
            v15 := GenS2{F0: fi13, F1: -2147483648}
            if false {
                var v16 int = (-(fi13))
                _ = v16
            } else {
                var v17 float64 = 1e10
                for fi18 := 0; fi18 < 17; fi18++ {
                    v14 -= (v14 / ((v14) | 1))
                    v19 := GenS2{F0: (^(fi13)), F1: 42}
                    var v20 uint32 = 65535
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v17 += v17
                _ = v17
            }
            _ = v14
            _ = v15
        }
        if false {
            if ((1000 == 1000) || false) {
                if (0 != -1) {
                    var v21 uint8 = (v5 | 255)
                    var v22 bool = true
                    _ = v21
                    _ = v22
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v3 = append(v3, 14)
            v3 = append(v3, 189)
        } else {
            var v23 int = 1
            _ = v23
        }
        _ = fi13
    }
    v5 -= v5
    for fi24 := 0; fi24 < 9; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            v1 -= (^(v1))
            switch "" {
            case "a":
                fmt.Println(v5, fi25, fi25)
            case "foo bar":
                v1 += (v1 * (v1 % ((v1) | 1)))
                fmt.Println(v5)
            case "":
                switch 1 {
                case 127:
                    v26 := GenS3{F0: 42, F1: (v1 << ((1) & 63)), F2: -332}
                    _ = v26
                case 32767:
                    var v27 int64 = -739
                    _ = v27
                }
            default:
                v1 = v1
            }
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println(v5, v1)
    v28 := GenS4{F0: -1, F1: 42}
    if (612 > -483) {
        var v29 bool = ((!false) && (true || false))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
