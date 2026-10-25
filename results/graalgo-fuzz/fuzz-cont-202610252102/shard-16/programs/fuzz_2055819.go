package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 int32, p2 int64) int8 {
    var v7 int8 = 0
    v8 := GenS1{F0: (p2 & -9223372036854775808), F1: uint32(p0), F2: (("\n\t" + "a") + (("a" + ("foo bar" + ("hello" + ("\n\t" + "hello")))) + "hello"))}
    switch 2147483647 {
    case 42:
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: -1, F1: 667, F2: "\n\t"}
            v11 := []int{(fi9 >> ((4294967295) & 63)), fi9, (fi9 | (-(fi9))), fi9}
            fmt.Println(p2)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 2
    v2 := []string{("\n\t" + "hello")}
    v3 := []int{1, 42, 0, -2147483648}
    var v4 int16 = 1
    v5 := make(map[uint32]bool)
    var v6 uint64 = 0
    v12 := GenS1{F0: 42, F1: uint32(v6), F2: ""}
    switch "foo bar" {
    case "a":
        if false {
            v13 := GenS1{F0: 0, F1: 2, F2: (("foo bar" + "foo bar") + ("a" + "a"))}
            v1 += v1
            _ = v13
        }
    default:
        v1 = (v1 >> ((1) & 63))
    }
    v14 := GenS2{F0: v4}
    v1 = v1
    var v15 int16 = -1
    if (!(81 > -1)) {
        v1 = (v1 >> ((v1) & 63))
    }
    v15 = (v4 - (-(v4)))
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            fmt.Println(fi17)
            v1 = 43
            _ = fi17
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            v4 -= (v4 << (((v1 % (((v1 * (v1 << (((v1 ^ (v1 | v1))) & 63)))) | 1))) & 63))
            switch "\n\t" {
            case "a":
                var v19 uint32 = 960
                _ = v19
            case "foo bar":
                delete(v5, 179)
            case "":
                fmt.Println(fi18, v4, fi18)
                var v20 int32 = -254
                var v21 uint16 = 32767
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v22 float64 = 0.1
                _ = v20
                _ = v21
                _ = v22
            default:
                var v23 int16 = (v15 * (v15 * v4))
                _ = v23
            }
            _ = fi18
        }
        v24 := GenS2{F0: (-(v4))}
        _ = fi16
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(v1, v6, v15)
        v1 += v1
        v4 -= v4
        _ = fi25
    }
    v26 := GenS1{F0: -1, F1: 42, F2: ("\n\t" + "\n\t")}
    fmt.Println(v1, v1)
    if (-1 >= 121) {
        if ((^(v6)) <= v6) {
            v4 = (v15 | (v15 / ((v4) | 1)))
        } else {
            fmt.Println(v1, v15, v1)
        }
        fmt.Println(v15, v6)
        if (35 == 127) {
            var v27 uint64 = v6
            _ = v27
        } else {
            v28 := GenS2{F0: (v15 << ((v1) & 63))}
            _ = len(v3)
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 uint8 = 128
                _ = fi29
                _ = v30
            }
            _ = v28
        }
    } else {
        v31 := GenS2{F0: 618}
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        switch fi32 {
        case 42:
            fmt.Println(v15, v6, v6)
        default:
            var v33 uint8 = 255
            _ = v33
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
