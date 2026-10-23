package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int64; F1 int64; F2 uint32 }
type GenS11 struct { F0 int16; F1 uint64; F2 string }
type GenS12 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS13 struct { F0 bool }
func genFn0(p0 int, p1 int64) bool {
    v4 := GenS1{F0: (!false), F1: 255, F2: 1}
    v5 := []int{p0, p0}
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: ((-1 >= -32768) && true), F1: 128, F2: 2}
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        var v9 int = p0
        v9 = p0
        v9 += v9
        _ = v9
    }
    if (637 == 193) {
        if true {
            switch 0 {
            case 128:
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint64 = 0
                    var v12 string = "foo bar"
                    if true {
                        var v13 int = (fi10 - p0)
                        var v14 int64 = 0
                        var v15 int16 = 170
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case -1:
                if (false && (((6 < 0) || true) && false)) {
                    v16 := GenS2{F0: -1, F1: 567, F2: 42}
                    v17 := GenS1{F0: (false || false), F1: 128, F2: 495}
                    _ = len(v5)
                    _ = v16
                    _ = v17
                }
            case 1:
                switch ("" + "") {
                case "":
                    var v18 uint32 = 705
                    _ = v18
                }
            }
        } else {
            fmt.Println(p1, p1)
            switch (p0 << ((18446744073709551615) & 63)) {
            case -952:
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
            default:
                if false {
                    v19 := GenS3{F0: 1, F1: ((p1 & p1) > p1), F2: 56}
                    _ = v19
                }
            }
            if (!(229 < 42)) {
                var v20 uint8 = 127
                _ = v20
            }
        }
    } else {
        p1 = p1
        var v21 uint16 = 65535
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (!(!(false || false)))
}
func genFn1(p0 bool, p1 uint8) string {
    if (0 <= 2) {
        fmt.Println(p0, p1)
    }
    var v22 bool = (-582 == 82)
    v23 := GenS4{F0: 1, F1: 0}
    v24 := []int{}
    switch "hello" {
    case "a":
        if false {
            p1 += (p1 & 15)
            var v25 uint8 = 10
            _ = v25
        } else {
            var v26 uint8 = p1
            v27 := GenS5{F0: 42, F1: -1}
            _ = v26
            _ = v27
        }
    default:
        v28 := GenS2{F0: 2147483647, F1: 2147483647, F2: 42}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    return "hello"
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    v3 := []uint32{}
    if true {
        var v29 uint32 = 65535
        var v30 uint = 2
        var v31 int32 = 2147483647
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 uint64 = 0
    v1 -= (^(v32))
    delete(v2, "foo bar")
    v33 := GenS2{F0: 2147483647, F1: -9223372036854775808, F2: -128}
    if (2 == 149) {
        for fi34 := 0; fi34 < 18; fi34++ {
            fmt.Println(v32, fi34)
            _ = len(v2)
            _ = fi34
        }
    }
    v3 = append(v3, 2147483648)
    var v35 uint16 = 0
    switch (((("" + "a") + "a") + (("" + ("foo bar" + "hello")) + "")) + (("foo bar" + "a") + ("foo bar" + (("foo bar" + "\n\t") + ("\n\t" + "foo bar"))))) {
    case "foo bar":
        for fi36 := 0; fi36 < 17; fi36++ {
            for fi37 := 0; fi37 < 17; fi37++ {
                switch 9223372036854775807 {
                case -101:
                    switch (fi36 ^ fi37) {
                    case 228:
                        var v38 uint8 = 193
                        _ = v38
                    }
                default:
                    var v39 int8 = 1
                    _ = v39
                }
                if (!(!(!true))) {
                    var v40 uint32 = 2147483648
                    _ = v40
                }
                var v41 int16 = -32768
                fmt.Println(fi36)
                delete(v2, "\n\t")
                var v42 uint64 = v32
                if ((2 <= 42) || true) {
                    v42 = (^(v32))
                    v42 -= v42
                }
                if ((true || (1 >= (^(v1)))) && (!true)) {
                    var v43 int16 = -1
                    v1 = v32
                    v44 := GenS6{F0: 0, F1: true}
                    var v45 bool = (!true)
                    v42 += (^(v32))
                    switch (-(fi37)) {
                    case 87:
                        switch "\n\t" {
                        case "foo bar":
                            v1 += 916
                        case "\n\t":
                            switch (("foo bar" + (("hello" + ("\n\t" + "foo bar")) + "\n\t")) + ("a" + "foo bar")) {
                            case "":
                                if v45 {
                                    fmt.Println(v43, v1, v35)
                                    var v46 string = (("a" + ("\n\t" + ("foo bar" + ""))) + "\n\t")
                                    switch fi36 {
                                    case -2147483648:
                                        v35 = v35
                                    case 2147483647:
                                        if (127 < -1) {
                                            fmt.Println(v45, v42)
                                            var v47 int = 4294967295
                                            _ = v47
                                        } else {
                                            switch (("\n\t" + ((("" + ("hello" + v46)) + (("" + "hello") + (v46 + v46))) + "foo bar")) + "foo bar") {
                                            case "foo bar":
                                                if (fi37 > (fi37 ^ fi37)) {
                                                    var v48 int = -779
                                                    v49 := v2["\n\t"]
                                                    _ = len(v3)
                                                    _ = v48
                                                    _ = v49
                                                }
                                            default:
                                                if v45 {
                                                    if (1e10 <= (-2.25 + 3.14159)) {
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483648
                                                        }
                                                        var v50 int32 = 42
                                                        _ = len(v2)
                                                        _ = v50
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 42
                                                    }
                                                    fmt.Println(fi37, fi36, v41)
                                                }
                                            }
                                            v51 := GenS7{F0: v41, F1: 1}
                                            _ = v51
                                        }
                                    case 0:
                                        if false {
                                            var v52 int16 = 1
                                            _ = v52
                                        }
                                        v53 := GenS8{F0: -594, F1: 9223372036854775807}
                                        _ = v53
                                    }
                                    _ = v46
                                } else {
                                    switch (-(fi37)) {
                                    case 0:
                                        switch (fi36 << ((80) & 63)) {
                                        case 0:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            var v54 uint = 0
                                            _ = v54
                                        }
                                    case 9223372036854775807:
                                        v35 -= v35
                                    }
                                    v55 := GenS5{F0: 42, F1: 1}
                                    _ = v55
                                }
                            case "hello":
                                v41 += (v43 >> ((2) & 63))
                            case "\n\t":
                                fmt.Println(v35, v35, v41)
                            }
                        default:
                            if ((0 != 2147483647) && v45) {
                                var v56 string = (((("a" + (("\n\t" + (("" + "\n\t") + "hello")) + (("a" + "") + ("" + ("" + (("\n\t" + "hello") + "foo bar")))))) + ("\n\t" + ("hello" + "hello"))) + "hello") + ("hello" + ""))
                                v43 += int16(v42)
                                v57 := GenS5{F0: 0, F1: 9223372036854775807}
                                _ = v56
                                _ = v57
                            } else {
                                fmt.Println(v43, v42, fi37)
                                v42 += (v1 << ((65535) & 63))
                                var v58 uint = 1
                                if (!(v58 == 0)) {
                                    v59 := GenS9{F0: 1, F1: 2147483647}
                                    _ = v59
                                } else {
                                    v41 = 829
                                    var v60 float64 = float64(float32(3.14159))
                                    _ = v60
                                }
                                v61 := GenS2{F0: -6, F1: 135, F2: -128}
                                if (v41 > v43) {
                                    var v62 int8 = -126
                                    _ = v62
                                } else {
                                    if (("\n\t" + ("hello" + "hello")) == ("a" + "")) {
                                        v32 = (^(v32))
                                        var v63 uint32 = 0
                                        var v64 string = "hello"
                                        _ = v63
                                        _ = v64
                                    }
                                    if v45 {
                                        if true {
                                            delete(v2, "\n\t")
                                            fmt.Println(v35, v58, v42)
                                        } else {
                                            fmt.Println(v43, v42)
                                            var v65 int64 = 703
                                            v42 += v1
                                            v2["foo bar"] = false
                                            _ = v65
                                        }
                                        var v66 uint16 = v35
                                        switch fi37 {
                                        case -778:
                                            var v67 uint64 = v32
                                            _ = v67
                                        }
                                        v42 = (v32 >> ((65535) & 63))
                                        var v68 int8 = 100
                                        _ = v66
                                        _ = v68
                                    }
                                    switch (fi37 / ((fi37) | 1)) {
                                    case 1:
                                        if false {
                                            fmt.Println(v58, v42)
                                        }
                                    default:
                                        v69 := GenS10{F0: -834, F1: -512, F2: 65535}
                                        _ = v69
                                    }
                                    var v70 uint = (v58 - (^(v58)))
                                    if false {
                                        v1 = 42
                                    } else {
                                        var v71 uint16 = v35
                                        v72 := GenS11{F0: (v43 * -1), F1: (v1 * (v1 | 1)), F2: ""}
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v70
                                }
                                _ = v58
                                _ = v61
                            }
                        }
                    case 446:
                        v43 += (v41 * v41)
                    case -9223372036854775808:
                        var v73 int32 = 42
                        _ = v73
                    }
                    if v45 {
                        if v45 {
                            var v74 int32 = -832
                            _ = v74
                        } else {
                            fmt.Println(fi36, v41, v35)
                            var v75 float32 = 1.5
                            v2["a"] = v45
                            v35 += v35
                            var v76 int16 = v43
                            switch -2147483648 {
                            case 32768:
                                if (42 == 127) {
                                    switch (fi37 ^ fi37) {
                                    case 127:
                                        v77 := GenS6{F0: 1, F1: v45}
                                        _ = v77
                                    case 128:
                                        v78 := v2[("hello" + ("a" + ("a" + ("foo bar" + "a"))))]
                                        _ = v78
                                    }
                                    var v79 int16 = v41
                                    v76 = (^(v43))
                                    _ = v79
                                }
                            case -1:
                                var v80 float32 = v75
                                _ = v80
                            default:
                                v81 := GenS4{F0: 127, F1: 127}
                                _ = v81
                            }
                            var v82 int16 = v76
                            _ = v75
                            _ = v76
                            _ = v82
                        }
                    }
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    var v83 uint = 1
                    if true {
                        v84 := GenS12{F0: 1000, F1: v1, F2: uint8(v83)}
                        fmt.Println(v1, v35)
                        v1 -= v32
                        if false {
                            var v85 uint64 = v32
                            _ = v85
                        }
                        fmt.Println(v32, v32, v35)
                        _ = v84
                    } else {
                        v41 = -32768
                        var v86 uint = v83
                        if true {
                            var v87 float64 = -2.25
                            var v88 float32 = ((1e10 / 1e10) + 1e10)
                            _ = v87
                            _ = v88
                        }
                        _ = v86
                    }
                    _ = v83
                }
                _ = fi37
                _ = v41
                _ = v42
            }
            v35 += v35
            for fi89 := 0; fi89 < 1; fi89++ {
                v35 += 42
                _ = fi89
            }
            _ = fi36
        }
    case "hello":
        v35 = v35
    }
    v90 := GenS9{F0: 2147483647, F1: 582}
    v1 += (v32 % ((v1) | 1))
    for fi91 := 0; fi91 < 16; fi91++ {
        fmt.Println(v1)
        if true {
            var v92 float32 = 0.0
            _ = v92
        } else {
            v93 := GenS3{F0: (^(v35)), F1: ("a" == ("\n\t" + "a")), F2: v1}
            _ = v93
        }
        _ = fi91
    }
    for fi94 := 0; fi94 < 5; fi94++ {
        var v95 bool = false
        _ = fi94
        _ = v95
    }
    v96 := GenS9{F0: 2147483647, F1: 4294967295}
    var v97 uint64 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v98 := GenS13{F0: (1 == 0)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v90
    _ = v96
    _ = v97
    _ = v98
}
