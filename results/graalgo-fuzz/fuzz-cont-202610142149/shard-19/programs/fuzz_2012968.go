package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 uint8; F2 uint16 }

func main() {
    var v1 float64 = 0.0
    var v2 bool = (true || (958 <= 648))
    var v3 uint = 255
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    v3 -= 2
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 1, F1: 0, F2: 1}
        fmt.Println(v1)
        fmt.Println(v2)
        _ = fi6
        _ = v7
    }
    var v8 float64 = -2.25
    if v2 {
        var v9 float64 = v1
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint32 = 4294967295
            v12 := []int{fi10}
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 1000}
        _ = v13
        _ = v9
    } else {
        v14 := GenS3{F0: 0}
        v15 := GenS4{F0: 215, F1: 100}
        v2 = (!(!(!v2)))
        _ = v14
        _ = v15
    }
    switch 42 {
    case 4294967295:
        switch "\n\t" {
        case "hello":
            switch "hello" {
            case "\n\t":
                var v16 int32 = -1
                _ = v16
            case "a":
                switch "hello" {
                case "\n\t":
                    v8 += 0.0
                case "hello":
                    switch "" {
                    case "\n\t":
                        fmt.Println(v2)
                    default:
                        v3 += (v3 / ((42) | 1))
                    }
                }
            case "hello":
                var v17 int32 = -1
                _ = v17
            default:
                switch 2147483647 {
                case 4294967295:
                    var v18 int32 = 0
                    _ = v18
                case 32768:
                    var v19 uint16 = 1
                    _ = v19
                case -119:
                    v2 = v2
                default:
                    if (("a" + (("" + "a") + "a")) == ("foo bar" + ("" + (("foo bar" + "") + "hello")))) {
                        if (!(1 == 32767)) {
                            v20 := []int{}
                            var v21 int32 = 42
                            _ = v20
                            _ = v21
                        }
                        switch "" {
                        case "hello":
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v1 = 0.1
                                _ = fi22
                            }
                        }
                    }
                }
            }
        }
    default:
        v1 -= (v1 - (v1 - v1))
    }
    v23 := GenS5{F0: v2, F1: 127}
    switch "hello" {
    case "a":
        var v24 uint = (v3 ^ (^(v3)))
        _ = v24
    case "":
        if (127 != 255) {
            switch -279 {
            case 32767:
                if (!false) {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 int16 = 42
                        _ = fi25
                        _ = v26
                    }
                } else {
                    var v27 uint16 = 468
                    _ = v27
                }
            case 2147483647:
                v28 := GenS6{F0: -2147483648, F1: 355, F2: 640}
                _ = v28
            default:
                var v29 uint = 2
                _ = v29
            }
            v1 = (-2.25 * v8)
            if ((((-2.25 * 3.14159) - 0.0) != (1.5 + float32(3.14159))) || (v2 || v2)) {
                v5[32767] = 1
                for fi30 := 0; fi30 < 5; fi30++ {
                    delete(v5, fi30)
                    var v31 uint8 = 154
                    var v32 uint = 4294967295
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := GenS5{F0: ((-1 >= 140) && true), F1: 127}
                fmt.Println(v1, v8, v3)
                _ = v33
            }
        }
    default:
        switch 32767 {
        case -9223372036854775808:
            if false {
                v34 := GenS3{F0: -9223372036854775808}
                if v2 {
                    var v35 uint32 = 7
                    _ = v35
                } else {
                    switch 127 {
                    case 128:
                        v36 := GenS7{F0: -32768}
                        _ = v36
                    }
                    v8 = 0.0
                    for fi37 := 0; fi37 < 14; fi37++ {
                        var v38 int8 = 127
                        _ = fi37
                        _ = v38
                    }
                }
                _ = v34
            } else {
                v4[-1] = 32767
            }
        case 32768:
            v39 := []int{-9223372036854775808, 298, -646, -462}
            _ = v39
        default:
            v40 := []int{0, 2147483647, 0}
            _ = v40
        }
    }
    switch ("hello" + "a") {
    case "":
        v8 -= v8
    default:
        delete(v4, -808)
    }
    var v41 uint8 = 128
    fmt.Println(v3, v1, v8)
    v42 := GenS8{F0: 255, F1: (^(v41)), F2: 325}
    var v43 uint8 = (^(v41))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
