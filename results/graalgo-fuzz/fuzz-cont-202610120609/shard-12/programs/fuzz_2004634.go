package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint32, p1 uint64, p2 string) uint64 {
    v7 := GenS1{F0: 0}
    var v8 float64 = float64(1.5)
    for fi9 := 0; fi9 < 14; fi9++ {
        p1 -= 0
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint16 = 254
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{-470, 848}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    return p1
}
func genFn1() int {
    var v13 int32 = -2147483648
    if false {
        if true {
            switch 4294967295 {
            case -1:
                var v14 int64 = -112
                _ = v14
            case 127:
                var v15 int32 = (v13 - (v13 >> ((1) & 63)))
                _ = v15
            }
            var v16 uint64 = 1
            v17 := []int{343, -614, 32768}
            _ = v16
            _ = v17
        } else {
            v18 := GenS1{F0: 42}
            _ = v18
        }
    }
    _ = v13
    return 127
}

func main() {
    var v1 float32 = (0.1 * 3.14159)
    v2 := []int{-265, 0, -731, 32768}
    var v3 int64 = 42
    var v4 int64 = (v3 | (v3 % ((412) | 1)))
    var v5 uint = 2
    var v6 uint64 = 42
    if (135 != 42) {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint8 = 128
            _ = fi19
            _ = v20
        }
    } else {
        if (221 <= 106) {
            if (1 != 32767) {
                var v21 uint = v5
                var v22 uint = (^(v21))
                _ = v21
                _ = v22
            }
            fmt.Println(v6, v4)
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                switch (fi23 * -2147483648) {
                case 845:
                    v24 := GenS1{F0: -1}
                    _ = v24
                case 32767:
                    if (!((2147483647 != 42) || false)) {
                        var v25 string = "\n\t"
                        v26 := GenS2{F0: v3, F1: 1, F2: -1}
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint8 = 255
                        _ = v27
                    }
                case 32768:
                    for fi28 := 0; fi28 < 17; fi28++ {
                        switch -654 {
                        case -2147483648:
                            var v29 bool = true
                            _ = v29
                        case 889:
                            switch ("hello" + ("foo bar" + "")) {
                            case "a":
                                fmt.Println(v4)
                            case "\n\t":
                                var v30 float32 = v1
                                var v31 float64 = 0.0
                                v32 := GenS1{F0: 1}
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        case 127:
                            var v33 int64 = (v4 + (v4 << (((v5 << ((735) & 63))) & 63)))
                            var v34 uint32 = 2147483648
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS1{F0: 166}
                        v36 := GenS1{F0: 19}
                        _ = fi28
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi23
            }
        }
        v37 := GenS1{F0: 140}
        _ = v37
    }
    v4 = v4
    if false {
        v5 = (v5 ^ v5)
    }
    var v38 int8 = -128
    for fi39 := 0; fi39 < 3; fi39++ {
        fmt.Println(v1)
        if false {
            if true {
                _ = len(v2)
                v2 = append(v2, (fi39 ^ 1))
            }
            v38 -= v38
            if ((v4 % ((int64(fi39)) | 1)) >= -9223372036854775808) {
                if (!(100 > 1)) {
                    if (-1 != 0) {
                        if true {
                            fmt.Println(v3)
                            v40 := GenS2{F0: v3, F1: 4294967295, F2: (v38 + (v38 << ((v5) & 63)))}
                            _ = v40
                        }
                        v3 += v4
                    } else {
                        for fi41 := 0; fi41 < 7; fi41++ {
                            switch "" {
                            case "":
                                v42 := GenS2{F0: int64(v5), F1: 997, F2: (-(v38))}
                                _ = v42
                            case "foo bar":
                                var v43 int = (fi39 >> (((v5 * 42)) & 63))
                                _ = v43
                            case "\n\t":
                                var v44 uint64 = v6
                                _ = v44
                            }
                            v45 := GenS1{F0: 1000}
                            _ = fi41
                            _ = v45
                        }
                        v46 := GenS2{F0: (v3 / ((v3) | 1)), F1: 700, F2: (v38 / (((v38 + (v38 / ((v38) | 1)))) | 1))}
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi47 := 0; fi47 < 14; fi47++ {
                                if true {
                                    var v48 int = fi39
                                    v2 = append(v2, fi47)
                                    _ = v48
                                } else {
                                    if false {
                                        if (0 != 32767) {
                                            v49 := GenS3{F0: 0, F1: 2147483647}
                                            fmt.Println(fi39)
                                            _ = v49
                                        } else {
                                            var v50 int32 = -630
                                            _ = v50
                                        }
                                    } else {
                                        var v51 uint64 = (v6 << (((v5 << (((v5 / ((v5) | 1))) & 63))) & 63))
                                        if len(v2) > 0 {
                                            v2[0] = -261
                                        }
                                        _ = v51
                                    }
                                }
                                v52 := GenS4{F0: true}
                                _ = fi47
                                _ = v52
                            }
                        } else {
                            v38 = 5
                        }
                        _ = v46
                    }
                    for fi53 := 0; fi53 < 1; fi53++ {
                        v3 -= (-(v3))
                        _ = fi53
                    }
                    v6 = v6
                } else {
                    fmt.Println(v5, fi39)
                    v3 = (v4 + v4)
                }
                v6 = 696
                fmt.Println(v6, v3)
            }
        }
        _ = fi39
    }
    switch genFn1() {
    case 32767:
        _ = len(v2)
    case 32768:
        _ = len(v2)
    }
    fmt.Println(v38, v38)
    if true {
        v2 = append(v2, 128)
    }
    if (!(128 < 2)) {
        if (!(!true)) {
            var v54 float64 = 1e10
            fmt.Println(v1, v3)
            _ = v54
        }
        var v55 float64 = ((1.5 + 1.5) - 3.14159)
        switch "a" {
        case "":
            var v56 int = -88
            _ = v56
        }
        _ = v55
    } else {
        var v57 int32 = 120
        _ = v57
    }
    v1 += float32(1.5)
    var v58 uint32 = 65535
    var v59 float32 = (0.0 + v1)
    fmt.Println(v1, v1, v4)
    v60 := GenS5{F0: 173, F1: 1000}
    if true {
        fmt.Println(v4)
        if true {
            for fi61 := 0; fi61 < 9; fi61++ {
                if true {
                    switch ("foo bar" + ("hello" + ("hello" + ""))) {
                    case "\n\t":
                        v4 = 9223372036854775807
                    case "a":
                        for fi62 := 0; fi62 < 8; fi62++ {
                            _ = len(v2)
                            v63 := GenS6{F0: 1, F1: (v6 / ((4294967295) | 1)), F2: 1}
                            switch "hello" {
                            case "":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v64 := GenS7{F0: true}
                                _ = v64
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                    var v65 uint8 = uint8(v5)
                    v3 = v3
                    _ = v65
                }
                _ = fi61
            }
            var v66 int = genFn1()
            _ = v66
        } else {
            v67 := GenS8{F0: v38, F1: ("foo bar" + ""), F2: int16(v6)}
            var v68 int64 = (v4 ^ (^(v4)))
            _ = v67
            _ = v68
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v69 uint16 = 501
    for fi70 := 0; fi70 < 4; fi70++ {
        fmt.Println(v69, v4, v59)
        var v71 string = ""
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v69
}
