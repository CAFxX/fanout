package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 uint8, p1 uint16) uint16 {
    fmt.Println(p0, p0)
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 12; fi6++ {
        switch fi6 {
        case 127:
            p0 -= (p0 - (p0 >> ((18446744073709551615) & 63)))
        case 42:
            v7 := GenS1{F0: -1, F1: 42, F2: 97}
            _ = v7
        case 0:
            var v8 int8 = -128
            _ = v8
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(p0)
            fmt.Println(fi9, p1, fi6)
            v5 -= 0.0
            var v10 uint8 = p0
            _ = fi9
            _ = v10
        }
        v11 := []int{32767, 32767}
        _ = fi6
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := []string{((v1 + v1) + (("foo bar" + v1) + (v1 + v1))), ("a" + v1), v1}
    v4 := make(map[uint32]int)
    switch 446 {
    case -2147483648:
        switch v1 {
        case "hello":
            var v12 string = v1
            _ = v12
        case "\n\t":
            var v13 int32 = 1
            _ = v13
        case "":
            for fi14 := 0; fi14 < 17; fi14++ {
                switch 128 {
                case -686:
                    fmt.Println(fi14, v1)
                case 0:
                    var v15 float32 = (((((float32(1e10) - float32(0.0)) - 1e10) + 1.5) * (3.14159 - (1e10 - (1.5 + ((0.0 - 1.5) * 1.5))))) * (1e10 / float32((-2.25 - 3.14159))))
                    _ = v15
                }
                if (-2.25 != (1e10 + 0.1)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v16 int16 = 0
                    _ = v16
                }
                _ = fi14
            }
        default:
            switch v1 {
            case "hello":
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 string = ("" + "")
                    v1 = v18
                    v19 := v4[2]
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "":
                if true {
                    if (2 <= 34) {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            for fi21 := 0; fi21 < 19; fi21++ {
                                v22 := GenS1{F0: 680, F1: 1, F2: 18446744073709551615}
                                _ = fi21
                                _ = v22
                            }
                            v1 = ("" + "a")
                            _ = fi20
                        }
                    } else {
                        if ((!true) || (!true)) {
                            v1 = v1
                            var v23 bool = false
                            _ = v23
                        } else {
                            var v24 int16 = 1
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 17; fi25++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = v1
                            _ = fi25
                        }
                        switch -559 {
                        case 127:
                            fmt.Println(v1)
                        case 42:
                            v26 := GenS1{F0: -1, F1: 65535, F2: 2}
                            _ = v26
                        }
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            var v29 string = ""
                            var v30 int32 = 42
                            switch (v29 + ("\n\t" + "")) {
                            case "\n\t":
                                v30 += v30
                                v30 = int32(fi27)
                                if (2147483648 < 734) {
                                    v29 = (v1 + ("foo bar" + ("hello" + v1)))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                switch (fi28 ^ fi27) {
                                case 127:
                                    v29 = ("hello" + ("\n\t" + ("hello" + v1)))
                                }
                            case "hello":
                                v29 = ((("hello" + (("foo bar" + "") + "hello")) + v29) + v29)
                                if (9223372036854775807 < 1) {
                                    fmt.Println(fi28, v1, fi28)
                                    _ = len(v3)
                                } else {
                                    v29 = "a"
                                }
                                v1 = v1
                            case "a":
                                if (!(false && true)) {
                                    v4[233] = (fi28 + fi28)
                                } else {
                                    fmt.Println(v1, v1, v1)
                                }
                            default:
                                var v31 int = fi28
                                _ = v31
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(v1, fi27)
                        _ = fi27
                    }
                }
            case "foo bar":
                v32 := GenS2{F0: 106, F1: 0, F2: (!(65535 <= 42))}
                _ = v32
            }
        }
    case 32768:
        v33 := GenS3{F0: 42}
        _ = v33
    default:
        v1 = ((v1 + v1) + ((v1 + v1) + v1))
    }
    v3 = append(v3, "foo bar")
    fmt.Println(v1)
    switch 32768 {
    case 0:
        var v34 int8 = 100
        _ = v34
    case -1:
        if true {
            var v35 int32 = 0
            v35 = v35
            _ = v35
        } else {
            if (65535 == 4294967295) {
                var v36 float64 = 1.5
                if ((-2147483648 == -2147483648) || (!(0 < 9223372036854775807))) {
                    switch ("a" + (v1 + (("\n\t" + "hello") + "\n\t"))) {
                    case "foo bar":
                        if false {
                            var v37 string = ((v1 + (v1 + "hello")) + ("\n\t" + (v1 + (v1 + ("\n\t" + (v1 + "a"))))))
                            _ = v37
                        }
                    case "\n\t":
                        var v38 uint8 = 255
                        _ = v38
                    }
                } else {
                    var v39 bool = true
                    _ = v39
                }
                _ = v36
            }
            v40 := v2[953]
            _ = v40
        }
    default:
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 uint16 = 32767
            if (!(1e10 >= float64(-2.25))) {
                if (-1 > -711) {
                    v43 := GenS4{F0: 184, F1: 255}
                    _ = len(v3)
                    var v44 int32 = 1
                    _ = v43
                    _ = v44
                }
                fmt.Println(v1, v42, fi41)
                for fi45 := 0; fi45 < 14; fi45++ {
                    if (907 != 42) {
                        if false {
                            v1 = "foo bar"
                            switch (fi45 - 9223372036854775807) {
                            case 127:
                                fmt.Println(fi45)
                            case 681:
                                var v46 float64 = 3.14159
                                _ = v46
                            }
                        } else {
                            var v47 uint = 577
                            switch (v1 + v1) {
                            case "\n\t":
                                v48 := GenS5{F0: 467, F1: -128}
                                _ = v48
                            case "hello":
                                fmt.Println(fi41, v1, v42)
                                v49 := GenS4{F0: 1, F1: (v47 >> ((v47) & 63))}
                                v50 := GenS6{F0: (v42 - v42)}
                                _ = v49
                                _ = v50
                            }
                            _ = v47
                        }
                        fmt.Println(fi45, fi45, fi41)
                    }
                    if true {
                        var v51 float64 = -2.25
                        switch (("\n\t" + (("foo bar" + v1) + ("" + "hello"))) + v1) {
                        case "\n\t":
                            if true {
                                switch v1 {
                                case "\n\t":
                                    v42 = 918
                                }
                            }
                        case "a":
                            var v52 uint32 = 2147483648
                            _ = v52
                        case "":
                            v53 := GenS7{F0: 0}
                            fmt.Println(fi41, v1)
                            _ = v53
                        default:
                            var v54 uint8 = 1
                            _ = v54
                        }
                        _ = v51
                    }
                    _ = fi45
                }
            } else {
                for fi55 := 0; fi55 < 10; fi55++ {
                    switch v1 {
                    case "foo bar":
                        if (2 != uint32(v42)) {
                            if false {
                                if true {
                                    v56 := GenS7{F0: -9223372036854775808}
                                    var v57 float64 = (3.14159 * (1e10 * float64(1e10)))
                                    fmt.Println(v1, v57, fi55)
                                    _ = v56
                                    _ = v57
                                }
                            } else {
                                v42 = genFn0(128, v42)
                            }
                            var v58 int16 = 757
                            _ = v58
                        }
                    case "":
                        switch fi41 {
                        case -640:
                            v1 = v1
                        }
                    case "a":
                        if true {
                            v1 = v1
                            var v59 uint32 = 2147483648
                            var v60 int = (-(fi41))
                            _ = v59
                            _ = v60
                        } else {
                            var v61 uint8 = 1
                            var v62 int64 = 0
                            _ = v61
                            _ = v62
                        }
                    }
                    var v63 uint32 = uint32(v42)
                    _ = fi55
                    _ = v63
                }
                v42 -= (v42 + v42)
            }
            v1 = (v1 + (v1 + ("hello" + (("foo bar" + "a") + ""))))
            _ = fi41
            _ = v42
        }
    }
    var v64 uint32 = 64
    switch 0 {
    case 1:
        switch ((v1 + v1) + "\n\t") {
        case "":
            var v65 string = (v1 + (v1 + "\n\t"))
            _ = v65
        default:
            if true {
                v66 := GenS2{F0: v64, F1: 127, F2: true}
                _ = v66
            }
        }
    case -186:
        for fi67 := 0; fi67 < 10; fi67++ {
            switch v1 {
            case "a":
                for fi68 := 0; fi68 < 18; fi68++ {
                    if true {
                        if true {
                            if true {
                                var v69 int16 = 32767
                                if ((((0.1 + float64(3.14159)) * (3.14159 / ((float64((3.14159 / (float32(1.5) + -2.25))) - (0.1 - float64(0.0))) + -2.25))) / (0.1 - 1e10)) == 3.14159) {
                                    fmt.Println(v1, v69)
                                    v70 := GenS5{F0: int(v69), F1: 1}
                                    v71 := GenS8{F0: int(v64), F1: 1, F2: 2147483647}
                                    _ = v70
                                    _ = v71
                                }
                                var v72 uint16 = 640
                                _ = v69
                                _ = v72
                            } else {
                                v2[(fi67 ^ -365)] = v64
                                v3 = append(v3, v1)
                                var v73 int32 = -2147483648
                                _ = v73
                            }
                            v74 := GenS3{F0: 564}
                            var v75 int64 = 902
                            var v76 int = -2147483648
                            _ = len(v3)
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            v1 = v1
                            fmt.Println(fi68)
                        }
                    }
                    var v77 int64 = -1
                    _ = fi68
                    _ = v77
                }
            case "foo bar":
                fmt.Println(v64)
            default:
                v78 := v2[fi67]
                _ = v78
            }
            var v79 uint8 = 127
            if true {
                v80 := GenS8{F0: fi67, F1: 4294967295, F2: -2147483648}
                var v81 int8 = 0
                v64 = v64
                _ = v80
                _ = v81
            }
            _ = fi67
            _ = v79
        }
    default:
        fmt.Println(v64, v64)
    }
    if false {
        v1 = "a"
        var v82 uint64 = 2
        var v83 string = ((v1 + (v1 + ((v1 + ("hello" + "\n\t")) + (v1 + "")))) + "")
        _ = v82
        _ = v83
    } else {
        v1 = (v1 + v1)
        fmt.Println(v64, v64)
    }
    fmt.Println(v1)
    var v84 float64 = 0.0
    v84 += ((1e10 / (-2.25 + v84)) - v84)
    var v85 uint8 = 255
    v85 += v85
    switch v1 {
    case "foo bar":
        for fi86 := 0; fi86 < 18; fi86++ {
            fmt.Println(v1)
            fmt.Println(v1, v1, v64)
            _ = fi86
        }
    }
    var v87 int = 4294967295
    if (0 >= 525) {
        var v88 uint32 = v64
        v88 -= 4294967295
        _ = v88
    }
    v64 = (^(v64))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v84
    _ = v85
    _ = v87
}
