package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 int64 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64 }
func genFn0() uint8 {
    var v7 int64 = 42
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{(fi8 - fi8)}
        var v10 uint32 = 2
        switch "foo bar" {
        case "foo bar":
            switch (^(fi8)) {
            case 32767:
                var v11 uint32 = v10
                _ = v11
            case 4294967295:
                v7 = v7
            default:
                var v12 uint16 = 0
                _ = v12
            }
        case "hello":
            if ("hello" != ("\n\t" + ("hello" + ""))) {
                v7 += (v7 + 1)
                var v13 uint8 = 85
                _ = v13
            } else {
                switch 9223372036854775807 {
                case -9223372036854775808:
                    v14 := GenS1{F0: 1}
                    _ = v14
                case -1:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 uint8 = 255
                        _ = fi15
                        _ = v16
                    }
                }
            }
        case "a":
            switch ((("\n\t" + "a") + ("hello" + "\n\t")) + "a") {
            case "a":
                var v17 uint32 = (v10 << ((2) & 63))
                _ = v17
            case "":
                var v18 uint16 = 42
                _ = v18
            case "foo bar":
                v19 := GenS2{F0: 1}
                _ = v19
            }
        default:
            v20 := GenS3{F0: 0, F1: (v10 / (((v10 ^ 2147483648)) | 1))}
            _ = v20
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if false {
            var v22 uint = 920
            v7 -= v7
            v22 += v22
            _ = v22
        } else {
            var v23 int32 = 1
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := []int{(fi24 * (fi21 - fi21))}
            v7 += (v7 + v7)
            switch ("\n\t" + "hello") {
            case "foo bar":
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
            case "a":
                var v26 float32 = (((1.5 - -2.25) * 0.1) * float32((3.14159 / (3.14159 / 0.1))))
                if (0 < v7) {
                    v26 -= -2.25
                    v27 := GenS4{F0: -2147483648, F1: 127, F2: v7}
                    _ = v27
                }
                v28 := GenS4{F0: 0, F1: 2, F2: v7}
                _ = v26
                _ = v28
            case "hello":
                v29 := GenS4{F0: -348, F1: 135, F2: 1}
                switch "hello" {
                case "\n\t":
                    if len(v25) > 0 {
                        v25[0] = -721
                    }
                case "foo bar":
                    if (!(178 <= 2)) {
                        var v30 int32 = 653
                        if (-128 <= -128) {
                            v30 = -671
                            v7 += v7
                            if (127 < 2) {
                                _ = len(v25)
                                v30 = -1
                            } else {
                                v25 = append(v25, fi24)
                                fmt.Println(v7, v30)
                                var v31 uint = 819
                                _ = v31
                            }
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            var v32 uint32 = 141
                            _ = v32
                        }
                        if (v30 < v30) {
                            v33 := GenS5{F0: (("a" + "hello") + "foo bar"), F1: 65535}
                            v34 := GenS3{F0: v7, F1: 4294967295}
                            _ = v33
                            _ = v34
                        } else {
                            fmt.Println(v7)
                            if len(v25) > 0 {
                                v25[0] = 127
                            }
                        }
                        _ = v30
                    }
                default:
                    if (1 > -521) {
                        v25 = append(v25, fi24)
                        if (1 == 400) {
                            v35 := GenS5{F0: "a", F1: 932}
                            if (1e10 <= 1e10) {
                                v36 := GenS6{F0: 2147483647, F1: 1, F2: 1}
                                _ = v36
                            } else {
                                if len(v25) > 0 {
                                    v25[0] = 9223372036854775807
                                }
                            }
                            _ = v35
                        }
                        if len(v25) > 0 {
                            v25[0] = -1
                        }
                    }
                }
                v37 := GenS7{F0: 42}
                v38 := GenS8{F0: 100, F1: 243}
                v7 += (v7 * 1)
                _ = v29
                _ = v37
                _ = v38
            }
            _ = fi24
            _ = v25
        }
        _ = fi21
    }
    fmt.Println(v7)
    _ = v7
    return 161
}
func genFn1(p0 int16, p1 int, p2 uint64) uint {
    fmt.Println(p2)
    v39 := []int{}
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS9{F0: p1, F1: (p2 & p2)}
        p1 -= (fi40 & fi40)
        fmt.Println(p2, p0, p1)
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return 2
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 127
    var v3 uint16 = 42
    var v4 int8 = 0
    var v5 float32 = ((float32(((0.0 - 0.1) / 0.1)) * (0.1 + 0.0)) * ((-2.25 * -2.25) * (0.1 / -2.25)))
    var v6 int8 = (-(v4))
    v42 := []int{}
    var v43 uint32 = 318
    v2 += (v2 & (v2 & v2))
    if ((v5 / v5) != 1.5) {
        fmt.Println(v5, v6)
        fmt.Println(v1)
    } else {
        var v44 uint8 = v2
        var v45 uint64 = 381
        v3 = 2
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        _ = len(v42)
        v2 = 2
        if ((^(v3)) >= v3) {
            v2 = (v2 % (((^(v2))) | 1))
            switch (("hello" + "hello") + "") {
            case "foo bar":
                v47 := GenS6{F0: uint32(v1), F1: 0, F2: 2}
                _ = v47
            case "a":
                _ = len(v42)
            }
            v3 += uint16(v2)
        } else {
            for fi48 := 0; fi48 < 17; fi48++ {
                v1 = (^(v1))
                fmt.Println(v6, v5)
                v4 -= v6
                if false {
                    if (!(true && ((true && true) || (v43 != v43)))) {
                        if true {
                            if ((v6 + (-(v6))) > (v4 >> ((v1) & 63))) {
                                var v49 uint8 = (v2 * (v2 >> ((987) & 63)))
                                _ = v49
                            }
                            if (!true) {
                                fmt.Println(v3, v2)
                            } else {
                                v6 -= v4
                                v1 = (v1 % ((v1) | 1))
                            }
                            var v50 int = 830
                            v51 := GenS10{F0: -236}
                            switch "a" {
                            case "\n\t":
                                v1 = 1
                            case "":
                                if len(v42) > 0 {
                                    v42[0] = fi48
                                }
                            case "a":
                                if (("a" + ("\n\t" + "foo bar")) == "a") {
                                    var v52 uint = genFn1(959, v50, 4294967295)
                                    _ = v52
                                }
                            }
                            _ = v50
                            _ = v51
                        }
                        var v53 float32 = float32(3.14159)
                        if (((0.1 < (1.5 + (1e10 * (0.1 / (1e10 - 1.5))))) && (v53 != (v5 - 3.14159))) && true) {
                            v54 := GenS9{F0: (fi48 * fi48), F1: 0}
                            var v55 int32 = 42
                            _ = v54
                            _ = v55
                        }
                        _ = v53
                    } else {
                        var v56 bool = true
                        if v56 {
                            var v57 uint = genFn1(32767, (^(fi48)), uint64(v1))
                            var v58 uint32 = (v43 << ((v57) & 63))
                            if true {
                                v42 = append(v42, fi46)
                            } else {
                                if (v56 || v56) {
                                    _ = len(v42)
                                    if len(v42) > 0 {
                                        v42[0] = 0
                                    }
                                }
                                _ = len(v42)
                                fmt.Println(fi48, v43, v4)
                                v59 := GenS2{F0: genFn0()}
                                v60 := GenS6{F0: v58, F1: 0, F2: 2}
                                v61 := GenS1{F0: 42}
                                v62 := GenS6{F0: 42, F1: -9223372036854775808, F2: 158}
                                _ = v59
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            _ = v57
                            _ = v58
                        } else {
                            if len(v42) > 0 {
                                v42[0] = fi48
                            }
                            v2 -= (v2 % (((v2 | v2)) | 1))
                        }
                        _ = v56
                    }
                    if (fi48 != (^(fi48))) {
                        if (2147483647 != v43) {
                            switch "hello" {
                            case "":
                                var v63 string = "hello"
                                v3 -= (v3 ^ v3)
                                var v64 float64 = float64((v5 + v5))
                                _ = v63
                                _ = v64
                            default:
                                if len(v42) >= 2 {
                                    v42 = v42[1:len(v42)]
                                }
                            }
                            v42 = append(v42, fi46)
                            var v65 uint = (v1 ^ genFn1(410, fi46, 951))
                            _ = v65
                        } else {
                            switch "a" {
                            case "hello":
                                v66 := GenS11{F0: 1}
                                _ = v66
                            case "foo bar":
                                var v67 uint = 255
                                _ = v67
                            case "":
                                fmt.Println(v1, v6, v43)
                            }
                            if false {
                                v42 = append(v42, (fi46 / (((fi46 - fi46)) | 1)))
                            }
                        }
                        v42 = append(v42, fi48)
                        fmt.Println(v43, v4, v43)
                    } else {
                        fmt.Println(v2)
                        v1 -= (v1 ^ 65535)
                        _ = len(v42)
                    }
                    v6 += v6
                }
                var v68 int64 = -297
                var v69 uint64 = 386
                v70 := GenS9{F0: fi48, F1: (v69 & (v69 ^ v69))}
                _ = fi48
                _ = v68
                _ = v69
                _ = v70
            }
            var v71 int8 = 65
            _ = v71
        }
        _ = fi46
    }
    if (true || (true && false)) {
        for fi72 := 0; fi72 < 3; fi72++ {
            _ = len(v42)
            switch int(v43) {
            case -2147483648:
                var v73 uint16 = (v3 | 32767)
                _ = v73
            }
            var v74 uint64 = 18446744073709551615
            _ = fi72
            _ = v74
        }
    }
    fmt.Println(v5)
    if true {
        fmt.Println(v3, v43)
        v75 := GenS12{F0: 0}
        v76 := GenS6{F0: (v43 * 340), F1: 0, F2: 217}
        _ = v75
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
