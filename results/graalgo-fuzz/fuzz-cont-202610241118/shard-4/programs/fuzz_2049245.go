package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 uint64, p1 uint) int {
    p1 = 42
    var v6 uint64 = p0
    if true {
        v7 := []int{128, -458}
        p0 += p0
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 string = ""
    var v2 uint64 = 691
    var v3 int32 = 802
    var v4 int64 = 9223372036854775807
    var v5 uint = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            var v9 float32 = -2.25
            switch (fi8 * (fi8 << ((v5) & 63))) {
            case 4294967295:
                v3 -= (-(v3))
            case -489:
                v10 := GenS1{F0: (fi8 >> (((v5 | v5)) & 63)), F1: 0, F2: -518}
                _ = v10
            default:
                v4 += 1
            }
            var v11 uint16 = 255
            _ = v11
            _ = v9
        }
        switch fi8 {
        case 0:
            var v12 uint16 = 65535
            _ = v12
        case 587:
            if ((((!(-899 != (fi8 + fi8))) && (!true)) || true) || false) {
                switch "foo bar" {
                case "\n\t":
                    if (((v1 + ("" + (v1 + v1))) + "\n\t") == v1) {
                        switch v1 {
                        case "a":
                            v1 = "hello"
                        case "hello":
                            if (!true) {
                                v1 = "a"
                            }
                        case "foo bar":
                            if ((-35 <= 1) && false) {
                                var v13 bool = false
                                v5 += v5
                                fmt.Println(v4)
                                _ = v13
                            } else {
                                var v14 float64 = 0.0
                                v15 := GenS1{F0: fi8, F1: v3, F2: -487}
                                var v16 float32 = (float32(1.5) + 0.1)
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        default:
                            if false {
                                v17 := GenS1{F0: fi8, F1: 0, F2: 1}
                                fmt.Println(v3, v5, fi8)
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    v19 := GenS2{F0: v4, F1: -1, F2: 1000}
                                    v20 := []int{2147483647, (fi18 << ((v5) & 63)), genFn0(v2, v5)}
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                                _ = v17
                            } else {
                                var v21 int64 = (-(v4))
                                _ = v21
                            }
                        }
                    } else {
                        var v22 float32 = -2.25
                        _ = v22
                    }
                }
                if true {
                    v23 := []int{int(v5), genFn0((v2 & (v2 * v2)), (v5 & 42))}
                    _ = v23
                }
                if true {
                    if false {
                        switch fi8 {
                        case -1:
                            v24 := []int{(fi8 / (((fi8 - (fi8 / (((-(fi8))) | 1)))) | 1)), (fi8 >> (((v5 >> (((v5 << ((1) & 63))) & 63))) & 63))}
                            _ = v24
                        }
                        var v25 float32 = 1.5
                        var v26 int32 = int32(fi8)
                        _ = v25
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 15; fi27++ {
                            if false {
                                var v28 uint64 = v2
                                var v29 int64 = (v4 ^ v4)
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS2{F0: (v4 | v4), F1: -1, F2: 107}
                            fmt.Println(v4)
                            _ = fi27
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 8; fi31++ {
                            switch 0 {
                            case 32767:
                                v32 := []int{(fi8 % ((fi8) | 1))}
                                _ = v32
                            default:
                                switch int(v3) {
                                case 9223372036854775807:
                                    v4 -= v4
                                    v33 := []int{}
                                    _ = v33
                                default:
                                    fmt.Println(v2, v3, v3)
                                }
                            }
                            var v34 uint = 136
                            var v35 int16 = -995
                            _ = fi31
                            _ = v34
                            _ = v35
                        }
                        var v36 uint32 = 65535
                        _ = v36
                    }
                    v37 := GenS3{F0: v1, F1: -49}
                    _ = v37
                }
            }
        case 32767:
            switch (v1 + "a") {
            case "foo bar":
                fmt.Println(v3, v3)
            case "\n\t":
                v4 = (v4 / ((-9223372036854775808) | 1))
            case "":
                v1 = v1
            default:
                v38 := GenS4{F0: "", F1: 42}
                _ = v38
            }
        }
        _ = fi8
    }
    var v39 float64 = (((1e10 / 0.1) * float64(-2.25)) - 3.14159)
    for fi40 := 0; fi40 < 6; fi40++ {
        if (0.1 > -2.25) {
            if false {
                for fi41 := 0; fi41 < 1; fi41++ {
                    if (v39 == (v39 + (1.5 * float64(3.14159)))) {
                        v4 = (v4 % (((^(v4))) | 1))
                    }
                    _ = fi41
                }
                v42 := []int{}
                _ = v42
            }
        }
        _ = fi40
    }
    var v43 uint8 = 114
    var v44 uint32 = 65535
    for fi45 := 0; fi45 < 4; fi45++ {
        if (!(false || (!false))) {
            v46 := []int{(fi45 << ((uint(v43)) & 63)), genFn0(v2, uint(v44)), (fi45 * (fi45 * (-(fi45))))}
            _ = v46
        } else {
            if true {
                v3 = v3
            } else {
                v39 -= float64(0.0)
                for fi47 := 0; fi47 < 17; fi47++ {
                    v3 += v3
                    v48 := []int{int(v5)}
                    fmt.Println(v43)
                    _ = fi47
                    _ = v48
                }
            }
        }
        _ = fi45
    }
    v39 += (v39 / 1.5)
    for fi49 := 0; fi49 < 13; fi49++ {
        for fi50 := 0; fi50 < 19; fi50++ {
            v4 = 0
            _ = fi50
        }
        if true {
            if (!(false && false)) {
                v39 += (v39 + v39)
                v44 = (v44 % (((v44 ^ 4294967295)) | 1))
            } else {
                v39 += v39
                fmt.Println(fi49, v4)
            }
            v51 := []int{fi49, (fi49 | genFn0((v2 * (v2 ^ v2)), (v5 / ((v5) | 1)))), fi49, (fi49 * fi49)}
            var v52 int8 = -1
            _ = v51
            _ = v52
        }
        for fi53 := 0; fi53 < 17; fi53++ {
            var v54 int16 = 0
            var v55 uint32 = v44
            fmt.Println(v4, v2, v39)
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi49
    }
    for fi56 := 0; fi56 < 16; fi56++ {
        if true {
            v1 = ("\n\t" + ((v1 + v1) + v1))
        }
        if (!(!true)) {
            switch ("hello" + v1) {
            case "\n\t":
                v57 := GenS3{F0: v1, F1: 100}
                _ = v57
            case "":
                v44 -= v44
            }
            switch v1 {
            case "a":
                var v58 bool = false
                _ = v58
            case "hello":
                v59 := []int{fi56, genFn0(726, v5), fi56, fi56}
                _ = v59
            }
        }
        _ = fi56
    }
    v60 := GenS5{F0: 2, F1: 128, F2: -32768}
    v61 := GenS5{F0: (v5 | v5), F1: 1, F2: 0}
    v62 := []int{0, 592, 1}
    if (!((954 == 667) && false)) {
        if (false && true) {
            fmt.Println(v4, v2)
            if true {
                v3 = (^(v3))
                for fi63 := 0; fi63 < 13; fi63++ {
                    v44 -= uint32(v2)
                    v5 -= 225
                    _ = fi63
                }
                if len(v62) > 0 {
                    v62[0] = 42
                }
            }
        } else {
            var v64 int64 = 1
            if (!false) {
                for fi65 := 0; fi65 < 4; fi65++ {
                    if len(v62) >= 2 {
                        v62 = v62[1:len(v62)]
                    }
                    fmt.Println(v5, fi65)
                    v4 += 51
                    _ = fi65
                }
                var v66 float64 = ((((((1e10 * (v39 * v39)) - v39) * (v39 + v39)) + float64(1.5)) * v39) - v39)
                _ = v66
            }
            _ = v64
        }
        var v67 float64 = v39
        _ = v67
    }
    v68 := GenS3{F0: "\n\t", F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v68
}
