package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []int{127, -1, 110}
    var v2 int64 = 0
    var v3 string = "a"
    for fi4 := 0; fi4 < 9; fi4++ {
        v2 = (v2 + v2)
        if false {
            v2 -= 9223372036854775807
            v5 := GenS1{F0: 2}
            v2 = 936
            _ = v5
        }
        switch 1 {
        case 4294967295:
            v3 = v3
        case 675:
            v2 -= (v2 / (((v2 << ((2) & 63))) | 1))
        default:
            v3 = "foo bar"
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        _ = fi6
    }
    v7 := GenS1{F0: 481}
    v2 = v2
    var v8 int8 = 100
    if true {
        v2 = -1
        v3 = v3
        var v9 int16 = -326
        _ = v9
    }
    v10 := GenS1{F0: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 65535}
        if (-685 == v2) {
            v8 = v8
        }
        var v13 uint64 = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = (2147483647 != 42)
    var v15 uint32 = 17
    switch ("foo bar" + (v3 + v3)) {
    case "foo bar":
        var v16 uint64 = 0
        _ = v16
    default:
        var v17 int16 = 884
        _ = v17
    }
    fmt.Println(v3, v2)
    v15 += (^(v15))
    v18 := GenS2{F0: true, F1: 728}
    var v19 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
