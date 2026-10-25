package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool }
func genFn0() int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int = fi6
        v8 := GenS1{F0: 18446744073709551615}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 180
    switch "foo bar" {
    case "a":
        fmt.Println(v9)
    case "foo bar":
        v10 := GenS2{F0: 52, F1: ("a" + "hello"), F2: -1}
        _ = v10
    }
    v11 := []int{32767, -1}
    _ = v11
    _ = v9
    return 40
}
func genFn1() uint16 {
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS3{F0: genFn0(), F1: 100, F2: 122}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := []int{(fi15 % ((1) | 1)), (fi15 & (fi15 + (fi15 * 982))), (fi15 | (-(fi15)))}
        v17 := GenS4{F0: ((("foo bar" + "a") + "foo bar") + ("hello" + "")), F1: 411, F2: 18446744073709551615}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v12
    return 42
}

func main() {
    v1 := []uint32{1}
    var v2 int8 = -1
    var v3 int = -914
    var v4 int32 = 1
    var v5 int8 = (v2 / ((42) | 1))
    v18 := GenS2{F0: genFn0(), F1: "a", F2: 234}
    v3 = (-(v3))
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint = 4294967295
        _ = fi20
        _ = v21
    }
    switch ("" + "\n\t") {
    case "foo bar":
        v2 = (v5 + genFn0())
    case "a":
        var v22 string = (("" + "hello") + "foo bar")
        _ = v22
    case "\n\t":
        var v23 float32 = 0.0
        _ = v23
    default:
        v24 := GenS1{F0: 286}
        _ = v24
    }
    v25 := GenS5{F0: 42, F1: v19}
    v1 = append(v1, 688)
    v2 = int8(v3)
    v26 := GenS6{F0: false}
    var v27 int64 = -1
    var v28 uint = 42
    var v29 int = (v3 / ((32767) | 1))
    v30 := GenS1{F0: v28}
    var v31 bool = (!(!false))
    var v32 float32 = float32((-2.25 * (0.1 - 0.0)))
    for fi33 := 0; fi33 < 13; fi33++ {
        if (!v31) {
            for fi34 := 0; fi34 < 6; fi34++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi34
            }
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            v27 = v27
            _ = len(v1)
            fmt.Println(v31, v27)
            _ = fi35
        }
        _ = fi33
    }
    var v36 uint = v28
    fmt.Println(v36, v28, v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
