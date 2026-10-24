package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int16; F2 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 int32; F2 int16 }
type GenS11 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint16; F1 uint }
func genFn0() uint8 {
    if (0 <= 128) {
        v7 := GenS1{F0: -1}
        v8 := GenS1{F0: -869}
        v9 := GenS2{F0: 1}
        _ = v7
        _ = v8
        _ = v9
    } else {
        if true {
            var v10 float32 = 3.14159
            v11 := []int{2147483647, -429, 32768, -2147483648}
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 string = "a"
            v14 = ((v14 + v14) + "")
            _ = fi13
            _ = v14
        }
        switch "hello" {
        case "\n\t":
            v15 := GenS3{F0: 546, F1: 0, F2: 128}
            _ = v15
        case "a":
            v16 := GenS4{F0: 215}
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := []int{127, fi17, fi12}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                v18 = append(v18, (fi17 ^ (^(fi12))))
                _ = fi17
                _ = v18
            }
        default:
            fmt.Println(fi12)
        }
        _ = fi12
    }
    var v19 float32 = 0.1
    switch "hello" {
    case "a":
        v20 := GenS4{F0: 128}
        _ = v20
    case "foo bar":
        if (-9223372036854775808 == -1) {
            v21 := []int{128, 128}
            _ = v21
        } else {
            var v22 uint8 = 128
            var v23 uint8 = (v22 & 42)
            for fi24 := 0; fi24 < 18; fi24++ {
                if false {
                    v25 := GenS5{F0: 2147483647, F1: 42, F2: "foo bar"}
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v23 = (v23 & (v22 / (((v22 % ((v23) | 1))) | 1)))
                        v19 = (v19 / v19)
                        _ = fi26
                    }
                    _ = v25
                }
                _ = fi24
            }
            _ = v22
            _ = v23
        }
    case "hello":
        v19 -= v19
    default:
        if (0 != 464) {
            var v27 int8 = 49
            _ = v27
        }
    }
    _ = v19
    return 24
}
func genFn1(p0 string, p1 int8) uint64 {
    switch ("a" + p0) {
    case "hello":
        if (true && false) {
            p1 = (p1 | -22)
        }
    case "":
        switch 128 {
        case 42:
            switch (p0 + p0) {
            case "hello":
                fmt.Println(p0, p1)
            default:
                v28 := GenS6{F0: -648}
                _ = v28
            }
        }
    case "\n\t":
        v29 := []int{2147483647, 9223372036854775807, 32768}
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v31 := []int{fi30, (fi30 / ((fi30) | 1)), -248, -9223372036854775808}
        _ = fi30
        _ = v31
    }
    p1 = (p1 / (((p1 >> ((65535) & 63))) | 1))
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 float32 = 0.0
        _ = fi32
        _ = v33
    }
    v34 := GenS7{F0: 895, F1: true}
    _ = p0
    _ = p1
    _ = v34
    return 0
}

func main() {
    var v1 int16 = -796
    var v2 int64 = -969
    var v3 float32 = (0.1 / (0.1 / 1e10))
    v4 := make(map[int]bool)
    var v5 int64 = (v2 / ((v2) | 1))
    var v6 string = "a"
    fmt.Println(v2)
    v35 := GenS8{F0: 65535}
    switch 128 {
    case -2147483648:
        for fi36 := 0; fi36 < 15; fi36++ {
            v1 = (v1 * 0)
            _ = fi36
        }
    default:
        delete(v4, 128)
    }
    fmt.Println(v6)
    v1 = v1
    v37 := []int{}
    var v38 string = "a"
    var v39 uint32 = 4294967295
    for fi40 := 0; fi40 < 16; fi40++ {
        if false {
            v41 := GenS9{F0: -1}
            var v42 uint = 2
            delete(v4, (fi40 | (fi40 % ((-1) | 1))))
            _ = v41
            _ = v42
        }
        v43 := GenS10{F0: int(v39), F1: 437, F2: (-(v1))}
        if (0 > v5) {
            v44 := GenS11{F0: 116, F1: (825 == 18446744073709551615), F2: 127}
            var v45 int8 = int8(v1)
            v45 += v45
            _ = v44
            _ = v45
        } else {
            v46 := v4[int(v1)]
            v3 = float32(float64((v3 * 0.0)))
            var v47 int32 = -1
            _ = v46
            _ = v47
        }
        _ = fi40
        _ = v43
    }
    switch 32767 {
    case -148:
        fmt.Println(v3, v3)
    case -68:
        v3 = float32((-2.25 * 0.1))
    }
    var v48 uint8 = 2
    v49 := GenS12{F0: 222, F1: 103}
    _ = len(v4)
    v39 = (v39 << ((990) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
