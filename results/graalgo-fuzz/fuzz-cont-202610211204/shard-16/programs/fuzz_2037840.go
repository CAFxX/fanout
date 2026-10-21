package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 int32) uint {
    var v4 int = 9223372036854775807
    v4 = v4
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(p0, fi5, v4)
        _ = fi5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    if true {
        var v8 float32 = 1e10
        v8 = 0.0
        _ = v8
    } else {
        var v9 uint8 = 137
        p0 = (-(p0))
        var v10 uint64 = uint64(v9)
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 130
}
func genFn1() int64 {
    v11 := GenS1{F0: "hello"}
    switch 32767 {
    case 975:
        var v12 uint32 = 1
        _ = v12
    default:
        var v13 uint64 = 869
        _ = v13
    }
    var v14 int32 = -532
    v15 := GenS1{F0: "a"}
    fmt.Println(v14)
    _ = v11
    _ = v14
    _ = v15
    return 917
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{32767, -86, 128}
    var v3 int16 = 0
    if (-1 != -32768) {
        fmt.Println(v3)
    } else {
        var v16 int64 = 1
        v16 = v16
        _ = v16
    }
    switch ((("" + "foo bar") + ("foo bar" + ("foo bar" + "hello"))) + "foo bar") {
    case "foo bar":
        var v17 uint8 = 2
        _ = v17
    case "a":
        fmt.Println(v3)
    }
    var v18 uint64 = 0
    _ = len(v2)
    if false {
        if (0.0 <= (-2.25 * 1.5)) {
            _ = len(v2)
        }
        var v19 bool = (false || true)
        _ = v19
    } else {
        var v20 string = ("a" + "\n\t")
        switch "foo bar" {
        case "hello":
            v21 := GenS1{F0: "a"}
            _ = v21
        default:
            v18 += v18
        }
        _ = len(v2)
        _ = v20
    }
    fmt.Println(v18, v3)
    var v22 float64 = (0.0 * -2.25)
    v3 = (v3 / ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
