package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint32 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if (!(false || false)) {
            v5 := GenS1{F0: 42}
            var v6 float64 = ((1e10 * 0.0) * ((((1.5 - 0.0) / -2.25) * ((float64(1e10) * float64(1.5)) / (float64(1e10) / 1e10))) + (0.1 - 1.5)))
            _ = v5
            _ = v6
        } else {
            var v7 int32 = 2147483647
            v8 := GenS1{F0: 65535}
            _ = v7
            _ = v8
        }
        if false {
            var v9 uint64 = 1
            _ = v9
        }
        var v10 uint8 = 255
        _ = fi4
        _ = v10
    }
    var v11 uint16 = 1
    _ = v11
    return 42
}
func genFn1(p0 int, p1 int) uint16 {
    switch p0 {
    case 127:
        if ((p1 | 1) <= -452) {
            if false {
                var v12 uint8 = 255
                v13 := GenS2{F0: 0}
                fmt.Println(v12, v12)
                _ = v12
                _ = v13
            }
        }
    case 9223372036854775807:
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint = 323
            _ = fi14
            _ = v15
        }
    default:
        if true {
            p0 = 4294967295
            var v16 uint16 = 2
            _ = v16
        }
    }
    v17 := []int{p1}
    switch (p0 << ((747) & 63)) {
    case 252:
        p0 = p0
    }
    var v18 float32 = 3.14159
    p1 = 15
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 388
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]uint32)
    var v3 uint = 0
    v19 := GenS1{F0: 65535}
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := []int{(fi20 * 42), fi20, (fi20 + 128)}
        if len(v21) > 0 {
            v21[0] = 42
        }
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        if ((!(0 >= 42)) && (true && false)) {
            delete(v2, fi22)
            if (!(true && ((140 <= 127) && false))) {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v2[fi23] = 42
                    switch int(v3) {
                    case 32768:
                        v24 := GenS1{F0: 362}
                        _ = v24
                    case -2147483648:
                        if true {
                            var v25 uint8 = 79
                            switch ("foo bar" + "foo bar") {
                            case "":
                                var v26 uint8 = v25
                                _ = v26
                            case "foo bar":
                                var v27 uint32 = 2
                                _ = v27
                            default:
                                v28 := v2[int(v3)]
                                _ = v28
                            }
                            _ = v25
                        }
                    default:
                        if false {
                            switch "foo bar" {
                            case "hello":
                                v1 = (v1 * v1)
                            case "":
                                _ = len(v2)
                                v1 -= (1e10 / 1.0)
                            case "a":
                                v29 := GenS2{F0: v3}
                                _ = v29
                            }
                            if (321 != -1) {
                                fmt.Println(v1, fi23)
                                v3 += v3
                                v1 += v1
                            }
                            var v30 float64 = ((0.1 / float64(1.5)) - (((1.5 + (-2.25 * 0.1)) + (float64(0.1) + 0.1)) + -2.25))
                            _ = v30
                        }
                    }
                    v3 += (v3 % ((824) | 1))
                    _ = fi23
                }
            } else {
                fmt.Println(v1, v3, v1)
                delete(v2, fi22)
                if (-1 <= -228) {
                    var v31 uint8 = 93
                    var v32 int64 = -1
                    _ = v31
                    _ = v32
                }
            }
        } else {
            v1 = (((1e10 - 1e10) / v1) - v1)
            v1 += v1
            for fi33 := 0; fi33 < 16; fi33++ {
                if ((false && (true || (((((v3 | v3) >= v3) && false) && (false && true)) || (18446744073709551615 != 982)))) && true) {
                    v1 = v1
                } else {
                    v34 := v2[1]
                    v1 -= (v1 / (v1 + 1.5))
                    _ = v34
                }
                if (-51 <= 42) {
                    v35 := GenS1{F0: v3}
                    switch (fi22 % ((127) | 1)) {
                    case 77:
                        var v36 int32 = 2147483647
                        _ = v36
                    case -540:
                        v37 := []int{(-(fi22)), fi22, fi33}
                        _ = v37
                    case 42:
                        v1 += v1
                        var v38 bool = (!false)
                        _ = v38
                    }
                    v39 := v2[(-(fi33))]
                    _ = v35
                    _ = v39
                } else {
                    if (true || (127 <= -1)) {
                        var v40 int16 = 1
                        v40 += (^(v40))
                        var v41 uint64 = 4294967295
                        var v42 string = ("" + "hello")
                        v43 := GenS2{F0: 255}
                        _ = v40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
                switch 32768 {
                case 127:
                    var v44 int32 = 2147483647
                    _ = v44
                case 42:
                    v45 := GenS3{F0: 0}
                    _ = v45
                }
                _ = fi33
            }
        }
        var v46 string = (("foo bar" + "hello") + ((("hello" + "foo bar") + "a") + "foo bar"))
        v47 := []int{fi22, 2147483647}
        _ = fi22
        _ = v46
        _ = v47
    }
    v48 := GenS4{F0: "hello"}
    v49 := []int{}
    var v50 uint16 = 2
    v50 = 32767
    switch 0 {
    case 42:
        v1 += 0.0
    case 0:
        if ((v3 / ((v3) | 1)) != (v3 >> ((18446744073709551615) & 63))) {
            for fi51 := 0; fi51 < 13; fi51++ {
                for fi52 := 0; fi52 < 17; fi52++ {
                    v3 += v3
                    _ = len(v2)
                    v3 = v3
                    var v53 uint64 = 230
                    var v54 bool = false
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi51
            }
            switch 32768 {
            case -596:
                v49 = append(v49, 0)
            case -94:
                var v55 int32 = -963
                _ = v55
            default:
                v50 = 42
            }
            for fi56 := 0; fi56 < 9; fi56++ {
                v57 := GenS1{F0: v3}
                if len(v49) > 0 {
                    v49[0] = 42
                }
                _ = fi56
                _ = v57
            }
        } else {
            if (0 < 31) {
                var v58 int16 = -32768
                v59 := GenS5{F0: 1}
                _ = v58
                _ = v59
            }
            v3 -= (v3 | (v3 & v3))
        }
    default:
        for fi60 := 0; fi60 < 8; fi60++ {
            for fi61 := 0; fi61 < 4; fi61++ {
                fmt.Println(fi60)
                var v62 float32 = v1
                var v63 uint32 = 2147483647
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = fi60
        }
    }
    switch ("foo bar" + "") {
    case "":
        v50 -= (v50 | v50)
    case "a":
        delete(v2, 32767)
    case "foo bar":
        fmt.Println(v1)
    default:
        switch "hello" {
        case "foo bar":
            var v64 bool = false
            _ = v64
        case "hello":
            switch "\n\t" {
            case "":
                var v65 uint = v3
                _ = v65
            default:
                if false {
                    var v66 uint = 1
                    if false {
                        for fi67 := 0; fi67 < 3; fi67++ {
                            switch "a" {
                            case "foo bar":
                                var v68 float64 = (3.14159 + 1.5)
                                _ = v68
                            }
                            _ = fi67
                        }
                    } else {
                        v50 = v50
                        v3 = v66
                        v3 = 567
                    }
                    _ = v66
                }
            }
        default:
            v69 := GenS6{F0: genFn0(), F1: 890}
            _ = v69
        }
    }
    _ = len(v49)
    var v70 uint32 = 2
    fmt.Println(v50, v50, v1)
    v71 := GenS1{F0: 65535}
    var v72 uint = v3
    _ = len(v49)
    v49 = append(v49, -573)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v70
    _ = v71
    _ = v72
}
