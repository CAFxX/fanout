package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16) bool {
    fmt.Println(p0)
    p0 = (p0 % ((p0) | 1))
    v6 := GenS1{F0: 562, F1: p0}
    _ = p0
    _ = v6
    return false
}
func genFn1(p0 int8, p1 uint32, p2 int32) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        switch ("a" + "a") {
        case "":
            v8 := []int{}
            _ = v8
        case "hello":
            v9 := GenS1{F0: 2, F1: -32768}
            _ = v9
        default:
            p0 = p0
        }
        p0 = -1
        p2 = (p2 & p2)
        _ = fi7
    }
    if false {
        fmt.Println(p1, p2)
        v10 := GenS1{F0: 212, F1: 0}
        if false {
            switch 128 {
            case 32768:
                v11 := GenS1{F0: 0, F1: -1}
                _ = v11
            }
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    v1 := []uint32{190, 766}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 0
    v5 := make(map[string]bool)
    v12 := GenS2{F0: 2}
    var v13 int = 128
    for fi14 := 0; fi14 < 11; fi14++ {
        if len(v1) > 0 {
            v1[0] = 165
        }
        _ = fi14
    }
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: 2}
            v17 := GenS3{F0: true}
            fmt.Println(fi15)
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 bool = false
        v4 = (v4 + v4)
        _ = v18
    } else {
        var v19 bool = (0 <= 2)
        _ = v19
    }
    var v20 int16 = 42
    if (!(1 <= 469)) {
        v20 -= v20
        v20 = v20
        if true {
            var v21 uint8 = 255
            var v22 uint16 = 890
            _ = v21
            _ = v22
        } else {
            var v23 uint32 = 0
            if false {
                if (!(true && true)) {
                    var v24 int32 = -1
                    _ = len(v5)
                    switch "a" {
                    case "":
                        v24 += v24
                    }
                    _ = v24
                } else {
                    v25 := GenS4{F0: v4, F1: v23, F2: v4}
                    _ = v25
                }
            } else {
                fmt.Println(v4, v13)
                v26 := GenS3{F0: true}
                var v27 uint8 = 255
                _ = v26
                _ = v27
            }
            _ = v23
        }
    }
    var v28 int8 = v4
    v20 = (v20 >> ((0) & 63))
    var v29 uint32 = 662
    if false {
        v5[("foo bar" + (("a" + "foo bar") + ("hello" + "hello")))] = false
        fmt.Println(v29, v13, v28)
    } else {
        fmt.Println(v20, v28)
        var v30 int = v13
        var v31 uint32 = 516
        _ = v30
        _ = v31
    }
    fmt.Println(v13)
    var v32 uint64 = 4294967295
    v33 := GenS5{F0: v28}
    v34 := GenS3{F0: ((1e10 <= (((0.1 / -2.25) * 0.1) / (3.14159 + ((-2.25 * (1.5 - 1.5)) * float64(-2.25))))) || true)}
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 uint8 = 42
        _ = fi35
        _ = v36
    }
    _ = len(v2)
    v32 += (v32 | (v32 << ((255) & 63)))
    if false {
        var v37 uint16 = 2
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
