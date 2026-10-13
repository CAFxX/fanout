package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
func genFn0(p0 uint16, p1 int16, p2 uint16) string {
    v7 := []int{}
    var v8 uint8 = 0
    p2 = 704
    v8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 int16 = 608
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483648
    var v4 string = "foo bar"
    var v5 string = v4
    v6 := make(map[int]bool)
    if false {
        var v9 int16 = v1
        _ = len(v6)
        _ = v9
    }
    v10 := GenS1{F0: (v5 + ((("foo bar" + "\n\t") + "\n\t") + v4)), F1: 1}
    v11 := GenS2{F0: 2}
    if (2 < 215) {
        var v12 float32 = (3.14159 - 0.1)
        _ = v12
    } else {
        if true {
            var v13 uint8 = 46
            switch -14 {
            case -921:
                if (-17 > 42) {
                    v3 = uint32(v13)
                    var v14 uint = 65535
                    _ = v14
                }
            default:
                v15 := []int{58, 32768}
                _ = v15
            }
            _ = v13
        }
        if false {
            var v16 uint = 4294967295
            switch "\n\t" {
            case "hello":
                v3 += v3
            case "":
                v17 := GenS3{F0: 1, F1: 100}
                _ = v17
            case "\n\t":
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch v5 {
                        case "":
                            v4 = v4
                        case "a":
                            switch v5 {
                            case "a":
                                v20 := GenS4{F0: v3, F1: uint8(v3)}
                                _ = v20
                            }
                        default:
                            delete(v2, (v3 / ((429) | 1)))
                        }
                        v3 = v3
                        var v21 uint16 = 0
                        if true {
                            v5 = v5
                        } else {
                            var v22 int32 = -258
                            v16 -= v16
                            if (true || (v16 <= (v16 ^ v16))) {
                                v22 += (v22 - v22)
                                if true {
                                    v1 -= v1
                                    var v23 uint64 = 0
                                    _ = v23
                                }
                                if false {
                                    v24 := []int{(-(fi18)), (fi19 - (fi19 % (((fi19 - 0)) | 1)))}
                                    v21 += v21
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    _ = v24
                                }
                                var v25 int32 = v22
                                switch (fi19 + (fi18 + 32768)) {
                                case 4294967295:
                                    v22 -= 154
                                case 32768:
                                    var v26 int = 9223372036854775807
                                    var v27 int8 = -1
                                    var v28 float32 = (0.1 * 0.0)
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                default:
                                    v2[65535] = 128
                                }
                                _ = v25
                            } else {
                                v29 := []int{(-(fi18)), 4294967295, 195}
                                v22 = (v22 / (((v22 * int32(fi18))) | 1))
                                v1 -= (-(v1))
                                _ = v29
                            }
                            _ = v22
                        }
                        _ = fi19
                        _ = v21
                    }
                    switch (v5 + v5) {
                    case "hello":
                        var v30 uint32 = v3
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 7; fi31++ {
                            var v32 bool = false
                            v33 := GenS5{F0: 127}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    v4 = (v4 + (v5 + ((v5 + "hello") + ((v4 + "foo bar") + v4))))
                    _ = fi18
                }
            default:
                fmt.Println(v4, v5)
            }
            _ = v16
        } else {
            var v34 float64 = 1e10
            v35 := []int{9223372036854775807, 128, -2147483648}
            v1 = v1
            _ = v34
            _ = v35
        }
        v5 = v4
    }
    var v36 int16 = -741
    v4 = (v5 + ((("a" + (v4 + v5)) + "hello") + "foo bar"))
    for fi37 := 0; fi37 < 6; fi37++ {
        if ((!true) && false) {
            if (!((!(2 <= 578)) && true)) {
                for fi38 := 0; fi38 < 15; fi38++ {
                    if true {
                        fmt.Println(fi38, fi38, v4)
                    }
                    switch (v4 + v4) {
                    case "hello":
                        fmt.Println(v1, v4)
                    case "":
                        v39 := []int{(fi37 % (((fi38 >> ((2) & 63))) | 1)), (fi38 + (fi38 - fi37)), fi38, -900}
                        _ = v39
                    default:
                        delete(v6, fi38)
                    }
                    v40 := GenS2{F0: 32767}
                    _ = fi38
                    _ = v40
                }
            }
            var v41 uint32 = (v3 & v3)
            _ = v41
        } else {
            var v42 float32 = (float32((0.0 - 0.0)) - -2.25)
            _ = v42
        }
        v43 := GenS6{F0: (v5 <= v4)}
        v3 = 2147483647
        _ = fi37
        _ = v43
    }
    if false {
        switch -777 {
        case 32768:
            var v44 uint16 = 995
            _ = v44
        case -290:
            v45 := GenS3{F0: 0, F1: -128}
            _ = v45
        default:
            switch "\n\t" {
            case "foo bar":
                v5 = (v5 + "\n\t")
            case "":
                v1 -= 491
            case "\n\t":
                switch (v4 + v5) {
                case "a":
                    if true {
                        var v46 bool = true
                        if (11 <= 1) {
                            v3 += 2147483647
                            var v47 float64 = 1.5
                            _ = v47
                        } else {
                            var v48 string = "a"
                            v1 = v36
                            v46 = v46
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 19; fi49++ {
                            for fi50 := 0; fi50 < 16; fi50++ {
                                var v51 int = fi49
                                fmt.Println(v3, v36, v3)
                                if (!(false && false)) {
                                    v5 = v5
                                }
                                _ = fi50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 15; fi52++ {
                                var v53 int8 = 100
                                var v54 int8 = v53
                                v55 := GenS7{F0: v46}
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            for fi56 := 0; fi56 < 15; fi56++ {
                                var v57 uint64 = 373
                                if (!((-2.25 + float64(-2.25)) == float64(0.0))) {
                                    var v58 uint8 = 255
                                    switch (fi49 + fi49) {
                                    case 32768:
                                        if true {
                                            v57 += (v57 - v57)
                                            v57 = v57
                                        } else {
                                            switch v5 {
                                            case "foo bar":
                                                var v59 uint64 = (v57 | 178)
                                                _ = v59
                                            case "":
                                                v36 += (v1 / (((^(v1))) | 1))
                                            }
                                        }
                                    case -9223372036854775808:
                                        if true {
                                            var v60 float64 = ((-2.25 * float64(3.14159)) / 1.0)
                                            if (!false) {
                                                var v61 uint16 = 32767
                                                v62 := []int{fi56, (fi49 * fi56), 1, -1}
                                                _ = v61
                                                _ = v62
                                            }
                                            v60 -= 1.5
                                            _ = v60
                                        } else {
                                            var v63 string = ""
                                            _ = v63
                                        }
                                    }
                                    if false {
                                        var v64 int64 = 42
                                        var v65 bool = ((v64 * (v64 ^ v64)) < (v64 * v64))
                                        _ = v64
                                        _ = v65
                                    } else {
                                        v66 := GenS2{F0: uint16(v57)}
                                        var v67 uint16 = 0
                                        _ = v66
                                        _ = v67
                                    }
                                    _ = v58
                                }
                                _ = fi56
                                _ = v57
                            }
                            _ = fi49
                        }
                        _ = v46
                    } else {
                        var v68 string = (v4 + v4)
                        var v69 uint64 = 1
                        switch -513 {
                        case 0:
                            var v70 bool = ((!false) && ((2147483648 >= (v3 ^ v3)) || false))
                            _ = v70
                        }
                        _ = v68
                        _ = v69
                    }
                case "foo bar":
                    switch -9223372036854775808 {
                    case 32767:
                        v71 := v6[-962]
                        _ = v71
                    default:
                        switch v4 {
                        case "\n\t":
                            var v72 int = 0
                            _ = v72
                        case "a":
                            var v73 uint32 = 0
                            _ = v73
                        }
                    }
                case "\n\t":
                    for fi74 := 0; fi74 < 5; fi74++ {
                        fmt.Println(fi74, v5)
                        var v75 float64 = (0.1 - -2.25)
                        if true {
                            var v76 uint64 = 4294967295
                            v77 := GenS8{F0: 2147483647}
                            _ = v76
                            _ = v77
                        }
                        _ = fi74
                        _ = v75
                    }
                }
            }
        }
        fmt.Println(v36)
        v78 := GenS3{F0: 1, F1: 127}
        _ = v78
    }
    var v79 float32 = 1e10
    fmt.Println(v4)
    v4 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
