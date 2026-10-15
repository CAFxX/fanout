package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 string; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 uint64 = 68
    var v4 uint32 = 2
    var v5 int = -459
    var v6 uint32 = 585
    switch -1 {
    case 441:
        v4 += (v4 & v6)
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v5 = (-(v5))
        v5 += (fi7 & fi7)
        _ = fi7
    }
    v8 := GenS1{F0: (("foo bar" + ("hello" + "")) + "foo bar"), F1: -9223372036854775808}
    var v9 uint16 = 824
    switch "hello" {
    case "a":
        v4 = v4
    case "\n\t":
        fmt.Println(v3, v9)
    default:
        fmt.Println(v3, v9, v4)
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = (v3 * v3)
        fmt.Println(v9, v5, v6)
        fmt.Println(v5)
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "foo bar":
        v5 = v5
    case "hello":
        switch (("hello" + ("foo bar" + "hello")) + "\n\t") {
        case "\n\t":
            v12 := GenS2{F0: 32767, F1: (v5 << ((255) & 63))}
            _ = v12
        case "":
            v13 := GenS2{F0: -1, F1: v5}
            _ = v13
        case "hello":
            v4 = 42
        default:
            v14 := v1[950]
            _ = v14
        }
    case "a":
        var v15 int64 = -1
        _ = v15
    default:
        var v16 int16 = 1000
        _ = v16
    }
    _ = len(v2)
    if false {
        var v17 int16 = 889
        _ = v17
    } else {
        _ = len(v1)
        if (3.14159 <= (3.14159 + 1e10)) {
            if (false || (("\n\t" < "hello") && false)) {
                var v18 string = "\n\t"
                _ = v18
            }
            v4 += v6
        } else {
            v19 := GenS2{F0: -32768, F1: (v5 + (v5 | v5))}
            v20 := v1[v6]
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = v19
            _ = v20
        }
        if false {
            fmt.Println(v3)
        }
    }
    if false {
        v21 := GenS1{F0: ("" + (("foo bar" + ("foo bar" + "a")) + (("a" + "foo bar") + "a"))), F1: (v5 / ((int(v9)) | 1))}
        _ = v21
    }
    var v22 float32 = (1.5 + 3.14159)
    if (true || (-32 == -128)) {
        v23 := GenS1{F0: "\n\t", F1: (v5 << ((42) & 63))}
        v24 := GenS3{F0: "foo bar", F1: "", F2: 255}
        _ = v23
        _ = v24
    }
    fmt.Println(v3)
    delete(v1, (v4 * v4))
    v25 := GenS1{F0: "", F1: v5}
    v22 -= v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
