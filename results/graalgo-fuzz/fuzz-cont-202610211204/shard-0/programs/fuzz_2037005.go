package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 bool; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 uint32 = 65535
    delete(v3, 4294967295)
    v4 -= v4
    if (!true) {
        fmt.Println(v4)
        switch ("a" + "") {
        case "foo bar":
            v5 := []int{896, 42}
            _ = v5
        case "\n\t":
            delete(v3, 0)
        }
        fmt.Println(v4)
    } else {
        switch ("a" + "foo bar") {
        case "a":
            v4 = v4
        case "hello":
            switch 473 {
            case 1:
                v6 := GenS1{F0: 32767, F1: 1}
                _ = v6
            case 127:
                v4 += v4
            case 32767:
                if false {
                    fmt.Println(v4)
                } else {
                    fmt.Println(v4)
                }
            }
        case "\n\t":
            _ = len(v1)
        default:
            if ((1 != -1) || true) {
                v7 := GenS1{F0: 0, F1: 2}
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS2{F0: -503, F1: (!false), F2: (fi8 | fi8)}
        _ = fi8
        _ = v9
    }
    var v10 float32 = -2.25
    v11 := GenS3{F0: -9223372036854775808}
    for fi12 := 0; fi12 < 4; fi12++ {
        if (true && true) {
            var v13 uint = 42
            _ = v13
        } else {
            v14 := []int{fi12}
            v4 -= (v4 >> ((2) & 63))
            _ = v14
        }
        v3[128] = "hello"
        _ = fi12
    }
    switch ("foo bar" + "") {
    case "\n\t":
        for fi15 := 0; fi15 < 3; fi15++ {
            switch ("foo bar" + ("a" + "")) {
            case "hello":
                var v16 float32 = ((0.0 * (0.0 * 1e10)) / 1.5)
                _ = v16
            case "foo bar":
                v4 += v4
            case "\n\t":
                for fi17 := 0; fi17 < 10; fi17++ {
                    v4 += v4
                    _ = fi17
                }
            default:
                var v18 int16 = 598
                _ = v18
            }
            fmt.Println(fi15, v4)
            _ = fi15
        }
    }
    var v19 string = "a"
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS4{F0: (fi20 / ((fi20) | 1)), F1: true, F2: 562}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        switch (v19 + (v19 + v19)) {
        case "a":
            fmt.Println(fi22, v4)
        default:
            var v23 uint32 = (v4 - (^(v4)))
            _ = v23
        }
        if true {
            v24 := []int{(fi22 % ((fi22) | 1)), fi22}
            _ = v24
        }
        _ = fi22
    }
    v25 := GenS1{F0: -1, F1: 18446744073709551615}
    fmt.Println(v4, v19)
    v26 := []int{-705}
    var v27 uint16 = 735
    v28 := GenS2{F0: 0, F1: (194 != (v27 + v27)), F2: 128}
    for fi29 := 0; fi29 < 16; fi29++ {
        switch v19 {
        case "a":
            if (!false) {
                var v30 uint = 18446744073709551615
                var v31 int32 = -1
                v32 := GenS5{F0: 42, F1: 208, F2: (true && false)}
                _ = v30
                _ = v31
                _ = v32
            }
        default:
            _ = len(v2)
        }
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
