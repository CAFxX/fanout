package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int64; F2 bool }
func genFn0(p0 int8, p1 uint8, p2 uint) uint64 {
    if true {
        p1 += p1
        switch ("\n\t" + "foo bar") {
        case "a":
            var v5 uint = (p2 / (((^(p2))) | 1))
            _ = v5
        case "hello":
            for fi6 := 0; fi6 < 9; fi6++ {
                fmt.Println(p0, p2)
                v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: -2147483648}
                _ = fi6
                _ = v7
            }
        }
    } else {
        v8 := GenS2{F0: p1}
        p2 = p2
        _ = v8
    }
    p0 += (p0 << (((p2 % (((p2 << ((2) & 63))) | 1))) & 63))
    p0 = p0
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 936, F1: -9223372036854775808, F2: (fi10 ^ fi10)}
            switch "a" {
            case "a":
                switch ("hello" + "hello") {
                case "":
                    if (2 < 2) {
                        p0 -= 0
                    }
                case "\n\t":
                    v12 := []int{}
                    _ = v12
                case "foo bar":
                    v13 := []int{fi10}
                    _ = v13
                default:
                    var v14 int64 = 600
                    _ = v14
                }
            }
            _ = fi10
            _ = v11
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := []int{(fi15 << (((p2 << ((2) & 63))) & 63))}
            _ = fi15
            _ = v16
        }
        v17 := []int{}
        _ = fi9
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 int8, p1 int8, p2 uint32) int32 {
    switch ((("" + "hello") + "\n\t") + "a") {
    case "hello":
        p0 -= p0
    case "":
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := []int{}
            _ = fi18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint8 = 127
            p2 = p2
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: 0, F1: -9223372036854775808, F2: -723}
    v23 := GenS2{F0: 99}
    var v24 int = -647
    var v25 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return -605
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 0.1
    var v3 float32 = (((-2.25 / 1.0) - -2.25) * (float32(v2) / 0.1))
    var v4 uint8 = 42
    v26 := []int{-2147483648, 1, 127, 1}
    if (false || false) {
        _ = len(v26)
        for fi27 := 0; fi27 < 3; fi27++ {
            if false {
                v3 += float32((((v2 - ((v2 - 3.14159) + ((0.1 * float64(v3)) * (v2 + (float64(v3) * v2))))) * v2) * (((v2 * float64(v3)) * (1e10 * (v2 - 1e10))) + v2)))
            } else {
                fmt.Println(v3)
                fmt.Println(fi27, v2)
            }
            var v28 int64 = -9223372036854775808
            _ = fi27
            _ = v28
        }
    }
    v1 = ""
    v29 := GenS2{F0: v4}
    v4 = (v4 | 128)
    v30 := GenS1{F0: -737, F1: -1, F2: -552}
    v1 = "hello"
    var v31 uint16 = 1
    if ((0.0 / v2) > (float64(((v3 / -2.25) - v3)) - 1e10)) {
        var v32 bool = true
        fmt.Println(v4, v3, v2)
        _ = v32
    } else {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            v34 := GenS3{F0: v1, F1: -595, F2: (false && true)}
            var v35 uint16 = v31
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
