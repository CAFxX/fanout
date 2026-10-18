package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint; F1 string; F2 bool }
func genFn0(p0 int32, p1 int64, p2 string) int64 {
    switch "hello" {
    case "a":
        p1 = p1
    case "foo bar":
        p2 = "foo bar"
    case "":
        v5 := GenS1{F0: ((p2 + p2) + ((("hello" + "foo bar") + ("\n\t" + (p2 + p2))) + p2))}
        _ = v5
    default:
        v6 := GenS2{F0: 42, F1: p1, F2: (!true)}
        _ = v6
    }
    if (3.14159 >= (-2.25 / 1.0)) {
        fmt.Println(p2, p1)
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                v8 := []int{}
                _ = v8
            }
            p2 = "foo bar"
            _ = fi7
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if (true && true) {
            v10 := GenS2{F0: 1, F1: (p1 * p1), F2: (int16(fi9) == -32768)}
            v11 := GenS1{F0: ((("" + "\n\t") + (("hello" + p2) + p2)) + ((p2 + (p2 + "hello")) + ("foo bar" + (p2 + p2))))}
            _ = v10
            _ = v11
        }
        var v12 int64 = 42
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1() bool {
    if true {
        var v13 uint64 = 1
        v13 += v13
        v13 = 18446744073709551615
        _ = v13
    } else {
        v14 := []int{32768, 1, 4294967295, 514}
        _ = v14
    }
    if false {
        var v15 bool = false
        _ = v15
    }
    var v16 uint64 = 2
    _ = v16
    return true
}

func main() {
    v1 := []string{"", "hello"}
    var v2 float64 = -2.25
    var v3 uint8 = 2
    var v4 int64 = -1
    var v17 int = 42
    for fi18 := 0; fi18 < 9; fi18++ {
        v17 = fi18
        if (((0.0 + 0.0) - 0.0) < -2.25) {
            var v19 int8 = -4
            _ = v19
        }
        if (!(232 == 2147483647)) {
            fmt.Println(v3, v2, v3)
            switch (v17 ^ 413) {
            case -302:
                for fi20 := 0; fi20 < 3; fi20++ {
                    v4 = 0
                    _ = fi20
                }
            case 638:
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 int64 = 267
                    v4 = (v4 & genFn0(-890, v22, "\n\t"))
                    _ = fi21
                    _ = v22
                }
            default:
                v2 -= float64(float32(float64(1e10)))
            }
            if (false && (42 < 916)) {
                switch "hello" {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    v2 += ((v2 + v2) / v2)
                default:
                    var v23 int = (^(fi18))
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    if true {
                        v25 := GenS1{F0: "\n\t"}
                        _ = v25
                    }
                    v26 := GenS3{F0: 230, F1: -661, F2: 42}
                    _ = fi24
                    _ = v26
                }
                var v27 int8 = 1
                _ = v27
            }
        } else {
            v3 = v3
        }
        _ = fi18
    }
    var v28 int = (-(v17))
    var v29 int8 = 100
    var v30 string = "a"
    var v31 uint32 = 797
    fmt.Println(v4, v2, v17)
    switch v30 {
    case "a":
        for fi32 := 0; fi32 < 8; fi32++ {
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(fi32, v2)
                _ = fi33
            }
            _ = fi32
        }
    case "":
        v3 = v3
    }
    if true {
        switch (v17 | (-(v17))) {
        case 127:
            var v34 bool = (!true)
            _ = v34
        case 826:
            switch 1 {
            case -517:
                v35 := GenS4{F0: v4, F1: v3}
                _ = v35
            case 128:
                v36 := GenS3{F0: v31, F1: -961, F2: v29}
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            if (4294967295 == 4294967295) {
                var v38 uint8 = v3
                switch "foo bar" {
                case "":
                    _ = len(v1)
                case "a":
                    var v39 uint16 = 0
                    _ = v39
                case "foo bar":
                    var v40 string = ""
                    _ = v40
                }
                switch 42 {
                case 127:
                    var v41 string = "hello"
                    _ = v41
                case 32767:
                    fmt.Println(fi37, fi37)
                case 0:
                    if true {
                        fmt.Println(v17, v4, v31)
                        v42 := GenS3{F0: (v31 ^ 2147483647), F1: -628, F2: (-(v29))}
                        _ = v42
                    }
                }
                _ = v38
            } else {
                var v43 float64 = (v2 - v2)
                v44 := GenS4{F0: v4, F1: v3}
                _ = v43
                _ = v44
            }
            _ = fi37
        }
    }
    v45 := GenS5{F0: 1, F1: v30, F2: false}
    var v46 int16 = 1
    var v47 uint64 = 4294967295
    if true {
        var v48 int8 = (v29 >> ((288) & 63))
        switch v30 {
        case "a":
            switch v17 {
            case 102:
                switch v30 {
                case "hello":
                    for fi49 := 0; fi49 < 1; fi49++ {
                        if true {
                            if ((v29 < 1) && (0.0 < 1e10)) {
                                var v50 int32 = -17
                                _ = v50
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v46 = 398
                        switch (v30 + (("" + (("hello" + (("a" + ("foo bar" + ("foo bar" + ((v30 + "foo bar") + (v30 + (v30 + ((v30 + (v30 + "")) + "hello"))))))) + (v30 + (v30 + v30)))) + (v30 + (v30 + ((v30 + v30) + v30))))) + "")) {
                        case "a":
                            if true {
                                var v51 uint64 = 42
                                if false {
                                    v52 := GenS1{F0: v30}
                                    if ((v29 * (^(v29))) > (v48 - (v29 + (v29 ^ (v48 >> ((637) & 63)))))) {
                                        for fi53 := 0; fi53 < 10; fi53++ {
                                            _ = len(v1)
                                            _ = fi53
                                        }
                                    } else {
                                        switch v30 {
                                        case "":
                                            var v54 uint16 = 42
                                            _ = v54
                                        case "foo bar":
                                            v55 := GenS1{F0: ("\n\t" + "")}
                                            _ = v55
                                        }
                                    }
                                    _ = v52
                                }
                                _ = v51
                            } else {
                                switch -9223372036854775808 {
                                case 128:
                                    v1 = append(v1, v30)
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "hello":
                            var v56 uint32 = v31
                            _ = v56
                        default:
                            for fi57 := 0; fi57 < 7; fi57++ {
                                if (v46 > v46) {
                                    fmt.Println(v4)
                                    var v58 int32 = -103
                                    var v59 uint16 = 65535
                                    v29 -= (v48 << ((887) & 63))
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 uint = 82
                                    _ = v60
                                }
                                _ = fi57
                            }
                        }
                        _ = fi49
                    }
                default:
                    _ = len(v1)
                }
            default:
                v30 = "\n\t"
            }
        default:
            fmt.Println(v3)
        }
        v61 := GenS3{F0: (v31 * (v31 * v31)), F1: 2147483647, F2: 70}
        _ = v48
        _ = v61
    }
    v62 := GenS3{F0: v31, F1: 2147483647, F2: int8(v3)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v62
}
