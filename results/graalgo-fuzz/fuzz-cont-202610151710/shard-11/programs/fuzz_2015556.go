package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint; F1 int8; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int64; F1 int; F2 bool }
type GenS8 struct { F0 string; F1 int32; F2 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 uint) uint {
    var v6 uint32 = 972
    switch "foo bar" {
    case "":
        v6 -= v6
    case "\n\t":
        var v7 int16 = -32768
        _ = v7
    case "a":
        v8 := GenS1{F0: ((p0 | 256) < (^(p0))), F1: 9223372036854775807, F2: 0}
        _ = v8
    default:
        if (42 >= 717) {
            v9 := []int{32768, -22, -9223372036854775808}
            var v10 int64 = 42
            _ = v10
            _ = v9
        } else {
            var v11 uint8 = 2
            p0 = 42
            _ = v11
        }
    }
    fmt.Println(v6)
    switch 0 {
    case -2147483648:
        var v12 uint = 65535
        _ = v12
    }
    p0 -= 42
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = (1.5 - 3.14159)
    v3 := make(map[int]int)
    var v4 int = 2147483647
    v5 := make(map[string]bool)
    if (0 < 127) {
        var v13 bool = true
        var v14 uint = 4294967295
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v2)
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v2, v4)
                _ = fi16
            }
            _ = fi15
        }
        _ = v13
        _ = v14
    }
    v17 := GenS1{F0: (!false), F1: (^(v4)), F2: 818}
    var v18 int8 = 50
    v19 := v3[(v4 << ((1) & 63))]
    fmt.Println(v2, v18, v19)
    if (true && false) {
        fmt.Println(v4)
        v18 += v18
        v20 := GenS2{F0: 32767, F1: 0}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 string = (("hello" + (("a" + "a") + "")) + "")
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS2{F0: 1, F1: 1}
    if true {
        switch "foo bar" {
        case "a":
            v24 := GenS3{F0: v19, F1: (v19 + v4), F2: 836}
            _ = v24
        case "foo bar":
            v18 = -128
        default:
            if (2147483647 <= -2147483648) {
                for fi25 := 0; fi25 < 9; fi25++ {
                    var v26 bool = (false && (2147483648 != 0))
                    _ = fi25
                    _ = v26
                }
                v18 = (v18 * int8(v19))
            }
        }
    } else {
        var v27 int64 = 9223372036854775807
        v18 -= v18
        _ = v27
    }
    v28 := GenS4{F0: "foo bar", F1: 1}
    var v29 float32 = (1.5 * float32((1.5 + 0.1)))
    v30 := GenS5{F0: 32, F1: v18, F2: 9223372036854775807}
    if false {
        if true {
            v31 := GenS4{F0: (("" + "\n\t") + ""), F1: 18446744073709551615}
            switch (-(v4)) {
            case -9223372036854775808:
                v4 = v4
            case 89:
                for fi32 := 0; fi32 < 4; fi32++ {
                    var v33 int64 = 42
                    v34 := []int{fi32, 1, -852, 1}
                    switch "a" {
                    case "\n\t":
                        v18 = -128
                    case "":
                        fmt.Println(v4, v2)
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            case 123:
                v18 -= -128
            }
            _ = v31
        } else {
            var v35 int8 = v18
            var v36 int8 = v35
            if (!false) {
                v19 = (v4 - (v4 << ((2) & 63)))
                if false {
                    v2 += float32(0.1)
                    var v37 bool = true
                    _ = v37
                } else {
                    if ((18446744073709551615 == 446) && false) {
                        if true {
                            v19 = v4
                        }
                        for fi38 := 0; fi38 < 3; fi38++ {
                            v39 := GenS2{F0: 65535, F1: 18446744073709551615}
                            if true {
                                for fi40 := 0; fi40 < 4; fi40++ {
                                    fmt.Println(v29)
                                    switch (v4 - fi40) {
                                    case -929:
                                        switch 9223372036854775807 {
                                        case 2147483647:
                                            v41 := GenS6{F0: 141, F1: 18446744073709551615, F2: 357}
                                            _ = v41
                                        }
                                    default:
                                        fmt.Println(v36, v4)
                                    }
                                    _ = fi40
                                }
                                v29 -= v2
                            } else {
                                v42 := v1["foo bar"]
                                v42 = v42
                                if (465 > 489) {
                                    if true {
                                        var v43 string = v42
                                        var v44 int8 = (^(v35))
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                _ = v42
                            }
                            v45 := []int{(v19 & fi38), (v4 / ((fi38) | 1))}
                            _ = fi38
                            _ = v39
                            _ = v45
                        }
                    } else {
                        v46 := GenS4{F0: "hello", F1: 4294967295}
                        _ = v46
                    }
                }
            } else {
                v47 := []int{-372, (v4 >> ((18446744073709551615) & 63))}
                for fi48 := 0; fi48 < 9; fi48++ {
                    var v49 int64 = 9223372036854775807
                    var v50 uint = 428
                    v4 = v4
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = v47
            }
            _ = v35
            _ = v36
        }
        fmt.Println(v19, v4)
        if false {
            switch "foo bar" {
            case "":
                v51 := GenS7{F0: 0, F1: (v4 << ((genFn0(0)) & 63)), F2: (true || (!(false || false)))}
                _ = v51
            }
            fmt.Println(v29)
            fmt.Println(v29, v19)
        } else {
            v19 += -599
        }
    } else {
        v29 = float32(-2.25)
        for fi52 := 0; fi52 < 10; fi52++ {
            v53 := GenS5{F0: 1, F1: -59, F2: 9223372036854775807}
            if (2 != 65535) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v54 := GenS8{F0: "hello", F1: 1, F2: 0}
                    _ = v54
                default:
                    switch "" {
                    case "":
                        v55 := GenS9{F0: 2}
                        _ = v55
                    case "hello":
                        v2 -= v2
                    case "\n\t":
                        v56 := []int{}
                        _ = v56
                    }
                }
                v57 := v1["\n\t"]
                _ = v57
            } else {
                v58 := GenS3{F0: fi52, F1: v4, F2: 582}
                v59 := GenS6{F0: 211, F1: 834, F2: 0}
                if true {
                    v2 -= (0.0 - (1.5 + v2))
                } else {
                    v29 = (0.0 / v29)
                }
                _ = v58
                _ = v59
            }
            if false {
                delete(v5, "hello")
                fmt.Println(v19, v4)
                for fi60 := 0; fi60 < 1; fi60++ {
                    var v61 float64 = (0.0 / -2.25)
                    if (true && true) {
                        v62 := GenS10{F0: 548, F1: 65535, F2: 1}
                        v61 += 0.1
                        if true {
                            v63 := []int{fi60, (fi60 / (((v19 | v4)) | 1)), v19, (fi60 % ((v4) | 1))}
                            _ = v63
                        }
                        _ = v62
                    }
                    var v64 int16 = 1
                    _ = fi60
                    _ = v61
                    _ = v64
                }
            }
            _ = fi52
            _ = v53
        }
    }
    switch (("\n\t" + ("hello" + (("foo bar" + ("a" + ("a" + ""))) + "foo bar"))) + "foo bar") {
    case "a":
        var v65 int16 = 418
        _ = v65
    case "foo bar":
        if true {
            if false {
                var v66 bool = false
                _ = v66
            }
            var v67 uint32 = 2147483648
            _ = v67
        }
    case "hello":
        for fi68 := 0; fi68 < 17; fi68++ {
            for fi69 := 0; fi69 < 8; fi69++ {
                var v70 int16 = -1
                if (317 >= -1) {
                    v19 = 377
                    _ = len(v1)
                }
                v71 := []int{(v19 % (((fi69 & v19)) | 1)), (fi68 << ((65535) & 63))}
                _ = fi69
                _ = v70
                _ = v71
            }
            v72 := []int{522, (fi68 + -758)}
            v29 = v2
            _ = fi68
            _ = v72
        }
    default:
        v4 = v4
    }
    var v73 string = "\n\t"
    v29 -= (float32(((1e10 / (-2.25 / ((((0.0 + float64(v29)) - (1e10 * 3.14159)) * 0.0) * 3.14159))) + 1e10)) + float32(0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v73:", v73)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v73
}
