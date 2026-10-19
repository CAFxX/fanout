package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 int16) uint64 {
    var v6 uint8 = 2
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    var v3 uint = 91
    var v4 uint32 = 42
    var v5 int32 = -2147483648
    var v7 int8 = -87
    if false {
        v3 = (^(v3))
        v7 = v7
    }
    var v8 int = 1
    v8 -= (v8 / (((v8 << ((672) & 63))) | 1))
    if true {
        v2 = append(v2, 0)
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int16 = 42
        var v11 uint8 = 1
        var v12 bool = true
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    switch ("" + "foo bar") {
    case "foo bar":
        switch -9223372036854775808 {
        case -307:
            v7 += 80
        case 128:
            switch "a" {
            case "":
                switch "a" {
                case "\n\t":
                    v5 = (v5 * 42)
                }
            case "hello":
                v13 := GenS1{F0: 2, F1: 730}
                _ = v13
            }
        default:
            var v14 uint16 = 42
            _ = v14
        }
    }
    delete(v1, v8)
    var v15 int32 = -1
    v3 -= (v3 + v3)
    v16 := GenS2{F0: v7, F1: "a"}
    switch -9223372036854775808 {
    case 1:
        if ((^(v3)) > v3) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v17 int32 = (v15 | (v15 & (v5 / ((v5) | 1))))
            v18 := GenS3{F0: 2, F1: 366}
            _ = v17
            _ = v18
        }
    }
    if (1 >= 102) {
        v19 := GenS4{F0: ((v5 / ((1) | 1)) < (^(v15)))}
        _ = v19
    }
    var v20 float64 = (1.5 * -2.25)
    v4 += v4
    var v21 int16 = -1
    var v22 int32 = (v5 + (v5 >> (((v3 & 18446744073709551615)) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
