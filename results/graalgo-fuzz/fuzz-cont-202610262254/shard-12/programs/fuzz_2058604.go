package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint32) uint16 {
    v5 := GenS1{F0: 42, F1: 1, F2: 648}
    v6 := GenS2{F0: 65535}
    var v7 int32 = 42
    if true {
        var v8 int16 = 42
        var v9 int32 = (^(v7))
        _ = v8
        _ = v9
    } else {
        switch 0 {
        case 127:
            v7 = v7
        default:
            v7 -= v7
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 457
}

func main() {
    v1 := []uint32{1, 1, 854, 42}
    var v2 int64 = 9223372036854775807
    var v3 int32 = 460
    v4 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint = 133
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS3{F0: 9}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if (141 >= genFn0((v2 | -783), 2147483647)) {
        v4[42] = (float64(0.1) == 1.5)
        switch ("a" + (((("hello" + ("\n\t" + "\n\t")) + "") + ("" + "foo bar")) + "hello")) {
        case "":
            fmt.Println(v3)
        case "foo bar":
            var v14 int8 = -1
            _ = v14
        }
    }
    fmt.Println(v3, v3)
    switch 32767 {
    case 127:
        for fi15 := 0; fi15 < 19; fi15++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            _ = fi15
        }
    case 9223372036854775807:
        v2 = v2
    case -844:
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 int8 = 42
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    default:
        var v19 uint = 536
        _ = v19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (42 != 0) {
        v3 = 2147483647
    } else {
        switch 42 {
        case -693:
            fmt.Println(v3, v2)
        case 2147483647:
            v3 += v3
        default:
            v3 -= (v3 & -933)
        }
        var v20 bool = ((!true) && false)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v20
    }
    if true {
        v2 -= v2
        if false {
            v21 := GenS1{F0: 0, F1: 0, F2: 499}
            var v22 float64 = (1.5 / 1.5)
            _ = v21
            _ = v22
        } else {
            v3 += 447
        }
    }
    fmt.Println(v2, v2)
    v23 := GenS4{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
