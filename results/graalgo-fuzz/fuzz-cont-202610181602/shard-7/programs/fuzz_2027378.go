package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0() uint32 {
    var v5 uint8 = 63
    v6 := []int{1, 128, 1, 2147483647}
    var v7 uint16 = 658
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1() uint32 {
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 int8 = 50
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    var v3 bool = false
    v4 := []uint32{2147483648, 2, 447}
    _ = len(v1)
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint = 416
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 255, F1: -333}
    v13 := GenS1{F0: 128, F1: -951}
    var v14 int64 = 9223372036854775807
    fmt.Println(v14, v14)
    var v15 bool = true
    if (2 == 0) {
        var v16 uint16 = 552
        v15 = (0.0 < 1.5)
        fmt.Println(v14)
        _ = v16
    }
    v14 -= v14
    var v17 uint8 = 255
    if len(v2) > 0 {
        v2[0] = 0
    }
    if v15 {
        v18 := GenS1{F0: (v17 + 255), F1: -346}
        _ = v18
    } else {
        delete(v1, 131)
        switch -914 {
        case 758:
            var v19 int32 = -2147483648
            _ = v19
        case 32767:
            var v20 uint64 = 18446744073709551615
            _ = v20
        case 2147483647:
            var v21 int16 = -1
            _ = v21
        }
        var v22 int16 = 42
        _ = v22
    }
    var v23 int16 = 42
    fmt.Println(v17)
    v3 = (!(!(73 <= -30)))
    for fi24 := 0; fi24 < 1; fi24++ {
        v3 = ((v23 / (((v23 - (-(v23)))) | 1)) < (-(v23)))
        for fi25 := 0; fi25 < 15; fi25++ {
            v3 = ((^(v17)) > (v17 * v17))
            _ = fi25
        }
        _ = fi24
    }
    switch 128 {
    case 9223372036854775807:
        v17 -= (v17 << ((65535) & 63))
    case -9223372036854775808:
        if false {
            fmt.Println(v15)
            v23 = v23
        }
    case 0:
        if true {
            var v26 uint16 = 255
            v14 = 9223372036854775807
            _ = v26
        }
    }
    v3 = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
