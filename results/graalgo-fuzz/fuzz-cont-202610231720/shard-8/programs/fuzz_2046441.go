package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8) int32 {
    v4 := []int{127}
    if (32767 != -9223372036854775808) {
        p0 = p0
        var v5 string = "foo bar"
        v6 := GenS1{F0: 234, F1: -911, F2: 100}
        _ = v5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -986
}

func main() {
    v1 := []int{946}
    var v2 uint = 609
    var v3 int64 = -1
    if ((^(v2)) == (v2 - 0)) {
        v2 -= 576
    } else {
        var v7 int32 = 42
        _ = v7
    }
    if (255 > 42) {
        var v8 uint16 = 255
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint32 = 660
        _ = fi9
        _ = v10
    }
    if (127 < 6) {
        fmt.Println(v3)
        v3 = v3
    } else {
        if (((3.14159 / 0.1) == 1e10) || (2 != 0)) {
            if true {
                v2 -= v2
                v2 = (v2 << (((^(v2))) & 63))
                if false {
                    switch "" {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = int(v2)
                        }
                    case "":
                        var v11 bool = false
                        _ = v11
                    case "a":
                        fmt.Println(v3)
                    }
                    var v12 string = ("foo bar" + ("foo bar" + "\n\t"))
                    _ = v12
                }
            } else {
                var v13 uint = (v2 | (v2 | (v2 | 772)))
                _ = v13
            }
            if ((477 < 2) || (258 != 1)) {
                if (true && true) {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        _ = len(v1)
                        _ = fi14
                    }
                    v3 -= (v3 >> ((v2) & 63))
                }
            }
            v3 = (v3 | v3)
        } else {
            var v15 float32 = -2.25
            _ = v15
        }
        v3 += (v3 ^ -9223372036854775808)
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch fi16 {
        case 318:
            var v17 uint16 = 0
            _ = v17
        default:
            var v18 float32 = 0.0
            _ = v18
        }
        _ = len(v1)
        _ = fi16
    }
    switch -9223372036854775808 {
    case 489:
        v3 = (v3 ^ v3)
    case 32768:
        var v19 uint32 = 65535
        _ = v19
    default:
        if ((502 <= 375) && false) {
            if (false || true) {
                fmt.Println(v2, v3)
            }
            v3 -= v3
        }
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v20 := GenS2{F0: -699}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
