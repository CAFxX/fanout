package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 int32; F2 int16 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint8) uint64 {
    var v6 uint16 = 850
    var v7 bool = (18446744073709551615 > 1)
    _ = p0
    _ = v6
    _ = v7
    return 404
}

func main() {
    var v1 uint8 = 202
    v2 := make(map[int]string)
    v3 := []int{-2147483648, 9223372036854775807, -1}
    var v4 uint16 = 65535
    v5 := make(map[string]bool)
    v8 := GenS1{F0: 968, F1: -1}
    v9 := GenS1{F0: 886, F1: 42}
    v1 -= v1
    if true {
        var v10 int16 = 32767
        v11 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = v10
        _ = v11
    }
    v4 += (v4 >> ((11) & 63))
    v4 += v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "" {
    case "\n\t":
        var v12 int64 = 127
        _ = v12
    case "":
        if len(v3) > 0 {
            v3[0] = -771
        }
    }
    v13 := GenS2{F0: int(v4), F1: 0, F2: 1}
    v1 -= v1
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v1 = (v1 >> ((734) & 63))
            v4 = v4
            v3 = append(v3, fi15)
            v16 := GenS3{F0: -625, F1: v4}
            v1 = (v1 >> ((18) & 63))
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = fi14
        _ = v17
    }
    _ = len(v3)
    _ = len(v5)
    switch 429 {
    case 4294967295:
        var v18 int32 = -70
        _ = v18
    case 2147483647:
        var v19 int8 = -128
        _ = v19
    case -850:
        v20 := GenS1{F0: 985, F1: 2147483647}
        _ = v20
    }
    v4 = (v4 * v4)
    v3 = append(v3, 2147483647)
    fmt.Println(v4)
    switch ((((("hello" + "a") + "foo bar") + "a") + "a") + "") {
    case "hello":
        switch "a" {
        case "foo bar":
            v1 += 2
        case "a":
            v4 += v4
        default:
            v1 += 85
        }
    case "foo bar":
        v21 := GenS1{F0: uint(v1), F1: 41}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
