package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint, p2 int16) int32 {
    v6 := GenS1{F0: 128, F1: 42, F2: (((0.0 * 1e10) / float32(1e10)) != 1.5)}
    p0 -= 426
    v7 := []int{2147483647, 826, -2147483648, 9223372036854775807}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = (p0 / (((p0 & (^(p0)))) | 1))
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if false {
        if (true || false) {
            p1 -= (^(p0))
            p2 = p2
        }
    } else {
        var v11 string = "a"
        var v12 uint = p1
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1() string {
    if false {
        var v13 bool = (!true)
        _ = v13
    }
    v14 := GenS1{F0: 2, F1: 42, F2: true}
    var v15 int32 = 2147483647
    switch 0 {
    case 0:
        v15 = (^(v15))
    case 9223372036854775807:
        if (-9223372036854775808 <= 9223372036854775807) {
            v16 := GenS2{F0: -9223372036854775808, F1: 161}
            _ = v16
        }
    case -2147483648:
        v17 := GenS1{F0: 2, F1: 42, F2: (-128 <= 80)}
        _ = v17
    default:
        var v18 int64 = 1
        _ = v18
    }
    _ = v14
    _ = v15
    return ("" + ("hello" + ("foo bar" + "a")))
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 uint = 42
    v4 := make(map[uint32]string)
    var v5 int32 = -1
    switch 4294967295 {
    case 128:
        var v19 int8 = 0
        _ = v19
    }
    switch "" {
    case "\n\t":
        v20 := GenS3{F0: 42, F1: (^(v5)), F2: (("" + (("" + ("\n\t" + "a")) + "a")) + ("" + "foo bar"))}
        _ = v20
    case "a":
        v21 := GenS4{F0: 42}
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        if (-128 <= -128) {
            var v23 int32 = (v5 << (((v3 & 255)) & 63))
            _ = v23
        } else {
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
        v24 := GenS1{F0: 0, F1: 42, F2: (2 != 2147483647)}
        _ = fi22
        _ = v24
    }
    v25 := GenS2{F0: 208, F1: 44}
    v3 = v3
    v5 += genFn0(18446744073709551615, (^(v3)), 42)
    if ("\n\t" == ("a" + "a")) {
        if true {
            v3 = (v3 << ((370) & 63))
        }
        v5 += 42
    }
    if true {
        v3 = v3
        switch 0 {
        case 4294967295:
            v5 = (^(v5))
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            v5 += (v5 * v5)
            v2[""] = (true || (false || ((42 > -1) && (!false))))
            _ = fi26
        }
    } else {
        v3 = (v3 | v3)
    }
    if false {
        v3 += v3
        var v27 uint = (v3 | 1)
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
