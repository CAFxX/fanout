package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int16; F2 uint32 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 uint, p2 uint32) uint64 {
    v5 := GenS1{F0: (587 > 65535), F1: p1, F2: ("foo bar" + (("" + "") + ""))}
    for fi6 := 0; fi6 < 9; fi6++ {
        if p0 {
            switch ((("foo bar" + "\n\t") + "foo bar") + "a") {
            case "foo bar":
                for fi7 := 0; fi7 < 14; fi7++ {
                    p0 = ((p0 || true) || p0)
                    _ = fi7
                }
            }
            if true {
                switch fi6 {
                case 42:
                    fmt.Println(p1, fi6, fi6)
                case -1:
                    for fi8 := 0; fi8 < 18; fi8++ {
                        if p0 {
                            v9 := GenS1{F0: p0, F1: 1, F2: "a"}
                            v10 := []int{(fi6 & fi8), -838}
                            p1 = p1
                            _ = v10
                            _ = v9
                        } else {
                            if (!true) {
                                p0 = p0
                                v11 := []int{(fi6 >> ((uint(p2)) & 63)), (fi8 << (((p1 - 1)) & 63))}
                                _ = v11
                            }
                            v12 := GenS2{F0: -128, F1: 18446744073709551615, F2: 65535}
                            _ = v12
                        }
                        var v13 bool = ("" >= (("\n\t" + ("" + "")) + "hello"))
                        _ = fi8
                        _ = v13
                    }
                case 32767:
                    if true {
                        var v14 float64 = (0.1 + (float64(((1.5 + -2.25) / 1.5)) / 1.0))
                        _ = v14
                    }
                default:
                    var v15 uint16 = 2
                    _ = v15
                }
                v16 := []int{(fi6 ^ (fi6 / ((fi6) | 1))), fi6}
                fmt.Println(p0)
                _ = v16
            } else {
                var v17 int = fi6
                _ = v17
            }
        } else {
            var v18 int8 = 81
            fmt.Println(p1, fi6)
            var v19 uint64 = 2
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: p2}
        _ = fi6
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    var v3 string = ("\n\t" + v2)
    var v4 uint = 65535
    if true {
        v4 -= (v4 * v4)
        var v21 uint64 = 4294967295
        fmt.Println(v3)
        _ = v21
    }
    v22 := []int{733, 235, 32768}
    if true {
        var v23 int64 = 0
        _ = v23
    } else {
        if true {
            v4 -= (v4 + v4)
        }
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if true {
        if true {
            fmt.Println(v3, v2)
            if ((v4 < 0) || ((-2.25 + -2.25) >= 0.1)) {
                v24 := GenS4{F0: v3, F1: -32768, F2: 0}
                var v25 uint8 = 169
                _ = v24
                _ = v25
            } else {
                v2 = ("hello" + ("foo bar" + ((v2 + ("hello" + v3)) + v3)))
                var v26 uint16 = 32767
                var v27 uint32 = 540
                _ = v26
                _ = v27
            }
            v28 := GenS5{F0: -924, F1: 1}
            _ = v28
        }
        if (-32768 <= -1) {
            fmt.Println(v3)
            switch 32767 {
            case 4294967295:
                if true {
                    if false {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v4 -= 65535
                            v22 = append(v22, fi29)
                            var v30 int32 = 404
                            _ = fi29
                            _ = v30
                        }
                        switch -210 {
                        case 9223372036854775807:
                            if (true || (1000 > 1000)) {
                                fmt.Println(v2, v3)
                            } else {
                                v2 = v3
                                if len(v22) > 0 {
                                    v22[0] = -9223372036854775808
                                }
                                var v31 int32 = int32(v4)
                                _ = v31
                            }
                        case -9223372036854775808:
                            var v32 int8 = 100
                            _ = v32
                        }
                    }
                }
            }
        } else {
            v33 := GenS6{F0: 42, F1: 349}
            _ = v33
        }
        fmt.Println(v3, v2, v2)
    } else {
        v4 = v4
        var v34 bool = (v2 < "\n\t")
        v4 = (v4 | (v4 % (((v4 | v4)) | 1)))
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int64 = -22
        if false {
            v2 = (((v3 + "hello") + (("foo bar" + v2) + v3)) + v2)
            fmt.Println(v3)
        }
        fmt.Println(v36, v4, fi35)
        _ = fi35
        _ = v36
    }
    if (100 <= 92) {
        for fi37 := 0; fi37 < 6; fi37++ {
            _ = len(v22)
            for fi38 := 0; fi38 < 15; fi38++ {
                if (758 > 406) {
                    v39 := GenS3{F0: 42}
                    var v40 float64 = float64(-2.25)
                    _ = v39
                    _ = v40
                } else {
                    _ = len(v1)
                    v41 := GenS1{F0: (-1 == int64(fi38)), F1: 255, F2: ("foo bar" + "foo bar")}
                    var v42 string = v2
                    _ = v41
                    _ = v42
                }
                v2 = v3
                if ((10 != -128) && true) {
                    if false {
                        if false {
                            _ = len(v1)
                            switch int(v4) {
                            case 32768:
                                switch fi37 {
                                case 128:
                                    switch v2 {
                                    case "":
                                        if false {
                                            var v43 float64 = 1.5
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            _ = v43
                                        } else {
                                            v44 := GenS5{F0: fi38, F1: -1}
                                            if false {
                                                v2 = v3
                                                switch (^(fi37)) {
                                                case 1:
                                                    if len(v22) > 0 {
                                                        v22[0] = fi37
                                                    }
                                                default:
                                                    fmt.Println(fi37, v2)
                                                }
                                            } else {
                                                switch (fi38 * (fi38 / ((fi37) | 1))) {
                                                case 42:
                                                    if len(v22) > 0 {
                                                        v22[0] = fi38
                                                    }
                                                case 4294967295:
                                                    var v45 string = "hello"
                                                    _ = v45
                                                case -2147483648:
                                                    var v46 string = ("a" + v2)
                                                    _ = v46
                                                default:
                                                    v4 += (v4 * (v4 + (v4 % ((v4) | 1))))
                                                }
                                            }
                                            _ = v44
                                        }
                                    case "foo bar":
                                        v4 = 2
                                    case "hello":
                                        if true {
                                            _ = len(v1)
                                            var v47 uint64 = 122
                                            v4 = (^(v4))
                                            _ = v47
                                        }
                                    default:
                                        v48 := GenS7{F0: v4, F1: (false || ((-191 >= 1) && true))}
                                        v4 += (v4 % ((572) | 1))
                                        var v49 uint16 = 65535
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            case 127:
                                var v50 int16 = 1
                                _ = v50
                            default:
                                v2 = v2
                            }
                        }
                    }
                    var v51 int = (fi37 / ((fi38) | 1))
                    _ = v51
                }
                if len(v22) > 0 {
                    v22[0] = (fi38 % ((fi37) | 1))
                }
                _ = fi38
            }
            if (714 == 42) {
                v52 := GenS1{F0: ((("foo bar" + (("foo bar" + v2) + v3)) + v2) > ("a" + v2)), F1: v4, F2: "a"}
                v4 += 608
                var v53 float64 = 0.0
                _ = v52
                _ = v53
            }
            _ = fi37
        }
        v54 := GenS2{F0: 73, F1: 0, F2: 1}
        if true {
            var v55 uint8 = 1
            _ = v55
        }
        _ = v54
    }
    var v56 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v56
}
