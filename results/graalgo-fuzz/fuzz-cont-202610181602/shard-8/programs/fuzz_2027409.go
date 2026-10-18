package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 uint8, p1 bool) uint8 {
    switch 32767 {
    case -2147483648:
        p1 = (127 != -9223372036854775808)
    case 32767:
        var v6 float32 = (1e10 - 1e10)
        _ = v6
    case -284:
        v7 := []int{393, 272, 42, -202}
        _ = v7
    }
    switch "foo bar" {
    case "a":
        var v8 int16 = 42
        _ = v8
    case "hello":
        p0 += (p0 >> ((42) & 63))
    case "":
        var v9 int = 128
        _ = v9
    default:
        v10 := GenS1{F0: 416}
        _ = v10
    }
    v11 := GenS1{F0: 65535}
    v12 := []int{4294967295}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}
func genFn1(p0 uint64, p1 int16, p2 bool) uint {
    fmt.Println(p1, p0, p1)
    var v13 int = 127
    fmt.Println(p0, v13, p2)
    p2 = true
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint = 540
    v3 := []string{"foo bar"}
    var v4 int8 = -79
    var v5 uint64 = 1
    v2 = (v2 & v2)
    for fi15 := 0; fi15 < 11; fi15++ {
        v1 -= -815
        v16 := GenS1{F0: v2}
        _ = fi15
        _ = v16
    }
    if false {
        v17 := GenS1{F0: (v2 + (^(v2)))}
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: (v2 / (((v2 ^ (v2 << (((v2 >> ((v2) & 63))) & 63)))) | 1))}
        _ = fi18
        _ = v19
    }
    fmt.Println(v5)
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS2{F0: ""}
            v5 = (^(v5))
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 float64 = 0.1
                v4 = v4
                v3 = append(v3, "\n\t")
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    case "\n\t":
        v2 -= (v2 % (((v2 - (v2 - (v2 ^ v2)))) | 1))
    case "foo bar":
        v24 := GenS3{F0: 127, F1: 42, F2: genFn1((v5 % ((2) | 1)), 42, (false && false))}
        _ = v24
    }
    switch "a" {
    case "":
        var v25 int32 = 42
        _ = v25
    case "foo bar":
        fmt.Println(v4, v1, v5)
    default:
        if false {
            var v26 float64 = 0.1
            var v27 string = "a"
            _ = v26
            _ = v27
        } else {
            if (v4 >= v4) {
                var v28 int32 = -236
                _ = v28
            } else {
                var v29 int16 = -1
                v30 := GenS4{F0: -2147483648, F1: 206, F2: 255}
                var v31 float32 = (-2.25 + (float32((3.14159 * 1e10)) / 0.1))
                _ = v29
                _ = v30
                _ = v31
            }
            v32 := GenS1{F0: v2}
            _ = v32
        }
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "")
    }
    switch 1 {
    case 87:
        var v33 uint8 = 165
        _ = v33
    }
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
