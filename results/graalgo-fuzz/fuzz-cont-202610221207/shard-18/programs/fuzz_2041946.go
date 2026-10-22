package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 uint) int16 {
    switch 9223372036854775807 {
    case 165:
        var v5 int8 = 42
        _ = v5
    case 609:
        var v6 uint32 = 0
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 730}
    switch 0 {
    case 710:
        var v8 string = "a"
        _ = v8
    case 32768:
        fmt.Println(p0)
    case 4294967295:
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: 42, F1: 2147483648}
            v11 := []int{fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 uint32 = 2
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 816
}

func main() {
    var v1 uint32 = 1
    var v2 int8 = 95
    var v3 int8 = 15
    v4 := []int{-9223372036854775808, -2147483648}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 127
        _ = fi13
        _ = v14
    }
    v3 -= (v2 % ((42) | 1))
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v15 uint = 65535
        _ = v15
    case "hello":
        v16 := GenS1{F0: 42, F1: (^(v1))}
        _ = v16
    case "":
        v17 := GenS1{F0: 1, F1: (v1 / ((1) | 1))}
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 bool = true
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint16 = 1
        var v22 int32 = -312
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: uint64(v1), F1: 690}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint64 = 0
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 4294967295, F1: (v1 ^ v1)}
    v1 -= (v1 % ((v1) | 1))
    v27 := GenS2{F0: 5, F1: -797, F2: -32768}
    var v28 uint32 = v1
    if true {
        fmt.Println(v1, v3, v28)
        v4 = append(v4, 32768)
    } else {
        var v29 bool = (true && true)
        v29 = (1 > 0)
        _ = v29
    }
    var v30 uint64 = 2
    if true {
        v31 := GenS2{F0: v30, F1: 1, F2: -162}
        v28 = (v1 % ((v1) | 1))
        var v32 string = "a"
        _ = v31
        _ = v32
    }
    fmt.Println(v28, v30)
    var v33 int64 = 9223372036854775807
    fmt.Println(v30, v1, v1)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}
