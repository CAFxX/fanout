package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int16, p1 int16, p2 int8) int8 {
    switch ("\n\t" + ("\n\t" + "foo bar")) {
    case "":
        var v4 uint64 = 2
        _ = v4
    default:
        fmt.Println(p2, p0, p1)
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        p1 -= p0
        _ = fi5
    }
    v6 := GenS1{F0: (p0 % ((p1) | 1)), F1: p0}
    var v7 string = ("\n\t" + "hello")
    if true {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}
func genFn1(p0 uint16) int64 {
    switch 124 {
    case -1:
        if false {
            v8 := []int{-165, 0}
            if true {
                var v9 float64 = (3.14159 - 0.0)
                _ = v9
            }
            var v10 uint64 = 0
            _ = v10
            _ = v8
        }
    case 32767:
        p0 = (p0 & (p0 & p0))
    case 9223372036854775807:
        v11 := []int{}
        _ = v11
    }
    v12 := GenS1{F0: 32767, F1: 0}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int64 = -46
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    v15 := []int{-1, 0}
    _ = p0
    _ = v12
    _ = v15
    return 926
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint64 = 2
        switch "foo bar" {
        case "":
            var v18 int = fi16
            _ = v18
        case "\n\t":
            switch "foo bar" {
            case "a":
                var v19 uint64 = (v17 & v17)
                _ = v19
            default:
                for fi20 := 0; fi20 < 1; fi20++ {
                    v17 += (v17 >> ((4294967295) & 63))
                    v21 := GenS2{F0: true, F1: 42}
                    _ = fi20
                    _ = v21
                }
            }
        }
        fmt.Println(fi16, v17)
        _ = fi16
        _ = v17
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v22 uint64 = 2
    _ = len(v2)
    v23 := GenS2{F0: true, F1: 2147483647}
    if true {
        switch -168 {
        case 9223372036854775807:
            var v24 uint64 = (v22 << ((42) & 63))
            _ = v24
        default:
            v22 += 1
        }
        v22 -= 4294967295
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1[892] = 127
    }
    v25 := v3[32768]
    if true {
        fmt.Println(v22, v22)
        v25 += v25
    }
    var v26 uint16 = 1
    var v27 uint32 = v25
    v27 -= v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
