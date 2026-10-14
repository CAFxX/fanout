package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint64 = 2
    v2 := []string{}
    var v3 int8 = 58
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 0, F1: 65535}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if false {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 int32 = v4
                switch (fi6 ^ -894) {
                case 207:
                    for fi8 := 0; fi8 < 16; fi8++ {
                        v1 += v1
                        _ = fi8
                    }
                default:
                    for fi9 := 0; fi9 < 1; fi9++ {
                        var v10 uint16 = 906
                        fmt.Println(v7, v7)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi9
                        _ = v10
                    }
                }
                v4 -= v7
                _ = fi6
                _ = v7
            }
            v11 := GenS2{F0: v4, F1: (^(v3)), F2: 0}
            _ = v11
        }
        v12 := GenS2{F0: (v4 * v4), F1: (v3 + v3), F2: 0}
        _ = v12
    }
    var v13 float64 = 1.5
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint16 = 32767
        _ = fi14
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = (("\n\t" + ("foo bar" + "hello")) + "hello")
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 uint8 = 127
    switch (((("a" + "foo bar") + "foo bar") + "") + "a") {
    case "\n\t":
        var v17 bool = true
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        v19 := GenS1{F0: 0, F1: 65535}
        v20 := GenS3{F0: false, F1: (v4 % ((v4) | 1))}
        _ = fi18
        _ = v19
        _ = v20
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
