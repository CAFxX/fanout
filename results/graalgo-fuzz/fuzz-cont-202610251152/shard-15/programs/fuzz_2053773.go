package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int; F1 uint64 }
func genFn0(p0 int8, p1 int32) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            var v6 uint = 4294967295
            v6 += v6
            _ = v6
        }
        if (32767 > 463) {
            switch 32767 {
            case 1:
                fmt.Println(p0)
            case 259:
                var v7 float32 = 3.14159
                _ = v7
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                fmt.Println(fi8, p1)
                var v9 uint32 = 0
                var v10 int16 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
            p1 = p1
        } else {
            if (235 >= 128) {
                switch "hello" {
                case "\n\t":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := []int{fi11, 128}
                        v13 := GenS1{F0: 625, F1: 1, F2: 32767}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                if true {
                    p0 += -7
                }
                p0 = p0
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := []int{(fi14 ^ (fi14 + fi5)), fi14, (fi14 | fi14), fi14}
                p0 += 127
                p0 = (p0 ^ -70)
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := []int{fi16, fi5}
                _ = fi16
                _ = v17
            }
        }
        _ = fi5
    }
    switch 4294967295 {
    case -698:
        var v18 uint16 = 609
        _ = v18
    case 4294967295:
        p1 = -1
    case 42:
        v19 := GenS1{F0: 42, F1: -9223372036854775808, F2: 649}
        _ = v19
    default:
        if (p0 != p0) {
            var v20 uint8 = 42
            _ = v20
        }
    }
    v21 := GenS2{F0: 0, F1: (!(596 >= 1))}
    _ = p0
    _ = p1
    _ = v21
    return 0
}
func genFn1(p0 int) string {
    p0 -= 1
    var v22 uint32 = 2147483647
    _ = p0
    _ = v22
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 56
    v3 := make(map[int]uint32)
    var v4 int32 = -2147483648
    fmt.Println(v4, v2)
    v23 := []int{9223372036854775807, 42}
    fmt.Println(v4, v2)
    switch 42 {
    case 42:
        v2 = v2
    case 128:
        _ = len(v3)
    case 9223372036854775807:
        v2 = v2
    default:
        fmt.Println(v4, v2)
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v2 = (-(v2))
        if (!(!false)) {
            v25 := GenS1{F0: 18446744073709551615, F1: 1, F2: 65535}
            if (-2.25 != 1e10) {
                v26 := v1[2]
                var v27 string = v26
                _ = v26
                _ = v27
            } else {
                v4 = (v4 * v4)
                v28 := GenS1{F0: 65535, F1: 42, F2: 32767}
                v29 := GenS1{F0: 582, F1: 1, F2: 1}
                _ = v28
                _ = v29
            }
            var v30 uint32 = 65535
            _ = v25
            _ = v30
        } else {
            if true {
                var v31 uint = 4294967295
                _ = v31
            }
            fmt.Println(v4, v4, fi24)
            switch 127 {
            case -435:
                if ((!(!(genFn0(v2, v4) > 1))) || (2 == 787)) {
                    delete(v1, 493)
                    _ = len(v3)
                } else {
                    v23 = append(v23, fi24)
                }
            case 2147483647:
                fmt.Println(v4, v4, v4)
            case 32767:
                switch ("foo bar" + "\n\t") {
                case "a":
                    switch fi24 {
                    case -176:
                        var v32 float32 = 0.1
                        _ = v32
                    case -357:
                        var v33 int32 = 42
                        _ = v33
                    case 127:
                        delete(v3, 32768)
                    default:
                        fmt.Println(fi24)
                    }
                case "":
                    v4 = (^(v4))
                case "\n\t":
                    if (((-1 < 1) || false) || ((("hello" + "hello") + "hello") < "foo bar")) {
                        if true {
                            v34 := GenS3{F0: 319, F1: 281}
                            v35 := GenS3{F0: 18446744073709551615, F1: 0}
                            v2 -= 42
                            _ = v34
                            _ = v35
                        }
                        switch fi24 {
                        case 138:
                            if false {
                                var v36 int8 = v2
                                v2 -= (v36 + v2)
                                v37 := GenS2{F0: 668, F1: true}
                                _ = v36
                                _ = v37
                            }
                        default:
                            var v38 string = "a"
                            _ = v38
                        }
                    } else {
                        for fi39 := 0; fi39 < 2; fi39++ {
                            v40 := GenS3{F0: 826, F1: 674}
                            v41 := GenS4{F0: false, F1: 1}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    }
                default:
                    v4 += int32(v2)
                }
            default:
                v1[307] = "\n\t"
            }
        }
        var v42 bool = (((0.1 / ((float32((1.5 / 1.5)) + 0.0) * (1e10 - 0.1))) != 0.0) || false)
        _ = fi24
        _ = v42
    }
    if ((9223372036854775807 < 1) && (42 == 0)) {
        v43 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 42}
        v4 += (-(v4))
        _ = v43
    } else {
        switch 32767 {
        case 4294967295:
            fmt.Println(v2, v2)
        }
        switch "a" {
        case "a":
            v44 := GenS1{F0: 223, F1: 42, F2: 474}
            _ = v44
        case "":
            if len(v23) > 0 {
                v23[0] = -378
            }
        }
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        if true {
            switch "hello" {
            case "\n\t":
                v4 += v4
            }
            var v46 bool = (0 <= 0)
            v47 := GenS3{F0: 0, F1: 255}
            _ = v46
            _ = v47
        } else {
            fmt.Println(fi45)
            v4 -= v4
        }
        v48 := GenS5{F0: (!true), F1: (fi45 * fi45)}
        _ = fi45
        _ = v48
    }
    v2 = (v2 - (v2 | v2))
    v49 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 902}
    v50 := GenS6{F0: 32767, F1: 4294967295}
    var v51 float32 = 3.14159
    var v52 string = "a"
    v3[-9223372036854775808] = 1
    fmt.Println(v51)
    for fi53 := 0; fi53 < 7; fi53++ {
        v4 += v4
        var v54 float64 = -2.25
        _ = fi53
        _ = v54
    }
    var v55 int32 = (-(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v55
}
