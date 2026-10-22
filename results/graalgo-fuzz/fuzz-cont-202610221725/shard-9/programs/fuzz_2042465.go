package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 string; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 int8) int64 {
    p0 = (p0 >> ((4294967295) & 63))
    var v4 bool = true
    if (!false) {
        var v5 uint64 = 42
        switch 4294967295 {
        case 32767:
            var v6 string = "a"
            _ = v6
        }
        var v7 int8 = p1
        _ = v5
        _ = v7
    }
    if true {
        var v8 uint16 = 2
        v9 := GenS1{F0: (-2147483648 < 1), F1: -2147483648}
        _ = v8
        _ = v9
    } else {
        switch "hello" {
        case "foo bar":
            var v10 int = 185
            _ = v10
        }
        v11 := GenS1{F0: (!v4), F1: 743}
        _ = v11
    }
    v4 = ((1 != 255) && ((75 != 2) || (!v4)))
    _ = p0
    _ = p1
    _ = v4
    return 9223372036854775807
}
func genFn1(p0 int8, p1 uint64) int32 {
    v12 := []int{4294967295, 42, -1, int(p1)}
    var v13 uint = 902
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 557
    var v3 int16 = -32768
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int16 = v3
        v16 := GenS1{F0: (-128 == 127), F1: 0}
        for fi17 := 0; fi17 < 1; fi17++ {
            switch "a" {
            case "":
                v18 := []int{(fi17 ^ 32768), fi17, fi14, (fi14 + (fi17 - (fi14 + (fi14 * fi14))))}
                var v19 uint16 = 255
                v2 = (v2 & 209)
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    switch (("foo bar" + "foo bar") + "") {
    case "hello":
        switch 127 {
        case 1:
            for fi20 := 0; fi20 < 13; fi20++ {
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint16 = 551
                    delete(v1, (v2 * v2))
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        case 0:
            if true {
                if true {
                    v3 -= (v3 + v3)
                    var v23 int8 = -128
                    v24 := GenS2{F0: 2147483647, F1: 1}
                    _ = v23
                    _ = v24
                } else {
                    var v25 int32 = -1
                    _ = v25
                }
                switch "" {
                case "\n\t":
                    if (298 == 569) {
                        v26 := GenS3{F0: 2}
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v1[v2] = "foo bar"
                            v3 += 0
                            _ = fi27
                        }
                        _ = v26
                    } else {
                        v28 := GenS4{F0: "hello", F1: "a", F2: 42}
                        _ = v28
                    }
                }
                v3 -= (v3 % ((v3) | 1))
            } else {
                delete(v1, (v2 / ((2) | 1)))
                if false {
                    if ((0.1 * ((float32(float64(float32((0.1 + -2.25)))) * 0.1) * (1e10 / 3.14159))) < (0.1 * 1e10)) {
                        switch -1 {
                        case 291:
                            for fi29 := 0; fi29 < 16; fi29++ {
                                if (-128 > int8(v3)) {
                                    fmt.Println(v2, v3)
                                } else {
                                    var v30 int16 = v3
                                    var v31 bool = (false && (!false))
                                    _ = v30
                                    _ = v31
                                }
                                switch (fi29 ^ (fi29 | (fi29 % ((int(v2)) | 1)))) {
                                case 127:
                                    var v32 bool = ((((1.5 / 1e10) - (((-2.25 / 1.0) * float32(((1.5 + (1e10 - (1.5 + float64(1e10)))) * ((3.14159 / (1.5 - 1e10)) - -2.25)))) * 1e10)) - 0.0) <= float32(1e10))
                                    _ = v32
                                }
                                _ = fi29
                            }
                        case 32767:
                            switch 127 {
                            case 0:
                                v33 := GenS2{F0: -33, F1: 1}
                                _ = v33
                            }
                        default:
                            v3 = v3
                        }
                    }
                } else {
                    if false {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v35 := GenS5{F0: v3}
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            }
        default:
            fmt.Println(v3, v2)
        }
    case "\n\t":
        fmt.Println(v3, v2)
    }
    v3 -= (v3 & (v3 - v3))
    v3 = v3
    if ((v2 >= 974) && ("foo bar" == "a")) {
        var v36 uint = 65535
        if true {
            if (!(18446744073709551615 < 1)) {
                v37 := []int{564, -2147483648}
                _ = v37
            }
            fmt.Println(v2, v2, v36)
        } else {
            fmt.Println(v2, v2, v2)
            v3 = (-(v3))
        }
        v3 -= v3
        _ = v36
    } else {
        if ((1.5 == 0.1) || (!true)) {
            for fi38 := 0; fi38 < 10; fi38++ {
                v3 += v3
                _ = fi38
            }
            switch "\n\t" {
            case "hello":
                for fi39 := 0; fi39 < 10; fi39++ {
                    fmt.Println(fi39, fi39)
                    for fi40 := 0; fi40 < 1; fi40++ {
                        v2 = v2
                        _ = fi40
                    }
                    _ = fi39
                }
            case "a":
                var v41 int16 = (v3 & (v3 | 922))
                _ = v41
            }
            v3 = v3
        } else {
            v2 -= v2
            v42 := GenS2{F0: 1, F1: 2}
            v3 = (v3 >> ((42) & 63))
            _ = v42
        }
    }
    if (false && false) {
        switch "foo bar" {
        case "foo bar":
            var v43 int16 = v3
            _ = v43
        }
        v44 := []int{32767, 32767, 42, 2147483647}
        for fi45 := 0; fi45 < 20; fi45++ {
            var v46 int8 = 0
            v47 := GenS1{F0: true, F1: 42}
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = v44
    }
    var v48 uint32 = 961
    v49 := GenS6{F0: (v3 / ((v3) | 1)), F1: 128}
    v2 = 0
    v50 := GenS3{F0: 65535}
    var v51 bool = false
    var v52 int16 = (-(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
