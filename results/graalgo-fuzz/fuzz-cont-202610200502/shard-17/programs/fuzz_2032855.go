package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint; F2 string }
type GenS7 struct { F0 uint8; F1 int8; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int64, p1 int) uint64 {
    v7 := []int{(p1 / (((p1 % (((p1 | (^(p1)))) | 1))) | 1))}
    v8 := GenS1{F0: (p1 & (-(p1)))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 768
}
func genFn1(p0 bool, p1 string) uint32 {
    switch p1 {
    case "foo bar":
        switch p1 {
        case "\n\t":
            p0 = p0
        default:
            if (2147483648 < 0) {
                var v9 int16 = 200
                switch ("" + "a") {
                case "hello":
                    if true {
                        fmt.Println(p1, p0, v9)
                        if true {
                            var v10 string = "a"
                            var v11 bool = ((42 <= -1) && (32768 <= 32767))
                            v12 := GenS2{F0: 483, F1: 1, F2: p0}
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    }
                case "a":
                    var v13 uint32 = 2147483647
                    _ = v13
                case "":
                    v14 := GenS1{F0: -16}
                    _ = v14
                default:
                    fmt.Println(p0)
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    p1 = ("" + p1)
                    v16 := GenS2{F0: 440, F1: 2147483647, F2: p0}
                    v17 := GenS3{F0: p0, F1: 171}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v9
            } else {
                p0 = p0
                var v18 uint64 = 1
                if p0 {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 int8 = 42
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 int64 = 42
                            p0 = false
                            v18 += (v18 >> ((255) & 63))
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 uint16 = 0
                        switch fi23 {
                        case 2147483647:
                            if (!((!(p0 && true)) && (true || p0))) {
                                var v25 float32 = 1e10
                                _ = v25
                            } else {
                                fmt.Println(p0, p1)
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    if (232 != fi23) {
                                        var v27 uint8 = 0
                                        if (2147483647 > 170) {
                                            v28 := GenS4{F0: (!false), F1: -2147483648, F2: (v24 << ((1) & 63))}
                                            p1 = "foo bar"
                                            var v29 bool = p0
                                            _ = v28
                                            _ = v29
                                        }
                                        switch int(v18) {
                                        case 2147483647:
                                            v24 += 32767
                                        case -9223372036854775808:
                                            v30 := GenS3{F0: (1 <= 837), F1: 42}
                                            _ = v30
                                        case -2147483648:
                                            fmt.Println(v24)
                                        }
                                        _ = v27
                                    }
                                    _ = fi26
                                }
                            }
                        case -1:
                            v18 -= v18
                        }
                        _ = fi23
                        _ = v24
                    }
                    var v31 int = 32768
                    _ = v31
                }
                _ = v18
            }
        }
    case "":
        p1 = (((p1 + p1) + p1) + "\n\t")
    }
    if (-2.25 < 1.5) {
        p0 = (!(!true))
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 462
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 string = ("a" + "\n\t")
    var v4 uint16 = 32767
    v5 := make(map[uint32]string)
    v6 := make(map[string]uint32)
    switch ("foo bar" + v3) {
    case "\n\t":
        v32 := GenS5{F0: -2147483648, F1: 42, F2: 42}
        _ = v32
    }
    if true {
        switch -760 {
        case 128:
            v3 = (v3 + "hello")
        case 42:
            if ((42 > -307) && true) {
                var v33 uint8 = 255
                switch v3 {
                case "":
                    v4 = v4
                case "hello":
                    var v34 uint32 = 42
                    _ = v34
                case "a":
                    v4 = (v4 >> ((18446744073709551615) & 63))
                default:
                    fmt.Println(v33, v4)
                }
                v4 -= (v4 * v4)
                _ = v33
            } else {
                v3 = (v3 + v3)
            }
        case 651:
            if (-1 >= -437) {
                v35 := GenS4{F0: false, F1: 2147483647, F2: v4}
                v4 += v4
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 17; fi36++ {
            if (!(true && false)) {
                v1 = 18446744073709551615
            } else {
                for fi37 := 0; fi37 < 4; fi37++ {
                    var v38 uint32 = 0
                    _ = fi37
                    _ = v38
                }
                var v39 string = "\n\t"
                _ = v39
            }
            _ = fi36
        }
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 uint16 = (v4 / (((^(v4))) | 1))
            _ = fi40
            _ = v41
        }
    }
    v42 := GenS3{F0: true, F1: 127}
    var v43 float32 = (0.0 / 1e10)
    if (42 > -1) {
        v1 -= 0
    }
    if ((v43 / (0.0 - v43)) > v43) {
        v44 := GenS5{F0: 1, F1: 0, F2: 42}
        v45 := []int{32767, -9223372036854775808, 32768}
        v45 = append(v45, 9223372036854775807)
        _ = v44
        _ = v45
    }
    v3 = "a"
    v4 += (v4 & (v4 % ((1) | 1)))
    if false {
        if (!false) {
            var v46 uint32 = 273
            var v47 int64 = -1
            switch v3 {
            case "\n\t":
                var v48 bool = (!(1 > -2147483648))
                _ = v48
            case "foo bar":
                var v49 int64 = (v47 + v47)
                _ = v49
            case "hello":
                v50 := GenS6{F0: 255, F1: 0, F2: (((v3 + v3) + ("\n\t" + v3)) + v3)}
                _ = v50
            }
            _ = v46
            _ = v47
        }
    }
    var v51 uint = uint(v4)
    switch -758 {
    case 128:
        switch -394 {
        case 787:
            var v52 string = v3
            _ = v52
        case -574:
            switch v3 {
            case "hello":
                var v53 uint = v51
                _ = v53
            case "foo bar":
                var v54 int32 = 42
                _ = v54
            }
        default:
            if (!((v51 >= 0) || (2 <= 1))) {
                v6[v3] = genFn1((false && true), (((((("hello" + "hello") + "a") + v3) + v3) + "a") + v3))
                v3 = ""
            } else {
                if false {
                    switch 420 {
                    case 844:
                        v4 -= (v4 + (^(v4)))
                    case -120:
                        var v55 int8 = 127
                        _ = v55
                    }
                    v4 = 568
                    v56 := GenS7{F0: 255, F1: -67, F2: 100}
                    _ = v56
                } else {
                    v57 := []int{-1}
                    v58 := v2[2]
                    _ = v57
                    _ = v58
                }
                var v59 string = v3
                _ = v59
            }
        }
    case 4294967295:
        if (1 <= -1) {
            var v60 uint = (^(v51))
            var v61 bool = false
            v62 := v6["\n\t"]
            _ = v60
            _ = v61
            _ = v62
        }
    default:
        v63 := GenS8{F0: 127, F1: (^(v1))}
        _ = v63
    }
    fmt.Println(v4)
    var v64 int = int(v4)
    v65 := GenS7{F0: 255, F1: 127, F2: -1}
    v66 := GenS5{F0: (v64 % (((v64 ^ (^(v64)))) | 1)), F1: -9223372036854775808, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
