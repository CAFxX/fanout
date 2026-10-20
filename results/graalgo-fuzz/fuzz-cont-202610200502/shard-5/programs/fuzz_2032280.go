package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint64; F2 int64 }

func main() {
    v1 := []string{("a" + "foo bar"), ("a" + ("" + "hello"))}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := []int{-987}
    if (1 >= 42) {
        var v5 int16 = 32767
        if (62 == 1) {
            v5 += v5
            switch -9223372036854775808 {
            case 1:
                v5 -= 1000
            default:
                var v6 float32 = 1e10
                _ = v6
            }
            switch 32768 {
            case 4294967295:
                var v7 string = "\n\t"
                _ = v7
            }
        } else {
            var v8 int32 = 1
            switch 1 {
            case 9223372036854775807:
                var v9 uint32 = 2147483647
                _ = v9
            default:
                if (0 == 259) {
                    switch "a" {
                    case "":
                        v10 := GenS1{F0: 4294967295}
                        _ = v10
                    default:
                        var v11 int = 127
                        _ = v11
                    }
                }
            }
            fmt.Println(v5, v8)
            _ = v8
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 += 1
            _ = fi12
        }
        _ = v5
    }
    var v13 uint8 = 255
    fmt.Println(v13)
    v13 += 128
    v14 := GenS1{F0: 18446744073709551615}
    fmt.Println(v13)
    if true {
        for fi15 := 0; fi15 < 16; fi15++ {
            switch (fi15 << ((65535) & 63)) {
            case 4294967295:
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v13, fi16)
                    var v17 int16 = 32767
                    v17 = (v17 | 1000)
                    _ = fi16
                    _ = v17
                }
            case 401:
                switch ("" + "\n\t") {
                case "foo bar":
                    v13 = v13
                case "":
                    v18 := GenS1{F0: 18446744073709551615}
                    _ = v18
                case "hello":
                    fmt.Println(fi15)
                default:
                    var v19 uint32 = 334
                    _ = v19
                }
            default:
                switch "hello" {
                case "foo bar":
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v21 := GenS1{F0: 899}
                        _ = fi20
                        _ = v21
                    }
                case "":
                    fmt.Println(fi15, fi15)
                default:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v13 += (v13 | v13)
                        _ = fi22
                    }
                }
            }
            _ = fi15
        }
        fmt.Println(v13)
        v23 := GenS2{F0: 32767, F1: 0, F2: 42}
        _ = v23
    } else {
        v24 := GenS3{F0: true, F1: false, F2: -2147483648}
        _ = v24
    }
    if true {
        v25 := GenS1{F0: 504}
        var v26 uint16 = 330
        _ = v25
        _ = v26
    } else {
        v27 := GenS3{F0: (2147483647 >= int32(v13)), F1: true, F2: -2147483648}
        for fi28 := 0; fi28 < 7; fi28++ {
            v13 = (v13 + (v13 << ((255) & 63)))
            if ((1 != 2) && ((((1e10 - 0.0) * 1e10) + float32(-2.25)) >= 1e10)) {
                for fi29 := 0; fi29 < 12; fi29++ {
                    if (9223372036854775807 < -55) {
                        v30 := GenS4{F0: -128}
                        var v31 uint32 = 42
                        _ = len(v2)
                        _ = v30
                        _ = v31
                    } else {
                        v13 -= v13
                    }
                    switch fi29 {
                    case 32767:
                        var v32 uint16 = 42
                        _ = v32
                    }
                    _ = fi29
                }
                if true {
                    v33 := GenS3{F0: (255 != 65535), F1: true, F2: 2147483647}
                    _ = v33
                }
            }
            if false {
                v13 = (v13 & (v13 << ((466) & 63)))
                switch fi28 {
                case 128:
                    for fi34 := 0; fi34 < 11; fi34++ {
                        v13 += v13
                        if false {
                            var v35 string = "a"
                            if true {
                                fmt.Println(fi28, v13, fi28)
                                var v36 string = v35
                                _ = v36
                            } else {
                                if (42 == 60) {
                                    v13 -= v13
                                }
                                var v37 uint8 = (v13 & 128)
                                _ = v37
                            }
                            v13 += v13
                            var v38 int = fi34
                            _ = v35
                            _ = v38
                        }
                        _ = fi34
                    }
                default:
                    delete(v3, 1)
                }
            } else {
                v39 := GenS5{F0: 0, F1: 42, F2: -558}
                fmt.Println(fi28)
                var v40 int64 = -9223372036854775808
                _ = v39
                _ = v40
            }
            _ = fi28
        }
        if ((!((42 < 1000) || (-756 > -1))) || true) {
            var v41 uint64 = 4294967295
            if true {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                var v42 float32 = 0.0
                _ = v42
            } else {
                v41 = (v41 % ((v41) | 1))
                v43 := GenS3{F0: true, F1: (-596 >= 2147483647), F2: 0}
                if false {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        var v45 uint16 = 1
                        _ = fi44
                        _ = v45
                    }
                    var v46 int32 = 1000
                    _ = v46
                }
                _ = v43
            }
            fmt.Println(v41, v13)
            _ = v41
        } else {
            var v47 float32 = 0.1
            _ = v47
        }
        _ = v27
    }
    switch 1 {
    case 1:
        if len(v1) > 0 {
            v1[0] = ((((("foo bar" + ((((("foo bar" + "hello") + "hello") + ("" + "a")) + "a") + "foo bar")) + "a") + "foo bar") + "\n\t") + "a")
        }
    case 449:
        fmt.Println(v13)
    default:
        for fi48 := 0; fi48 < 14; fi48++ {
            delete(v3, 39)
            _ = fi48
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
