package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint8) int16 {
    v5 := GenS1{F0: 40, F1: (32768 == 42), F2: -12}
    fmt.Println(p0)
    p0 += 201
    v6 := []int{2147483647}
    var v7 int16 = 32767
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (v7 << ((1) & 63))
}
func genFn1() uint64 {
    switch "foo bar" {
    case "":
        switch 0 {
        case 32768:
            v8 := GenS1{F0: 774, F1: true, F2: -128}
            _ = v8
        }
    case "hello":
        var v9 int = -1
        _ = v9
    }
    var v10 uint8 = 11
    var v11 uint32 = 1
    if (true || false) {
        var v12 bool = false
        for fi13 := 0; fi13 < 1; fi13++ {
            if (v11 < v11) {
                switch "a" {
                case "hello":
                    v14 := []int{fi13}
                    _ = v14
                case "a":
                    v15 := []int{-9223372036854775808}
                    _ = v15
                case "foo bar":
                    v11 -= v11
                default:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        if (0 <= 0) {
                            var v17 int8 = 42
                            _ = v17
                        }
                        _ = fi16
                    }
                }
                v12 = v12
                if (true || false) {
                    v12 = (!v12)
                    v10 = v10
                } else {
                    v18 := GenS2{F0: 1, F1: (v10 - (v10 & v10)), F2: ((1e10 / (float64(((-2.25 - (1.5 * 3.14159)) * (float32(1e10) / (0.0 - 1e10)))) / 1.0)) <= 0.1)}
                    v19 := GenS1{F0: v11, F1: true, F2: 127}
                    if false {
                        v11 -= 161
                        switch (("\n\t" + ("a" + "hello")) + "\n\t") {
                        case "foo bar":
                            switch ((("a" + ("\n\t" + "foo bar")) + "hello") + "\n\t") {
                            case "foo bar":
                                var v20 int32 = 42
                                _ = v20
                            case "\n\t":
                                v21 := GenS2{F0: -128, F1: v10, F2: ("a" != "foo bar")}
                                _ = v21
                            case "a":
                                v22 := GenS2{F0: 42, F1: (v10 & v10), F2: (0 == 1)}
                                _ = v22
                            }
                        }
                        if (3.14159 >= (0.1 / 1.0)) {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                var v24 int16 = 600
                                v10 = v10
                                _ = fi23
                                _ = v24
                            }
                        }
                    } else {
                        v11 += 4294967295
                    }
                    _ = v18
                    _ = v19
                }
            } else {
                if v12 {
                    if v12 {
                        var v25 uint32 = v11
                        _ = v25
                    }
                    v26 := GenS3{F0: fi13, F1: 734, F2: 42}
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                switch int(v10) {
                case -9223372036854775808:
                    v28 := []int{}
                    _ = v28
                default:
                    switch (fi13 + fi27) {
                    case -2147483648:
                        v29 := []int{1, (fi13 ^ (fi13 / ((fi27) | 1))), 431, (fi13 << ((255) & 63))}
                        _ = v29
                    }
                }
                v11 = (v11 % ((v11) | 1))
                v11 -= 631
                v30 := GenS4{F0: v10, F1: 2147483647}
                _ = fi27
                _ = v30
            }
            _ = fi13
        }
        _ = v12
    } else {
        switch "hello" {
        case "a":
            if false {
                fmt.Println(v11)
                v10 = (v10 % ((v10) | 1))
            }
        case "hello":
            var v31 int64 = -1
            _ = v31
        case "\n\t":
            v32 := []int{-1, -2147483648}
            _ = v32
        default:
            v33 := []int{2147483647, 127}
            _ = v33
        }
    }
    switch 32767 {
    case 32767:
        v34 := GenS1{F0: v11, F1: (false && (!(4294967295 != -2147483648))), F2: 101}
        _ = v34
    default:
        v11 = v11
    }
    _ = v10
    _ = v11
    return 858
}

