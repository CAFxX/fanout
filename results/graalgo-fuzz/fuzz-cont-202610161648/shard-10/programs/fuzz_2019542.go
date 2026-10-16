package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
func genFn0(p0 int) int16 {
    var v6 int = (p0 + 32768)
    if true {
        v7 := GenS1{F0: v6, F1: 400, F2: 757}
        _ = v7
    } else {
        fmt.Println(v6, p0)
    }
    _ = p0
    _ = v6
    return -32768
}

func main() {
    v1 := []int{-161, 127, -1, 4294967295}
    var v2 uint16 = 901
    var v3 uint8 = 109
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    var v8 int64 = 9223372036854775807
    v2 = v2
    fmt.Println(v2)
    var v9 int16 = 1000
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: genFn0(-415), F1: 65535}
        v8 = (v8 / ((-90) | 1))
        v12 := GenS3{F0: "\n\t"}
        _ = fi10
        _ = v11
        _ = v12
    }
    v8 -= (v8 % ((v8) | 1))
    _ = len(v1)
    if false {
        fmt.Println(v2)
        v13 := GenS3{F0: "foo bar"}
        _ = v13
    } else {
        if false {
            var v14 int64 = (v8 - (-(v8)))
            v15 := GenS2{F0: v9, F1: 566}
            _ = v14
            _ = v15
        } else {
            v3 = 239
            var v16 int = 9223372036854775807
            _ = v16
        }
        if true {
            var v17 uint32 = 4294967295
            fmt.Println(v9, v17)
            fmt.Println(v8, v3)
            _ = v17
        } else {
            v3 = (v3 * 174)
        }
    }
    v18 := GenS1{F0: 1, F1: 467, F2: uint64(v2)}
    v2 = v2
    v19 := GenS4{F0: "hello"}
    v5["\n\t"] = 526
    fmt.Println(v9, v2)
    v2 = v2
    var v20 uint64 = 2
    var v21 int8 = 127
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 float32 = float32(1.5)
        for fi24 := 0; fi24 < 9; fi24++ {
            v5[("foo bar" + "foo bar")] = 42
            _ = fi24
        }
        fmt.Println(v20, v21, v20)
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
