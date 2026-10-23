package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8, p1 int64) string {
    p1 += p1
    var v4 int16 = -32768
    _ = p0
    _ = p1
    _ = v4
    return "a"
}
func genFn1(p0 uint8) uint16 {
    var v5 bool = ((true || true) || false)
    v5 = true
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: -1, F1: 96, F2: 2147483647}
        var v8 bool = false
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: (fi9 + (fi6 + 402)), F1: 123, F2: 42}
            var v11 uint16 = 941
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 32767 {
    case 913:
        v12 := GenS1{F0: 127, F1: -70, F2: 1}
        _ = v12
    case -2147483648:
        if (("foo bar" + "foo bar") >= ("" + "\n\t")) {
            var v13 int64 = int64(p0)
            v14 := GenS3{F0: "a", F1: 4294967295, F2: (p0 + (^(p0)))}
            var v15 uint64 = 0
            _ = v13
            _ = v14
            _ = v15
        } else {
            if (!(0.1 == float64(0.1))) {
                p0 = (p0 * p0)
            } else {
                switch "" {
                case "":
                    if (107 >= 2) {
                        v16 := []int{}
                        v5 = (p0 < 1)
                        _ = v16
                    } else {
                        if true {
                            for fi17 := 0; fi17 < 3; fi17++ {
                                switch "a" {
                                case "a":
                                    if true {
                                        p0 = p0
                                        v18 := GenS3{F0: "\n\t", F1: 0, F2: (p0 / ((255) | 1))}
                                        _ = v18
                                    } else {
                                        var v19 int16 = -32768
                                        var v20 float64 = 1.5
                                        _ = v19
                                        _ = v20
                                    }
                                case "hello":
                                    for fi21 := 0; fi21 < 8; fi21++ {
                                        v22 := []int{fi21, fi17}
                                        _ = fi21
                                        _ = v22
                                    }
                                case "":
                                    v23 := []int{(fi17 % ((fi17) | 1)), 0}
                                    _ = v23
                                }
                                switch (("" + "") + "a") {
                                case "hello":
                                    v5 = (-77 < -128)
                                case "\n\t":
                                    p0 = (p0 | 127)
                                case "a":
                                    p0 = (p0 << ((255) & 63))
                                }
                                _ = fi17
                            }
                            if (88 < 32767) {
                                p0 = p0
                                var v24 uint = 745
                                _ = v24
                            } else {
                                switch 555 {
                                case 44:
                                    var v25 uint16 = 65535
                                    _ = v25
                                case -518:
                                    v26 := []int{281}
                                    _ = v26
                                case 42:
                                    v27 := GenS3{F0: "a", F1: 2, F2: p0}
                                    _ = v27
                                default:
                                    fmt.Println(p0, v5)
                                }
                            }
                            switch "a" {
                            case "foo bar":
                                v28 := []int{1, 2147483647, 785}
                                _ = v28
                            case "":
                                var v29 string = "a"
                                _ = v29
                            case "a":
                                switch 9223372036854775807 {
                                case 0:
                                    v30 := []int{}
                                    _ = v30
                                }
                            }
                        } else {
                            fmt.Println(p0, p0)
                        }
                        var v31 uint64 = 735
                        v32 := []int{}
                        _ = v31
                        _ = v32
                    }
                case "a":
                    if v5 {
                        fmt.Println(v5)
                    } else {
                        p0 -= 2
                        if (!v5) {
                            fmt.Println(v5, v5)
                            p0 = 128
                        }
                        switch "a" {
                        case "\n\t":
                            p0 = 128
                        case "":
                            p0 = p0
                        case "foo bar":
                            p0 += 128
                        }
                    }
                case "hello":
                    for fi33 := 0; fi33 < 20; fi33++ {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            fmt.Println(v5, fi33, v5)
                            _ = fi34
                        }
                        v5 = true
                        if (!(128 > fi33)) {
                            v35 := []int{42}
                            switch "\n\t" {
                            case "hello":
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                    switch -165 {
                                    case -721:
                                        var v37 bool = false
                                        _ = v37
                                    default:
                                        fmt.Println(fi36)
                                    }
                                    _ = fi36
                                }
                            case "a":
                                var v38 uint8 = p0
                                _ = v38
                            case "":
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    v5 = true
                                    switch (("foo bar" + ("" + "\n\t")) + "foo bar") {
                                    case "foo bar":
                                        v35 = append(v35, (fi33 << ((42) & 63)))
                                    case "":
                                        var v40 uint64 = 378
                                        _ = v40
                                    case "a":
                                        v5 = v5
                                    }
                                    if true {
                                        var v41 bool = true
                                        v42 := GenS2{F0: fi39, F1: 1, F2: 670}
                                        v35 = append(v35, (fi39 >> ((308) & 63)))
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if v5 {
                                            _ = len(v35)
                                            if false {
                                                fmt.Println(fi39, p0, p0)
                                                v43 := GenS2{F0: fi33, F1: 1, F2: 809}
                                                _ = v43
                                            } else {
                                                switch fi33 {
                                                case -333:
                                                    var v44 string = ""
                                                    if (1 == 100) {
                                                        switch "\n\t" {
                                                        case "":
                                                            v44 = "\n\t"
                                                        case "hello":
                                                            if true {
                                                                var v45 float64 = 1e10
                                                                fmt.Println(v44)
                                                                _ = v45
                                                            } else {
                                                                fmt.Println(p0, fi33, p0)
                                                            }
                                                        case "a":
                                                            if (516 < 42) {
                                                                v5 = v5
                                                            }
                                                            v46 := GenS3{F0: v44, F1: 1, F2: (p0 & (^(p0)))}
                                                            _ = v46
                                                        default:
                                                            var v47 int8 = 127
                                                            _ = v47
                                                        }
                                                        var v48 float64 = float64(-2.25)
                                                        var v49 uint64 = 4294967295
                                                        v35 = append(v35, fi39)
                                                        var v50 uint64 = (v49 ^ 2)
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS4{F0: 2}
                                                        _ = v51
                                                    }
                                                    fmt.Println(p0, fi39)
                                                    v44 = "\n\t"
                                                    _ = v44
                                                case 32768:
                                                    v5 = false
                                                case 732:
                                                    if len(v35) >= 2 {
                                                        v35 = v35[1:len(v35)]
                                                    }
                                                    var v52 string = "a"
                                                    _ = v52
                                                }
                                                var v53 float32 = (3.14159 - 1e10)
                                                var v54 int8 = 42
                                                _ = len(v35)
                                                var v55 int32 = 1
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    }
                                    _ = fi39
                                }
                            default:
                                switch (fi33 * (fi33 - fi33)) {
                                case 0:
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                case 785:
                                    p0 = 2
                                }
                            }
                            _ = v35
                        } else {
                            v5 = v5
                            var v56 int8 = -1
                            v56 = 100
                            _ = v56
                        }
                        _ = fi33
                    }
                }
            }
            v5 = v5
            switch "a" {
            case "a":
                for fi57 := 0; fi57 < 17; fi57++ {
                    if true {
                        p0 = p0
                        v58 := GenS5{F0: 1}
                        if v5 {
                            v5 = v5
                            v59 := []int{(fi57 >> ((393) & 63)), fi57, (^(fi57))}
                            _ = v59
                        } else {
                            fmt.Println(fi57, fi57)
                            fmt.Println(v5, v5)
                            for fi60 := 0; fi60 < 9; fi60++ {
                                p0 = p0
                                p0 = (^(p0))
                                _ = fi60
                            }
                        }
                        _ = v58
                    }
                    p0 -= (^(p0))
                    v61 := []int{fi57, fi57, 127}
                    _ = fi57
                    _ = v61
                }
            case "hello":
                fmt.Println(p0)
            default:
                v62 := GenS4{F0: 74}
                _ = v62
            }
        }
    case 4294967295:
        switch 128 {
        case 32768:
            fmt.Println(v5, p0)
        default:
            for fi63 := 0; fi63 < 4; fi63++ {
                p0 = (p0 & (p0 | (^(p0))))
                v64 := []int{(-(fi63))}
                var v65 uint32 = 4294967295
                _ = fi63
                _ = v64
                _ = v65
            }
        }
    }
    _ = p0
    _ = v5
    return 213
}

func main() {
    v1 := []uint32{631, 1, 2}
    var v2 int32 = -470
    var v3 int64 = 42
    v1 = append(v1, 2)
    v66 := GenS2{F0: 42, F1: -32768, F2: 1}
    var v67 uint8 = 1
    v67 = v67
    v68 := GenS3{F0: (("\n\t" + "a") + "foo bar"), F1: 768, F2: (v67 ^ 128)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    v69 := GenS3{F0: "a", F1: 4294967295, F2: v67}
    var v70 string = "a"
    switch "\n\t" {
    case "\n\t":
        switch 0 {
        case -2147483648:
            v3 += 9223372036854775807
        default:
            var v71 int = 2147483647
            _ = v71
        }
    }
    var v72 int = -913
    v67 -= (^(v67))
    if true {
        v2 = v2
        v2 -= v2
        v3 = 42
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        v67 -= 216
    } else {
        var v73 bool = true
        v3 = v3
        v70 = "a"
        _ = v73
    }
    if ("" > "a") {
        var v74 uint16 = 65535
        v2 = -2147483648
        var v75 int8 = 15
        _ = v74
        _ = v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v72
}
