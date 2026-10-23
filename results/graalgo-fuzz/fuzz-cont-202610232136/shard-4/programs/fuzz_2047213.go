package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int; F2 int }
func genFn0(p0 uint16) uint {
    p0 = p0
    p0 = 255
    p0 = (p0 * p0)
    _ = p0
    return 2
}
func genFn1(p0 bool, p1 uint8, p2 int64) int64 {
    switch "hello" {
    case "hello":
        v6 := []int{}
        _ = v6
    case "":
        if false {
            p2 -= 0
            p0 = true
        } else {
            p1 += (p1 % (((p1 << ((664) & 63))) | 1))
            var v7 uint = 1
            _ = v7
        }
    default:
        var v8 int16 = 1
        _ = v8
    }
    fmt.Println(p0, p0, p2)
    if p0 {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 bool = (1 != uint32(p1))
            if (1000 <= -107) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 int32 = 2147483647
                    var v13 int = fi9
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            fmt.Println(p0, fi9, p2)
            _ = fi9
            _ = v10
        }
        v14 := []int{575, 2147483647, 128}
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int32 = -2147483648
        var v17 int32 = (v16 - v16)
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 * p2)
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = 1e10
    var v3 uint32 = 0
    v4 := make(map[int]uint32)
    var v5 string = ("a" + "hello")
    var v18 int64 = 9223372036854775807
    if false {
        if (v3 != (v3 + 4294967295)) {
            var v19 uint64 = 2
            fmt.Println(v2, v5, v5)
            _ = v19
        } else {
            fmt.Println(v1, v2, v3)
            v20 := []int{460, -1, 0}
            _ = v20
        }
        v5 = (v5 + (v5 + v5))
    } else {
        var v21 uint16 = 65535
        v2 += 1.5
        v21 = (v21 & v21)
        _ = v21
    }
    var v22 int16 = 32767
    fmt.Println(v3, v22)
    var v23 float32 = v1
    fmt.Println(v2, v2)
    if (!(0 <= -92)) {
        var v24 string = v5
        fmt.Println(v1, v5, v1)
        for fi25 := 0; fi25 < 11; fi25++ {
            if true {
                v24 = v5
                var v26 float64 = ((0.0 - v2) + 0.1)
                fmt.Println(v1)
                _ = v26
            }
            var v27 int32 = 2147483647
            var v28 int16 = (v22 >> ((255) & 63))
            _ = fi25
            _ = v27
            _ = v28
        }
        _ = v24
    } else {
        v2 -= (v2 - v2)
        v29 := GenS1{F0: ((v3 + 0) != (v3 >> ((0) & 63)))}
        _ = v29
    }
    var v30 uint16 = 527
    v18 -= int64(v3)
    if (42 == v3) {
        v31 := []int{}
        _ = v31
    }
    delete(v4, 32768)
    switch "hello" {
    case "a":
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 uint8 = 128
            v4[fi32] = 4294967295
            _ = fi32
            _ = v33
        }
    case "hello":
        fmt.Println(v2, v23)
    default:
        for fi34 := 0; fi34 < 4; fi34++ {
            v35 := []int{1, 42, (-(fi34))}
            v5 = ("foo bar" + ("foo bar" + ((v5 + v5) + v5)))
            var v36 uint8 = 128
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    switch 32767 {
    case -2147483648:
        var v37 float32 = 1.5
        _ = v37
    default:
        v38 := GenS2{F0: (v18 % ((v18) | 1)), F1: 128, F2: -282}
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 uint16 = (v30 | v30)
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
