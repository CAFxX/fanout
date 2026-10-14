package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 bool) uint64 {
    switch "\n\t" {
    case "foo bar":
        if p0 {
            var v6 string = ""
            var v7 string = ""
            _ = v6
            _ = v7
        }
    default:
        var v8 int64 = 525
        _ = v8
    }
    v9 := GenS1{F0: ("foo bar" + ""), F1: 1}
    if (0 != 0) {
        v10 := []int{-483, 127, 903}
        _ = v10
    } else {
        fmt.Println(p0)
        var v11 int16 = -672
        _ = v11
    }
    var v12 bool = (!true)
    _ = p0
    _ = v12
    _ = v9
    return 42
}
func genFn1(p0 string) uint64 {
    p0 = p0
    v13 := []int{-761, -282, 137, 569}
    _ = p0
    _ = v13
    return genFn0((!false))
}

func main() {
    v1 := []uint32{}
    var v2 int64 = -9223372036854775808
    var v3 bool = true
    var v4 string = "hello"
    var v5 int64 = v2
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint32 = 891
        _ = fi14
        _ = v15
    }
    v3 = (v3 || v3)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = -2.25
        v4 = (v4 + (v4 + ""))
        v18 := GenS2{F0: 197, F1: 2}
        _ = fi16
        _ = v17
        _ = v18
    }
    v5 += v5
    _ = len(v1)
    v4 = v4
    var v19 int64 = (v5 % (((v5 + v5)) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
