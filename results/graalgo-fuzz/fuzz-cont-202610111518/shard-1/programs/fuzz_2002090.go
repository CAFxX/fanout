package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 2
    var v3 float32 = (0.0 + (1e10 + 0.0))
    var v4 string = v1
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(v1, v1, v4)
        v6 := GenS1{F0: 42, F1: 4294967295}
        if false {
            v3 = 1.5
            if false {
                v2 = v2
                for fi7 := 0; fi7 < 17; fi7++ {
                    fmt.Println(fi7, fi5)
                    v4 = v4
                    _ = fi7
                }
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint = 255
                    v10 := []int{fi5, (fi8 | (fi5 % (((fi8 | (fi5 % ((127) | 1)))) | 1))), (fi8 + (fi5 % ((9223372036854775807) | 1)))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            v11 := GenS2{F0: 32767, F1: 693, F2: 255}
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    v12 := []int{-2147483648, 2147483647, 32768}
    switch v1 {
    case "hello":
        var v13 uint64 = (v2 ^ v2)
        _ = v13
    case "foo bar":
        v14 := GenS3{F0: 0, F1: 922, F2: 42}
        _ = v14
    case "":
        v1 = v1
    default:
        v3 -= 3.14159
    }
    var v15 int = -645
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS3{F0: 746, F1: 0, F2: 71}
        v18 := GenS2{F0: 42, F1: 65535, F2: 65535}
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        v19 := GenS1{F0: 127, F1: 18446744073709551615}
        v4 = "hello"
        _ = v19
    } else {
        for fi20 := 0; fi20 < 14; fi20++ {
            fmt.Println(v3, v4, fi20)
            fmt.Println(v3)
            _ = fi20
        }
        if (((!false) || false) && false) {
            var v21 float32 = 1e10
            _ = v21
        }
        if (42 < v2) {
            v22 := GenS1{F0: 42, F1: 1}
            _ = v22
        } else {
            if (!((!(!((128 <= 127) && true))) && false)) {
                if (0 > 42) {
                    if (152 == 114) {
                        v1 = (v1 + "a")
                    }
                    v15 += (-(v15))
                }
                v3 = 0.0
            } else {
                for fi23 := 0; fi23 < 12; fi23++ {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        fmt.Println(v1)
                        if (1 >= 100) {
                            if false {
                                fmt.Println(fi24)
                                var v25 string = "\n\t"
                                _ = v25
                            } else {
                                v15 = -808
                            }
                            v26 := GenS1{F0: -2147483648, F1: 35}
                            _ = v26
                        }
                        _ = len(v12)
                        _ = fi24
                    }
                    switch ("\n\t" + "a") {
                    case "":
                        for fi27 := 0; fi27 < 15; fi27++ {
                            var v28 uint16 = 2
                            var v29 float64 = (1e10 * (0.1 - -2.25))
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        fmt.Println(v2, v4, v15)
                    }
                    v4 = "\n\t"
                    _ = fi23
                }
                fmt.Println(v4, v3)
            }
            var v30 uint32 = 42
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := GenS1{F0: 2147483647, F1: 421}
        _ = fi31
        _ = v32
    }
    _ = len(v12)
    switch v4 {
    case "\n\t":
        switch v1 {
        case "":
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v2)
                _ = fi33
            }
        }
    case "":
        for fi34 := 0; fi34 < 7; fi34++ {
            v4 = (v4 + v4)
            var v35 uint32 = 0
            _ = fi34
            _ = v35
        }
    case "a":
        v36 := GenS2{F0: 1, F1: 65535, F2: 1}
        _ = v36
    }
    if false {
        v37 := GenS4{F0: 2147483647, F1: (2147483648 < 984)}
        if len(v12) > 0 {
            v12[0] = (v15 / (((v15 >> ((0) & 63))) | 1))
        }
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 int32 = -2147483648
        _ = fi38
        _ = v39
    }
    fmt.Println(v15, v15, v2)
    if false {
        var v40 int64 = 738
        _ = v40
    } else {
        v15 = (v15 % ((v15) | 1))
        var v41 float64 = 0.0
        v2 = v2
        _ = v41
    }
    switch 478 {
    case 504:
        for fi42 := 0; fi42 < 4; fi42++ {
            var v43 int = fi42
            _ = fi42
            _ = v43
        }
    case 0:
        fmt.Println(v15)
    case -198:
        var v44 uint = 42
        _ = v44
    default:
        _ = len(v12)
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        var v46 int8 = 0
        if (false || false) {
            v2 = 849
            v47 := GenS1{F0: -2147483648, F1: 2}
            _ = v47
        } else {
            if true {
                v15 += (-(v15))
                var v48 string = v1
                _ = v48
            }
            var v49 uint64 = v2
            _ = v49
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
