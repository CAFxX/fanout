package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint, p2 uint16) int16 {
    var v5 uint64 = 42
    v5 = 2
    switch 9223372036854775807 {
    case 32767:
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int64 = 567
            if ((p2 << (((^(p1))) & 63)) != (p2 + p2)) {
                fmt.Println(v5, fi6, v7)
            } else {
                var v8 bool = (212 == p0)
                if v8 {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        if (int64(p0) <= (v7 + 0)) {
                            var v10 uint64 = v5
                            _ = v10
                        }
                        _ = fi9
                    }
                    for fi11 := 0; fi11 < 11; fi11++ {
                        var v12 string = "foo bar"
                        v13 := GenS1{F0: 2, F1: true, F2: ("hello" + (((v12 + "a") + "\n\t") + "hello"))}
                        switch v12 {
                        case "foo bar":
                            switch v12 {
                            case "foo bar":
                                var v14 int8 = 0
                                _ = v14
                            default:
                                var v15 string = v12
                                fmt.Println(v5, fi6, fi11)
                                var v16 int16 = 334
                                _ = v15
                                _ = v16
                            }
                        case "a":
                            var v17 int8 = 42
                            _ = v17
                        case "hello":
                            v12 = (v12 + v12)
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                fmt.Println(fi6, p1)
                _ = v8
            }
            _ = fi6
            _ = v7
        }
    case 167:
        var v18 float32 = (0.1 / -2.25)
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS1{F0: p0, F1: (!true), F2: ("hello" + "foo bar")}
            if (false && false) {
                var v21 bool = false
                _ = v21
            } else {
                if true {
                    switch "" {
                    case "\n\t":
                        if (-1 <= 1) {
                            fmt.Println(p2)
                            if (true && false) {
                                v22 := GenS1{F0: 255, F1: ((fi19 >= fi19) || (false || (v5 == (v5 & 116)))), F2: ("" + "hello")}
                                var v23 int = fi19
                                _ = v22
                                _ = v23
                            }
                            if false {
                                p2 += 255
                                v24 := GenS1{F0: (p0 ^ (p0 ^ p0)), F1: true, F2: ("\n\t" + "")}
                                _ = v24
                            } else {
                                switch (("\n\t" + "\n\t") + "foo bar") {
                                case "foo bar":
                                    p2 += uint16(p0)
                                default:
                                    p2 -= (p2 / ((p2) | 1))
                                }
                            }
                        }
                    case "foo bar":
                        p2 = (p2 >> ((uint(p0)) & 63))
                    case "a":
                        var v25 uint8 = (p0 & (p0 ^ p0))
                        _ = v25
                    default:
                        fmt.Println(p1, p0, fi19)
                    }
                } else {
                    v5 = (v5 * v5)
                    p0 = p0
                    var v26 bool = false
                    _ = v26
                }
                v27 := GenS1{F0: (^(p0)), F1: true, F2: ""}
                _ = v27
            }
            v28 := []int{}
            _ = fi19
            _ = v20
            _ = v28
        }
    } else {
        var v29 int64 = 727
        if (176 > p2) {
            fmt.Println(p0)
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint32 = 2147483648
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 uint = p1
                    fmt.Println(p2)
                    v33 = v33
                    var v34 int64 = v29
                    if (float64(-2.25) > 3.14159) {
                        var v35 bool = false
                        _ = v35
                    } else {
                        switch ((("a" + ("" + "foo bar")) + (("\n\t" + ("a" + "hello")) + "\n\t")) + "foo bar") {
                        case "foo bar":
                            v36 := GenS1{F0: p0, F1: ((^(p2)) >= uint16(v33)), F2: ("foo bar" + "\n\t")}
                            _ = v36
                        case "a":
                            var v37 uint32 = 42
                            _ = v37
                        case "\n\t":
                            v38 := []int{(fi32 % ((136) | 1))}
                            _ = v38
                        }
                    }
                    v39 := GenS1{F0: (^(p0)), F1: (!false), F2: ("foo bar" + ("" + "hello"))}
                    fmt.Println(v33, p1, p1)
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v39
                }
                _ = fi30
                _ = v31
            }
            v40 := GenS1{F0: (p0 % ((p0) | 1)), F1: false, F2: ""}
            _ = v40
        } else {
            p1 += (p1 % (((p1 | p1)) | 1))
            switch -440 {
            case 384:
                v41 := []int{}
                _ = v41
            case 4294967295:
                for fi42 := 0; fi42 < 20; fi42++ {
                    v43 := []int{-608, fi42}
                    _ = fi42
                    _ = v43
                }
            case -990:
                var v44 uint32 = 1
                _ = v44
            }
        }
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 586
}

func main() {
    var v1 int8 = 0
    v2 := []string{"foo bar"}
    var v3 int8 = (v1 + 100)
    v4 := []int{636, 525, 2147483647, 377}
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 int64 = 42
        v3 -= 0
        _ = fi45
        _ = v46
    }
    switch -623 {
    case -9223372036854775808:
        v3 += v1
    case 127:
        v1 = (v3 | -128)
    default:
        v47 := GenS2{F0: -1, F1: -32768, F2: -602}
        _ = v47
    }
    v3 -= (v3 & -12)
    var v48 uint16 = 65535
    if true {
        v4 = append(v4, 964)
        v49 := GenS3{F0: v48, F1: (!(!(false && (587 < 147)))), F2: -32768}
        v3 -= (v1 >> ((42) & 63))
        _ = v49
    }
    var v50 uint32 = 0
    v51 := GenS4{F0: 839}
    switch "\n\t" {
    case "":
        v4 = append(v4, 997)
    case "\n\t":
        var v52 float32 = (0.1 * 0.1)
        _ = v52
    case "foo bar":
        for fi53 := 0; fi53 < 3; fi53++ {
            var v54 uint64 = 2
            _ = fi53
            _ = v54
        }
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v50
    _ = v51
}
