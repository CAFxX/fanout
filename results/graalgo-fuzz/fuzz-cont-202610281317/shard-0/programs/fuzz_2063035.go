package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int32 }

func main() {
    var v1 uint = 0
    var v2 int32 = 0
    v3 := make(map[string]string)
    var v4 float32 = (1.5 * 0.0)
    var v5 int32 = v2
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: (fi6 + fi6), F1: (fi6 & -403)}
        v8 := GenS2{F0: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch 128 {
            case -9223372036854775808:
                var v11 int16 = 1
                _ = v11
            case -1:
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v1, v1)
                    case "":
                        var v12 int = (fi10 & fi9)
                        _ = v12
                    default:
                        v13 := []int{(-(fi9)), fi10, (-(fi10))}
                        _ = v13
                    }
                } else {
                    fmt.Println(fi10)
                }
            case 4294967295:
                var v14 float64 = 0.0
                _ = v14
            default:
                fmt.Println(v4)
            }
            var v15 float32 = (((v4 - v4) / ((float32(((0.0 * 1.5) + 1e10)) - 0.0) + v4)) + v4)
            _ = fi10
            _ = v15
        }
        if false {
            fmt.Println(fi9, v1, v2)
            fmt.Println(v2)
            v1 -= (v1 + (v1 << (((v1 ^ 120)) & 63)))
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(fi9, v1, v2)
            var v17 float32 = ((v4 / 1.0) * ((-2.25 * 1e10) - v4))
            _ = fi16
            _ = v17
        }
        _ = fi9
    }
    v5 = (v5 | (v2 + v5))
    v2 -= (v2 % ((-744) | 1))
    v2 -= 1
    fmt.Println(v4, v1)
    v4 -= v4
    v18 := GenS3{F0: 0, F1: v2}
    v19 := GenS2{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
