package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 bool) uint16 {
    var v4 uint16 = 32767
    v5 := []int{-9223372036854775808, -972}
    if p0 {
        fmt.Println(v4, p0, p0)
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 244}
            var v8 int = -1
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: -815}
        _ = v9
    } else {
        switch "foo bar" {
        case "foo bar":
            v10 := GenS2{F0: 4294967295}
            _ = v10
        case "\n\t":
            var v11 uint64 = 422
            _ = v11
        case "a":
            p0 = (1e10 == 1e10)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int64, p1 uint8) uint64 {
    p1 = p1
    var v12 int8 = 97
    var v13 int = 128
    if (486 < 2) {
        switch ("foo bar" + ("foo bar" + ("foo bar" + ""))) {
        case "a":
            p0 += p0
        default:
            p0 = int64(p1)
        }
        fmt.Println(p0, v12, p0)
        v13 = v13
    } else {
        v14 := []int{32768, (v13 - (^(v13)))}
        v12 += (v12 & v12)
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := []int{(v13 ^ (v13 >> ((65535) & 63)))}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := []uint32{0, 42, 1}
    var v2 float64 = float64(((0.1 - -2.25) + 0.0))
    var v3 string = ("a" + "foo bar")
    switch "\n\t" {
    case "hello":
        switch 1 {
        case 754:
            var v17 uint8 = 195
            _ = v17
        case 9223372036854775807:
            v2 -= v2
        default:
            var v18 uint = 1
            _ = v18
        }
    case "foo bar":
        var v19 uint16 = 2
        _ = v19
    case "a":
        var v20 int8 = -74
        _ = v20
    default:
        v21 := GenS3{F0: -2147483648}
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = "foo bar"
    var v22 int8 = -17
    var v23 uint8 = 127
    v1 = append(v1, 2147483648)
    var v24 int16 = -32768
    v24 -= v24
    var v25 uint64 = 143
    v23 = v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
