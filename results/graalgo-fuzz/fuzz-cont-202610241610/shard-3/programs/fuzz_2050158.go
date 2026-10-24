package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 float32 = 0.0
    v4 := []string{"a", "\n\t", "", "\n\t"}
    v5 := make(map[int]int)
    var v6 uint = 42
    v6 = (v6 << ((v6) & 63))
    if (!((v2 >> ((v6) & 63)) < v2)) {
        var v7 int16 = 0
        _ = v7
    }
    var v8 int64 = v2
    fmt.Println(v8)
    fmt.Println(v2)
    delete(v1, "")
    if false {
        if ((-1 <= 32767) && (1.5 == (v3 + 0.1))) {
            var v9 uint = v6
            if false {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                switch "a" {
                case "":
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        fmt.Println(v8, v6)
                    }
                case "a":
                    switch "hello" {
                    case "\n\t":
                        for fi10 := 0; fi10 < 12; fi10++ {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                v12 := GenS1{F0: 0, F1: 100}
                                var v13 string = (("foo bar" + ("hello" + "hello")) + "a")
                                v9 += (v9 / (((v9 << (((v9 & (v6 % ((255) | 1)))) & 63))) | 1))
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            var v14 bool = false
                            _ = fi10
                            _ = v14
                        }
                    }
                default:
                    v3 += v3
                }
            }
            v8 = 644
            _ = v9
        } else {
            var v15 float64 = -2.25
            _ = v15
        }
        v8 += (v8 << (((v6 - (v6 - (v6 << (((v6 & 2)) & 63))))) & 63))
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS1{F0: v6, F1: 42}
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v6, v8)
    fmt.Println(v8, v3)
    if true {
        var v18 uint8 = 128
        v19 := GenS2{F0: v6, F1: 1000, F2: -9223372036854775808}
        _ = v18
        _ = v19
    }
    var v20 int32 = 1
    v21 := GenS3{F0: 121, F1: 65535}
    if false {
        var v22 int8 = 127
        if false {
            v23 := GenS3{F0: -37, F1: v6}
            _ = v23
        }
        _ = v22
    }
    if false {
        v24 := GenS1{F0: 1, F1: 1}
        switch 1 {
        case -748:
            switch "foo bar" {
            case "hello":
                var v25 string = "hello"
                _ = v25
            case "a":
                var v26 uint = (v6 - v6)
                _ = v26
            case "foo bar":
                v2 = v2
            default:
                for fi27 := 0; fi27 < 1; fi27++ {
                    fmt.Println(v8, v2, v20)
                    var v28 uint8 = 77
                    _ = fi27
                    _ = v28
                }
            }
        case 1:
            switch "hello" {
            case "hello":
                switch "foo bar" {
                case "":
                    if (true && false) {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            for fi30 := 0; fi30 < 8; fi30++ {
                                var v31 float32 = (v3 + ((((1e10 - (v3 * v3)) + 3.14159) / v3) - 0.0))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v8 -= (v2 | 0)
                                _ = fi30
                                _ = v31
                            }
                            var v32 bool = (false && false)
                            _ = fi29
                            _ = v32
                        }
                        _ = len(v5)
                        v33 := GenS4{F0: (42 < 1)}
                        _ = v33
                    } else {
                        fmt.Println(v8)
                        fmt.Println(v6, v8)
                    }
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case "\n\t":
                for fi34 := 0; fi34 < 18; fi34++ {
                    for fi35 := 0; fi35 < 9; fi35++ {
                        fmt.Println(v2, v20, v2)
                        v20 = int32(v2)
                        _ = fi35
                    }
                    v8 -= -1
                    _ = fi34
                }
            case "a":
                switch 42 {
                case 626:
                    var v36 float32 = 3.14159
                    _ = v36
                }
            default:
                var v37 uint64 = 4294967295
                _ = v37
            }
        case 128:
            if (0 >= 128) {
                var v38 int32 = v20
                _ = v38
            } else {
                v39 := GenS5{F0: v20}
                for fi40 := 0; fi40 < 15; fi40++ {
                    var v41 bool = (!true)
                    _ = fi40
                    _ = v41
                }
                for fi42 := 0; fi42 < 6; fi42++ {
                    v2 -= v2
                    for fi43 := 0; fi43 < 17; fi43++ {
                        fmt.Println(fi43, fi42)
                        v5[(fi42 << ((2) & 63))] = fi42
                        fmt.Println(fi42, v3, v8)
                        v4 = append(v4, "\n\t")
                        v2 = v2
                        var v44 int32 = (^(v20))
                        v45 := GenS1{F0: 255, F1: -1}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = fi42
                }
                _ = v39
            }
        default:
            var v46 int8 = 42
            _ = v46
        }
        if (true && (0 != (-(v8)))) {
            var v47 int32 = v20
            _ = v47
        }
        _ = v24
    }
    v48 := GenS6{F0: 255, F1: -128}
    switch ("" + "a") {
    case "\n\t":
        if (0.0 < 3.14159) {
            if ((true || ((v20 >= (v20 >> ((65535) & 63))) && (true || (true || (true && (1.5 > 1e10)))))) || false) {
                if ((!(218 <= 1)) || (42 > 2147483647)) {
                    fmt.Println(v8, v8, v8)
                } else {
                    if (!false) {
                        v49 := GenS2{F0: 844, F1: 1000, F2: -739}
                        _ = v49
                    }
                    switch "hello" {
                    case "hello":
                        var v50 uint16 = 263
                        _ = v50
                    case "\n\t":
                        _ = len(v1)
                    case "a":
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    default:
                        v1["\n\t"] = 1
                    }
                }
                var v51 uint = 0
                _ = v51
            } else {
                v2 = -9223372036854775808
                delete(v1, (("foo bar" + ("foo bar" + "foo bar")) + ""))
                v52 := v1["hello"]
                _ = v52
            }
        } else {
            v53 := GenS7{F0: true}
            for fi54 := 0; fi54 < 11; fi54++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v55 int8 = 108
                _ = fi54
                _ = v55
            }
            _ = v53
        }
    case "":
        switch "" {
        case "a":
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "\n\t")
            }
        case "":
            if true {
                var v56 int = 32767
                var v57 int8 = -1
                for fi58 := 0; fi58 < 13; fi58++ {
                    for fi59 := 0; fi59 < 11; fi59++ {
                        v60 := GenS8{F0: (^(v57)), F1: v57, F2: 911}
                        _ = fi59
                        _ = v60
                    }
                    _ = fi58
                }
                _ = v56
                _ = v57
            } else {
                var v61 bool = false
                v20 -= 2147483647
                v8 += (v8 / (((v8 | v2)) | 1))
                _ = v61
            }
        }
    default:
        for fi62 := 0; fi62 < 6; fi62++ {
            fmt.Println(v2, v2)
            var v63 int8 = -1
            _ = fi62
            _ = v63
        }
    }
    var v64 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v64
    _ = v8
}
