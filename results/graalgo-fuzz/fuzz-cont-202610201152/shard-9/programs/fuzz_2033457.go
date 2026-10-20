package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int; F2 uint16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int16 = -982
    v4 := []uint32{846, 2, 65535, 522}
    if false {
        if true {
            v5 := GenS1{F0: 0, F1: "a"}
            switch "hello" {
            case "a":
                if (137 > 128) {
                    for fi6 := 0; fi6 < 20; fi6++ {
                        if false {
                            if false {
                                v7 := GenS2{F0: -172, F1: 502}
                                var v8 float32 = (1e10 + (1e10 + float32(3.14159)))
                                _ = v7
                                _ = v8
                            } else {
                                v3 = (^(v3))
                            }
                        } else {
                            if false {
                                switch 4294967295 {
                                case -135:
                                    v3 = v3
                                case 127:
                                    var v9 int = (fi6 ^ fi6)
                                    _ = v9
                                case 128:
                                    for fi10 := 0; fi10 < 1; fi10++ {
                                        if (-2.25 < ((float64(-2.25) * (1.5 - ((0.0 / (3.14159 + 1e10)) - float64(1.5)))) * 3.14159)) {
                                            v3 = -1
                                        } else {
                                            fmt.Println(fi10, fi10)
                                        }
                                        v3 += 173
                                        _ = fi10
                                    }
                                default:
                                    v3 = 1000
                                }
                            } else {
                                delete(v2, 42)
                                v11 := v2[0]
                                fmt.Println(v3, fi6)
                                _ = v11
                            }
                        }
                        v12 := GenS3{F0: 1, F1: 2}
                        _ = fi6
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS1{F0: 1, F1: "hello"}
                        v3 = 0
                        switch (("hello" + "foo bar") + "foo bar") {
                        case "":
                            switch -2147483648 {
                            case -2147483648:
                                for fi15 := 0; fi15 < 8; fi15++ {
                                    if len(v4) > 0 {
                                        v4[0] = 700
                                    }
                                    v3 = v3
                                    _ = fi15
                                }
                            case -351:
                                v16 := v2[4294967295]
                                _ = v16
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                }
            case "\n\t":
                if (255 <= 86) {
                    switch 536 {
                    case -196:
                        var v17 int8 = 37
                        _ = v17
                    case 662:
                        for fi18 := 0; fi18 < 20; fi18++ {
                            v19 := GenS4{F0: (v3 & v3)}
                            for fi20 := 0; fi20 < 7; fi20++ {
                                switch "" {
                                case "":
                                    if true {
                                        v2[4294967295] = 102
                                    } else {
                                        fmt.Println(fi20)
                                        var v21 uint16 = 255
                                        _ = v21
                                    }
                                default:
                                    var v22 float64 = 3.14159
                                    _ = v22
                                }
                                _ = fi20
                            }
                            v3 = (v3 ^ (v3 << ((1) & 63)))
                            _ = fi18
                            _ = v19
                        }
                    case 460:
                        v3 = (v3 | v3)
                    default:
                        if true {
                            fmt.Println(v3)
                            var v23 int16 = -1
                            v23 = (v23 ^ (v3 - (v3 * 0)))
                            _ = v23
                        }
                    }
                    v3 += 32767
                } else {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            switch ("hello" + "hello") {
                            case "hello":
                                var v26 float32 = -2.25
                                _ = v26
                            }
                            fmt.Println(fi25)
                            _ = fi25
                        }
                        _ = fi24
                    }
                }
            default:
                fmt.Println(v3)
            }
            v3 += (v3 / (((v3 - v3)) | 1))
            _ = v5
        } else {
            if len(v4) > 0 {
                v4[0] = 722
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := GenS3{F0: 9223372036854775807, F1: 991}
                v29 := GenS5{F0: 0, F1: fi27, F2: 787}
                for fi30 := 0; fi30 < 17; fi30++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v31 bool = (752 == 1)
                    var v32 int = -9223372036854775808
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    v3 = (-(v3))
    switch "foo bar" {
    case "foo bar":
        v33 := GenS1{F0: -44, F1: ""}
        _ = v33
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    switch 1 {
    case 128:
        switch "\n\t" {
        case "hello":
            var v34 uint32 = 2
            _ = v34
        case "\n\t":
            v3 += -805
        default:
            var v35 uint = 1
            _ = v35
        }
    case -684:
        if true {
            v4 = append(v4, 4294967295)
            v3 += (v3 % (((^(v3))) | 1))
            if (!((0 >= 111) || false)) {
                var v36 int32 = 42
                _ = v36
            }
        } else {
            v37 := GenS6{F0: 4294967295, F1: ("hello" + ("\n\t" + (("hello" + ("a" + "a")) + "hello")))}
            for fi38 := 0; fi38 < 12; fi38++ {
                if true {
                    for fi39 := 0; fi39 < 13; fi39++ {
                        v3 += v3
                        var v40 bool = true
                        _ = fi39
                        _ = v40
                    }
                    if (1 <= 0) {
                        if true {
                            var v41 int = 127
                            _ = v41
                        }
                    } else {
                        var v42 int32 = 1
                        v3 -= 0
                        v43 := GenS5{F0: (v42 >> ((65535) & 63)), F1: (fi38 << ((255) & 63)), F2: 271}
                        _ = v42
                        _ = v43
                    }
                }
                var v44 int8 = 100
                _ = fi38
                _ = v44
            }
            v45 := v1[128]
            _ = v37
            _ = v45
        }
    case 23:
        if (!true) {
            var v46 bool = false
            _ = v46
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v47 := GenS7{F0: 42, F1: 0}
            _ = v47
        }
    }
    var v48 int8 = -67
    var v49 uint64 = 719
    v50 := GenS7{F0: -2147483648, F1: v49}
    for fi51 := 0; fi51 < 16; fi51++ {
        var v52 uint = uint(v49)
        if true {
            var v53 int32 = 2147483647
            v54 := GenS8{F0: v48}
            delete(v2, 2147483647)
            _ = v53
            _ = v54
        } else {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v55 int = fi51
                _ = v55
            }
        }
        _ = fi51
        _ = v52
    }
    v48 = v48
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
