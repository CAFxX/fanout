package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 uint64 = 325
    v5 := make(map[string]uint32)
    switch 402 {
    case -33:
        if (!(-2147483648 == 42)) {
            v3 = append(v3, 9223372036854775807)
        } else {
            v6 := GenS1{F0: "a"}
            fmt.Println(v4, v4)
            _ = v6
        }
    case -259:
        v1 -= v1
    }
    var v7 int = -837
    if true {
        v8 := GenS1{F0: "foo bar"}
        var v9 float32 = -2.25
        _ = v8
        _ = v9
    }
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := v5["foo bar"]
        var v13 string = "a"
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = 2147483647
    v3 = append(v3, 128)
    switch "foo bar" {
    case "foo bar":
        _ = len(v2)
    default:
        if true {
            for fi15 := 0; fi15 < 12; fi15++ {
                if true {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: "foo bar"}
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            v18 := GenS2{F0: -1}
            var v19 float32 = (0.1 / 3.14159)
            _ = v18
            _ = v19
        } else {
            v1 -= (v1 * v1)
            v4 -= 286
        }
    }
    v20 := GenS3{F0: (v1 + 18446744073709551615), F1: 1}
    v7 += (v7 ^ v7)
    if true {
        delete(v2, ((("" + "a") + "foo bar") + ""))
        v7 = (v7 << (((v1 % (((v1 * (v1 * 255))) | 1))) & 63))
        v10 -= v10
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int = -290
            switch "a" {
            case "\n\t":
                var v23 bool = (v22 == (v7 / ((v22) | 1)))
                _ = v23
            default:
                fmt.Println(v1, fi21, v1)
            }
            v4 += 0
            _ = fi21
            _ = v22
        }
        switch -826 {
        case 32768:
            v1 = 162
        case 127:
            v3 = append(v3, 128)
        case -9223372036854775808:
            v24 := GenS1{F0: "hello"}
            _ = v24
        }
        var v25 bool = false
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
