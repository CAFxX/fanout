package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 int16, p1 int8) uint32 {
    p1 += (p1 >> ((379) & 63))
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int16 = -1
        var v6 int16 = (^(v5))
        v7 := GenS1{F0: 65535}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v8
    return 2
}

func main() {
    v1 := []int{0}
    v2 := []string{((("" + "hello") + ("foo bar" + "")) + "foo bar"), ("" + "a"), "hello", "\n\t"}
    v3 := make(map[int]bool)
    if (-1 <= 0) {
        var v9 int32 = 1
        switch 536 {
        case 128:
            var v10 int32 = v9
            _ = v10
        case 0:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 2147483647:
            v9 -= (-(v9))
        default:
            switch 378 {
            case -2147483648:
                for fi11 := 0; fi11 < 4; fi11++ {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        delete(v3, fi11)
                        _ = fi12
                    }
                    _ = fi11
                }
            case 590:
                var v13 uint16 = 258
                _ = v13
            case -9223372036854775808:
                for fi14 := 0; fi14 < 1; fi14++ {
                    if ((("a" + "a") + ("foo bar" + "\n\t")) == ("" + ("hello" + ((((("a" + ((("foo bar" + "a") + ("a" + "hello")) + "foo bar")) + "foo bar") + "foo bar") + "foo bar") + "")))) {
                        delete(v3, (fi14 * fi14))
                    } else {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS1{F0: 465}
                            if false {
                                var v17 uint16 = 220
                                _ = v17
                            } else {
                                v9 -= (v9 + (v9 & (v9 | (v9 % ((v9) | 1)))))
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = fi14
                }
            }
        }
        _ = v9
    } else {
        var v18 int32 = 42
        _ = v18
    }
    if (!(1e10 >= (0.0 + 0.1))) {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS1{F0: 0}
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS2{F0: -327, F1: fi19, F2: 0}
                var v23 float32 = 0.1
                if (42 == 1) {
                    var v24 uint64 = 2
                    _ = v24
                } else {
                    switch fi21 {
                    case 0:
                        v23 = v23
                    }
                    v25 := GenS3{F0: -195}
                    fmt.Println(fi21)
                    if (true || false) {
                        var v26 int64 = 995
                        _ = v26
                    }
                    _ = v25
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            var v27 string = "hello"
            _ = fi19
            _ = v20
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            if true {
                var v29 uint64 = 1
                v29 = (v29 << ((662) & 63))
                _ = v29
            }
            var v30 uint8 = 0
            v31 := GenS1{F0: 1}
            _ = fi28
            _ = v30
            _ = v31
        }
    } else {
        var v32 uint64 = 2
        if (127 <= 42) {
            for fi33 := 0; fi33 < 16; fi33++ {
                v3[-9223372036854775808] = (102 <= 127)
                v34 := GenS3{F0: 1000}
                _ = fi33
                _ = v34
            }
            v2 = append(v2, "foo bar")
        } else {
            v2 = append(v2, "foo bar")
            fmt.Println(v32)
            if false {
                fmt.Println(v32)
                v32 = (v32 + (v32 << ((4294967295) & 63)))
                var v35 int16 = 42
                _ = v35
            } else {
                var v36 uint8 = 240
                v37 := GenS4{F0: 2147483647, F1: 873}
                _ = v36
                _ = v37
            }
        }
        _ = v32
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 16; fi39++ {
            fmt.Println(fi38)
            if false {
                var v40 int64 = 0
                v41 := GenS1{F0: 2}
                v40 = 930
                v42 := GenS1{F0: 2}
                _ = v40
                _ = v41
                _ = v42
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi39
        }
        var v43 float32 = float32(((float64((0.0 * 0.1)) + (0.0 + 0.0)) / 3.14159))
        _ = fi38
        _ = v43
    }
    if (18446744073709551615 != 653) {
        var v44 uint8 = 2
        switch -9223372036854775808 {
        case 2147483647:
            var v45 int64 = -9223372036854775808
            _ = v45
        case 32767:
            switch "a" {
            case "":
                if false {
                    v44 = v44
                }
            default:
                v46 := GenS4{F0: 65535, F1: 2147483647}
                _ = v46
            }
        default:
            if (((("foo bar" + ("hello" + "")) + (((("\n\t" + ("foo bar" + "a")) + "hello") + ("foo bar" + ("hello" + "a"))) + ("foo bar" + "foo bar"))) + "") == ("" + ((("" + ("hello" + "foo bar")) + "foo bar") + (("a" + "a") + "")))) {
                if true {
                    v44 -= (^(v44))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v47 uint = 255
                    _ = v47
                } else {
                    v48 := GenS3{F0: -91}
                    _ = v48
                }
                for fi49 := 0; fi49 < 1; fi49++ {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v2 = append(v2, "foo bar")
                        v44 = v44
                        _ = fi50
                    }
                    v51 := GenS5{F0: 2147483648, F1: 4294967295}
                    _ = fi49
                    _ = v51
                }
            }
        }
        switch 9223372036854775807 {
        case 128:
            fmt.Println(v44)
        }
        _ = v44
    }
    var v52 float32 = -2.25
    if (-115 < -880) {
        v52 += ((v52 * v52) - (3.14159 / (((3.14159 * (1.5 + float32((1e10 * (1e10 * (float64(1e10) - float64(v52))))))) / ((0.0 * (((3.14159 / (1e10 * v52)) / 3.14159) - 0.1)) * 1e10)) / ((((v52 / v52) + -2.25) / -2.25) / 0.1))))
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 int8 = -46
            _ = fi53
            _ = v54
        }
        v55 := GenS6{F0: 1, F1: ""}
        _ = v55
    } else {
        v52 = v52
        for fi56 := 0; fi56 < 8; fi56++ {
            v57 := GenS4{F0: 836, F1: fi56}
            var v58 uint32 = 2147483648
            v59 := GenS1{F0: (v58 - 0)}
            _ = fi56
            _ = v57
            _ = v58
            _ = v59
        }
        v60 := GenS7{F0: (false || (1 > 127)), F1: "hello", F2: 343}
        _ = v60
    }
    var v61 string = "a"
    fmt.Println(v61)
    v1 = append(v1, 9223372036854775807)
    var v62 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v61
    _ = v62
}
