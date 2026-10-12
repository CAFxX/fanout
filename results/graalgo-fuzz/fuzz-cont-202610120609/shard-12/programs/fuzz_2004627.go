package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int; F2 int32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 string, p1 int, p2 string) int16 {
    var v4 uint32 = 499
    v5 := []int{(p1 ^ 4294967295), (p1 - p1), p1}
    v6 := GenS1{F0: "foo bar", F1: -365}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1() string {
    if false {
        v7 := GenS2{F0: 100}
        _ = v7
    }
    if true {
        switch "hello" {
        case "foo bar":
            var v8 string = "foo bar"
            _ = v8
        case "a":
            v9 := GenS1{F0: "foo bar", F1: -1}
            _ = v9
        case "":
            var v10 int16 = 42
            _ = v10
        }
        var v11 uint = 591
        for fi12 := 0; fi12 < 5; fi12++ {
            v11 = (^(v11))
            _ = fi12
        }
        _ = v11
    }
    if true {
        switch "" {
        case "hello":
            if true {
                if (-128 >= 0) {
                    v13 := GenS1{F0: "\n\t", F1: -2147483648}
                    _ = v13
                } else {
                    var v14 float64 = (3.14159 + 1e10)
                    var v15 uint64 = 1
                    fmt.Println(v15, v15)
                    _ = v14
                    _ = v15
                }
            }
        case "foo bar":
            v16 := GenS1{F0: "foo bar", F1: -16}
            _ = v16
        case "\n\t":
            if true {
                v17 := []int{4294967295, 1, -428, 9223372036854775807}
                _ = v17
            } else {
                v18 := GenS2{F0: -120}
                _ = v18
            }
        }
        var v19 float32 = 0.0
        var v20 uint8 = 128
        _ = v19
        _ = v20
    }
    v21 := []int{547}
    _ = v21
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 int64 = -628
    switch 32768 {
    case 931:
        if ((-336 < 32767) || ((((128 <= 169) || (42 <= 0)) && false) || (false && (((0 == 885) && true) && true)))) {
            var v22 int16 = genFn0("", int(v2), (("a" + "") + ("foo bar" + ("\n\t" + (("" + "hello") + "foo bar")))))
            v23 := v1[128]
            _ = v22
            _ = v23
        } else {
            v24 := []int{4294967295}
            if true {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := GenS1{F0: "foo bar", F1: 1}
                    if ((v3 | (v3 + 601)) >= (v3 % ((v3) | 1))) {
                        switch fi25 {
                        case 9223372036854775807:
                            var v27 int = 0
                            _ = v27
                        default:
                            switch ("\n\t" + "hello") {
                            case "":
                                v28 := GenS2{F0: int8(v2)}
                                _ = v28
                            case "\n\t":
                                v24 = append(v24, -1)
                            case "foo bar":
                                if false {
                                    for fi29 := 0; fi29 < 18; fi29++ {
                                        v3 = v3
                                        _ = len(v24)
                                        v30 := GenS1{F0: (("foo bar" + ("hello" + "")) + ""), F1: v2}
                                        _ = fi29
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 11; fi31++ {
                                        v32 := GenS1{F0: (("foo bar" + ("hello" + "hello")) + "\n\t"), F1: -1}
                                        v33 := v1[fi31]
                                        var v34 uint8 = 42
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                } else {
                                    switch 2147483647 {
                                    case -693:
                                        v35 := GenS1{F0: (("hello" + "\n\t") + "foo bar"), F1: v2}
                                        _ = v35
                                    }
                                    fmt.Println(v3, v2)
                                    v2 -= v2
                                }
                            default:
                                v36 := GenS1{F0: "hello", F1: (-(v2))}
                                _ = v36
                            }
                        }
                    } else {
                        v3 = v3
                        switch fi25 {
                        case 2147483647:
                            if (255 > 0) {
                                fmt.Println(fi25)
                            }
                        default:
                            for fi37 := 0; fi37 < 9; fi37++ {
                                var v38 string = "foo bar"
                                _ = len(v24)
                                v1[(fi25 << ((65535) & 63))] = 748
                                _ = fi37
                                _ = v38
                            }
                        }
                        v39 := GenS1{F0: "", F1: v2}
                        _ = v39
                    }
                    v3 += (^(v3))
                    _ = fi25
                    _ = v26
                }
                v2 -= 42
                delete(v1, 106)
            } else {
                _ = len(v24)
                var v40 int32 = -1
                _ = v40
            }
            var v41 uint = 177
            _ = v24
            _ = v41
        }
    case -199:
        if (!false) {
            v42 := []int{42, int(v3), -477}
            for fi43 := 0; fi43 < 8; fi43++ {
                if ((-33 == 1) || (true && true)) {
                    var v44 bool = (0 <= 196)
                    if true {
                        var v45 uint32 = 283
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 12; fi46++ {
                            var v47 string = "\n\t"
                            var v48 uint = 198
                            if ((^(v48)) < 1) {
                                switch v47 {
                                case "hello":
                                    if (!true) {
                                        if false {
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            v47 = (((v47 + v47) + v47) + ((v47 + v47) + v47))
                                            var v49 uint16 = 1
                                            _ = v49
                                        } else {
                                            var v50 bool = ((v3 | (v3 >> ((v48) & 63))) >= (v3 << ((42) & 63)))
                                            if (v50 || v44) {
                                                var v51 int8 = -17
                                                var v52 int32 = 2147483647
                                                _ = v51
                                                _ = v52
                                            } else {
                                                fmt.Println(v2, v50)
                                            }
                                            _ = v50
                                        }
                                    }
                                    fmt.Println(v44, v3)
                                case "\n\t":
                                    v44 = v44
                                }
                            }
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        switch "\n\t" {
                        case "hello":
                            for fi53 := 0; fi53 < 20; fi53++ {
                                v44 = (false && (!v44))
                                var v54 float64 = 0.1
                                _ = fi53
                                _ = v54
                            }
                        }
                        var v55 uint = 65535
                        _ = v55
                    }
                    var v56 int8 = 101
                    _ = v44
                    _ = v56
                }
                delete(v1, 4294967295)
                if (false && ((("foo bar" <= "\n\t") || (true && false)) && false)) {
                    var v57 uint8 = 255
                    _ = v57
                } else {
                    v58 := GenS3{F0: v2, F1: fi43, F2: v2}
                    _ = v58
                }
                _ = fi43
            }
            for fi59 := 0; fi59 < 13; fi59++ {
                if true {
                    var v60 bool = false
                    v1[fi59] = 42
                    switch "foo bar" {
                    case "foo bar":
                        for fi61 := 0; fi61 < 14; fi61++ {
                            var v62 float64 = 1.5
                            _ = fi61
                            _ = v62
                        }
                    case "\n\t":
                        v63 := GenS2{F0: 43}
                        _ = v63
                    case "":
                        if true {
                            v64 := v1[(fi59 + (fi59 + fi59))]
                            v65 := GenS4{F0: (fi59 / (((fi59 & 128)) | 1)), F1: 42}
                            _ = v64
                            _ = v65
                        }
                    default:
                        v42 = append(v42, 32768)
                    }
                    _ = v60
                } else {
                    if false {
                        v66 := GenS5{F0: 128}
                        v67 := GenS2{F0: -1}
                        v68 := GenS5{F0: 0}
                        _ = v66
                        _ = v67
                        _ = v68
                    } else {
                        fmt.Println(fi59, fi59, fi59)
                        v1[int(v3)] = 42
                    }
                }
                _ = fi59
            }
            _ = v42
        }
    default:
        if (784 < 722) {
            if true {
                var v69 bool = true
                for fi70 := 0; fi70 < 3; fi70++ {
                    var v71 int = (fi70 ^ -2147483648)
                    if v69 {
                        v69 = v69
                    }
                    _ = fi70
                    _ = v71
                }
                v72 := GenS6{F0: 213}
                _ = v69
                _ = v72
            } else {
                var v73 uint = 65535
                _ = v73
            }
            v2 -= v2
            var v74 float32 = float32(0.0)
            _ = v74
        }
    }
    var v75 int16 = 32767
    _ = len(v1)
    if (true && (0 == 744)) {
        for fi76 := 0; fi76 < 8; fi76++ {
            v77 := []int{(fi76 & (fi76 - 42))}
            _ = fi76
            _ = v77
        }
        var v78 int8 = -1
        fmt.Println(v78, v75, v2)
        _ = v78
    }
    v3 = 611
    v79 := GenS1{F0: "hello", F1: (v2 & v2)}
    fmt.Println(v3)
    var v80 float32 = (((0.0 / 0.1) + (1.5 - 3.14159)) / 0.1)
    fmt.Println(v75)
    var v81 uint = 65535
    fmt.Println(v75)
    v75 = v75
    if ((!false) && ("a" == (("\n\t" + "\n\t") + "a"))) {
        if (255 >= 2) {
            var v82 string = (((("\n\t" + ("\n\t" + ((("\n\t" + "\n\t") + "") + ("a" + "a")))) + "foo bar") + (("a" + (("hello" + "hello") + ("foo bar" + "a"))) + ("\n\t" + "\n\t"))) + "foo bar")
            _ = v82
        }
        switch 2147483647 {
        case 32767:
            var v83 float64 = 1e10
            _ = v83
        case 4294967295:
            if false {
                v84 := GenS7{F0: 0, F1: (v81 | (v81 & (v81 % (((v81 * (v81 & 248))) | 1))))}
                fmt.Println(v75, v75)
                for fi85 := 0; fi85 < 18; fi85++ {
                    if true {
                        v75 = (v75 & v75)
                        v86 := GenS4{F0: (fi85 - 32768), F1: int8(v75)}
                        _ = v86
                    } else {
                        var v87 uint16 = 65535
                        switch -757 {
                        case -955:
                            v88 := GenS3{F0: v2, F1: (^(fi85)), F2: v2}
                            _ = v88
                        default:
                            v89 := GenS2{F0: 127}
                            _ = v89
                        }
                        _ = v87
                    }
                    var v90 string = "foo bar"
                    if ((0.0 > -2.25) || ((^(fi85)) > 32767)) {
                        v90 = "hello"
                    }
                    _ = fi85
                    _ = v90
                }
                _ = v84
            }
        case 2147483647:
            v2 = 825
        default:
            v91 := GenS8{F0: "", F1: 700}
            _ = v91
        }
    } else {
        v92 := []int{32767, 32768, 0}
        v93 := GenS9{F0: 537}
        v75 = v75
        _ = v92
        _ = v93
    }
    var v94 int64 = v3
    v75 = genFn0("a", 4294967295, "foo bar")
    v80 -= 0.0
    if (1 < (^(v75))) {
        switch 2147483647 {
        case -396:
            v81 += v81
        case 771:
            v95 := GenS3{F0: (v2 >> (((^(v81))) & 63)), F1: 32768, F2: v2}
            _ = v95
        case -2147483648:
            switch "\n\t" {
            case "\n\t":
                var v96 uint8 = 206
                _ = v96
            }
        default:
            v75 = genFn0(("" + "hello"), 4294967295, "a")
        }
        fmt.Println(v3, v80, v3)
    } else {
        v2 = (v2 << ((v81) & 63))
        var v97 uint64 = 1
        v98 := []int{4294967295, 0, 32768}
        _ = v97
        _ = v98
    }
    switch (("a" + "\n\t") + "") {
    case "foo bar":
        fmt.Println(v75, v81)
    default:
        for fi99 := 0; fi99 < 9; fi99++ {
            v94 += (v94 << (((v81 + (^(v81)))) & 63))
            for fi100 := 0; fi100 < 20; fi100++ {
                v101 := []int{}
                switch -534 {
                case 2147483647:
                    _ = len(v1)
                case 9223372036854775807:
                    _ = len(v101)
                default:
                    switch ("" + (("\n\t" + "\n\t") + "foo bar")) {
                    case "hello":
                        fmt.Println(v3)
                        var v102 uint32 = 2147483648
                        fmt.Println(v75, v94, v94)
                        _ = v102
                    case "\n\t":
                        var v103 string = ("foo bar" + "")
                        _ = v103
                    }
                }
                _ = fi100
                _ = v101
            }
            if false {
                fmt.Println(v81, v80, v94)
            } else {
                v75 += (v75 | (-(v75)))
            }
            _ = fi99
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v75
    _ = v79
    _ = v80
    _ = v81
    _ = v94
}
