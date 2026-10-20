package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int }
func genFn0() uint64 {
    v7 := []int{1}
    if (true || (!false)) {
        if false {
            var v8 uint8 = 2
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint = 1
        var v11 uint8 = 128
        var v12 uint64 = 1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    v5 := []int{32767}
    var v6 int8 = 97
    var v13 uint = 847
    v14 := GenS1{F0: 898, F1: 9, F2: 2}
    v15 := GenS2{F0: 127, F1: 459, F2: 1}
    fmt.Println(v13)
    switch "a" {
    case "hello":
        for fi16 := 0; fi16 < 18; fi16++ {
            if ((-1 != -1) && true) {
                v1["\n\t"] = "a"
            } else {
                if (437 > 780) {
                    v4[fi16] = 42
                    if false {
                        switch "a" {
                        case "":
                            for fi17 := 0; fi17 < 8; fi17++ {
                                v18 := GenS1{F0: int32(v2), F1: 2, F2: 2}
                                fmt.Println(v6)
                                _ = fi17
                                _ = v18
                            }
                        case "\n\t":
                            v13 -= v13
                        default:
                            fmt.Println(v13, v6)
                        }
                    }
                    var v19 int32 = 2147483647
                    _ = v19
                } else {
                    if true {
                        if false {
                            for fi20 := 0; fi20 < 8; fi20++ {
                                v21 := GenS3{F0: -40}
                                var v22 uint = v13
                                delete(v1, (("foo bar" + "foo bar") + ("foo bar" + "hello")))
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            var v23 int32 = -742
                            var v24 uint = v13
                            _ = v23
                            _ = v24
                        }
                        _ = len(v4)
                        v25 := GenS4{F0: 564, F1: 4294967295, F2: 524}
                        _ = v25
                    } else {
                        fmt.Println(v6)
                    }
                    _ = len(v5)
                    v26 := GenS4{F0: 32767, F1: 216, F2: 4294967295}
                    _ = v26
                }
                var v27 int = (fi16 - 700)
                _ = v27
            }
            switch "a" {
            case "a":
                var v28 uint64 = 18446744073709551615
                _ = v28
            case "":
                var v29 uint8 = 42
                _ = v29
            default:
                if (true || true) {
                    v2 = (^(v2))
                    for fi30 := 0; fi30 < 1; fi30++ {
                        if false {
                            var v31 int8 = v6
                            v13 = v13
                            switch "a" {
                            case "foo bar":
                                v13 = v13
                            case "\n\t":
                                delete(v1, "hello")
                            }
                            _ = v31
                        }
                        var v32 uint = v13
                        if true {
                            v33 := GenS5{F0: 127, F1: 389}
                            v6 = 100
                            _ = v33
                        } else {
                            fmt.Println(fi30)
                            if (!(255 < 1)) {
                                switch (fi30 - fi16) {
                                case -687:
                                    if true {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                case 600:
                                    var v34 uint = (v13 << ((v32) & 63))
                                    _ = v34
                                case 468:
                                    var v35 float32 = 1e10
                                    _ = v35
                                }
                            }
                            switch "hello" {
                            case "a":
                                v2 += v2
                            case "foo bar":
                                fmt.Println(fi16)
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                        var v36 bool = false
                        var v37 int32 = 1
                        v38 := GenS6{F0: (v2 >> (((^(v32))) & 63)), F1: 128}
                        var v39 uint32 = 2147483647
                        _ = fi30
                        _ = v32
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    if (1 <= 42) {
                        v40 := GenS7{F0: (("\n\t" + (("hello" + "\n\t") + (("" + "hello") + "\n\t"))) + "foo bar")}
                        v41 := GenS4{F0: 0, F1: 4294967295, F2: 2}
                        _ = v40
                        _ = v41
                    } else {
                        v2 += v2
                    }
                }
            }
            if ("foo bar" >= ("a" + "\n\t")) {
                fmt.Println(fi16)
                switch (fi16 >> (((v13 >> (((^(v13))) & 63))) & 63)) {
                case -194:
                    fmt.Println(fi16)
                }
                _ = len(v5)
            } else {
                switch "foo bar" {
                case "a":
                    v2 += (^(v2))
                case "foo bar":
                    v13 = v13
                }
                v2 = v2
                delete(v1, "\n\t")
            }
            _ = fi16
        }
    case "":
        switch -452 {
        case 436:
            var v42 uint8 = 255
            _ = v42
        }
    default:
        fmt.Println(v6)
    }
    v43 := GenS8{F0: "\n\t", F1: 4294967295}
    var v44 uint8 = 127
    v6 += (^(v6))
    var v45 string = "\n\t"
    v45 = (v45 + "")
    v46 := v4[32768]
    fmt.Println(v6)
    switch v45 {
    case "a":
        for fi47 := 0; fi47 < 6; fi47++ {
            for fi48 := 0; fi48 < 14; fi48++ {
                v6 -= -128
                v49 := GenS3{F0: v2}
                _ = fi48
                _ = v49
            }
            v45 = v45
            if false {
                _ = len(v5)
                v46 = (v46 << (((^(v13))) & 63))
            } else {
                v45 = v45
                fmt.Println(fi47)
                var v50 bool = true
                _ = v50
            }
            _ = fi47
        }
    }
    if (false && ((v6 >> ((v13) & 63)) > (v6 % ((v6) | 1)))) {
        var v51 int = 4294967295
        fmt.Println(v45, v2)
        v51 += 128
        _ = v51
    }
    v2 = v2
    if (false && (true || false)) {
        if (((0.0 - (1e10 / (float32(3.14159) + 1.5))) > 1e10) && false) {
            v2 -= v2
        }
        delete(v4, 42)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        var v52 int64 = 296
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
