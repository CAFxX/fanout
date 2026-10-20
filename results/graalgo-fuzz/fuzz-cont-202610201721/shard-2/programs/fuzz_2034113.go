package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
func genFn0(p0 int32, p1 int32) string {
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 += p1
        v6 := []int{(fi5 / ((61) | 1))}
        _ = fi5
        _ = v6
    }
    if (1e10 >= (-2.25 - 0.0)) {
        var v7 uint8 = 1
        _ = v7
    }
    p1 = (p0 % ((0) | 1))
    switch "\n\t" {
    case "hello":
        v8 := GenS1{F0: 0, F1: 18446744073709551615, F2: -128}
        _ = v8
    case "\n\t":
        if false {
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(p1, p1)
                p0 = (p1 << ((65535) & 63))
                _ = fi9
            }
            v10 := GenS1{F0: 76, F1: 42, F2: 0}
            var v11 int = 32768
            _ = v10
            _ = v11
        } else {
            fmt.Println(p0, p0)
            v12 := []int{}
            fmt.Println(p1)
            _ = v12
        }
    default:
        v13 := []int{}
        _ = v13
    }
    _ = p0
    _ = p1
    return "hello"
}
func genFn1(p0 uint32, p1 int32, p2 uint64) bool {
    var v14 uint64 = p2
    if false {
        v15 := []int{982, -688}
        _ = v15
    } else {
        v16 := []int{2147483647}
        _ = v16
    }
    v17 := GenS1{F0: 502, F1: 372, F2: -43}
    var v18 int32 = (p1 + (p1 >> ((255) & 63)))
    var v19 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 int64 = 1
    var v2 bool = false
    var v3 int = 1
    var v4 int16 = 1
    fmt.Println(v1)
    for fi20 := 0; fi20 < 14; fi20++ {
        switch "\n\t" {
        case "\n\t":
            var v21 int8 = 100
            _ = v21
        }
        v22 := GenS1{F0: 32767, F1: 1, F2: 42}
        _ = fi20
        _ = v22
    }
    v1 = (v1 - (v1 | 1))
    v23 := GenS1{F0: 65535, F1: 4294967295, F2: 1}
    for fi24 := 0; fi24 < 8; fi24++ {
        switch "foo bar" {
        case "\n\t":
            for fi25 := 0; fi25 < 5; fi25++ {
                v2 = v2
                _ = fi25
            }
        case "":
            v26 := GenS1{F0: 1, F1: 4294967295, F2: 100}
            _ = v26
        case "a":
            v27 := GenS1{F0: 255, F1: 223, F2: 0}
            _ = v27
        default:
            switch "\n\t" {
            case "":
                var v28 uint8 = 255
                _ = v28
            case "foo bar":
                var v29 string = "foo bar"
                _ = v29
            case "\n\t":
                var v30 int32 = -1
                _ = v30
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 int8 = 100
                    var v33 bool = (v2 || false)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        }
        _ = fi24
    }
    fmt.Println(v4, v3)
    v34 := []int{v3}
    v34 = append(v34, (v3 + 1))
    v4 -= v4
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
}
