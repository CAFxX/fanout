package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = 0.0
    var v3 int16 = 42
    var v4 uint8 = 1
    v4 = 128
    v5 := GenS1{F0: 127}
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS2{F0: int32(v4), F1: "foo bar", F2: 1}
        _ = fi6
        _ = v7
    }
    if (("a" == ("hello" + ("a" + "a"))) && true) {
        v8 := GenS2{F0: -1, F1: "a", F2: 18}
        v9 := GenS2{F0: v1, F1: "hello", F2: 18446744073709551615}
        v1 = v1
        _ = v8
        _ = v9
    }
    if ("foo bar" < "") {
        var v10 int = 4294967295
        v11 := GenS2{F0: v1, F1: "", F2: 781}
        var v12 int64 = -424
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v1, v2, v1)
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 5; fi14++ {
        v13 = ((((v13 + (v13 + ("\n\t" + ((v13 + "\n\t") + (((v13 + "a") + "\n\t") + v13))))) + ((v13 + "") + v13)) + v13) + "hello")
        _ = fi14
    }
    v13 = "a"
    if false {
        var v15 uint32 = 65535
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int64 = 9223372036854775807
            for fi18 := 0; fi18 < 19; fi18++ {
                if (true || false) {
                    fmt.Println(v1, v15, fi16)
                } else {
                    var v19 int64 = v17
                    _ = v19
                }
                switch v13 {
                case "":
                    v17 = v17
                case "hello":
                    var v20 int = (^(fi16))
                    v21 := []int{(fi16 / (((v20 << ((42) & 63))) | 1)), (-(fi18)), int(v15), (fi18 % (((^(fi16))) | 1))}
                    _ = v20
                    _ = v21
                }
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            v4 = (v4 << ((243) & 63))
            v13 = "foo bar"
            if (!((v3 != v3) && ((0.0 - ((3.14159 * 1.5) / float64(v2))) == (0.0 + ((1.5 - -2.25) + -2.25))))) {
                var v23 uint8 = v4
                var v24 int32 = v1
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        _ = v15
    }
    v4 = 2
    v25 := []int{-9223372036854775808}
    if ((uint(v4) <= 646) && (!(1.5 > (-2.25 * float32(-2.25))))) {
        v2 += float32(((0.0 + -2.25) * 3.14159))
        v13 = (v13 + "foo bar")
    } else {
        v26 := GenS3{F0: v4}
        _ = v26
    }
    v27 := GenS2{F0: v1, F1: v13, F2: 609}
    if (488 >= 170) {
        var v28 uint8 = 42
        var v29 uint32 = 2147483648
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint8 = 127
            v13 = "foo bar"
            _ = fi30
            _ = v31
        }
        if false {
            v2 = 1e10
            if len(v25) > 0 {
                v25[0] = 2147483647
            }
            if (942 <= 128) {
                if (2 > 393) {
                    if len(v25) > 0 {
                        v25[0] = -792
                    }
                    fmt.Println(v13)
                    v2 -= float32(((3.14159 * 1e10) / 1.5))
                }
                v13 = "a"
                _ = len(v25)
            } else {
                fmt.Println(v2)
                switch -366 {
                case 4294967295:
                    var v32 string = ((v13 + (v13 + v13)) + v13)
                    _ = v32
                case -384:
                    if ((v3 & 775) <= 1) {
                        if false {
                            v2 -= (3.14159 + v2)
                        }
                        v13 = v13
                        if false {
                            v4 = 127
                        } else {
                            fmt.Println(v3)
                            var v33 int64 = int64(v3)
                            _ = v33
                        }
                    }
                case 2147483647:
                    var v34 uint8 = v4
                    _ = v34
                default:
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                }
            }
        }
        if true {
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 uint16 = 65535
                if (!(v1 != (v1 ^ 2147483647))) {
                    if (!(fi35 == 127)) {
                        if (175 > 65535) {
                            for fi37 := 0; fi37 < 11; fi37++ {
                                if true {
                                    v2 = (-2.25 - -2.25)
                                    if false {
                                        var v38 int32 = v1
                                        if true {
                                            var v39 int32 = v38
                                            fmt.Println(v13, fi35)
                                            _ = v39
                                        }
                                        _ = v38
                                    }
                                }
                                var v40 uint32 = 320
                                v3 -= (-(v3))
                                _ = fi37
                                _ = v40
                            }
                            var v41 string = (("\n\t" + "") + "foo bar")
                            _ = v41
                        } else {
                            v36 -= (v36 % ((v36) | 1))
                            v13 = (v13 + v13)
                            var v42 string = "a"
                            _ = v42
                        }
                        v43 := GenS4{F0: true, F1: 42, F2: true}
                        if true {
                            v44 := GenS2{F0: -1, F1: v13, F2: 0}
                            fmt.Println(v3)
                            switch "hello" {
                            case "foo bar":
                                v13 = ((v13 + "") + ((v13 + "hello") + v13))
                            default:
                                fmt.Println(v2)
                            }
                            _ = v44
                        }
                        _ = v43
                    } else {
                        v45 := GenS4{F0: false, F1: -9223372036854775808, F2: true}
                        for fi46 := 0; fi46 < 7; fi46++ {
                            var v47 int16 = (v3 >> ((65535) & 63))
                            fmt.Println(v3, v4, v47)
                            if ((v47 >> ((557) & 63)) != (-(v3))) {
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                            } else {
                                if len(v25) > 0 {
                                    v25[0] = (^(fi35))
                                }
                                var v48 float64 = 1e10
                                _ = v48
                            }
                            if false {
                                var v49 string = v13
                                v50 := GenS5{F0: 127}
                                _ = v49
                                _ = v50
                            }
                            _ = fi46
                            _ = v47
                        }
                        v25 = append(v25, fi35)
                        _ = v45
                    }
                } else {
                    var v51 bool = (((0.0 * 0.1) + ((((0.0 / 1e10) + 3.14159) - 0.0) / 3.14159)) != (((0.0 / 0.1) - 1.5) + 3.14159))
                    switch (fi35 ^ 42) {
                    case -1:
                        for fi52 := 0; fi52 < 5; fi52++ {
                            switch v13 {
                            case "":
                                if ((-(v3)) == (v3 * (v3 % (((^(v3))) | 1)))) {
                                    if (v51 && v51) {
                                        v53 := GenS1{F0: fi35}
                                        v54 := GenS6{F0: v13}
                                        v2 = 1e10
                                        _ = v53
                                        _ = v54
                                    }
                                    v25 = append(v25, (fi52 + (fi35 & (-(fi35)))))
                                }
                            case "\n\t":
                                v1 -= (v1 | 2147483647)
                            default:
                                v1 = 0
                            }
                            _ = fi52
                        }
                    case 0:
                        v36 += 0
                    }
                    var v55 uint64 = 670
                    _ = v51
                    _ = v55
                }
                _ = fi35
                _ = v36
            }
        } else {
            v56 := GenS3{F0: v4}
            _ = v56
        }
    }
    v25 = append(v25, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
