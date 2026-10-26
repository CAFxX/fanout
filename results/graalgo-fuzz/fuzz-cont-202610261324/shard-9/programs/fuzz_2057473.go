package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
type GenS4 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint = 1
    var v2 int32 = 714
    var v3 int64 = -1
    v4 := []string{"", "", "", "hello"}
    v5 := make(map[string]string)
    if (float64(3.14159) == (((1.5 / 1.5) - 3.14159) - 0.1)) {
        v2 = (v2 | (^(v2)))
        var v6 uint16 = 32767
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int32 = v2
                _ = len(v4)
                v3 = v3
                var v10 int8 = 42
                switch (^(fi7)) {
                case -398:
                    v2 = 0
                }
                var v11 uint32 = 2
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v5["a"] = "foo bar"
            _ = fi7
        }
        _ = v6
    }
    v12 := GenS1{F0: 32767, F1: 4294967295}
    v13 := GenS1{F0: 42, F1: 1}
    v1 = v1
    if len(v4) > 0 {
        v4[0] = ("hello" + "a")
    }
    v4 = append(v4, ("" + "\n\t"))
    for fi14 := 0; fi14 < 7; fi14++ {
        if (!((false && (true && false)) && true)) {
            if true {
                v15 := GenS2{F0: int8(v3)}
                if (0 != 1) {
                    _ = len(v5)
                    var v16 float64 = 1.5
                    if (4294967295 >= 124) {
                        v16 = 0.0
                    } else {
                        v1 -= (v1 / (((v1 ^ v1)) | 1))
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 uint = 255
                            var v19 bool = (42 != 1)
                            delete(v5, "foo bar")
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = v16
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v20 float32 = -2.25
                    _ = v20
                }
                _ = v15
            }
            v21 := GenS1{F0: -9223372036854775808, F1: 0}
            _ = v21
        } else {
            fmt.Println(fi14, v1)
            if (-752 <= 1) {
                var v22 int64 = v3
                v23 := GenS1{F0: (fi14 << ((v1) & 63)), F1: 18446744073709551615}
                var v24 uint8 = 128
                _ = v22
                _ = v23
                _ = v24
            } else {
                switch "\n\t" {
                case "\n\t":
                    v2 -= v2
                case "":
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v1 += v1
                        v1 = (v1 * 107)
                        _ = fi25
                    }
                case "a":
                    v5["hello"] = "hello"
                default:
                    var v26 int16 = -276
                    _ = v26
                }
                v27 := GenS3{F0: 4294967295, F1: ((true && true) || false), F2: (v2 | (-(v2)))}
                _ = v27
            }
        }
        var v28 uint16 = 2
        _ = fi14
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v4 = append(v4, "hello")
        _ = fi29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        if (false && (65535 != 1)) {
            var v31 uint32 = 65535
            var v32 float64 = -2.25
            var v33 uint16 = 47
            _ = v31
            _ = v32
            _ = v33
        } else {
            v2 += (v2 >> ((v1) & 63))
        }
        var v34 int = (-(fi30))
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 string = "a"
            var v37 float32 = 1.5
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi30
        _ = v34
    }
    if true {
        var v38 uint32 = 2147483648
        _ = v38
    } else {
        v39 := GenS2{F0: -1}
        fmt.Println(v1, v3)
        var v40 int32 = v2
        _ = v39
        _ = v40
    }
    if (!false) {
        var v41 uint64 = 1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            if (v41 < v41) {
                switch 436 {
                case 32768:
                    if true {
                        if ((0.1 / -2.25) < 1.5) {
                            v43 := GenS2{F0: 100}
                            for fi44 := 0; fi44 < 3; fi44++ {
                                var v45 int32 = 121
                                _ = fi44
                                _ = v45
                            }
                            _ = v43
                        }
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    } else {
                        if false {
                            v46 := v5["hello"]
                            v1 -= 527
                            _ = v46
                        }
                        delete(v5, "hello")
                    }
                default:
                    var v47 int16 = 1000
                    _ = v47
                }
                v48 := GenS2{F0: 113}
                for fi49 := 0; fi49 < 11; fi49++ {
                    v50 := GenS4{F0: 447, F1: 958}
                    var v51 uint64 = v41
                    fmt.Println(v2)
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                _ = v48
            }
            _ = fi42
        }
        _ = v41
    }
    v52 := GenS4{F0: 32767, F1: 608}
    var v53 bool = (!false)
    v53 = (128 == 9223372036854775807)
    v54 := GenS3{F0: v1, F1: v53, F2: (v2 << ((v1) & 63))}
    _ = len(v5)
    var v55 bool = v53
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
