package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
type GenS4 struct { F0 int8; F1 int8; F2 bool }
func genFn0(p0 bool, p1 uint8) int8 {
    p0 = (p1 != (p1 & (p1 | p1)))
    p1 = p1
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 uint = 4294967295
    var v5 uint16 = 65535
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "\n\t" {
        case "a":
            fmt.Println(fi6)
        default:
            var v7 uint16 = v5
            _ = v7
        }
        v8 := GenS1{F0: "", F1: (false && (((("\n\t" + (("" + (("" + "") + ("foo bar" + (("\n\t" + ("hello" + "\n\t")) + "hello")))) + ("" + "a"))) + ("\n\t" + "foo bar")) + ("foo bar" + "")) > ""))}
        var v9 float32 = 1.5
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 int8 = -114
    if true {
        v10 = v10
        v11 := GenS1{F0: "hello", F1: (uint64(v5) != 21)}
        _ = v11
    } else {
        switch -2147483648 {
        case 0:
            switch "foo bar" {
            case "\n\t":
                v5 = (v5 * v5)
            case "a":
                if false {
                    var v12 uint16 = v5
                    _ = v12
                } else {
                    var v13 string = "foo bar"
                    _ = len(v1)
                    _ = v13
                }
            }
        case -577:
            v4 = v4
        case 32768:
            v14 := []int{-9223372036854775808, 127}
            _ = v14
        }
        v15 := GenS2{F0: 9223372036854775807, F1: v5}
        _ = v15
    }
    v10 -= (^(v10))
    switch "a" {
    case "\n\t":
        _ = len(v1)
    case "":
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(fi16)
            if true {
                var v17 uint16 = 159
                fmt.Println(v4, v4, fi16)
                var v18 uint64 = 18446744073709551615
                _ = v17
                _ = v18
            } else {
                v5 += (v5 * (v5 & (v5 + v5)))
                v19 := GenS2{F0: -1, F1: (^(v5))}
                _ = v19
            }
            _ = fi16
        }
    }
    var v20 uint64 = 732
    if true {
        var v21 float64 = 0.1
        _ = v21
    }
    var v22 float32 = float32(-2.25)
    fmt.Println(v4, v10)
    v10 += 42
    v22 -= (v22 + float32(((1.5 * 0.0) + -2.25)))
    v23 := GenS3{F0: v20, F1: -933, F2: 391}
    v3[1] = 2147483648
    v24 := GenS3{F0: v20, F1: 1, F2: 765}
    if true {
        v25 := GenS4{F0: v10, F1: v10, F2: (0 < 0)}
        v10 = 42
        v26 := []int{}
        _ = v25
        _ = v26
    } else {
        v20 = 42
        v27 := GenS4{F0: -1, F1: 62, F2: (-209 <= -190)}
        _ = v27
    }
    switch "foo bar" {
    case "\n\t":
        v28 := GenS3{F0: 0, F1: -2147483648, F2: 772}
        _ = v28
    case "foo bar":
        v29 := []int{}
        _ = v29
    case "":
        v4 = (v4 - 255)
    }
    fmt.Println(v22, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
