package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() uint {
    var v4 uint64 = 248
    v4 -= v4
    v4 -= (^(v4))
    v4 -= v4
    _ = v4
    return 1
}
func genFn1(p0 bool, p1 string, p2 int16) uint16 {
    v5 := []int{}
    var v6 int = 4294967295
    var v7 float64 = float64(-2.25)
    var v8 int64 = -9223372036854775808
    if len(v5) > 0 {
        v5[0] = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 107
}

func main() {
    v1 := []string{"foo bar", "foo bar", (("" + "\n\t") + "foo bar")}
    v2 := make(map[string]uint32)
    var v3 int8 = -108
    var v9 uint8 = 151
    for fi10 := 0; fi10 < 10; fi10++ {
        v3 = 100
        v11 := GenS1{F0: 1, F1: 544}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            switch -288 {
            case -2147483648:
                var v14 float64 = ((1.5 * ((1.5 + 3.14159) / (float64(-2.25) / 1.0))) + 0.1)
                _ = v14
            default:
                v3 = v3
            }
            switch ((("hello" + ("a" + "\n\t")) + ("hello" + ("\n\t" + "a"))) + "a") {
            case "\n\t":
                switch 42 {
                case 128:
                    if len(v1) > 0 {
                        v1[0] = ("a" + "foo bar")
                    }
                case -357:
                    v15 := v2[("hello" + "")]
                    _ = v15
                default:
                    var v16 uint32 = 65535
                    _ = v16
                }
            case "a":
                v9 = v9
            default:
                v9 = 5
            }
            _ = fi13
        }
        if true {
            v9 = (^(v9))
            v17 := GenS2{F0: uint64(v9)}
            v18 := GenS1{F0: -2147483648, F1: 4294967295}
            _ = v17
            _ = v18
        } else {
            switch 42 {
            case 2147483647:
                var v19 uint32 = 65535
                _ = v19
            case 685:
                v20 := GenS3{F0: -20}
                _ = v20
            }
            _ = len(v1)
        }
        _ = fi12
    }
    switch 32767 {
    case -179:
        for fi21 := 0; fi21 < 1; fi21++ {
            v9 += 175
            fmt.Println(v9)
            var v22 uint = 4294967295
            _ = fi21
            _ = v22
        }
    case 9223372036854775807:
        v23 := GenS1{F0: 791, F1: 776}
        _ = v23
    default:
        v3 += (v3 ^ v3)
    }
    v3 = (v3 / (((v3 >> ((516) & 63))) | 1))
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 + (v3 & -128))
    default:
        var v24 int8 = (v3 >> ((42) & 63))
        _ = v24
    }
    var v25 string = ""
    var v26 string = (v25 + (v25 + "hello"))
    v25 = "\n\t"
    v26 = "hello"
    var v27 int32 = 53
    v9 = v9
    v25 = v25
    v27 = v27
    if len(v1) > 0 {
        v1[0] = (v26 + ("a" + "hello"))
    }
    _ = len(v2)
    v28 := v2[v26]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
