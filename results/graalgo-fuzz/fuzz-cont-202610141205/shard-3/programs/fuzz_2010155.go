package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }
type GenS4 struct { F0 int; F1 uint8; F2 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483647
    var v4 float64 = 0.1
    var v5 float32 = -2.25
    var v6 float32 = v5
    v5 += v5
    var v7 int32 = -870
    var v8 int64 = 9223372036854775807
    switch 32768 {
    case 4294967295:
        v9 := GenS1{F0: (v3 % ((0) | 1)), F1: 0, F2: 127}
        _ = v9
    case 0:
        v3 = v3
    }
    var v10 int32 = (v7 + v7)
    var v11 int8 = 1
    v5 = 0.0
    fmt.Println(v10, v10, v10)
    v10 = v10
    v12 := []int{-2147483648, 32767}
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            delete(v1, ("" + "\n\t"))
            v8 = v8
            _ = fi13
        }
        var v14 int = -2147483648
        _ = v14
    }
    if true {
        switch -605 {
        case -844:
            v15 := GenS2{F0: 255, F1: 65535, F2: (v8 * -942)}
            _ = v15
        case 42:
            fmt.Println(v11, v8, v7)
        }
    }
    if false {
        var v16 uint64 = uint64(v3)
        var v17 float32 = v6
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v3 += (^(v3))
        switch "a" {
        case "foo bar":
            switch "a" {
            case "":
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            case "hello":
                if (!((("a" + ("foo bar" + "foo bar")) + ("a" + ("a" + ("" + ("a" + ("\n\t" + "")))))) != "")) {
                    _ = len(v1)
                } else {
                    v19 := GenS3{F0: 0, F1: 0, F2: 753}
                    v20 := GenS2{F0: 127, F1: 988, F2: (v8 % (((v8 / ((v8) | 1))) | 1))}
                    _ = v19
                    _ = v20
                }
            case "\n\t":
                var v21 int64 = v8
                _ = v21
            }
        case "hello":
            fmt.Println(v11)
        default:
            for fi22 := 0; fi22 < 12; fi22++ {
                v3 = v3
                switch "hello" {
                case "hello":
                    v10 = (-(v10))
                }
                _ = fi22
            }
        }
        _ = fi18
    }
    fmt.Println(v7, v3)
    var v23 uint32 = (^(v3))
    v24 := GenS4{F0: 42, F1: 255, F2: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
