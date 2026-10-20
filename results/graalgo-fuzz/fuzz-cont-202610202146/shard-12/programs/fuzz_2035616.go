package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 int8 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0() int16 {
    if (42 <= 65535) {
        v7 := []int{-725}
        var v8 int16 = 1
        if false {
            switch "a" {
            case "foo bar":
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 float64 = 0.1
                    v10 -= -2.25
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v10 += (((v10 + 0.0) * 0.0) - v10)
                        var v12 int = fi11
                        _ = fi11
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
            case "a":
                fmt.Println(v8)
            default:
                var v13 float32 = 1.5
                _ = v13
            }
            var v14 int = -701
            switch "hello" {
            case "hello":
                switch 42 {
                case 32768:
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v8 = -869
                        v8 = (v8 & (-(v8)))
                        switch (fi15 * (v14 >> ((329) & 63))) {
                        case 416:
                            fmt.Println(v8, fi15)
                        case 32767:
                            for fi16 := 0; fi16 < 3; fi16++ {
                                v17 := GenS1{F0: 470}
                                switch ("a" + "a") {
                                case "a":
                                    if false {
                                        v14 += (fi15 << ((255) & 63))
                                        var v18 string = "a"
                                        _ = v18
                                    }
                                case "":
                                    v19 := GenS1{F0: 2}
                                    _ = v19
                                default:
                                    var v20 float64 = ((0.1 / -2.25) + float64(((1e10 - 1e10) / 3.14159)))
                                    _ = v20
                                }
                                var v21 uint8 = 0
                                _ = fi16
                                _ = v17
                                _ = v21
                            }
                        default:
                            var v22 int32 = -1
                            _ = v22
                        }
                        _ = fi15
                    }
                case 127:
                    v8 += (v8 >> ((65535) & 63))
                default:
                    var v23 int32 = -1
                    _ = v23
                }
            case "a":
                switch "foo bar" {
                case "hello":
                    v14 += (v14 >> ((69) & 63))
                case "a":
                    v24 := GenS2{F0: 0, F1: -2147483648}
                    _ = v24
                case "\n\t":
                    v8 = (v8 >> ((2) & 63))
                default:
                    v8 = (v8 * v8)
                }
            case "":
                if len(v7) > 0 {
                    v7[0] = 1
                }
            }
            _ = v14
        }
        _ = v7
        _ = v8
    }
    v25 := []int{-635, 32767}
    _ = v25
    return -808
}
func genFn1() bool {
    var v26 bool = true
    if false {
        fmt.Println(v26)
    } else {
        v26 = (!((1.5 < 1.5) && v26))
        v26 = (!(!v26))
    }
    if v26 {
        v27 := []int{}
        var v28 int8 = -75
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: 4294967295, F1: -839, F2: 42}
    _ = v26
    _ = v29
    return v26
}

func main() {
    var v1 uint64 = 42
    var v2 int = -2147483648
    var v3 int8 = -1
    var v4 int32 = 1
    var v5 uint = 353
    var v6 bool = false
    v30 := GenS2{F0: 46, F1: (^(v4))}
    for fi31 := 0; fi31 < 12; fi31++ {
        v4 += (v4 | (v4 & (v4 | v4)))
        v3 = 1
        var v32 int64 = -1
        _ = fi31
        _ = v32
    }
    var v33 uint16 = 664
    var v34 int16 = genFn0()
    v1 = v1
    v35 := GenS4{F0: (v4 + (v4 | v4)), F1: (v3 - v3)}
    v36 := []int{int(v5), (v2 % (((v2 | v2)) | 1)), v2}
    v33 -= (^(v33))
    switch "\n\t" {
    case "hello":
        var v37 int8 = v3
        _ = v37
    default:
        var v38 bool = v6
        _ = v38
    }
    v33 = v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
