package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
func genFn0(p0 bool, p1 uint64, p2 bool) uint8 {
    v7 := GenS1{F0: -72, F1: 52}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 0
        p0 = ((42 >= -128) && p0)
        _ = fi8
        _ = v9
    }
    var v10 bool = false
    v11 := []int{127, 34}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 236
}
func genFn1(p0 int, p1 string, p2 uint16) int {
    if (0 == -837) {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(p2)
            if (375 >= 42) {
                if false {
                    v13 := GenS1{F0: (p0 << ((2) & 63)), F1: 38}
                    var v14 uint = 1
                    _ = v13
                    _ = v14
                }
            }
            var v15 uint8 = genFn0(true, 494, false)
            _ = fi12
            _ = v15
        }
        if (395 > 2) {
            var v16 float32 = (((float32(((3.14159 * (-2.25 / -2.25)) / 3.14159)) / float32(1e10)) / float32(3.14159)) * float32((0.1 / ((1.5 / 1.0) - 1e10))))
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        p2 = p2
        v18 := []int{4294967295, fi17, (p0 | 32767)}
        var v19 uint16 = 551
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return (p0 / (((p0 * (p0 & (p0 | p0)))) | 1))
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 184
    v3 := []uint32{}
    var v4 int16 = -1
    var v5 uint64 = 961
    var v6 uint64 = (v5 + v5)
    fmt.Println(v4, v4, v1)
    v2 += (^(v2))
    var v20 int = -408
    var v21 int = v20
    var v22 int32 = 1
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi23
    }
    if true {
        v24 := GenS1{F0: (v20 | 43), F1: -1}
        _ = v24
    } else {
        switch "" {
        case "":
            v25 := GenS1{F0: v21, F1: 100}
            _ = v25
        }
    }
    var v26 uint32 = uint32(v1)
    v27 := GenS1{F0: (-(v21)), F1: -1}
    for fi28 := 0; fi28 < 10; fi28++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v21 -= 32767
        v29 := GenS1{F0: genFn1(fi28, "\n\t", 910), F1: 42}
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: v2}
    var v31 int8 = 100
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v20 += v20
        _ = fi32
    }
    switch "hello" {
    case "a":
        switch "a" {
        case "foo bar":
            v33 := GenS3{F0: 255, F1: (v21 - 2147483647), F2: 205}
            _ = v33
        case "\n\t":
            var v34 float32 = 0.0
            _ = v34
        case "a":
            for fi35 := 0; fi35 < 9; fi35++ {
                v5 = (v5 << ((uint(v5)) & 63))
                v36 := GenS2{F0: (v2 << (((v2 + 567)) & 63))}
                var v37 uint32 = (^(v26))
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    default:
        var v38 uint8 = v1
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
