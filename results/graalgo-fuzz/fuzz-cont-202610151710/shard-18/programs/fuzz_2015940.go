package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint32; F2 string }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int8 }
func genFn0(p0 bool, p1 string, p2 uint) int64 {
    v6 := GenS1{F0: 2147483647}
    p2 += 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int32) int32 {
    if false {
        v7 := GenS2{F0: -1, F1: "foo bar"}
        var v8 uint8 = 214
        var v9 float64 = 1.5
        _ = v7
        _ = v8
        _ = v9
    }
    p0 = (p0 & (p0 ^ -2147483648))
    var v10 uint64 = 223
    v11 := []int{-672, 128, 42, 977}
    var v12 int32 = (-(p0))
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return p0
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 uint32 = 2
    var v4 int8 = -85
    var v5 uint64 = 852
    if (100 == -1) {
        fmt.Println(v3)
    }
    fmt.Println(v4, v4)
    v13 := GenS3{F0: false}
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int = fi14
        for fi16 := 0; fi16 < 11; fi16++ {
            if (fi16 == -520) {
                v17 := GenS4{F0: v4}
                var v18 int32 = -1
                if false {
                    if true {
                        v1 -= (float64((1e10 - 1e10)) * (v1 * (v1 / float64(1.5))))
                    }
                }
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v19 := GenS2{F0: 0, F1: "hello"}
    switch "\n\t" {
    case "foo bar":
        _ = len(v2)
    }
    if true {
        v5 = (v5 >> ((4294967295) & 63))
    }
    v20 := GenS5{F0: 0}
    v3 = v3
    v21 := GenS4{F0: v4}
    var v22 float32 = 0.0
    if ("\n\t" > (("hello" + ("" + "")) + "foo bar")) {
        var v23 int16 = 0
        if true {
            var v24 uint8 = 2
            if false {
                var v25 int64 = 1
                fmt.Println(v25)
                switch ("\n\t" + "hello") {
                case "":
                    if true {
                        var v26 uint64 = v5
                        v23 = (v23 + 615)
                        v23 = (v23 - v23)
                        _ = v26
                    } else {
                        var v27 int32 = -681
                        _ = v27
                    }
                }
                _ = v25
            }
            if (!false) {
                v28 := GenS1{F0: 0}
                var v29 int64 = 9223372036854775807
                _ = v28
                _ = v29
            }
            _ = v24
        } else {
            v2[32767] = false
            var v30 uint16 = 493
            for fi31 := 0; fi31 < 17; fi31++ {
                switch fi31 {
                case 128:
                    var v32 uint = 0
                    _ = v32
                case 430:
                    var v33 float32 = 3.14159
                    _ = v33
                case 395:
                    fmt.Println(v1, v30)
                default:
                    if false {
                        v34 := []int{}
                        switch "hello" {
                        case "foo bar":
                            v35 := GenS3{F0: (v5 > 2)}
                            _ = v35
                        case "":
                            var v36 int16 = (v23 + v23)
                            _ = v36
                        }
                        v5 += v5
                        _ = v34
                    }
                }
                var v37 uint64 = 42
                _ = fi31
                _ = v37
            }
            _ = v30
        }
        _ = v23
    }
    if (296 <= 18446744073709551615) {
        var v38 int32 = genFn1(0)
        v2[32767] = false
        v39 := GenS4{F0: 127}
        _ = v38
        _ = v39
    } else {
        switch 128 {
        case 0:
            fmt.Println(v22, v1, v5)
        case 42:
            var v40 uint = 372
            _ = v40
        case 1:
            for fi41 := 0; fi41 < 17; fi41++ {
                v4 += 19
                var v42 int8 = v4
                _ = fi41
                _ = v42
            }
        default:
            var v43 int8 = v4
            _ = v43
        }
    }
    switch "foo bar" {
    case "\n\t":
        v3 += (v3 / (((^(v3))) | 1))
    case "":
        if false {
            var v44 bool = (true && true)
            v45 := v2[32767]
            _ = v44
            _ = v45
        } else {
            v22 = v22
        }
    default:
        var v46 int32 = genFn1(-974)
        _ = v46
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        if false {
            var v48 float32 = v22
            v48 = 0.1
            v49 := GenS6{F0: 1000, F1: v3, F2: ("hello" + "hello")}
            _ = v48
            _ = v49
        }
        switch "foo bar" {
        case "":
            fmt.Println(v5, v4, fi47)
        default:
            switch -2147483648 {
            case 9223372036854775807:
                switch fi47 {
                case 0:
                    if true {
                        if (false && (!true)) {
                            v50 := []int{fi47, fi47, fi47, 2147483647}
                            if ((v3 / ((v3) | 1)) >= 42) {
                                switch (fi47 & 544) {
                                case 127:
                                    fmt.Println(fi47, v4)
                                case 42:
                                    v5 -= (v5 >> ((uint(v5)) & 63))
                                case -998:
                                    if len(v50) > 0 {
                                        v50[0] = -15
                                    }
                                default:
                                    var v51 int8 = v4
                                    _ = v51
                                }
                                v52 := GenS3{F0: (42 > 2147483647)}
                                _ = v52
                            } else {
                                if len(v50) > 0 {
                                    v50[0] = (fi47 & fi47)
                                }
                            }
                            _ = v50
                        } else {
                            var v53 string = "\n\t"
                            _ = v53
                        }
                    } else {
                        var v54 uint16 = 255
                        for fi55 := 0; fi55 < 4; fi55++ {
                            v56 := []int{(fi47 * fi47), fi55}
                            v54 = (v54 * (v54 * (v54 << ((4294967295) & 63))))
                            _ = fi55
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 19; fi57++ {
                            if false {
                                var v58 uint = 338
                                fmt.Println(v58, v22, v58)
                                _ = v58
                            } else {
                                v59 := GenS7{F0: 2147483647, F1: 4294967295}
                                v2[(fi57 + 32768)] = (1 != v54)
                                _ = v59
                            }
                            var v60 int8 = v4
                            var v61 int64 = genFn0(false, "\n\t", 42)
                            _ = fi57
                            _ = v60
                            _ = v61
                        }
                        _ = v54
                    }
                }
            }
        }
        _ = fi47
    }
    if false {
        v3 = (v3 << ((255) & 63))
        if false {
            if true {
                v5 = (^(v5))
                fmt.Println(v5)
            }
            v1 = 1e10
        }
        _ = len(v2)
    } else {
        switch "" {
        case "\n\t":
            if true {
                v3 -= 2
            }
        case "a":
            switch "a" {
            case "a":
                v1 += (-2.25 + (-2.25 / (0.1 / ((((v1 + v1) + -2.25) - v1) - 1e10))))
            case "\n\t":
                var v62 int64 = 20
                _ = v62
            case "hello":
                var v63 int8 = 0
                _ = v63
            default:
                for fi64 := 0; fi64 < 16; fi64++ {
                    var v65 int8 = (v4 * (v4 % ((0) | 1)))
                    _ = fi64
                    _ = v65
                }
            }
        case "":
            switch -9223372036854775808 {
            case 274:
                v66 := GenS8{F0: (v4 - v4)}
                _ = v66
            case 4294967295:
                for fi67 := 0; fi67 < 3; fi67++ {
                    fmt.Println(v4, v1, v5)
                    var v68 float64 = float64(v22)
                    _ = fi67
                    _ = v68
                }
            default:
                if false {
                    for fi69 := 0; fi69 < 5; fi69++ {
                        v70 := []int{fi69}
                        _ = fi69
                        _ = v70
                    }
                    delete(v2, 127)
                } else {
                    fmt.Println(v5)
                    var v71 int64 = 9223372036854775807
                    _ = v71
                }
            }
        default:
            var v72 int8 = 100
            _ = v72
        }
    }
    var v73 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v73
}
