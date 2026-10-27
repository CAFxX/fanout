package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 uint16, p2 int16) int8 {
    fmt.Println(p1)
    v7 := []int{32768}
    fmt.Println(p0, p1, p2)
    switch 128 {
    case 181:
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(p0, p2, p1)
            _ = fi8
        }
    case -163:
        p1 += 255
    case -84:
        p2 = (p2 << ((255) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -128
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 bool = (!(32767 != 529))
    var v5 uint32 = 517
    v6 := make(map[string]bool)
    v9 := GenS1{F0: -924, F1: genFn0(1, 2, -1), F2: genFn0(788, 655, -32768)}
    v10 := v3[127]
    if (!((-9223372036854775808 > 967) && v4)) {
        var v11 int = -2147483648
        v12 := v2[(v11 >> ((42) & 63))]
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: 0, F1: 1, F2: 28}
            if v4 {
                var v15 float64 = 1e10
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint16 = 2
                    switch v12 {
                    case "foo bar":
                        var v18 int64 = -9223372036854775808
                        _ = v18
                    default:
                        var v19 int32 = 0
                        _ = v19
                    }
                    var v20 int = (v11 >> ((42) & 63))
                    _ = fi16
                    _ = v17
                    _ = v20
                }
                _ = v15
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                _ = len(v2)
                _ = fi21
            }
            _ = fi13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    v4 = (v1 > v1)
    _ = len(v6)
    if v4 {
        var v22 int16 = -32768
        _ = v22
    } else {
        v5 += (^(v5))
        v23 := GenS2{F0: v1, F1: -1}
        _ = v23
    }
    v5 -= (^(v5))
    fmt.Println(v1, v1)
    v10 = (((("\n\t" + "hello") + (v10 + ((v10 + "\n\t") + ((("" + "") + v10) + "hello")))) + "a") + "foo bar")
    v24 := GenS3{F0: 125, F1: (v1 + (v1 ^ (v1 | (v1 ^ 2147483647))))}
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS4{F0: 1}
        var v27 uint32 = (v5 / ((65535) | 1))
        if v4 {
            switch ((v10 + "\n\t") + "") {
            case "a":
                var v28 int8 = -128
                _ = v28
            case "foo bar":
                var v29 string = ((v10 + v10) + (v10 + v10))
                _ = v29
            default:
                fmt.Println(v1, v5, v5)
            }
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
