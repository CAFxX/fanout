package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 uint16, p1 bool, p2 int16) uint {
    v5 := GenS1{F0: -766, F1: (p0 % (((p0 % ((p0) | 1))) | 1))}
    switch 42 {
    case 1:
        v6 := []int{-654}
        _ = v6
    case 4294967295:
        var v7 float32 = (1.5 + 1.5)
        _ = v7
    }
    switch "" {
    case "\n\t":
        if p1 {
            var v8 uint32 = uint32(p0)
            if (0 != 128) {
                v9 := []int{271, -2147483648, -2147483648}
                if p1 {
                    fmt.Println(p1, p2, p0)
                    var v10 uint = 2
                    _ = v10
                } else {
                    _ = len(v9)
                    p1 = true
                }
                _ = v9
            }
            _ = v8
        }
    case "a":
        var v11 uint = 2
        _ = v11
    }
    fmt.Println(p0, p1, p0)
    switch "hello" {
    case "\n\t":
        var v12 uint64 = 514
        _ = v12
    case "hello":
        p0 = p0
    case "":
        p2 += p2
    default:
        v13 := GenS1{F0: 0, F1: p0}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}
func genFn1(p0 int32, p1 uint32) int8 {
    for fi14 := 0; fi14 < 6; fi14++ {
        p1 = 717
        p1 -= p1
        _ = fi14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int32 = (p0 ^ p0)
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = ((("hello" + "\n\t") + ((("hello" + ("a" + "")) + "a") + "a")) + ("" + "hello"))
    v2 := []uint32{2147483648, 2147483648, 61, 2}
    v3 := make(map[int]uint32)
    var v4 uint64 = 334
    switch 403 {
    case 4294967295:
        for fi17 := 0; fi17 < 14; fi17++ {
            v4 = 4294967295
            _ = fi17
        }
    }
    v18 := GenS1{F0: -1, F1: 32767}
    fmt.Println(v4, v4)
    v1 = v1
    switch "hello" {
    case "":
        var v19 uint8 = 42
        _ = v19
    case "foo bar":
        v1 = v1
    case "a":
        v1 = "foo bar"
    default:
        if len(v2) > 0 {
            v2[0] = 1
        }
    }
    v4 -= 1
    if false {
        for fi20 := 0; fi20 < 17; fi20++ {
            if true {
                v1 = v1
                if false {
                    v21 := GenS2{F0: 356, F1: true}
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 float32 = 1e10
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi20
        }
        var v24 float32 = 1e10
        fmt.Println(v24)
        _ = v24
    }
    var v25 uint8 = 127
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            v1 = v1
            var v28 int = (fi26 ^ (-(fi27)))
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    if (42 != 1) {
        v4 += 539
    } else {
        fmt.Println(v4, v4)
        v4 += (v4 >> ((348) & 63))
        v2 = append(v2, 229)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v29 uint64 = v4
    var v30 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
