package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 string; F1 int; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint, p1 int8, p2 bool) uint16 {
    if p2 {
        if (0 > 1) {
            if p2 {
                v5 := GenS1{F0: "hello"}
                _ = v5
            } else {
                v6 := []int{-576, 42, 4294967295, 9223372036854775807}
                _ = v6
            }
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            if p2 {
                p0 -= (^(p0))
                p2 = p2
                v8 := []int{42}
                _ = v8
            } else {
                var v9 uint16 = 236
                if (!(667 < 1)) {
                    v10 := []int{fi7, fi7, (fi7 - 0), (^(fi7))}
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        fmt.Println(fi11)
                        p0 -= p0
                        if (fi7 == (fi7 ^ (fi7 & int(p1)))) {
                            v9 -= v9
                            v12 := GenS2{F0: 397, F1: p0, F2: (p0 | 1)}
                            _ = v12
                        } else {
                            v9 = (v9 * (v9 | (v9 - v9)))
                        }
                        _ = fi11
                    }
                }
                p0 = p0
                _ = v9
            }
            _ = fi7
        }
    } else {
        v13 := GenS3{F0: (true && (3.14159 != 3.14159))}
        switch "a" {
        case "":
            v14 := GenS4{F0: 2, F1: 4294967295}
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS5{F0: "hello", F1: fi15, F2: (fi15 ^ fi15)}
                _ = fi15
                _ = v16
            }
        default:
            var v17 float64 = 1.5
            _ = v17
        }
        var v18 float64 = 0.0
        _ = v13
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        p0 += (p0 & p0)
        _ = fi19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        p2 = p2
        if ((1.5 < 0.1) && p2) {
            p2 = p2
        } else {
            for fi21 := 0; fi21 < 11; fi21++ {
                if true {
                    v22 := GenS2{F0: 42, F1: 2, F2: p0}
                    _ = v22
                } else {
                    v23 := []int{(fi21 - (fi20 + 1)), (-(fi21))}
                    var v24 uint32 = 2147483648
                    v25 := GenS1{F0: ((("" + "foo bar") + ("foo bar" + ("" + "\n\t"))) + ("" + "foo bar"))}
                    _ = v23
                    _ = v24
                    _ = v25
                }
                p0 += (p0 / (((p0 % ((671) | 1))) | 1))
                if (fi20 <= fi21) {
                    if ((p0 - (p0 % (((p0 << ((p0) & 63))) | 1))) != 65535) {
                        var v26 uint32 = 349
                        switch 4294967295 {
                        case -1:
                            p2 = p2
                        case 224:
                            var v27 uint8 = 116
                            _ = v27
                        default:
                            v26 = v26
                        }
                        var v28 int32 = 765
                        _ = v26
                        _ = v28
                    } else {
                        fmt.Println(p2, p1)
                    }
                    v29 := []int{}
                    switch "hello" {
                    case "":
                        switch -9 {
                        case 0:
                            var v30 int16 = 1
                            _ = v30
                        case -9223372036854775808:
                            var v31 string = ("" + "")
                            _ = v31
                        default:
                            v32 := GenS2{F0: 2147483647, F1: (p0 % ((p0) | 1)), F2: p0}
                            _ = v32
                        }
                    case "foo bar":
                        switch (fi20 / (((fi20 << ((p0) & 63))) | 1)) {
                        case 127:
                            v33 := GenS5{F0: "a", F1: (fi21 | 32767), F2: int(p1)}
                            _ = v33
                        case 2147483647:
                            v29 = append(v29, fi20)
                        }
                    }
                    if (false && true) {
                        var v34 uint64 = 2
                        _ = v34
                    } else {
                        if p2 {
                            p0 = p0
                            switch ("hello" + "a") {
                            case "\n\t":
                                v35 := GenS1{F0: "hello"}
                                _ = v35
                            case "a":
                                p1 = (p1 % ((p1) | 1))
                            case "hello":
                                p2 = ((p0 * p0) > (p0 | 618))
                            }
                            if (p2 || (!true)) {
                                v36 := GenS3{F0: true}
                                var v37 uint16 = 596
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS1{F0: "foo bar"}
                                v29 = append(v29, fi21)
                                p2 = (154 <= 128)
                                _ = v38
                            }
                            v29 = append(v29, 309)
                        } else {
                            if false {
                                fmt.Println(p2, p1, fi20)
                                if true {
                                    var v39 string = ""
                                    _ = v39
                                }
                                p1 -= p1
                            }
                            var v40 float64 = 1e10
                            _ = v40
                        }
                    }
                    _ = len(v29)
                    _ = len(v29)
                    switch -9223372036854775808 {
                    case -611:
                        v41 := GenS3{F0: (42 == -595)}
                        _ = v41
                    }
                    _ = v29
                } else {
                    var v42 int = (fi20 >> (((p0 / ((p0) | 1))) & 63))
                    fmt.Println(p0, p2, fi20)
                    if ((!p2) || ((0.1 + (3.14159 / 1.0)) < -2.25)) {
                        if p2 {
                            p2 = (1e10 <= float32(1.5))
                            fmt.Println(fi20)
                            switch (^(fi21)) {
                            case 207:
                                switch (fi20 % (((^(fi21))) | 1)) {
                                case -781:
                                    var v43 int32 = -683
                                    _ = v43
                                case 9223372036854775807:
                                    v44 := []int{fi21}
                                    _ = v44
                                case 127:
                                    fmt.Println(v42)
                                }
                            case 616:
                                var v45 uint64 = 131
                                _ = v45
                            case 850:
                                fmt.Println(p0, p0, fi21)
                            }
                        } else {
                            p1 -= p1
                        }
                    }
                    _ = v42
                }
                _ = fi21
            }
        }
        _ = fi20
    }
    if p2 {
        var v46 int32 = -1
        for fi47 := 0; fi47 < 13; fi47++ {
            v48 := GenS1{F0: "hello"}
            p0 += (p0 % ((p0) | 1))
            var v49 int32 = v46
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v46
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]string)
    var v3 uint = 2
    var v4 uint = 42
    switch -1 {
    case 42:
        var v50 int64 = v1
        _ = v50
    default:
        v1 = (v1 + v1)
    }
    delete(v2, -9223372036854775808)
    v3 += 42
    var v51 uint16 = 65535
    var v52 uint64 = 18446744073709551615
    if true {
        fmt.Println(v52)
    } else {
        v3 += v3
        for fi53 := 0; fi53 < 6; fi53++ {
            v54 := GenS6{F0: (v4 >> ((0) & 63))}
            _ = fi53
            _ = v54
        }
        v55 := GenS1{F0: ((("\n\t" + (("\n\t" + "foo bar") + "\n\t")) + ("foo bar" + "foo bar")) + (("hello" + "") + "a"))}
        _ = v55
    }
    v1 = (v1 - (v1 >> ((v4) & 63)))
    var v56 string = ("foo bar" + ("\n\t" + ((("" + ("foo bar" + "hello")) + "") + ((("hello" + "foo bar") + "foo bar") + ("a" + ("\n\t" + ((("\n\t" + "\n\t") + "\n\t") + ("\n\t" + ("foo bar" + ("" + ("a" + "hello")))))))))))
    if (2 <= 698) {
        var v57 uint8 = 88
        v57 = 127
        _ = v57
    } else {
        if (1 != 1) {
            if true {
                v58 := []int{520, 32767, 205}
                var v59 int64 = (v1 / ((v1) | 1))
                if false {
                    var v60 uint16 = (^(v51))
                    switch 128 {
                    case 127:
                        fmt.Println(v60, v1, v52)
                    case 2147483647:
                        v61 := GenS7{F0: 2, F1: 42}
                        _ = v61
                    }
                    _ = v60
                }
                _ = v58
                _ = v59
            } else {
                for fi62 := 0; fi62 < 17; fi62++ {
                    v63 := v2[fi62]
                    _ = fi62
                    _ = v63
                }
                for fi64 := 0; fi64 < 2; fi64++ {
                    fmt.Println(v56)
                    v51 -= (v51 >> ((v4) & 63))
                    _ = fi64
                }
                fmt.Println(v1, v3, v51)
            }
            switch 127 {
            case 2147483647:
                var v65 int8 = 42
                _ = v65
            default:
                var v66 uint64 = 0
                _ = v66
            }
        } else {
            v4 = v4
        }
        if (-2.25 != 0.0) {
            var v67 float32 = 3.14159
            switch 955 {
            case -268:
                v56 = ""
            case 1:
                v52 = v52
            }
            _ = v67
        }
    }
    v2[4294967295] = v56
    switch 32767 {
    case 127:
        switch v56 {
        case "":
            if true {
                var v68 int64 = v1
                v52 = 18446744073709551615
                _ = v68
            }
        case "\n\t":
            v69 := GenS5{F0: "", F1: 2147483647, F2: 432}
            _ = v69
        default:
            for fi70 := 0; fi70 < 19; fi70++ {
                v1 = 0
                _ = fi70
            }
        }
    case 318:
        switch "foo bar" {
        case "":
            fmt.Println(v3, v51, v51)
        case "hello":
            var v71 uint16 = v51
            _ = v71
        case "foo bar":
            fmt.Println(v52, v3, v1)
        default:
            var v72 int64 = (v1 << (((^(v3))) & 63))
            _ = v72
        }
    case 0:
        v1 += v1
    default:
        v73 := []int{-304, -394}
        _ = v73
    }
    var v74 uint64 = uint64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v56
    _ = v74
}
