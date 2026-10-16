package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64) int16 {
    v6 := GenS1{F0: 2147483648, F1: 2}
    v7 := GenS1{F0: 0, F1: 65535}
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint32, p2 uint) int64 {
    v8 := []int{(p0 % ((p0) | 1)), p0, p0, (p0 * -89)}
    var v9 float64 = (1e10 / 3.14159)
    if ((p1 << (((p2 % ((2) | 1))) & 63)) < p1) {
        for fi10 := 0; fi10 < 16; fi10++ {
            p2 = (p2 & p2)
            if (!(!(true && false))) {
                p1 = 316
                var v11 bool = false
                _ = v11
            }
            fmt.Println(p2, p0, v9)
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 858
}

func main() {
    v1 := []uint32{430, 42}
    var v2 string = "\n\t"
    var v3 float64 = -2.25
    v4 := make(map[int]string)
    var v5 uint64 = 791
    v12 := GenS1{F0: 2147483648, F1: 112}
    var v13 int64 = 427
    v14 := GenS2{F0: false, F1: 0, F2: (v13 << ((415) & 63))}
    switch v2 {
    case "hello":
        var v15 int = -386
        _ = v15
    case "\n\t":
        var v16 int64 = (-(v13))
        _ = v16
    case "a":
        fmt.Println(v13, v2, v2)
    default:
        v17 := GenS2{F0: (!false), F1: v13, F2: v13}
        _ = v17
    }
    fmt.Println(v5)
    if true {
        if false {
            switch (v2 + ("" + v2)) {
            case "hello":
                delete(v4, 1)
            default:
                v18 := GenS3{F0: 100}
                _ = v18
            }
            if (true && true) {
                v19 := v4[1]
                _ = v19
            } else {
                fmt.Println(v3)
            }
        } else {
            v3 = float64(float32(v3))
            v13 = (-(v13))
        }
        var v20 int16 = 32767
        var v21 float64 = float64((float32(v3) - (-2.25 + 1.5)))
        _ = v20
        _ = v21
    }
    var v22 float64 = (v3 * float64(((0.1 * 1.5) + 1.5)))
    var v23 int16 = 1
    v5 -= v5
    switch v2 {
    case "foo bar":
        _ = len(v1)
    case "":
        if (false && (!true)) {
            switch 122 {
            case 0:
                var v24 int64 = (v13 >> ((490) & 63))
                _ = v24
            }
        }
    }
    if (true && ((-43 > -720) || (true && true))) {
        v25 := GenS1{F0: 65535, F1: 686}
        var v26 uint16 = 802
        for fi27 := 0; fi27 < 7; fi27++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi27
        }
        _ = v25
        _ = v26
    } else {
        v28 := GenS3{F0: 104}
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        fmt.Println(v2, v5)
        if (("hello" + ((v2 + v2) + v2)) != (("a" + "hello") + "")) {
            var v30 uint = 520
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 float64 = v22
                var v33 uint8 = 1
                _ = fi31
                _ = v32
                _ = v33
            }
            var v34 int16 = (^(v23))
            _ = v30
            _ = v34
        } else {
            fmt.Println(v22)
            var v35 string = v2
            v36 := GenS4{F0: 42}
            _ = v35
            _ = v36
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
