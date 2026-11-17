package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    var v3 int64 = 416
    v4 := make(map[string]uint32)
    var v5 int = 9223372036854775807
    if false {
        var v6 int32 = -2147483648
        v7 := GenS1{F0: (false || false), F1: 42, F2: 1}
        _ = v6
        _ = v7
    } else {
        v8 := []int{}
        for fi9 := 0; fi9 < 4; fi9++ {
            if (true || (59 <= 243)) {
                if true {
                    if len(v8) > 0 {
                        v8[0] = (fi9 & fi9)
                    }
                    fmt.Println(v1, fi9)
                }
                if false {
                    v1 = "a"
                } else {
                    v5 = fi9
                    var v10 uint32 = 4294967295
                    v5 -= -2147483648
                    _ = v10
                }
                v5 = 497
            }
            _ = fi9
        }
        if (false && true) {
            fmt.Println(v3, v1)
        } else {
            var v11 uint8 = 255
            _ = v11
        }
        _ = v8
    }
    if true {
        fmt.Println(v1, v5, v5)
        var v12 string = v1
        var v13 uint64 = 1
        _ = v12
        _ = v13
    }
    var v14 int32 = 747
    var v15 int = (v5 % (((v5 ^ (v5 / ((32767) | 1)))) | 1))
    v3 += (v3 >> ((0) & 63))
    switch (v5 << ((198) & 63)) {
    case 710:
        switch "a" {
        case "a":
            switch "a" {
            case "a":
                var v16 float32 = -2.25
                _ = v16
            case "":
                var v17 bool = (!(!true))
                _ = v17
            }
        case "hello":
            delete(v4, ("hello" + v1))
        case "\n\t":
            v14 += v14
        default:
            v15 = (v5 / (((v15 % ((v5) | 1))) | 1))
        }
    default:
        if ((((("" + ((v1 + v1) + "a")) + ((v1 + v1) + v1)) + ("" + ((("" + (v1 + v1)) + v1) + ""))) + v1) > v1) {
            if true {
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 int32 = (-(v14))
                    v20 := []int{(v5 >> ((759) & 63))}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v5 = 127
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 int8 = -1
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: (0 > 1), F1: 128, F2: 2}
                if true {
                    v24 := GenS2{F0: 32767, F1: 9223372036854775807}
                    var v25 bool = true
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint64 = 796
                        _ = fi26
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                } else {
                    var v28 string = v1
                    v29 := GenS2{F0: 1000, F1: (v3 | -1)}
                    _ = v28
                    _ = v29
                }
                fmt.Println(v14, v1, v5)
                _ = v23
            }
        }
    }
    switch 128 {
    case 32768:
        var v30 uint32 = 2147483648
        _ = v30
    case -1:
        var v31 uint8 = 128
        _ = v31
    case 750:
        var v32 int32 = -2147483648
        _ = v32
    }
    v14 = (v14 | (-(v14)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
