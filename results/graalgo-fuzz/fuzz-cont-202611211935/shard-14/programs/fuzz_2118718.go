package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(((("" + (("\n\t" + "\n\t") + "")) + "foo bar") + "a") + "foo bar"), "", "foo bar", "hello"}
    var v3 int8 = -53
    v4 := make(map[uint32]uint32)
    fmt.Println(v3)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 255}
        var v7 int16 = 474
        fmt.Println(v3, fi5)
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        switch (fi8 << ((4294967295) & 63)) {
        case -736:
            for fi9 := 0; fi9 < 4; fi9++ {
                switch -1 {
                case -9223372036854775808:
                    fmt.Println(fi8)
                default:
                    v3 = v3
                }
                _ = len(v2)
                var v10 int8 = v3
                _ = fi9
                _ = v10
            }
        case 32767:
            fmt.Println(v3, fi8)
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                v3 -= (v3 << ((65535) & 63))
                _ = fi11
            }
        }
        _ = fi8
    }
    v3 -= -128
    switch 127 {
    case -547:
        v3 += v3
    default:
        if false {
            var v12 int8 = (^(v3))
            _ = v12
        } else {
            var v13 uint16 = 42
            v13 = (^(v13))
            var v14 uint = 897
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v3)
    if false {
        v15 := v1[0]
        _ = v15
    }
    fmt.Println(v3)
    v3 = (v3 / (((v3 << ((255) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
