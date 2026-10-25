package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    var v5 float32 = 0.1
    var v6 uint16 = 0
    fmt.Println(v6)
    if (!true) {
        var v7 bool = true
        fmt.Println(v5)
        v8 := GenS1{F0: 791}
        _ = v7
        _ = v8
    }
    var v9 int = -699
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int32 = 1
        v2 = 255
        var v12 float64 = (((v1 / float64(-2.25)) + 0.1) - 0.0)
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v6, v9, v5)
    for fi13 := 0; fi13 < 12; fi13++ {
        v6 += (v6 + (^(v6)))
        var v14 string = ("\n\t" + "hello")
        var v15 int16 = -116
        _ = fi13
        _ = v14
        _ = v15
    }
    v6 = (v6 ^ (v6 << ((708) & 63)))
    var v16 int = -1
    v17 := GenS2{F0: 597}
    v18 := GenS1{F0: 18446744073709551615}
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
