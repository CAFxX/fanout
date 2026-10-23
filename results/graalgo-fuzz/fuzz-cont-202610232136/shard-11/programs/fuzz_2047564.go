package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
func genFn0() uint32 {
    v5 := GenS1{F0: -951}
    var v6 int8 = 42
    var v7 uint8 = 167
    var v8 int = 42
    var v9 int = -193
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float64 = 0.1
    var v2 string = ""
    var v3 int32 = -32
    v4 := make(map[uint32]bool)
    fmt.Println(v2)
    v1 = 0.0
    var v10 uint16 = 803
    v11 := GenS1{F0: v3}
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint16 = (v10 / (((^(v10))) | 1))
        _ = fi12
        _ = v13
    }
    if (0.1 < 1e10) {
        switch 42 {
        case -1:
            v4[42] = (true && true)
        }
    } else {
        switch v2 {
        case "hello":
            v14 := GenS1{F0: (-(v3))}
            _ = v14
        case "foo bar":
            var v15 float32 = 0.1
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 12; fi16++ {
                if ((0.0 * v1) != 0.0) {
                    var v17 int64 = 208
                    _ = v17
                }
                _ = fi16
            }
        }
    }
    var v18 int16 = 738
    v19 := GenS2{F0: v10, F1: ((127 == 234) && true), F2: (!(50 > 2))}
    var v20 uint = 375
    v10 -= v10
    v10 = 0
    v4[2147483648] = true
    switch "" {
    case "":
        if false {
            switch v2 {
            case "\n\t":
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint64 = 0
                    v10 += (v10 + (v10 * 1))
                    v4[2147483647] = false
                    _ = fi21
                    _ = v22
                }
            case "":
                switch ("foo bar" + "") {
                case "a":
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v1 = 1e10
                        _ = fi23
                    }
                default:
                    switch v2 {
                    case "":
                        var v24 float32 = ((0.0 * 0.1) + 0.0)
                        _ = v24
                    case "hello":
                        fmt.Println(v2, v18, v2)
                    case "\n\t":
                        var v25 int16 = 32767
                        _ = v25
                    }
                }
            case "foo bar":
                v26 := GenS1{F0: v3}
                _ = v26
            }
            var v27 uint32 = 812
            v28 := GenS3{F0: (^(v10)), F1: v27}
            _ = v27
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            switch "" {
            case "foo bar":
                fmt.Println(fi29, v18, v3)
            case "\n\t":
                switch fi29 {
                case 1:
                    var v30 string = ("foo bar" + ("a" + v2))
                    _ = v30
                case 203:
                    var v31 int64 = -468
                    _ = v31
                }
            }
            v32 := []int{(fi29 - (fi29 << ((v20) & 63))), -1, -2147483648}
            _ = fi29
            _ = v32
        }
    }
    v33 := GenS4{F0: -9223372036854775808, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
}
