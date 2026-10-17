package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint8) int64 {
    var v5 int = 1
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 489
    var v3 int32 = 930
    var v4 int32 = 785
    v2 = -2147483648
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v2)
        v4 += v4
        _ = fi7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = 42
        for fi10 := 0; fi10 < 5; fi10++ {
            if (42 == 163) {
                switch "" {
                case "foo bar":
                    fmt.Println(v3)
                }
                switch v2 {
                case 32767:
                    v11 := []int{(^(fi8))}
                    _ = v11
                case 9223372036854775807:
                    v2 += fi10
                }
                v3 = (v4 & 1)
            }
            if (!(!(434 != 714))) {
                if false {
                    v9 -= v9
                    v12 := v1[fi8]
                    _ = v12
                }
            } else {
                v13 := []int{fi8}
                var v14 int = v2
                v3 = (v3 - 425)
                fmt.Println(v14, v9)
                _ = v13
                _ = v14
            }
            if (false || ("foo bar" != "foo bar")) {
                if (!true) {
                    v4 += v3
                    if false {
                        v15 := GenS1{F0: (!true), F1: 38}
                        v9 -= v9
                        _ = v15
                    }
                    var v16 int = (fi8 - (fi10 - (fi10 % ((fi8) | 1))))
                    _ = v16
                }
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    if false {
        var v17 int8 = 127
        if false {
            v17 -= v17
        } else {
            v2 = (v2 ^ (v2 - v2))
            fmt.Println(v17)
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: ((!(2147483648 <= 65535)) || (1 == 32767)), F1: 65535}
                v3 = -2147483648
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    }
    v20 := []int{(v2 * (v2 ^ v2)), -715}
    v4 = v3
    if (!(!(328 > 65535))) {
        fmt.Println(v4)
        v21 := GenS2{F0: 760, F1: 42}
        _ = v21
    } else {
        switch (v2 >> ((255) & 63)) {
        case 128:
            var v22 int8 = -55
            _ = v22
        case 32767:
            var v23 int8 = 112
            _ = v23
        case 42:
            for fi24 := 0; fi24 < 3; fi24++ {
                switch -2147483648 {
                case -663:
                    v2 -= (^(fi24))
                case 133:
                    var v25 float64 = 1.5
                    _ = v25
                default:
                    v3 = -1
                }
                _ = fi24
            }
        default:
            var v26 uint16 = 0
            _ = v26
        }
    }
    var v27 uint8 = 0
    if false {
        v27 = (v27 >> ((255) & 63))
        v28 := GenS2{F0: 9223372036854775807, F1: 65535}
        if (float64((((1.5 - float32(3.14159)) / -2.25) / 1e10)) != (3.14159 * 1.5)) {
            _ = len(v20)
            v27 = v27
            _ = len(v1)
        }
        _ = v28
    } else {
        _ = len(v20)
        v4 += 151
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
