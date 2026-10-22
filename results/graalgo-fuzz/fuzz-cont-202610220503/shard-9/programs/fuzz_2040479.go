package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int }
func genFn0() int16 {
    var v6 uint16 = 842
    v6 -= (v6 + (v6 << ((255) & 63)))
    var v7 uint = 2
    _ = v6
    _ = v7
    return int16(v7)
}
func genFn1(p0 uint64, p1 int16) bool {
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint16 = 65535
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = v11
    return (255 == 13)
}

func main() {
    v1 := []uint32{65535, 2, 2147483647}
    var v2 int = 0
    var v3 uint8 = 1
    v4 := []string{"\n\t", (("foo bar" + (("hello" + "hello") + "")) + "a")}
    var v5 int16 = -166
    _ = len(v4)
    v12 := GenS2{F0: 707, F1: v2}
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: 1, F1: fi13}
        fmt.Println(v2, v2)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v3 -= v3
        if (false && (uint(v3) < 4294967295)) {
            fmt.Println(v2, v5)
            switch (-(fi15)) {
            case -1:
                v5 += 42
            case -823:
                v3 = (v3 - v3)
            case 4294967295:
                var v16 uint16 = 640
                _ = v16
            default:
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v2, v2)
                default:
                    var v17 uint32 = 956
                    _ = v17
                }
            }
            fmt.Println(v3)
        } else {
            var v18 float64 = (1e10 + ((0.0 / 1.5) - (3.14159 * 0.0)))
            _ = v18
        }
        v3 = v3
        _ = fi15
    }
    v19 := GenS2{F0: 1, F1: v2}
    fmt.Println(v3, v2)
    var v20 uint64 = 42
    v21 := GenS2{F0: 255, F1: -9223372036854775808}
    switch v2 {
    case 96:
        v22 := GenS1{F0: (v3 & 128)}
        _ = v22
    case 933:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 32767:
        var v23 int32 = 651
        _ = v23
    }
    var v24 bool = false
    _ = len(v4)
    v1 = append(v1, 500)
    v2 -= 0
    fmt.Println(v20)
    var v25 uint8 = v3
    var v26 uint64 = (^(v20))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
