package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    var v3 uint = 4294967295
    var v4 string = "\n\t"
    var v5 string = "hello"
    var v6 string = "hello"
    v5 = v5
    var v7 uint32 = 2147483648
    if true {
        v8 := GenS1{F0: uint16(v7), F1: (v7 - (v7 / ((v7) | 1))), F2: 32767}
        v9 := []int{-2147483648, 708, 908, 4294967295}
        var v10 uint8 = 157
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int64 = 42
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(v7, v1, v4)
            v14 := GenS2{F0: -497, F1: 77, F2: 4294967295}
            var v15 int64 = (v11 << (((v3 << (((v3 / (((v3 / ((v3) | 1))) | 1))) & 63))) & 63))
            _ = fi13
            _ = v14
            _ = v15
        }
        v4 = ("" + (((v4 + v5) + "hello") + v6))
        var v16 int16 = 0
        _ = fi12
        _ = v16
    }
    v17 := []int{32768, 9223372036854775807, 32768, 42}
    for fi18 := 0; fi18 < 12; fi18++ {
        v17 = append(v17, -9223372036854775808)
        _ = fi18
    }
    if v1 {
        if ((((!(1 != -386)) && ((2147483647 <= 127) && false)) || (!(0 == 0))) || false) {
            fmt.Println(v5, v4, v3)
        } else {
            v19 := GenS2{F0: -32768, F1: 2, F2: 787}
            var v20 uint8 = 127
            _ = v19
            _ = v20
        }
        var v21 float64 = 1e10
        _ = v21
    } else {
        v22 := GenS1{F0: 250, F1: 1, F2: 32767}
        _ = v22
    }
    if v1 {
        if v1 {
            if (1 != 2) {
                var v23 string = v6
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 float32 = 0.0
                    v7 -= 682
                    if (33 == 366) {
                        for fi26 := 0; fi26 < 6; fi26++ {
                            switch -952 {
                            case 0:
                                if (18446744073709551615 != 18446744073709551615) {
                                    var v27 int16 = -422
                                    var v28 uint64 = 1
                                    _ = v27
                                    _ = v28
                                } else {
                                    fmt.Println(v25, fi26)
                                }
                            case 2147483647:
                                v29 := GenS1{F0: 42, F1: v7, F2: 32767}
                                _ = v29
                            default:
                                var v30 int64 = (v11 + int64(v7))
                                _ = v30
                            }
                            v31 := GenS1{F0: 0, F1: 4294967295, F2: -413}
                            _ = fi26
                            _ = v31
                        }
                        var v32 bool = (2 <= 42)
                        switch "a" {
                        case "\n\t":
                            var v33 uint32 = v7
                            _ = v33
                        case "":
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        default:
                            switch v23 {
                            case "hello":
                                var v34 uint = (v3 - (v3 + v3))
                                _ = v34
                            case "a":
                                v35 := GenS1{F0: 2, F1: 798, F2: 1}
                                _ = v35
                            }
                        }
                        _ = v32
                    }
                    _ = fi24
                    _ = v25
                }
                v23 = v6
                _ = v23
            }
        } else {
            var v36 uint = 42
            _ = v36
        }
        v4 = (v5 + "\n\t")
        _ = len(v2)
    }
    v37 := GenS1{F0: 2, F1: (v7 << ((v3) & 63)), F2: 0}
    if v1 {
        if v1 {
            if (v7 < (v7 << (((v3 & (^(v3)))) & 63))) {
                v38 := GenS2{F0: 32767, F1: 128, F2: 43}
                _ = v38
            } else {
                var v39 uint64 = 0
                _ = v39
            }
        } else {
            if v1 {
                switch "\n\t" {
                case "a":
                    if true {
                        v40 := GenS3{F0: v1, F1: -128, F2: 18446744073709551615}
                        _ = len(v2)
                        _ = v40
                    } else {
                        v5 = v6
                    }
                case "":
                    v41 := GenS1{F0: 65535, F1: v7, F2: 1}
                    _ = v41
                default:
                    switch 128 {
                    case 127:
                        for fi42 := 0; fi42 < 20; fi42++ {
                            var v43 int = int(v11)
                            v4 = v4
                            switch 4294967295 {
                            case 32768:
                                var v44 int8 = 0
                                _ = v44
                            case 4294967295:
                                v4 = "\n\t"
                            }
                            _ = fi42
                            _ = v43
                        }
                    case 0:
                        var v45 float32 = 1e10
                        _ = v45
                    default:
                        v2[1] = (42 == v3)
                    }
                }
                fmt.Println(v4)
            } else {
                var v46 float64 = 0.1
                _ = v46
            }
        }
        v6 = ""
        var v47 int16 = -32768
        _ = v47
    }
    v48 := GenS1{F0: 257, F1: (v7 + (v7 + v7)), F2: 42}
    v49 := GenS2{F0: 42, F1: 175, F2: 4294967295}
    for fi50 := 0; fi50 < 9; fi50++ {
        v51 := GenS1{F0: 255, F1: (^(v7)), F2: 0}
        if false {
            var v52 uint = (v3 | (v3 << ((v3) & 63)))
            v7 = v7
            v1 = false
            _ = v52
        } else {
            if v1 {
                delete(v2, (fi50 + -1))
            } else {
                for fi53 := 0; fi53 < 20; fi53++ {
                    var v54 uint64 = 42
                    _ = fi53
                    _ = v54
                }
            }
            switch -827 {
            case 1:
                v3 = 295
            }
        }
        v55 := GenS2{F0: 347, F1: 255, F2: 1}
        _ = fi50
        _ = v51
        _ = v55
    }
    fmt.Println(v11)
    if (!(!(79 >= 42))) {
        fmt.Println(v3, v5)
        fmt.Println(v6, v4, v6)
        var v56 int16 = 32767
        _ = v56
    }
    var v57 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
