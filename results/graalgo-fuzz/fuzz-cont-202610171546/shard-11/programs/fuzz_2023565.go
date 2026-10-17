package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) int64 {
    if ((p0 && true) && (!p0)) {
        v7 := []int{128}
        v8 := GenS1{F0: 255, F1: 2147483647}
        _ = v7
        _ = v8
    } else {
        v9 := GenS2{F0: 65535}
        _ = v9
    }
    v10 := []int{1, 9223372036854775807, 42, 4294967295}
    if false {
        if p0 {
            p0 = p0
            fmt.Println(p0)
        }
    }
    var v11 uint = 697
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]uint32)
    v3 := []int{261}
    var v4 float32 = 0.0
    var v5 int8 = 100
    var v6 uint = (^(v1))
    if (v5 <= v5) {
        v5 += (v5 * (v5 | -1))
        var v12 uint16 = 32767
        v5 = v5
        _ = v12
    } else {
        v1 = v1
    }
    var v13 uint8 = 128
    fmt.Println(v5, v1, v1)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int8 = (^(v5))
        v1 = 418
        switch -140 {
        case 127:
            var v16 uint8 = v13
            _ = v16
        case -1:
            switch "foo bar" {
            case "":
                var v17 bool = true
                _ = v17
            case "foo bar":
                if (1 != 12) {
                    v1 += 0
                    if len(v3) > 0 {
                        v3[0] = (fi14 << ((v1) & 63))
                    }
                } else {
                    _ = len(v3)
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v15 += (v5 ^ 80)
                        _ = fi18
                    }
                }
            default:
                var v19 uint64 = 42
                _ = v19
            }
        case -2147483648:
            v4 = float32((1.5 - 1e10))
        default:
            v20 := GenS2{F0: 169}
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    if (true && true) {
        var v21 uint = v6
        v22 := GenS1{F0: v13, F1: 42}
        _ = v21
        _ = v22
    }
    var v23 int8 = (v5 - (v5 / ((v5) | 1)))
    delete(v2, 9223372036854775807)
    var v24 bool = false
    if len(v3) > 0 {
        v3[0] = -670
    }
    v5 = 100
    for fi25 := 0; fi25 < 19; fi25++ {
        v5 = v23
        v5 = (v5 >> (((v6 % ((1) | 1))) & 63))
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
