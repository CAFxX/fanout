package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []string{"\n\t", "a", (("a" + ("hello" + ("\n\t" + ("a" + ((("\n\t" + "") + "a") + "a"))))) + "")}
    var v2 int16 = -197
    v3 := make(map[uint32]bool)
    var v4 int = -273
    fmt.Println(v4)
    if true {
        var v5 uint8 = 102
        v5 += (v5 << ((42) & 63))
        _ = v5
    } else {
        v1 = append(v1, "a")
        delete(v3, 2147483648)
    }
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 int16 = (v2 * (-(v2)))
                v2 += (v8 - v2)
                v9 := GenS1{F0: 18446744073709551615}
                _ = fi7
                _ = v8
                _ = v9
            }
            v4 = (v4 & fi6)
            _ = fi6
        }
        fmt.Println(v2, v2)
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v2, v2)
            _ = fi10
        }
    } else {
        _ = len(v1)
        switch "foo bar" {
        case "":
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 uint64 = 18446744073709551615
                _ = fi11
                _ = v12
            }
        default:
            v3[535] = true
        }
        v2 -= (v2 / (((v2 & (-(v2)))) | 1))
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int = (v4 << ((720) & 63))
        v15 := GenS2{F0: 458, F1: 0}
        v14 += (fi13 ^ (-(v14)))
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 0, F1: 18446744073709551615}
    v17 := GenS2{F0: v2, F1: 65535}
    var v18 int32 = -1
    fmt.Println(v4, v18)
    if (!(2 > 27)) {
        for fi19 := 0; fi19 < 4; fi19++ {
            v2 = v2
            fmt.Println(v2)
            _ = fi19
        }
        v20 := v3[417]
        var v21 int64 = -743
        _ = v20
        _ = v21
    }
    if (255 < 255) {
        v4 -= (v4 | v4)
        if (0 != 290) {
            v18 += (v18 << ((1) & 63))
        } else {
            v4 = v4
            var v22 uint8 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
        }
        v23 := GenS3{F0: 42}
        _ = v23
    }
    var v24 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
