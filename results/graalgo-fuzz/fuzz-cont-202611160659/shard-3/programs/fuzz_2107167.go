package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }
type GenS6 struct { F0 int8; F1 int32 }

func main() {
    var v1 int64 = -1
    v2 := []uint32{671}
    var v3 float64 = -2.25
    v3 = float64((0.1 * -2.25))
    var v4 uint16 = 133
    if false {
        var v5 uint32 = 908
        v1 = (v1 >> ((uint(v4)) & 63))
        for fi6 := 0; fi6 < 15; fi6++ {
            v3 = float64(float32((((v3 / 1e10) * v3) - 1.5)))
            _ = fi6
        }
        _ = v5
    }
    var v7 uint8 = 128
    switch ("a" + "foo bar") {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 2
        }
    case "\n\t":
        v7 -= (^(v7))
    case "a":
        v3 = ((0.1 - v3) + float64((1.5 + (3.14159 + 1.5))))
    default:
        var v8 int = 621
        _ = v8
    }
    v3 = v3
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v7, v3)
        _ = fi9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi11 {
            case 127:
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            default:
                var v12 int16 = 0
                _ = v12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            switch fi10 {
            case -2147483648:
                fmt.Println(v1, fi13)
            case 604:
                if (fi10 != fi13) {
                    v14 := GenS1{F0: v7, F1: (v7 / ((v7) | 1)), F2: v7}
                    if ((^(v7)) <= 1) {
                        if false {
                            if true {
                                _ = len(v2)
                            }
                            v15 := GenS2{F0: 100, F1: 0}
                            _ = v15
                        }
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                        v3 = 0.0
                    } else {
                        v1 = v1
                    }
                    _ = v14
                } else {
                    v16 := GenS3{F0: (!(!false))}
                    if false {
                        v17 := GenS4{F0: 0, F1: -993}
                        _ = v17
                    } else {
                        if ((1e10 + 1e10) == 1e10) {
                            fmt.Println(v3, fi10, v7)
                            v18 := GenS1{F0: 127, F1: 2, F2: 255}
                            _ = v18
                        }
                        if false {
                            if (float64((0.1 - float32(v3))) >= v3) {
                                switch fi13 {
                                case 160:
                                    v19 := GenS1{F0: v7, F1: v7, F2: v7}
                                    _ = v19
                                case -9223372036854775808:
                                    v7 += (v7 / ((42) | 1))
                                default:
                                    v3 -= (v3 / 3.14159)
                                }
                                var v20 int32 = 42
                                switch "" {
                                case "hello":
                                    v1 -= (^(v1))
                                case "":
                                    v21 := GenS1{F0: 9, F1: v7, F2: 0}
                                    _ = v21
                                case "a":
                                    var v22 string = "hello"
                                    _ = v22
                                }
                                _ = v20
                            }
                        } else {
                            fmt.Println(fi13)
                        }
                    }
                    v23 := GenS3{F0: true}
                    _ = v16
                    _ = v23
                }
            case 1:
                v24 := GenS3{F0: true}
                _ = v24
            }
            var v25 uint16 = v4
            _ = fi13
            _ = v25
        }
        _ = fi10
    }
    var v26 uint64 = 18446744073709551615
    switch 32768 {
    case 9223372036854775807:
        v7 -= 0
    case -571:
        for fi27 := 0; fi27 < 8; fi27++ {
            if true {
                var v28 uint = 18446744073709551615
                fmt.Println(v26, v4, v7)
                var v29 string = "\n\t"
                _ = v28
                _ = v29
            }
            v30 := GenS4{F0: 32767, F1: 2147483647}
            _ = fi27
            _ = v30
        }
    case -905:
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 uint16 = (v4 | uint16(v26))
                fmt.Println(v3)
                if true {
                    v7 -= (v7 % ((v7) | 1))
                    var v34 int8 = 104
                    var v35 int16 = -1
                    var v36 float32 = 0.0
                    var v37 bool = false
                    switch fi32 {
                    case 991:
                        var v38 int = (fi32 | 2147483647)
                        _ = v38
                    case -9223372036854775808:
                        v26 = v26
                    }
                    _ = v34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v39 := GenS3{F0: false}
                _ = fi32
                _ = v33
                _ = v39
            }
            _ = fi31
        }
    default:
        for fi40 := 0; fi40 < 13; fi40++ {
            v41 := GenS1{F0: 255, F1: 1, F2: 72}
            v4 = v4
            fmt.Println(v1)
            _ = fi40
            _ = v41
        }
    }
    v7 = (v7 + 127)
    v3 = (v3 - v3)
    var v42 int64 = v1
    fmt.Println(v42, v26, v3)
    for fi43 := 0; fi43 < 6; fi43++ {
        if ("hello" <= "\n\t") {
            switch "foo bar" {
            case "foo bar":
                if ((!(v4 >= (v4 + 65535))) && false) {
                    v26 += v26
                    v26 = (v26 % (((v26 % (((^(v26))) | 1))) | 1))
                } else {
                    var v44 bool = (100 > 100)
                    var v45 float64 = float64(-2.25)
                    _ = v44
                    _ = v45
                }
            }
            v46 := GenS5{F0: (fi43 << ((255) & 63)), F1: false, F2: 386}
            _ = v46
        } else {
            for fi47 := 0; fi47 < 5; fi47++ {
                if true {
                    if true {
                        v26 = 18446744073709551615
                        var v48 int64 = v42
                        _ = v48
                    }
                    v3 -= 0.0
                    if (42 >= 0) {
                        fmt.Println(fi43, v42)
                        var v49 int64 = v1
                        _ = v49
                    }
                } else {
                    var v50 uint = 65535
                    var v51 float64 = v3
                    _ = v50
                    _ = v51
                }
                v4 = v4
                if true {
                    v52 := GenS1{F0: v7, F1: v7, F2: v7}
                    _ = v52
                }
                if (18446744073709551615 != 42) {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if true {
                        var v53 int8 = 100
                        v42 = (-(v42))
                        v42 = v42
                        _ = v53
                    }
                    if false {
                        if (!true) {
                            var v54 int64 = -9223372036854775808
                            _ = v54
                        } else {
                            v42 = v1
                        }
                        var v55 uint = uint(v7)
                        var v56 uint64 = (^(v26))
                        fmt.Println(v4)
                        fmt.Println(v4)
                        v57 := GenS6{F0: -80, F1: int32(v55)}
                        fmt.Println(v3)
                        var v58 float64 = float64(-2.25)
                        _ = v55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                }
                _ = len(v2)
                _ = fi47
            }
            v2 = append(v2, 888)
        }
        var v59 int32 = 0
        fmt.Println(v1, fi43)
        _ = fi43
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v7
}
