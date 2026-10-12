package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int64; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 int16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16; F1 string; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    var v3 float64 = float64((float32(0.1) / 1.5))
    var v4 int64 = 9223372036854775807
    fmt.Println(v2, v4)
    if true {
        var v5 uint8 = 99
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 2, F2: 2147483647}
    for fi7 := 0; fi7 < 11; fi7++ {
        if (-125 >= 100) {
            for fi8 := 0; fi8 < 10; fi8++ {
                switch (fi8 & (fi8 + fi8)) {
                case 4294967295:
                    if true {
                        if false {
                            v9 := v1[18]
                            _ = v9
                        }
                    }
                default:
                    v10 := GenS1{F0: 42, F1: 912, F2: int32(v4)}
                    _ = v10
                }
                v11 := GenS1{F0: 1, F1: 42, F2: -1}
                var v12 int = fi7
                switch "foo bar" {
                case "":
                    switch ("foo bar" + (("\n\t" + "\n\t") + (((("\n\t" + ("hello" + ("a" + "\n\t"))) + "\n\t") + "") + "a"))) {
                    case "foo bar":
                        v12 += v12
                    case "a":
                        if ("foo bar" < "\n\t") {
                            v2 -= v2
                        } else {
                            var v13 int64 = (v4 / ((v4) | 1))
                            var v14 int8 = 1
                            _ = v13
                            _ = v14
                        }
                    default:
                        v15 := []int{fi7, int(v2)}
                        _ = v15
                    }
                case "foo bar":
                    v16 := GenS2{F0: 255}
                    _ = v16
                }
                _ = fi8
                _ = v11
                _ = v12
            }
        } else {
            v3 = (float64(1e10) + -2.25)
            var v17 bool = (65535 <= 413)
            _ = v17
        }
        switch fi7 {
        case -463:
            var v18 uint32 = 2
            _ = v18
        default:
            if true {
                if true {
                    var v19 uint8 = 255
                    v20 := GenS1{F0: 1, F1: 65535, F2: -2147483648}
                    _ = v19
                    _ = v20
                } else {
                    v3 = ((v3 * v3) * ((v3 * float64((((0.1 + 1e10) / 1.0) / ((0.1 + 1e10) * 3.14159)))) - 0.0))
                }
                if false {
                    v21 := []int{fi7}
                    _ = v21
                }
            } else {
                if (!true) {
                    var v22 uint32 = 603
                    for fi23 := 0; fi23 < 6; fi23++ {
                        fmt.Println(v3, fi7, v22)
                        var v24 float32 = (0.1 * (0.1 * 1.5))
                        if false {
                            switch (fi23 / ((fi23) | 1)) {
                            case 9223372036854775807:
                                switch fi7 {
                                case 4294967295:
                                    v4 += (v4 << ((750) & 63))
                                case -9223372036854775808:
                                    v25 := []int{fi23}
                                    _ = v25
                                case 2147483647:
                                    v22 += (^(v22))
                                    var v26 int32 = 1
                                    v27 := GenS3{F0: 42, F1: (v4 * v4), F2: 0}
                                    _ = v26
                                    _ = v27
                                }
                            case 504:
                                delete(v1, v22)
                            default:
                                var v28 uint64 = 42
                                _ = v28
                            }
                        }
                        if true {
                            if false {
                                fmt.Println(v3)
                            } else {
                                v24 += (v24 + (v24 * float32((float64((v24 + (3.14159 / (-2.25 + (-2.25 + 0.1))))) * 1e10))))
                            }
                            v29 := []int{(fi23 / (((fi23 % ((fi23) | 1))) | 1)), (fi7 << ((130) & 63)), 42}
                            if false {
                                var v30 bool = true
                                var v31 uint8 = 168
                                switch (fi23 >> ((255) & 63)) {
                                case 0:
                                    v2 = v2
                                }
                                if ((v31 <= v31) && ((v30 || (("hello" <= "a") || v30)) || v30)) {
                                    switch "" {
                                    case "a":
                                        var v32 uint64 = (^(v2))
                                        _ = v32
                                    case "\n\t":
                                        v30 = (!v30)
                                    default:
                                        if true {
                                            if false {
                                                fmt.Println(v31, v2)
                                            } else {
                                                v3 = v3
                                            }
                                            var v33 float32 = v24
                                            _ = v33
                                        } else {
                                            var v34 string = ""
                                            v35 := GenS3{F0: -691, F1: 0, F2: 65535}
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    var v36 uint64 = (v2 + (v2 / ((148) | 1)))
                                    _ = v36
                                }
                                var v37 string = "hello"
                                _ = v30
                                _ = v31
                                _ = v37
                            }
                            _ = v29
                        }
                        fmt.Println(v2, v4, v24)
                        _ = fi23
                        _ = v24
                    }
                    v2 = (v2 * 1)
                    _ = v22
                }
                v4 = (v4 << ((1) & 63))
                for fi38 := 0; fi38 < 16; fi38++ {
                    v39 := GenS4{F0: 951}
                    var v40 uint16 = 255
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
        }
        _ = fi7
    }
    var v41 uint16 = 2
    for fi42 := 0; fi42 < 3; fi42++ {
        var v43 uint16 = 32767
        for fi44 := 0; fi44 < 2; fi44++ {
            v4 = v4
            fmt.Println(v3, v4, v43)
            _ = fi44
        }
        if true {
            for fi45 := 0; fi45 < 9; fi45++ {
                v41 = (v41 | v43)
                _ = len(v1)
                _ = fi45
            }
        } else {
            var v46 uint = uint(v2)
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    var v47 uint = 471
    v47 -= v47
    if (-97 != -1) {
        var v48 string = "hello"
        _ = v48
    } else {
        var v49 uint16 = (v41 << ((v47) & 63))
        v50 := []int{32768}
        v51 := GenS5{F0: 9223372036854775807, F1: uint64(v41), F2: 32767}
        _ = v49
        _ = v50
        _ = v51
    }
    _ = len(v1)
    v47 = (v47 ^ (v47 - v47))
    if (v2 >= v2) {
        v52 := GenS6{F0: (("foo bar" + "hello") + ("hello" + "hello")), F1: (v47 % (((^(v47))) | 1))}
        _ = v52
    } else {
        var v53 uint = v47
        v2 += v2
        _ = v53
    }
    v54 := GenS6{F0: "a", F1: (^(v47))}
    v55 := GenS7{F0: (v41 + 2), F1: "", F2: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v54
    _ = v55
    _ = v6
}
