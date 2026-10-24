package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 string; F1 uint; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int32; F2 string }
type GenS7 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []uint32{}
    if len(v3) > 0 {
        v3[0] = 844
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "hello" {
    case "\n\t":
        switch "foo bar" {
        case "":
            v4 := GenS1{F0: 65535, F1: -2147483648, F2: -1}
            _ = v4
        case "\n\t":
            var v5 uint16 = 255
            _ = v5
        case "a":
            var v6 int16 = -270
            _ = v6
        default:
            var v7 int32 = 1
            _ = v7
        }
    case "":
        if true {
            var v8 int8 = -82
            var v9 int8 = v8
            _ = v8
            _ = v9
        }
    }
    v10 := GenS2{F0: "", F1: 2, F2: (-7 != 2147483647)}
    var v11 int32 = 2147483647
    if false {
        v12 := GenS3{F0: (-(v11))}
        v11 += (^(v11))
        fmt.Println(v11)
        _ = v12
    }
    if true {
        var v13 float64 = ((3.14159 * (((-2.25 / 0.1) + 0.1) * (float64(1.5) / 1.0))) + ((1e10 * 1.5) + (-2.25 / (float64(float32(3.14159)) * 1.5))))
        var v14 uint32 = 913
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            switch "hello" {
            case "\n\t":
                var v16 bool = true
                _ = v16
            case "hello":
                v11 = v11
            }
            v1[199] = (255 <= 18446744073709551615)
            _ = fi15
        }
        if (65535 >= 0) {
            delete(v2, ("" + "a"))
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int32 = v11
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int16 = -655
                if true {
                    _ = len(v1)
                    v20 -= -903
                }
                if (2 != 0) {
                    if (42 != 1) {
                        fmt.Println(v20, v11, v11)
                        _ = len(v3)
                    }
                    if true {
                        fmt.Println(v20)
                        if true {
                            switch (fi19 % (((fi19 * fi19)) | 1)) {
                            case 4294967295:
                                v21 := GenS3{F0: (^(v11))}
                                _ = v21
                            case 198:
                                for fi22 := 0; fi22 < 14; fi22++ {
                                    v23 := GenS4{F0: 65535, F1: 0}
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            v24 := GenS2{F0: ("foo bar" + ("hello" + ("" + ""))), F1: 893, F2: (fi19 < -634)}
                            _ = v24
                        }
                        v25 := GenS5{F0: true, F1: 935, F2: 1}
                        _ = v25
                    }
                } else {
                    if (!false) {
                        v20 += v20
                    } else {
                        if (false && (true && (!true))) {
                            var v26 uint64 = 2
                            var v27 string = "\n\t"
                            _ = v26
                            _ = v27
                        }
                        v20 += (v20 | (v20 << ((4294967295) & 63)))
                        switch 9223372036854775807 {
                        case 1:
                            v11 = 42
                        case 32768:
                            var v28 float64 = ((1e10 + -2.25) / -2.25)
                            _ = v28
                        case 9223372036854775807:
                            v20 = 452
                        }
                    }
                }
                _ = fi19
                _ = v20
            }
        } else {
            if (true && false) {
                v11 -= v11
                var v29 float64 = 1e10
                fmt.Println(v11)
                _ = v29
            } else {
                var v30 int8 = -128
                _ = v30
            }
            if false {
                switch 4294967295 {
                case 32768:
                    for fi31 := 0; fi31 < 3; fi31++ {
                        fmt.Println(v11)
                        _ = fi31
                    }
                case 0:
                    switch "foo bar" {
                    case "":
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            var v32 uint64 = 2
                            _ = v32
                        case "a":
                            for fi33 := 0; fi33 < 8; fi33++ {
                                fmt.Println(v11)
                                _ = fi33
                            }
                        default:
                            for fi34 := 0; fi34 < 8; fi34++ {
                                v35 := GenS3{F0: v11}
                                v11 = -503
                                _ = fi34
                                _ = v35
                            }
                        }
                    default:
                        v36 := GenS5{F0: true, F1: 65535, F2: 2}
                        _ = v36
                    }
                default:
                    v11 -= 2147483647
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    if (true || false) {
                        v38 := GenS2{F0: (("hello" + (("a" + "a") + ("" + "hello"))) + ""), F1: 4294967295, F2: (!((((3.14159 / (1e10 - (1.5 - 0.1))) / -2.25) - (1.5 - (((1e10 * 1e10) + -2.25) / (3.14159 / 1.5)))) <= (1e10 - ((1.5 / 1.0) / ((-2.25 / 1.0) - 1.5)))))}
                        _ = v38
                    }
                    var v39 int64 = 0
                    _ = fi37
                    _ = v39
                }
            } else {
                if (false && false) {
                    v40 := GenS1{F0: 0, F1: 127, F2: 2147483647}
                    _ = v40
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v41 string = ""
                _ = v41
            }
        }
        if true {
            v42 := GenS6{F0: 0, F1: v11, F2: "a"}
            for fi43 := 0; fi43 < 15; fi43++ {
                fmt.Println(v11)
                fmt.Println(v11, v11)
                if len(v3) > 0 {
                    v3[0] = 42
                }
                _ = fi43
            }
            for fi44 := 0; fi44 < 3; fi44++ {
                v1[42] = ((2147483647 == v11) || ((504 >= 42) && (!true)))
                _ = fi44
            }
            _ = v42
        } else {
            if (((!(!false)) && false) || ((42 <= 18446744073709551615) && true)) {
                for fi45 := 0; fi45 < 9; fi45++ {
                    fmt.Println(fi45, fi45)
                    v11 += v11
                    _ = fi45
                }
            } else {
                v11 = v11
                var v46 int8 = 100
                _ = v46
            }
        }
    }
    if true {
        for fi47 := 0; fi47 < 5; fi47++ {
            var v48 float64 = float64(0.0)
            v48 = v48
            v48 -= v48
            _ = fi47
            _ = v48
        }
    } else {
        var v49 uint8 = 128
        var v50 int64 = 973
        _ = v49
        _ = v50
    }
    v51 := GenS5{F0: true, F1: 65535, F2: 997}
    if true {
        v11 -= (^(v11))
        v11 = (v11 << ((18446744073709551615) & 63))
    }
    var v52 float32 = 1e10
    for fi53 := 0; fi53 < 20; fi53++ {
        switch -9223372036854775808 {
        case 1:
            var v54 uint16 = 42
            _ = v54
        case 32767:
            var v55 int32 = (v11 + v11)
            _ = v55
        }
        v11 -= 885
        if true {
            var v56 uint32 = 633
            var v57 int = int(v56)
            _ = v56
            _ = v57
        }
        _ = fi53
    }
    var v58 uint32 = 1
    v59 := GenS7{F0: 42}
    fmt.Println(v58, v11, v11)
    var v60 uint8 = 1
    v11 = (v11 >> ((1) & 63))
    v52 -= (v52 * v52)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v58
    _ = v59
    _ = v60
}
