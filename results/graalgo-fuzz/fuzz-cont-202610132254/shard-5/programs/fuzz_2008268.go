package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint; F1 int16; F2 int16 }
type GenS7 struct { F0 uint32; F1 int16; F2 int32 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 uint16 = 591
    v3 = 65535
    switch "hello" {
    case "a":
        if false {
            if false {
                v4 := v2[4294967295]
                switch ("" + ((("foo bar" + "\n\t") + "a") + ("\n\t" + "hello"))) {
                case "":
                    v5 := GenS1{F0: 1, F1: 42, F2: 389}
                    _ = v5
                case "\n\t":
                    for fi6 := 0; fi6 < 7; fi6++ {
                        v4 = true
                        v3 = (v3 - (^(v3)))
                        _ = fi6
                    }
                default:
                    fmt.Println(v3, v4)
                }
                _ = v4
            }
            if false {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3)
                }
            } else {
                var v7 int8 = 0
                v7 -= (-(v7))
                _ = v7
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                switch fi8 {
                case 2147483647:
                    var v9 int8 = -128
                    _ = v9
                case 4294967295:
                    v3 = v3
                }
                fmt.Println(v3, v3)
                delete(v1, "")
                _ = fi8
            }
        } else {
            v3 += 1
            v3 = v3
            v10 := GenS1{F0: 4294967295, F1: 1000, F2: v3}
            _ = v10
        }
    }
    if false {
        switch 264 {
        case 9223372036854775807:
            switch "a" {
            case "a":
                fmt.Println(v3)
            case "":
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(fi11)
                    _ = fi11
                }
            case "foo bar":
                if ((65535 >= 31) || (33 > (^(v3)))) {
                    var v12 uint8 = 2
                    _ = v12
                }
            default:
                v13 := GenS2{F0: -1, F1: 1}
                _ = v13
            }
        case 4294967295:
            v3 = (v3 - v3)
        case -364:
            var v14 float64 = (-2.25 * (0.1 * 1e10))
            _ = v14
        default:
            var v15 float32 = 3.14159
            _ = v15
        }
    }
    v3 -= 0
    v16 := GenS3{F0: 82}
    switch "" {
    case "foo bar":
        v3 -= 32767
    case "":
        fmt.Println(v3)
    default:
        if false {
            v3 -= (v3 & 255)
            v3 = 255
        } else {
            switch 32768 {
            case 42:
                switch 1 {
                case 42:
                    v17 := GenS4{F0: 403, F1: 2147483647}
                    _ = v17
                case 1:
                    var v18 string = ("a" + ("" + (((("" + ("a" + ("\n\t" + "hello"))) + "foo bar") + ("foo bar" + "\n\t")) + "hello")))
                    _ = v18
                case 128:
                    var v19 uint8 = 1
                    _ = v19
                default:
                    v20 := GenS5{F0: "hello", F1: -1}
                    _ = v20
                }
            case -515:
                var v21 uint16 = v3
                _ = v21
            default:
                v3 = (v3 - 1)
            }
            v22 := GenS2{F0: 0, F1: -107}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS6{F0: 0, F1: 834, F2: -1}
        switch ("hello" + "") {
        case "hello":
            var v25 uint64 = 0
            _ = v25
        case "\n\t":
            var v26 float64 = 3.14159
            _ = v26
        default:
            v27 := GenS2{F0: -2147483648, F1: 127}
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    switch "foo bar" {
    case "hello":
        v3 = (v3 & 1)
    case "\n\t":
        v3 += 794
    case "foo bar":
        var v28 int32 = 2147483647
        _ = v28
    }
    v3 -= v3
    v29 := GenS7{F0: 704, F1: 389, F2: -718}
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 uint64 = 290
        _ = fi30
        _ = v31
    }
    fmt.Println(v3)
    v3 -= v3
    fmt.Println(v3)
    for fi32 := 0; fi32 < 2; fi32++ {
        v3 = (^(v3))
        v3 = v3
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 int32 = -2147483648
            v35 := GenS8{F0: 4294967295, F1: (v3 % (((^(v3))) | 1)), F2: 32}
            if true {
                switch "a" {
                case "":
                    v36 := []int{}
                    _ = v36
                case "foo bar":
                    var v37 uint16 = (v3 | 997)
                    _ = v37
                case "hello":
                    v38 := GenS2{F0: (v34 ^ v34), F1: 127}
                    _ = v38
                default:
                    var v39 int32 = (-(v34))
                    _ = v39
                }
                v40 := GenS3{F0: 35}
                _ = v40
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
}
