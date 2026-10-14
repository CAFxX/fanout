package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint16, p2 uint16) string {
    p0 += (p0 & p0)
    if false {
        var v4 float32 = -2.25
        var v5 int16 = 356
        if (false || true) {
            fmt.Println(p2, v5, v4)
        }
        _ = v4
        _ = v5
    } else {
        var v6 uint32 = 42
        switch ("" + ("a" + "a")) {
        case "foo bar":
            var v7 uint8 = 255
            _ = v7
        case "hello":
            v8 := []int{128, 589, 4294967295}
            _ = v8
        }
        p0 -= 1
        _ = v6
    }
    switch ("" + "hello") {
    case "a":
        var v9 int8 = 0
        _ = v9
    case "hello":
        switch (("hello" + "hello") + "a") {
        case "\n\t":
            var v10 int64 = 637
            _ = v10
        case "":
            p2 = (p2 + p2)
        case "a":
            p1 = p1
        default:
            p2 = p2
        }
    case "foo bar":
        p1 = (p2 & 42)
    default:
        v11 := GenS1{F0: 1, F1: (301 > 4294967295)}
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(p0, p1)
        var v13 float32 = float32(0.0)
        v14 := []int{(^(fi12)), (fi12 % ((0) | 1))}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: p0, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return ("\n\t" + "a")
}

func main() {
    v1 := []string{"foo bar", (("a" + "foo bar") + ("a" + "hello")), ("hello" + "\n\t"), "foo bar"}
    v2 := make(map[string]bool)
    v3 := []uint32{333, 1}
    switch "hello" {
    case "":
        var v16 int16 = 32767
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 int64 = 9223372036854775807
            var v19 float32 = (float32((-2.25 + 1e10)) + 0.0)
            fmt.Println(v18)
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        var v20 int64 = 0
        _ = v20
    }
    v21 := GenS3{F0: 32767}
    var v22 int = 4294967295
    v22 = (v22 << ((0) & 63))
    var v23 int16 = 1000
    v23 = v23
    fmt.Println(v23)
    fmt.Println(v23)
    fmt.Println(v23, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
