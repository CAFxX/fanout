package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 float64 = 1.5
    var v4 int64 = 42
    v5 := GenS1{F0: 32767}
    var v6 int8 = -1
    if (!v1) {
        v7 := GenS2{F0: (-(v4))}
        _ = v7
    }
    v8 := GenS3{F0: 4294967295, F1: -143}
    v9 := GenS4{F0: 1}
    var v10 uint = 255
    v1 = (!false)
    fmt.Println(v4)
    v11 := []int{-502, 2147483647}
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: v4}
        v3 -= -2.25
        _ = fi12
        _ = v13
    }
    var v14 int32 = 1
    fmt.Println(v10, v4, v3)
    if v1 {
        v15 := GenS2{F0: (v4 / ((42) | 1))}
        v1 = false
        if true {
            for fi16 := 0; fi16 < 7; fi16++ {
                switch "\n\t" {
                case "foo bar":
                    var v17 uint16 = 2
                    _ = v17
                case "hello":
                    v10 = (v10 / ((65535) | 1))
                case "\n\t":
                    switch fi16 {
                    case 9223372036854775807:
                        _ = len(v11)
                    case 545:
                        v2["hello"] = "\n\t"
                    case -2147483648:
                        fmt.Println(v6, v1, fi16)
                    }
                default:
                    if true {
                        var v18 uint = (v10 + (v10 << ((0) & 63)))
                        fmt.Println(v1, v14)
                        var v19 uint8 = 128
                        _ = v18
                        _ = v19
                    } else {
                        _ = len(v11)
                        var v20 float32 = 1e10
                        var v21 uint64 = 588
                        _ = v20
                        _ = v21
                    }
                }
                v1 = v1
                _ = fi16
            }
            v1 = false
        } else {
            v1 = (1 < 1)
        }
        _ = v15
    } else {
        if (29 != 0) {
            switch 9223372036854775807 {
            case 240:
                switch "foo bar" {
                case "hello":
                    if (v1 && ((2147483647 < 1) || v1)) {
                        v22 := GenS5{F0: 128, F1: 2}
                        fmt.Println(v4)
                        _ = v22
                    } else {
                        _ = len(v2)
                        fmt.Println(v6, v10, v10)
                        v23 := v2[("a" + "a")]
                        _ = v23
                    }
                }
            case 64:
                _ = len(v11)
            default:
                delete(v2, (("\n\t" + ("\n\t" + ("hello" + ("hello" + "\n\t")))) + ""))
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(v14)
                _ = fi24
            }
        }
    }
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
