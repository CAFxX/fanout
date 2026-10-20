package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
func genFn0(p0 uint16, p1 int32) int64 {
    v6 := []int{}
    var v7 uint8 = 255
    if (!(((p1 >> ((255) & 63)) >= (p1 | 215)) || true)) {
        if (!((!true) && (5 >= (v7 & (v7 - 219))))) {
            _ = len(v6)
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int64) int16 {
    v8 := []int{}
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v9 int8 = 0
        fmt.Println(v9)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            p0 += p0
            fmt.Println(fi10, p0)
            _ = fi10
        }
        var v11 uint32 = 747
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 480
}

func main() {
    var v1 int16 = 1
    var v2 string = "\n\t"
    v3 := []uint32{731, 65535, 42}
    var v4 uint64 = 4294967295
    var v5 string = v2
    v12 := GenS1{F0: false}
    if false {
        v13 := GenS1{F0: (-673 >= 127)}
        v14 := GenS2{F0: 0, F1: genFn0(65535, 42), F2: -2147483648}
        _ = v13
        _ = v14
    } else {
        v3 = append(v3, 989)
        var v15 uint16 = 560
        v16 := GenS2{F0: 108, F1: -9223372036854775808, F2: 2147483647}
        _ = v15
        _ = v16
    }
    switch v2 {
    case "hello":
        v17 := GenS3{F0: -9223372036854775808, F1: 127, F2: v4}
        _ = v17
    case "foo bar":
        v2 = v5
    case "a":
        switch v2 {
        case "":
            v3 = append(v3, 980)
        case "\n\t":
            switch -1 {
            case 128:
                v1 += (v1 % (((v1 - (^(v1)))) | 1))
            case 543:
                var v18 string = v5
                _ = v18
            case -277:
                var v19 uint32 = 778
                _ = v19
            default:
                v5 = v2
            }
        case "hello":
            switch -9223372036854775808 {
            case -385:
                switch v5 {
                case "\n\t":
                    var v20 uint16 = 65535
                    _ = v20
                case "":
                    v21 := GenS3{F0: 428, F1: 128, F2: v4}
                    _ = v21
                case "a":
                    if true {
                        var v22 int32 = -1
                        fmt.Println(v4)
                        v23 := GenS2{F0: -1, F1: 1, F2: 936}
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int16 = -18
                        v24 += v1
                        var v25 int64 = -908
                        _ = v24
                        _ = v25
                    }
                }
            default:
                var v26 uint64 = v4
                _ = v26
            }
        }
    }
    if true {
        switch "hello" {
        case "\n\t":
            fmt.Println(v2)
        default:
            var v27 int32 = 2147483647
            _ = v27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 float64 = 0.0
            if false {
                fmt.Println(v29)
                var v30 string = v5
                if (v4 <= v4) {
                    _ = len(v3)
                } else {
                    fmt.Println(v29, v30, v29)
                    for fi31 := 0; fi31 < 8; fi31++ {
                        switch fi28 {
                        case 127:
                            switch 2147483647 {
                            case -9223372036854775808:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 32768:
                                if false {
                                    v5 = "a"
                                    var v32 string = (v2 + ((((((v30 + "foo bar") + v30) + (v30 + (v2 + v2))) + v5) + "foo bar") + ((v30 + "a") + v2)))
                                    _ = v32
                                } else {
                                    var v33 int32 = 2147483647
                                    _ = v33
                                }
                            case 1:
                                if (false || true) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        _ = len(v3)
                                    case "foo bar":
                                        v34 := GenS3{F0: 904, F1: (fi31 & 9223372036854775807), F2: 1000}
                                        _ = v34
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 977
                                    }
                                    var v35 uint = 477
                                    _ = v35
                                } else {
                                    fmt.Println(v5, fi28)
                                }
                            }
                        case -120:
                            v4 += (v4 % ((v4) | 1))
                            v29 -= v29
                            var v36 string = (v2 + v30)
                            _ = v36
                        }
                        var v37 uint32 = 0
                        _ = fi31
                        _ = v37
                    }
                }
                _ = v30
            }
            _ = fi28
            _ = v29
        }
    }
    v4 = 0
    var v38 uint8 = 0
    var v39 float32 = (float32(0.1) * (0.1 / 1e10))
    v2 = v5
    var v40 int16 = v1
    var v41 int16 = genFn1(-9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
