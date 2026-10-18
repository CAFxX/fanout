package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 bool, p1 uint16) uint64 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 0, F1: 1}
        switch "\n\t" {
        case "hello":
            p1 = p1
        }
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        p1 = (p1 / ((p1) | 1))
        _ = fi6
    }
    p0 = (p0 || ((!(((p1 << ((657) & 63)) <= p1) || p0)) || p0))
    switch "foo bar" {
    case "a":
        if p0 {
            fmt.Println(p1, p1)
        }
    default:
        var v7 int = 32767
        _ = v7
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() uint64 {
    var v8 int32 = 2147483647
    v8 = -1
    v9 := []int{128}
    var v10 float64 = -2.25
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -96
    var v2 uint16 = 1
    v3 := []uint32{2, 65535, 0, 2147483648}
    var v11 int16 = -300
    for fi12 := 0; fi12 < 14; fi12++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v13 int32 = 42
        _ = fi12
        _ = v13
    }
    _ = len(v3)
    var v14 int32 = 2147483647
    v11 += v11
    v11 = v11
    if false {
        switch -927 {
        case -389:
            v1 += v1
        case 9223372036854775807:
            var v15 uint16 = (v2 << ((1) & 63))
            _ = v15
        }
    }
    var v16 int8 = v1
    switch 726 {
    case 1:
        var v17 uint8 = 243
        _ = v17
    case 32768:
        for fi18 := 0; fi18 < 18; fi18++ {
            v11 = v11
            _ = fi18
        }
    case 367:
        v16 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
}
