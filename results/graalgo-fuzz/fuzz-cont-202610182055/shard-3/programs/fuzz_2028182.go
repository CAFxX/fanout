package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 int64 }
type GenS5 struct { F0 string; F1 uint32; F2 string }
type GenS6 struct { F0 int16; F1 string; F2 int16 }
type GenS7 struct { F0 int32; F1 bool }
func genFn0(p0 int32) int32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -233}
    _ = p0
    _ = v9
    return -304
}
func genFn1(p0 uint32, p1 uint32) int {
    v10 := []int{}
    if false {
        p0 -= p1
        v11 := GenS2{F0: true, F1: 18446744073709551615}
        _ = v11
    }
    var v12 int8 = 42
    if true {
        v13 := GenS2{F0: false, F1: 1}
        var v14 uint64 = 545
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return 813
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -2147483648
    v3 := []uint32{1, 42, 0}
    var v4 float64 = 1.5
    v5 := []string{}
    var v6 bool = (("foo bar" + "") == (("foo bar" + "hello") + "\n\t"))
    v4 = (v4 / (((((v4 + 0.1) - 0.1) / -2.25) + v4) * (v4 - 3.14159)))
    v15 := GenS3{F0: 518, F1: -332}
    for fi16 := 0; fi16 < 15; fi16++ {
        v4 = 0.1
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 bool = (((1e10 * -2.25) * 0.1) == 1e10)
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v2, v2)
            _ = fi19
        }
        _ = fi16
    }
    var v20 uint8 = 38
    v21 := v1["hello"]
    v6 = v6
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 int = (fi22 % ((genFn1((^(v21)), 342)) | 1))
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: 0, F1: 32767, F2: 42}
    if true {
        v6 = v6
        fmt.Println(v2, v20, v21)
        for fi25 := 0; fi25 < 7; fi25++ {
            if false {
                _ = len(v1)
                if ((1.5 - (3.14159 + 1e10)) != 1.5) {
                    v26 := GenS5{F0: "a", F1: 276, F2: "a"}
                    var v27 float32 = (((1.5 * (-2.25 + -2.25)) / (1.5 * (-2.25 + 0.0))) + (1.5 + (1.5 - 3.14159)))
                    _ = v26
                    _ = v27
                } else {
                    if false {
                        v2 = fi25
                        switch "a" {
                        case "":
                            for fi28 := 0; fi28 < 12; fi28++ {
                                var v29 int8 = 100
                                var v30 float64 = (float64(1e10) / v4)
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                        case "\n\t":
                            var v31 float64 = (-2.25 + (v4 + 0.0))
                            _ = v31
                        default:
                            v32 := GenS6{F0: 32767, F1: "a", F2: 320}
                            _ = v32
                        }
                        v33 := GenS4{F0: 0, F1: 1000, F2: -955}
                        _ = v33
                    } else {
                        var v34 int = (v2 & 251)
                        _ = v34
                    }
                }
            }
            var v35 int8 = 42
            _ = fi25
            _ = v35
        }
    }
    v36 := GenS7{F0: 42, F1: (100 >= -128)}
    for fi37 := 0; fi37 < 14; fi37++ {
        if (100 < 42) {
            v1[""] = (^(v21))
            var v38 uint64 = 836
            if v6 {
                if v6 {
                    v38 = 702
                    if v6 {
                        for fi39 := 0; fi39 < 18; fi39++ {
                            switch "\n\t" {
                            case "foo bar":
                                var v40 bool = true
                                _ = v40
                            case "a":
                                var v41 uint32 = v21
                                _ = v41
                            case "\n\t":
                                v3 = append(v3, (v21 | v21))
                            }
                            var v42 uint32 = (v21 - v21)
                            switch fi37 {
                            case 2147483647:
                                v1["\n\t"] = v42
                            case 127:
                                var v43 float64 = v4
                                _ = v43
                            case 1:
                                var v44 uint64 = v38
                                _ = v44
                            default:
                                v45 := GenS1{F0: 32768}
                                _ = v45
                            }
                            _ = fi39
                            _ = v42
                        }
                        switch v2 {
                        case -9223372036854775808:
                            var v46 uint16 = 1
                            _ = v46
                        }
                    }
                } else {
                    delete(v1, "")
                }
            }
            _ = v38
        } else {
            v47 := GenS5{F0: "a", F1: 4294967295, F2: (("a" + "foo bar") + "")}
            v20 = v20
            v2 -= 32767
            _ = v47
        }
        _ = fi37
    }
    var v48 int = (v2 << ((510) & 63))
    v1[("" + "a")] = (v21 ^ (v21 + (v21 & v21)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
