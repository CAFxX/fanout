package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 int64, p2 int) int16 {
    var v5 uint8 = 42
    var v6 int16 = 42
    p1 = p1
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{}
        v9 := GenS1{F0: ((v6 << ((65535) & 63)) == v6), F1: 42, F2: 963}
        v6 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10, v6)
        if (255 != 42) {
            var v11 int64 = p1
            var v12 int32 = -736
            v13 := GenS1{F0: ((("foo bar" + "\n\t") >= "hello") && (42 < v5)), F1: 901, F2: 4294967295}
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: (2 > 2), F1: 983, F2: 420}
            var v15 int = p0
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v6 >> ((0) & 63))
}
func genFn1() int {
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint32 = 494
        _ = fi16
        _ = v17
    }
    v18 := []int{32767, -1}
    switch 32767 {
    case 4294967295:
        if true {
            var v19 uint32 = 285
            _ = v19
        }
    default:
        var v20 int32 = 989
        _ = v20
    }
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 int = 128
    v2 := []uint32{279, 1}
    var v3 uint = 695
    v4 := []string{("hello" + "a")}
    var v21 int8 = 35
    v21 = v21
    switch -2147483648 {
    case -959:
        v22 := GenS1{F0: ((false || (201 != 42)) || true), F1: 2147483648, F2: (v3 ^ (v3 | v3))}
        _ = v22
    case 127:
        for fi23 := 0; fi23 < 15; fi23++ {
            switch "hello" {
            case "\n\t":
                var v24 uint64 = 452
                _ = v24
            default:
                v21 = 1
            }
            if true {
                for fi25 := 0; fi25 < 9; fi25++ {
                    if true {
                        fmt.Println(fi23, v21, v21)
                    }
                    _ = len(v4)
                    _ = fi25
                }
            }
            _ = fi23
        }
    default:
        for fi26 := 0; fi26 < 8; fi26++ {
            v1 = fi26
            v4 = append(v4, "hello")
            v27 := GenS2{F0: v3, F1: 18446744073709551615, F2: -1}
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS1{F0: ("a" <= ""), F1: 2147483647, F2: v3}
    v29 := GenS3{F0: ("" + "\n\t"), F1: (850 < 1)}
    for fi30 := 0; fi30 < 5; fi30++ {
        _ = len(v4)
        var v31 bool = false
        var v32 float32 = 0.0
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v1, v21, v1)
    v1 = (v1 * (v1 & v1))
    var v33 float64 = ((-2.25 + 3.14159) - float64(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
