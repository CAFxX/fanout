package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    switch 865 {
    case -937:
        v4 := GenS1{F0: 4294967295, F1: "", F2: 0}
        _ = v4
    default:
        switch (("" + "foo bar") + ("a" + (("foo bar" + (("" + (("" + (("foo bar" + (("" + "hello") + ("hello" + "a"))) + "\n\t")) + "hello")) + "\n\t")) + "a"))) {
        case "foo bar":
            var v5 float64 = 3.14159
            _ = v5
        case "":
            var v6 uint32 = 42
            _ = v6
        case "hello":
            fmt.Println(v3)
        }
    }
    fmt.Println(v3)
    switch 32768 {
    case 4294967295:
        switch 1 {
        case 1:
            v7 := []int{128, 32768, 1, -2147483648}
            _ = v7
        case -723:
            fmt.Println(v3)
        }
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            v3 = v3
            fmt.Println(v3, v3)
            for fi9 := 0; fi9 < 1; fi9++ {
                if true {
                    var v10 int16 = -32768
                    _ = v10
                }
                v11 := GenS2{F0: 2, F1: 42}
                _ = fi9
                _ = v11
            }
            _ = fi8
        }
    }
    v12 := []int{631}
    var v13 float64 = 1e10
    switch "foo bar" {
    case "hello":
        switch "a" {
        case "a":
            v14 := GenS3{F0: 2147483648, F1: 4294967295}
            _ = v14
        case "hello":
            if true {
                _ = len(v2)
                var v15 uint32 = 0
                var v16 uint64 = 18446744073709551615
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: 65535, F1: "foo bar", F2: 18446744073709551615}
                var v18 uint8 = (v3 * (v3 & v3))
                if (32767 == 65535) {
                    v19 := GenS3{F0: 372, F1: 1}
                    var v20 int32 = -2147483648
                    v21 := v2[-9223372036854775808]
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = v17
                _ = v18
            }
        default:
            var v22 int = 42
            _ = v22
        }
    case "foo bar":
        _ = len(v1)
    case "":
        v3 = v3
    default:
        v13 += 1.5
    }
    switch 32767 {
    case 127:
        var v23 string = "a"
        _ = v23
    }
    v13 = 1e10
    if (230 >= 2) {
        v24 := GenS3{F0: 416, F1: 398}
        v13 -= (((v13 * 0.1) - 0.1) + (0.0 * -2.25))
        var v25 int = 128
        _ = v24
        _ = v25
    }
    fmt.Println(v13, v3)
    var v26 float32 = 1e10
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            v13 += v13
            _ = fi28
        }
        v13 += float64(((1e10 / v26) / 3.14159))
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
}
