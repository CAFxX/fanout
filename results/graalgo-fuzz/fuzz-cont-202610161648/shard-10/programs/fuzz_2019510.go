package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0() int32 {
    var v7 float64 = -2.25
    if ((!false) || false) {
        v8 := []int{}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 float32 = -2.25
            for fi11 := 0; fi11 < 13; fi11++ {
                v10 -= v10
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        v7 += 0.1
    }
    var v12 int32 = 319
    _ = v12
    _ = v7
    return (v12 >> ((65535) & 63))
}
func genFn1(p0 uint) uint16 {
    p0 = 255
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(p0, fi13)
        switch 934 {
        case -330:
            fmt.Println(p0)
        case 712:
            p0 = 0
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            switch "\n\t" {
            case "a":
                var v15 uint8 = 127
                _ = v15
            case "\n\t":
                fmt.Println(fi14, fi13, fi13)
            default:
                var v16 float32 = 3.14159
                _ = v16
            }
            p0 += p0
            p0 += 20
            _ = fi14
        }
        _ = fi13
    }
    var v17 uint32 = 612
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: 2147483647, F1: (v17 / ((v17) | 1)), F2: -76}
        fmt.Println(fi18, p0, v17)
        var v20 int8 = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v17
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    var v3 int64 = 417
    var v4 uint32 = 617
    v5 := make(map[int]bool)
    var v6 float32 = (0.0 / 1.0)
    v21 := GenS2{F0: v3}
    fmt.Println(v2, v3, v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        switch "a" {
        case "":
            v2 -= -128
        }
        _ = len(v1)
        _ = fi22
    }
    fmt.Println(v3, v4, v4)
    v1["\n\t"] = 4294967295
    var v23 int8 = v2
    v24 := []int{1}
    if false {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint32 = v4
            _ = fi25
            _ = v26
        }
        var v27 float32 = v6
        _ = v27
    }
    var v28 float64 = 1.5
    fmt.Println(v28, v3, v3)
    v29 := GenS3{F0: 0, F1: (v3 | (v3 - v3)), F2: 128}
    var v30 uint = 693
    switch "" {
    case "a":
        v31 := GenS2{F0: 775}
        _ = v31
    case "hello":
        v2 += (v2 / ((v2) | 1))
    case "foo bar":
        switch "hello" {
        case "\n\t":
            v32 := GenS4{F0: 0, F1: (^(v23))}
            _ = v32
        case "":
            var v33 string = "a"
            _ = v33
        }
    default:
        var v34 int = 4294967295
        _ = v34
    }
    v3 = (v3 & (v3 >> ((255) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
