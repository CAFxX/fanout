package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
func genFn0() uint32 {
    v4 := []int{}
    if false {
        v5 := GenS1{F0: -32768, F1: 112}
        var v6 int8 = 53
        _ = v5
        _ = v6
    }
    _ = v4
    return 700
}
func genFn1(p0 int) uint32 {
    p0 -= 32767
    v7 := []int{(p0 ^ 4294967295), p0}
    if false {
        switch p0 {
        case 0:
            if len(v7) > 0 {
                v7[0] = (p0 / ((p0) | 1))
            }
        case 4294967295:
            fmt.Println(p0)
        case 32768:
            for fi8 := 0; fi8 < 16; fi8++ {
                if len(v7) > 0 {
                    v7[0] = (p0 % ((0) | 1))
                }
                var v9 float64 = -2.25
                v10 := GenS2{F0: 0}
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS3{F0: 65535, F1: ("foo bar" + "hello")}
                p0 = 648
                fmt.Println(fi11, fi11)
                _ = fi11
                _ = v12
            }
        }
        p0 = (^(p0))
        if (153 > 128) {
            var v13 float64 = (1.5 / 1e10)
            _ = v13
        }
    }
    p0 -= p0
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    switch "\n\t" {
    case "a":
        if (true && (!((!(219 != 1)) && true))) {
            v14 := v1["\n\t"]
            switch 236 {
            case -601:
                var v15 bool = false
                _ = v15
            case -972:
                switch "a" {
                case "a":
                    v16 := GenS4{F0: 42}
                    _ = v16
                case "foo bar":
                    if false {
                        switch "" {
                        case "\n\t":
                            v14 -= v14
                        case "":
                            v14 = v14
                        case "hello":
                            if true {
                                for fi17 := 0; fi17 < 4; fi17++ {
                                    switch "hello" {
                                    case "foo bar":
                                        var v18 float64 = 1.5
                                        _ = v18
                                    case "\n\t":
                                        v19 := GenS2{F0: 0}
                                        _ = v19
                                    case "hello":
                                        var v20 int16 = -32768
                                        _ = v20
                                    }
                                    if (!true) {
                                        switch "hello" {
                                        case "foo bar":
                                            switch "hello" {
                                            case "foo bar":
                                                var v21 int8 = 0
                                                _ = v21
                                            case "a":
                                                fmt.Println(v14, v14)
                                            default:
                                                for fi22 := 0; fi22 < 17; fi22++ {
                                                    v23 := v2[(^(fi22))]
                                                    _ = fi22
                                                    _ = v23
                                                }
                                            }
                                        case "\n\t":
                                            v24 := GenS4{F0: -1}
                                            _ = v24
                                        case "hello":
                                            v14 -= v14
                                        default:
                                            v25 := []int{fi17}
                                            _ = v25
                                        }
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v14, fi17)
                                        case "foo bar":
                                            v14 = genFn0()
                                        }
                                        switch (fi17 - fi17) {
                                        case -9223372036854775808:
                                            v14 = (v14 - genFn0())
                                        case -745:
                                            fmt.Println(v14)
                                        case 32768:
                                            if true {
                                                if true {
                                                    var v26 int = fi17
                                                    for fi27 := 0; fi27 < 13; fi27++ {
                                                        v26 = fi27
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v28 int16 = 262
                                                            v29 := []int{-629, v26}
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        _ = len(v3)
                                                        _ = fi27
                                                    }
                                                    for fi30 := 0; fi30 < 2; fi30++ {
                                                        v31 := GenS1{F0: 1, F1: 934}
                                                        _ = fi30
                                                        _ = v31
                                                    }
                                                    _ = v26
                                                }
                                            } else {
                                                switch fi17 {
                                                case 0:
                                                    v32 := GenS3{F0: 4294967295, F1: ""}
                                                    _ = v32
                                                }
                                                for fi33 := 0; fi33 < 8; fi33++ {
                                                    v14 -= 2147483648
                                                    v14 = (v14 - 65535)
                                                    fmt.Println(fi33, fi17, v14)
                                                    _ = fi33
                                                }
                                            }
                                        default:
                                            v14 -= v14
                                        }
                                        fmt.Println(fi17, fi17)
                                    }
                                    _ = fi17
                                }
                            }
                        default:
                            if true {
                                if (!((0.1 != float64((float32(((float64(1.5) + -2.25) - 3.14159)) * float32((float64(-2.25) + 3.14159))))) && false)) {
                                    if false {
                                        var v34 uint16 = 837
                                        _ = v34
                                    }
                                    _ = len(v3)
                                    switch -620 {
                                    case 590:
                                        if true {
                                            v35 := []int{-2147483648, 127}
                                            v36 := GenS1{F0: 0, F1: uint64(v14)}
                                            if (!(-32768 < -32768)) {
                                                _ = len(v3)
                                            }
                                            _ = v35
                                            _ = v36
                                        }
                                    default:
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            for fi38 := 0; fi38 < 12; fi38++ {
                                                fmt.Println(fi37, v14, fi37)
                                                v14 += v14
                                                _ = fi38
                                            }
                                            v14 -= v14
                                            _ = fi37
                                        }
                                    }
                                } else {
                                    var v39 uint16 = 32767
                                    var v40 int64 = -867
                                    if false {
                                        if false {
                                            var v41 uint16 = v39
                                            switch -219 {
                                            case 2147483647:
                                                fmt.Println(v39, v39, v41)
                                            }
                                            _ = v41
                                        } else {
                                            fmt.Println(v39)
                                            var v42 uint = uint(v39)
                                            if (false && (!false)) {
                                                v42 = 2
                                                v43 := []int{0, -228}
                                                _ = v43
                                            } else {
                                                v42 = (v42 ^ 0)
                                                fmt.Println(v14, v40, v42)
                                                switch "" {
                                                case "\n\t":
                                                    var v44 bool = false
                                                    _ = v44
                                                case "":
                                                    v1["hello"] = (v14 & (v14 - 65535))
                                                default:
                                                    if (!(42 <= 42)) {
                                                        v42 -= 255
                                                        v39 -= (^(v39))
                                                        if false {
                                                            var v45 int64 = -33
                                                            var v46 uint = (v42 * v42)
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            v47 := GenS1{F0: -32768, F1: 2}
                                                            _ = v47
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v42
                                        }
                                    } else {
                                        v14 = 474
                                        v48 := v2[128]
                                        var v49 float64 = 3.14159
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = v39
                                    _ = v40
                                }
                            }
                        }
                        for fi50 := 0; fi50 < 12; fi50++ {
                            var v51 uint16 = 255
                            _ = fi50
                            _ = v51
                        }
                        if true {
                            if (1 != 81) {
                                v14 -= v14
                                var v52 uint = 0
                                _ = v52
                            }
                            v14 += v14
                            v14 = v14
                        } else {
                            for fi53 := 0; fi53 < 2; fi53++ {
                                if true {
                                    var v54 int16 = 42
                                    v55 := GenS4{F0: -21}
                                    if ("hello" >= "") {
                                        for fi56 := 0; fi56 < 6; fi56++ {
                                            var v57 float64 = (float64(0.1) / 1e10)
                                            _ = fi56
                                            _ = v57
                                        }
                                        v54 -= (^(v54))
                                    } else {
                                        for fi58 := 0; fi58 < 19; fi58++ {
                                            v59 := v3["hello"]
                                            fmt.Println(v54, v59)
                                            var v60 int16 = (v54 | (^(v54)))
                                            v60 += 0
                                            switch fi53 {
                                            case -2147483648:
                                                v1["hello"] = v14
                                            case 0:
                                                if v59 {
                                                    v61 := GenS5{F0: 42}
                                                    v62 := GenS6{F0: genFn1(32768)}
                                                    _ = v61
                                                    _ = v62
                                                }
                                            default:
                                                delete(v1, "\n\t")
                                            }
                                            _ = fi58
                                            _ = v59
                                            _ = v60
                                        }
                                        v14 = v14
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                var v63 int32 = -1
                                switch -1 {
                                case 4294967295:
                                    for fi64 := 0; fi64 < 2; fi64++ {
                                        fmt.Println(v63, fi64, v63)
                                        _ = fi64
                                    }
                                case 2147483647:
                                    v63 += (^(v63))
                                default:
                                    var v65 int16 = 1000
                                    _ = v65
                                }
                                _ = fi53
                                _ = v63
                            }
                        }
                    } else {
                        var v66 uint = 2
                        _ = v66
                    }
                default:
                    v14 -= genFn1(108)
                }
            }
            v67 := GenS7{F0: ("" + "")}
            _ = v14
            _ = v67
        }
    case "\n\t":
        v68 := GenS5{F0: -1}
        _ = v68
    default:
        v1["hello"] = genFn0()
    }
    if false {
        v69 := []int{529, 128, -2147483648}
        var v70 float32 = 1e10
        switch (("\n\t" + "hello") + ("foo bar" + "\n\t")) {
        case "a":
            v70 -= v70
        case "":
            v70 = float32(1.5)
        case "foo bar":
            for fi71 := 0; fi71 < 11; fi71++ {
                switch -581 {
                case -774:
                    fmt.Println(fi71)
                case -2147483648:
                    var v72 string = ("hello" + "")
                    _ = v72
                default:
                    v70 += v70
                }
                _ = fi71
            }
        default:
            v70 -= float32((3.14159 - 0.1))
        }
        _ = v69
        _ = v70
    }
    var v73 int8 = 127
    var v74 uint32 = 269
    v3[""] = false
    v3[""] = false
    v74 = (v74 * (^(v74)))
    for fi75 := 0; fi75 < 20; fi75++ {
        v76 := GenS5{F0: 42}
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 1; fi77++ {
        var v78 int32 = -594
        _ = fi77
        _ = v78
    }
    fmt.Println(v74)
    for fi79 := 0; fi79 < 7; fi79++ {
        v74 = v74
        _ = fi79
    }
    var v80 uint16 = 2
    if true {
        if false {
            var v81 uint8 = 42
            v73 = v73
            var v82 uint8 = (^(v81))
            _ = v81
            _ = v82
        } else {
            if (v73 >= (v73 - int8(v74))) {
                var v83 uint64 = 925
                for fi84 := 0; fi84 < 5; fi84++ {
                    var v85 uint = 720
                    _ = fi84
                    _ = v85
                }
                _ = v83
            }
            _ = len(v2)
        }
    } else {
        v73 -= (^(v73))
        if true {
            v73 = v73
            var v86 uint32 = 485
            var v87 uint = 223
            _ = v86
            _ = v87
        } else {
            v80 = v80
            var v88 uint = 201
            _ = v88
        }
        v80 = v80
    }
    v74 = v74
    fmt.Println(v74)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v80
}
