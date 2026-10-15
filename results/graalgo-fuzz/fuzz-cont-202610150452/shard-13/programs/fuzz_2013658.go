package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0() uint8 {
    v4 := GenS1{F0: 1000}
    v5 := []int{42, 0, -135, 42}
    v6 := GenS1{F0: -253}
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1() uint {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: genFn0(), F1: 851}
    var v10 float32 = 1.5
    fmt.Println(v10)
    _ = v10
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 46
    fmt.Println(v3)
    _ = len(v2)
    v3 += 145
    if true {
        var v11 string = ""
        var v12 float64 = 1.5
        v13 := v1[1]
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := GenS3{F0: 893, F1: 12}
        var v15 uint64 = 1
        _ = v14
        _ = v15
    }
    v16 := []int{2147483647, 42}
    if (1 <= -9223372036854775808) {
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS4{F0: 410, F1: v3}
            _ = fi17
            _ = v18
        }
        v19 := GenS5{F0: 0, F1: 877}
        var v20 int = 32767
        _ = v19
        _ = v20
    }
    switch "hello" {
    case "\n\t":
        switch "\n\t" {
        case "":
            if true {
                var v21 uint64 = 599
                _ = v21
            }
        }
    case "":
        var v22 int64 = 143
        _ = v22
    default:
        var v23 uint = 1
        _ = v23
    }
    v3 += v3
    for fi24 := 0; fi24 < 18; fi24++ {
        v3 += (v3 - (v3 << ((954) & 63)))
        _ = fi24
    }
    v3 = genFn0()
    fmt.Println(v3)
    var v25 float64 = (1e10 + 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
