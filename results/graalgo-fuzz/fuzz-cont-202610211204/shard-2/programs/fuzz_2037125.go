package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8; F1 string; F2 string }
type GenS3 struct { F0 int8 }
func genFn0() string {
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint64 = v6
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v6, fi9)
        _ = fi9
    }
    switch "hello" {
    case "":
        var v10 uint = 912
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: -846, F1: 179, F2: -32768}
        _ = v11
    default:
        switch 42 {
        case 9223372036854775807:
            fmt.Println(v6)
        case 128:
            if false {
                var v12 float64 = 3.14159
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: (fi13 | fi13), F1: 255, F2: 42}
                    var v15 float64 = v12
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                v6 -= (v6 ^ (v6 << ((1) & 63)))
            }
        }
    }
    var v16 uint8 = 209
    _ = v16
    _ = v6
    return ""
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 255
    v3 := []uint32{0}
    v4 := make(map[string]bool)
    v5 := []int{127, 359, 4294967295}
    v2 += v2
    v2 += v2
    var v17 int64 = 42
    v1 = (v1 | 1)
    v18 := GenS1{F0: 128, F1: 1, F2: -32768}
    var v19 uint = 544
    v20 := GenS1{F0: -107, F1: 0, F2: -932}
    var v21 uint8 = (v2 | (v2 % ((uint8(v19)) | 1)))
    fmt.Println(v17, v2)
    v22 := GenS2{F0: int8(v21), F1: "foo bar", F2: ("hello" + ("a" + ("foo bar" + "a")))}
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v2, v1, v1)
        var v24 uint64 = v1
        _ = fi23
        _ = v24
    }
    switch "a" {
    case "\n\t":
        v4[("foo bar" + "\n\t")] = false
    case "hello":
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint8 = (^(v21))
            _ = fi25
            _ = v26
        }
    case "foo bar":
        var v27 uint64 = v1
        _ = v27
    default:
        _ = len(v4)
    }
    switch "" {
    case "":
        switch 0 {
        case 4294967295:
            v17 += (-(v17))
        case 395:
            var v28 uint64 = v1
            _ = v28
        case -2147483648:
            _ = len(v5)
        }
    case "a":
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int = int(v17)
            v31 := GenS3{F0: -102}
            _ = fi29
            _ = v30
            _ = v31
        }
    default:
        v4["hello"] = true
    }
    delete(v4, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
