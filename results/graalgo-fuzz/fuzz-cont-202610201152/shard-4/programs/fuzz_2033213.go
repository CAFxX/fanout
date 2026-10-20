package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    var v1 int64 = 624
    v2 := []int{116, 42, -231, 259}
    var v3 uint16 = 32767
    v4 := GenS1{F0: 51, F1: 18446744073709551615, F2: v3}
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 = (v3 & (v3 ^ (v3 / (((v3 * 2)) | 1))))
        _ = fi5
    }
    v1 += v1
    fmt.Println(v3)
    var v6 int = 0
    var v7 string = ""
    var v8 string = (v7 + v7)
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 13; fi10++ {
        v8 = v7
        _ = fi10
    }
    v11 := GenS2{F0: 100, F1: v3, F2: v3}
    v3 -= v3
    var v12 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
