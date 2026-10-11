package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string }

func main() {
    v1 := []int{2147483647}
    var v2 bool = (!true)
    var v3 bool = v2
    var v4 uint = 65535
    if v3 {
        for fi5 := 0; fi5 < 17; fi5++ {
            _ = len(v1)
            v1 = append(v1, (fi5 / ((fi5) | 1)))
            _ = fi5
        }
        v6 := GenS1{F0: 255, F1: 32767, F2: 4294967295}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v1 = append(v1, (fi7 & (fi7 / ((fi7) | 1))))
            v3 = (!(v3 || (100 >= -1)))
            _ = fi7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            switch "a" {
            case "foo bar":
                v3 = v3
            case "":
                if len(v1) > 0 {
                    v1[0] = -1
                }
            case "\n\t":
                fmt.Println(v3)
            default:
                fmt.Println(v2)
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                switch "\n\t" {
                case "foo bar":
                    if false {
                        v4 -= v4
                        v3 = (v4 > (v4 + (v4 + v4)))
                        _ = len(v1)
                    }
                case "a":
                    var v10 uint32 = uint32(v4)
                    _ = v10
                case "":
                    v11 := GenS1{F0: 127, F1: 979, F2: 2147483647}
                    _ = v11
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v2 = true
                v12 := GenS1{F0: 1, F1: 32767, F2: 0}
                _ = fi9
                _ = v12
            }
            _ = fi8
        }
        _ = len(v1)
    }
    var v13 int16 = 32767
    if (v3 || true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v13 = (v13 * (v13 * (^(v13))))
    switch -9223372036854775808 {
    case -327:
        var v14 int16 = (v13 % ((v13) | 1))
        _ = v14
    default:
        switch "\n\t" {
        case "a":
            var v15 float32 = 1.5
            _ = v15
        case "foo bar":
            v4 = (v4 << (((^(v4))) & 63))
        }
    }
    switch "\n\t" {
    case "":
        v16 := GenS1{F0: 42, F1: 65535, F2: 4294967295}
        _ = v16
    case "foo bar":
        var v17 float64 = 0.0
        _ = v17
    default:
        v18 := GenS2{F0: -566, F1: (v4 ^ (v4 << ((2) & 63)))}
        _ = v18
    }
    v4 = 0
    var v19 int8 = 127
    fmt.Println(v3, v2)
    v3 = ((v4 ^ 0) < (v4 & v4))
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS2{F0: 32768, F1: 0}
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: (((((("foo bar" + "") + "a") + "") + "foo bar") + ("foo bar" + "hello")) + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
