package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }
func genFn0(p0 int32, p1 int16) int16 {
    v4 := []int{}
    switch 9223372036854775807 {
    case 610:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v4 = append(v4, -33)
    var v5 int32 = (p0 | (-(p0)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 uint, p1 uint) uint64 {
    v6 := []int{-2147483648, 32767, 127, -569}
    var v7 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int8 = 127
    var v2 uint64 = 2
    var v3 int8 = 108
    var v8 bool = (("hello" + "\n\t") <= ("\n\t" + "foo bar"))
    v1 = v3
    for fi9 := 0; fi9 < 3; fi9++ {
        v2 -= 4294967295
        _ = fi9
    }
    v3 -= (v1 - v1)
    v8 = ((!v8) && (42 != 65535))
    var v10 string = ""
    var v11 uint = 4294967295
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v3 -= 42
    case "a":
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v11)
            _ = fi12
        }
    default:
        var v13 uint16 = uint16(v2)
        _ = v13
    }
    v1 = 0
    v14 := []int{128, int(v2), 32767, -24}
    v15 := GenS1{F0: 1000, F1: -858}
    switch 9223372036854775807 {
    case -617:
        v14 = append(v14, 127)
    case 32768:
        fmt.Println(v10, v11, v11)
    case 32767:
        fmt.Println(v1, v1)
    default:
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 float32 = 1e10
                var v19 int8 = (v3 % (((^(v3))) | 1))
                v20 := GenS2{F0: 42, F1: 549, F2: 42}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi16
        }
    }
    var v21 uint16 = 1
    if v8 {
        v22 := GenS1{F0: genFn0(-2147483648, 0), F1: 39}
        _ = v22
    } else {
        var v23 float32 = ((1.5 + 0.0) * ((0.0 - (3.14159 - (0.0 * 3.14159))) * 1.5))
        _ = v23
    }
    v2 += v2
    var v24 bool = ((((0 > -293) && v8) || (!v8)) && v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v8
}
