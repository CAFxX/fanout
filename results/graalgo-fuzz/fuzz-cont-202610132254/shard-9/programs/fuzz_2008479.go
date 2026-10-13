package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 string }
func genFn0() int16 {
    if false {
        var v4 uint16 = 42
        var v5 uint64 = 1
        var v6 bool = (1.5 == 1.5)
        _ = v4
        _ = v5
        _ = v6
    } else {
        if (false && false) {
            v7 := GenS1{F0: 32767, F1: 18446744073709551615, F2: 716}
            _ = v7
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint64 = 904
            v10 := []int{fi8}
            var v11 float32 = float32(1.5)
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int = fi12
            v13 = (fi12 | 127)
            _ = fi12
            _ = v13
        }
    }
    var v14 string = "a"
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint16 = 1
        var v17 bool = (false && (false || true))
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v14
    return -32768
}
func genFn1(p0 uint16, p1 int32, p2 uint32) int32 {
    switch "foo bar" {
    case "":
        fmt.Println(p2, p2)
    case "hello":
        v18 := []int{}
        _ = v18
    case "a":
        v19 := GenS2{F0: 444}
        _ = v19
    }
    if false {
        v20 := GenS2{F0: 0}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint32 = 1
    var v2 int64 = -405
    var v3 float64 = float64(float32(3.14159))
    var v21 uint16 = 42
    v21 = (^(v21))
    v22 := GenS1{F0: (v21 * (v21 & (^(v21)))), F1: 85, F2: 423}
    var v23 float64 = (v3 - v3)
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint16 = 570
        var v26 int16 = -32768
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = -1
    v28 := GenS3{F0: (v2 >> ((156) & 63))}
    v2 = v2
    v29 := []int{}
    var v30 bool = false
    v23 = float64(-2.25)
    for fi31 := 0; fi31 < 16; fi31++ {
        if true {
            v32 := GenS2{F0: 377}
            v33 := GenS4{F0: "", F1: 32767}
            v29 = append(v29, fi31)
            _ = v32
            _ = v33
        } else {
            fmt.Println(v21, v30, v21)
            if (true && v30) {
                v34 := GenS4{F0: "", F1: 32767}
                switch 9223372036854775807 {
                case 1:
                    v35 := GenS4{F0: "a", F1: 1000}
                    _ = v35
                case 32768:
                    if v30 {
                        switch "" {
                        case "a":
                            v21 -= 470
                        case "hello":
                            switch "a" {
                            case "\n\t":
                                v36 := GenS3{F0: v2}
                                _ = v36
                            case "a":
                                switch (fi31 - (^(fi31))) {
                                case 9223372036854775807:
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                default:
                                    var v37 int8 = -1
                                    _ = v37
                                }
                            case "hello":
                                if len(v29) > 0 {
                                    v29[0] = fi31
                                }
                            }
                        default:
                            if ((false || v30) && v30) {
                                v38 := GenS4{F0: ("a" + ("foo bar" + "foo bar")), F1: int16(v2)}
                                if (v21 == (v21 ^ v21)) {
                                    v1 += 445
                                }
                                _ = v38
                            }
                        }
                        var v39 int16 = 896
                        v40 := GenS5{F0: (v1 + (^(v1)))}
                        _ = v39
                        _ = v40
                    }
                case -504:
                    v41 := GenS3{F0: (v2 + int64(v1))}
                    _ = v41
                }
                for fi42 := 0; fi42 < 11; fi42++ {
                    switch 2147483647 {
                    case -813:
                        v43 := GenS6{F0: 708}
                        _ = v43
                    }
                    if len(v29) > 0 {
                        v29[0] = fi42
                    }
                    _ = fi42
                }
                _ = v34
            } else {
                v21 += v21
                for fi44 := 0; fi44 < 18; fi44++ {
                    fmt.Println(v30, v2, v21)
                    v3 = (1.5 / (v3 + 1e10))
                    if ("a" <= "") {
                        switch 9223372036854775807 {
                        case 2147483647:
                            v21 -= v21
                        case -991:
                            fmt.Println(v30, v30, v30)
                        default:
                            if (v27 > int8(fi44)) {
                                v23 = ((v3 * 1e10) / 1e10)
                            } else {
                                v23 = 3.14159
                                var v45 int32 = genFn1(v21, 42, (v1 % ((v1) | 1)))
                                fmt.Println(v27, fi31, fi31)
                                _ = v45
                            }
                        }
                    } else {
                        v2 = v2
                    }
                    fmt.Println(v23, fi44, fi31)
                    _ = fi44
                }
            }
        }
        _ = fi31
    }
    switch ("" + ("\n\t" + "")) {
    case "hello":
        v23 = float64((1e10 * -2.25))
    case "foo bar":
        fmt.Println(v3)
    case "a":
        fmt.Println(v21)
    }
    switch ("foo bar" + ("a" + "a")) {
    case "":
        v46 := GenS7{F0: 32767, F1: "a"}
        _ = v46
    }
    v1 = v1
    var v47 uint64 = 796
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v47
}
