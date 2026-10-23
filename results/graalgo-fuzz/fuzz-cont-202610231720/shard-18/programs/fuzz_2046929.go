package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{}
    var v2 uint = 2
    v3 := make(map[int]uint32)
    switch -40 {
    case 42:
        fmt.Println(v2)
    }
    v2 = 355
    if true {
        if false {
            v4 := GenS1{F0: 42, F1: -385}
            _ = v4
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            v5 := GenS2{F0: (v2 & v2)}
            _ = v5
        case "hello":
            if true {
                var v6 uint64 = 18446744073709551615
                fmt.Println(v6)
                _ = v6
            } else {
                var v7 uint32 = 0
                _ = v7
            }
        }
    } else {
        v2 -= (v2 % ((18446744073709551615) | 1))
        delete(v3, -2147483648)
        v2 = v2
    }
    delete(v3, -9223372036854775808)
    var v8 float32 = ((0.0 * 0.0) / float32(0.1))
    v8 += float32(1.5)
    v1 = append(v1, 993)
    v9 := GenS2{F0: v2}
    switch 32768 {
    case 32767:
        var v10 int16 = 0
        _ = v10
    default:
        if true {
            v8 = (((-2.25 + v8) - 3.14159) + ((v8 - 1.5) * float32(1.5)))
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if ((!(1 <= 42)) || (184 == 221)) {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                v2 -= 387
                switch -1 {
                case 42:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v2 -= v2
                        var v12 int32 = -54
                        _ = fi11
                        _ = v12
                    }
                default:
                    fmt.Println(v8, v2)
                }
            }
        }
    }
    switch -565 {
    case -2147483648:
        var v13 uint16 = 1
        _ = v13
    case 9223372036854775807:
        v8 -= ((1.5 * 3.14159) * v8)
    case 32768:
        if len(v1) > 0 {
            v1[0] = 65535
        }
    }
    if true {
        v2 = (v2 << ((v2) & 63))
    }
    var v14 uint32 = 152
    var v15 bool = (false || ((("a" + "hello") + "\n\t") < (("" + ("\n\t" + ("foo bar" + ""))) + "a")))
    fmt.Println(v8)
    var v16 int64 = -1
    v8 -= (0.1 + (v8 + v8))
    if len(v1) > 0 {
        v1[0] = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
