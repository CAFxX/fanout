package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint64; F2 int64 }
func genFn0(p0 uint64, p1 uint16) uint64 {
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: 1, F1: 2}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS3{F0: 249}
        var v11 uint = 4294967295
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p1, p0)
    var v12 int16 = -32768
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return (p0 / (((p0 << ((103) & 63))) | 1))
}

func main() {
    v1 := []string{"", ("hello" + "a"), (("\n\t" + "") + "a")}
    v2 := make(map[int]bool)
    var v3 uint32 = 65535
    var v4 float64 = 3.14159
    v5 := make(map[int]bool)
    v6 := []string{"foo bar", (("" + ("foo bar" + (("" + "") + ("" + "foo bar")))) + ("" + "a"))}
    if false {
        if (true || false) {
            var v13 string = ("" + "\n\t")
            var v14 int16 = -696
            v14 = (v14 % (((v14 % ((int16(v3)) | 1))) | 1))
            _ = v13
            _ = v14
        } else {
            v3 = 42
        }
        switch (("" + "hello") + ("" + "hello")) {
        case "foo bar":
            v15 := GenS1{F0: 18446744073709551615}
            _ = v15
        }
        if (!true) {
            var v16 int16 = 1000
            if (!(true && (-9223372036854775808 <= 42))) {
                var v17 int32 = -2147483648
                _ = v17
            }
            _ = v16
        } else {
            if ((!true) && (((false && true) && false) && (82 <= -128))) {
                if false {
                    v18 := GenS3{F0: 1}
                    if (-1 < 100) {
                        v1 = append(v1, ("\n\t" + ""))
                        var v19 uint64 = 42
                        fmt.Println(v19)
                        _ = v19
                    }
                    var v20 uint8 = 0
                    _ = v18
                    _ = v20
                }
            } else {
                switch 9223372036854775807 {
                case 488:
                    v3 = v3
                default:
                    fmt.Println(v4)
                }
            }
        }
    }
    fmt.Println(v4)
    for fi21 := 0; fi21 < 13; fi21++ {
        v4 += 1.5
        _ = fi21
    }
    switch "a" {
    case "hello":
        v22 := GenS4{F0: 499}
        _ = v22
    case "":
        v3 -= v3
    case "\n\t":
        delete(v5, -1)
    }
    v1 = append(v1, "a")
    if true {
        fmt.Println(v4, v4)
        var v23 int64 = -884
        v24 := GenS4{F0: 535}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := GenS5{F0: 4294967295, F1: (v3 / (((v3 - (^(v3)))) | 1))}
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(fi25, v3)
            switch "a" {
            case "\n\t":
                delete(v5, fi27)
            case "":
                if false {
                    var v28 int32 = 2147483647
                    var v29 uint16 = 65535
                    var v30 float64 = v4
                    _ = len(v5)
                    _ = v28
                    _ = v29
                    _ = v30
                }
            case "foo bar":
                if ((("\n\t" != "foo bar") || false) || (v4 <= 0.0)) {
                    v3 = (v3 * v3)
                }
            }
            switch "hello" {
            case "\n\t":
                v3 += (v3 >> ((0) & 63))
            case "a":
                if false {
                    fmt.Println(v4, v3)
                    v4 = (v4 - (-2.25 / ((-2.25 * v4) / ((((v4 - float64((((-2.25 * 3.14159) - ((0.1 + (3.14159 / 1e10)) - 1e10)) - 0.1))) + v4) - 3.14159) + ((0.1 / v4) - v4)))))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "foo bar":
                v31 := GenS6{F0: 1000, F1: 18446744073709551615, F2: 0}
                if false {
                    var v32 int64 = -301
                    _ = v32
                }
                switch fi25 {
                case 876:
                    fmt.Println(v3)
                default:
                    var v33 uint64 = 4294967295
                    _ = v33
                }
                _ = v31
            }
            _ = fi27
        }
        for fi34 := 0; fi34 < 13; fi34++ {
            v4 = -2.25
            _ = len(v2)
            var v35 float64 = v4
            _ = fi34
            _ = v35
        }
        _ = fi25
        _ = v26
    }
    v3 = (v3 + (v3 | v3))
    if true {
        v3 -= 65535
        var v36 uint16 = 2
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 5; fi38++ {
                switch fi38 {
                case 2147483647:
                    var v39 uint64 = 0
                    _ = v39
                case -9223372036854775808:
                    delete(v5, fi37)
                case 4294967295:
                    switch fi38 {
                    case 128:
                        if ((((-2.25 + float32(v4)) * (3.14159 * (1.5 - 0.0))) - 3.14159) >= ((0.0 * 1e10) * ((1.5 * -2.25) * (-2.25 * 0.1)))) {
                            if (-1 != 426) {
                                if (271 == 0) {
                                    v36 = (v36 + (v36 & (^(v36))))
                                    var v40 uint32 = 657
                                    _ = v40
                                } else {
                                    _ = len(v2)
                                }
                            } else {
                                var v41 int8 = 100
                                _ = v41
                            }
                        }
                    default:
                        switch 517 {
                        case 1:
                            v42 := v2[fi38]
                            _ = v42
                        case -1:
                            var v43 int8 = 59
                            _ = v43
                        case 2147483647:
                            var v44 uint = 66
                            _ = v44
                        }
                    }
                default:
                    v36 += v36
                }
                _ = fi38
            }
            _ = fi37
        }
        _ = v36
    } else {
        for fi45 := 0; fi45 < 2; fi45++ {
            if false {
                switch ("hello" + (("\n\t" + "") + (("hello" + "") + "foo bar"))) {
                case "":
                    v3 += v3
                case "hello":
                    v4 += v4
                case "a":
                    if false {
                        var v46 int32 = 2147483647
                        var v47 uint = 1
                        _ = v46
                        _ = v47
                    }
                }
                fmt.Println(fi45, v3)
            }
            for fi48 := 0; fi48 < 6; fi48++ {
                v3 = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (true || true) {
                    v3 += 521
                    v49 := GenS6{F0: 32767, F1: 497, F2: 1}
                    v50 := GenS3{F0: 42}
                    _ = v49
                    _ = v50
                }
                _ = fi48
            }
            _ = fi45
        }
    }
    var v51 uint16 = 585
    v52 := GenS5{F0: 437, F1: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
