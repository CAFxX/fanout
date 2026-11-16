package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string; F1 int8; F2 uint32 }
type GenS6 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int16) string {
    v5 := []int{0, 127, 0, -2147483648}
    v6 := GenS1{F0: (false && true), F1: (((p0 ^ (p0 << ((4294967295) & 63))) < p0) && false)}
    switch "" {
    case "hello":
        var v7 int16 = 1000
        _ = v7
    default:
        var v8 int16 = (p0 | (p0 >> ((4294967295) & 63)))
        _ = v8
    }
    var v9 uint = 255
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return "foo bar"
}
func genFn1() uint32 {
    var v10 string = ("" + "foo bar")
    if true {
        var v11 int64 = -987
        v12 := GenS2{F0: (-2147483648 > -2147483648)}
        v11 = 42
        _ = v11
        _ = v12
    }
    v10 = "hello"
    v13 := GenS3{F0: v10}
    v10 = v10
    _ = v10
    _ = v13
    return 97
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-1}
    var v3 int32 = 829
    var v4 int = 935
    fmt.Println(v4)
    v14 := v1[""]
    if true {
        v15 := GenS4{F0: v3, F1: -128}
        v3 = v3
        _ = v15
    } else {
        if true {
            switch "foo bar" {
            case "\n\t":
                v16 := GenS3{F0: "\n\t"}
                _ = v16
            case "a":
                v3 += (v3 * (v3 | (v3 + v3)))
            }
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 int8 = -17
                v19 := GenS5{F0: "foo bar", F1: v18, F2: v14}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint = 18446744073709551615
            _ = v20
        }
        if (0 != 42) {
            switch "a" {
            case "hello":
                var v21 uint32 = v14
                _ = v21
            default:
                v3 = 0
            }
            var v22 string = "\n\t"
            fmt.Println(v3, v3)
            _ = v22
        } else {
            var v23 uint64 = 1
            v4 = -1
            _ = v23
        }
        v4 += 32767
    }
    var v24 int64 = 9223372036854775807
    for fi25 := 0; fi25 < 13; fi25++ {
        v1[""] = v14
        v26 := GenS6{F0: (v3 + v3), F1: 127, F2: 65535}
        _ = fi25
        _ = v26
    }
    var v27 uint = 2
    var v28 int64 = v24
    var v29 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
