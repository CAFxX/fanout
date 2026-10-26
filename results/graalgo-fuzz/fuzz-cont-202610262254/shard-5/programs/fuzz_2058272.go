package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 string, p1 uint, p2 uint32) bool {
    for fi5 := 0; fi5 < 15; fi5++ {
        if true {
            p0 = p0
            p0 = ((p0 + (((p0 + p0) + (p0 + "hello")) + "hello")) + (p0 + p0))
            p0 = p0
        }
        v6 := []int{fi5, -2147483648, fi5}
        var v7 int16 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{-843}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}
func genFn1(p0 bool, p1 uint16, p2 int) int64 {
    if p0 {
        switch p2 {
        case 9223372036854775807:
            v9 := GenS1{F0: 100, F1: "foo bar", F2: (p1 * p1)}
            _ = v9
        case 0:
            v10 := []int{p2, (p2 / ((p2) | 1))}
            _ = v10
        default:
            fmt.Println(p2)
        }
        var v11 bool = (0 <= 108)
        switch ("\n\t" + "a") {
        case "hello":
            if (0 != 70) {
                fmt.Println(p0, v11, v11)
                if ((1.5 < ((float64(-2.25) + -2.25) * 1e10)) && ("a" >= "")) {
                    switch "hello" {
                    case "foo bar":
                        var v12 uint32 = 258
                        _ = v12
                    default:
                        var v13 int64 = 9223372036854775807
                        _ = v13
                    }
                    if false {
                        var v14 int64 = int64(p1)
                        p0 = ((v14 >> ((4294967295) & 63)) > -9223372036854775808)
                        v15 := []int{p2, 0, p2, p2}
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 int32 = 1
                    _ = v16
                }
            } else {
                var v17 uint16 = p1
                for fi18 := 0; fi18 < 1; fi18++ {
                    fmt.Println(p2, v17, p0)
                    _ = fi18
                }
                _ = v17
            }
        }
        _ = v11
    }
    p0 = (!false)
    p2 = p2
    if (0 != 1) {
        if (127 != 105) {
            fmt.Println(p0)
            p2 = 32767
        } else {
            var v19 int64 = 559
            var v20 float64 = 0.1
            v21 := GenS1{F0: 42, F1: "foo bar", F2: (^(p1))}
            _ = v19
            _ = v20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []uint32{0, 581, 42}
    var v2 uint = 1
    var v3 uint64 = 224
    var v4 uint16 = 32767
    v1 = append(v1, uint32(v3))
    var v22 uint8 = 92
    v23 := GenS2{F0: true, F1: true, F2: v3}
    switch -913 {
    case 32767:
        if (!true) {
            fmt.Println(v22)
        } else {
            v2 -= 65535
        }
    case 42:
        v24 := GenS1{F0: 0, F1: "hello", F2: v4}
        _ = v24
    case -9223372036854775808:
        v25 := GenS3{F0: genFn1((-9223372036854775808 > -9223372036854775808), 65535, 511), F1: 588, F2: (int8(v22) == 127)}
        _ = v25
    default:
        for fi26 := 0; fi26 < 9; fi26++ {
            switch (fi26 / ((456) | 1)) {
            case 32768:
                v2 -= 255
            case 32767:
                fmt.Println(v3, v4, v22)
            default:
                var v27 int8 = 127
                _ = v27
            }
            _ = fi26
        }
    }
    v3 += (^(v3))
    fmt.Println(v4, v22, v2)
    var v28 uint = (v2 >> ((971) & 63))
    for fi29 := 0; fi29 < 15; fi29++ {
        switch "a" {
        case "":
            v30 := GenS4{F0: 118}
            _ = v30
        }
        var v31 uint8 = 75
        _ = fi29
        _ = v31
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
}
