package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int; F1 uint }

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = 0.0
    var v3 int = 773
    if (!true) {
        var v4 string = "a"
        v3 = 128
        _ = v4
    } else {
        v5 := GenS1{F0: 18446744073709551615, F1: 0, F2: -32768}
        if (false || false) {
            for fi6 := 0; fi6 < 18; fi6++ {
                if ((true || (!true)) && (!(363 != 917))) {
                    var v7 int8 = 1
                    _ = v7
                } else {
                    var v8 int8 = 89
                    v3 -= (v3 | 32767)
                    fmt.Println(v8, v8, fi6)
                    _ = v8
                }
                for fi9 := 0; fi9 < 12; fi9++ {
                    v1 -= ((3.14159 / ((v2 + v2) - float64(((1.5 * 0.0) * 0.1)))) * float64((3.14159 / (3.14159 / -2.25))))
                    _ = fi9
                }
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 uint8 = 52
                    _ = fi10
                    _ = v11
                }
                _ = fi6
            }
        } else {
            fmt.Println(v1)
        }
        _ = v5
    }
    var v12 uint32 = 2
    if (v3 <= -689) {
        var v13 uint16 = 42
        _ = v13
    }
    v14 := GenS2{F0: true}
    switch "hello" {
    case "foo bar":
        for fi15 := 0; fi15 < 15; fi15++ {
            v3 -= (v3 % ((v3) | 1))
            v1 = ((1.5 + (((0.0 - ((v1 / v2) * 0.1)) * (v1 / (v2 / (v1 + v1)))) * 0.1)) + (v2 / float64((1.5 - -2.25))))
            var v16 uint64 = 290
            _ = fi15
            _ = v16
        }
    case "":
        var v17 int = (v3 << ((uint(v12)) & 63))
        _ = v17
    case "\n\t":
        v18 := []int{(v3 << ((2) & 63))}
        _ = v18
    default:
        if false {
            v1 -= v2
            var v19 string = ""
            _ = v19
        } else {
            v20 := GenS3{F0: -883, F1: 1}
            v21 := GenS4{F0: 898, F1: 2}
            _ = v20
            _ = v21
        }
    }
    v12 = v12
    v12 += v12
    if (!((true && true) && true)) {
        if (0.0 != 1.5) {
            for fi22 := 0; fi22 < 11; fi22++ {
                if false {
                    v12 -= (v12 / ((4294967295) | 1))
                }
                _ = fi22
            }
        } else {
            v3 = v3
            v12 += v12
            var v23 bool = (!(1 == -9223372036854775808))
            _ = v23
        }
    } else {
        switch (("hello" + "foo bar") + (("a" + "a") + "a")) {
        case "\n\t":
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                if true {
                    var v24 float32 = (0.1 * (float32(1.5) * -2.25))
                    switch "foo bar" {
                    case "a":
                        v2 -= v2
                    case "foo bar":
                        fmt.Println(v3, v12, v24)
                    }
                    _ = v24
                } else {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v1)
                    case "hello":
                        if (true && (!false)) {
                            v25 := GenS4{F0: (v12 / (((v12 | v12)) | 1)), F1: 2}
                            fmt.Println(v1, v3)
                            _ = v25
                        } else {
                            switch v3 {
                            case -9223372036854775808:
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            case 128:
                                if ((-2.25 / (1.5 - 0.0)) > 1e10) {
                                    if false {
                                        var v27 int64 = -1
                                        var v28 uint = 18446744073709551615
                                        _ = v27
                                        _ = v28
                                    }
                                }
                            case -992:
                                v29 := []int{}
                                _ = v29
                            default:
                                var v30 float64 = v2
                                _ = v30
                            }
                            var v31 uint32 = v12
                            _ = v31
                        }
                    }
                    if true {
                        fmt.Println(v3, v12)
                        v32 := GenS5{F0: (-(v3)), F1: 0}
                        _ = v32
                    }
                }
            case "hello":
                fmt.Println(v2, v1)
            default:
                v3 -= 2147483647
            }
        }
        v33 := []int{-2147483648, (v3 | (-(v3)))}
        _ = v33
    }
    v3 = (-(v3))
    switch v3 {
    case 32768:
        if false {
            var v34 int8 = -118
            _ = v34
        }
    case 127:
        var v35 float64 = 0.1
        _ = v35
    default:
        v36 := []int{v3}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
