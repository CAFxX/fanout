package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8; F2 string }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 uint32 }
func genFn0() int32 {
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 string = ("" + ("a" + (("hello" + "") + ("" + (("a" + "") + ("a" + "foo bar"))))))
            fmt.Println(v7, v7)
            _ = fi6
            _ = v7
        }
    }
    var v8 int32 = -2147483648
    v9 := []int{618}
    v8 = -766
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 float32 = (3.14159 / 1e10)
            _ = fi11
            _ = v12
        }
        var v13 uint32 = 2
        _ = fi10
        _ = v13
    }
    _ = v8
    _ = v9
    return (v8 * (^(v8)))
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    var v4 float32 = -2.25
    v5 := make(map[uint32]bool)
    if false {
        var v14 uint64 = 507
        v1 = v1
        for fi15 := 0; fi15 < 4; fi15++ {
            switch "a" {
            case "foo bar":
                v16 := GenS1{F0: 2, F1: 1, F2: "foo bar"}
                _ = v16
            default:
                fmt.Println(v3, v4, v14)
            }
            var v17 uint64 = 4294967295
            _ = fi15
            _ = v17
        }
        _ = v14
    } else {
        switch (("foo bar" + "\n\t") + "hello") {
        case "foo bar":
            var v18 uint16 = 149
            _ = v18
        case "":
            var v19 int32 = (v1 | 812)
            _ = v19
        }
        v1 -= 0
    }
    switch "\n\t" {
    case "":
        var v20 int = 424
        _ = v20
    case "a":
        v21 := []int{42, -9223372036854775808}
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 int16 = (v3 | (v3 / ((0) | 1)))
            for fi24 := 0; fi24 < 9; fi24++ {
                v23 -= 42
                v25 := GenS2{F0: false}
                v26 := GenS3{F0: v1, F1: 100, F2: ("hello" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t")))}
                fmt.Println(v4, fi22)
                var v27 uint32 = 266
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi22
            _ = v23
        }
    default:
        v3 -= 42
    }
    v28 := []int{}
    var v29 float32 = (v4 + (float32(0.1) + (1e10 * ((v4 / v4) / (-2.25 - (-2.25 + 3.14159))))))
    delete(v2, 145)
    var v30 string = "a"
    v31 := GenS4{F0: (-(v3)), F1: 0}
    var v32 int = 2147483647
    v28 = append(v28, (v32 + int(v3)))
    v1 = v1
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    switch "hello" {
    case "a":
        fmt.Println(v4, v32, v29)
    case "\n\t":
        var v33 int32 = (^(v1))
        _ = v33
    case "hello":
        v34 := GenS4{F0: (v3 & (^(v3))), F1: 504}
        _ = v34
    default:
        _ = len(v28)
    }
    v29 += v29
    v35 := GenS5{F0: 231, F1: v1, F2: 457}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
