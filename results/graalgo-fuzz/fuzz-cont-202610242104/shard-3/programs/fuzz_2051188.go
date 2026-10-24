package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool, p1 uint8) string {
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    var v6 float32 = 0.0
    fmt.Println(p1, v6, v6)
    _ = p0
    _ = p1
    _ = v6
    return "a"
}

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 128
    v3 := []uint32{572}
    v4 := make(map[uint32]int)
    v5 := []int{32767, 2147483647, -2147483648}
    fmt.Println(v1, v1)
    var v7 float32 = (float32((1.5 * 1e10)) * 0.1)
    var v8 int = 42
    var v9 float64 = ((3.14159 + float64(0.1)) * 0.1)
    if false {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: true, F1: 0, F2: 2}
            _ = fi10
            _ = v11
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1)
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v8 = fi12
            fmt.Println(fi12, v8)
            _ = fi12
        }
    }
    var v13 uint8 = (v2 / (((v2 % ((v2) | 1))) | 1))
    v14 := GenS1{F0: (!true), F1: 1, F2: 2}
    if ((-1 > 42) && true) {
        v8 = v8
        switch 0 {
        case -9223372036854775808:
            var v15 uint8 = v13
            _ = v15
        case -33:
            switch "" {
            case "hello":
                v9 = (0.0 + float64(1.5))
            case "":
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS1{F0: false, F1: 743, F2: 4294967295}
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        switch 9223372036854775807 {
                        case 0:
                            var v18 int64 = -9223372036854775808
                            _ = v18
                        }
                    case -432:
                        if false {
                            var v19 bool = ((!(v7 == (v7 - v7))) && false)
                            if true {
                                v20 := GenS1{F0: ((!(42 < 42)) || (!(42 <= 0))), F1: 65535, F2: 1}
                                _ = v20
                            }
                            _ = v19
                        }
                    default:
                        v8 = v8
                    }
                    _ = fi16
                    _ = v17
                }
            case "a":
                if (2147483647 != 0) {
                    _ = len(v4)
                    var v21 int8 = 55
                    if (1000 > -323) {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            if (42 != (v2 / ((0) | 1))) {
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    v13 = 128
                                    _ = fi23
                                }
                                fmt.Println(fi22, fi22, v2)
                            } else {
                                v24 := GenS2{F0: 1, F1: "a"}
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    v9 = (v9 * v9)
                                    _ = fi25
                                }
                                _ = v24
                            }
                            v26 := GenS2{F0: 32767, F1: ((("hello" + (("hello" + (((("hello" + "a") + "foo bar") + "foo bar") + "a")) + ("a" + ("\n\t" + "hello")))) + "") + "")}
                            _ = fi22
                            _ = v26
                        }
                        switch "\n\t" {
                        case "hello":
                            var v27 int32 = int32(v1)
                            _ = v27
                        case "\n\t":
                            var v28 string = "hello"
                            _ = v28
                        }
                    }
                    _ = v21
                } else {
                    v13 += (v2 | (^(v2)))
                }
            }
        default:
            if (1 != -496) {
                if (v9 > 0.0) {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        switch 128 {
                        case -942:
                            v30 := GenS2{F0: 255, F1: "a"}
                            _ = v30
                        case -2147483648:
                            if (-2.25 != 3.14159) {
                                switch 616 {
                                case 2147483647:
                                    v1 -= (v1 + v1)
                                case 127:
                                    switch (("foo bar" + "hello") + "hello") {
                                    case "":
                                        fmt.Println(v8, v7)
                                    default:
                                        if false {
                                            v9 += v9
                                            switch (((("" + "a") + "\n\t") + ("" + ((("\n\t" + "hello") + ("a" + "a")) + ("hello" + "hello")))) + "foo bar") {
                                            case "":
                                                var v31 uint8 = v13
                                                _ = v31
                                            default:
                                                for fi32 := 0; fi32 < 16; fi32++ {
                                                    var v33 int64 = 997
                                                    switch "" {
                                                    case "hello":
                                                        v34 := GenS3{F0: "", F1: (v13 ^ (^(v2))), F2: 32767}
                                                        _ = v34
                                                    case "":
                                                        v7 -= (0.0 + 1e10)
                                                    }
                                                    _ = fi32
                                                    _ = v33
                                                }
                                            }
                                            if false {
                                                fmt.Println(v2, v2)
                                                for fi35 := 0; fi35 < 14; fi35++ {
                                                    v1 = v1
                                                    if true {
                                                        v13 -= (v13 >> ((255) & 63))
                                                    }
                                                    fmt.Println(v13)
                                                    _ = fi35
                                                }
                                                for fi36 := 0; fi36 < 1; fi36++ {
                                                    v7 += v7
                                                    _ = fi36
                                                }
                                            }
                                        } else {
                                            fmt.Println(v9)
                                            var v37 string = (("" + "") + ((("foo bar" + "hello") + "hello") + "foo bar"))
                                            _ = v37
                                        }
                                    }
                                }
                                var v38 uint8 = (v13 ^ 117)
                                fmt.Println(v7, v8)
                                _ = v38
                            } else {
                                fmt.Println(v7, v13)
                            }
                        }
                        _ = fi29
                    }
                }
            } else {
                v39 := GenS2{F0: 1, F1: "foo bar"}
                v4[1] = v8
                if false {
                    v40 := GenS2{F0: 32767, F1: ""}
                    if (true && (!true)) {
                        v41 := GenS3{F0: "", F1: v2, F2: 65535}
                        switch (v8 | 32768) {
                        case 32768:
                            v9 = float64((v7 - v7))
                        case 4294967295:
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                        }
                        _ = v41
                    } else {
                        v13 = v13
                        for fi42 := 0; fi42 < 13; fi42++ {
                            fmt.Println(fi42, v7)
                            fmt.Println(fi42)
                            _ = fi42
                        }
                    }
                    _ = v40
                } else {
                    delete(v4, 0)
                }
                _ = v39
            }
        }
        var v43 int16 = 1000
        _ = v43
    }
    v44 := GenS4{F0: 2}
    if true {
        if true {
            for fi45 := 0; fi45 < 7; fi45++ {
                v9 = float64(v7)
                v46 := GenS5{F0: 255, F1: 0, F2: v2}
                _ = fi45
                _ = v46
            }
        } else {
            var v47 int64 = -9223372036854775808
            var v48 int8 = -45
            _ = v47
            _ = v48
        }
        if (-1 > 1) {
            var v49 int64 = 454
            fmt.Println(v8)
            v50 := GenS6{F0: ((("a" + ("" + "\n\t")) < ("\n\t" + ("hello" + ""))) || true), F1: 835}
            _ = v49
            _ = v50
        }
        var v51 int64 = 0
        _ = v51
    } else {
        v3 = append(v3, 2147483648)
    }
    v7 = (1.5 / 0.1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        delete(v4, 2147483647)
        if false {
            v1 = 547
            fmt.Println(v7)
            switch v8 {
            case 1:
                var v52 uint16 = 390
                _ = v52
            default:
                var v53 int8 = -128
                _ = v53
            }
        }
    } else {
        var v54 int = v8
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4[1] = v54
        _ = v54
    }
    if false {
        v55 := GenS3{F0: "a", F1: v2, F2: 2}
        var v56 float32 = float32((v9 + v9))
        _ = v55
        _ = v56
    } else {
        v13 = v13
        v7 = (0.1 / (v7 - (v7 - v7)))
        fmt.Println(v9, v7, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
