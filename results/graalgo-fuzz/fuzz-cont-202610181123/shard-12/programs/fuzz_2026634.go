package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := []int{}
    v2 := []uint32{1}
    v3 := make(map[string]bool)
    v4 := []uint32{821, 2147483648, 1}
    var v5 uint64 = 0
    var v6 int32 = 0
    var v7 int32 = v6
    v8 := GenS1{F0: -76, F1: 84, F2: 306}
    var v9 uint = 2
    switch "" {
    case "":
        var v10 int8 = -128
        _ = v10
    case "hello":
        switch 127 {
        case 32767:
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(fi11, v7)
                _ = fi11
            }
        default:
            v9 = 1
        }
    }
    fmt.Println(v6, v9, v5)
    v12 := GenS1{F0: 42, F1: 2147483648, F2: 1}
    _ = len(v3)
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 string = ("hello" + "\n\t")
        _ = fi13
        _ = v14
    }
    var v15 uint = (^(v9))
    switch "" {
    case "":
        var v16 int64 = -493
        _ = v16
    case "foo bar":
        fmt.Println(v15)
    case "hello":
        var v17 uint32 = 42
        _ = v17
    default:
        v15 = 492
    }
    v18 := GenS1{F0: 127, F1: 42, F2: -688}
    if ("foo bar" != "\n\t") {
        switch 1 {
        case 4294967295:
            v5 = 2
        case 507:
            _ = len(v1)
        default:
            v19 := GenS2{F0: "foo bar", F1: 100}
            _ = v19
        }
        switch 799 {
        case 32767:
            var v20 uint32 = 495
            _ = v20
        case 586:
            switch 212 {
            case -2147483648:
                switch "foo bar" {
                case "":
                    v6 = (v6 / ((v7) | 1))
                case "\n\t":
                    var v21 uint16 = 498
                    _ = v21
                }
            case 9223372036854775807:
                var v22 int = 0
                _ = v22
            case 1:
                if (("foo bar" == ("foo bar" + "a")) || (!false)) {
                    if false {
                        var v23 bool = (0 <= 255)
                        _ = v23
                    }
                } else {
                    v5 = uint64(v9)
                    v24 := GenS3{F0: 765, F1: -16}
                    _ = v24
                }
            }
        }
        var v25 string = ("a" + "foo bar")
        _ = v25
    }
    _ = len(v3)
    v7 += 1
    var v26 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
