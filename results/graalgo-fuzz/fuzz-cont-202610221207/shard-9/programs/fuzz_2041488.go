package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int32; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0(p0 int64, p1 int32) int64 {
    p1 = p1
    v6 := GenS1{F0: -1}
    v7 := []int{32768, -2147483648, 32768}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            p0 = p0
            _ = fi9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            switch fi10 {
            case 32767:
                v7 = append(v7, 32767)
            default:
                p0 += p0
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi10
        }
        _ = fi8
    }
    p0 -= (p0 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    var v11 int16 = 42
    v11 -= (v11 & (v11 >> ((439) & 63)))
    _ = p0
    _ = v11
    return v11
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = 0.0
    var v3 uint64 = 0
    v4 := []uint32{149, 0}
    var v5 int8 = -1
    var v12 int32 = 42
    var v13 uint32 = 4294967295
    var v14 int64 = 1
    v13 = (v13 / (((v13 >> ((65535) & 63))) | 1))
    v15 := GenS2{F0: (("a" + "foo bar") + ""), F1: (^(v13))}
    for fi16 := 0; fi16 < 11; fi16++ {
        if true {
            if (true || (!false)) {
                switch "foo bar" {
                case "\n\t":
                    var v17 uint = 18446744073709551615
                    _ = v17
                case "":
                    v4 = append(v4, (^(v13)))
                case "foo bar":
                    v3 += (v3 - 867)
                default:
                    var v18 float64 = v2
                    _ = v18
                }
                v1 += uint8(v13)
            } else {
                if false {
                    v1 -= (^(v1))
                }
            }
            v12 = 220
            v5 -= (-(v5))
        } else {
            _ = len(v4)
            fmt.Println(v13)
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v2, v2, v13)
                _ = len(v4)
                if true {
                    if true {
                        v20 := GenS2{F0: "hello", F1: 4294967295}
                        var v21 uint64 = v3
                        v21 = 4294967295
                        _ = v20
                        _ = v21
                    }
                    if (!true) {
                        v1 -= v1
                        var v22 bool = (!true)
                        v22 = v22
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v22
                    } else {
                        var v23 uint8 = 82
                        _ = v23
                    }
                    var v24 int = (fi19 ^ (^(fi19)))
                    _ = v24
                } else {
                    var v25 uint64 = (v3 - (v3 >> ((1) & 63)))
                    if true {
                        if (false || true) {
                            if true {
                                v13 = (^(v13))
                                var v26 int32 = v12
                                _ = v26
                            }
                            switch "foo bar" {
                            case "":
                                var v27 string = "hello"
                                _ = v27
                            default:
                                var v28 int = (-(fi19))
                                _ = v28
                            }
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                v29 := GenS3{F0: 42, F1: (-(v12))}
                                _ = v29
                            case "\n\t":
                                _ = len(v4)
                            case "a":
                                if true {
                                    v30 := GenS2{F0: "", F1: (v13 / ((2) | 1))}
                                    v31 := GenS4{F0: (v12 - int32(v25)), F1: int16(fi16), F2: 32767}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if false {
                                        fmt.Println(v3, v14)
                                        v14 = v14
                                        var v32 uint = 1
                                        _ = v32
                                    }
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v33 int8 = (v5 * v5)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v3 -= v25
                                    _ = v33
                                }
                            }
                        } else {
                            var v34 uint32 = 48
                            _ = v34
                        }
                        if false {
                            v35 := GenS3{F0: 0, F1: -520}
                            if (367 <= 4294967295) {
                                v4 = append(v4, (v13 >> ((65535) & 63)))
                                v25 = (v25 - (v25 ^ 656))
                                v13 += (v13 << ((18446744073709551615) & 63))
                            }
                            var v36 int64 = 1
                            _ = v35
                            _ = v36
                        } else {
                            var v37 uint16 = 2
                            _ = v37
                        }
                    }
                    _ = v25
                }
                _ = fi19
            }
        }
        var v38 int64 = (v14 | v14)
        _ = fi16
        _ = v38
    }
    switch 4294967295 {
    case -614:
        switch "foo bar" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 609
            }
        case "foo bar":
            v39 := GenS1{F0: 0}
            _ = v39
        case "a":
            v40 := GenS5{F0: v13, F1: 65535}
            _ = v40
        default:
            switch "hello" {
            case "":
                v3 = (v3 | (v3 * v3))
            default:
                fmt.Println(v13)
            }
        }
    case 127:
        var v41 uint8 = 1
        _ = v41
    case -2147483648:
        if true {
            var v42 float64 = float64(0.0)
            v12 += v12
            _ = v42
        }
    default:
        var v43 bool = ((v5 | (v5 * 42)) <= 42)
        _ = v43
    }
    v44 := GenS3{F0: 589, F1: 42}
    var v45 bool = true
    var v46 bool = v45
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
