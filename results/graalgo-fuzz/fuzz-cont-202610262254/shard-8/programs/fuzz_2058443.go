package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32; F1 uint; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", "a", "foo bar"}
    var v3 bool = (false && (float32(-2.25) == 0.0))
    v4 := []int{}
    v5 := make(map[int]int)
    var v6 int64 = -152
    v7 := GenS1{F0: 42}
    v3 = (2147483648 < 4294967295)
    v3 = (v3 && v3)
    if v3 {
        var v8 uint = 896
        _ = v8
    } else {
        fmt.Println(v3, v6)
    }
    v6 = 297
    v9 := GenS1{F0: 42}
    switch -1 {
    case -9223372036854775808:
        var v10 int64 = 0
        _ = v10
    default:
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    fmt.Println(v6, v6)
    if false {
        fmt.Println(v6)
    } else {
        if (false || v3) {
            v4 = append(v4, 32767)
        } else {
            switch 42 {
            case 127:
                v12 := GenS1{F0: 1}
                _ = v12
            case -1:
                if false {
                    if (-2147483648 > 0) {
                        v3 = v3
                        for fi13 := 0; fi13 < 15; fi13++ {
                            for fi14 := 0; fi14 < 16; fi14++ {
                                v15 := GenS1{F0: 892}
                                _ = len(v2)
                                if v3 {
                                    v4 = append(v4, (fi14 + fi14))
                                }
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 4; fi16++ {
                                v6 = (v6 >> ((65535) & 63))
                                _ = fi16
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi13
                        }
                    }
                    var v17 int32 = 2147483647
                    _ = v17
                } else {
                    if (true || v3) {
                        for fi18 := 0; fi18 < 20; fi18++ {
                            v2 = append(v2, ("hello" + "hello"))
                            for fi19 := 0; fi19 < 9; fi19++ {
                                fmt.Println(fi19, v6, fi19)
                                v6 += (v6 * v6)
                                _ = fi19
                            }
                            v3 = v3
                            _ = fi18
                        }
                        v3 = v3
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v6 = v6
                            v3 = false
                            for fi21 := 0; fi21 < 20; fi21++ {
                                v6 = v6
                                _ = fi21
                            }
                            _ = fi20
                        }
                    } else {
                        fmt.Println(v6, v6)
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v6 += v6
                            fmt.Println(v6)
                            _ = fi22
                        }
                        v5[-2147483648] = 32767
                    }
                }
            default:
                var v23 string = "hello"
                _ = v23
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                for fi25 := 0; fi25 < 19; fi25++ {
                    fmt.Println(v6)
                    if (v3 && v3) {
                        v5[194] = fi25
                        if (((!(-128 > int8(fi25))) || false) && v3) {
                            fmt.Println(fi25)
                            if v3 {
                                fmt.Println(fi25, fi25, v3)
                                v26 := v1[fi25]
                                switch 4294967295 {
                                case 32768:
                                    v6 = v6
                                case 513:
                                    fmt.Println(v26, v26)
                                case -2147483648:
                                    v26 = 4294967295
                                }
                                _ = v26
                            } else {
                                if (65535 != 1) {
                                    v27 := GenS1{F0: 491}
                                    switch (("foo bar" + "\n\t") + ("\n\t" + ("foo bar" + (("foo bar" + "a") + "foo bar")))) {
                                    case "a":
                                        var v28 uint16 = 2
                                        _ = v28
                                    case "\n\t":
                                        var v29 int = 32768
                                        _ = v29
                                    case "hello":
                                        v30 := GenS1{F0: 2}
                                        _ = v30
                                    }
                                    fmt.Println(v3, v6, v6)
                                    _ = v27
                                }
                            }
                        }
                        if false {
                            var v31 uint16 = 42
                            v32 := GenS1{F0: 18446744073709551615}
                            _ = v31
                            _ = v32
                        } else {
                            v6 = (v6 / ((9223372036854775807) | 1))
                            var v33 float64 = (float64(1e10) + 0.1)
                            v6 = v6
                            _ = v33
                        }
                    }
                    v3 = (285 >= 255)
                    _ = fi25
                }
                delete(v5, (fi24 / (((-(fi24))) | 1)))
                for fi34 := 0; fi34 < 2; fi34++ {
                    if v3 {
                        if true {
                            fmt.Println(fi34, fi34, fi24)
                            v3 = v3
                        }
                        v35 := GenS1{F0: 2}
                        var v36 uint32 = 2147483647
                        _ = v35
                        _ = v36
                    } else {
                        if v3 {
                            switch (fi34 * (-(fi24))) {
                            case -49:
                                var v37 int = (fi24 - (-(fi34)))
                                _ = v37
                            default:
                                v38 := GenS2{F0: 173, F1: 65535}
                                _ = v38
                            }
                        }
                        switch 127 {
                        case 9223372036854775807:
                            var v39 float32 = (0.0 - 3.14159)
                            _ = v39
                        case 42:
                            v1[9223372036854775807] = 2147483647
                        }
                    }
                    v3 = v3
                    _ = fi34
                }
                _ = fi24
            }
        }
        v5[32767] = -735
    }
    var v40 float32 = float32(3.14159)
    var v41 uint = 17
    if v3 {
        if (0.0 != 1e10) {
            if v3 {
                v40 = v40
                if (!v3) {
                    _ = len(v5)
                    v6 = 124
                    fmt.Println(v41, v6, v40)
                } else {
                    v6 = v6
                    if (199 == 1) {
                        var v42 int8 = 123
                        if (v3 && true) {
                            fmt.Println(v3, v6, v40)
                        } else {
                            v43 := GenS2{F0: 42, F1: 630}
                            _ = v43
                        }
                        switch ((("a" + (("foo bar" + (((("foo bar" + "foo bar") + (("foo bar" + "\n\t") + ("\n\t" + ("a" + ("hello" + "\n\t"))))) + "\n\t") + "foo bar")) + "")) + "hello") + "hello") {
                        case "hello":
                            fmt.Println(v41, v3)
                        }
                        _ = v42
                    } else {
                        var v44 uint64 = 0
                        _ = v44
                    }
                }
                fmt.Println(v6)
            }
        } else {
            v45 := GenS2{F0: 38, F1: 947}
            var v46 int8 = 127
            _ = v45
            _ = v46
        }
        if (!(0 < 42)) {
            fmt.Println(v41)
            _ = len(v2)
            switch 907 {
            case -1:
                v41 = v41
            case -9223372036854775808:
                for fi47 := 0; fi47 < 10; fi47++ {
                    v4 = append(v4, (fi47 - -2147483648))
                    _ = fi47
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = (("\n\t" + "") + "a")
            }
            v48 := GenS1{F0: 42}
            _ = v48
        }
    } else {
        switch 903 {
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -73:
            v6 -= v6
        default:
            fmt.Println(v40, v40)
        }
        if true {
            var v49 float64 = 3.14159
            if v3 {
                if v3 {
                    v50 := GenS3{F0: -784, F1: (v41 - (v41 | (v41 * (v41 % ((v41) | 1)))))}
                    _ = v50
                } else {
                    if v3 {
                        v6 -= v6
                        v51 := GenS1{F0: 2}
                        _ = v51
                    }
                    v41 += v41
                }
                if ((v3 && v3) && false) {
                    if len(v2) > 0 {
                        v2[0] = ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + ("a" + "foo bar"))
                    }
                } else {
                    var v52 uint64 = 42
                    _ = v52
                }
            } else {
                var v53 string = "foo bar"
                for fi54 := 0; fi54 < 2; fi54++ {
                    v1[(-(fi54))] = 942
                    _ = fi54
                }
                var v55 int16 = 32767
                _ = v53
                _ = v55
            }
            _ = v49
        } else {
            for fi56 := 0; fi56 < 20; fi56++ {
                for fi57 := 0; fi57 < 16; fi57++ {
                    if (!(float64(v40) <= -2.25)) {
                        var v58 int32 = -1
                        _ = v58
                    }
                    v59 := GenS4{F0: 2, F1: (v41 % ((v41) | 1)), F2: v3}
                    v40 -= v40
                    _ = fi57
                    _ = v59
                }
                _ = fi56
            }
            v60 := GenS1{F0: 42}
            _ = v60
        }
        switch 2147483647 {
        case 9223372036854775807:
            for fi61 := 0; fi61 < 1; fi61++ {
                var v62 float64 = 1e10
                _ = fi61
                _ = v62
            }
        case 42:
            v41 += 42
        }
    }
    v41 -= v41
    if v3 {
        var v63 uint64 = uint64(v41)
        fmt.Println(v40, v6)
        v6 -= 42
        _ = v63
    } else {
        v40 += v40
    }
    v64 := GenS2{F0: 457, F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v64
    _ = v7
    _ = v9
}
