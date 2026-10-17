package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 int8 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0() string {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{}
        if false {
            var v9 string = "foo bar"
            v10 := GenS1{F0: -128, F1: (42 >= 18446744073709551615), F2: 18446744073709551615}
            var v11 uint = 65535
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v12 float64 = ((-2.25 + ((1.5 + (0.1 * float64(0.1))) * float64((0.1 - (0.0 + (1.5 - 1.5)))))) - -2.25)
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int8 = 0
        _ = fi13
        _ = v14
    }
    v12 = float64(1.5)
    var v15 string = (("" + "hello") + (("foo bar" + (("\n\t" + ("a" + ("" + "\n\t"))) + "foo bar")) + "hello"))
    _ = v12
    _ = v15
    return "hello"
}
func genFn1(p0 int, p1 uint32, p2 int) string {
    var v16 uint = 809
    var v17 uint = 0
    v17 = (v17 << (((v16 << ((18446744073709551615) & 63))) & 63))
    switch int(v17) {
    case 1:
        var v18 float32 = (0.1 / 1.0)
        _ = v18
    case -254:
        var v19 int = (-(p2))
        _ = v19
    case 104:
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS1{F0: 0, F1: (p2 <= 127), F2: (v16 % ((45) | 1))}
            var v22 int = p2
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        var v23 uint16 = 3
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 int64 = 1
    var v4 uint16 = 32767
    v5 := make(map[int]uint32)
    var v6 int32 = 486
    var v24 int64 = v3
    for fi25 := 0; fi25 < 10; fi25++ {
        switch ("a" + "\n\t") {
        case "a":
            var v26 uint8 = 0
            _ = v26
        case "":
            if false {
                var v27 uint8 = 1
                _ = v27
            } else {
                switch fi25 {
                case 2147483647:
                    var v28 float32 = (v2 / (-2.25 - (v2 * v2)))
                    _ = v28
                case 128:
                    if (1 < 829) {
                        if false {
                            var v29 float64 = 0.1
                            _ = v29
                        } else {
                            var v30 float64 = float64(1e10)
                            switch "foo bar" {
                            case "foo bar":
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v32 := GenS1{F0: 1, F1: ((v3 & v24) >= (v24 / (((-(v24))) | 1))), F2: 1}
                                    if false {
                                        v4 = v4
                                    } else {
                                        v30 -= (float64(-2.25) - (1.5 * (v30 - 0.0)))
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            case "hello":
                                if false {
                                    var v33 uint32 = 42
                                    v6 += 961
                                    _ = v33
                                } else {
                                    var v34 int16 = 0
                                    if ("a" <= "\n\t") {
                                        for fi35 := 0; fi35 < 14; fi35++ {
                                            switch ((("\n\t" + ("a" + ((("foo bar" + (("hello" + "a") + "a")) + "\n\t") + ((("a" + "hello") + "\n\t") + ("hello" + ""))))) + "foo bar") + ((("a" + "a") + "foo bar") + "foo bar")) {
                                            case "a":
                                                var v36 int = fi25
                                                _ = v36
                                            }
                                            _ = fi35
                                        }
                                        if (!true) {
                                            v37 := []int{fi25}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4 += 42
                                            case "hello":
                                                v2 -= 0.0
                                            case "\n\t":
                                                var v38 int8 = 100
                                                _ = v38
                                            }
                                            v6 = (v6 / ((357) | 1))
                                            _ = v37
                                        }
                                        v39 := GenS1{F0: -1, F1: false, F2: 18446744073709551615}
                                        _ = v39
                                    } else {
                                        var v40 uint32 = 2
                                        if (65535 < 930) {
                                            switch "hello" {
                                            case "hello":
                                                for fi41 := 0; fi41 < 12; fi41++ {
                                                    fmt.Println(v24)
                                                    var v42 int32 = v6
                                                    var v43 uint64 = uint64(v4)
                                                    _ = fi41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            case "\n\t":
                                                v44 := []int{}
                                                _ = v44
                                            default:
                                                if true {
                                                    fmt.Println(fi25)
                                                } else {
                                                    var v45 int16 = (v34 - v34)
                                                    _ = v45
                                                }
                                            }
                                            delete(v1, "a")
                                            if (4294967295 <= 2) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v46 int = -926
                                                    _ = v46
                                                case "foo bar":
                                                    var v47 int64 = 42
                                                    _ = v47
                                                case "":
                                                    var v48 int32 = v6
                                                    _ = v48
                                                default:
                                                    v4 -= (v4 / (((v4 << ((133) & 63))) | 1))
                                                }
                                                if false {
                                                    var v49 int8 = 127
                                                    _ = v49
                                                }
                                                v30 = v30
                                            }
                                        }
                                        v3 -= -1
                                        _ = v40
                                    }
                                    _ = v34
                                }
                            default:
                                if true {
                                    var v50 int16 = -32768
                                    var v51 float64 = ((0.1 / float64((float32(v30) * (1e10 * v2)))) + float64(0.1))
                                    var v52 int8 = int8(v6)
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 int8 = 0
                                    var v54 bool = true
                                    _ = v53
                                    _ = v54
                                }
                            }
                            v55 := GenS1{F0: -1, F1: ("foo bar" > "hello"), F2: 65535}
                            _ = v30
                            _ = v55
                        }
                        fmt.Println(v3, v3)
                        if true {
                            v56 := []int{(^(fi25)), -114, fi25, fi25}
                            _ = v56
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            switch (fi25 >> ((42) & 63)) {
                            case 558:
                                var v57 uint16 = v4
                                _ = v57
                            case -614:
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    v2 -= -2.25
                                    var v59 int = fi58
                                    fmt.Println(v6, v59, v2)
                                    _ = fi58
                                    _ = v59
                                }
                            case 42:
                                v60 := GenS1{F0: 42, F1: false, F2: 42}
                                _ = v60
                            }
                        }
                        for fi61 := 0; fi61 < 3; fi61++ {
                            v2 = (v2 + v2)
                            var v62 int32 = (v6 - (v6 * (v6 << ((718) & 63))))
                            v63 := GenS2{F0: fi25}
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        var v64 int = fi25
                        _ = v64
                    }
                }
                var v65 uint8 = 255
                _ = v65
            }
        case "\n\t":
            if true {
                if false {
                    v66 := GenS3{F0: 289, F1: 128, F2: (v4 ^ (^(v4)))}
                    var v67 float32 = (v2 - (-2.25 - float32(1.5)))
                    if true {
                        v68 := GenS3{F0: 65535, F1: 127, F2: v4}
                        if (v6 > v6) {
                            v67 = float32((1e10 - float64(v2)))
                        }
                        _ = v68
                    } else {
                        fmt.Println(v4, v4, fi25)
                        var v69 float32 = 3.14159
                        v70 := GenS4{F0: v4, F1: 765, F2: 127}
                        _ = v69
                        _ = v70
                    }
                    _ = v66
                    _ = v67
                } else {
                    var v71 float32 = 1.5
                    _ = v71
                }
                for fi72 := 0; fi72 < 10; fi72++ {
                    switch (fi25 + -235) {
                    case 128:
                        v3 -= 870
                    case 934:
                        if ((v24 ^ (v24 % (((v24 / ((v3) | 1))) | 1))) == int64(v4)) {
                            v73 := GenS3{F0: 2147483648, F1: 0, F2: (^(v4))}
                            _ = v73
                        } else {
                            if (true && false) {
                                var v74 uint64 = 1
                                if false {
                                    var v75 uint64 = v74
                                    v6 += (v6 ^ (-(v6)))
                                    _ = v75
                                }
                                var v76 float64 = 1.5
                                _ = v74
                                _ = v76
                            } else {
                                fmt.Println(v24)
                                if false {
                                    if false {
                                        var v77 int32 = v6
                                        v78 := GenS5{F0: (-(v24)), F1: v4, F2: 2}
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                var v79 uint16 = v4
                                _ = v79
                            }
                            v6 = (v6 & (v6 / (((v6 | (v6 << ((624) & 63)))) | 1)))
                        }
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            v80 := []int{fi72, fi72, (fi25 | fi72), 0}
                            var v81 bool = (!(!false))
                            _ = v80
                            _ = v81
                        case "\n\t":
                            v82 := GenS5{F0: (v24 & 9223372036854775807), F1: (v4 | v4), F2: 1}
                            _ = v82
                        default:
                            v83 := []int{fi25, fi25}
                            _ = v83
                        }
                    }
                    var v84 int8 = 1
                    _ = fi72
                    _ = v84
                }
                v3 += v24
            }
        }
        _ = fi25
    }
    for fi85 := 0; fi85 < 8; fi85++ {
        fmt.Println(v4)
        _ = fi85
    }
    for fi86 := 0; fi86 < 18; fi86++ {
        switch 488 {
        case 127:
            var v87 int16 = -1
            _ = v87
        }
        _ = len(v5)
        delete(v5, (fi86 >> ((1) & 63)))
        _ = fi86
    }
    _ = len(v1)
    for fi88 := 0; fi88 < 17; fi88++ {
        switch 133 {
        case 128:
            v89 := GenS4{F0: v4, F1: fi88, F2: 0}
            _ = v89
        case -174:
            fmt.Println(v6)
        }
        _ = fi88
    }
    for fi90 := 0; fi90 < 9; fi90++ {
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "foo bar":
            var v91 string = "\n\t"
            _ = v91
        case "hello":
            var v92 float32 = (v2 * v2)
            _ = v92
        case "":
            for fi93 := 0; fi93 < 19; fi93++ {
                if true {
                    var v94 int16 = 478
                    var v95 uint8 = 1
                    _ = v94
                    _ = v95
                }
                _ = fi93
            }
        default:
            v96 := []int{(^(fi90)), -1}
            _ = v96
        }
        v24 = v24
        v6 += (-(v6))
        _ = fi90
    }
    v97 := v5[309]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
