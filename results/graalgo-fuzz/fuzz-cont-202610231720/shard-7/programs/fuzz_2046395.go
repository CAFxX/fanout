package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
func genFn0(p0 uint16, p1 int, p2 int16) uint {
    for fi5 := 0; fi5 < 9; fi5++ {
        if false {
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 float32 = 0.1
                var v8 int64 = 42
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := []int{873}
            _ = v9
        }
        _ = fi5
    }
    p1 -= 0
    v10 := []int{p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 uint8, p1 uint8) int {
    v11 := []int{2147483647, 128, -9223372036854775808, -73}
    if (3.14159 <= (3.14159 - 0.0)) {
        p0 -= (^(p0))
        var v12 uint8 = p1
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint16 = 0
        if (!false) {
            if true {
                if (("\n\t" + "hello") >= "hello") {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 uint32 = 4294967295
                        _ = fi15
                        _ = v16
                    }
                    var v17 int16 = 1
                    _ = v17
                } else {
                    _ = len(v11)
                    v18 := GenS1{F0: (fi13 / ((fi13) | 1)), F1: -128}
                    _ = v18
                }
                for fi19 := 0; fi19 < 3; fi19++ {
                    p0 = p0
                    p1 = (p1 * (p0 % ((1) | 1)))
                    v20 := GenS2{F0: 42, F1: 1}
                    _ = fi19
                    _ = v20
                }
            } else {
                switch "\n\t" {
                case "":
                    var v21 string = ("hello" + "hello")
                    _ = v21
                case "a":
                    var v22 int32 = -587
                    _ = v22
                case "\n\t":
                    v23 := GenS3{F0: fi13}
                    _ = v23
                default:
                    switch (fi13 + 32768) {
                    case -1:
                        if (42 == 42) {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v25 := GenS3{F0: 1}
                                if len(v11) > 0 {
                                    v11[0] = fi13
                                }
                                var v26 bool = (false && (false && (false || (1 < 42))))
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            var v27 int32 = 42
                            p1 -= p1
                            _ = v27
                        } else {
                            var v28 float64 = (0.1 + 0.0)
                            var v29 uint64 = 1
                            _ = v28
                            _ = v29
                        }
                    case -2147483648:
                        switch "\n\t" {
                        case "a":
                            var v30 uint = 265
                            _ = v30
                        case "":
                            if len(v11) > 0 {
                                v11[0] = fi13
                            }
                        case "\n\t":
                            fmt.Println(p1, p0)
                        }
                    case 4294967295:
                        v14 += (v14 / ((v14) | 1))
                    }
                }
                var v31 int16 = 32767
                _ = v31
            }
            if ((3.14159 >= 1e10) && true) {
                fmt.Println(p0, p1, p0)
                var v32 int64 = 42
                _ = v32
            } else {
                var v33 int8 = -1
                for fi34 := 0; fi34 < 10; fi34++ {
                    v35 := GenS3{F0: fi13}
                    var v36 uint16 = 2
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = v33
            }
            if true {
                var v37 float64 = float64(0.1)
                v11 = append(v11, fi13)
                switch fi13 {
                case -9223372036854775808:
                    if (!false) {
                        var v38 uint64 = 403
                        var v39 float64 = v37
                        switch "foo bar" {
                        case "":
                            fmt.Println(fi13, p0, v37)
                        default:
                            p1 = 127
                        }
                        _ = v38
                        _ = v39
                    } else {
                        var v40 bool = (true && false)
                        if len(v11) > 0 {
                            v11[0] = (^(fi13))
                        }
                        _ = v40
                    }
                }
                _ = v37
            } else {
                v14 += v14
            }
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    var v1 bool = (!true)
    var v2 uint8 = 2
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    v41 := []int{32768, 643, 2147483647}
    v2 = v2
    v41 = append(v41, 2147483647)
    v3[1] = false
    var v42 string = ""
    var v43 float64 = ((0.1 * 0.1) + (-2.25 * 1.5))
    fmt.Println(v43, v43)
    for fi44 := 0; fi44 < 11; fi44++ {
        fmt.Println(v1, v1, fi44)
        _ = fi44
    }
    _ = len(v3)
    v2 = (v2 >> ((65535) & 63))
    var v45 bool = (v1 && (!v1))
    var v46 int = -2147483648
    if v1 {
        v41 = append(v41, genFn1(201, (v2 % (((v2 % ((151) | 1))) | 1))))
        fmt.Println(v45, v42, v46)
    } else {
        v1 = v45
        switch (v42 + (v42 + (("foo bar" + ("a" + "foo bar")) + "\n\t"))) {
        case "\n\t":
            for fi47 := 0; fi47 < 15; fi47++ {
                switch (v46 / ((genFn1((v2 & v2), (v2 + 96))) | 1)) {
                case -878:
                    switch (v42 + "") {
                    case "foo bar":
                        switch -459 {
                        case 0:
                            _ = len(v4)
                        case 1:
                            _ = len(v4)
                        case 393:
                            switch 0 {
                            case 2147483647:
                                for fi48 := 0; fi48 < 14; fi48++ {
                                    if false {
                                        v41 = append(v41, genFn1(0, v2))
                                        v49 := GenS4{F0: 1, F1: v42, F2: -32768}
                                        _ = v49
                                    } else {
                                        v50 := v3[42]
                                        _ = v50
                                    }
                                    v51 := GenS3{F0: genFn1(0, 42)}
                                    _ = fi48
                                    _ = v51
                                }
                            case -798:
                                var v52 int8 = 127
                                _ = v52
                            default:
                                fmt.Println(v1, v1)
                            }
                        default:
                            for fi53 := 0; fi53 < 13; fi53++ {
                                var v54 float64 = 0.0
                                if (589 != -160) {
                                    if false {
                                        fmt.Println(fi53, v2, fi47)
                                        fmt.Println(fi53, v2)
                                    }
                                    v55 := GenS1{F0: fi53, F1: -1}
                                    _ = v55
                                }
                                _ = fi53
                                _ = v54
                            }
                        }
                    case "\n\t":
                        var v56 bool = false
                        _ = v56
                    case "hello":
                        v1 = (false || v1)
                    default:
                        v42 = ("hello" + ((("\n\t" + "") + "") + v42))
                    }
                case 856:
                    v45 = false
                case 0:
                    var v57 float64 = (3.14159 * v43)
                    _ = v57
                default:
                    v1 = v1
                }
                var v58 float32 = 0.0
                _ = fi47
                _ = v58
            }
        case "":
            v59 := GenS2{F0: 42, F1: 42}
            _ = v59
        default:
            var v60 int64 = -1
            _ = v60
        }
    }
    v61 := GenS1{F0: (v46 | v46), F1: -128}
    var v62 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v61
    _ = v62
}
