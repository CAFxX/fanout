package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint32 }
func genFn0() int64 {
    var v5 uint16 = 506
    if true {
        fmt.Println(v5)
        v6 := GenS1{F0: 0, F1: 1, F2: 641}
        var v7 int64 = 0
        _ = v6
        _ = v7
    } else {
        v5 = v5
        var v8 string = ("\n\t" + "foo bar")
        _ = v8
    }
    _ = v5
    return -252
}
func genFn1() int16 {
    var v9 float64 = 1e10
    var v10 int8 = 0
    _ = v10
    _ = v9
    return -921
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -1
    v3 := make(map[int]string)
    var v4 uint8 = 255
    v11 := GenS2{F0: (v4 % ((v4) | 1)), F1: "\n\t"}
    for fi12 := 0; fi12 < 7; fi12++ {
        v4 += v4
        v13 := GenS3{F0: fi12, F1: 0}
        v14 := []int{v2, (fi12 << ((42) & 63)), (fi12 | (fi12 << ((269) & 63))), v2}
        _ = fi12
        _ = v13
        _ = v14
    }
    if (genFn1() < genFn1()) {
        var v15 uint16 = uint16(v4)
        switch ("hello" + "a") {
        case "foo bar":
            v15 += (v15 % (((v15 << ((255) & 63))) | 1))
        }
        v4 += v4
        _ = v15
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 int64 = 404
            _ = len(v3)
            _ = fi16
            _ = v17
        }
        var v18 float64 = (((-2.25 - (((-2.25 / -2.25) / 1.5) * 1e10)) + 0.1) / 1.0)
        _ = v18
    }
    v4 += (v4 + v4)
    var v19 float64 = 3.14159
    v20 := GenS1{F0: 687, F1: 9223372036854775807, F2: v2}
    fmt.Println(v19, v2, v19)
    if (255 == 4294967295) {
        fmt.Println(v2, v4, v2)
        switch "a" {
        case "":
            if (1 == 876) {
                v21 := GenS2{F0: v4, F1: "\n\t"}
                _ = v21
            } else {
                v22 := []int{-364}
                switch "a" {
                case "":
                    fmt.Println(v19, v2)
                default:
                    switch "hello" {
                    case "\n\t":
                        v4 = (v4 / ((255) | 1))
                    default:
                        v22 = append(v22, (v2 << ((636) & 63)))
                    }
                }
                if true {
                    var v23 int = 664
                    for fi24 := 0; fi24 < 12; fi24++ {
                        if false {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi24)
                            case "a":
                                v19 += v19
                            case "hello":
                                if (18446744073709551615 == 0) {
                                    if true {
                                        switch 32768 {
                                        case 4294967295:
                                            if false {
                                                _ = len(v22)
                                                v23 += v2
                                            }
                                        default:
                                            var v25 int32 = 0
                                            _ = v25
                                        }
                                        switch 990 {
                                        case -769:
                                            var v26 uint = 18446744073709551615
                                            _ = v26
                                        case 271:
                                            v19 = (v19 + 0.1)
                                        case 1:
                                            var v27 uint64 = 200
                                            _ = v27
                                        default:
                                            v2 += fi24
                                        }
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    }
                                    switch ("hello" + "hello") {
                                    case "a":
                                        for fi28 := 0; fi28 < 17; fi28++ {
                                            v22 = append(v22, v23)
                                            _ = fi28
                                        }
                                    case "":
                                        if (genFn0() == -315) {
                                            fmt.Println(fi24)
                                            var v29 uint64 = 866
                                            _ = v29
                                        }
                                    default:
                                        switch -701 {
                                        case 32768:
                                            for fi30 := 0; fi30 < 1; fi30++ {
                                                v19 = (-2.25 + v19)
                                                var v31 uint16 = 150
                                                v32 := GenS4{F0: 1}
                                                _ = fi30
                                                _ = v31
                                                _ = v32
                                            }
                                        case -9223372036854775808:
                                            var v33 int = (-(v23))
                                            _ = v33
                                        default:
                                            delete(v3, (fi24 << ((255) & 63)))
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v4)
                            }
                            var v34 float64 = 1e10
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 2; fi35++ {
                            var v36 uint16 = 42
                            var v37 string = ("" + "foo bar")
                            fmt.Println(v23)
                            v38 := GenS1{F0: 42, F1: -9223372036854775808, F2: (v2 & (v23 & v2))}
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        _ = fi24
                    }
                    v4 = (^(v4))
                    _ = v23
                }
                _ = v22
            }
        case "\n\t":
            var v39 uint = 15
            _ = v39
        case "hello":
            v40 := v3[(v2 + (v2 - int(v4)))]
            _ = v40
        }
        var v41 uint32 = 230
        _ = v41
    } else {
        v42 := GenS4{F0: 0}
        _ = v42
    }
    var v43 uint8 = 128
    fmt.Println(v19, v19, v19)
    v43 = 2
    var v44 int8 = 100
    fmt.Println(v43)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
