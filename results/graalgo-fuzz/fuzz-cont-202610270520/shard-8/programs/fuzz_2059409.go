package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0(p0 int16, p1 int64) int16 {
    v7 := GenS1{F0: 127, F1: -60, F2: 142}
    var v8 uint16 = 127
    v9 := GenS2{F0: 348, F1: (p1 << ((4294967295) & 63)), F2: 475}
    p1 += p1
    if (!false) {
        p1 = p1
    } else {
        v10 := GenS2{F0: 127, F1: 42, F2: 2}
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(p1, v8, fi11)
            v8 = 896
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 208
}
func genFn1(p0 int32, p1 int16, p2 uint32) bool {
    switch "foo bar" {
    case "a":
        p2 += 65535
    case "\n\t":
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: -494, F1: 42, F2: 808}
            _ = fi12
            _ = v13
        }
    case "foo bar":
        for fi14 := 0; fi14 < 7; fi14++ {
            if (!false) {
                var v15 float32 = 1.5
                fmt.Println(fi14, p1)
                v16 := []int{}
                _ = v15
                _ = v16
            }
            switch "foo bar" {
            case "":
                if false {
                    switch (fi14 + fi14) {
                    case -932:
                        v17 := GenS1{F0: fi14, F1: -128, F2: 6}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        p0 = -2147483648
                        v19 := GenS3{F0: 32767, F1: 1, F2: (p1 / ((p1) | 1))}
                        fmt.Println(p1)
                        _ = fi18
                        _ = v19
                    }
                } else {
                    v20 := []int{(fi14 - (^(fi14)))}
                    if false {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    _ = v20
                }
            default:
                v21 := []int{}
                _ = v21
            }
            var v22 uint8 = 128
            _ = fi14
            _ = v22
        }
    default:
        p0 -= (p0 | p0)
    }
    fmt.Println(p2, p1)
    v23 := []int{42, 2147483647}
    if (p1 > (p1 * p1)) {
        var v24 string = "\n\t"
        v25 := GenS2{F0: -125, F1: 1, F2: 180}
        p0 += (p0 * (^(p0)))
        _ = v24
        _ = v25
    }
    switch "" {
    case "hello":
        v26 := GenS4{F0: 551, F1: 255}
        _ = v26
    case "":
        var v27 int8 = -20
        _ = v27
    default:
        var v28 float64 = float64(1.5)
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return true
}

func main() {
    v1 := []uint32{2147483647, 2, 252}
    v2 := []int{}
    v3 := make(map[uint32]string)
    v4 := []int{42, 1, -1}
    var v5 int32 = 0
    var v6 int8 = 127
    var v29 uint16 = 438
    var v30 int = 9223372036854775807
    if (!true) {
        var v31 uint16 = v29
        _ = v31
    } else {
        var v32 uint8 = 128
        if ((!false) && false) {
            var v33 uint32 = 2147483647
            v34 := GenS5{F0: v6, F1: 157, F2: -662}
            var v35 int8 = (v6 * -128)
            _ = v33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        fmt.Println(v5, v5, v30)
        _ = fi36
    }
    switch "foo bar" {
    case "a":
        switch "hello" {
        case "\n\t":
            var v37 uint32 = 0
            _ = v37
        default:
            v6 -= int8(v5)
        }
    case "\n\t":
        var v38 int16 = 32767
        _ = v38
    }
    v29 = 65535
    var v39 uint32 = 2147483648
    v40 := GenS5{F0: (v6 * 127), F1: 18446744073709551615, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
