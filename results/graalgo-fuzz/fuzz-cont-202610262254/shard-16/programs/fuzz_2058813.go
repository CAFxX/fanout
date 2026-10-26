package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 uint) int {
    var v4 int32 = -252
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 = p1
        _ = fi5
    }
    v6 := GenS1{F0: 779, F1: 70}
    v7 := GenS2{F0: 1, F1: 255, F2: 62}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16) int {
    v8 := []int{127}
    v8 = append(v8, -952)
    _ = p0
    _ = v8
    return 579
}

func main() {
    var v1 string = ""
    var v2 float64 = -2.25
    v3 := make(map[int]int)
    v1 = "a"
    v9 := GenS1{F0: 0, F1: 1}
    var v10 uint32 = 65535
    switch genFn0(42, 4294967295) {
    case 393:
        if false {
            var v11 uint16 = 2
            _ = v11
        } else {
            var v12 string = v1
            for fi13 := 0; fi13 < 10; fi13++ {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS2{F0: -9223372036854775808, F1: 255, F2: 97}
                    if true {
                        fmt.Println(fi14, v1)
                        fmt.Println(fi13)
                    } else {
                        var v16 uint8 = 42
                        v16 = (v16 << ((716) & 63))
                        v17 := GenS1{F0: 388, F1: 750}
                        v18 := []int{(^(fi14)), (fi14 * 2147483647), -9223372036854775808, genFn1(32767)}
                        switch (-(fi13)) {
                        case -820:
                            var v19 float64 = (-2.25 + (((float64(float32(3.14159)) - v2) + v2) + (float64(-2.25) - 1e10)))
                            _ = v19
                        case 4294967295:
                            fmt.Println(fi13)
                        }
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v20 int32 = 42
                    _ = fi14
                    _ = v15
                    _ = v20
                }
                v12 = ("hello" + "foo bar")
                _ = fi13
            }
            _ = v12
        }
    case 1:
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := []int{genFn0(4294967295, 405), (fi22 << ((114) & 63))}
                _ = len(v23)
                v2 = v2
                _ = fi22
                _ = v23
            }
            if true {
                v24 := GenS1{F0: 1, F1: 0}
                if false {
                    v25 := GenS3{F0: (v10 * v10), F1: 1, F2: 32767}
                    _ = v25
                }
                if (-1 == -85) {
                    v2 = 3.14159
                }
                _ = v24
            }
            _ = fi21
        }
    default:
        switch "foo bar" {
        case "\n\t":
            v2 += (v2 / (-2.25 * v2))
        }
    }
    v3[9223372036854775807] = -1
    var v26 int32 = -333
    var v27 int = -2147483648
    switch ("a" + "foo bar") {
    case "foo bar":
        v28 := v3[(^(v27))]
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 uint32 = v10
        if (255 > 360) {
            if (617 == 255) {
                for fi31 := 0; fi31 < 11; fi31++ {
                    if true {
                        var v32 float32 = 1e10
                        _ = v32
                    } else {
                        var v33 float32 = 0.0
                        _ = v33
                    }
                    _ = fi31
                }
            }
        }
        _ = fi29
        _ = v30
    }
    var v34 float32 = -2.25
    var v35 uint32 = (v10 ^ (v10 % (((v10 - 2147483648)) | 1)))
    switch v1 {
    case "\n\t":
        if (false || (-512 >= 42)) {
            var v36 int32 = -1
            _ = len(v3)
            _ = v36
        }
    case "foo bar":
        if true {
            v37 := GenS1{F0: 314, F1: 65535}
            v1 = "foo bar"
            fmt.Println(v10, v26, v10)
            _ = v37
        } else {
            v10 -= (v10 << ((18446744073709551615) & 63))
        }
    case "a":
        if true {
            v38 := GenS1{F0: 42, F1: 65535}
            fmt.Println(v26)
            _ = v38
        } else {
            v39 := []int{v27, (v27 >> ((483) & 63)), int(v10)}
            if ((v26 & -469) > v26) {
                switch v1 {
                case "foo bar":
                    var v40 int64 = 9223372036854775807
                    _ = v40
                case "hello":
                    for fi41 := 0; fi41 < 15; fi41++ {
                        v42 := GenS4{F0: 1}
                        switch (((v1 + v1) + v1) + "\n\t") {
                        case "":
                            for fi43 := 0; fi43 < 1; fi43++ {
                                fmt.Println(v34, v26, v1)
                                var v44 uint32 = v35
                                switch genFn1(32767) {
                                case 2147483647:
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        var v45 bool = (212 > 1000)
                                        if v45 {
                                            if (!(65535 < 65535)) {
                                                var v46 int64 = 42
                                                v35 += (^(v10))
                                                if (!v45) {
                                                    if (v1 < v1) {
                                                        var v47 string = ((v1 + "a") + v1)
                                                        var v48 uint16 = 165
                                                        var v49 int64 = (^(v46))
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                } else {
                                                    _ = len(v3)
                                                }
                                                _ = v46
                                            }
                                        } else {
                                            var v50 int16 = -32768
                                            var v51 bool = v45
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v45
                                    }
                                }
                                _ = fi43
                                _ = v44
                            }
                        case "hello":
                            for fi52 := 0; fi52 < 20; fi52++ {
                                switch v1 {
                                case "a":
                                    switch v1 {
                                    case "a":
                                        fmt.Println(v26, v26)
                                    default:
                                        switch v1 {
                                        case "":
                                            switch v1 {
                                            case "\n\t":
                                                v53 := GenS3{F0: (v10 >> ((1) & 63)), F1: int32(v27), F2: 42}
                                                _ = v53
                                            case "hello":
                                                switch (fi41 & v27) {
                                                case -2147483648:
                                                    if ((true && true) && false) {
                                                        switch "" {
                                                        case "\n\t":
                                                            v10 -= v35
                                                        case "foo bar":
                                                            v54 := GenS3{F0: v10, F1: -1, F2: 32767}
                                                            _ = v54
                                                        }
                                                    }
                                                case 32768:
                                                    if (false || (42 != 4294967295)) {
                                                        switch v1 {
                                                        case "\n\t":
                                                            var v55 float64 = (-2.25 + ((float64(3.14159) - (0.0 + v2)) * v2))
                                                            _ = v55
                                                        case "a":
                                                            fmt.Println(v27)
                                                        case "":
                                                            var v56 int32 = v26
                                                            _ = v56
                                                        }
                                                        var v57 int32 = v26
                                                        if len(v39) >= 2 {
                                                            v39 = v39[1:len(v39)]
                                                        }
                                                        _ = v57
                                                    }
                                                case 4294967295:
                                                    v2 += -2.25
                                                default:
                                                    v58 := GenS5{F0: 641}
                                                    _ = v58
                                                }
                                            case "a":
                                                var v59 uint8 = 240
                                                _ = v59
                                            default:
                                                v60 := GenS6{F0: (true && (!((v34 < 1e10) && (!false))))}
                                                _ = v60
                                            }
                                        case "\n\t":
                                            v26 = (v26 * 2147483647)
                                        default:
                                            var v61 uint8 = 246
                                            _ = v61
                                        }
                                    }
                                default:
                                    switch (v1 + v1) {
                                    case "a":
                                        var v62 int32 = 42
                                        _ = v62
                                    default:
                                        if (177 <= 903) {
                                            v63 := GenS2{F0: -705, F1: 255, F2: 1}
                                            if false {
                                                v26 -= 2147483647
                                                v35 = (v35 << ((0) & 63))
                                            }
                                            v10 += v10
                                            v64 := GenS7{F0: (-(v26)), F1: 42, F2: 1000}
                                            _ = v63
                                            _ = v64
                                        } else {
                                            fmt.Println(v26, v1)
                                            if true {
                                                var v65 int64 = 0
                                                fmt.Println(v27)
                                                _ = v65
                                            }
                                        }
                                    }
                                }
                                var v66 int32 = v26
                                _ = fi52
                                _ = v66
                            }
                        default:
                            if false {
                                for fi67 := 0; fi67 < 11; fi67++ {
                                    fmt.Println(fi41, v1, fi67)
                                    var v68 int = (-(v27))
                                    v35 = v10
                                    _ = fi67
                                    _ = v68
                                }
                                v27 += 127
                                var v69 float32 = 0.0
                                _ = v69
                            }
                        }
                        v10 += (v10 | 2147483647)
                        _ = fi41
                        _ = v42
                    }
                case "\n\t":
                    var v70 bool = true
                    _ = v70
                }
            }
            _ = v39
        }
    default:
        if false {
            if ((v26 << ((0) & 63)) != 42) {
                v10 = (v35 & v35)
            }
            fmt.Println(v10, v10, v2)
        } else {
            var v71 int = v27
            delete(v3, (v27 << ((65535) & 63)))
            for fi72 := 0; fi72 < 13; fi72++ {
                fmt.Println(v1, v35, v71)
                _ = fi72
            }
            _ = v71
        }
    }
    var v73 float64 = v2
    if (false || true) {
        var v74 int8 = 3
        var v75 int = 788
        v26 = -1
        _ = v74
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v73
    _ = v9
}
