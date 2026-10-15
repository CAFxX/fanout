package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64; F2 uint }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
type GenS4 struct { F0 uint64 }
func genFn0() bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint32 = 0
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 188
    var v7 int16 = 1
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(fi9, v7, v7)
        var v10 float64 = (((1.5 - -2.25) * 1.5) - 0.0)
        v11 := GenS1{F0: v6}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = ("hello" + ("a" + "foo bar"))
    v3 := make(map[uint32]uint32)
    v1 = (v1 + (v1 + (^(v1))))
    v12 := GenS1{F0: 18446744073709551615}
    fmt.Println(v2)
    v13 := []int{}
    if (0 == 507) {
        if ((false || true) && (127 != -2147483648)) {
            var v14 float32 = (3.14159 / (1.5 + ((0.0 - 1e10) - (-2.25 * float32(0.0)))))
            _ = v14
        } else {
            if false {
                v1 -= v1
                if (42 <= 0) {
                    var v15 float32 = float32(1.5)
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                fmt.Println(v1, fi16, fi16)
                switch fi16 {
                case 32768:
                    v17 := GenS1{F0: 18446744073709551615}
                    _ = v17
                case 127:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        switch "" {
                        case "a":
                            var v19 uint32 = (v1 - v1)
                            switch (((("hello" + v2) + v2) + "") + v2) {
                            case "\n\t":
                                var v20 float64 = (-2.25 - 0.1)
                                _ = v20
                            case "a":
                                if true {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                } else {
                                    if true {
                                        v1 -= (v1 - 4294967295)
                                        switch (v2 + v2) {
                                        case "foo bar":
                                            v21 := v3[(v19 ^ v19)]
                                            _ = v21
                                        case "a":
                                            v22 := v3[(v19 - (v1 / ((133) | 1)))]
                                            _ = v22
                                        default:
                                            delete(v3, 65535)
                                            if len(v13) > 0 {
                                                v13[0] = (fi16 >> ((2) & 63))
                                            }
                                        }
                                        v23 := GenS1{F0: 18446744073709551615}
                                        _ = v23
                                    } else {
                                        v2 = (v2 + ((v2 + ("hello" + "hello")) + "a"))
                                        v2 = v2
                                    }
                                }
                                v2 = v2
                                var v24 int32 = 2147483647
                                var v25 int64 = -9223372036854775808
                                _ = v24
                                _ = v25
                            case "hello":
                                var v26 float64 = (((0.1 / 1e10) - 3.14159) - 1.5)
                                _ = v26
                            default:
                                if (false || false) {
                                    if false {
                                        if true {
                                            fmt.Println(v1)
                                            v27 := GenS2{F0: fi18, F1: uint64(v19), F2: 1}
                                            switch (fi16 * 1) {
                                            case -2147483648:
                                                v19 = 65535
                                            case 507:
                                                v1 = 0
                                            case 2147483647:
                                                delete(v3, 436)
                                                v28 := GenS3{F0: 740, F1: -1, F2: 552}
                                                var v29 float64 = (((0.0 / -2.25) + 1e10) * ((3.14159 + -2.25) + 0.0))
                                                _ = v28
                                                _ = v29
                                            default:
                                                v19 -= v1
                                            }
                                            _ = v27
                                        } else {
                                            if (1 < 2) {
                                                fmt.Println(v2)
                                                _ = len(v3)
                                                var v30 uint8 = 1
                                                _ = v30
                                            }
                                            var v31 int32 = -2147483648
                                            _ = v31
                                        }
                                        if (1 != -1) {
                                            fmt.Println(fi16, v19, v1)
                                        }
                                        v19 = 2147483648
                                    }
                                    v32 := GenS3{F0: -935, F1: 82, F2: 0}
                                    fmt.Println(fi16, v2)
                                    _ = v32
                                }
                            }
                            v33 := GenS1{F0: 789}
                            _ = v19
                            _ = v33
                        }
                        _ = fi18
                    }
                }
                _ = fi16
            }
        }
        for fi34 := 0; fi34 < 13; fi34++ {
            v35 := GenS1{F0: 793}
            _ = fi34
            _ = v35
        }
    } else {
        if false {
            var v36 int64 = 984
            if (!true) {
                v37 := GenS2{F0: -1, F1: 18446744073709551615, F2: 0}
                _ = v37
            } else {
                if false {
                    var v38 int16 = 42
                    _ = v38
                }
            }
            _ = v36
        } else {
            v2 = v2
        }
    }
    v39 := GenS2{F0: 83, F1: 721, F2: 1000}
    v2 = (("\n\t" + (("a" + "a") + v2)) + (v2 + v2))
    v2 = v2
    if true {
        v1 += (^(v1))
    }
    v1 = 2147483648
    fmt.Println(v2)
    if ((v1 / ((65535) | 1)) < v1) {
        var v40 int64 = -1
        v2 = (v2 + (v2 + v2))
        _ = v40
    } else {
        _ = len(v13)
        var v41 bool = (-2.25 == (-2.25 + 0.0))
        _ = v41
    }
    v42 := GenS3{F0: -1, F1: 0, F2: 4294967295}
    fmt.Println(v2, v1)
    var v43 uint64 = 2
    if true {
        switch -216 {
        case 4294967295:
            v44 := GenS2{F0: 4294967295, F1: (^(v43)), F2: 255}
            _ = v44
        case 2147483647:
            for fi45 := 0; fi45 < 11; fi45++ {
                var v46 float64 = ((-2.25 / 1.0) / 1.5)
                if (-720 > 130) {
                    var v47 uint = 65535
                    _ = v47
                }
                if true {
                    var v48 string = v2
                    v49 := GenS2{F0: (fi45 / (((fi45 & 42)) | 1)), F1: (v43 << ((4294967295) & 63)), F2: 1}
                    _ = v48
                    _ = v49
                }
                _ = fi45
                _ = v46
            }
        case 0:
            for fi50 := 0; fi50 < 4; fi50++ {
                var v51 uint16 = 65535
                v52 := GenS4{F0: (v43 % (((v43 / ((2) | 1))) | 1))}
                if true {
                    delete(v3, (v1 * (^(v1))))
                    v2 = "\n\t"
                }
                _ = fi50
                _ = v51
                _ = v52
            }
        default:
            v43 = v43
        }
    } else {
        var v53 int32 = -991
        _ = v53
    }
    var v54 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v42
    _ = v43
    _ = v54
}
