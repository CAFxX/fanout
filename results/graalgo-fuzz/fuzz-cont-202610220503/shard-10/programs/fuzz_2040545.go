package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t"}
    var v3 string = "a"
    v4 := make(map[string]bool)
    var v5 uint16 = 32767
    v6 := GenS1{F0: v5}
    fmt.Println(v3, v5)
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = (^(v5))
        var v8 uint = 18446744073709551615
        _ = fi7
        _ = v8
    }
    if len(v2) > 0 {
        v2[0] = ((v3 + (v3 + v3)) + (v3 + ""))
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int32 = -1
        var v11 uint32 = 146
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 440 {
    case 128:
        v3 = (("\n\t" + "a") + "foo bar")
    case -451:
        switch 603 {
        case 2147483647:
            v5 += 65535
        case -9223372036854775808:
            fmt.Println(v3)
        case 9223372036854775807:
            var v12 float32 = float32(0.0)
            _ = v12
        }
    default:
        var v13 string = v3
        _ = v13
    }
    if true {
        var v14 uint = 255
        _ = v14
    } else {
        var v15 int64 = 1
        fmt.Println(v15, v15)
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint32 = 1
            var v18 uint64 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    if (120 >= 0) {
        v3 = v3
        v2 = append(v2, "\n\t")
        v19 := GenS1{F0: v5}
        _ = v19
    }
    v20 := GenS1{F0: 65535}
    v3 = "a"
    fmt.Println(v3)
    var v21 uint8 = 84
    v22 := GenS1{F0: 524}
    _ = len(v4)
    var v23 float64 = 0.1
    switch 2147483647 {
    case 4294967295:
        v24 := GenS2{F0: 42}
        _ = v24
    case 42:
        v23 = v23
    case -2147483648:
        var v25 int8 = 42
        _ = v25
    default:
        if false {
            v26 := GenS3{F0: -754, F1: v3}
            switch "a" {
            case "foo bar":
                v27 := GenS2{F0: 100}
                _ = v27
            case "\n\t":
                fmt.Println(v21)
            }
            _ = v26
        }
    }
    v3 = v3
    v28 := GenS3{F0: -803, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
