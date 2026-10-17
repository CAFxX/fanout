package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int64) uint8 {
    fmt.Println(p0)
    p0 = p0
    var v5 bool = true
    v6 := GenS1{F0: 9223372036854775807, F1: 42, F2: (p0 / (((p0 % ((p0) | 1))) | 1))}
    var v7 uint8 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{"", "\n\t", "a", ("a" + "\n\t")}
    v2 := make(map[int]bool)
    var v3 uint64 = 4294967295
    var v4 bool = (922 == 32767)
    if (42 == 42) {
        v4 = v4
        v3 = v3
        v4 = (genFn0(9223372036854775807) < 2)
    } else {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "hello")
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint32 = 696
            var v10 uint16 = 331
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int8 = 0
        _ = v11
    }
    v3 += (v3 >> ((255) & 63))
    fmt.Println(v4)
    if false {
        if (!((v4 || true) && (!v4))) {
            var v12 float32 = 1e10
            for fi13 := 0; fi13 < 6; fi13++ {
                v3 = (v3 % (((^(v3))) | 1))
                if v4 {
                    var v14 uint8 = 255
                    _ = v14
                }
                var v15 uint8 = 2
                _ = fi13
                _ = v15
            }
            delete(v2, -2147483648)
            _ = v12
        } else {
            v16 := GenS2{F0: ("\n\t" + "a")}
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            if true {
                v4 = false
            }
            v18 := GenS3{F0: "hello", F1: 816}
            _ = fi17
            _ = v18
        }
        switch -9223372036854775808 {
        case 668:
            if v4 {
                var v19 float64 = (1.5 - (((0.0 / ((float64((3.14159 * 0.0)) - -2.25) / 0.1)) * 0.1) / (-2.25 * float64(1.5))))
                for fi20 := 0; fi20 < 14; fi20++ {
                    v19 -= 1e10
                    _ = fi20
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 int32 = 42
                    _ = fi21
                    _ = v22
                }
                _ = v19
            }
        case 32768:
            v4 = v4
        case 2147483647:
            var v23 int16 = -32768
            _ = v23
        default:
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v3)
            case "\n\t":
                var v24 uint8 = 161
                _ = v24
            case "hello":
                var v25 float32 = ((0.1 - 1e10) + ((0.0 / ((float32(0.0) - 3.14159) + 0.0)) / 0.1))
                _ = v25
            default:
                v3 -= (v3 / (((v3 & (v3 % (((^(v3))) | 1)))) | 1))
            }
        }
    } else {
        v4 = false
        var v26 int8 = -1
        fmt.Println(v4, v26, v4)
        _ = v26
    }
    var v27 uint16 = 746
    var v28 bool = (v4 && v4)
    var v29 uint16 = v27
    if v28 {
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int8 = 127
            _ = fi30
            _ = v31
        }
        var v32 uint8 = 253
        _ = v32
    } else {
        v33 := v2[127]
        v28 = (2147483648 > 42)
        v34 := GenS3{F0: (("a" + "foo bar") + "hello"), F1: 1}
        _ = v33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
