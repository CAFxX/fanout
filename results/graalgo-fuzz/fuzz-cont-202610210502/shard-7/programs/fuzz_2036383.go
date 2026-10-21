package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0() int16 {
    var v5 uint64 = 2
    var v6 int8 = 1
    switch "hello" {
    case "hello":
        fmt.Println(v5, v5)
    case "a":
        if ((!false) || (-1 > v6)) {
            switch 42 {
            case -2147483648:
                var v7 int8 = (v6 + v6)
                _ = v7
            }
            fmt.Println(v6, v5)
        }
    case "foo bar":
        switch 95 {
        case 128:
            fmt.Println(v6)
        case -1:
            fmt.Println(v5, v6)
        }
    default:
        v8 := []int{-389, 127}
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 850, F1: false}
        _ = fi9
        _ = v10
    }
    v5 += v5
    _ = v5
    _ = v6
    return 71
}
func genFn1(p0 uint8, p1 int8, p2 int) uint {
    p1 = p1
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(fi11, p0, fi12)
            var v13 float64 = (-2.25 - 0.1)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return 616
}

func main() {
    var v1 uint8 = 133
    var v2 int = -892
    var v3 uint16 = 65535
    v4 := []uint32{2147483647, 4294967295}
    v1 = 32
    v14 := GenS2{F0: (v3 + v3)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v15 := GenS2{F0: (v3 >> ((0) & 63))}
    v1 = v1
    var v16 float32 = (((1.5 * (1.5 + 0.1)) + -2.25) + 0.0)
    v17 := GenS3{F0: 792, F1: 42, F2: 1}
    if false {
        var v18 int64 = -1
        v19 := GenS2{F0: v3}
        _ = v18
        _ = v19
    }
    v16 -= v16
    var v20 float64 = 3.14159
    switch "" {
    case "foo bar":
        var v21 string = ("hello" + "")
        _ = v21
    case "":
        v1 += v1
    case "hello":
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS4{F0: -1, F1: -32768}
            _ = fi22
            _ = v23
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (("a" + "a") == ("a" + "hello")) {
        fmt.Println(v3)
    } else {
        var v24 uint64 = 1
        switch "foo bar" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = 65535
            }
        case "":
            v16 = v16
        case "hello":
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 string = ("foo bar" + "foo bar")
                v26 = v26
                if false {
                    fmt.Println(v2, v20)
                }
                _ = fi25
                _ = v26
            }
        }
        _ = v24
    }
    var v27 float64 = v20
    if true {
        _ = len(v4)
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v2 += (v2 + 128)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
