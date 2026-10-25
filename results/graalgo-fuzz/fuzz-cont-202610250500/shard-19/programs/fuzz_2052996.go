package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint32; F1 string; F2 uint8 }
func genFn0(p0 int16) bool {
    switch ("hello" + "a") {
    case "foo bar":
        v7 := []int{}
        _ = v7
    case "":
        var v8 uint32 = 0
        _ = v8
    }
    switch -1 {
    case 32768:
        var v9 uint64 = 0
        _ = v9
    case 2147483647:
        v10 := GenS1{F0: -1, F1: 127, F2: p0}
        _ = v10
    case 42:
        var v11 int32 = 1
        _ = v11
    default:
        fmt.Println(p0)
    }
    var v12 int16 = 42
    if true {
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(fi13, v12, v12)
            v14 := GenS1{F0: 0, F1: 42, F2: (v12 & v12)}
            _ = fi13
            _ = v14
        }
        var v15 float32 = 0.1
        v16 := []int{-400, 2147483647}
        _ = v15
        _ = v16
    }
    if false {
        v17 := []int{2147483647, -2147483648}
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                fmt.Println(fi18, v12, v12)
                fmt.Println(p0, v12)
            } else {
                fmt.Println(fi18, fi18)
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                if false {
                    var v20 uint = 255
                    v20 += 2
                    _ = v20
                }
                var v21 uint64 = 4294967295
                v21 = (v21 << ((1) & 63))
                switch ("hello" + ("hello" + "")) {
                case "a":
                    var v22 float64 = ((1.5 + 1e10) + 1e10)
                    _ = v22
                case "hello":
                    switch fi19 {
                    case -2147483648:
                        if false {
                            var v23 uint64 = (v21 + 4294967295)
                            if (!true) {
                                v24 := GenS1{F0: 1, F1: 42, F2: v12}
                                _ = v24
                            } else {
                                fmt.Println(fi18, fi18, p0)
                            }
                            v23 = (^(v21))
                            fmt.Println(v12, p0, v21)
                            var v25 int16 = (p0 & (p0 ^ p0))
                            _ = v23
                            _ = v25
                        } else {
                            var v26 int32 = 42
                            switch "hello" {
                            case "a":
                                v17 = append(v17, (fi18 >> ((777) & 63)))
                            case "foo bar":
                                if ((919 <= 32767) && true) {
                                    if (!true) {
                                        fmt.Println(fi18)
                                        switch (fi18 + fi19) {
                                        case 246:
                                            fmt.Println(v12, v26, v21)
                                        case 0:
                                            fmt.Println(p0)
                                        }
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        if false {
                                            v12 += 32767
                                            v27 := GenS1{F0: 9223372036854775807, F1: 42, F2: v12}
                                            v12 = -1
                                            _ = v27
                                        }
                                        var v28 int = fi18
                                        v17 = append(v17, -320)
                                        v29 := GenS1{F0: 1, F1: 255, F2: (p0 ^ v12)}
                                        v30 := GenS2{F0: fi18}
                                        if true {
                                            if false {
                                                p0 += int16(v26)
                                                v21 -= (v21 + v21)
                                                _ = len(v17)
                                            }
                                        } else {
                                            v31 := GenS2{F0: -527}
                                            v26 -= 632
                                            _ = v31
                                        }
                                        fmt.Println(v21, p0)
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            default:
                                switch (fi18 << ((2) & 63)) {
                                case -662:
                                    if true {
                                        fmt.Println(v21, v26)
                                        v26 -= v26
                                    }
                                case 140:
                                    v21 += 42
                                }
                            }
                            _ = v26
                        }
                    case 4294967295:
                        var v32 float32 = -2.25
                        _ = v32
                    case 128:
                        if ((!(!(!true))) && (false || (true || true))) {
                            v12 = v12
                            var v33 uint16 = 32767
                            _ = v33
                        }
                    default:
                        switch fi19 {
                        case 32768:
                            if (true || ((v21 / ((42) | 1)) == (v21 - 230))) {
                                var v34 int8 = 1
                                v35 := GenS3{F0: v21, F1: -2147483648, F2: 1}
                                var v36 uint16 = 1
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                if true {
                                    v37 := GenS2{F0: fi19}
                                    _ = v37
                                }
                            }
                        case -850:
                            p0 = (v12 - -1)
                        case 32767:
                            v38 := GenS3{F0: v21, F1: -886, F2: 65535}
                            _ = v38
                        default:
                            v21 = v21
                        }
                    }
                case "foo bar":
                    v21 -= v21
                }
                _ = fi19
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    } else {
        v39 := []int{28, 128}
        fmt.Println(p0)
        _ = v39
    }
    _ = p0
    _ = v12
    return (true && false)
}
func genFn1(p0 string, p1 int8, p2 int64) int16 {
    var v40 int64 = 692
    var v41 int32 = -2147483648
    switch 42 {
    case 0:
        switch p0 {
        case "foo bar":
            v42 := GenS4{F0: (v41 * v41)}
            _ = v42
        case "":
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 uint16 = 0
                _ = fi43
                _ = v44
            }
        default:
            var v45 uint64 = 978
            _ = v45
        }
    case 42:
        v46 := []int{-583, 127, 611}
        _ = v46
    default:
        v40 = (p2 + v40)
    }
    var v47 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    _ = v41
    _ = v47
    return (v47 << ((2) & 63))
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768}
    var v3 uint32 = 42
    v4 := []uint32{v3, 1, (v3 - v3), (v1 << ((2) & 63))}
    v5 := []int{9223372036854775807, 1, 1, 102}
    v6 := make(map[string]bool)
    var v48 bool = true
    var v49 uint = 776
    if v48 {
        v3 -= (v1 * (v3 / (((^(v1))) | 1)))
        switch "foo bar" {
        case "hello":
            _ = len(v5)
        case "\n\t":
            var v50 string = "foo bar"
            _ = v50
        default:
            v3 = (v3 - 42)
        }
        var v51 uint8 = 154
        _ = v51
    }
    v52 := GenS4{F0: -72}
    for fi53 := 0; fi53 < 4; fi53++ {
        v48 = v48
        _ = len(v2)
        v3 = v3
        _ = fi53
    }
    v54 := GenS5{F0: (("hello" + "foo bar") + "foo bar"), F1: 4294967295}
    if false {
        var v55 int8 = -1
        var v56 int16 = 1000
        fmt.Println(v1)
        _ = v55
        _ = v56
    } else {
        var v57 uint16 = 2
        var v58 int32 = 42
        switch 1 {
        case 168:
            for fi59 := 0; fi59 < 1; fi59++ {
                if (127 >= 95) {
                    if v48 {
                        v6[("\n\t" + "\n\t")] = false
                        if v48 {
                            v60 := GenS6{F0: (v57 + v57), F1: int32(fi59)}
                            var v61 string = ""
                            _ = v60
                            _ = v61
                        }
                    } else {
                        fmt.Println(v1)
                    }
                    for fi62 := 0; fi62 < 10; fi62++ {
                        v63 := GenS4{F0: v58}
                        var v64 int32 = (v58 * v58)
                        var v65 float32 = 0.0
                        _ = len(v6)
                        switch ("" + "hello") {
                        case "a":
                            v58 += (v64 & (^(v64)))
                        case "foo bar":
                            v58 = (v64 + (v64 & (-(v58))))
                            v48 = (-32768 > genFn1((((("hello" + "foo bar") + "foo bar") + "") + "hello"), -128, -827))
                            var v66 int16 = -210
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                            _ = v66
                        case "":
                            fmt.Println(v1)
                        default:
                            switch (fi59 << (((v49 >> (((v49 + (^(v49)))) & 63))) & 63)) {
                            case 4294967295:
                                v67 := GenS7{F0: v1, F1: ("\n\t" + "foo bar"), F2: 2}
                                _ = v67
                            case -155:
                                switch fi59 {
                                case -2147483648:
                                    v68 := GenS6{F0: (v57 ^ (^(v57))), F1: v64}
                                    _ = v68
                                case 382:
                                    switch (fi62 - (-(fi59))) {
                                    case -1:
                                        v65 = (v65 + (0.0 / v65))
                                    case 149:
                                        fmt.Println(v48, fi59)
                                    }
                                case 1:
                                    var v69 int16 = -832
                                    _ = v69
                                default:
                                    var v70 uint32 = v1
                                    _ = v70
                                }
                                var v71 int32 = -1
                                _ = v71
                            case 240:
                                delete(v6, (("hello" + "hello") + "\n\t"))
                            }
                            var v72 float64 = float64((0.0 + float32(1.5)))
                            var v73 uint8 = 249
                            _ = v72
                            _ = v73
                        }
                        _ = fi62
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                } else {
                    v74 := GenS3{F0: 914, F1: (-(v58)), F2: (v57 & 120)}
                    for fi75 := 0; fi75 < 3; fi75++ {
                        v5 = append(v5, fi75)
                        _ = fi75
                    }
                    _ = v74
                }
                switch (("a" + ("\n\t" + "\n\t")) + "a") {
                case "hello":
                    var v76 uint8 = 127
                    _ = v76
                default:
                    v77 := GenS6{F0: 65535, F1: v58}
                    _ = v77
                }
                _ = fi59
            }
        case 0:
            for fi78 := 0; fi78 < 11; fi78++ {
                v79 := GenS1{F0: 1, F1: 1, F2: 1000}
                _ = fi78
                _ = v79
            }
        case 42:
            v58 = v58
        }
        _ = v57
        _ = v58
    }
    v80 := v6[""]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v54
    _ = v6
    _ = v80
}
