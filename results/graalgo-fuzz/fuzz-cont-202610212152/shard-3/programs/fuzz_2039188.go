package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
func genFn0(p0 string, p1 uint8, p2 bool) uint {
    if true {
        p1 += (p1 / ((p1) | 1))
        p2 = p2
    } else {
        if p2 {
            var v5 uint32 = 0
            v6 := []int{-2147483648, 128, -9223372036854775808}
            _ = v5
            _ = v6
        } else {
            var v7 uint64 = 0
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            p1 += p1
            var v10 float64 = 1.5
            var v11 uint64 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        p2 = p2
        _ = fi8
    }
    v12 := []int{}
    fmt.Println(p1, p1, p1)
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 315
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{""}
    var v3 int8 = 119
    var v4 int16 = -32768
    var v13 uint = 98
    var v14 int = 9223372036854775807
    var v15 bool = (1000 < 42)
    switch "hello" {
    case "a":
        switch ("hello" + "") {
        case "a":
            var v16 bool = (42 != 1)
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 10; fi17++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi17
            }
        default:
            if false {
                if (!v15) {
                    fmt.Println(v3, v3, v1)
                } else {
                    if (2147483648 > 891) {
                        v18 := GenS1{F0: 255, F1: uint32(v13), F2: 42}
                        var v19 uint = genFn0((("hello" + "hello") + ("" + "foo bar")), 127, v15)
                        if len(v2) > 0 {
                            v2[0] = ("\n\t" + "\n\t")
                        }
                        _ = v18
                        _ = v19
                    } else {
                        v4 += -1
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v21 := GenS1{F0: 1, F1: 903, F2: 1}
                        _ = fi20
                        _ = v21
                    }
                }
            } else {
                var v22 string = ""
                v23 := GenS1{F0: 2, F1: 453, F2: 109}
                v24 := GenS1{F0: 130, F1: 2147483648, F2: 0}
                _ = v22
                _ = v23
                _ = v24
            }
        }
    case "foo bar":
        v14 -= 818
    case "":
        var v25 bool = (!(!false))
        _ = v25
    }
    v26 := GenS2{F0: 42, F1: 476, F2: 1}
    if ((v13 - 1) != genFn0("\n\t", uint8(v13), (!true))) {
        var v27 bool = ((v4 < (v4 << (((v13 << ((2) & 63))) & 63))) || ((v14 - v14) < 452))
        fmt.Println(v27, v15)
        v28 := GenS2{F0: 1, F1: 65535, F2: 141}
        _ = v27
        _ = v28
    }
    var v29 int64 = -9223372036854775808
    v13 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
