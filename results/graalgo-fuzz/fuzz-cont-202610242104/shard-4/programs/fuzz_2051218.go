package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-579, 4294967295}
    v3 := []int{1, 113, 32768, 2147483647}
    var v4 int32 = -2147483648
    fmt.Println(v4)
    for fi5 := 0; fi5 < 3; fi5++ {
        if true {
            v4 = -281
        }
        var v6 int64 = -1
        _ = fi5
        _ = v6
    }
    v4 = 2147483647
    _ = len(v2)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 42, F1: 0, F2: -897}
        v9 := GenS2{F0: (^(v4))}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 555 {
    case 1:
        v4 = (v4 - (-(v4)))
    case 653:
        var v10 uint = 778
        _ = v10
    }
    var v11 int16 = 783
    v12 := GenS1{F0: 0, F1: 1, F2: 0}
    if (!true) {
        v4 = v4
        var v13 bool = ("hello" != "foo bar")
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 bool = v13
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v16 int16 = (v11 + (^(v11)))
    v16 -= (v16 + v16)
    v16 -= (v16 ^ v11)
    v1[65535] = -193
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "hello" {
        case "foo bar":
            var v18 uint64 = 217
            _ = v18
        case "":
            var v19 string = ""
            _ = v19
        case "a":
            var v20 int = 4294967295
            _ = v20
        default:
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int = (fi17 | fi17)
                _ = len(v1)
                fmt.Println(v22, v22, fi21)
                _ = fi21
                _ = v22
            }
        }
        _ = fi17
    }
    v4 = (v4 | (v4 - v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
