package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32) int {
    var v6 uint8 = 0
    var v7 bool = (false || (-180 <= -339))
    fmt.Println(v6)
    v6 += 24
    _ = p0
    _ = v6
    _ = v7
    return 32768
}
func genFn1(p0 int16, p1 bool, p2 int8) uint64 {
    v8 := []int{32767}
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "hello" {
        case "":
            var v10 uint32 = 348
            _ = v10
        default:
            v11 := GenS1{F0: "hello", F1: (("foo bar" + "hello") + ("hello" + "hello")), F2: 2}
            _ = v11
        }
        var v12 uint8 = 99
        var v13 uint16 = 42
        _ = fi9
        _ = v12
        _ = v13
    }
    p0 = (p0 / (((p0 - (p0 << ((42) & 63)))) | 1))
    p2 = p2
    p0 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    var v1 int8 = 42
    var v2 float32 = (3.14159 * float32((((float64(0.0) / 3.14159) + ((0.0 / 1.0) * 3.14159)) + float64(3.14159))))
    var v3 float32 = -2.25
    v4 := []uint32{4294967295, 377, 2147483647, 42}
    v5 := make(map[int]bool)
    if len(v4) > 0 {
        v4[0] = 65535
    }
    v2 = v3
    fmt.Println(v3, v3)
    var v14 bool = true
    v15 := GenS1{F0: "", F1: "", F2: 224}
    fmt.Println(v14, v14, v3)
    if ("a" != "") {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS2{F0: 1}
            v2 = v3
            v3 = (v2 + (v2 + 1.5))
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            switch "hello" {
            case "\n\t":
                var v19 int64 = int64(v1)
                _ = v19
            case "foo bar":
                v14 = false
            case "a":
                v20 := GenS2{F0: 2}
                _ = v20
            }
            v21 := v5[fi18]
            v22 := GenS1{F0: "hello", F1: "a", F2: 127}
            _ = fi18
            _ = v21
            _ = v22
        }
        var v23 uint64 = 220
        _ = v23
    }
    v24 := GenS1{F0: ("foo bar" + "foo bar"), F1: (("foo bar" + ("\n\t" + "a")) + (("foo bar" + "\n\t") + "foo bar")), F2: 42}
    switch "a" {
    case "":
        var v25 uint16 = 772
        _ = v25
    case "hello":
        fmt.Println(v3, v14, v14)
    case "\n\t":
        fmt.Println(v2, v3, v1)
    default:
        v26 := GenS1{F0: "", F1: ("" + "a"), F2: 0}
        _ = v26
    }
    v3 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
