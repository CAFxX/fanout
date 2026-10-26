package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0(p0 uint64, p1 int8) uint32 {
    v5 := GenS1{F0: int16(p1), F1: -307}
    fmt.Println(p1)
    switch "\n\t" {
    case "\n\t":
        p1 += p1
    case "foo bar":
        p0 -= (p0 >> ((476) & 63))
    case "":
        var v6 uint = uint(p0)
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 349
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 79
    var v3 bool = (!(-1 < 567))
    var v4 float32 = -2.25
    v2 = v2
    if (v3 || (32767 < 1)) {
        v7 := GenS1{F0: 479, F1: 639}
        _ = v7
    }
    fmt.Println(v4)
    var v8 uint64 = v2
    v4 -= float32(1.5)
    for fi9 := 0; fi9 < 2; fi9++ {
        v8 -= v8
        var v10 uint32 = 571
        fmt.Println(v3, v8, v2)
        _ = fi9
        _ = v10
    }
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 12; fi12++ {
        if ((v3 || v3) || v3) {
            v4 += 3.14159
        }
        v8 += v2
        _ = fi12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := []int{}
        if false {
            fmt.Println(v4, v8, v8)
            v15 := GenS2{F0: -9223372036854775808, F1: 848, F2: 1}
            if len(v14) > 0 {
                v14[0] = (fi13 << ((18446744073709551615) & 63))
            }
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(fi16, fi16)
                _ = fi16
            }
            switch -1 {
            case 2147483647:
                var v17 uint8 = 255
                _ = v17
            default:
                fmt.Println(fi13)
            }
        }
        _ = fi13
        _ = v14
    }
    if (v3 || v3) {
        switch "" {
        case "a":
            var v18 float32 = (v4 * ((v4 * v4) * 3.14159))
            _ = v18
        case "":
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int = fi19
                switch "a" {
                case "":
                    v21 := GenS2{F0: 42, F1: -929, F2: 175}
                    _ = v21
                case "foo bar":
                    var v22 int16 = 42
                    _ = v22
                default:
                    v11 = 1e10
                }
                _ = fi19
                _ = v20
            }
        case "foo bar":
            v4 -= -2.25
        }
        switch ("a" + "foo bar") {
        case "foo bar":
            v1[65535] = 2147483647
        }
        switch 128 {
        case -64:
            switch 1 {
            case 0:
                if (-417 <= 0) {
                    v23 := []int{255, 0, 128, 0}
                    v23 = append(v23, -2147483648)
                    v4 = v4
                    _ = v23
                }
            case 128:
                v4 = v4
            default:
                v2 = (v2 ^ (v2 - (v8 - v2)))
            }
        case -490:
            fmt.Println(v11, v2)
        default:
            v8 -= (v2 * v8)
        }
    }
    var v24 uint16 = 65535
    v8 -= (v8 + (v8 | (v2 * v8)))
    var v25 string = "a"
    v8 = (v2 ^ v2)
    v25 = "a"
    fmt.Println(v3)
    delete(v1, 65535)
    v26 := GenS2{F0: 0, F1: 42, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
