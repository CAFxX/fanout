package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int64; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int; F1 uint64; F2 string }
type GenS14 struct { F0 int8; F1 int; F2 int16 }
func genFn0(p0 string, p1 int32) string {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float32 = 1e10
        var v8 float32 = ((1.5 / (v7 + 0.1)) - -2.25)
        v9 := GenS1{F0: 1, F1: 255, F2: 0}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = 1.5
    if true {
        if true {
            v11 := GenS2{F0: 42, F1: "foo bar", F2: 255}
            switch "foo bar" {
            case "\n\t":
                switch p0 {
                case "foo bar":
                    switch (p0 + ("a" + "")) {
                    case "a":
                        if false {
                            switch int(p1) {
                            case 0:
                                var v12 float32 = -2.25
                                _ = v12
                            case -2147483648:
                                var v13 bool = (!false)
                                _ = v13
                            case 128:
                                var v14 int32 = 0
                                _ = v14
                            }
                            p1 -= 131
                        } else {
                            switch 4294967295 {
                            case 32768:
                                switch "foo bar" {
                                case "foo bar":
                                    v15 := []int{4294967295}
                                    _ = v15
                                case "a":
                                    v16 := GenS2{F0: 850, F1: (p0 + ("\n\t" + (p0 + ""))), F2: 255}
                                    _ = v16
                                case "":
                                    v17 := []int{1, 1, 648}
                                    _ = v17
                                }
                            default:
                                var v18 int = 42
                                _ = v18
                            }
                            fmt.Println(v10)
                        }
                    }
                }
            case "hello":
                p0 = "a"
            case "foo bar":
                if false {
                    v19 := []int{0, 9223372036854775807, 42}
                    v20 := GenS3{F0: 471, F1: 42, F2: 2}
                    _ = v19
                    _ = v20
                }
            default:
                var v21 uint16 = 65535
                _ = v21
            }
            v22 := GenS4{F0: 857, F1: 101}
            _ = v11
            _ = v22
        }
        fmt.Println(p0, p1)
    } else {
        v23 := []int{4294967295, 2147483647, -9223372036854775808, -1}
        if false {
            v24 := GenS5{F0: 171, F1: 1, F2: p1}
            for fi25 := 0; fi25 < 19; fi25++ {
                fmt.Println(v10, p0)
                fmt.Println(p1, p1, fi25)
                _ = fi25
            }
            for fi26 := 0; fi26 < 8; fi26++ {
                v10 -= v10
                fmt.Println(v10, v10)
                _ = fi26
            }
            _ = v24
        }
        v27 := GenS4{F0: 18446744073709551615, F1: 255}
        _ = v23
        _ = v27
    }
    if (!(0 < 128)) {
        for fi28 := 0; fi28 < 9; fi28++ {
            switch (p0 + ((p0 + ((p0 + p0) + p0)) + "a")) {
            case "\n\t":
                v29 := []int{}
                _ = v29
            case "":
                p0 = ((p0 + ("" + ("\n\t" + p0))) + p0)
            default:
                v30 := GenS6{F0: 613, F1: 191}
                _ = v30
            }
            p1 += -422
            p0 = "hello"
            _ = fi28
        }
        if true {
            fmt.Println(v10, p1, v10)
            v10 = 0.1
        }
    } else {
        var v31 uint32 = 522
        var v32 uint16 = 32767
        _ = v31
        _ = v32
    }
    v33 := GenS7{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v33
    return "foo bar"
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 1
    var v3 uint16 = 469
    var v4 int16 = 32767
    v5 := make(map[string]bool)
    v1 = 9223372036854775807
    var v34 uint64 = 1
    v35 := GenS7{F0: 62}
    v36 := GenS3{F0: v1, F1: v34, F2: v34}
    if true {
        switch 4294967295 {
        case 455:
            var v37 uint32 = 1
            _ = v37
        case 2147483647:
            if (!false) {
                var v38 int64 = (v1 + (v1 * v1))
                _ = v38
            } else {
                v39 := GenS5{F0: 128, F1: -9, F2: 1}
                _ = v39
            }
        default:
            var v40 uint16 = (^(v3))
            _ = v40
        }
        v34 += (^(v34))
        var v41 int32 = int32(v2)
        _ = v41
    } else {
        v42 := v5["\n\t"]
        var v43 string = "foo bar"
        v42 = (v34 >= (v34 * (^(v34))))
        _ = v42
        _ = v43
    }
    v44 := GenS1{F0: 42, F1: 76, F2: (v4 & (v4 >> ((42) & 63)))}
    var v45 uint = 2
    var v46 bool = true
    v47 := []int{128, -19}
    v48 := v5["hello"]
    v49 := GenS8{F0: 2147483648}
    switch -9223372036854775808 {
    case 128:
        for fi50 := 0; fi50 < 16; fi50++ {
            v51 := GenS4{F0: (^(v34)), F1: (v3 ^ 2)}
            v34 = (v34 % (((v34 ^ v34)) | 1))
            _ = fi50
            _ = v51
        }
    case 95:
        switch "a" {
        case "hello":
            for fi52 := 0; fi52 < 12; fi52++ {
                v53 := GenS9{F0: (v3 - (v3 + v3)), F1: -128}
                var v54 bool = v48
                if false {
                    switch ("hello" + "hello") {
                    case "":
                        for fi55 := 0; fi55 < 12; fi55++ {
                            if true {
                                var v56 bool = (!v54)
                                _ = v56
                            }
                            v5[""] = (2147483647 != -373)
                            var v57 uint16 = (v3 >> ((uint(v34)) & 63))
                            _ = fi55
                            _ = v57
                        }
                    case "hello":
                        var v58 int32 = 1
                        _ = v58
                    case "foo bar":
                        v2 = v2
                    default:
                        v59 := GenS9{F0: 1, F1: int8(v34)}
                        _ = v59
                    }
                } else {
                    if true {
                        fmt.Println(v48)
                        for fi60 := 0; fi60 < 1; fi60++ {
                            if ((v4 >= (^(v4))) || false) {
                                v61 := GenS10{F0: 42, F1: (^(v34)), F2: 1}
                                _ = v61
                            } else {
                                v2 = 127
                                v62 := GenS11{F0: -9223372036854775808, F1: uint16(v34)}
                                _ = v62
                            }
                            v2 = v2
                            switch "foo bar" {
                            case "":
                                switch fi60 {
                                case -739:
                                    if (-80 >= (-(v2))) {
                                        var v63 int8 = int8(v4)
                                        v64 := GenS12{F0: 0, F1: -180}
                                        v65 := GenS13{F0: int(v2), F1: 0, F2: "\n\t"}
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    } else {
                                        if v46 {
                                            var v66 uint = (^(v45))
                                            _ = v66
                                        }
                                        v54 = false
                                        var v67 int64 = (v1 + v1)
                                        fmt.Println(v1, v54, v67)
                                        v68 := GenS10{F0: 128, F1: v34, F2: 1}
                                        _ = v67
                                        _ = v68
                                    }
                                case 128:
                                    v4 += (v4 - (-(v4)))
                                default:
                                    switch fi52 {
                                    case 1:
                                        v46 = (2147483647 >= 2147483647)
                                    default:
                                        v48 = v54
                                    }
                                }
                            case "a":
                                if (v34 != 2) {
                                    v69 := GenS3{F0: v1, F1: v34, F2: (v34 - (v34 | (^(v34))))}
                                    _ = v69
                                } else {
                                    var v70 float32 = float32(float64(float32(3.14159)))
                                    if (!(!true)) {
                                        if len(v47) > 0 {
                                            v47[0] = (fi52 & fi60)
                                        }
                                        if len(v47) >= 2 {
                                            v47 = v47[1:len(v47)]
                                        }
                                        v48 = (v1 < -9223372036854775808)
                                    }
                                    var v71 int32 = -2147483648
                                    v1 = 42
                                    switch fi52 {
                                    case 128:
                                        fmt.Println(v3)
                                    }
                                    _ = v70
                                    _ = v71
                                }
                            case "foo bar":
                                var v72 float32 = 0.0
                                _ = v72
                            default:
                                var v73 uint64 = v34
                                _ = v73
                            }
                            _ = fi60
                        }
                        v74 := GenS2{F0: (v3 + 533), F1: "\n\t", F2: 255}
                        _ = v74
                    } else {
                        for fi75 := 0; fi75 < 17; fi75++ {
                            fmt.Println(v54, fi52, v45)
                            switch (fi52 * int(v3)) {
                            case -139:
                                var v76 int = 42
                                _ = v76
                            case 32768:
                                fmt.Println(v34, v4)
                                v34 += 423
                            }
                            fmt.Println(v2)
                            _ = fi75
                        }
                        var v77 string = ("" + "hello")
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 1; fi78++ {
                        _ = len(v47)
                        v47 = append(v47, 127)
                        _ = fi78
                    }
                }
                _ = fi52
                _ = v53
                _ = v54
            }
        case "\n\t":
            for fi79 := 0; fi79 < 12; fi79++ {
                var v80 int8 = (-(v2))
                _ = fi79
                _ = v80
            }
        }
    default:
        v34 = 0
    }
    fmt.Println(v3)
    _ = len(v47)
    switch "a" {
    case "hello":
        var v81 int32 = 42
        _ = v81
    case "":
        for fi82 := 0; fi82 < 4; fi82++ {
            for fi83 := 0; fi83 < 10; fi83++ {
                fmt.Println(fi83)
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v1, v46, v4)
                default:
                    v34 += (v34 << ((4294967295) & 63))
                }
                switch "hello" {
                case "":
                    v46 = ((true || v48) && (v48 && false))
                default:
                    var v84 float64 = 0.1
                    _ = v84
                }
                _ = fi83
            }
            var v85 uint64 = (v34 | (^(v34)))
            fmt.Println(v85, v2, v3)
            _ = fi82
            _ = v85
        }
    case "a":
        v1 += (^(v1))
    }
    v86 := GenS14{F0: (v2 / ((v2) | 1)), F1: int(v34), F2: -37}
    if ((((("hello" + "foo bar") + "\n\t") + "hello") + "") <= "a") {
        var v87 float32 = 1e10
        switch "\n\t" {
        case "":
            if ("\n\t" <= ("hello" + "foo bar")) {
                fmt.Println(v46, v45, v4)
            }
        case "foo bar":
            var v88 int16 = 1000
            _ = v88
        case "a":
            _ = len(v5)
        }
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v86
}
