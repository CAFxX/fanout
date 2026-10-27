package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 int8; F1 string }
func genFn0() uint16 {
    var v5 int32 = -2147483648
    switch ("\n\t" + (("hello" + "\n\t") + "")) {
    case "hello":
        if false {
            var v6 int64 = 7
            if (0.1 > ((-2.25 / 1.5) / -2.25)) {
                for fi7 := 0; fi7 < 4; fi7++ {
                    var v8 uint = 65535
                    _ = fi7
                    _ = v8
                }
            } else {
                switch -1 {
                case 2147483647:
                    switch (("" + "foo bar") + "\n\t") {
                    case "\n\t":
                        switch 127 {
                        case 9223372036854775807:
                            var v9 int64 = v6
                            _ = v9
                        case -1:
                            v10 := GenS1{F0: -604, F1: 42, F2: -66}
                            _ = v10
                        default:
                            if true {
                                fmt.Println(v6)
                            }
                        }
                    default:
                        switch ("a" + "hello") {
                        case "foo bar":
                            switch "a" {
                            case "\n\t":
                                v11 := []int{0, 42, -104}
                                _ = v11
                            case "foo bar":
                                if (!(!true)) {
                                    var v12 int16 = -589
                                    _ = v12
                                }
                            default:
                                var v13 int = -22
                                _ = v13
                            }
                        case "\n\t":
                            v5 += 2147483647
                        }
                    }
                case 1:
                    if true {
                        fmt.Println(v6)
                    }
                case 771:
                    v14 := GenS2{F0: 963, F1: false, F2: int16(v6)}
                    _ = v14
                default:
                    v15 := GenS1{F0: 128, F1: 645, F2: 47}
                    _ = v15
                }
            }
            if (0 > 255) {
                var v16 uint = 255
                var v17 float64 = -2.25
                v16 -= 799
                _ = v16
                _ = v17
            }
            _ = v6
        } else {
            var v18 uint64 = 18446744073709551615
            v18 = v18
            if ((v5 ^ v5) == v5) {
                v5 = 2147483647
                var v19 int8 = -26
                v20 := GenS2{F0: 0, F1: true, F2: -699}
                _ = v19
                _ = v20
            }
            _ = v18
        }
    }
    v21 := []int{9223372036854775807, 840, 42, 42}
    v5 += 0
    _ = v21
    _ = v5
    return 32767
}

func main() {
    var v1 bool = false
    var v2 uint8 = 191
    v3 := make(map[string]bool)
    var v4 uint16 = 792
    v2 += (v2 & 1)
    v2 -= v2
    _ = len(v3)
    v22 := []int{-1}
    var v23 float64 = (1e10 - (1.5 - 0.0))
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(v1, v23, v23)
        switch "hello" {
        case "foo bar":
            v25 := GenS2{F0: 2147483648, F1: v1, F2: 0}
            _ = v25
        case "hello":
            if len(v22) > 0 {
                v22[0] = (fi24 << ((540) & 63))
            }
        case "a":
            switch "" {
            case "a":
                for fi26 := 0; fi26 < 9; fi26++ {
                    if true {
                        var v27 uint32 = 1
                        var v28 int = (fi24 - fi26)
                        if true {
                            v29 := v3["a"]
                            v29 = (!((genFn0() == genFn0()) && (v29 || (true && v29))))
                            v22 = append(v22, -2147483648)
                            _ = v29
                        }
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
            }
        }
        v30 := GenS3{F0: ("\n\t" + "hello"), F1: uint(v4)}
        _ = fi24
        _ = v30
    }
    v31 := GenS1{F0: -884, F1: (v4 << ((2) & 63)), F2: 45}
    fmt.Println(v4, v23)
    v32 := GenS4{F0: (v2 % ((v2) | 1)), F1: 73, F2: (^(v2))}
    if false {
        v33 := GenS3{F0: "", F1: 255}
        v34 := v3[("foo bar" + "")]
        v35 := GenS2{F0: 824, F1: true, F2: 32767}
        _ = v33
        _ = v34
        _ = v35
    }
    switch ("a" + "a") {
    case "hello":
        var v36 int8 = 1
        _ = v36
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    switch "hello" {
    case "":
        v37 := GenS5{F0: 98, F1: "\n\t"}
        _ = v37
    case "foo bar":
        v4 += uint16(v2)
    case "hello":
        switch "hello" {
        case "foo bar":
            var v38 uint64 = 1
            _ = v38
        case "a":
            v4 = v4
        case "":
            var v39 uint64 = 4294967295
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
