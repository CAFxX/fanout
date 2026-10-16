package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 int; F2 int64 }
func genFn0(p0 uint16) string {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 bool = true
        var v7 string = "foo bar"
        v8 := GenS1{F0: v7, F1: 1, F2: 4294967295}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 566
    _ = p0
    _ = v9
    return (("foo bar" + "\n\t") + "foo bar")
}

func main() {
    var v1 float32 = (0.1 / 3.14159)
    v2 := make(map[int]int)
    v3 := []uint32{65535, 42, 4294967295}
    var v4 float64 = 3.14159
    if true {
        v10 := GenS1{F0: "hello", F1: 1, F2: 255}
        v4 = v4
        _ = v10
    } else {
        fmt.Println(v4, v1)
        var v11 float64 = ((v4 + ((3.14159 + (v4 - v4)) + v4)) / (v4 - v4))
        switch 42 {
        case 1:
            v12 := GenS2{F0: 973, F1: 2147483647, F2: ("\n\t" + ("\n\t" + "\n\t"))}
            _ = v12
        default:
            v11 = (1.5 / v4)
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: "\n\t", F1: 0, F2: 4294967295}
        v15 := GenS1{F0: "\n\t", F1: 1, F2: 156}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v1, v1)
        _ = fi16
    }
    v1 -= 1e10
    var v17 uint16 = 92
    if (!true) {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v1)
            _ = fi18
        }
        v19 := GenS2{F0: 0, F1: 1, F2: "a"}
        var v20 string = "foo bar"
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint64 = 18446744073709551615
            var v23 int = (fi21 * (fi21 % (((fi21 >> ((911) & 63))) | 1)))
            v1 = (v1 - 0.1)
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS3{F0: 161, F1: fi24, F2: 425}
            _ = fi24
            _ = v25
        }
    }
    var v26 int16 = 42
    v27 := GenS3{F0: 168, F1: 531, F2: -1}
    var v28 uint64 = 1
    v17 -= v17
    v4 = (float64((1e10 * 1e10)) + -2.25)
    var v29 float32 = (((0.0 * 3.14159) / 3.14159) - float32(v4))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
