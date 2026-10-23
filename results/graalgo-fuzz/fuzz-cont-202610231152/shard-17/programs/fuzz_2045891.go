package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 uint64, p1 int8, p2 uint64) uint16 {
    switch 32767 {
    case 1:
        p1 -= (^(p1))
    case 371:
        v5 := GenS1{F0: (p1 % ((p1) | 1)), F1: 899, F2: -9223372036854775808}
        _ = v5
    case 0:
        var v6 float32 = 0.1
        _ = v6
    default:
        p0 = 18446744073709551615
    }
    v7 := []int{4294967295}
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}
func genFn1(p0 int32, p1 uint) int8 {
    v8 := []int{-2147483648, -89, 994, -9223372036854775808}
    fmt.Println(p1, p0)
    var v9 string = ("" + ("a" + ("foo bar" + "a")))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 string = "a"
    v2 := []uint32{65535, 0, 1, 2147483647}
    var v3 int64 = 42
    v4 := make(map[uint32]uint32)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := v4[42]
        if (255 != 161) {
            switch fi10 {
            case 805:
                v12 := GenS1{F0: -128, F1: 139, F2: fi10}
                _ = v12
            default:
                v1 = v1
            }
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v3, v3)
    fmt.Println(v1)
    if false {
        v3 = (v3 % ((42) | 1))
        if (!true) {
            delete(v4, 1)
            switch 1 {
            case 128:
                v3 = (v3 << ((697) & 63))
            }
        } else {
            switch 127 {
            case 475:
                fmt.Println(v3)
            default:
                v13 := GenS1{F0: 18, F1: 4294967295, F2: 32768}
                _ = v13
            }
        }
    } else {
        fmt.Println(v3, v1)
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if (-128 >= genFn1(241, 677)) {
            for fi15 := 0; fi15 < 2; fi15++ {
                v1 = (v1 + ((v1 + v1) + "hello"))
                switch ((v1 + v1) + v1) {
                case "":
                    var v16 uint = 42
                    _ = v16
                case "a":
                    var v17 string = v1
                    _ = v17
                }
                _ = fi15
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 float32 = (float32(0.1) - (3.14159 * 0.1))
                _ = fi18
                _ = v19
            }
        } else {
            v20 := GenS2{F0: 0, F1: 1}
            _ = v20
        }
        _ = fi14
    }
    switch "a" {
    case "foo bar":
        switch -2147483648 {
        case -1:
            if (!((0.1 / 0.1) != (0.0 * 1e10))) {
                _ = len(v2)
            }
        case -2147483648:
            var v21 uint = 1
            _ = v21
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "hello":
        var v22 uint32 = 4294967295
        _ = v22
    case "a":
        _ = len(v4)
    }
    if len(v2) > 0 {
        v2[0] = 376
    }
    v23 := GenS2{F0: 287, F1: 42}
    fmt.Println(v3, v3)
    var v24 int = 2147483647
    v25 := GenS2{F0: -1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
