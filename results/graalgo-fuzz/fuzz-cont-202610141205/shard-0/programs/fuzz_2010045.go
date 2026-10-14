package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int8; F2 int8 }
type GenS7 struct { F0 bool; F1 uint32; F2 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8 }
func genFn0(p0 uint32) uint {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 1, F1: 0, F2: fi5}
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{1, (fi7 & 127), (fi5 >> ((1) & 63))}
            if true {
                switch fi5 {
                case -2147483648:
                    var v9 float32 = 3.14159
                    _ = v9
                case 32767:
                    p0 += p0
                case 923:
                    var v10 float64 = ((3.14159 / (1e10 + 3.14159)) / float64(0.1))
                    _ = v10
                }
            }
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v11 := GenS2{F0: 255}
    v12 := GenS3{F0: 2}
    switch "hello" {
    case "foo bar":
        if true {
            v13 := []int{2147483647}
            _ = v13
        }
    case "\n\t":
        fmt.Println(p0)
    case "a":
        p0 = (p0 / ((p0) | 1))
    }
    var v14 float32 = 0.0
    _ = p0
    _ = v11
    _ = v12
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := []int{fi15, (fi16 + -9223372036854775808)}
            var v18 int16 = 32767
            if (14 != -1) {
                v19 := GenS1{F0: 275, F1: 412, F2: 767}
                if (false && true) {
                    v4 = v4
                    v20 := GenS4{F0: 1, F1: fi15, F2: v4}
                    if (true || false) {
                        switch (fi15 % ((fi16) | 1)) {
                        case 127:
                            var v21 int64 = 1
                            _ = v21
                        default:
                            v18 = (-(v18))
                        }
                        v22 := GenS5{F0: 4294967295}
                        _ = v22
                    }
                    _ = v20
                } else {
                    v23 := GenS6{F0: ("foo bar" + "\n\t"), F1: 1, F2: -1}
                    v4 = ""
                    v24 := GenS1{F0: 0, F1: 909, F2: fi16}
                    v4 = "hello"
                    v4 = (v4 + v4)
                    _ = v23
                    _ = v24
                }
                v25 := GenS6{F0: "", F1: 1, F2: -91}
                v4 = (v4 + (v4 + v4))
                _ = v19
                _ = v25
            } else {
                var v26 int8 = -96
                switch fi15 {
                case 4294967295:
                    v4 = ((v4 + "\n\t") + (v4 + (("" + v4) + v4)))
                default:
                    if len(v17) > 0 {
                        v17[0] = fi16
                    }
                }
                _ = v26
            }
            if false {
                var v27 int32 = 1
                if (1 == 756) {
                    v4 = ((v4 + (("a" + v4) + v4)) + v4)
                    v4 = (v4 + v4)
                } else {
                    v27 = v27
                    var v28 uint8 = 0
                    v17 = append(v17, 32767)
                    _ = v28
                }
                var v29 uint8 = 255
                _ = v27
                _ = v29
            } else {
                v18 = 42
                if true {
                    switch (fi15 & 9223372036854775807) {
                    case 128:
                        v3[93] = 65535
                    default:
                        v30 := v2[2147483648]
                        _ = v30
                    }
                }
            }
            var v31 string = (("hello" + v4) + v4)
            _ = fi16
            _ = v17
            _ = v18
            _ = v31
        }
        _ = fi15
    }
    if false {
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 uint16 = 32767
            switch 39 {
            case 128:
                fmt.Println(v4)
            case -263:
                v1[4294967295] = ""
            }
            fmt.Println(fi32)
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        switch v4 {
        case "":
            v4 = v4
        }
        _ = fi34
    }
    delete(v3, 288)
    v3[42] = 2147483648
    var v35 int = 940
    var v36 uint32 = 2
    var v37 uint16 = 2
    if (float32(-2.25) == ((-2.25 - 1.5) * 3.14159)) {
        v37 -= v37
    }
    var v38 float64 = (1e10 + 0.0)
    switch "foo bar" {
    case "foo bar":
        for fi39 := 0; fi39 < 8; fi39++ {
            fmt.Println(v36, fi39, fi39)
            if (9223372036854775807 < 0) {
                v40 := []int{}
                var v41 float64 = v38
                _ = v40
                _ = v41
            } else {
                fmt.Println(v38, v38)
                if ((127 != 41) || (440 == 18446744073709551615)) {
                    delete(v3, (v36 % ((v36) | 1)))
                } else {
                    v42 := GenS6{F0: v4, F1: 1, F2: -1}
                    v1[(v36 & v36)] = v4
                    v35 -= 32768
                    _ = v42
                }
            }
            delete(v3, v36)
            _ = fi39
        }
    case "hello":
        var v43 int = -1
        _ = v43
    }
    fmt.Println(v36, v35)
    if (2147483647 >= 42) {
        v35 -= (v35 | v35)
        for fi44 := 0; fi44 < 7; fi44++ {
            fmt.Println(v4, v37, v36)
            switch v4 {
            case "hello":
                v45 := GenS3{F0: 2}
                _ = v45
            case "foo bar":
                var v46 int16 = 0
                _ = v46
            case "\n\t":
                v4 = v4
            default:
                _ = len(v3)
            }
            _ = fi44
        }
    }
    if true {
        fmt.Println(v4, v35, v4)
    } else {
        v47 := GenS7{F0: (128 <= 1), F1: (^(v36)), F2: v35}
        _ = v47
    }
    v38 = v38
    v35 = -2147483648
    switch (^(v35)) {
    case -9223372036854775808:
        for fi48 := 0; fi48 < 11; fi48++ {
            for fi49 := 0; fi49 < 9; fi49++ {
                if true {
                    v50 := GenS5{F0: 255}
                    v4 = v4
                    if true {
                        _ = len(v1)
                        v37 = (v37 - (v37 << ((genFn0((v36 + v36))) & 63)))
                    } else {
                        _ = len(v1)
                        if (false || (2147483648 != 2147483647)) {
                            v51 := GenS8{F0: 618}
                            fmt.Println(v35)
                            _ = v51
                        } else {
                            var v52 float64 = float64(-2.25)
                            var v53 int64 = -1
                            var v54 float64 = v38
                            v38 -= 0.1
                            v55 := []int{fi49, v35, -217}
                            _ = v52
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        if (v37 > v37) {
                            var v56 uint = uint(v37)
                            if (!false) {
                                v57 := GenS9{F0: 128}
                                _ = v57
                            } else {
                                fmt.Println(v36)
                            }
                            _ = v56
                        } else {
                            v58 := []int{(fi48 >> ((uint(v37)) & 63)), fi49}
                            v59 := GenS3{F0: 128}
                            var v60 uint16 = v37
                            switch 4294967295 {
                            case 42:
                                v60 -= (v60 ^ 32767)
                            case 9223372036854775807:
                                if (9223372036854775807 < 42) {
                                    v35 = (fi49 * v35)
                                }
                            case 32767:
                                v58 = append(v58, v35)
                            default:
                                v38 = (v38 / ((float64(float32(0.1)) + (float64(0.0) - v38)) / v38))
                            }
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                    }
                    _ = v50
                }
                fmt.Println(v35, v4, v36)
                _ = fi49
            }
            var v61 int64 = 9223372036854775807
            _ = fi48
            _ = v61
        }
    case 32768:
        for fi62 := 0; fi62 < 2; fi62++ {
            var v63 int8 = 1
            _ = fi62
            _ = v63
        }
    default:
        v64 := v2[(v36 << ((genFn0(v36)) & 63))]
        _ = v64
    }
    if (!((!(!(false || (v4 == (v4 + v4))))) || (-672 != 675))) {
        v36 += v36
        v65 := []int{127}
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
