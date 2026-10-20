package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint64; F2 string }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 string; F1 int16; F2 int32 }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 1
    var v3 uint8 = 1
    var v4 uint16 = 92
    v5 := make(map[int]string)
    v6 := GenS1{F0: "a"}
    if (false || (true || ((v3 * v3) >= v3))) {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(v4, fi7)
            for fi8 := 0; fi8 < 20; fi8++ {
                if (78 == 100) {
                    if true {
                        v3 = (v3 ^ v3)
                        v9 := GenS1{F0: ("a" + "hello")}
                        fmt.Println(fi7)
                        v3 += 153
                        _ = v9
                    }
                    v10 := GenS2{F0: (^(v4))}
                    v11 := GenS3{F0: v2, F1: "a"}
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            _ = fi7
        }
        var v12 int32 = v2
        _ = v12
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint8 = v3
            v4 = v4
            fmt.Println(v3, fi13, v4)
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if ((-(v2)) <= v2) {
                var v16 float64 = 1e10
                v17 := GenS4{F0: v2}
                v3 = 2
                _ = v16
                _ = v17
            }
            if ((v3 / ((247) | 1)) > 1) {
                fmt.Println(v4)
                v18 := GenS4{F0: v2}
                v19 := GenS5{F0: uint8(v4)}
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        v20 := GenS6{F0: v3, F1: 0}
        _ = v20
    }
    if (!(2 != 70)) {
        var v21 bool = false
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        switch fi22 {
        case 127:
            if (0 != 18446744073709551615) {
                var v23 int32 = v2
                var v24 int16 = 42
                _ = v23
                _ = v24
            }
        case 734:
            switch ("hello" + "foo bar") {
            case "a":
                for fi25 := 0; fi25 < 4; fi25++ {
                    switch "a" {
                    case "":
                        if true {
                            var v26 bool = (false || (((!(v4 < (v4 % ((v4) | 1)))) || ((-1 == 1) && (-710 != -1))) || ("a" >= ((((("" + "foo bar") + "hello") + "a") + ("foo bar" + "foo bar")) + (((("a" + "") + (("\n\t" + (("foo bar" + ("hello" + "foo bar")) + (("a" + "\n\t") + ("foo bar" + ("a" + ("a" + "hello")))))) + "a")) + (("foo bar" + "a") + ("hello" + ""))) + "a")))))
                            v27 := GenS4{F0: 774}
                            _ = v26
                            _ = v27
                        } else {
                            var v28 int32 = v2
                            _ = v28
                        }
                    }
                    _ = fi25
                }
            case "":
                if (!(!true)) {
                    fmt.Println(v4, v3)
                    if false {
                        var v29 int64 = 0
                        if false {
                            var v30 int64 = v29
                            var v31 float32 = -2.25
                            var v32 int16 = 386
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            switch "a" {
                            case "a":
                                v2 = (v2 - (v2 >> ((539) & 63)))
                            case "":
                                var v33 uint8 = v3
                                _ = v33
                            default:
                                v34 := []int{127, int(v2), (fi22 >> ((0) & 63))}
                                _ = v34
                            }
                            switch "foo bar" {
                            case "a":
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    if ((v3 * (v3 & v3)) != v3) {
                                        if true {
                                            v4 = (v4 - v4)
                                            if ((0.0 - 0.0) > (1e10 - (1.5 - 1e10))) {
                                                if ((3.14159 - (1e10 / float32((float64(((-2.25 + ((0.0 + float32(0.1)) / -2.25)) - 3.14159)) * 1.5)))) < 0.0) {
                                                    v36 := v1[0]
                                                    v37 := GenS4{F0: (v2 & (v2 >> ((255) & 63)))}
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v38 := GenS2{F0: 65535}
                                                    var v39 uint = 833
                                                    if true {
                                                        v40 := GenS7{F0: int64(v3), F1: uint64(v4), F2: "hello"}
                                                        v3 = v3
                                                        v29 = 1
                                                        _ = v40
                                                    }
                                                    v41 := []int{fi22, fi35, 555}
                                                    if (true && true) {
                                                        if (0 >= (v3 * (v3 << ((255) & 63)))) {
                                                            switch fi35 {
                                                            case 737:
                                                                fmt.Println(fi35)
                                                                var v42 uint32 = 65535
                                                                _ = v42
                                                            default:
                                                                var v43 int64 = int64(v39)
                                                                _ = v43
                                                            }
                                                            if true {
                                                                v29 -= (^(v29))
                                                            } else {
                                                                var v44 uint16 = (v4 * (v4 | (v4 * (v4 ^ (v4 + 711)))))
                                                                _ = v44
                                                            }
                                                            v2 = -2147483648
                                                        }
                                                    } else {
                                                        switch fi35 {
                                                        case -1:
                                                            _ = len(v41)
                                                        case -63:
                                                            var v45 uint8 = v3
                                                            _ = v45
                                                        case 9223372036854775807:
                                                            if (1 > 42) {
                                                                v46 := GenS6{F0: v3, F1: 1}
                                                                _ = v46
                                                            }
                                                        default:
                                                            var v47 int16 = 42
                                                            _ = v47
                                                        }
                                                        v29 -= v29
                                                        v48 := GenS4{F0: v2}
                                                        _ = v48
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                    _ = v41
                                                }
                                                v49 := v1[0]
                                                v50 := GenS7{F0: v29, F1: 42, F2: v49}
                                                v51 := GenS5{F0: (v3 / ((128) | 1))}
                                                v52 := v1[uint32(v3)]
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            } else {
                                                var v53 int = 32768
                                                _ = v53
                                            }
                                            if (!false) {
                                                v4 = 1
                                            } else {
                                                switch (fi22 ^ (fi35 + -2147483648)) {
                                                case 23:
                                                    v29 -= v29
                                                case 127:
                                                    v54 := GenS7{F0: v29, F1: 2, F2: ("" + "")}
                                                    _ = v54
                                                default:
                                                    var v55 bool = false
                                                    _ = v55
                                                }
                                                if false {
                                                    var v56 int = -2147483648
                                                    _ = v56
                                                }
                                            }
                                        } else {
                                            if ("\n\t" == (((("a" + "hello") + ("" + "\n\t")) + ("a" + "")) + (("\n\t" + ("a" + ("foo bar" + ("" + (("a" + "foo bar") + "\n\t"))))) + (("\n\t" + "a") + ((("foo bar" + "foo bar") + "\n\t") + "a"))))) {
                                                v3 = (^(v3))
                                                if ((v3 ^ v3) == (v3 + 2)) {
                                                    var v57 float64 = 0.1
                                                    var v58 uint32 = 2147483648
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    var v59 float32 = 0.1
                                                    if (v29 >= (v29 << ((0) & 63))) {
                                                        if (!false) {
                                                            var v60 uint8 = v3
                                                            _ = v60
                                                        }
                                                        v61 := []int{fi22}
                                                        fmt.Println(v59)
                                                        _ = v61
                                                    }
                                                    _ = v59
                                                }
                                                var v62 float32 = 1e10
                                                _ = v62
                                            } else {
                                                fmt.Println(v4, v4, v29)
                                                v29 = 0
                                                if (866 >= 42) {
                                                    switch "a" {
                                                    case "hello":
                                                        v3 = 0
                                                    }
                                                    v2 = -2147483648
                                                } else {
                                                    var v63 int = fi35
                                                    _ = v63
                                                }
                                            }
                                            v64 := []int{}
                                            if (("hello" == "foo bar") && (!(-139 < (v2 << ((1) & 63))))) {
                                                var v65 uint16 = (v4 ^ (v4 | (v4 + 161)))
                                                v2 -= (v2 ^ v2)
                                                var v66 uint64 = 42
                                                fmt.Println(v4)
                                                _ = v65
                                                _ = v66
                                            }
                                            switch (("a" + "hello") + ("a" + ("" + "hello"))) {
                                            case "":
                                                _ = len(v1)
                                            }
                                            v2 -= v2
                                            fmt.Println(v29)
                                            _ = v64
                                        }
                                        fmt.Println(v3, v2, v2)
                                    } else {
                                        v2 = (-(v2))
                                    }
                                    switch (-(fi35)) {
                                    case 4294967295:
                                        v67 := GenS7{F0: 9223372036854775807, F1: uint64(v4), F2: "hello"}
                                        v2 -= v2
                                        v68 := v1[1]
                                        if ((v4 <= v4) || (4294967295 > 65535)) {
                                            var v69 float32 = ((-2.25 - (float32(((0.1 - -2.25) + 3.14159)) / ((0.1 / 1.0) - (float32(1.5) / 0.1)))) - 3.14159)
                                            _ = v69
                                        } else {
                                            var v70 uint16 = (v4 - v4)
                                            v71 := []int{fi22}
                                            _ = v70
                                            _ = v71
                                        }
                                        _ = v67
                                        _ = v68
                                    default:
                                        fmt.Println(v2, v2)
                                    }
                                    v2 -= (-(v2))
                                    _ = fi35
                                }
                            case "":
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    var v73 int32 = -2147483648
                                    _ = fi72
                                    _ = v73
                                }
                            }
                        }
                        for fi74 := 0; fi74 < 10; fi74++ {
                            v75 := GenS8{F0: 42, F1: v2}
                            var v76 int32 = v2
                            var v77 uint8 = uint8(v4)
                            _ = fi74
                            _ = v75
                            _ = v76
                            _ = v77
                        }
                        _ = v29
                    } else {
                        for fi78 := 0; fi78 < 6; fi78++ {
                            v79 := GenS9{F0: "foo bar", F1: int16(v3), F2: (v2 >> ((205) & 63))}
                            v80 := GenS1{F0: ""}
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                        if ((!true) || ((!(int8(v3) != 83)) && (-1 < 100))) {
                            v4 -= v4
                        } else {
                            var v81 int = 2147483647
                            switch "a" {
                            case "hello":
                                for fi82 := 0; fi82 < 6; fi82++ {
                                    v83 := []int{fi22}
                                    v3 = v3
                                    var v84 uint64 = 18446744073709551615
                                    _ = fi82
                                    _ = v83
                                    _ = v84
                                }
                            default:
                                if (v81 != (v81 * (v81 >> ((42) & 63)))) {
                                    v81 += 1
                                } else {
                                    var v85 int16 = 32767
                                    switch ((("foo bar" + ("foo bar" + "foo bar")) + "\n\t") + "hello") {
                                    case "hello":
                                        v2 -= v2
                                    case "":
                                        var v86 float64 = 0.0
                                        _ = v86
                                    case "\n\t":
                                        if true {
                                            var v87 uint32 = 2
                                            fmt.Println(fi22, v2, v3)
                                            fmt.Println(v87)
                                            _ = v87
                                        } else {
                                            v88 := GenS9{F0: "", F1: v85, F2: v2}
                                            v89 := GenS3{F0: v2, F1: "hello"}
                                            fmt.Println(v3, v81, v3)
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                    _ = v85
                                }
                            }
                            _ = v81
                        }
                        var v90 bool = false
                        _ = v90
                    }
                }
            default:
                var v91 int8 = 84
                _ = v91
            }
        case 32768:
            var v92 float32 = float32(1.5)
            _ = v92
        default:
            v4 -= v4
        }
        if false {
            switch fi22 {
            case 1:
                switch fi22 {
                case 1:
                    var v93 int32 = (v2 << ((1) & 63))
                    _ = v93
                case 0:
                    v4 -= (v4 - (^(v4)))
                }
            default:
                var v94 bool = false
                _ = v94
            }
            v95 := GenS10{F0: int32(v4), F1: "foo bar"}
            _ = v95
        }
        v96 := GenS11{F0: "a"}
        _ = fi22
        _ = v96
    }
    var v97 uint = 4294967295
    delete(v5, 42)
    var v98 float64 = 0.0
    v99 := GenS5{F0: 2}
    var v100 bool = true
    v4 -= (v4 * v4)
    fmt.Println(v98)
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v101 int64 = 1
        _ = v101
    }
    var v102 uint8 = uint8(v4)
    v103 := []int{}
    var v104 int8 = 100
    fmt.Println(v3, v4)
    v105 := GenS5{F0: 255}
    var v106 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v104:", v104)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
    _ = v98
    _ = v99
}
