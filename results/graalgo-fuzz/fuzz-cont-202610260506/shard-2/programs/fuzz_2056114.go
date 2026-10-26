package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int16; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint8; F1 uint }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 bool = (1 > 1)
    switch -9223372036854775808 {
    case -528:
        v5 := GenS1{F0: 2147483647, F1: 127}
        _ = v5
    case 127:
        var v6 float64 = (float64(float32(0.0)) + 3.14159)
        _ = v6
    }
    var v7 int = 381
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if ("\n\t" <= "a") {
            if true {
                v7 = 32768
            }
            switch v7 {
            case -832:
                v9 := GenS1{F0: 1, F1: (v7 * v7)}
                _ = v9
            case 0:
                if (32767 < 1000) {
                    delete(v3, "\n\t")
                    v10 := GenS1{F0: 722, F1: (-(fi8))}
                    var v11 int = fi8
                    _ = v10
                    _ = v11
                } else {
                    v7 = (fi8 - 1)
                    v12 := GenS1{F0: 2, F1: (^(fi8))}
                    if true {
                        v7 -= -1
                        if ((!false) || false) {
                            if (!(fi8 < v7)) {
                                if (!true) {
                                    v4 = true
                                    if v4 {
                                        if (v4 && (2147483647 > 2147483648)) {
                                            if v4 {
                                                var v13 uint16 = 421
                                                v13 = 32767
                                                fmt.Println(v4, v4)
                                                _ = v13
                                            } else {
                                                v1 = append(v1, 0)
                                                for fi14 := 0; fi14 < 4; fi14++ {
                                                    var v15 uint8 = 255
                                                    fmt.Println(v15)
                                                    _ = fi14
                                                    _ = v15
                                                }
                                                v4 = (!(776 == 884))
                                            }
                                        }
                                        var v16 int32 = -498
                                        _ = v16
                                    } else {
                                        if v4 {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v4 = (!false)
                                            _ = len(v3)
                                        }
                                        v17 := GenS2{F0: 65535, F1: 65535}
                                        _ = v17
                                    }
                                    if (!(2147483647 > 1)) {
                                        if (((float64((1.5 * 1e10)) * 1.5) / 1e10) <= -2.25) {
                                            var v18 int8 = 1
                                            var v19 uint64 = 4294967295
                                            for fi20 := 0; fi20 < 12; fi20++ {
                                                v21 := GenS3{F0: (v4 && (650 < 65535))}
                                                var v22 string = "a"
                                                if true {
                                                    fmt.Println(v18, v18, v19)
                                                    fmt.Println(v19, v7, v22)
                                                }
                                                _ = fi20
                                                _ = v21
                                                _ = v22
                                            }
                                            _ = v18
                                            _ = v19
                                        }
                                        delete(v2, 42)
                                    } else {
                                        v23 := GenS3{F0: (1 > 1)}
                                        var v24 float64 = 0.0
                                        _ = v23
                                        _ = v24
                                    }
                                } else {
                                    var v25 int32 = 2147483647
                                    v26 := GenS3{F0: false}
                                    _ = v25
                                    _ = v26
                                }
                            } else {
                                if true {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v4)
                                    case "a":
                                        v4 = v4
                                    default:
                                        if v4 {
                                            if false {
                                                if true {
                                                    var v27 int = (v7 / (((fi8 >> ((526) & 63))) | 1))
                                                    _ = v27
                                                }
                                            }
                                            if (42 == 2147483648) {
                                                for fi28 := 0; fi28 < 4; fi28++ {
                                                    v7 -= -393
                                                    _ = fi28
                                                }
                                                v4 = true
                                                var v29 int16 = 0
                                                _ = v29
                                            }
                                        }
                                    }
                                    for fi30 := 0; fi30 < 16; fi30++ {
                                        v31 := v2[42]
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v32 int8 = 42
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 uint32 = 4294967295
                                    _ = v33
                                }
                            }
                        } else {
                            var v34 float64 = 0.1
                            var v35 int = (fi8 % (((fi8 & 2147483647)) | 1))
                            v35 -= v35
                            _ = v34
                            _ = v35
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v12
                }
            default:
                v36 := GenS1{F0: 4294967295, F1: (v7 >> ((18446744073709551615) & 63))}
                _ = v36
            }
            v37 := GenS4{F0: 0, F1: 533}
            _ = v37
        }
        _ = fi8
    }
    if v4 {
        var v38 uint32 = 4294967295
        _ = v38
    } else {
        v4 = v4
        fmt.Println(v4)
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 uint16 = 49
        _ = fi39
        _ = v40
    }
    v41 := GenS5{F0: 42, F1: -403}
    for fi42 := 0; fi42 < 3; fi42++ {
        v7 = 1
        v43 := GenS1{F0: 65, F1: fi42}
        _ = fi42
        _ = v43
    }
    if v4 {
        fmt.Println(v4)
        v4 = (!v4)
        if true {
            var v44 int64 = 1
            v45 := GenS6{F0: -619, F1: 18446744073709551615, F2: -945}
            var v46 uint32 = 923
            _ = v44
            _ = v45
            _ = v46
        } else {
            var v47 bool = ((!v4) || (105 >= 66))
            v47 = (v47 && false)
            var v48 float32 = 0.1
            _ = v47
            _ = v48
        }
    } else {
        for fi49 := 0; fi49 < 14; fi49++ {
            v4 = true
            v7 -= (v7 - fi49)
            _ = fi49
        }
        v4 = true
        v4 = (!(!((!((!(!(-9223372036854775808 >= 9223372036854775807))) || (0 > 0))) || (813 != 4294967295))))
    }
    v4 = v4
    v7 += -1
    for fi50 := 0; fi50 < 11; fi50++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v51 := GenS7{F0: 0, F1: 61}
        _ = fi50
        _ = v51
    }
    v4 = (-9223372036854775808 == 0)
    v52 := GenS2{F0: 4294967295, F1: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v52
    _ = v7
}
