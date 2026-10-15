package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
func genFn0(p0 int64) uint16 {
    fmt.Println(p0)
    p0 = (-(p0))
    v5 := GenS1{F0: 237, F1: 0}
    v6 := GenS2{F0: 42, F1: -1}
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 = p0
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 388
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = ((-2.25 + 0.0) * 3.14159)
    v3 := make(map[int]uint32)
    var v4 uint8 = 2
    v1 = 1
    var v8 bool = true
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{}
        v4 = v4
        v4 = (v4 ^ v4)
        _ = fi9
        _ = v10
    }
    var v11 int8 = 124
    if false {
        var v12 int16 = 32767
        fmt.Println(v2, v4)
        _ = v12
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 18; fi14++ {
        if true {
            if false {
                switch "foo bar" {
                case "foo bar":
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v11 += 42
                        var v16 uint32 = 2147483648
                        _ = fi15
                        _ = v16
                    }
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 float32 = (0.1 - ((0.1 + (0.0 / (1e10 / -2.25))) * 0.1))
                            switch v13 {
                            case "foo bar":
                                var v19 float64 = v2
                                _ = v19
                            case "\n\t":
                                var v20 float64 = v2
                                _ = v20
                            default:
                                fmt.Println(v8, v18)
                            }
                            var v21 int8 = (v11 | v11)
                            _ = fi17
                            _ = v18
                            _ = v21
                        }
                    case "foo bar":
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 int = (fi22 * (^(fi22)))
                            fmt.Println(v13, v1, fi22)
                            switch 2147483647 {
                            case 9223372036854775807:
                                v24 := GenS3{F0: 42}
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                    case "":
                        for fi25 := 0; fi25 < 19; fi25++ {
                            switch v13 {
                            case "hello":
                                v1 += (^(v1))
                            case "\n\t":
                                v26 := []int{fi25, fi14, fi14}
                                _ = v26
                            default:
                                if (((!v8) && (32767 < 1)) && v8) {
                                    var v27 float32 = 3.14159
                                    v13 = v13
                                    switch fi14 {
                                    case -132:
                                        fmt.Println(v13)
                                    }
                                    var v28 int = fi14
                                    switch 32768 {
                                    case 127:
                                        fmt.Println(v4, v27, v1)
                                    default:
                                        switch ((v13 + ("" + v13)) + v13) {
                                        case "foo bar":
                                            delete(v3, fi25)
                                        case "hello":
                                            fmt.Println(v11, v28)
                                        case "\n\t":
                                            if (2147483647 < uint32(v1)) {
                                                switch v13 {
                                                case "\n\t":
                                                    v29 := []int{fi25, (v28 >> ((42) & 63)), fi25, (v28 - int(v4))}
                                                    _ = v29
                                                default:
                                                    if ((v13 + (v13 + v13)) != "hello") {
                                                        v1 -= (v1 ^ v1)
                                                        if true {
                                                            v11 = (^(v11))
                                                        } else {
                                                            var v30 uint8 = 2
                                                            _ = v30
                                                        }
                                                    } else {
                                                        v1 += v1
                                                    }
                                                }
                                                var v31 uint32 = 2147483647
                                                var v32 uint8 = v4
                                                _ = v31
                                                _ = v32
                                            } else {
                                                var v33 int64 = -689
                                                if false {
                                                    var v34 int64 = (v33 - v33)
                                                    _ = v34
                                                }
                                                if false {
                                                    var v35 bool = ((fi25 & -118) == 0)
                                                    _ = v35
                                                } else {
                                                    v28 = fi14
                                                    if (v8 && (42 < 2)) {
                                                        var v36 int32 = -82
                                                        v8 = false
                                                        _ = v36
                                                    } else {
                                                        var v37 uint64 = 706
                                                        v38 := GenS4{F0: v8}
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                                _ = v33
                                            }
                                        }
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                var v39 int8 = (v11 / ((1) | 1))
                                switch v13 {
                                case "\n\t":
                                    var v40 int32 = 2147483647
                                    _ = v40
                                case "a":
                                    var v41 string = "foo bar"
                                    _ = v41
                                default:
                                    fmt.Println(v1, v2, v2)
                                }
                                _ = v39
                            }
                            _ = fi25
                        }
                    default:
                        v42 := v3[42]
                        _ = v42
                    }
                }
                for fi43 := 0; fi43 < 5; fi43++ {
                    if v8 {
                        var v44 uint = 805
                        v45 := v3[-1]
                        fmt.Println(fi43)
                        _ = v44
                        _ = v45
                    } else {
                        v46 := GenS2{F0: -1, F1: fi14}
                        if (0.1 <= v2) {
                            var v47 string = v13
                            if (67 < -1) {
                                fmt.Println(v13, fi43)
                                var v48 uint = 42
                                v2 = (-2.25 * v2)
                                _ = v48
                            } else {
                                v13 = "foo bar"
                                var v49 int = (fi14 ^ fi43)
                                fmt.Println(fi43, v1)
                                _ = v49
                            }
                            _ = v47
                        } else {
                            v50 := GenS2{F0: 1, F1: int(v11)}
                            v51 := []int{fi43, (^(fi43)), (fi43 / ((fi43) | 1))}
                            _ = v50
                            _ = v51
                        }
                        _ = v46
                    }
                    _ = fi43
                }
                v52 := []int{(fi14 + fi14), 4294967295}
                _ = v52
            } else {
                var v53 float64 = (1.5 / 1e10)
                _ = v53
            }
            v54 := GenS5{F0: (v4 | 210)}
            var v55 uint8 = v4
            _ = v54
            _ = v55
        } else {
            v11 += (v11 + 1)
        }
        _ = fi14
    }
    fmt.Println(v11, v13)
    if v8 {
        v4 -= 121
        v8 = v8
        v56 := GenS3{F0: 2}
        _ = v56
    } else {
        var v57 uint = 577
        var v58 int8 = 100
        _ = v57
        _ = v58
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        v60 := GenS6{F0: 4294967295}
        _ = fi59
        _ = v60
    }
    v2 = 0.1
    var v61 int32 = 2147483647
    if v8 {
        if (v8 && v8) {
            var v62 uint32 = 2147483647
            _ = v62
        } else {
            v3[-48] = 65535
            switch "hello" {
            case "hello":
                v61 = v61
            case "\n\t":
                var v63 uint16 = v1
                _ = v63
            }
        }
        for fi64 := 0; fi64 < 13; fi64++ {
            switch v13 {
            case "":
                v13 = ""
            }
            fmt.Println(v11, v11, v11)
            v65 := []int{}
            _ = fi64
            _ = v65
        }
        switch -253 {
        case 9223372036854775807:
            for fi66 := 0; fi66 < 6; fi66++ {
                v67 := GenS3{F0: 1}
                _ = fi66
                _ = v67
            }
        case 128:
            fmt.Println(v13, v11, v8)
        default:
            v68 := GenS1{F0: (v1 / ((v1) | 1)), F1: 330}
            _ = v68
        }
    } else {
        v69 := GenS2{F0: -56, F1: 171}
        _ = v69
    }
    v70 := GenS5{F0: (v4 >> ((2) & 63))}
    if (32768 <= -1) {
        if false {
            v71 := []int{4294967295, 1, 0}
            _ = v71
        }
    }
    for fi72 := 0; fi72 < 4; fi72++ {
        switch fi72 {
        case 2147483647:
            v73 := []int{(fi72 - fi72)}
            _ = v73
        }
        v74 := []int{}
        _ = fi72
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v70
    _ = v8
}
