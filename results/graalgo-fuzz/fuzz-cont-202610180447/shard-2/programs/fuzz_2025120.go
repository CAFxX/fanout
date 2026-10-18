package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint8; F1 string }
func genFn0(p0 int8) int32 {
    var v4 int = -547
    var v5 uint = 676
    v6 := GenS1{F0: 1, F1: 1}
    v7 := []int{v4, (^(v4))}
    var v8 float64 = 3.14159
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 188
}
func genFn1(p0 uint, p1 int) uint64 {
    p1 += (p1 | 1)
    v9 := GenS2{F0: 1, F1: "hello", F2: 796}
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS2{F0: 2147483647, F1: "hello", F2: 65535}
        v12 := GenS3{F0: -99}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = (p0 + (^(p0)))
    switch "hello" {
    case "":
        v14 := GenS4{F0: 456}
        _ = v14
    case "a":
        v15 := []int{-9223372036854775808, (p1 << (((p0 - p0)) & 63))}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 1
    var v2 bool = (false && true)
    v3 := []string{}
    v2 = (!v2)
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v1, v2)
        if (0.1 < 3.14159) {
            for fi17 := 0; fi17 < 2; fi17++ {
                if (float32((1.5 - 3.14159)) > 1e10) {
                    fmt.Println(v1)
                    v18 := GenS5{F0: false}
                    fmt.Println(v2, fi17)
                    _ = v18
                }
                if v2 {
                    switch fi16 {
                    case -1:
                        var v19 float64 = 1.5
                        _ = v19
                    case 42:
                        var v20 string = ("hello" + "\n\t")
                        _ = v20
                    case 32767:
                        switch "\n\t" {
                        case "foo bar":
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    var v21 uint16 = 2
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    switch "" {
                                    case "hello":
                                        v21 -= v21
                                    case "foo bar":
                                        var v22 uint = 930
                                        _ = v22
                                    case "":
                                        var v23 uint16 = 2
                                        _ = v23
                                    }
                                    _ = v21
                                }
                            case "hello":
                                fmt.Println(fi17, v1, v2)
                            case "a":
                                if v2 {
                                    v2 = v2
                                    var v24 uint32 = (v1 << ((42) & 63))
                                    _ = v24
                                } else {
                                    switch (fi16 * int(v1)) {
                                    case 2147483647:
                                        if (1.5 == (1.5 * 0.0)) {
                                            v2 = v2
                                            v25 := GenS6{F0: 81}
                                            v1 += (^(v1))
                                            var v26 float32 = 0.0
                                            v27 := GenS5{F0: (!((fi16 << ((597) & 63)) > (fi16 >> ((968) & 63))))}
                                            if (!v2) {
                                                var v28 uint32 = v1
                                                _ = v28
                                            }
                                            v2 = v2
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                    case -772:
                                        var v29 int64 = -9223372036854775808
                                        _ = v29
                                    case -498:
                                        switch "hello" {
                                        case "foo bar":
                                            if (!true) {
                                                var v30 uint8 = uint8(v1)
                                                _ = v30
                                            }
                                        case "":
                                            var v31 int = (fi16 + 32767)
                                            _ = v31
                                        }
                                        if (1e10 != ((1e10 / 3.14159) / (1e10 / 0.1))) {
                                            var v32 uint32 = v1
                                            if (!v2) {
                                                v33 := GenS7{F0: 42, F1: (v1 ^ v1), F2: 1}
                                                _ = len(v3)
                                                _ = v33
                                            }
                                            var v34 uint8 = 30
                                            _ = v32
                                            _ = v34
                                        }
                                    }
                                    var v35 uint = 1
                                    _ = v35
                                }
                            default:
                                switch fi16 {
                                case 9223372036854775807:
                                    var v36 uint8 = 0
                                    v36 -= v36
                                    _ = v36
                                case 1:
                                    v37 := GenS8{F0: 0, F1: "\n\t"}
                                    _ = v37
                                }
                            }
                        case "hello":
                            v1 = (v1 << ((128) & 63))
                        default:
                            v38 := GenS1{F0: 42, F1: 833}
                            _ = v38
                        }
                    }
                }
                _ = fi17
            }
            var v39 int16 = -365
            v2 = v2
            _ = v39
        }
        _ = fi16
    }
    _ = len(v3)
    v40 := GenS5{F0: false}
    var v41 int32 = 1
    switch 4294967295 {
    case -2147483648:
        v3 = append(v3, "hello")
    case 385:
        v1 = v1
    case 127:
        for fi42 := 0; fi42 < 6; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                var v44 int32 = -659
                v45 := GenS5{F0: v2}
                var v46 int32 = 72
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            _ = fi42
        }
    default:
        var v47 uint = 18446744073709551615
        _ = v47
    }
    _ = len(v3)
    var v48 int32 = genFn0(1)
    for fi49 := 0; fi49 < 17; fi49++ {
        if (true && true) {
            v48 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if true {
                var v50 bool = (-2.25 == 0.1)
                fmt.Println(v1)
                _ = v50
            }
            if (v2 || (831 > 2)) {
                if ((!((1.5 / float32(0.1)) < -2.25)) && v2) {
                    v41 += 0
                    v51 := GenS4{F0: 0}
                    var v52 uint = 4294967295
                    _ = v51
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 7; fi53++ {
                        switch fi53 {
                        case 4294967295:
                            _ = len(v3)
                        default:
                            var v54 uint = 254
                            _ = v54
                        }
                        fmt.Println(v1, v2)
                        v2 = (128 < 34)
                        v1 += v1
                        _ = fi53
                    }
                }
                v3 = append(v3, "")
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v55 int64 = 9223372036854775807
                var v56 uint64 = 90
                fmt.Println(v41, v1)
                _ = v55
                _ = v56
            }
        }
        var v57 uint16 = 255
        var v58 int = (fi49 / (((fi49 + (fi49 / ((32768) | 1)))) | 1))
        _ = fi49
        _ = v57
        _ = v58
    }
    if (!(0 != 1)) {
        var v59 float32 = (0.1 * float32(float64(1e10)))
        switch 84 {
        case -1:
            v41 -= v41
        }
        _ = v59
    }
    if (("hello" + ("a" + "hello")) < "") {
        v2 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v48
}
