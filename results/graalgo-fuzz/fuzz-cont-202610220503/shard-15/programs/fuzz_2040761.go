package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 int32, p1 bool) int16 {
    var v6 int32 = p0
    if true {
        if (!(208 > 145)) {
            v6 = (p0 - p0)
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: p1, F1: p1, F2: -236}
                for fi9 := 0; fi9 < 7; fi9++ {
                    fmt.Println(fi9)
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
        } else {
            v10 := GenS2{F0: 2147483647, F1: "", F2: -1}
            _ = v10
        }
        v11 := []int{-9223372036854775808}
        _ = v11
    }
    if false {
        p0 = (p0 & (-(p0)))
    } else {
        p1 = (!(42 >= -1))
        v12 := []int{1, 4294967295, -1, 1}
        _ = v12
    }
    var v13 uint32 = 0
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 95
    v3 := []string{"foo bar", (("foo bar" + ("\n\t" + "foo bar")) + ""), ("foo bar" + ("hello" + "")), "foo bar"}
    v4 := make(map[int]bool)
    var v5 uint64 = 714
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            delete(v1, "")
            _ = fi15
        }
        if (1 != -1) {
            v2 -= v2
            v2 = (-(v2))
        } else {
            var v16 int64 = -9223372036854775808
            _ = v16
        }
        if (!((("hello" + "") + "") > ("foo bar" + ("hello" + "hello")))) {
            v17 := GenS3{F0: 2147483648}
            var v18 int8 = 1
            var v19 uint8 = 128
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 uint16 = 65535
                if ((!((float32(3.14159) - 0.0) > float32(((((3.14159 + (3.14159 * (-2.25 / 0.1))) * float64((3.14159 + 1e10))) + -2.25) * (((1.5 * 1e10) * (float64(0.1) - -2.25)) + float64(3.14159)))))) && (!(v2 >= (v2 + 1)))) {
                    var v22 int32 = (^(v2))
                    var v23 uint32 = 2
                    delete(v4, (-(fi14)))
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS4{F0: v2}
                    var v25 bool = ((^(fi14)) != fi20)
                    var v26 float32 = -2.25
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = fi20
                _ = v21
            }
            var v27 uint16 = 255
            _ = v27
        }
        _ = fi14
    }
    v5 += v5
    v2 = (v2 * v2)
    fmt.Println(v5, v5)
    v5 += 0
    if false {
        delete(v4, int(v5))
    } else {
        fmt.Println(v5, v5)
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v5, v2, fi28)
        for fi29 := 0; fi29 < 15; fi29++ {
            _ = len(v3)
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            var v30 uint64 = (v5 >> ((255) & 63))
            v31 := GenS5{F0: v30, F1: 42}
            _ = fi29
            _ = v30
            _ = v31
        }
        if false {
            v1["a"] = fi28
        }
        _ = fi28
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        fmt.Println(v2)
        var v33 string = ("\n\t" + "foo bar")
        if true {
            v3 = append(v3, "")
            var v34 string = v33
            _ = v34
        } else {
            fmt.Println(v2, v2, v2)
        }
        _ = fi32
        _ = v33
    }
    v2 += v2
    for fi35 := 0; fi35 < 3; fi35++ {
        if true {
            var v36 string = "foo bar"
            v2 = 1
            _ = v36
        }
        v2 += -347
        _ = fi35
    }
    fmt.Println(v2, v5)
    if true {
        if false {
            if (4294967295 > 123) {
                v1[""] = 127
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 uint16 = 32767
                    switch fi37 {
                    case 9223372036854775807:
                        var v39 uint = 678
                        _ = v39
                    case -9223372036854775808:
                        var v40 string = "hello"
                        _ = v40
                    case 32768:
                        delete(v4, (fi37 - 127))
                    }
                    _ = fi37
                    _ = v38
                }
                for fi41 := 0; fi41 < 9; fi41++ {
                    switch "hello" {
                    case "foo bar":
                        v42 := GenS5{F0: v5, F1: 824}
                        _ = v42
                    case "":
                        var v43 uint32 = 2
                        _ = v43
                    }
                    v5 -= 129
                    _ = fi41
                }
            }
            var v44 uint16 = 2
            var v45 int = 4294967295
            _ = v44
            _ = v45
        } else {
            var v46 int32 = -1
            v47 := GenS6{F0: ("foo bar" + "a"), F1: -32768}
            v2 = v46
            _ = v46
            _ = v47
        }
        var v48 float64 = (3.14159 + (1.5 / 1.0))
        v49 := GenS4{F0: v2}
        _ = v48
        _ = v49
    }
    if (954 != 816) {
        v2 = 1
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 int64 = 0
            var v52 string = ""
            _ = fi50
            _ = v51
            _ = v52
        }
        v2 = v2
    } else {
        var v53 float32 = -2.25
        v5 += v5
        var v54 int8 = -128
        _ = v53
        _ = v54
    }
    switch "" {
    case "a":
        fmt.Println(v2)
    default:
        for fi55 := 0; fi55 < 17; fi55++ {
            v56 := GenS2{F0: 155, F1: "a", F2: (fi55 * fi55)}
            for fi57 := 0; fi57 < 18; fi57++ {
                var v58 int16 = genFn0(v2, true)
                _ = fi57
                _ = v58
            }
            _ = fi55
            _ = v56
        }
    }
    if (!true) {
        for fi59 := 0; fi59 < 11; fi59++ {
            for fi60 := 0; fi60 < 6; fi60++ {
                v5 = (v5 % (((v5 % ((v5) | 1))) | 1))
                _ = fi60
            }
            var v61 int = (^(fi59))
            _ = fi59
            _ = v61
        }
        switch "\n\t" {
        case "\n\t":
            v2 = v2
        default:
            fmt.Println(v5, v5)
        }
        switch -285 {
        case 81:
            v62 := GenS1{F0: false, F1: true, F2: 486}
            _ = v62
        case 9223372036854775807:
            for fi63 := 0; fi63 < 6; fi63++ {
                fmt.Println(fi63, fi63, fi63)
                _ = fi63
            }
        case -1:
            for fi64 := 0; fi64 < 10; fi64++ {
                var v65 uint = 2
                v66 := GenS6{F0: ("hello" + "hello"), F1: -299}
                _ = fi64
                _ = v65
                _ = v66
            }
        }
    } else {
        if false {
            v67 := GenS4{F0: v2}
            for fi68 := 0; fi68 < 6; fi68++ {
                fmt.Println(v5)
                _ = fi68
            }
            _ = v67
        } else {
            if (v2 <= v2) {
                v5 -= (v5 | (v5 >> ((2) & 63)))
                var v69 int64 = 0
                _ = v69
            } else {
                var v70 uint64 = 4
                if (2147483648 >= 42) {
                    _ = len(v4)
                } else {
                    fmt.Println(v2, v5, v5)
                }
                if len(v3) > 0 {
                    v3[0] = ("a" + "\n\t")
                }
                _ = v70
            }
            v71 := GenS7{F0: -9223372036854775808}
            _ = v71
        }
        v72 := v4[0]
        _ = v72
    }
    for fi73 := 0; fi73 < 3; fi73++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi73
    }
    if false {
        var v74 uint8 = 190
        switch 588 {
        case 4294967295:
            switch "foo bar" {
            case "\n\t":
                v75 := v1[("" + "hello")]
                _ = v75
            case "a":
                v76 := GenS2{F0: (v2 / (((v2 / (((v2 >> ((2) & 63))) | 1))) | 1)), F1: ("foo bar" + "hello"), F2: 9223372036854775807}
                _ = v76
            case "":
                fmt.Println(v5, v5, v5)
            }
        case 0:
            for fi77 := 0; fi77 < 1; fi77++ {
                var v78 int32 = (^(v2))
                _ = fi77
                _ = v78
            }
        default:
            for fi79 := 0; fi79 < 6; fi79++ {
                var v80 bool = (960 == 0)
                var v81 float64 = (0.0 / -2.25)
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi79
                _ = v80
                _ = v81
            }
        }
        for fi82 := 0; fi82 < 9; fi82++ {
            var v83 string = (("\n\t" + ("foo bar" + ("" + "foo bar"))) + "a")
            fmt.Println(v5, v5)
            _ = fi82
            _ = v83
        }
        _ = v74
    } else {
        var v84 int16 = genFn0(-2147483648, true)
        fmt.Println(v2)
        _ = v84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
