package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 int32, p1 int32) uint16 {
    if (0 == 1) {
        p1 = p0
        p0 += (p1 & p0)
    } else {
        switch "foo bar" {
        case "\n\t":
            var v5 int8 = 7
            _ = v5
        default:
            for fi6 := 0; fi6 < 10; fi6++ {
                p1 -= (p0 % (((^(p1))) | 1))
                _ = fi6
            }
        }
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 0
        p0 -= p0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        p0 -= (p0 * p1)
        var v10 int64 = 1
        var v11 uint16 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v12
    return 65535
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 uint64 = 740
    v4 := make(map[string]string)
    switch 220 {
    case 42:
        v3 -= (v3 << ((1) & 63))
    }
    if false {
        v13 := GenS1{F0: 0, F1: -9223372036854775808}
        v14 := GenS1{F0: 806, F1: 1}
        v3 -= (v3 & v3)
        _ = v13
        _ = v14
    }
    v3 += v3
    v3 += 1
    var v15 string = (("foo bar" + "hello") + "\n\t")
    v15 = (v15 + v15)
    fmt.Println(v15, v15)
    var v16 int16 = -101
    v3 = (v3 - (^(v3)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 4294967295 {
    case -9223372036854775808:
        for fi17 := 0; fi17 < 1; fi17++ {
            fmt.Println(fi17)
            v18 := GenS1{F0: -802, F1: -9223372036854775808}
            _ = fi17
            _ = v18
        }
    case 1:
        switch 135 {
        case 127:
            v16 += v16
        case 4294967295:
            var v19 int = 4294967295
            _ = v19
        }
    default:
        v20 := GenS2{F0: 1, F1: false, F2: (v16 % ((v16) | 1))}
        _ = v20
    }
    var v21 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
