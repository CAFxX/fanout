package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 int16) uint8 {
    var v7 uint16 = 230
    var v8 uint32 = 42
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int32 = 0
    v2 := []int{}
    var v3 int32 = v1
    var v4 uint16 = 1
    var v5 uint8 = 127
    var v6 uint32 = 4294967295
    if false {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 string = ("hello" + ("a" + ("\n\t" + "hello")))
            _ = fi9
            _ = v10
        }
    }
    v5 = (v5 / (((v5 - (v5 ^ (^(v5))))) | 1))
    v4 = (v4 + v4)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int8 = 42
        _ = len(v2)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: (v5 << ((153) & 63))}
    if true {
        var v14 float32 = (1e10 + 0.0)
        _ = v14
    }
    switch "hello" {
    case "hello":
        if (!(true && false)) {
            if true {
                var v15 bool = false
                _ = v15
            }
            var v16 uint8 = 176
            _ = v16
        }
    case "\n\t":
        v5 -= (^(v5))
    case "a":
        v4 -= (v4 + 2)
    default:
        var v17 string = "foo bar"
        _ = v17
    }
    var v18 bool = (v6 <= v6)
    if (!false) {
        switch 4294967295 {
        case -174:
            v1 = v3
        case 2147483647:
            if v18 {
                v19 := GenS2{F0: (283 > 42)}
                for fi20 := 0; fi20 < 19; fi20++ {
                    if (v5 == (v5 & v5)) {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            if len(v2) > 0 {
                                v2[0] = fi21
                            }
                            fmt.Println(v4, v3)
                            _ = fi21
                        }
                    }
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v23 int32 = v3
                        var v24 int8 = -32
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi20
                }
                _ = v19
            }
        }
    }
    switch 4294967295 {
    case -1:
        fmt.Println(v18)
    default:
        v2 = append(v2, 32768)
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        for fi25 := 0; fi25 < 17; fi25++ {
            v18 = v18
            _ = fi25
        }
    default:
        v26 := GenS1{F0: v4, F1: (v5 & genFn0(v5, -32768))}
        _ = v26
    }
    v3 -= (v1 | 1)
    if ((float32(3.14159) * ((((float32(0.0) + ((0.1 - (0.0 / 0.1)) * 1e10)) * (((0.0 * 1.5) - 1.5) / (((((1e10 / 1.5) / float32(0.1)) + float32(1e10)) / 1e10) + (-2.25 / (1e10 - 1.5))))) * 0.1) + 3.14159)) != 3.14159) {
        fmt.Println(v1)
        var v27 int32 = v3
        _ = v27
    }
    v4 = v4
    v28 := GenS2{F0: v18}
    var v29 int = -824
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
