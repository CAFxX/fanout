package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
func genFn0(p0 uint64, p1 uint16) bool {
    v4 := GenS1{F0: "\n\t", F1: (("hello" + "") >= "")}
    var v5 bool = false
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (true || v5)
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    v3 := []int{227, -440}
    var v6 int8 = -104
    v7 := GenS1{F0: "hello", F1: false}
    var v8 int32 = -317
    v9 := GenS1{F0: "a", F1: (1.5 >= 1e10)}
    if (1 == 2) {
        if true {
            var v10 string = "\n\t"
            _ = v10
        } else {
            var v11 int16 = 1
            if false {
                fmt.Println(v8, v6)
                v12 := v2[1]
                _ = v12
            }
            _ = v11
        }
    } else {
        v2[858] = "foo bar"
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v6)
        var v14 uint64 = 2
        _ = fi13
        _ = v14
    }
    if true {
        v8 = (v8 & 42)
        var v15 float32 = 3.14159
        v15 = (v15 + v15)
        _ = v15
    } else {
        v16 := GenS1{F0: "foo bar", F1: true}
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if (0 < (v8 >> ((47) & 63))) {
            fmt.Println(v8, v8)
        }
        _ = fi17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS1{F0: ("" + (("a" + "hello") + "foo bar")), F1: (2 <= 2147483648)}
        _ = fi18
        _ = v19
    }
    v6 = v6
    if (930 > 10) {
        v8 += (v8 + v8)
        v6 = (v6 % ((v6) | 1))
    }
    var v20 int64 = -628
    v21 := GenS1{F0: "hello", F1: true}
    var v22 string = "foo bar"
    v6 += v6
    v23 := GenS2{F0: v8, F1: false}
    if (!false) {
        v2[2147483648] = v22
        if false {
            switch "a" {
            case "\n\t":
                v24 := v2[2147483648]
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 6; fi25++ {
                    if true {
                        var v26 int = fi25
                        v27 := GenS3{F0: 0, F1: (fi25 / (((^(fi25))) | 1))}
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS2{F0: v8, F1: (65535 >= 254)}
                        v20 -= v20
                        if (368 != 42) {
                            if (-32768 == 447) {
                                switch (v22 + v22) {
                                case "hello":
                                    fmt.Println(fi25)
                                case "foo bar":
                                    v22 = (v22 + v22)
                                default:
                                    switch (-(fi25)) {
                                    case 9223372036854775807:
                                        var v29 int64 = v20
                                        _ = v29
                                    default:
                                        var v30 uint = 33
                                        _ = v30
                                    }
                                }
                                v8 = (-(v8))
                            } else {
                                v20 = v20
                            }
                            for fi31 := 0; fi31 < 11; fi31++ {
                                v6 += 42
                                var v32 bool = false
                                _ = fi31
                                _ = v32
                            }
                            v33 := GenS4{F0: 128, F1: v6, F2: 18}
                            _ = v33
                        } else {
                            var v34 int64 = -9223372036854775808
                            _ = v34
                        }
                        _ = v28
                    }
                    fmt.Println(v20, v8)
                    v22 = v22
                    _ = fi25
                }
            case "a":
                switch (v22 + (("" + "") + (v22 + ("a" + "")))) {
                case "\n\t":
                    v35 := GenS4{F0: 128, F1: -1, F2: 1}
                    _ = v35
                default:
                    v8 = v8
                }
            default:
                v3 = append(v3, 128)
            }
            v36 := GenS1{F0: v22, F1: true}
            if (32768 < -742) {
                if (1 > 42) {
                    if true {
                        var v37 bool = false
                        v6 = v6
                        v22 = ("a" + v22)
                        _ = v37
                    }
                    v38 := GenS3{F0: 1, F1: 42}
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        fmt.Println(v6, fi39)
                        if true {
                            v40 := v2[4294967295]
                            if true {
                                var v41 float64 = (1.5 + 0.1)
                                v8 -= 372
                                for fi42 := 0; fi42 < 19; fi42++ {
                                    var v43 string = v40
                                    v20 = -1
                                    v22 = v22
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                            _ = v40
                        }
                        _ = fi39
                    }
                }
            }
            _ = v36
        }
        switch v22 {
        case "":
            switch 379 {
            case -2147483648:
                switch -92 {
                case 4294967295:
                    var v44 float64 = 0.0
                    _ = v44
                default:
                    var v45 int64 = -9223372036854775808
                    _ = v45
                }
            case 0:
                v46 := GenS5{F0: 794}
                _ = v46
            }
        case "foo bar":
            fmt.Println(v22, v20)
        case "hello":
            if ((3.14159 != 0.0) || ((42 == 1) || (2 < 42))) {
                if false {
                    v8 += v8
                    v8 += (v8 - -1)
                }
                if len(v3) > 0 {
                    v3[0] = -882
                }
                if ((v6 - (v6 % ((v6) | 1))) == (^(v6))) {
                    fmt.Println(v8)
                    for fi47 := 0; fi47 < 15; fi47++ {
                        if len(v3) > 0 {
                            v3[0] = fi47
                        }
                        _ = fi47
                    }
                } else {
                    v22 = "foo bar"
                    v48 := GenS3{F0: 1, F1: -9223372036854775808}
                    var v49 int16 = -32768
                    _ = v48
                    _ = v49
                }
            }
        default:
            var v50 bool = true
            _ = v50
        }
    } else {
        fmt.Println(v22, v6, v22)
        switch v22 {
        case "\n\t":
            var v51 int16 = -32768
            _ = v51
        case "foo bar":
            v52 := GenS6{F0: v22}
            _ = v52
        case "a":
            v6 = v6
        default:
            v1 = append(v1, ("" + v22))
        }
        v53 := v2[1]
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
