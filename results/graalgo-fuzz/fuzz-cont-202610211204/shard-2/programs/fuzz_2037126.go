package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 int32; F1 uint16; F2 uint64 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) uint64 {
    if (true || true) {
        if false {
            fmt.Println(p0)
        }
        switch ("\n\t" + (("hello" + ("\n\t" + "\n\t")) + "a")) {
        case "\n\t":
            p1 = 1
        }
    } else {
        var v6 int32 = 0
        p2 -= p2
        _ = v6
    }
    p2 -= 0
    switch "foo bar" {
    case "":
        if (-9223372036854775808 != -541) {
            if true {
                if true {
                    v7 := GenS1{F0: -627, F1: uint16(p0), F2: 0}
                    var v8 uint8 = (p1 >> ((4294967295) & 63))
                    _ = v7
                    _ = v8
                }
                var v9 int64 = 1
                switch 657 {
                case 128:
                    var v10 uint8 = p1
                    _ = v10
                case -1:
                    v11 := []int{3, 4294967295, -1, -660}
                    _ = v11
                case 32767:
                    var v12 int = 42
                    _ = v12
                }
                _ = v9
            }
        } else {
            v13 := GenS1{F0: 9223372036854775807, F1: p2, F2: 0}
            _ = v13
        }
    case "hello":
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 float64 = 3.14159
            _ = fi14
            _ = v15
        }
    case "\n\t":
        for fi16 := 0; fi16 < 2; fi16++ {
            p1 = (p1 * p1)
            var v17 float32 = 0.1
            var v18 uint = 255
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}
func genFn1(p0 int, p1 int64, p2 uint8) uint64 {
    var v19 int = p0
    v19 = p0
    v20 := GenS2{F0: p2, F1: (p1 << ((255) & 63)), F2: 397}
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS1{F0: p1, F1: 903, F2: p2}
        v23 := []int{-592, -1, p0, p0}
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 float32 = float32(0.0)
    v3 := []uint32{2147483647, 815}
    var v4 uint64 = 69
    v5 := make(map[int]bool)
    v1 -= v1
    var v24 int = 4294967295
    switch "" {
    case "a":
        v4 -= v4
    case "foo bar":
        v25 := v5[v24]
        _ = v25
    default:
        fmt.Println(v1, v1, v1)
    }
    v4 += (v4 << ((uint(v4)) & 63))
    var v26 uint16 = v1
    var v27 float64 = 1.5
    for fi28 := 0; fi28 < 14; fi28++ {
        v3 = append(v3, 0)
        v1 = (^(v1))
        v29 := GenS2{F0: 159, F1: 0, F2: (v4 << ((255) & 63))}
        _ = fi28
        _ = v29
    }
    v1 = (v26 << ((674) & 63))
    var v30 int64 = 42
    v31 := GenS3{F0: int16(v30), F1: 29, F2: 1}
    fmt.Println(v30, v4, v24)
    v4 = (v4 / (((v4 ^ (^(v4)))) | 1))
    v32 := GenS4{F0: -606, F1: (v26 << ((91) & 63)), F2: 2}
    v33 := GenS4{F0: -1, F1: (^(v1)), F2: v4}
    var v34 int64 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
