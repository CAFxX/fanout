package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint16; F1 string; F2 uint }
type GenS5 struct { F0 uint64; F1 int32; F2 uint }
func genFn0(p0 string, p1 int) int64 {
    var v6 uint64 = 4294967295
    v7 := []int{(-(p1)), 527, p1}
    p1 = (^(p1))
    var v8 uint32 = 713
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 int32, p2 string) string {
    v9 := []int{-1, 802, 128}
    fmt.Println(p1, p2)
    _ = len(v9)
    p2 = ("foo bar" + (((((p2 + "foo bar") + p2) + ((p2 + p2) + p2)) + (p2 + p2)) + ("foo bar" + p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (("a" + "") + ("hello" + p2))
}

func main() {
    var v1 string = ("a" + ("foo bar" + "\n\t"))
    v2 := make(map[uint32]bool)
    var v3 int = 32767
    v4 := make(map[int]string)
    var v5 float64 = 1.5
    switch (^(v3)) {
    case -637:
        var v10 uint = 0
        _ = v10
    case -32:
        v11 := GenS1{F0: (v1 + "\n\t"), F1: 1}
        _ = v11
    default:
        v3 += (v3 >> ((328) & 63))
    }
    v3 = v3
    var v12 int = -929
    v13 := GenS1{F0: "a", F1: 32767}
    var v14 int8 = 100
    if true {
        for fi15 := 0; fi15 < 7; fi15++ {
            v3 = fi15
            _ = fi15
        }
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v14, v12)
        }
    }
    _ = len(v2)
    var v16 int8 = v14
    if (42 >= 1) {
        switch (v3 * 32767) {
        case -1:
            fmt.Println(v14, v12, v5)
        }
        v3 -= (v12 ^ -1)
    }
    v1 = ""
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int8 = v16
        v14 = -123
        _ = fi17
        _ = v18
    }
    switch (v3 + (^(v3))) {
    case 1:
        v19 := GenS1{F0: ("" + ""), F1: 0}
        _ = v19
    case 127:
        _ = len(v2)
    }
    fmt.Println(v1, v16, v12)
    var v20 int8 = v14
    switch v3 {
    case -959:
        if ((!true) || true) {
            for fi21 := 0; fi21 < 19; fi21++ {
                v14 += (v14 * (v14 << ((1) & 63)))
                v22 := []int{}
                _ = fi21
                _ = v22
            }
            v23 := GenS2{F0: 42}
            _ = v23
        } else {
            if false {
                v24 := []int{}
                _ = v24
            }
        }
    case 127:
        var v25 float32 = 1e10
        _ = v25
    case 42:
        v14 = v16
    default:
        switch v1 {
        case "foo bar":
            switch v1 {
            case "foo bar":
                v14 = 127
            }
        case "\n\t":
            v26 := GenS3{F0: (v16 | 20), F1: -1}
            _ = v26
        case "":
            v27 := GenS4{F0: 1, F1: (v1 + "a"), F2: 42}
            _ = v27
        default:
            v28 := GenS5{F0: 688, F1: 147, F2: 0}
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
