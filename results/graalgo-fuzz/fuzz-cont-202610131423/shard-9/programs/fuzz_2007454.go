package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
func genFn0(p0 int8) int16 {
    if false {
        switch 526 {
        case 721:
            var v7 int = 2147483647
            _ = v7
        case 128:
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint32 = 2147483647
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := []int{(fi8 | (^(fi10))), 996}
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        default:
            v12 := []int{}
            _ = v12
        }
        p0 = (p0 / ((0) | 1))
        if true {
            v13 := GenS1{F0: 346, F1: 150}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 uint = 1
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v17 := GenS1{F0: 432, F1: 32767}
    var v18 uint = 42
    v19 := []int{-59, 127, 0, 32767}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return -32768
}

func main() {
    var v1 bool = true
    var v2 float32 = 3.14159
    var v3 int = 42
    v4 := make(map[int]int)
    var v5 int32 = -1
    v6 := []string{("" + "hello"), "", ("\n\t" + ("a" + ""))}
    for fi20 := 0; fi20 < 7; fi20++ {
        if v1 {
            if ("hello" < "foo bar") {
                var v21 uint16 = 83
                if (!v1) {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        var v23 int = 1
                        _ = len(v4)
                        _ = fi22
                        _ = v23
                    }
                } else {
                    var v24 int8 = -82
                    v25 := GenS2{F0: 0, F1: (v5 << ((65535) & 63)), F2: 42}
                    if ((899 == int64(v24)) && (1.5 == 3.14159)) {
                        v26 := GenS3{F0: (v21 % (((v21 | 32767)) | 1)), F1: -92}
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            var v28 bool = v1
                            v29 := GenS4{F0: 2}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        if (v2 != ((v2 - v2) - ((1.5 / v2) / 3.14159))) {
                            switch 0 {
                            case -752:
                                var v30 uint = 899
                                _ = v30
                            case -9223372036854775808:
                                v31 := GenS3{F0: 431, F1: 42}
                                _ = v31
                            case 128:
                                fmt.Println(v3)
                            default:
                                v6 = append(v6, (("" + ("hello" + "\n\t")) + ("a" + ("a" + "a"))))
                            }
                        } else {
                            v32 := GenS1{F0: v5, F1: 2}
                            _ = v32
                        }
                        var v33 int8 = (v24 >> ((42) & 63))
                        _ = v33
                    }
                    _ = v24
                    _ = v25
                }
                _ = len(v6)
                _ = v21
            }
            v1 = (245 >= 482)
        } else {
            if v1 {
                var v34 int8 = -52
                switch "a" {
                case "a":
                    switch ("\n\t" + "foo bar") {
                    case "":
                        v35 := GenS4{F0: 18446744073709551615}
                        _ = v35
                    case "foo bar":
                        for fi36 := 0; fi36 < 8; fi36++ {
                            v3 += (fi20 / ((fi20) | 1))
                            _ = fi36
                        }
                    case "a":
                        for fi37 := 0; fi37 < 14; fi37++ {
                            var v38 float64 = (-2.25 + 3.14159)
                            var v39 int32 = (v5 ^ v5)
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                case "hello":
                    var v40 float32 = (v2 * float32(0.0))
                    _ = v40
                case "":
                    v4[fi20] = (v3 << ((610) & 63))
                default:
                    var v41 float64 = 1e10
                    _ = v41
                }
                _ = v34
            }
            switch ("" + "") {
            case "foo bar":
                v3 -= 0
            }
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 float32 = float32(((3.14159 - 3.14159) / (1e10 / (-2.25 * 1.5))))
                switch v3 {
                case 42:
                    var v44 uint16 = 2
                    _ = v44
                case 9223372036854775807:
                    fmt.Println(v5, fi20, v2)
                case -9223372036854775808:
                    if true {
                        var v45 int32 = -657
                        fmt.Println(fi42)
                        fmt.Println(fi20, v2, v2)
                        _ = v45
                    } else {
                        v43 = v2
                    }
                default:
                    v5 += (v5 - v5)
                }
                fmt.Println(v43, v5)
                _ = fi42
                _ = v43
            }
        }
        var v46 uint8 = 91
        v1 = false
        _ = fi20
        _ = v46
    }
    fmt.Println(v5)
    if (v3 >= 32768) {
        var v47 int = (v3 + (v3 | v3))
        if (27 >= -128) {
            var v48 int16 = 0
            _ = v48
        } else {
            fmt.Println(v47, v2)
        }
        _ = v47
    }
    var v49 int = v3
    var v50 uint32 = 206
    v50 -= (v50 - v50)
    for fi51 := 0; fi51 < 5; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            v53 := GenS5{F0: 2}
            var v54 int64 = 937
            _ = len(v6)
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = fi51
    }
    var v55 int = int(v5)
    if (v5 > (v5 | v5)) {
        for fi56 := 0; fi56 < 17; fi56++ {
            v57 := GenS6{F0: 690}
            for fi58 := 0; fi58 < 5; fi58++ {
                v59 := GenS5{F0: 656}
                _ = fi58
                _ = v59
            }
            v60 := GenS5{F0: 42}
            _ = fi56
            _ = v57
            _ = v60
        }
        var v61 uint64 = 4294967295
        var v62 int8 = -128
        _ = v61
        _ = v62
    } else {
        v5 -= (v5 | v5)
        for fi63 := 0; fi63 < 10; fi63++ {
            if (!(!(!(float32(-2.25) >= (0.1 / 1.5))))) {
                if v1 {
                    v50 -= (v50 - (v50 << ((2) & 63)))
                } else {
                    v6 = append(v6, ("" + "hello"))
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v2 += v2
                        _ = fi64
                    }
                    fmt.Println(v50, v5, v1)
                }
                v2 = v2
                fmt.Println(v2, v55)
            } else {
                if (!(1 > 33)) {
                    fmt.Println(v1)
                    v3 -= -2147483648
                    fmt.Println(v50, v3, v55)
                } else {
                    _ = len(v4)
                    var v65 string = "foo bar"
                    _ = v65
                }
                v66 := v4[-1]
                for fi67 := 0; fi67 < 1; fi67++ {
                    var v68 uint64 = 42
                    _ = fi67
                    _ = v68
                }
                _ = v66
            }
            _ = fi63
        }
        fmt.Println(v1, v3, v3)
    }
    var v69 int32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v55
    _ = v6
    _ = v69
}
