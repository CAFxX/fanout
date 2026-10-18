package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint16; F1 string; F2 string }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint8 }
func genFn0(p0 int32, p1 uint) uint64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            p1 += (p1 & (p1 << (((p1 << ((2) & 63))) & 63)))
            switch "hello" {
            case "":
                p0 = 0
            }
            _ = fi7
        }
        v8 := []int{fi6}
        _ = fi6
        _ = v8
    }
    p0 = (-(p0))
    switch "foo bar" {
    case "":
        var v9 float32 = (1e10 * (-2.25 - ((3.14159 - ((0.1 * (1.5 * 1e10)) / ((1.5 - (3.14159 + ((3.14159 + (3.14159 * (-2.25 - (1e10 + 0.1)))) / -2.25))) - (float32(3.14159) * (((-2.25 / 1.0) + 0.0) * float32(1.5)))))) - (0.1 - (-2.25 - (-2.25 * (float32(3.14159) - 1.5)))))))
        _ = v9
    }
    _ = p0
    _ = p1
    return uint64(p1)
}
func genFn1() bool {
    if (770 <= 42) {
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int16 = 0
                v12 := GenS1{F0: 1, F1: 184, F2: ("" + (("" + "foo bar") + ""))}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint32 = 1
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            if true {
                v15 := GenS2{F0: -766, F1: -1, F2: 0}
                _ = v15
            } else {
                var v16 int8 = 127
                _ = v16
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 string = "hello"
                _ = fi17
                _ = v18
            }
            var v19 int8 = -106
            _ = fi14
            _ = v19
        }
        if false {
            switch "foo bar" {
            case "foo bar":
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(fi20)
                    fmt.Println(fi20)
                    v21 := GenS1{F0: 2, F1: fi20, F2: "foo bar"}
                    _ = fi20
                    _ = v21
                }
            case "":
                switch -2147483648 {
                case 0:
                    var v22 uint32 = 2147483648
                    _ = v22
                case -1:
                    var v23 int = 42
                    _ = v23
                case 42:
                    var v24 uint = 1
                    _ = v24
                default:
                    v25 := []int{128, 32768, -1}
                    _ = v25
                }
            }
            var v26 bool = false
            _ = v26
        } else {
            var v27 int64 = 0
            v28 := GenS3{F0: -1}
            var v29 int = 526
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 uint32 = 42
        _ = v30
    }
    var v31 int8 = -32
    var v32 bool = (0.0 > -2.25)
    _ = v31
    _ = v32
    return v32
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = -603
    v3 := []string{"a", "foo bar"}
    var v4 uint = 255
    v5 := make(map[string]uint32)
    if (1 >= 255) {
        var v33 uint32 = v1
        v34 := GenS2{F0: v2, F1: v2, F2: (-(v2))}
        if false {
            for fi35 := 0; fi35 < 9; fi35++ {
                fmt.Println(v2)
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 uint32 = v33
                    fmt.Println(v1, v37)
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            fmt.Println(v2, v2)
        }
        _ = v33
        _ = v34
    }
    switch 9223372036854775807 {
    case 0:
        var v38 int16 = v2
        _ = v38
    case -9223372036854775808:
        switch "foo bar" {
        case "":
            if (!false) {
                var v39 float64 = (1.5 / (1e10 / 1.0))
                var v40 bool = true
                v2 += v2
                _ = v39
                _ = v40
            }
        }
    case 32767:
        v41 := GenS1{F0: 127, F1: 9223372036854775807, F2: "hello"}
        _ = v41
    default:
        var v42 int = 73
        _ = v42
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println(v2, v4, v4)
    v5[("hello" + (("hello" + "hello") + "hello"))] = v1
    if true {
        switch ("" + "") {
        case "\n\t":
            fmt.Println(v2)
        case "foo bar":
            fmt.Println(v1, v2, v1)
        default:
            v43 := GenS2{F0: v2, F1: v2, F2: 1000}
            _ = v43
        }
        fmt.Println(v1, v1, v2)
        _ = len(v3)
    }
    v44 := GenS4{F0: ("" <= "a")}
    if true {
        if true {
            var v45 float32 = float32((((-2.25 - ((-2.25 + (1e10 / float64(1e10))) + float64((0.1 - 0.0)))) - 0.0) * 0.0))
            if true {
                var v46 float64 = 1e10
                _ = v46
            } else {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            }
            _ = v45
        } else {
            switch 322 {
            case 128:
                var v47 uint8 = 255
                _ = v47
            default:
                var v48 int32 = 2147483647
                _ = v48
            }
            v49 := v5["a"]
            _ = v49
        }
        var v50 float64 = 0.0
        for fi51 := 0; fi51 < 13; fi51++ {
            var v52 int16 = v2
            var v53 bool = ((v2 ^ (v2 / ((v2) | 1))) <= v52)
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = v50
    }
    var v54 int16 = v2
    var v55 int = 42
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v56 string = "foo bar"
    if false {
        v4 = 2
        fmt.Println(v56, v56, v55)
    }
    v2 -= 0
    for fi57 := 0; fi57 < 19; fi57++ {
        v58 := GenS5{F0: (v1 / ((v1) | 1))}
        fmt.Println(v4)
        for fi59 := 0; fi59 < 7; fi59++ {
            var v60 float64 = (((1e10 * -2.25) + 1.5) * 1.5)
            _ = fi59
            _ = v60
        }
        _ = fi57
        _ = v58
    }
    for fi61 := 0; fi61 < 19; fi61++ {
        if false {
            if (1 != v1) {
                if (1 > 2) {
                    var v62 int = v55
                    var v63 float64 = 1.5
                    var v64 int8 = -47
                    _ = v62
                    _ = v63
                    _ = v64
                } else {
                    switch 42 {
                    case 4294967295:
                        fmt.Println(v56, v56, v2)
                    default:
                        var v65 uint64 = 18446744073709551615
                        _ = v65
                    }
                    var v66 int16 = v54
                    if false {
                        switch (^(fi61)) {
                        case -549:
                            v54 -= v66
                        case 1:
                            if false {
                                fmt.Println(v54)
                                var v67 uint32 = 65535
                                _ = v67
                            } else {
                                v68 := GenS6{F0: 0, F1: 42, F2: 42}
                                v69 := v5["hello"]
                                for fi70 := 0; fi70 < 13; fi70++ {
                                    var v71 uint32 = v69
                                    _ = fi70
                                    _ = v71
                                }
                                _ = v68
                                _ = v69
                            }
                        case 9223372036854775807:
                            v72 := GenS7{F0: 1, F1: v56, F2: ""}
                            _ = v72
                        }
                        var v73 uint16 = 1
                        _ = v73
                    } else {
                        v66 = -2
                    }
                    _ = v66
                }
            }
            switch (fi61 % (((v55 | (-(v55)))) | 1)) {
            case 860:
                fmt.Println(v4)
            default:
                v74 := v5["\n\t"]
                _ = v74
            }
            v56 = v56
        } else {
            fmt.Println(v54, v2)
            v55 = (v55 + (v55 - 136))
        }
        v75 := GenS6{F0: 0, F1: 42, F2: 2}
        if ((42 == 255) && ((!(-1 <= 9223372036854775807)) || (1e10 != (0.1 / 1.0)))) {
            v55 -= fi61
            switch ("a" + (((v56 + "") + v56) + "a")) {
            case "":
                v76 := GenS8{F0: -548, F1: (v2 | v2)}
                _ = v76
            default:
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            }
            var v77 float64 = (1e10 * 0.1)
            _ = v77
        }
        _ = fi61
        _ = v75
    }
    for fi78 := 0; fi78 < 5; fi78++ {
        var v79 float64 = 1.5
        v80 := GenS9{F0: int(v54)}
        switch (^(fi78)) {
        case 128:
            var v81 string = (v56 + ((v56 + v56) + v56))
            _ = v81
        }
        _ = fi78
        _ = v79
        _ = v80
    }
    v82 := GenS10{F0: -498, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v82
}
