package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 uint32 = 696
    var v2 int64 = -1
    v3 := []uint32{65535, (v1 ^ v1), (v1 / ((581) | 1)), v1}
    var v4 uint8 = 0
    var v5 uint = 65535
    if len(v3) > 0 {
        v3[0] = 427
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float64 = 1.5
        v8 := GenS1{F0: 42, F1: 750, F2: int16(v4)}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v9 int64 = 42
        _ = v9
    }
    var v10 int32 = -968
    fmt.Println(v10)
    v11 := GenS2{F0: "a", F1: v1, F2: v1}
    v4 -= 42
    if true {
        for fi12 := 0; fi12 < 5; fi12++ {
            v5 = 664
            _ = len(v3)
            if true {
                var v13 int16 = -263
                v10 = v10
                _ = v13
            } else {
                v2 = v2
                v1 += (v1 << ((v5) & 63))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi12
        }
        var v14 bool = true
        var v15 string = "\n\t"
        _ = v14
        _ = v15
    } else {
        if len(v3) > 0 {
            v3[0] = (v1 + (^(v1)))
        }
        v4 = v4
    }
    var v16 int = -735
    v2 += (v2 | 0)
    var v17 uint8 = 1
    fmt.Println(v2)
    v18 := GenS2{F0: "\n\t", F1: (v1 | 2147483647), F2: 2}
    v5 -= (v5 ^ 42)
    if true {
        var v19 uint64 = 224
        var v20 uint16 = 334
        var v21 int16 = 0
        _ = v19
        _ = v20
        _ = v21
    } else {
        switch "" {
        case "a":
            fmt.Println(v10, v10, v16)
        default:
            switch -375 {
            case 979:
                if ("foo bar" != (("foo bar" + "foo bar") + "a")) {
                    v22 := GenS3{F0: 717, F1: v10, F2: (^(v1))}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v22
                }
            default:
                var v23 string = ("" + ("\n\t" + ("a" + ("" + "hello"))))
                _ = v23
            }
        }
        fmt.Println(v5, v1, v17)
        v4 = (v17 * 127)
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
