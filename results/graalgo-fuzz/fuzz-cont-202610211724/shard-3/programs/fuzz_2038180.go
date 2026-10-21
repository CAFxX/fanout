package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint32) uint8 {
    var v5 int16 = -1
    if false {
        if true {
            v6 := []int{2147483647, 157, 344}
            _ = v6
        }
        var v7 int32 = -2147483648
        _ = v7
    }
    var v8 uint32 = p0
    if ((-2.25 * (1.5 / 3.14159)) != 1.5) {
        fmt.Println(v8, p0, v5)
    }
    _ = p0
    _ = v5
    _ = v8
    return 118
}

func main() {
    var v1 uint64 = 42
    var v2 int = 9223372036854775807
    v3 := make(map[string]int)
    v4 := []uint32{}
    if len(v4) > 0 {
        v4[0] = 475
    }
    switch 32768 {
    case 1:
        switch "" {
        case "":
            if (((0.0 * float32(-2.25)) + (0.0 / (0.0 - -2.25))) < 1.5) {
                v2 += v2
            }
        case "\n\t":
            if (236 <= 32) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int8 = 63
                    for fi11 := 0; fi11 < 19; fi11++ {
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                        var v12 uint32 = 65535
                        _ = fi11
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
                v1 += 42
                switch v2 {
                case 2147483647:
                    if false {
                        fmt.Println(v2)
                        v1 = 42
                        v13 := GenS1{F0: (0.1 <= 1e10)}
                        _ = v13
                    }
                case -9223372036854775808:
                    var v14 uint64 = (v1 << ((686) & 63))
                    _ = v14
                case 42:
                    var v15 int = v2
                    _ = v15
                default:
                    v16 := GenS1{F0: (-344 <= 1)}
                    _ = v16
                }
            }
        case "a":
            v1 -= v1
        }
    case 9223372036854775807:
        v1 = v1
    case 32767:
        if false {
            v1 += (^(v1))
        } else {
            v3[("a" + "a")] = 9223372036854775807
        }
    }
    fmt.Println(v1, v1)
    var v17 uint = 113
    switch (v2 / ((v2) | 1)) {
    case 32767:
        var v18 int16 = -731
        _ = v18
    case 296:
        fmt.Println(v2)
    default:
        var v19 int8 = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        if (-132 != 1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            v21 := GenS1{F0: true}
            delete(v3, "foo bar")
            _ = v21
        }
        var v22 int8 = -63
        v23 := v3["a"]
        _ = fi20
        _ = v22
        _ = v23
    }
    var v24 uint8 = 255
    switch v2 {
    case -9223372036854775808:
        v25 := GenS1{F0: ((0.0 < ((0.0 * 0.1) * (1.5 / float64((3.14159 - 0.0))))) || true)}
        _ = v25
    case -408:
        v26 := GenS2{F0: 1, F1: v24}
        _ = v26
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
