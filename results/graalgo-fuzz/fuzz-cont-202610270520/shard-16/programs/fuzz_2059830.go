package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 bool; F2 bool }
func genFn0(p0 uint) uint {
    v4 := GenS1{F0: 126, F1: 0, F2: (p0 << ((p0) & 63))}
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: 127, F1: -621, F2: (p0 + (p0 + p0))}
        _ = fi5
        _ = v6
    }
    p0 -= (p0 / ((p0) | 1))
    if false {
        p0 += p0
        if (0 >= 940) {
            if (32767 >= 32767) {
                fmt.Println(p0)
                switch -2147483648 {
                case -69:
                    for fi7 := 0; fi7 < 6; fi7++ {
                        var v8 uint = 662
                        var v9 float32 = 1e10
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                case 264:
                    p0 += 1
                default:
                    var v10 bool = false
                    _ = v10
                }
                p0 -= p0
            } else {
                if true {
                    if (624 >= 42) {
                        v11 := []int{2147483647, -682, -2147483648, 196}
                        p0 = p0
                        if true {
                            fmt.Println(p0)
                            p0 -= (p0 % ((p0) | 1))
                        }
                        _ = v11
                    }
                    v12 := GenS2{F0: 32767, F1: 128}
                    _ = v12
                } else {
                    v13 := []int{-9223372036854775808, 2147483647, -9223372036854775808, -415}
                    if false {
                        fmt.Println(p0)
                        fmt.Println(p0)
                    } else {
                        if (true || true) {
                            for fi14 := 0; fi14 < 20; fi14++ {
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    p0 -= (^(p0))
                                    _ = fi15
                                }
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    var v17 float32 = 0.0
                                    _ = fi16
                                    _ = v17
                                }
                                var v18 uint64 = 2
                                _ = fi14
                                _ = v18
                            }
                        } else {
                            p0 = p0
                            var v19 uint32 = 65535
                            _ = v19
                        }
                    }
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 uint64 = 4294967295
                        v22 := GenS3{F0: -1, F1: 861, F2: 32767}
                        var v23 string = "hello"
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v13
                }
            }
            var v24 int32 = 45
            _ = v24
        }
        switch "" {
        case "foo bar":
            v25 := GenS2{F0: 405, F1: 128}
            _ = v25
        case "\n\t":
            p0 += p0
        default:
            v26 := GenS4{F0: 65535, F1: 32767}
            _ = v26
        }
    }
    p0 += p0
    _ = p0
    _ = v4
    return p0
}
func genFn1(p0 int32) uint16 {
    for fi27 := 0; fi27 < 8; fi27++ {
        if false {
            v28 := []int{-182, (fi27 * (fi27 + fi27)), fi27, fi27}
            var v29 uint64 = 42
            if false {
                for fi30 := 0; fi30 < 13; fi30++ {
                    var v31 float64 = (0.1 * -2.25)
                    fmt.Println(v29, fi30)
                    _ = fi30
                    _ = v31
                }
                p0 += (p0 & p0)
            } else {
                switch "" {
                case "a":
                    v28 = append(v28, (fi27 - -9223372036854775808))
                case "\n\t":
                    switch "" {
                    case "a":
                        var v32 uint32 = 1
                        _ = v32
                    case "hello":
                        v33 := GenS2{F0: 32767, F1: 159}
                        _ = v33
                    case "foo bar":
                        if (4294967295 != 0) {
                            p0 = 492
                            switch (-(fi27)) {
                            case 0:
                                fmt.Println(p0, v29, p0)
                            case 9223372036854775807:
                                for fi34 := 0; fi34 < 4; fi34++ {
                                    v35 := GenS3{F0: 1, F1: 2147483647, F2: 761}
                                    v29 = (v29 ^ v29)
                                    _ = fi34
                                    _ = v35
                                }
                            case 4294967295:
                                v29 += (v29 % (((v29 - 326)) | 1))
                            }
                        }
                    }
                case "":
                    var v36 float64 = 0.1
                    _ = v36
                }
                var v37 uint = 2
                _ = v37
            }
            _ = v28
            _ = v29
        }
        p0 -= 42
        var v38 float64 = float64(1.5)
        _ = fi27
        _ = v38
    }
    v39 := GenS5{F0: -330, F1: -94, F2: 1}
    _ = p0
    _ = v39
    return 2
}

func main() {
    var v1 int64 = 0
    var v2 float32 = 1.5
    var v3 uint = 65535
    if true {
        v3 = 572
        fmt.Println(v1, v1)
    } else {
        var v40 int16 = -254
        v41 := []int{814}
        var v42 int16 = (v40 << ((v3) & 63))
        _ = v40
        _ = v41
        _ = v42
    }
    v43 := GenS1{F0: 128, F1: 0, F2: (v3 + v3)}
    var v44 string = "foo bar"
    for fi45 := 0; fi45 < 5; fi45++ {
        if (4294967295 > 2147483647) {
            v1 = v1
            v3 += genFn0(4294967295)
        } else {
            switch "\n\t" {
            case "foo bar":
                v46 := []int{(fi45 % ((fi45) | 1)), 42}
                _ = v46
            case "hello":
                v47 := []int{}
                _ = v47
            default:
                v2 = v2
            }
            if (!(529 > 13)) {
                v48 := GenS5{F0: (v1 | (v1 << ((4294967295) & 63))), F1: -1, F2: (v3 * genFn0(v3))}
                if (1.5 >= (v2 - v2)) {
                    switch (v44 + ((v44 + (v44 + v44)) + ("foo bar" + ((((("" + v44) + v44) + ("a" + v44)) + v44) + "foo bar")))) {
                    case "a":
                        v2 += v2
                    case "":
                        var v49 bool = (v44 == v44)
                        _ = v49
                    case "\n\t":
                        v1 += 42
                    default:
                        fmt.Println(v2, v44)
                    }
                } else {
                    v2 -= (v2 * v2)
                    var v50 uint = genFn0((^(v3)))
                    if false {
                        if false {
                            var v51 int8 = -128
                            v1 += (v1 | -1)
                            _ = v51
                        }
                    } else {
                        v3 -= v50
                    }
                    _ = v50
                }
                _ = v48
            } else {
                var v52 int8 = 0
                v53 := GenS6{F0: (v3 & (v3 >> ((1) & 63))), F1: v52}
                var v54 uint16 = 2
                _ = v52
                _ = v53
                _ = v54
            }
        }
        v55 := GenS7{F0: 2}
        v3 += (^(v3))
        _ = fi45
        _ = v55
    }
    var v56 float32 = (float32(((-2.25 * 1.5) / 1.0)) / 1.0)
    if ((float64(v56) * 0.0) < float64((0.1 + 0.0))) {
        var v57 float64 = (1.5 + 1.5)
        _ = v57
    }
    v58 := GenS8{F0: -1, F1: false, F2: (!false)}
    for fi59 := 0; fi59 < 3; fi59++ {
        var v60 int64 = 9223372036854775807
        var v61 int32 = 2147483647
        var v62 uint8 = 0
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v56
    _ = v58
}
