package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0() string {
    v4 := []int{-475, -2147483648, -1}
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float32 = float32((3.14159 - 0.0))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    switch 127 {
    case 32767:
        var v8 float32 = 3.14159
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 230, F1: "\n\t"}
        if false {
            var v10 int32 = 42
            _ = v10
        }
        var v11 uint8 = 200
        _ = v11
        _ = v9
    }
    _ = v4
    return ("" + "a")
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = ((0.1 - -2.25) <= float64(((float32((((-2.25 * (0.0 * 0.1)) - float64((float32((-2.25 / 1.0)) / 1.0))) - 0.0)) * 1.5) / float32(1e10))))
    var v3 uint64 = 42
    switch "foo bar" {
    case "foo bar":
        if (100 > -196) {
            var v12 bool = v2
            _ = v12
        }
    case "":
        v13 := GenS1{F0: 255, F1: ""}
        _ = v13
    case "\n\t":
        var v14 float32 = (3.14159 + ((0.1 / 1.0) / float32(((0.0 + 1e10) / 0.1))))
        _ = v14
    default:
        fmt.Println(v2)
    }
    switch "\n\t" {
    case "a":
        var v15 float64 = 1.5
        _ = v15
    default:
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := v1["foo bar"]
                if (((^(v3)) < (v3 + v3)) && true) {
                    switch 42 {
                    case -992:
                        v19 := v1["\n\t"]
                        _ = v19
                    case 2147483647:
                        fmt.Println(v3)
                    default:
                        v2 = (((1e10 + 0.0) + float64((0.1 * (0.0 * (float32(3.14159) / 1.0))))) == float64(-2.25))
                    }
                } else {
                    var v20 uint = 1
                    _ = v20
                }
                var v21 int64 = 9223372036854775807
                _ = fi17
                _ = v18
                _ = v21
            }
            var v22 string = ("hello" + ("a" + ("\n\t" + "a")))
            _ = fi16
            _ = v22
        }
    }
    v2 = (2147483647 != -877)
    v3 = 1
    var v23 int32 = 1
    var v24 int64 = 0
    v23 -= 205
    switch "\n\t" {
    case "foo bar":
        switch -99 {
        case 1:
            for fi25 := 0; fi25 < 12; fi25++ {
                v2 = false
                v26 := GenS2{F0: 100, F1: 0}
                v27 := GenS3{F0: (v23 / (((v23 % (((-(v23))) | 1))) | 1))}
                _ = fi25
                _ = v26
                _ = v27
            }
        case 127:
            var v28 uint32 = 201
            _ = v28
        case -2147483648:
            if (!true) {
                var v29 int64 = (v24 ^ (v24 >> ((65535) & 63)))
                _ = v29
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
