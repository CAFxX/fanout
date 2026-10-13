package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 int16 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint }
func genFn0() int {
    switch 127 {
    case 4294967295:
        v7 := GenS1{F0: 65535, F1: 2, F2: 0}
        _ = v7
    case 32768:
        if false {
            v8 := GenS2{F0: "", F1: 0}
            _ = v8
        }
    }
    v9 := []int{4294967295}
    var v10 uint = 255
    var v11 uint = 42
    _ = v10
    _ = v11
    _ = v9
    return -267
}
func genFn1(p0 uint8, p1 bool, p2 int) int32 {
    v12 := []int{(p2 + (p2 & genFn0())), -1, (p2 ^ genFn0()), p2}
    for fi13 := 0; fi13 < 11; fi13++ {
        if (true && p1) {
            if p1 {
                p2 -= (^(fi13))
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint64 = 2
                v15 = (^(v15))
                var v16 uint16 = 255
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v17 := GenS3{F0: 448, F1: p0}
        v18 := GenS4{F0: 690, F1: "a", F2: 894}
        _ = fi13
        _ = v17
        _ = v18
    }
    var v19 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v19
    return (v19 + 0)
}

func main() {
    v1 := []uint32{2, 123, 65535}
    var v2 float32 = (0.0 / float32(float64(3.14159)))
    v3 := make(map[uint32]uint32)
    v4 := []uint32{0, 2147483647, 42}
    var v5 uint64 = 2
    v6 := make(map[int]bool)
    if (!(!true)) {
        v5 = v5
        if true {
            v1 = append(v1, 65535)
            fmt.Println(v5)
        }
    } else {
        var v20 float32 = ((v2 / (v2 * v2)) - float32((-2.25 - (0.0 / 1e10))))
        _ = v20
    }
    v5 += (v5 / ((v5) | 1))
    var v21 uint = 255
    _ = len(v6)
    v5 -= v5
    for fi22 := 0; fi22 < 17; fi22++ {
        v5 = v5
        _ = fi22
    }
    v2 = v2
    var v23 uint8 = 35
    var v24 uint16 = 0
    switch "hello" {
    case "foo bar":
        if false {
            _ = len(v1)
            v24 = 65535
        } else {
            v25 := GenS5{F0: (v21 << (((v21 % ((978) | 1))) & 63)), F1: int16(v21)}
            _ = v25
        }
    case "\n\t":
        fmt.Println(v24)
    case "hello":
        if ("hello" <= (((("\n\t" + "foo bar") + "foo bar") + ("" + (("\n\t" + "\n\t") + ("foo bar" + "foo bar")))) + "hello")) {
            v26 := GenS2{F0: "", F1: -151}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v23 = v23
                    _ = fi28
                }
                _ = fi27
            }
            v24 = (^(v24))
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS6{F0: v5, F1: (v23 & v23), F2: (v21 | uint(v5))}
                _ = fi29
                _ = v30
            }
        }
    }
    v21 -= v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
