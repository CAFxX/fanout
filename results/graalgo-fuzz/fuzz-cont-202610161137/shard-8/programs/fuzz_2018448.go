package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 int64; F1 int32 }
func genFn0(p0 bool, p1 int32, p2 uint64) uint64 {
    var v5 uint8 = 128
    v6 := GenS1{F0: v5}
    if (0 == 42) {
        v7 := GenS1{F0: 49}
        for fi8 := 0; fi8 < 12; fi8++ {
            p1 = (p1 & (p1 << ((1) & 63)))
            _ = fi8
        }
        v9 := GenS2{F0: -102, F1: 0}
        _ = v7
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 uint16 = 884
    var v2 int8 = -128
    var v3 int64 = 1
    v4 := []uint32{2, 1, 65535, 510}
    v3 -= 42
    switch "hello" {
    case "":
        for fi10 := 0; fi10 < 15; fi10++ {
            if len(v4) > 0 {
                v4[0] = 554
            }
            for fi11 := 0; fi11 < 15; fi11++ {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v1 = v1
                case "foo bar":
                    var v12 int32 = -2147483648
                    _ = v12
                case "a":
                    v13 := GenS3{F0: (fi11 ^ fi11), F1: (v1 * 32767), F2: v2}
                    _ = v13
                }
                var v14 bool = ((2 < 1) || (false || (992 >= 18446744073709551615)))
                v3 = (v3 & v3)
                _ = fi11
                _ = v14
            }
            v15 := GenS4{F0: v3, F1: 42}
            _ = fi10
            _ = v15
        }
    case "\n\t":
        switch 434 {
        case -9223372036854775808:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 2147483647:
            v1 = 0
        default:
            _ = len(v4)
        }
    case "foo bar":
        if true {
            fmt.Println(v1)
            v2 += 42
            var v16 uint32 = 42
            _ = v16
        } else {
            var v17 uint16 = v1
            _ = v17
        }
    default:
        v1 = (v1 + 0)
    }
    if true {
        var v18 int = -9223372036854775808
        v18 -= (v18 + v18)
        v3 -= v3
        _ = v18
    }
    var v19 uint32 = 4294967295
    var v20 uint16 = (v1 << ((1) & 63))
    switch ("hello" + "hello") {
    case "":
        var v21 int = -9223372036854775808
        _ = v21
    case "\n\t":
        v1 = v20
    case "hello":
        fmt.Println(v19)
    default:
        if len(v4) > 0 {
            v4[0] = (v19 ^ v19)
        }
    }
    v2 = v2
    v2 += (v2 / (((-(v2))) | 1))
    v4 = append(v4, v19)
    var v22 int8 = 100
    var v23 string = (("a" + ("hello" + ("hello" + ("hello" + "")))) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
