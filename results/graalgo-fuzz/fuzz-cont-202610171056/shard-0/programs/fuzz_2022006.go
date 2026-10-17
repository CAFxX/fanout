package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0(p0 uint8, p1 int64) uint8 {
    if ((p1 >> ((178) & 63)) == -127) {
        var v5 uint8 = (p0 << ((0) & 63))
        _ = v5
    }
    p0 -= p0
    v6 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v6
    return (p0 & (^(p0)))
}
func genFn1(p0 string) uint64 {
    v7 := GenS1{F0: 1}
    p0 = p0
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 string = "a"
        _ = fi8
        _ = v9
    }
    p0 = (p0 + (p0 + p0))
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = 42
    var v3 string = ("foo bar" + ("a" + ""))
    v4 := []uint32{2147483647, 65535}
    fmt.Println(v3, v2)
    if (!true) {
        switch 32768 {
        case -1:
            if false {
                if len(v4) > 0 {
                    v4[0] = 2
                }
            }
        default:
            fmt.Println(v3)
        }
        v3 = v3
        v4 = append(v4, 666)
    } else {
        var v10 uint32 = 186
        var v11 string = v3
        _ = v10
        _ = v11
    }
    var v12 int32 = 2147483647
    fmt.Println(v3)
    fmt.Println(v2)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v12 = (v12 << ((134) & 63))
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 int = (fi13 * 1)
                    v2 = -1
                    _ = fi14
                    _ = v15
                }
                v1 = v1
                _ = fi13
            }
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        }
        v3 = v3
    } else {
        v16 := GenS2{F0: (v12 * 42), F1: 883, F2: 2}
        _ = v16
    }
    fmt.Println(v1)
    _ = len(v4)
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
