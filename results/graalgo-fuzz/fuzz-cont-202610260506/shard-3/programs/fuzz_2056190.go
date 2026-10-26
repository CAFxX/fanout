package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 uint) uint16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        p0 = (^(p0))
        _ = fi5
    }
    p0 += 4294967295
    var v6 int8 = 100
    _ = p0
    _ = v6
    return 65535
}
func genFn1(p0 uint64) bool {
    v7 := []int{4294967295, 9223372036854775807, 128}
    var v8 int = 42
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 bool = (false && (42 <= 126))
        v11 := GenS1{F0: 42, F1: "foo bar", F2: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint64 = (p0 * 188)
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    var v13 int64 = 584
    v13 = 9
    _ = len(v4)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            v4[4294967295] = 2
            _ = fi14
        }
    case "a":
        for fi15 := 0; fi15 < 8; fi15++ {
            v3 = (v3 && v3)
            v16 := GenS1{F0: int32(fi15), F1: "", F2: -88}
            var v17 uint16 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
    case "\n\t":
        if ((true && (102 != -9223372036854775808)) || true) {
            switch "hello" {
            case "hello":
                switch 32767 {
                case 9223372036854775807:
                    v18 := GenS1{F0: -2147483648, F1: "a", F2: -119}
                    _ = v18
                case -1:
                    fmt.Println(v13, v2, v3)
                }
            case "\n\t":
                v13 -= 1
            default:
                switch 127 {
                case 32767:
                    v19 := GenS2{F0: (^(v2)), F1: 4294967295}
                    _ = v19
                }
            }
        }
    default:
        v3 = v3
    }
    var v20 int8 = 100
    if v3 {
        if false {
            v21 := []int{0, -596}
            _ = v21
        } else {
            var v22 uint16 = 1
            _ = v22
        }
        v23 := []int{-2147483648, 127}
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = v23
    }
    if (1e10 >= -2.25) {
        if (2 != 4294967295) {
            v2 = (v2 & (v2 | 42))
        }
        var v24 int = 2147483647
        _ = v24
    } else {
        v25 := []int{}
        var v26 float64 = (0.0 + 0.1)
        v3 = (!((!(2 == 91)) || (1 < 65535)))
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (v2 << ((255) & 63)), F1: "a", F2: -107}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
