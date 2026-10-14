package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool; F1 string; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 uint64; F1 int }
func genFn0(p0 uint8, p1 bool, p2 uint16) int32 {
    v6 := GenS1{F0: uint32(p2), F1: "a"}
    p0 = (p0 & p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint8, p1 uint8, p2 uint64) int64 {
    fmt.Println(p1, p0)
    p1 += (p0 >> ((618) & 63))
    var v7 int32 = 0
    if false {
        v8 := GenS1{F0: 599, F1: ("foo bar" + "a")}
        v9 := GenS1{F0: 224, F1: ""}
        _ = v8
        _ = v9
    } else {
        var v10 bool = (0.1 == 0.1)
        switch int(p0) {
        case 4294967295:
            if (!((("hello" + ("a" + "hello")) <= "\n\t") && v10)) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 uint = 18446744073709551615
                    _ = fi11
                    _ = v12
                }
            } else {
                v13 := []int{32768}
                _ = v13
            }
        case -576:
            fmt.Println(p0)
        case 128:
            switch -9223372036854775808 {
            case 1:
                p2 -= (p2 << ((834) & 63))
            default:
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v10, v7)
                    _ = fi14
                }
            }
        }
        _ = v10
    }
    if (42 == 2147483648) {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!true) {
                v16 := []int{fi15}
                switch (fi15 - -2147483648) {
                case -2147483648:
                    v17 := GenS2{F0: uint32(p2), F1: "", F2: "hello"}
                    _ = v17
                case -9223372036854775808:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 float32 = (0.0 / (1.5 - 0.0))
                        _ = fi18
                        _ = v19
                    }
                case 9223372036854775807:
                    if (false && true) {
                        fmt.Println(p1)
                        p0 = p1
                    }
                }
                p0 -= (p0 << ((18446744073709551615) & 63))
                _ = v16
            } else {
                v20 := []int{-415, 32768, (fi15 >> ((161) & 63)), fi15}
                if ((!false) && true) {
                    if len(v20) > 0 {
                        v20[0] = fi15
                    }
                }
                switch ("a" + "\n\t") {
                case "\n\t":
                    v21 := GenS1{F0: 2147483647, F1: ""}
                    _ = v21
                case "":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        p2 = 18446744073709551615
                        var v23 int32 = 392
                        _ = fi22
                        _ = v23
                    }
                default:
                    v24 := GenS3{F0: -9223372036854775808, F1: 341}
                    _ = v24
                }
                _ = v20
            }
            var v25 int8 = -8
            p2 -= p2
            _ = fi15
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 86
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 uint8 = 1
    var v5 int64 = 9223372036854775807
    v5 += 9223372036854775807
    delete(v2, "foo bar")
    v26 := []int{9223372036854775807}
    if (!false) {
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                if false {
                    var v29 bool = true
                    var v30 string = "hello"
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS4{F0: true, F1: "", F2: 18446744073709551615}
                    fmt.Println(v5, fi28)
                    var v32 float32 = 3.14159
                    _ = v31
                    _ = v32
                }
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi28
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                v1 = v4
                var v34 int = fi33
                v1 -= (v4 ^ (v1 + 255))
                v2["a"] = "foo bar"
                _ = len(v26)
                var v35 uint64 = 2
                _ = fi33
                _ = v34
                _ = v35
            }
            v36 := GenS4{F0: false, F1: (("a" + "foo bar") + ("hello" + "hello")), F2: 4294967295}
            _ = fi27
            _ = v36
        }
        fmt.Println(v1, v4)
        var v37 uint = 42
        _ = v37
    }
    if false {
        var v38 uint64 = 18446744073709551615
        _ = v38
    } else {
        v4 -= (^(v1))
    }
    var v39 int8 = -121
    fmt.Println(v4, v5)
    switch 42 {
    case 124:
        if true {
            switch "" {
            case "\n\t":
                switch ("" + "foo bar") {
                case "hello":
                    v2["\n\t"] = ""
                case "\n\t":
                    v4 = (v4 & v1)
                case "foo bar":
                    delete(v2, ((("" + "hello") + "foo bar") + ""))
                default:
                    if false {
                        if false {
                            v4 += v4
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v41 := GenS5{F0: 32767}
                                if ((-2147483648 != 2147483647) && false) {
                                    v5 += (v5 >> ((18446744073709551615) & 63))
                                    for fi42 := 0; fi42 < 8; fi42++ {
                                        _ = len(v26)
                                        v43 := GenS6{F0: int(v4), F1: (v39 & -37)}
                                        v44 := GenS7{F0: (^(v39)), F1: (v5 * (^(v5)))}
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                    var v45 bool = true
                                    _ = v45
                                }
                                _ = fi40
                                _ = v41
                            }
                        } else {
                            if (1 <= 4294967295) {
                                for fi46 := 0; fi46 < 13; fi46++ {
                                    if true {
                                        var v47 uint = 18446744073709551615
                                        _ = v47
                                    }
                                    _ = fi46
                                }
                                var v48 uint = 481
                                v49 := GenS8{F0: 430, F1: -2147483648}
                                _ = v48
                                _ = v49
                            }
                            v39 -= -1
                            var v50 uint32 = 2147483648
                            _ = v50
                        }
                        var v51 int32 = 1
                        _ = v51
                    }
                }
            default:
                fmt.Println(v39, v4, v39)
            }
            if true {
                v5 -= (^(v5))
                for fi52 := 0; fi52 < 19; fi52++ {
                    var v53 string = ("hello" + "")
                    _ = len(v26)
                    _ = len(v2)
                    _ = fi52
                    _ = v53
                }
            } else {
                for fi54 := 0; fi54 < 4; fi54++ {
                    v26 = append(v26, (fi54 / (((fi54 & 32768)) | 1)))
                    var v55 uint32 = 4294967295
                    var v56 int64 = genFn1(v1, v1, 672)
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                var v57 int64 = genFn1((v4 << ((1) & 63)), (v4 + v1), 123)
                if (2147483647 != 977) {
                    fmt.Println(v4)
                    var v58 int32 = 0
                    delete(v3, "foo bar")
                    _ = v58
                } else {
                    if len(v26) > 0 {
                        v26[0] = 4294967295
                    }
                    v4 += 255
                }
                _ = v57
            }
        } else {
            if true {
                v5 = v5
            }
            if false {
                var v59 int32 = -244
                var v60 int16 = 42
                _ = v59
                _ = v60
            }
        }
    }
    for fi61 := 0; fi61 < 16; fi61++ {
        v1 = (v4 + 2)
        for fi62 := 0; fi62 < 9; fi62++ {
            var v63 uint32 = 378
            _ = fi62
            _ = v63
        }
        _ = fi61
    }
    v5 += (v5 * v5)
    var v64 int64 = v5
    var v65 float32 = -2.25
    var v66 uint = 1
    v5 = 536
    if true {
        var v67 uint32 = 65535
        _ = v67
    } else {
        if true {
            for fi68 := 0; fi68 < 11; fi68++ {
                delete(v3, "")
                var v69 bool = (float32(0.1) < 3.14159)
                _ = fi68
                _ = v69
            }
        }
        if len(v26) > 0 {
            v26[0] = -1
        }
    }
    var v70 int = 2147483647
    var v71 int = (v70 << (((v66 / ((v66) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v70
    _ = v71
}
