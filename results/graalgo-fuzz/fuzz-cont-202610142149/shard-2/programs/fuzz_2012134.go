package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int }
func genFn0() int8 {
    v7 := []int{-300}
    v8 := GenS1{F0: "a"}
    var v9 uint16 = 222
    v7 = append(v7, 32767)
    if true {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS2{F0: 42}
            v12 := GenS3{F0: 128}
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{1, 0}
    var v2 int16 = 971
    var v3 uint64 = 4294967295
    var v4 bool = true
    var v5 uint8 = 187
    v6 := make(map[string]uint32)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int64 = 42
        var v15 string = "\n\t"
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: (v5 ^ (v5 & v5))}
    v17 := v6[("hello" + "a")]
    v18 := GenS4{F0: v2, F1: 228}
    if v4 {
        v19 := GenS4{F0: (^(v2)), F1: (v5 & v5)}
        v20 := GenS1{F0: ("" + ("a" + "foo bar"))}
        v3 -= (v3 + (v3 << ((663) & 63)))
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint32 = 458
            var v23 float32 = 0.0
            v24 := GenS5{F0: (v3 * v3), F1: (fi21 % ((4294967295) | 1))}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v5 += (^(v5))
    }
    v1 = append(v1, (v17 % (((v17 / ((v17) | 1))) | 1)))
    var v25 bool = false
    v2 = 1000
    for fi26 := 0; fi26 < 5; fi26++ {
        v2 = v2
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
