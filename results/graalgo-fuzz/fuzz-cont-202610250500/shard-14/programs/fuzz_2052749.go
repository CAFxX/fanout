package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 float32 = 1e10
    var v2 int64 = -1
    var v3 uint = 1
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    var v6 uint16 = 141
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 += 1
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 -= int64(v6)
            _ = fi8
        }
        _ = fi7
    }
    if false {
        var v9 uint = v3
        v9 = v3
        if false {
            v9 = v9
        }
        _ = v9
    } else {
        if (!(!(127 != 1))) {
            var v10 bool = false
            var v11 uint16 = v6
            v12 := GenS1{F0: 896, F1: uint32(v6)}
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 uint32 = 65535
        _ = v13
    }
    v14 := []int{}
    var v15 int32 = 0
    v16 := GenS1{F0: 718, F1: 796}
    switch ("foo bar" + "") {
    case "hello":
        v6 -= (^(v6))
    case "foo bar":
        switch 32767 {
        case 4294967295:
            v6 -= (v6 ^ (^(v6)))
        }
    case "\n\t":
        v17 := GenS1{F0: 2147483648, F1: 42}
        _ = v17
    }
    var v18 uint = v3
    var v19 int = 818
    if (false || false) {
        delete(v5, 0)
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS2{F0: 65535, F1: 42}
            v23 := GenS2{F0: v3, F1: v2}
            _ = fi21
            _ = v22
            _ = v23
        }
        if (!false) {
            v24 := GenS1{F0: 2147483647, F1: uint32(v18)}
            _ = v24
        } else {
            var v25 float64 = 0.1
            _ = v25
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
