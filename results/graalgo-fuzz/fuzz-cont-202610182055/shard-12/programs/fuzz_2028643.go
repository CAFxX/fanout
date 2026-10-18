package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0() string {
    v5 := GenS1{F0: 20, F1: 18446744073709551615, F2: -1}
    v6 := []int{-1, 4294967295, 32767}
    var v7 uint16 = 42
    v8 := GenS2{F0: 42, F1: 42}
    v7 = (v7 >> ((255) & 63))
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 string = (("" + "hello") + "foo bar")
    var v2 string = ("\n\t" + (v1 + "\n\t"))
    v3 := make(map[string]int)
    var v4 int8 = 89
    if true {
        if false {
            v1 = "hello"
            _ = len(v3)
            v1 = ("a" + ((v1 + v1) + ((v1 + "a") + v1)))
        }
    } else {
        _ = len(v3)
        for fi9 := 0; fi9 < 20; fi9++ {
            switch fi9 {
            case 0:
                for fi10 := 0; fi10 < 15; fi10++ {
                    v4 += 42
                    if (18446744073709551615 < 2) {
                        var v11 uint8 = 102
                        if false {
                            v12 := []int{fi10, fi10, 127}
                            switch v1 {
                            case "foo bar":
                                if true {
                                    v2 = v2
                                    v13 := GenS1{F0: 42, F1: 18446744073709551615, F2: -118}
                                    _ = len(v3)
                                    _ = v13
                                }
                            case "a":
                                switch (fi10 & fi10) {
                                case -2147483648:
                                    if len(v12) > 0 {
                                        v12[0] = fi10
                                    }
                                default:
                                    if (562 > -1) {
                                        v4 -= 1
                                    }
                                }
                            }
                            v2 = v2
                            var v14 uint64 = 1
                            _ = v12
                            _ = v14
                        } else {
                            if true {
                                var v15 string = "foo bar"
                                _ = v15
                            } else {
                                switch fi10 {
                                case 128:
                                    var v16 int32 = 42
                                    _ = v16
                                case 0:
                                    fmt.Println(fi10)
                                case -430:
                                    var v17 int = (fi9 % (((fi9 + (fi9 | fi9))) | 1))
                                    _ = v17
                                default:
                                    fmt.Println(fi9, v1, fi9)
                                }
                                v18 := []int{-844, (fi9 ^ fi10), (fi10 ^ -1), -566}
                                _ = v18
                            }
                            var v19 uint64 = 2
                            _ = v19
                        }
                        fmt.Println(v2, fi10)
                        v20 := GenS3{F0: 308, F1: v2}
                        _ = v11
                        _ = v20
                    }
                    switch v2 {
                    case "hello":
                        v4 = int8(fi10)
                    case "foo bar":
                        var v21 uint8 = 127
                        _ = v21
                    case "":
                        v1 = ((v1 + ("\n\t" + v2)) + "hello")
                    }
                    _ = fi10
                }
            case -816:
                v22 := GenS4{F0: 2, F1: 1, F2: (532 < 1)}
                _ = v22
            default:
                v23 := GenS5{F0: 1}
                _ = v23
            }
            _ = fi9
        }
    }
    var v24 int16 = 0
    var v25 uint16 = 977
    if (float32((-2.25 - 0.0)) == 1e10) {
        v26 := GenS6{F0: 0, F1: 2147483647, F2: 65535}
        v1 = "foo bar"
        _ = v26
    }
    var v27 int = 128
    var v28 uint8 = 128
    switch (v27 >> ((4294967295) & 63)) {
    case -2147483648:
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v28)
            for fi30 := 0; fi30 < 19; fi30++ {
                switch "a" {
                case "":
                    fmt.Println(v27)
                default:
                    if true {
                        fmt.Println(fi29, v25, v27)
                        if (113 <= 255) {
                            v25 = (^(v25))
                            fmt.Println(fi30)
                            var v31 uint = 246
                            var v32 string = (v1 + v2)
                            var v33 uint16 = v25
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            if true {
                                if false {
                                    var v34 int16 = 861
                                    v2 = "\n\t"
                                    v35 := []int{}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v25 -= 32767
                                }
                                var v36 int64 = 42
                                _ = v36
                            }
                            switch (((v1 + "\n\t") + "foo bar") + v2) {
                            case "foo bar":
                                switch v1 {
                                case "":
                                    v4 = (v4 / (((v4 - v4)) | 1))
                                case "hello":
                                    switch fi29 {
                                    case 4294967295:
                                        var v37 float32 = -2.25
                                        _ = v37
                                    default:
                                        fmt.Println(v1)
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v24)
                            }
                        }
                    } else {
                        if false {
                            v38 := GenS7{F0: (v25 + v25), F1: -1, F2: 13}
                            v2 = "hello"
                            _ = v38
                        }
                        v39 := GenS4{F0: (v25 % ((0) | 1)), F1: 42, F2: (142 <= 556)}
                        fmt.Println(v28, v25)
                        if true {
                            var v40 uint8 = 0
                            _ = v40
                        } else {
                            v41 := GenS4{F0: 32767, F1: 65535, F2: (18446744073709551615 > 18446744073709551615)}
                            _ = v41
                        }
                        switch (-(v27)) {
                        case 4294967295:
                            v1 = (("" + (v1 + "a")) + v1)
                        }
                        _ = v39
                    }
                }
                v28 -= (v28 ^ v28)
                if (true || (0.0 != (0.0 - 0.0))) {
                    switch (fi29 / (((fi30 + (fi29 << ((18446744073709551615) & 63)))) | 1)) {
                    case 377:
                        v28 = v28
                    case -2147483648:
                        var v42 int16 = -1
                        _ = v42
                    case -1:
                        var v43 int64 = 336
                        _ = v43
                    default:
                        fmt.Println(v28, v24, v24)
                    }
                    var v44 uint = 1
                    _ = v44
                } else {
                    v45 := GenS1{F0: v28, F1: 42, F2: 9223372036854775807}
                    v4 -= (v4 - (^(v4)))
                    if (((v1 + v2) + "hello") != "foo bar") {
                        if false {
                            fmt.Println(v2, fi30, v1)
                        }
                    }
                    _ = v45
                }
                _ = fi30
            }
            fmt.Println(v2, v25)
            _ = fi29
        }
    case 2147483647:
        if (false || (true && true)) {
            v27 = 1
            v1 = "hello"
            delete(v3, ((v2 + v2) + v2))
        }
    default:
        v46 := GenS3{F0: 2147483648, F1: v1}
        _ = v46
    }
    v2 = "hello"
    if true {
        var v47 float32 = ((float32(-2.25) + (3.14159 * 3.14159)) / -2.25)
        _ = v47
    } else {
        v28 = (v28 / (((v28 + (v28 << ((750) & 63)))) | 1))
        v25 = v25
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS7{F0: 805, F1: v24, F2: 1}
            var v50 bool = (91 < (v4 - v4))
            if v50 {
                var v51 uint32 = 174
                v3[v2] = (v27 % ((42) | 1))
                _ = v51
            }
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    fmt.Println(v2, v28)
    if false {
        v52 := GenS6{F0: 42, F1: 2147483647, F2: 793}
        var v53 float32 = 1e10
        _ = v52
        _ = v53
    }
    var v54 uint = 42
    v55 := GenS1{F0: 127, F1: (v54 * (v54 * 65535)), F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
