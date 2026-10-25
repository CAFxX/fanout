package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int16, p1 uint32, p2 string) int32 {
    switch p2 {
    case "\n\t":
        fmt.Println(p2, p2)
    }
    if ((p0 << ((4294967295) & 63)) < p0) {
        v5 := []int{-1, -9223372036854775808}
        v5 = append(v5, -1)
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if ((1.5 + 0.1) <= 1e10) {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 bool = true
                var v9 uint64 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 int16 = (^(p0))
            for fi11 := 0; fi11 < 12; fi11++ {
                switch (^(fi6)) {
                case 1:
                    fmt.Println(p1, p0)
                case 32767:
                    switch fi6 {
                    case 676:
                        var v12 int8 = 16
                        _ = v12
                    case -16:
                        v13 := GenS1{F0: 167, F1: false}
                        _ = v13
                    }
                default:
                    if true {
                        if (((0.0 * 3.14159) / 1e10) > 0.1) {
                            if false {
                                if (!true) {
                                    if true {
                                        if (false && true) {
                                            var v14 int8 = 127
                                            _ = v14
                                        }
                                        p0 = v10
                                        fmt.Println(fi11, p2, fi6)
                                    } else {
                                        v15 := []int{464, 189, (fi11 & 32767), 4294967295}
                                        fmt.Println(v10)
                                        v10 = -178
                                        _ = v15
                                    }
                                    var v16 uint64 = 1
                                    v17 := []int{(fi11 * 443), fi6, (fi11 << ((1) & 63)), fi6}
                                    _ = v16
                                    _ = v17
                                } else {
                                    var v18 int8 = -128
                                    _ = v18
                                }
                                v19 := []int{int(v10), 32767, fi11}
                                _ = v19
                            }
                        } else {
                            switch p2 {
                            case "\n\t":
                                var v20 int64 = -244
                                _ = v20
                            }
                            v21 := GenS1{F0: 0, F1: (!true)}
                            _ = v21
                        }
                        v22 := GenS1{F0: 255, F1: false}
                        var v23 int = fi11
                        _ = v22
                        _ = v23
                    } else {
                        if true {
                            if (0 >= 317) {
                                p0 -= p0
                                var v24 uint = 255
                                _ = v24
                            }
                            var v25 bool = false
                            _ = v25
                        } else {
                            var v26 bool = true
                            var v27 uint = 4294967295
                            var v28 float32 = float32(1e10)
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                }
                _ = fi11
            }
            _ = v10
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                if (!(-9223372036854775808 == 1)) {
                    var v30 float32 = 0.1
                    v30 -= ((v30 + float32(1e10)) / v30)
                    _ = v30
                }
                v31 := GenS2{F0: 2}
                fmt.Println(fi29, fi29)
                _ = fi29
                _ = v31
            }
        }
        var v32 string = (p2 + p2)
        for fi33 := 0; fi33 < 3; fi33++ {
            if (615 < -9223372036854775808) {
                p1 = p1
                v34 := GenS1{F0: 148, F1: (0.1 <= (1e10 + 3.14159))}
                _ = v34
            } else {
                p1 -= p1
                if false {
                    var v35 uint32 = p1
                    _ = v35
                }
            }
            var v36 uint32 = (p1 & (p1 | 2147483648))
            _ = fi33
            _ = v36
        }
        _ = fi6
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v1 = v1
    if ("" > "\n\t") {
        for fi37 := 0; fi37 < 19; fi37++ {
            if false {
                v1 += (v1 << ((v1) & 63))
            }
            v38 := GenS1{F0: 1, F1: false}
            fmt.Println(v1, fi37)
            _ = fi37
            _ = v38
        }
        switch "foo bar" {
        case "foo bar":
            for fi39 := 0; fi39 < 20; fi39++ {
                v1 = 340
                var v40 uint = v1
                _ = fi39
                _ = v40
            }
        case "":
            for fi41 := 0; fi41 < 20; fi41++ {
                if true {
                    var v42 uint = (v1 * v1)
                    _ = v42
                }
                if false {
                    v43 := GenS3{F0: 255, F1: -123}
                    v44 := GenS4{F0: -1, F1: 1000, F2: 2}
                    _ = v43
                    _ = v44
                }
                if false {
                    var v45 int32 = -1
                    _ = v45
                }
                _ = fi41
            }
        default:
            v46 := GenS5{F0: (v1 & (v1 / ((v1) | 1)))}
            _ = v46
        }
        _ = len(v2)
    }
    switch ("hello" + "hello") {
    case "foo bar":
        if (false && true) {
            var v47 uint8 = 42
            _ = v47
        } else {
            if (0 > 65535) {
                v48 := v2[-556]
                var v49 int = 218
                v50 := GenS1{F0: 2, F1: (v1 <= v1)}
                _ = v48
                _ = v49
                _ = v50
            }
            for fi51 := 0; fi51 < 10; fi51++ {
                v52 := GenS6{F0: 32767}
                _ = fi51
                _ = v52
            }
            switch 1 {
            case -1:
                var v53 uint16 = 42
                _ = v53
            case 32767:
                var v54 uint = 42
                _ = v54
            default:
                var v55 string = ""
                _ = v55
            }
        }
    default:
        if (803 < 840) {
            if (true && (float32((-2.25 - (-2.25 / (float64((0.1 + 1e10)) * ((3.14159 / 1e10) / 0.1))))) < 1.5)) {
                if true {
                    v1 += 579
                    if true {
                        v56 := GenS5{F0: v1}
                        _ = len(v4)
                        var v57 string = "a"
                        _ = v56
                        _ = v57
                    }
                    var v58 bool = false
                    _ = v58
                } else {
                    if (!("hello" == ("hello" + "\n\t"))) {
                        v1 += (v1 + 65535)
                    }
                }
                for fi59 := 0; fi59 < 5; fi59++ {
                    var v60 int32 = -501
                    _ = fi59
                    _ = v60
                }
                v1 = 2
            } else {
                v61 := GenS7{F0: 594, F1: 65535, F2: 2}
                fmt.Println(v1)
                _ = v61
            }
            for fi62 := 0; fi62 < 5; fi62++ {
                fmt.Println(fi62, v1)
                _ = fi62
            }
            if (-9223372036854775808 < 4294967295) {
                var v63 int8 = 100
                v1 -= v1
                v64 := []int{2147483647, -250}
                _ = v63
                _ = v64
            } else {
                var v65 bool = (59 == 57)
                switch ((("hello" + (("foo bar" + "a") + ("\n\t" + ""))) + ("foo bar" + "foo bar")) + "foo bar") {
                case "":
                    fmt.Println(v65)
                case "foo bar":
                    fmt.Println(v1, v1)
                default:
                    v66 := GenS6{F0: 255}
                    _ = v66
                }
                var v67 int = -1
                _ = v65
                _ = v67
            }
        } else {
            fmt.Println(v1)
        }
    }
    var v68 string = (((("foo bar" + "hello") + ("foo bar" + "\n\t")) + ("\n\t" + "\n\t")) + ("foo bar" + "hello"))
    switch 755 {
    case -860:
        v1 = 2
    default:
        for fi69 := 0; fi69 < 20; fi69++ {
            v70 := GenS8{F0: -144}
            v71 := GenS1{F0: 1, F1: (fi69 >= (^(fi69)))}
            delete(v4, -9223372036854775808)
            _ = fi69
            _ = v70
            _ = v71
        }
    }
    var v72 string = ("hello" + (v68 + v68))
    for fi73 := 0; fi73 < 1; fi73++ {
        for fi74 := 0; fi74 < 7; fi74++ {
            if (-128 != 42) {
                fmt.Println(fi74, v1)
            }
            _ = fi74
        }
        var v75 int = fi73
        _ = fi73
        _ = v75
    }
    var v76 uint64 = 4294967295
    switch -1 {
    case 855:
        v77 := []int{133, 294}
        _ = v77
    }
    if (((3.14159 / 1e10) > 0.0) && ((false && (true || false)) || (2 == 550))) {
        v78 := GenS5{F0: v1}
        fmt.Println(v76)
        _ = v78
    }
    for fi79 := 0; fi79 < 11; fi79++ {
        v80 := GenS6{F0: 923}
        _ = fi79
        _ = v80
    }
    if false {
        v1 += (v1 | 95)
        var v81 bool = false
        var v82 int16 = 0
        _ = v81
        _ = v82
    } else {
        switch v72 {
        case "foo bar":
            for fi83 := 0; fi83 < 18; fi83++ {
                v1 -= v1
                fmt.Println(v72, fi83, v68)
                if ((1 <= 869) || (-32768 == -640)) {
                    switch 9223372036854775807 {
                    case 42:
                        switch fi83 {
                        case 32768:
                            v76 = v76
                        case -1:
                            v84 := GenS7{F0: 65535, F1: 0, F2: 255}
                            _ = v84
                        default:
                            fmt.Println(v72)
                        }
                    case -587:
                        v85 := GenS6{F0: 42}
                        _ = v85
                    }
                    for fi86 := 0; fi86 < 18; fi86++ {
                        v72 = ("a" + (("a" + "\n\t") + (v72 + (v72 + "a"))))
                        _ = fi86
                    }
                    var v87 int8 = 80
                    _ = v87
                }
                _ = fi83
            }
        }
    }
    _ = len(v2)
    switch v72 {
    case "a":
        v88 := v4[886]
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v72
    _ = v76
}
