package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 bool, p1 bool, p2 int) uint32 {
    p1 = (4294967295 <= (-(p2)))
    if p1 {
        p2 += (-(p2))
        if (false && true) {
            v7 := GenS1{F0: -1, F1: 42, F2: 9223372036854775807}
            var v8 uint64 = 0
            _ = v7
            _ = v8
        }
    } else {
        var v9 string = ("" + "a")
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        p0 = false
        var v11 int16 = 1000
        switch fi10 {
        case -1:
            var v12 int = (p2 * fi10)
            _ = v12
        case 128:
            p1 = true
        default:
            if p0 {
                if p0 {
                    v11 = v11
                    p1 = ((0 < 9223372036854775807) && (!p1))
                } else {
                    p0 = p1
                    v11 += -1
                }
                switch ("hello" + "a") {
                case "a":
                    for fi13 := 0; fi13 < 18; fi13++ {
                        switch p2 {
                        case 32767:
                            var v14 string = "hello"
                            _ = v14
                        case 128:
                            p0 = p0
                            v15 := []int{(fi10 ^ p2), (fi13 << ((474) & 63)), 127, fi10}
                            var v16 string = "\n\t"
                            _ = v15
                            _ = v16
                        default:
                            v17 := []int{2147483647, (-(fi10))}
                            _ = v17
                        }
                        _ = fi13
                    }
                case "":
                    v18 := GenS1{F0: -40, F1: 408, F2: -1}
                    _ = v18
                case "foo bar":
                    var v19 int = 42
                    _ = v19
                default:
                    p0 = (0.0 < -2.25)
                }
            }
        }
        _ = fi10
        _ = v11
    }
    var v20 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 1
}
func genFn1(p0 uint) uint {
    v21 := GenS2{F0: 2147483647, F1: 111, F2: 450}
    p0 += (p0 + (p0 ^ 2))
    if (!((p0 * p0) < p0)) {
        p0 += p0
    } else {
        switch ((("" + "\n\t") + "hello") + "hello") {
        case "a":
            v22 := GenS1{F0: -1, F1: 0, F2: 1}
            _ = v22
        case "hello":
            p0 -= (p0 | (p0 + 42))
        }
    }
    p0 = 42
    switch ("a" + "foo bar") {
    case "a":
        var v23 bool = true
        _ = v23
    case "":
        p0 = (p0 | 0)
    default:
        if (1000 != 52) {
            if (((1.5 * -2.25) - (0.0 / float64(0.1))) >= (((0.1 * 3.14159) / 1e10) / 3.14159)) {
                v24 := GenS3{F0: -9223372036854775808, F1: 32767}
                _ = v24
            }
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v21
    return 2
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 4294967295
    var v3 uint8 = 0
    var v4 uint64 = v2
    var v5 uint32 = 4294967295
    var v6 uint32 = 4294967295
    v2 += v2
    v25 := GenS4{F0: 0}
    for fi26 := 0; fi26 < 10; fi26++ {
        if false {
            v1 += (v1 * v1)
        }
        var v27 uint64 = 679
        _ = fi26
        _ = v27
    }
    switch 32768 {
    case 1:
        var v28 float32 = 1.5
        _ = v28
    case -9223372036854775808:
        fmt.Println(v5)
    case -998:
        for fi29 := 0; fi29 < 18; fi29++ {
            if (-2.25 < 0.1) {
                fmt.Println(v3, v1, v2)
            } else {
                v5 -= v6
            }
            v30 := GenS1{F0: (-(v1)), F1: 65535, F2: -959}
            v4 += (^(v2))
            _ = fi29
            _ = v30
        }
    default:
        v4 = v2
    }
    v31 := GenS2{F0: 0, F1: 13, F2: uint16(v5)}
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS1{F0: v1, F1: uint16(v6), F2: 0}
        _ = fi32
        _ = v33
    }
    switch "a" {
    case "":
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 uint64 = 180
            _ = fi34
            _ = v35
        }
    }
    if false {
        v6 += 4294967295
        if (!true) {
            v36 := []int{-764, 42, 387}
            switch ("" + "") {
            case "a":
                v37 := GenS4{F0: 737}
                _ = v37
            }
            var v38 uint = 383
            _ = v36
            _ = v38
        }
    } else {
        var v39 int64 = 1
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 int = 0
            var v42 uint = 255
            _ = fi40
            _ = v41
            _ = v42
        }
        if (32767 > 2) {
            v2 = (v2 ^ (v2 | v2))
        }
        _ = v39
    }
    v43 := GenS3{F0: 1, F1: -2147483648}
    var v44 string = "foo bar"
    v45 := []int{128, 42, 42, 242}
    _ = len(v45)
    v4 = v4
    v46 := GenS5{F0: false, F1: true}
    v2 += v4
    _ = len(v45)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
