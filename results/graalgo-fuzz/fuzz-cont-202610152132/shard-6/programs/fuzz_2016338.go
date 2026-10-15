package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int16; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 uint8, p1 bool, p2 string) uint16 {
    v5 := []int{4294967295, 9223372036854775807, -2147483648, 0}
    var v6 uint = 65535
    var v7 int64 = -64
    switch 128 {
    case -1:
        var v8 int8 = 0
        _ = v8
    case 32768:
        p2 = ("" + "foo bar")
    default:
        v9 := GenS1{F0: -1, F1: 42}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 417
}
func genFn1() uint64 {
    v10 := []int{814}
    v11 := GenS1{F0: 1, F1: 0}
    if (0 > 255) {
        var v12 int8 = 114
        if false {
            switch 0 {
            case 9223372036854775807:
                v12 += v12
            case 42:
                v13 := GenS2{F0: "\n\t"}
                _ = v13
            case 4294967295:
                switch "hello" {
                case "a":
                    if len(v10) > 0 {
                        v10[0] = 32768
                    }
                case "foo bar":
                    var v14 int32 = -641
                    _ = v14
                case "hello":
                    if len(v10) > 0 {
                        v10[0] = -827
                    }
                }
            default:
                v12 = (v12 & (v12 - -27))
            }
            v15 := GenS2{F0: ("\n\t" + ("" + "\n\t"))}
            _ = v15
        } else {
            if ((42 > -911) || ("foo bar" >= "hello")) {
                if (!false) {
                    var v16 int64 = -1
                    var v17 uint = 65535
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 uint8 = 0
                        _ = fi18
                        _ = v19
                    }
                }
            }
            v20 := GenS1{F0: -1, F1: -1}
            var v21 float32 = 1e10
            _ = v20
            _ = v21
        }
        if (104 == 2) {
            fmt.Println(v12)
        }
        _ = v12
    }
    if true {
        v22 := GenS2{F0: "\n\t"}
        v23 := GenS3{F0: 128}
        _ = v22
        _ = v23
    } else {
        var v24 uint = 42
        for fi25 := 0; fi25 < 19; fi25++ {
            v24 = v24
            _ = fi25
        }
        _ = v24
    }
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (2 > 18446744073709551615)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 3.14159
    for fi26 := 0; fi26 < 18; fi26++ {
        v4 += float64(1e10)
        var v27 int16 = 268
        _ = fi26
        _ = v27
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 int8 = -63
                _ = fi29
                _ = v30
            }
            v31 := GenS4{F0: "\n\t"}
            var v32 uint64 = 2
            _ = fi28
            _ = v31
            _ = v32
        }
    default:
        v4 = 0.1
    }
    fmt.Println(v2, v4)
    if false {
        v4 = 0.1
        v33 := []int{32767, 32768}
        if v2 {
            if v2 {
                fmt.Println(v4)
            } else {
                switch 127 {
                case 127:
                    v34 := GenS5{F0: 1, F1: 613}
                    _ = v34
                case 42:
                    switch -2147483648 {
                    case -822:
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            v33 = append(v33, 4294967295)
                        case 2147483647:
                            switch -670 {
                            case 725:
                                fmt.Println(v2)
                            case 0:
                                switch "foo bar" {
                                case "":
                                    if len(v33) >= 2 {
                                        v33 = v33[1:len(v33)]
                                    }
                                default:
                                    fmt.Println(v2, v4)
                                }
                            }
                        case 4294967295:
                            v33 = append(v33, 32768)
                        default:
                            v2 = false
                        }
                    default:
                        v2 = v2
                    }
                }
                for fi35 := 0; fi35 < 19; fi35++ {
                    var v36 uint32 = 0
                    for fi37 := 0; fi37 < 1; fi37++ {
                        fmt.Println(fi37, fi35)
                        var v38 bool = v2
                        if ((!false) && v2) {
                            var v39 int16 = 32767
                            delete(v3, v36)
                            if v38 {
                                v36 = (v36 * (v36 - (v36 << ((42) & 63))))
                                v40 := GenS4{F0: "foo bar"}
                                switch "" {
                                case "foo bar":
                                    var v41 uint64 = 1
                                    switch ("" + (((("a" + ("hello" + "\n\t")) + ("a" + ("a" + (("hello" + ("\n\t" + "foo bar")) + "foo bar")))) + (("foo bar" + "a") + "hello")) + "\n\t")) {
                                    case "hello":
                                        fmt.Println(v4, v39, v39)
                                    case "\n\t":
                                        if false {
                                            if v38 {
                                                v42 := GenS6{F0: v39, F1: 0, F2: 0}
                                                var v43 int16 = (v39 % ((int16(v41)) | 1))
                                                var v44 int = fi37
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                    default:
                                        if false {
                                            v45 := GenS3{F0: 128}
                                            var v46 uint32 = uint32(v41)
                                            fmt.Println(v38, fi35, v4)
                                            switch fi35 {
                                            case 42:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v47 int64 = -561
                                                    _ = v47
                                                case "":
                                                    var v48 bool = v38
                                                    _ = v48
                                                case "a":
                                                    var v49 string = "a"
                                                    _ = v49
                                                }
                                            case 32768:
                                                if (1.5 < -2.25) {
                                                    if v2 {
                                                        v38 = v2
                                                    }
                                                    var v50 uint = 0
                                                    _ = v50
                                                } else {
                                                    fmt.Println(v39)
                                                    var v51 float64 = 1e10
                                                    var v52 uint64 = genFn1()
                                                    v53 := GenS2{F0: "\n\t"}
                                                    _ = len(v33)
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                            case 9223372036854775807:
                                                v2 = v2
                                            }
                                            fmt.Println(fi37, v38)
                                            v38 = false
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v2 = true
                                            v41 += v41
                                        }
                                    }
                                    _ = v41
                                case "":
                                    fmt.Println(v39, fi35)
                                case "hello":
                                    v38 = (!(42 >= -723))
                                }
                                _ = v40
                            } else {
                                fmt.Println(v36, v4, v2)
                                v54 := GenS6{F0: 1, F1: 962, F2: 262}
                                _ = v54
                            }
                            _ = v39
                        } else {
                            var v55 uint32 = 1
                            _ = v55
                        }
                        _ = fi37
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v4, v4)
            }
            if v2 {
                var v56 float64 = v4
                _ = v56
            } else {
                v57 := GenS7{F0: 255}
                v4 -= v4
                _ = v57
            }
            switch -2147483648 {
            case 127:
                var v58 uint16 = 0
                _ = v58
            case 4294967295:
                for fi59 := 0; fi59 < 18; fi59++ {
                    v33 = append(v33, fi59)
                    var v60 int32 = -740
                    v4 = v4
                    _ = fi59
                    _ = v60
                }
            case 241:
                v4 = 1.5
            default:
                for fi61 := 0; fi61 < 10; fi61++ {
                    for fi62 := 0; fi62 < 14; fi62++ {
                        delete(v1, fi62)
                        v2 = (42 > 4294967295)
                        _ = fi62
                    }
                    _ = fi61
                }
            }
        } else {
            v2 = v2
        }
        _ = v33
    }
    _ = len(v3)
    v1[-27] = 32768
    v2 = true
    if (!(false || v2)) {
        for fi63 := 0; fi63 < 13; fi63++ {
            if (((-991 < -249) && (!v2)) && ("a" >= ("\n\t" + "foo bar"))) {
                switch "foo bar" {
                case "":
                    for fi64 := 0; fi64 < 15; fi64++ {
                        if ((fi64 > -9223372036854775808) || (("" + ("" + ("" + ((("\n\t" + "a") + "a") + "foo bar")))) == ("a" + "a"))) {
                            var v65 int8 = 114
                            v4 = v4
                            v66 := []int{-1, fi63, (fi64 ^ 128)}
                            _ = v65
                            _ = v66
                        }
                        v3[185] = 991
                        var v67 int64 = 9223372036854775807
                        _ = fi64
                        _ = v67
                    }
                case "hello":
                    v2 = false
                case "a":
                    v2 = ((3.14159 * ((3.14159 - 0.1) - 3.14159)) <= 1e10)
                default:
                    v68 := []int{}
                    _ = v68
                }
                if v2 {
                    delete(v1, -852)
                    v69 := GenS8{F0: 42, F1: (1 == 32767), F2: ("hello" + "")}
                    if ("\n\t" > "a") {
                        if ("hello" <= "hello") {
                            v70 := GenS5{F0: 255, F1: 32768}
                            _ = v70
                        }
                    }
                    _ = v69
                } else {
                    switch (fi63 ^ (fi63 + (^(fi63)))) {
                    case 127:
                        v71 := []int{}
                        _ = v71
                    }
                    v72 := []int{}
                    _ = len(v72)
                    _ = v72
                }
            }
            v73 := GenS3{F0: 1}
            switch -340 {
            case 9223372036854775807:
                v74 := GenS2{F0: "a"}
                _ = v74
            default:
                for fi75 := 0; fi75 < 6; fi75++ {
                    switch "a" {
                    case "foo bar":
                        var v76 int64 = 954
                        _ = v76
                    }
                    var v77 int64 = -514
                    v2 = v2
                    _ = fi75
                    _ = v77
                }
            }
            _ = fi63
            _ = v73
        }
    } else {
        v3[101] = 0
        v2 = (!v2)
        v78 := GenS4{F0: ""}
        _ = v78
    }
    var v79 int16 = -1
    if (!((v79 >> ((836) & 63)) > 0)) {
        v2 = (v2 && v2)
        var v80 float32 = (1.5 + 3.14159)
        _ = v80
    } else {
        if (0 > genFn0(0, true, ((("a" + "hello") + "\n\t") + ""))) {
            switch "foo bar" {
            case "foo bar":
                if (0 != -847) {
                    if v2 {
                        for fi81 := 0; fi81 < 14; fi81++ {
                            v79 = 15
                            v82 := GenS4{F0: "foo bar"}
                            _ = fi81
                            _ = v82
                        }
                    } else {
                        if (1 != 69) {
                            var v83 uint32 = 2147483648
                            var v84 float64 = v4
                            _ = v83
                            _ = v84
                        }
                        v85 := GenS3{F0: 2}
                        _ = v85
                    }
                    var v86 string = "a"
                    v87 := []int{-1, 4294967295}
                    _ = v86
                    _ = v87
                }
            }
            switch (("\n\t" + ("" + ("hello" + "foo bar"))) + "") {
            case "\n\t":
                for fi88 := 0; fi88 < 10; fi88++ {
                    v79 = (v79 / ((v79) | 1))
                    v79 -= 452
                    _ = fi88
                }
            }
        }
    }
    var v89 int = -9223372036854775808
    var v90 uint32 = 4294967295
    var v91 uint8 = 0
    for fi92 := 0; fi92 < 5; fi92++ {
        v91 = (^(v91))
        v90 += (v90 ^ v90)
        _ = fi92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v79
    _ = v89
    _ = v90
    _ = v91
}
