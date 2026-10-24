package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 42
    var v3 string = ""
    var v4 int32 = -864
    v5 := make(map[int]int)
    switch v3 {
    case "hello":
        v3 = ("hello" + "a")
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        v2 = 0
        v2 = (v2 << ((1) & 63))
        _ = fi6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint32 = 2147483647
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v8, v8)
            if (v4 <= (v4 << ((1) & 63))) {
                v10 := GenS1{F0: -32768, F1: (v4 << ((0) & 63))}
                var v11 bool = false
                var v12 uint64 = 2
                v3 = v3
                _ = v10
                _ = v11
                _ = v12
            } else {
                fmt.Println(fi7, v3, fi9)
                var v13 int8 = 106
                _ = v13
            }
            _ = fi9
        }
        var v14 bool = ((v8 << ((4294967295) & 63)) != 0)
        _ = fi7
        _ = v14
        _ = v8
    }
    var v15 int64 = v2
    if true {
        if true {
            v15 = (v2 | (v15 + (v15 + (-(v15)))))
        }
        v16 := GenS1{F0: 1, F1: int32(v15)}
        _ = len(v5)
        _ = v16
    }
    v15 = (v2 * v15)
    var v17 int64 = v15
    v15 -= (v17 << ((255) & 63))
    if false {
        switch v3 {
        case "hello":
            var v18 int = -911
            _ = v18
        case "":
            v19 := GenS1{F0: 42, F1: (v4 - v4)}
            _ = v19
        }
    } else {
        v17 -= (v17 % ((v17) | 1))
    }
    var v20 float64 = ((1.5 * 3.14159) / -2.25)
    var v21 string = ((v3 + ((("\n\t" + "") + v3) + v3)) + v3)
    v4 = (^(v4))
    if (!false) {
        switch 883 {
        case 128:
            var v22 uint64 = 262
            _ = v22
        }
    } else {
        if false {
            v5[32768] = 32767
            switch 395 {
            case 127:
                v4 -= v4
            case 688:
                v23 := GenS2{F0: (v4 + v4), F1: false, F2: 247}
                _ = v23
            default:
                v24 := []int{2147483647, -740}
                _ = v24
            }
        } else {
            v25 := []int{-9223372036854775808, 945}
            _ = v25
        }
    }
    var v26 uint16 = 968
    v20 = v20
    fmt.Println(v2, v2)
    v17 += -9223372036854775808
    v3 = (((("hello" + "") + v3) + "foo bar") + (v21 + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
