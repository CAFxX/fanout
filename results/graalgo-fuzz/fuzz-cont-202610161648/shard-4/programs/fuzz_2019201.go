package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 int8 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 636
    var v3 uint8 = 0
    v4 := make(map[int]string)
    v5 := GenS1{F0: false, F1: 0}
    var v6 int = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 = v3
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(v3)
            v6 += (v6 & (fi8 % ((-2147483648) | 1)))
            _ = len(v4)
            _ = fi9
        }
        v10 := GenS2{F0: v1, F1: 479}
        if true {
            fmt.Println(fi8, v3, v6)
            var v11 bool = (!false)
            _ = v11
        }
        _ = fi8
        _ = v10
    }
    var v12 int8 = 127
    v13 := v4[(-(v6))]
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{}
        switch v13 {
        case "a":
            var v16 uint32 = v2
            _ = v16
        case "":
            switch v13 {
            case "a":
                v17 := GenS3{F0: -1, F1: fi14, F2: 2}
                _ = v17
            default:
                v6 = v6
            }
        case "hello":
            var v18 uint16 = 0
            _ = v18
        }
        v19 := GenS2{F0: (v1 ^ (^(v1))), F1: (v2 % ((0) | 1))}
        _ = fi14
        _ = v15
        _ = v19
    }
    v20 := GenS4{F0: 255}
    v21 := GenS5{F0: int32(v12), F1: (v1 | 788), F2: v12}
    var v22 int = (v6 % (((v6 << ((42) & 63))) | 1))
    switch (("\n\t" + v13) + v13) {
    case "":
        v23 := []int{32767}
        _ = v23
    case "foo bar":
        if false {
            var v24 uint32 = (v2 + v2)
            v13 = v13
            switch v13 {
            case "\n\t":
                fmt.Println(v24)
            }
            _ = v24
        } else {
            v12 += v12
        }
    case "\n\t":
        delete(v4, 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
