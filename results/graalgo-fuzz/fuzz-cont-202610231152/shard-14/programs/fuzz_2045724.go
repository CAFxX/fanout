package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8; F1 string; F2 int32 }
type GenS6 struct { F0 int }
func genFn0(p0 uint) int64 {
    var v6 float32 = 0.1
    var v7 int8 = 42
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-2147483648}
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    if true {
        var v8 string = "\n\t"
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                if len(v2) > 0 {
                    v2[0] = fi9
                }
                switch "a" {
                case "\n\t":
                    v11 := v5[4294967295]
                    _ = v11
                case "hello":
                    v12 := GenS1{F0: -9223372036854775808, F1: 255}
                    _ = v12
                default:
                    v13 := GenS1{F0: fi10, F1: 255}
                    _ = v13
                }
                _ = fi10
            }
            var v14 int16 = 32767
            _ = fi9
            _ = v14
        }
    }
    var v15 uint32 = 256
    v15 = 0
    fmt.Println(v15)
    switch ("" + "\n\t") {
    case "a":
        switch 42 {
        case 0:
            switch ("foo bar" + "hello") {
            case "a":
                v16 := GenS2{F0: 18446744073709551615}
                _ = v16
            default:
                if (-2147483648 != -2147483648) {
                    v15 += v15
                }
            }
        case -770:
            var v17 float32 = 1e10
            _ = v17
        }
    default:
        v15 += v15
    }
    v18 := GenS3{F0: 685}
    var v19 int16 = -32768
    for fi20 := 0; fi20 < 6; fi20++ {
        switch "hello" {
        case "":
            v15 -= (v15 % ((1) | 1))
        default:
            if true {
                v19 += int16(v15)
            }
        }
        if false {
            var v21 int64 = 9223372036854775807
            if true {
                fmt.Println(v21)
                var v22 uint32 = v15
                v21 -= v21
                _ = v22
            } else {
                v23 := GenS4{F0: 102, F1: 1, F2: 42}
                _ = v23
            }
            var v24 int32 = -332
            _ = v21
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 string = "hello"
                var v27 int64 = 9223372036854775807
                _ = fi25
                _ = v26
                _ = v27
            }
            v2 = append(v2, (fi20 >> ((1) & 63)))
            v15 = 1
        }
        _ = fi20
    }
    fmt.Println(v19, v15)
    v28 := GenS5{F0: 127, F1: "\n\t", F2: -1}
    var v29 uint16 = 1
    v30 := GenS6{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
