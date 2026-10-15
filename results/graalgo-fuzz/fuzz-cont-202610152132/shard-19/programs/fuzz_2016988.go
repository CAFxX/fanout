package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int; F1 int16; F2 string }
type GenS5 struct { F0 int; F1 int }
func genFn0() int64 {
    var v4 int = 4294967295
    v5 := GenS1{F0: 2}
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32) int16 {
    switch "a" {
    case "\n\t":
        v6 := GenS2{F0: 128, F1: 4294967295, F2: ""}
        _ = v6
    case "foo bar":
        v7 := GenS2{F0: 42, F1: 626, F2: ""}
        _ = v7
    }
    v8 := GenS2{F0: 245, F1: 568, F2: "a"}
    p0 -= (p0 >> ((2) & 63))
    if (0 <= -1) {
        v9 := GenS2{F0: 25, F1: 4294967295, F2: ("a" + "\n\t")}
        _ = v9
    } else {
        var v10 uint16 = 32767
        fmt.Println(v10, v10)
        var v11 uint = 255
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: ("" + "foo bar"), F1: p0}
    _ = p0
    _ = v12
    _ = v8
    return 32767
}

func main() {
    var v1 int16 = 959
    v2 := make(map[uint32]uint32)
    var v3 int = -1
    switch "a" {
    case "a":
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS4{F0: v3, F1: (v1 + (v1 >> ((1) & 63))), F2: "hello"}
            var v15 int16 = v1
            _ = fi13
            _ = v14
            _ = v15
        }
    case "hello":
        if (65535 <= 471) {
            v16 := []int{v3, (v3 ^ v3), v3}
            _ = v16
        } else {
            switch "hello" {
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                var v17 uint32 = 42
                _ = v17
            default:
                v3 = (v3 % ((-759) | 1))
            }
        }
    case "\n\t":
        var v18 uint16 = 0
        _ = v18
    }
    v19 := GenS5{F0: 9223372036854775807, F1: 897}
    v1 = v1
    if false {
        var v20 string = ((("hello" + (("a" + "\n\t") + (("a" + ("foo bar" + "\n\t")) + "\n\t"))) + "") + "\n\t")
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(v1, v20)
            v20 = v20
            var v22 float32 = 1.5
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        switch v3 {
        case -548:
            var v23 uint64 = 615
            _ = v23
        default:
            var v24 uint64 = 378
            _ = v24
        }
        fmt.Println(v3)
        v25 := GenS3{F0: "foo bar", F1: -9}
        _ = v25
    }
    v3 -= (^(v3))
    switch ("\n\t" + "\n\t") {
    case "hello":
        fmt.Println(v1, v3)
    case "\n\t":
        var v26 uint8 = 42
        _ = v26
    case "":
        v27 := GenS5{F0: (^(v3)), F1: v3}
        _ = v27
    }
    v1 = (v1 + v1)
    v28 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
}
