package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 uint }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 int) uint32 {
    v4 := []int{p0, p0, p1}
    var v5 float64 = -2.25
    p0 -= (p0 * (p0 | (p0 | p0)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int) uint16 {
    v6 := GenS1{F0: "hello"}
    var v7 uint8 = 128
    switch p0 {
    case 336:
        var v8 int32 = 2147483647
        _ = v8
    case -9223372036854775808:
        var v9 int8 = -1
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 4; fi10++ {
            if true {
                switch "foo bar" {
                case "hello":
                    v11 := []int{p0, 4294967295}
                    _ = v11
                case "a":
                    v12 := []int{p0}
                    _ = v12
                default:
                    if true {
                        var v13 float64 = 3.14159
                        _ = v13
                    }
                }
                var v14 uint64 = 4294967295
                var v15 uint16 = 1
                _ = v14
                _ = v15
            }
            _ = fi10
        }
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            switch p0 {
            case 4294967295:
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int8 = -57
                    switch ("foo bar" + ("foo bar" + "\n\t")) {
                    case "":
                        v7 = 1
                    case "\n\t":
                        var v19 uint32 = 272
                        _ = v19
                    case "hello":
                        if true {
                            if true {
                                if (true || false) {
                                    fmt.Println(v7, v7)
                                    var v20 string = "hello"
                                    _ = v20
                                }
                                v21 := []int{p0, fi17, p0, fi16}
                                switch (fi17 ^ p0) {
                                case 42:
                                    fmt.Println(fi17, v7, v18)
                                case 128:
                                    v22 := GenS2{F0: "\n\t", F1: 282, F2: v7}
                                    _ = v22
                                default:
                                    if len(v21) > 0 {
                                        v21[0] = 42
                                    }
                                }
                                _ = v21
                            } else {
                                fmt.Println(v7, p0)
                            }
                        } else {
                            p0 -= 50
                            fmt.Println(p0)
                            if false {
                                v18 = v18
                                fmt.Println(fi17, v7)
                            }
                        }
                    default:
                        var v23 string = "hello"
                        v18 -= (v18 >> ((65535) & 63))
                        v18 = v18
                        _ = v23
                    }
                    v24 := GenS3{F0: 790, F1: "hello", F2: 0}
                    switch "hello" {
                    case "\n\t":
                        var v25 int16 = -32768
                        _ = v25
                    case "foo bar":
                        fmt.Println(p0, v18, fi16)
                    }
                    _ = fi17
                    _ = v18
                    _ = v24
                }
            case 1:
                fmt.Println(v7)
            default:
                var v26 uint16 = 0
                _ = v26
            }
            _ = fi16
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 983
    v2 := []uint32{2147483648, 42}
    v3 := make(map[int]int)
    for fi27 := 0; fi27 < 7; fi27++ {
        if true {
            var v28 int64 = -448
            if (((("" + "foo bar") + "hello") + "\n\t") >= "hello") {
                var v29 string = "a"
                switch v29 {
                case "":
                    v29 = v29
                default:
                    switch v29 {
                    case "hello":
                        if (!(!(65535 < 57))) {
                            if false {
                                v1 = v1
                            } else {
                                _ = len(v3)
                                v29 = v29
                            }
                            var v30 int16 = 0
                            _ = v30
                        }
                    default:
                        v31 := GenS1{F0: v29}
                        _ = v31
                    }
                }
                v3[42] = fi27
                _ = v29
            }
            _ = v28
        } else {
            v32 := GenS4{F0: 2, F1: (!(fi27 == 127))}
            fmt.Println(v1, v1)
            if false {
                v1 = (^(v1))
            }
            _ = v32
        }
        var v33 uint8 = 0
        fmt.Println(v33, v1)
        _ = fi27
        _ = v33
    }
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS3{F0: 42, F1: "foo bar", F2: 0}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v1 = v1
    v1 = (v1 | v1)
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := GenS5{F0: "hello"}
        var v39 float64 = (1.5 / (3.14159 - 0.1))
        var v40 uint16 = v1
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 bool = true
    if (1 >= 2147483648) {
        v1 -= genFn1(127)
        if v41 {
            v41 = (false && (false || v41))
            v41 = (-407 >= 950)
        } else {
            v42 := GenS1{F0: ""}
            var v43 uint32 = 4294967295
            _ = v42
            _ = v43
        }
    } else {
        _ = len(v3)
        switch "a" {
        case "\n\t":
            switch (("foo bar" + "foo bar") + "foo bar") {
            case "":
                if true {
                    fmt.Println(v41)
                    switch 32767 {
                    case 860:
                        if false {
                            if ((0.0 != (((1.5 * 0.0) - 1e10) * float32(float64((-2.25 + 0.0))))) && ("" > "\n\t")) {
                                delete(v3, -314)
                                var v44 int8 = 127
                                v45 := GenS4{F0: 4294967295, F1: true}
                                _ = v44
                                _ = v45
                            } else {
                                var v46 int32 = -1
                                _ = v46
                            }
                            fmt.Println(v41)
                            var v47 uint32 = 2
                            _ = v47
                        } else {
                            var v48 uint32 = 4294967295
                            _ = v48
                        }
                    }
                } else {
                    fmt.Println(v41)
                }
            case "foo bar":
                fmt.Println(v41, v41)
            }
        case "hello":
            if (4294967295 > 255) {
                _ = len(v2)
                if (true || (false && (-170 == -1))) {
                    if false {
                        v1 = genFn1(0)
                        var v49 uint32 = 174
                        if (18446744073709551615 == 65535) {
                            v2 = append(v2, v49)
                            v50 := GenS4{F0: 255, F1: (!false)}
                            if (v49 < v49) {
                                var v51 uint8 = 128
                                v52 := GenS5{F0: "hello"}
                                v53 := GenS2{F0: "hello", F1: (v49 + (v49 - (v49 << ((88) & 63)))), F2: 172}
                                _ = v51
                                _ = v52
                                _ = v53
                            } else {
                                var v54 int16 = 1
                                _ = v54
                            }
                            _ = v50
                        }
                        _ = v49
                    }
                    fmt.Println(v41, v1)
                    delete(v3, -376)
                }
            }
        case "a":
            v55 := GenS3{F0: int16(v1), F1: "", F2: 444}
            _ = v55
        }
        for fi56 := 0; fi56 < 5; fi56++ {
            var v57 bool = v41
            for fi58 := 0; fi58 < 9; fi58++ {
                var v59 uint16 = v1
                _ = fi58
                _ = v59
            }
            _ = fi56
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
}
