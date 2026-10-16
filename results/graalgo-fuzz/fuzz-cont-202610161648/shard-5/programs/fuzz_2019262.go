package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint8; F2 bool }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint; F1 int64 }

func main() {
    var v1 bool = (true || ((!false) || ((-1 > -87) || true)))
    v2 := make(map[uint32]int)
    var v3 int32 = 2147483647
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v1)
        v3 += v3
        for fi7 := 0; fi7 < 6; fi7++ {
            v1 = v1
            _ = fi7
        }
        _ = fi6
    }
    if v1 {
        var v8 uint16 = 570
        delete(v2, 2147483647)
        var v9 uint = 882
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 77
        var v11 string = "\n\t"
        _ = v10
        _ = v11
    }
    fmt.Println(v3, v3)
    var v12 uint8 = 2
    v13 := GenS1{F0: 1, F1: 42, F2: "hello"}
    var v14 float32 = 3.14159
    v5["a"] = (("hello" + "\n\t") + "foo bar")
    v12 = v12
    switch -110 {
    case 2147483647:
        v15 := GenS2{F0: 18446744073709551615}
        _ = v15
    case -684:
        switch 127 {
        case 0:
            if false {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := v5["hello"]
                    _ = fi16
                    _ = v17
                }
            } else {
                fmt.Println(v12)
                fmt.Println(v1, v14)
            }
        }
    default:
        v14 -= float32((1.5 - (float64(-2.25) / 3.14159)))
    }
    if false {
        switch "foo bar" {
        case "hello":
            v1 = v1
        case "\n\t":
            var v18 int16 = 42
            _ = v18
        }
        v19 := GenS3{F0: 128, F1: 2, F2: v1}
        _ = v19
    } else {
        v20 := GenS4{F0: 0, F1: 42}
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := []int{}
        _ = fi21
        _ = v22
    }
    switch 0 {
    case -74:
        v23 := GenS5{F0: 4294967295, F1: 578, F2: 42}
        _ = v23
    case -9223372036854775808:
        switch -937 {
        case 128:
            for fi24 := 0; fi24 < 5; fi24++ {
                for fi25 := 0; fi25 < 13; fi25++ {
                    switch "a" {
                    case "":
                        var v26 uint = 18446744073709551615
                        _ = v26
                    case "hello":
                        v12 += (v12 - 42)
                    }
                    var v27 int8 = 0
                    _ = fi25
                    _ = v27
                }
                v1 = v1
                _ = fi24
            }
        default:
            fmt.Println(v12)
        }
    case 855:
        v28 := v2[42]
        _ = v28
    default:
        v12 += v12
    }
    delete(v4, 231)
    v29 := v2[649]
    v30 := GenS6{F0: 255, F1: 0}
    v3 = v3
    var v31 int = v29
    var v32 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
