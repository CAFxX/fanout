package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0() uint32 {
    switch ("foo bar" + "hello") {
    case "foo bar":
        switch -1 {
        case -63:
            var v5 bool = ("hello" > ((("" + "foo bar") + "") + "a"))
            _ = v5
        default:
            var v6 float32 = -2.25
            _ = v6
        }
    case "hello":
        var v7 uint = 18446744073709551615
        _ = v7
    default:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    var v9 uint16 = 440
    v9 = 32767
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    v3 := make(map[string]bool)
    var v4 uint8 = 255
    switch -44 {
    case 4294967295:
        switch "a" {
        case "hello":
            v10 := []int{-5, 1, -464, -2147483648}
            _ = v10
        case "a":
            var v11 int32 = (v2 << ((18446744073709551615) & 63))
            _ = v11
        case "foo bar":
            var v12 float32 = 1.5
            _ = v12
        }
    case 2147483647:
        var v13 int32 = -1
        _ = v13
    default:
        var v14 int = -2147483648
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    v15 = 1
    if (!false) {
        fmt.Println(v2, v4, v15)
        var v16 uint64 = (v15 | 18446744073709551615)
        _ = v16
    } else {
        v17 := GenS1{F0: v2}
        if (true || true) {
            v2 = (v2 << ((18446744073709551615) & 63))
            if ((v4 | 131) > v4) {
                var v18 int = 2147483647
                _ = v18
            } else {
                v19 := []int{32768, 0, 2147483647}
                v20 := GenS2{F0: 42}
                _ = v19
                _ = v20
            }
        }
        _ = len(v1)
        _ = v17
    }
    v21 := GenS3{F0: 65535, F1: 840, F2: 271}
    v22 := []int{}
    v23 := GenS4{F0: 0, F1: "\n\t"}
    v24 := GenS1{F0: v2}
    switch "" {
    case "\n\t":
        v4 = (v4 + v4)
    case "a":
        if len(v22) > 0 {
            v22[0] = 4294967295
        }
    case "foo bar":
        var v25 string = "hello"
        _ = v25
    default:
        if true {
            var v26 string = ""
            var v27 uint8 = uint8(v15)
            v27 -= (v4 ^ 0)
            _ = v26
            _ = v27
        }
    }
    if true {
        fmt.Println(v15, v15, v15)
        for fi28 := 0; fi28 < 10; fi28++ {
            _ = len(v1)
            v4 -= 128
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS3{F0: 917, F1: 2, F2: 1}
                fmt.Println(v15, fi28)
                v31 := GenS5{F0: 0, F1: (v2 * v2)}
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
