package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() uint64 {
    v6 := []int{}
    var v7 uint8 = 53
    var v8 int16 = 42
    v6 = append(v6, 319)
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int, p1 int) int64 {
    v9 := []int{}
    fmt.Println(p0, p0)
    p0 -= p0
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int16 = 585
        fmt.Println(fi10, fi10)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 104
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    var v3 int64 = 42
    var v4 uint8 = uint8(v2)
    var v5 int32 = -786
    if false {
        v12 := GenS1{F0: v4, F1: v5}
        _ = v12
    }
    v13 := GenS1{F0: v4, F1: (v5 ^ -2147483648)}
    switch "" {
    case "foo bar":
        v14 := GenS2{F0: 4294967295}
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: 0, F1: v5}
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int8 = 100
            _ = fi17
            _ = v18
        }
        v4 -= (v4 >> ((v2) & 63))
        _ = fi15
        _ = v16
    }
    v19 := GenS2{F0: 18446744073709551615}
    switch "\n\t" {
    case "\n\t":
        v20 := v1[2]
        _ = v20
    }
    v2 = v2
    if true {
        switch int(v3) {
        case 127:
            v21 := []int{42, 1, 128, -140}
            _ = v21
        }
        fmt.Println(v5, v4, v2)
        v22 := GenS1{F0: uint8(v2), F1: (-(v5))}
        _ = v22
    } else {
        fmt.Println(v3, v2, v5)
        fmt.Println(v2)
        v2 += (v2 & 255)
    }
    switch 0 {
    case 1:
        v23 := GenS3{F0: false, F1: 32767}
        _ = v23
    case 4294967295:
        var v24 uint16 = 42
        _ = v24
    case 42:
        v5 = v5
    }
    v4 -= (^(v4))
    v2 += (v2 & v2)
    v4 = (v4 ^ 255)
    if false {
        v4 = (v4 / ((42) | 1))
    } else {
        var v25 bool = false
        var v26 uint64 = 0
        v5 += -2147483648
        _ = v25
        _ = v26
    }
    v1[4294967295] = 128
    fmt.Println(v5, v5, v4)
    v3 += (v3 - 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
