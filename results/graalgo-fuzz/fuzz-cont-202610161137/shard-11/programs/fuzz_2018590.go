package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64) uint32 {
    v7 := []int{1, 204, 4294967295}
    p0 -= p0
    var v8 int32 = 373
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int16, p1 int16) int16 {
    var v9 uint32 = 466
    var v10 int16 = (^(p0))
    var v11 string = ""
    if true {
        v10 -= v10
        p1 = 403
        v10 += (p0 - 1)
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            switch (v11 + v11) {
            case "a":
                if (-131 > -56) {
                    if true {
                        v10 -= (v10 + p1)
                    } else {
                        v11 = ""
                        p0 += p1
                        v13 := []int{0, fi12, (^(fi12)), fi12}
                        _ = v13
                    }
                }
            }
            _ = fi12
        }
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint8 = 0
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int8 = int8(v15)
            var v18 bool = (3.14159 <= 0.1)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (v10 % ((v10) | 1))
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "\n\t", "", ""}
    var v2 float64 = 0.0
    var v3 float64 = 3.14159
    var v4 bool = ((!false) && false)
    var v5 int = 1
    v6 := []string{}
    var v19 string = "hello"
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v2, fi20, v4)
        var v21 uint8 = 190
        fmt.Println(v19, v4, v3)
        _ = fi20
        _ = v21
    }
    _ = len(v6)
    switch v5 {
    case -246:
        var v22 int64 = -581
        _ = v22
    case -2147483648:
        for fi23 := 0; fi23 < 11; fi23++ {
            switch ((v19 + "") + v19) {
            case "hello":
                var v24 string = ("" + (("hello" + v19) + ""))
                _ = v24
            case "\n\t":
                v5 += -197
            case "a":
                var v25 uint16 = 42
                _ = v25
            }
            var v26 bool = (-570 < -14)
            _ = fi23
            _ = v26
        }
    case 4294967295:
        var v27 uint = 18446744073709551615
        _ = v27
    default:
        var v28 float32 = 3.14159
        _ = v28
    }
    switch -1 {
    case 128:
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint32 = 466
                var v32 int8 = 11
                _ = len(v6)
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 int16 = 0
            if ((1 <= 619) || v4) {
                v33 -= v33
                v34 := GenS1{F0: 95}
                v35 := GenS1{F0: -66}
                _ = v34
                _ = v35
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi29
            _ = v33
        }
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 float64 = v2
            switch 630 {
            case 4294967295:
                switch v19 {
                case "":
                    var v38 uint8 = 42
                    _ = v38
                default:
                    v39 := GenS2{F0: 413, F1: v4, F2: 42}
                    _ = v39
                }
            case 42:
                for fi40 := 0; fi40 < 9; fi40++ {
                    _ = len(v1)
                    v5 += (^(fi36))
                    _ = fi40
                }
            }
            _ = fi36
            _ = v37
        }
    }
    var v41 uint = 387
    v42 := GenS3{F0: ("" + v19), F1: 32767, F2: 1}
    for fi43 := 0; fi43 < 2; fi43++ {
        if v4 {
            if true {
                if (96 <= 0) {
                    v5 -= v5
                }
            }
        }
        _ = fi43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        if v4 {
            for fi45 := 0; fi45 < 19; fi45++ {
                v4 = (!v4)
                fmt.Println(v4, fi44, fi45)
                v46 := GenS1{F0: -128}
                _ = fi45
                _ = v46
            }
            v47 := GenS3{F0: v19, F1: -32768, F2: 4294967295}
            _ = v47
        } else {
            var v48 int16 = -32768
            _ = v48
        }
        for fi49 := 0; fi49 < 3; fi49++ {
            v6 = append(v6, v19)
            v50 := GenS1{F0: 100}
            v2 += ((float64((-2.25 / 1e10)) - 0.0) / 0.1)
            _ = fi49
            _ = v50
        }
        if v4 {
            var v51 uint64 = 2
            if ((0 != 42) || v4) {
                for fi52 := 0; fi52 < 19; fi52++ {
                    v19 = "\n\t"
                    fmt.Println(v2, fi44)
                    v19 = v19
                    _ = fi52
                }
                fmt.Println(v3, v3)
                if len(v1) > 0 {
                    v1[0] = v19
                }
            } else {
                fmt.Println(fi44, fi44, v51)
                v19 = "a"
            }
            v53 := GenS1{F0: -128}
            _ = v51
            _ = v53
        }
        _ = fi44
    }
    var v54 uint64 = 958
    var v55 float32 = 0.1
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            for fi56 := 0; fi56 < 13; fi56++ {
                v57 := GenS4{F0: 1, F1: 1}
                var v58 int64 = 499
                fmt.Println(v55, v19)
                _ = fi56
                _ = v57
                _ = v58
            }
            var v59 int16 = 42
            _ = v59
        }
        for fi60 := 0; fi60 < 17; fi60++ {
            switch "a" {
            case "hello":
                for fi61 := 0; fi61 < 14; fi61++ {
                    v62 := GenS1{F0: 42}
                    v63 := GenS2{F0: 2, F1: v4, F2: 2}
                    v64 := GenS2{F0: 65535, F1: (!(!((0.1 > (v55 + 0.1)) && v4))), F2: uint8(v54)}
                    v54 += (v54 % ((v54) | 1))
                    _ = fi61
                    _ = v62
                    _ = v63
                    _ = v64
                }
            }
            if ((!(4294967295 >= 1)) || (42 > v54)) {
                if v4 {
                    for fi65 := 0; fi65 < 18; fi65++ {
                        v6 = append(v6, v19)
                        switch (-(fi60)) {
                        case 128:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            _ = len(v6)
                        }
                        _ = fi65
                    }
                }
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
            }
            v66 := GenS3{F0: v19, F1: genFn1(42, genFn1(634, 1)), F2: 827}
            _ = fi60
            _ = v66
        }
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        fmt.Println(v5, v4, v2)
        v54 += 2
        _ = fi67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
