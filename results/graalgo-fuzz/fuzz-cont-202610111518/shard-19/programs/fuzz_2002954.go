package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int16; F1 string }
func genFn0() int64 {
    var v4 string = "foo bar"
    v4 = ("foo bar" + v4)
    var v5 int8 = 100
    v5 = (^(v5))
    if true {
        if (1.5 != 1e10) {
            v5 = 1
        }
        if (false && false) {
            var v6 int32 = int32(v5)
            if (1e10 == 1e10) {
                var v7 int8 = 100
                for fi8 := 0; fi8 < 17; fi8++ {
                    v5 -= (v5 ^ 100)
                    for fi9 := 0; fi9 < 5; fi9++ {
                        v10 := []int{fi9, -975, (^(fi9))}
                        switch (fi8 >> ((4294967295) & 63)) {
                        case 928:
                            if (float64(3.14159) == 1e10) {
                                v6 = (v6 % (((v6 << ((265) & 63))) | 1))
                            }
                        case 128:
                            var v11 uint32 = 2
                            _ = v11
                        default:
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                    v12 := GenS1{F0: true, F1: false, F2: 1}
                    _ = fi8
                    _ = v12
                }
                v13 := []int{4294967295, -382}
                _ = v13
                _ = v7
            }
            var v14 uint8 = 0
            _ = v14
            _ = v6
        }
    } else {
        v4 = ((v4 + v4) + "\n\t")
        if (!true) {
            switch v4 {
            case "":
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := []int{}
                    _ = fi15
                    _ = v16
                }
            case "foo bar":
                v4 = (v4 + (v4 + v4))
            case "\n\t":
                v4 = ""
            }
        }
    }
    _ = v4
    _ = v5
    return -1
}
func genFn1() string {
    v17 := []int{9223372036854775807}
    if ("hello" == "a") {
        switch 32768 {
        case -458:
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 int32 = 1
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = fi18
                _ = v19
            }
        case 0:
            for fi20 := 0; fi20 < 14; fi20++ {
                if len(v17) > 0 {
                    v17[0] = (-(fi20))
                }
                var v21 uint64 = 1
                if (!true) {
                    _ = len(v17)
                    if false {
                        fmt.Println(v21)
                        v22 := GenS1{F0: (32767 > 65535), F1: (v21 <= v21), F2: 1}
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: v21, F1: 107, F2: "hello"}
                    var v24 string = "foo bar"
                    _ = v23
                    _ = v24
                }
                _ = fi20
                _ = v21
            }
        }
        var v25 float32 = (((-2.25 / 1.0) / 1.5) / 1.0)
        var v26 uint8 = 1
        _ = v25
        _ = v26
    }
    var v27 uint16 = 729
    _ = v17
    _ = v27
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -190
    var v3 uint32 = 827
    var v28 int8 = -19
    v3 = v3
    fmt.Println(v2)
    if true {
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    v3 += 1
    var v29 uint32 = 42
    switch ("foo bar" + "hello") {
    case "hello":
        var v30 int16 = 968
        _ = v30
    case "a":
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 float32 = 1e10
            _ = fi31
            _ = v32
        }
    case "foo bar":
        for fi33 := 0; fi33 < 15; fi33++ {
            if ((0.0 / 1.0) != 0.0) {
                v34 := []int{128, fi33, fi33}
                for fi35 := 0; fi35 < 18; fi35++ {
                    if (590 == v29) {
                        var v36 uint64 = 0
                        var v37 uint64 = v36
                        var v38 uint8 = 128
                        switch -1 {
                        case 0:
                            switch ("\n\t" + "hello") {
                            case "a":
                                switch "" {
                                case "foo bar":
                                    var v39 float32 = 0.0
                                    _ = v39
                                case "a":
                                    var v40 int64 = (v2 + v2)
                                    _ = v40
                                }
                                v2 = 525
                                v37 += (v37 / ((v36) | 1))
                            case "hello":
                                switch fi33 {
                                case 271:
                                    v41 := GenS2{F0: (v37 >> ((2) & 63)), F1: 42, F2: "foo bar"}
                                    _ = v41
                                case 9223372036854775807:
                                    v42 := GenS3{F0: 0}
                                    _ = v42
                                case 0:
                                    v43 := GenS4{F0: 42, F1: v29}
                                    _ = v43
                                }
                            }
                        case 42:
                            v29 += v3
                        case 32768:
                            if false {
                                v2 -= v2
                                var v44 float32 = 1e10
                                _ = v44
                            } else {
                                v3 -= 299
                                switch ("hello" + "") {
                                case "hello":
                                    v28 -= v28
                                case "a":
                                    if true {
                                        var v45 uint16 = uint16(v36)
                                        var v46 int16 = -32768
                                        _ = v45
                                        _ = v46
                                    } else {
                                        v36 += (v36 - v37)
                                    }
                                default:
                                    v47 := GenS2{F0: (v37 % ((4294967295) | 1)), F1: 42, F2: "\n\t"}
                                    _ = v47
                                }
                                switch fi35 {
                                case 128:
                                    v37 += v37
                                default:
                                    if len(v34) >= 2 {
                                        v34 = v34[1:len(v34)]
                                    }
                                }
                                _ = len(v1)
                            }
                        default:
                            v29 += 473
                        }
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v48 int64 = (v2 << ((0) & 63))
                    _ = fi35
                    _ = v48
                }
                var v49 uint8 = 128
                _ = v34
                _ = v49
            } else {
                var v50 uint32 = (v29 - v29)
                _ = v50
            }
            _ = fi33
        }
    }
    if false {
        if false {
            v2 += v2
            v51 := GenS5{F0: -1, F1: (("" + ((("\n\t" + ("a" + ("a" + "foo bar"))) + "a") + "hello")) + "a")}
            fmt.Println(v28, v3, v28)
            _ = v51
        } else {
            v1[0] = 629
        }
    }
    var v52 uint64 = 18446744073709551615
    switch 2147483647 {
    case 32767:
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 bool = true
            _ = fi53
            _ = v54
        }
    case 853:
        if false {
            v29 = 0
        }
    case 90:
        var v55 int16 = 0
        _ = v55
    default:
        v56 := GenS2{F0: 4294967295, F1: -48, F2: ""}
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v52
}
