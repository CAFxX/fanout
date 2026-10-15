package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint; F1 uint64; F2 int }
func genFn0(p0 uint64, p1 int64, p2 uint32) string {
    v7 := GenS1{F0: "a", F1: (true || ("" != "")), F2: 0}
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("" + "foo bar")
}
func genFn1() string {
    var v9 float64 = 3.14159
    v10 := GenS1{F0: "a", F1: false, F2: -1}
    fmt.Println(v9)
    fmt.Println(v9)
    _ = v10
    _ = v9
    return "hello"
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    v4 := []int{}
    var v5 bool = false
    v6 := make(map[int]string)
    if v5 {
        var v11 uint8 = 92
        _ = v11
    }
    if len(v4) > 0 {
        v4[0] = -137
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if ((3.14159 / 1e10) > ((((-2.25 * ((0.1 / 3.14159) - float32(3.14159))) * 0.1) - -2.25) * 1.5)) {
            if v5 {
                var v13 uint64 = 0
                _ = v13
            } else {
                switch ("foo bar" + "foo bar") {
                case "":
                    fmt.Println(fi12, v3)
                case "\n\t":
                    v5 = ((!(0 < -128)) || v5)
                default:
                    if v5 {
                        fmt.Println(v3, v5, v1)
                        fmt.Println(v5)
                        switch "hello" {
                        case "\n\t":
                            for fi14 := 0; fi14 < 10; fi14++ {
                                switch fi12 {
                                case 32768:
                                    v1 = v1
                                }
                                var v15 uint8 = 2
                                _ = fi14
                                _ = v15
                            }
                        }
                    }
                }
                if v5 {
                    fmt.Println(v3, v5)
                    var v16 float32 = -2.25
                    if v5 {
                        var v17 int = 0
                        for fi18 := 0; fi18 < 19; fi18++ {
                            if false {
                                var v19 int16 = -957
                                _ = v19
                            }
                            var v20 int32 = -595
                            var v21 float64 = (1.5 / 0.1)
                            _ = fi18
                            _ = v20
                            _ = v21
                        }
                        _ = v17
                    } else {
                        if v5 {
                            var v22 uint16 = 2
                            v6[(fi12 & (fi12 - -9223372036854775808))] = ("foo bar" + ("hello" + ("" + "foo bar")))
                            _ = v22
                        } else {
                            var v23 uint16 = 4
                            _ = v23
                        }
                        var v24 int = 1
                        _ = v24
                    }
                    _ = v16
                } else {
                    v25 := GenS1{F0: "foo bar", F1: (v5 && (v5 || (v1 <= v1))), F2: 0}
                    if len(v4) > 0 {
                        v4[0] = (fi12 | fi12)
                    }
                    _ = v25
                }
                v5 = ((fi12 | (fi12 << ((42) & 63))) <= 0)
            }
            v26 := GenS2{F0: 127, F1: false}
            switch "a" {
            case "":
                v27 := GenS3{F0: 108, F1: true, F2: 395}
                _ = v27
            case "foo bar":
                v4 = append(v4, fi12)
            default:
                if v5 {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        if v5 {
                            v3 = (^(v3))
                            if true {
                                var v29 int64 = 0
                                _ = v29
                            }
                        }
                        switch fi12 {
                        case 4294967295:
                            if v5 {
                                v3 = 245
                                var v30 float32 = 0.1
                                if (v5 && ((!v5) && true)) {
                                    v30 = v30
                                    fmt.Println(v3, v1, v1)
                                } else {
                                    v31 := GenS1{F0: ("foo bar" + ("a" + "hello")), F1: v5, F2: 100}
                                    fmt.Println(v30)
                                    _ = v31
                                }
                                _ = v30
                            }
                        default:
                            v32 := v6[(fi12 / ((-738) | 1))]
                            _ = v32
                        }
                        _ = fi28
                    }
                    switch (("\n\t" + ("\n\t" + "a")) + "foo bar") {
                    case "foo bar":
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 uint32 = 65535
                            _ = fi33
                            _ = v34
                        }
                    case "a":
                        for fi35 := 0; fi35 < 13; fi35++ {
                            if (!false) {
                                if len(v4) > 0 {
                                    v4[0] = fi12
                                }
                                v1 = (^(v1))
                            } else {
                                v5 = (1 <= 0)
                            }
                            if (255 == 42) {
                                v36 := GenS4{F0: 412}
                                _ = v36
                            }
                            v37 := GenS5{F0: ("" + ""), F1: v5}
                            _ = fi35
                            _ = v37
                        }
                    case "hello":
                        v1 -= v1
                    }
                } else {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        if v5 {
                            v3 = (-(v3))
                            var v39 int8 = 44
                            v3 = v3
                            _ = v39
                        } else {
                            v40 := GenS6{F0: fi38, F1: 1, F2: (v3 * 0)}
                            var v41 uint16 = 65535
                            fmt.Println(v41, v41, fi12)
                            _ = v40
                            _ = v41
                        }
                        _ = len(v2)
                        if ((65535 >= 255) && false) {
                            v3 = -32768
                            if len(v4) > 0 {
                                v4[0] = fi12
                            }
                            var v42 int64 = v1
                            _ = v42
                        }
                        _ = fi38
                    }
                }
            }
            _ = v26
        } else {
            v5 = ((fi12 + -2147483648) != 0)
            v43 := GenS3{F0: 2147483647, F1: true, F2: 902}
            for fi44 := 0; fi44 < 13; fi44++ {
                fmt.Println(fi12, fi44, fi12)
                _ = fi44
            }
            _ = v43
        }
        _ = fi12
    }
    var v45 int32 = 73
    _ = len(v4)
    if v5 {
        var v46 int = 314
        for fi47 := 0; fi47 < 3; fi47++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi47
        }
        _ = v46
    }
    v5 = (0 < 0)
    fmt.Println(v5, v5, v45)
    for fi48 := 0; fi48 < 17; fi48++ {
        for fi49 := 0; fi49 < 3; fi49++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v50 := v2[fi49]
            var v51 int16 = (^(v3))
            _ = fi49
            _ = v50
            _ = v51
        }
        v1 -= v1
        fmt.Println(fi48, v3, v3)
        _ = fi48
    }
    v52 := GenS6{F0: 0, F1: 18446744073709551615, F2: (v3 / (((v3 | 32767)) | 1))}
    for fi53 := 0; fi53 < 19; fi53++ {
        var v54 uint64 = 354
        v3 = v3
        _ = fi53
        _ = v54
    }
    if v5 {
        for fi55 := 0; fi55 < 2; fi55++ {
            var v56 int16 = 1
            _ = fi55
            _ = v56
        }
    }
    v57 := GenS7{F0: 255, F1: 226, F2: int(v45)}
    v2[-601] = 1
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v57
    _ = v6
}
