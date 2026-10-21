package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint16; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64; F1 bool; F2 int }
func genFn0(p0 string) int64 {
    var v5 uint = 842
    if false {
        var v6 uint16 = 65535
        _ = v6
    } else {
        if (((false && (1.5 >= -2.25)) && true) && false) {
            v7 := []int{128, 283, 9223372036854775807}
            switch -2147483648 {
            case -552:
                var v8 float32 = ((1e10 - 0.1) * 1.5)
                _ = v8
            case 2147483647:
                var v9 float64 = 1.5
                _ = v9
            default:
                v10 := GenS1{F0: 0, F1: p0, F2: -2147483648}
                _ = v10
            }
            _ = v7
        } else {
            var v11 int32 = 42
            _ = v11
        }
        if ((true && (!true)) || (true && (v5 == (v5 * (v5 ^ (v5 & v5)))))) {
            v5 += v5
        } else {
            var v12 uint8 = 0
            for fi13 := 0; fi13 < 12; fi13++ {
                if true {
                    var v14 string = "\n\t"
                    _ = v14
                } else {
                    var v15 uint32 = 2147483648
                    v16 := []int{int(v12)}
                    _ = v15
                    _ = v16
                }
                v17 := GenS1{F0: 32767, F1: p0, F2: 2147483647}
                _ = fi13
                _ = v17
            }
            _ = v12
        }
        var v18 uint8 = 0
        _ = v18
    }
    p0 = (p0 + "hello")
    _ = p0
    _ = v5
    return 509
}

func main() {
    var v1 string = "a"
    var v2 int16 = 0
    var v3 uint8 = 1
    var v4 int64 = 1
    var v19 int = 9223372036854775807
    if ((true || (!(-1 != 100))) && (!true)) {
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := GenS2{F0: (v4 - v4)}
                v19 -= (fi20 << ((65535) & 63))
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    if (371 < 0) {
                        v3 = v3
                        var v24 int64 = v4
                        var v25 float32 = -2.25
                        _ = v24
                        _ = v25
                    } else {
                        if (true || false) {
                            var v26 uint64 = 1
                            v27 := []int{995}
                            fmt.Println(v26)
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS3{F0: 0}
                            _ = v28
                        }
                    }
                    var v29 int16 = (v2 >> ((0) & 63))
                    v1 = (v1 + "foo bar")
                    _ = v29
                } else {
                    if ((v1 <= "a") && (false && true)) {
                        switch fi20 {
                        case 475:
                            var v30 bool = false
                            _ = v30
                        case 9223372036854775807:
                            v31 := []int{(^(fi20)), 0, (-(v19))}
                            var v32 int8 = 123
                            switch v1 {
                            case "hello":
                                var v33 int16 = (v2 | v2)
                                _ = v33
                            default:
                                v34 := GenS1{F0: 18, F1: v1, F2: -262}
                                _ = v34
                            }
                            _ = v31
                            _ = v32
                        case 128:
                            v35 := GenS2{F0: v4}
                            _ = v35
                        }
                    }
                }
                _ = fi23
            }
            switch ("" + ("" + v1)) {
            case "foo bar":
                if false {
                    fmt.Println(v1)
                    var v36 uint32 = 1
                    v37 := GenS1{F0: 422, F1: v1, F2: -2147483648}
                    _ = v36
                    _ = v37
                }
            default:
                if false {
                    var v38 uint32 = 65535
                    v1 = v1
                    _ = v38
                }
            }
            _ = fi20
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            if (v19 < (fi39 >> ((4294967295) & 63))) {
                v40 := GenS4{F0: 1, F1: (v1 + (v1 + v1)), F2: 4294967295}
                _ = v40
            } else {
                v41 := GenS2{F0: (-(v4))}
                var v42 uint8 = (v3 % ((255) | 1))
                _ = v41
                _ = v42
            }
            _ = fi39
        }
        v19 -= v19
    } else {
        v43 := []int{(v19 * (v19 * (-(v19)))), 9223372036854775807, (v19 % (((v19 & 32768)) | 1))}
        _ = v43
    }
    v4 += -9223372036854775808
    v44 := GenS4{F0: 163, F1: v1, F2: 18446744073709551615}
    v45 := GenS5{F0: (-(v2))}
    v46 := GenS1{F0: v2, F1: "hello", F2: 42}
    v47 := GenS6{F0: "foo bar", F1: 818, F2: (v19 & 2147483647)}
    if (2 >= 724) {
        var v48 string = v1
        _ = v48
    }
    if (1.5 < 0.1) {
        v3 = v3
        switch v19 {
        case 128:
            switch v1 {
            case "":
                if false {
                    var v49 int64 = genFn0(v1)
                    _ = v49
                } else {
                    if true {
                        if true {
                            v2 -= 1000
                        }
                    }
                }
            case "a":
                v50 := []int{v19}
                _ = v50
            case "hello":
                var v51 int = v19
                _ = v51
            }
        case -1:
            if (!(255 != 650)) {
                v52 := []int{(v19 % ((-1) | 1))}
                v53 := GenS7{F0: (v4 % (((-(v4))) | 1))}
                switch v1 {
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    switch v19 {
                    case 358:
                        for fi54 := 0; fi54 < 9; fi54++ {
                            v55 := GenS1{F0: 716, F1: v1, F2: 2147483647}
                            var v56 bool = (-720 > -1)
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    case 2147483647:
                        var v57 string = ""
                        _ = v57
                    }
                default:
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                }
                _ = v52
                _ = v53
            } else {
                fmt.Println(v19, v2, v19)
                v1 = (v1 + ("foo bar" + v1))
            }
        case 4294967295:
            var v58 int64 = -9223372036854775808
            _ = v58
        }
        v19 = -750
    } else {
        v19 += v19
    }
    var v59 int = v19
    var v60 int8 = -1
    if false {
        v59 -= v19
    } else {
        fmt.Println(v2, v4, v1)
        fmt.Println(v4)
        if true {
            var v61 string = "\n\t"
            for fi62 := 0; fi62 < 6; fi62++ {
                var v63 bool = true
                v64 := GenS8{F0: (-(v4)), F1: v63, F2: 1}
                switch fi62 {
                case 329:
                    if ((v4 > (v4 * v4)) && v63) {
                        v59 = (^(v19))
                        for fi65 := 0; fi65 < 8; fi65++ {
                            var v66 uint32 = 65535
                            _ = fi65
                            _ = v66
                        }
                    }
                case 9223372036854775807:
                    var v67 int16 = (v2 >> ((0) & 63))
                    _ = v67
                }
                _ = fi62
                _ = v63
                _ = v64
            }
            _ = v61
        }
    }
    if (true || (!(1e10 >= -2.25))) {
        var v68 bool = true
        var v69 uint8 = (v3 - v3)
        _ = v68
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v59
    _ = v60
}