func main() {
    var v1 uint8 = 194
    var v2 bool = true
    var v3 int32 = 42
    var v4 uint32 = 2147483648
    var v35 int8 = 0
    v3 = (v3 + v3)
    v36 := []int{-9223372036854775808}
    switch -929 {
    case 42:
        v2 = (v2 && (1.5 == 1.5))
    }
    v1 = 231
    if len(v36) > 0 {
        v36[0] = 2147483647
    }
    if false {
        if true {
            fmt.Println(v1, v4)
            v37 := GenS3{F0: 42, F1: 0, F2: 2}
            if v2 {
                v38 := GenS3{F0: 32767, F1: 1, F2: 275}
                _ = v38
            } else {
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
            }
            _ = v37
        } else {
            v35 += v35
            var v39 int32 = (v3 | v3)
            var v40 int8 = (^(v35))
            _ = v39
            _ = v40
        }
        var v41 float32 = 3.14159
        _ = v41
    }
    v42 := GenS2{F0: v35, F1: v1, F2: v2}
    v4 = 2
    for fi43 := 0; fi43 < 7; fi43++ {
        v44 := GenS2{F0: (v35 + v35), F1: 1, F2: true}
        switch (fi43 % ((fi43) | 1)) {
        case -251:
            var v45 uint = 18446744073709551615
            _ = v45
        case -2147483648:
            v2 = v2
        case 1:
            if (v4 > 0) {
                v2 = v2
                v46 := GenS2{F0: 127, F1: (v1 << ((42) & 63)), F2: false}
                if (v2 && (v2 || v2)) {
                    for fi47 := 0; fi47 < 2; fi47++ {
                        v35 -= (^(v35))
                        if (18446744073709551615 <= 752) {
                            var v48 uint8 = v1
                            var v49 uint64 = uint64(v1)
                            _ = v48
                            _ = v49
                        } else {
                            var v50 uint16 = 605
                            if v2 {
                                if true {
                                    var v51 int8 = 0
                                    var v52 int64 = 9223372036854775807
                                    var v53 uint8 = v1
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    switch ("\n\t" + ("a" + "")) {
                                    case "hello":
                                        switch "hello" {
                                        case "hello":
                                            _ = len(v36)
                                        case "a":
                                            var v54 float32 = (-2.25 * 0.0)
                                            _ = v54
                                        case "":
                                            var v55 uint32 = v4
                                            _ = v55
                                        }
                                    case "\n\t":
                                        v4 = v4
                                    case "foo bar":
                                        _ = len(v36)
                                    default:
                                        if (v4 > 0) {
                                            var v56 uint32 = 258
                                            _ = v56
                                        }
                                    }
                                }
                                if (true && (true || v2)) {
                                    var v57 int64 = 42
                                    _ = v57
                                } else {
                                    v4 -= 949
                                }
                                v36 = append(v36, fi43)
                                v3 = v3
                            } else {
                                if v2 {
                                    if v2 {
                                        if ((v50 | (v50 - (v50 ^ v50))) < 128) {
                                            var v58 float64 = (0.0 - ((3.14159 + 0.0) / 0.1))
                                            _ = v58
                                        } else {
                                            _ = len(v36)
                                        }
                                    }
                                }
                                fmt.Println(v3)
                                var v59 int = fi43
                                _ = v59
                            }
                            _ = len(v36)
                            _ = v50
                        }
                        _ = fi47
                    }
                }
                _ = v46
            }
        default:
            switch "a" {
            case "":
                v4 = (v4 & (v4 + v4))
            case "a":
                for fi60 := 0; fi60 < 18; fi60++ {
                    switch (fi60 | (-(fi60))) {
                    case 32767:
                        v4 -= 2
                    default:
                        _ = len(v36)
                    }
                    if v2 {
                        var v61 uint16 = 65535
                        if v2 {
                            if true {
                                fmt.Println(fi60, fi60)
                            }
                            if (!v2) {
                                fmt.Println(v3)
                            }
                            var v62 string = "foo bar"
                            _ = v62
                        } else {
                            var v63 float32 = -2.25
                            v36 = append(v36, (fi60 >> ((255) & 63)))
                            _ = v63
                        }
                        v64 := GenS5{F0: int8(v4), F1: -32768}
                        _ = v61
                        _ = v64
                    } else {
                        var v65 uint = 334
                        _ = v65
                    }
                    if v2 {
                        if true {
                            v66 := GenS6{F0: v35}
                            var v67 int32 = (v3 / ((v3) | 1))
                            _ = v66
                            _ = v67
                        } else {
                            if ((!(v4 != v4)) && v2) {
                                var v68 uint64 = genFn1()
                                v35 += v35
                                if (v3 <= (v3 + -1)) {
                                    var v69 uint32 = 769
                                    var v70 int64 = -1
                                    var v71 int = fi43
                                    switch (^(fi43)) {
                                    case 1:
                                        var v72 uint16 = 0
                                        _ = v72
                                    case 0:
                                        v2 = (int16(v1) != 137)
                                    case 27:
                                        if ((-1 <= (-(v70))) || false) {
                                            if false {
                                                v36 = append(v36, fi60)
                                                fmt.Println(fi60, fi60)
                                                v35 = (^(v35))
                                            }
                                            if len(v36) > 0 {
                                                v36[0] = v71
                                            }
                                        }
                                    }
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                                v73 := GenS7{F0: (!false), F1: ("a" + ("foo bar" + (((("" + "") + "") + ("hello" + "")) + "a")))}
                                _ = v68
                                _ = v73
                            }
                            _ = len(v36)
                            var v74 int = fi60
                            _ = v74
                        }
                        v36 = append(v36, fi60)
                        fmt.Println(fi43, v4)
                    } else {
                        var v75 string = "\n\t"
                        if v2 {
                            fmt.Println(v75, v75, v1)
                            v4 -= v4
                        }
                        var v76 int64 = 6
                        v3 = 0
                        fmt.Println(v3, v75, v2)
                        _ = v75
                        _ = v76
                    }
                    _ = fi60
                }
            default:
                for fi77 := 0; fi77 < 4; fi77++ {
                    v3 -= v3
                    v3 = (v3 | -868)
                    v35 = (v35 ^ 43)
                    _ = fi77
                }
            }
        }
        var v78 float64 = 0.1
        _ = fi43
        _ = v44
        _ = v78
    }
    switch (("a" + "") + "") {
    case "foo bar":
        v3 = (v3 * (v3 + v3))
    case "":
        v2 = ((v1 >> ((2) & 63)) <= (v1 << ((18446744073709551615) & 63)))
    case "\n\t":
        fmt.Println(v4, v1)
    default:
        if false {
            v79 := GenS2{F0: v35, F1: 0, F2: (false || v2)}
            var v80 bool = v2
            v1 -= (v1 ^ 42)
            _ = v79
            _ = v80
        } else {
            for fi81 := 0; fi81 < 12; fi81++ {
                if (v4 > v4) {
                    fmt.Println(v4, v3, v2)
                    var v82 int16 = 0
                    _ = v82
                }
                for fi83 := 0; fi83 < 9; fi83++ {
                    v36 = append(v36, fi81)
                    switch ("\n\t" + "a") {
                    case "a":
                        if true {
                            fmt.Println(fi81, fi81, v1)
                            var v84 uint64 = 893
                            _ = v84
                        }
                    case "hello":
                        v85 := GenS8{F0: -326}
                        _ = v85
                    }
                    if ((v1 << ((4294967295) & 63)) != 2) {
                        var v86 uint16 = 736
                        var v87 int64 = 42
                        fmt.Println(v3)
                        _ = v86
                        _ = v87
                    }
                    if (!(0.0 < 0.0)) {
                        var v88 int = (fi83 ^ fi81)
                        var v89 uint8 = 255
                        _ = v88
                        _ = v89
                    }
                    var v90 int32 = v3
                    _ = fi83
                    _ = v90
                }
                v1 = (v1 | (v1 >> ((475) & 63)))
                _ = fi81
            }
        }
    }
    for fi91 := 0; fi91 < 13; fi91++ {
        if (v2 && v2) {
            for fi92 := 0; fi92 < 13; fi92++ {
                v1 += (v1 - (v1 << ((18446744073709551615) & 63)))
                _ = fi92
            }
            var v93 uint32 = (v4 | v4)
            _ = v93
        } else {
            if len(v36) > 0 {
                v36[0] = fi91
            }
            if true {
                var v94 float32 = 1.5
                v95 := GenS1{F0: v4, F1: v2, F2: (v35 * -1)}
                _ = v94
                _ = v95
            }
        }
        fmt.Println(v4, v35, v1)
        if (((float32((0.1 - ((0.0 / 1.5) / (1e10 / 1.5)))) / 1.5) - ((3.14159 * ((-2.25 / (-2.25 * float32(1.5))) * 3.14159)) / 0.1)) > 0.1) {
            var v96 int8 = 100
            _ = v96
        } else {
            fmt.Println(v3)
        }
        _ = fi91
    }
    v97 := GenS6{F0: v35}
    var v98 uint = 0
    v99 := GenS7{F0: v2, F1: "a"}
    fmt.Println(v2, v98)
    v98 = (v98 & v98)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v97
    _ = v98
    _ = v99
}
