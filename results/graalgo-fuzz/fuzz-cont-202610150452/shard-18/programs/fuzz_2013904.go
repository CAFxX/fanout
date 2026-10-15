package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{127, 32767}
    var v3 int32 = -1
    var v4 float32 = 0.0
    v5 := GenS1{F0: 42, F1: 1}
    v6 := GenS1{F0: 1, F1: 135}
    v1[4294967295] = ""
    v1[65535] = "a"
    if true {
        if false {
            var v7 int32 = v3
            var v8 uint = 0
            _ = v7
            _ = v8
        }
        if true {
            var v9 float32 = (v4 / (v4 / (((v4 - v4) + 0.1) * v4)))
            _ = v9
        }
    }
    var v10 int = -9223372036854775808
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(fi11, v10, fi11)
            _ = fi12
        }
        v13 := GenS2{F0: 42}
        _ = fi11
        _ = v13
    }
    switch v10 {
    case 0:
        switch v10 {
        case -798:
            if false {
                v14 := GenS1{F0: 1, F1: 255}
                var v15 int8 = 0
                _ = v14
                _ = v15
            }
        case 881:
            switch "foo bar" {
            case "":
                fmt.Println(v4, v3, v10)
            case "\n\t":
                var v16 uint = 2
                _ = v16
            default:
                v4 = v4
            }
        case 1:
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int = v10
                var v19 uint8 = 2
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            fmt.Println(v3, v3)
        }
    case -907:
        _ = len(v1)
    default:
        _ = len(v2)
    }
    v1[2147483647] = "hello"
    for fi20 := 0; fi20 < 5; fi20++ {
        v3 -= -1
        var v21 uint = 817
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        switch "foo bar" {
        case "a":
            v10 = (fi22 % ((-1) | 1))
        case "\n\t":
            var v23 bool = (!false)
            _ = v23
        case "foo bar":
            if (!true) {
                _ = len(v1)
            } else {
                v1[773] = "\n\t"
                var v24 uint8 = 2
                var v25 int64 = 42
                _ = v24
                _ = v25
            }
        }
        _ = fi22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS2{F0: 869}
    var v27 uint32 = 483
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
