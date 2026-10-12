package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 646
    v2 := []int{4294967295, 42, 2147483647, -2147483648}
    var v3 int16 = 1
    var v4 uint64 = 18446744073709551615
    if true {
        switch int(v4) {
        case 229:
            switch ("foo bar" + ("a" + "foo bar")) {
            case "hello":
                v3 -= (v3 - (^(v3)))
            case "a":
                v4 -= v4
            default:
                fmt.Println(v3)
            }
        case 906:
            var v5 uint32 = 2147483647
            _ = v5
        case -2147483648:
            v3 = (^(v3))
        }
        var v6 uint = 2
        if true {
            if true {
                if (((1 < -55) && (false || true)) && (-9223372036854775808 >= 1)) {
                    v1 = (v1 - v1)
                }
            }
            switch -72 {
            case -423:
                v4 -= (v4 & v4)
            }
        } else {
            v7 := GenS1{F0: "\n\t", F1: (128 < 179), F2: 255}
            v1 -= v1
            _ = v7
        }
        _ = v6
    } else {
        v4 -= 4294967295
    }
    if (454 > 42) {
        v3 = (^(v3))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v1 = (v1 % ((65535) | 1))
        v4 -= (v4 * (v4 >> ((42) & 63)))
        _ = fi8
    }
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v10 uint8 = uint8(v1)
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v1, v4)
    switch 945 {
    case -2147483648:
        var v11 int8 = -1
        _ = v11
    case 42:
        var v12 uint32 = 695
        _ = v12
    case 2147483647:
        fmt.Println(v4)
    default:
        var v13 int = 632
        _ = v13
    }
    v4 += (^(v4))
    var v14 float64 = (3.14159 + 3.14159)
    switch "" {
    case "\n\t":
        v1 = v1
    case "hello":
        var v15 uint32 = v1
        _ = v15
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v16 int = 2147483647
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch "a" {
    case "hello":
        v14 = v14
    case "\n\t":
        if (128 > int(v4)) {
            v17 := GenS2{F0: v4}
            var v18 int32 = 42
            _ = v17
            _ = v18
        }
    default:
        _ = len(v2)
    }
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
