package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint32) int64 {
    v5 := GenS1{F0: true, F1: false, F2: 198}
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS2{F0: 128}
            _ = fi7
            _ = v8
        }
        fmt.Println(p0)
        v9 := GenS2{F0: 1}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        p0 = p0
        _ = fi10
    }
    _ = p0
    _ = v5
    return -585
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 128
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483648
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int = fi11
        v2 = v2
        v3[v4] = ""
        _ = fi11
        _ = v12
    }
    switch 42 {
    case 128:
        if ((!(!false)) && false) {
            v3[(^(v4))] = (("a" + "hello") + "foo bar")
            switch "a" {
            case "":
                switch 128 {
                case 4294967295:
                    fmt.Println(v2)
                case 32768:
                    v4 -= v4
                case 0:
                    v4 += v4
                default:
                    switch -1 {
                    case 0:
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            v4 = 2
                        case "hello":
                            if ("foo bar" > "\n\t") {
                                var v13 float32 = 0.1
                                for fi14 := 0; fi14 < 17; fi14++ {
                                    v2 += 127
                                    if (-1 == -520) {
                                        v15 := []int{32768, (fi14 / (((fi14 - fi14)) | 1)), (fi14 / (((fi14 ^ fi14)) | 1))}
                                        switch fi14 {
                                        case 2147483647:
                                            for fi16 := 0; fi16 < 12; fi16++ {
                                                v13 = v13
                                                var v17 float64 = -2.25
                                                _ = fi16
                                                _ = v17
                                            }
                                        case -1:
                                            v2 -= 42
                                        default:
                                            if len(v15) > 0 {
                                                v15[0] = fi14
                                            }
                                        }
                                        if false {
                                            v15 = append(v15, (fi14 - fi14))
                                            var v18 float64 = 1.5
                                            delete(v3, 1)
                                            _ = v18
                                        } else {
                                            v4 -= v4
                                            var v19 string = ("hello" + "\n\t")
                                            _ = v19
                                        }
                                        _ = v15
                                    } else {
                                        if false {
                                            var v20 int16 = 1000
                                            _ = v20
                                        }
                                        _ = len(v3)
                                    }
                                    _ = fi14
                                }
                                var v21 uint8 = 160
                                _ = v13
                                _ = v21
                            }
                        default:
                            if true {
                                var v22 int = -2147483648
                                _ = v22
                            } else {
                                v23 := GenS2{F0: (v2 - v2)}
                                v24 := GenS2{F0: v2}
                                _ = v23
                                _ = v24
                            }
                        }
                    case 32768:
                        var v25 float64 = 1e10
                        _ = v25
                    default:
                        v2 = 201
                    }
                }
            }
        }
    }
    if false {
        v4 = v4
        v4 -= (v4 >> ((42) & 63))
        v26 := GenS2{F0: 217}
        _ = v26
    }
    switch "hello" {
    case "\n\t":
        var v27 int32 = 2147483647
        _ = v27
    default:
        fmt.Println(v4, v2)
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 uint16 = 2
        fmt.Println(v29, v29)
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v4)
    for fi30 := 0; fi30 < 11; fi30++ {
        v2 = v2
        fmt.Println(v2, v2)
        if true {
            for fi31 := 0; fi31 < 14; fi31++ {
                v32 := GenS3{F0: 65535, F1: 18446744073709551615}
                v4 -= (v4 - (^(v4)))
                _ = fi31
                _ = v32
            }
        }
        _ = fi30
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 uint8 = (v2 | (v2 - 1))
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        switch ("" + "") {
        case "":
            var v36 int32 = 62
            _ = v36
        case "foo bar":
            for fi37 := 0; fi37 < 15; fi37++ {
                fmt.Println(fi35, v4, v4)
                v38 := []int{(fi35 | (fi35 - (fi37 << ((18446744073709551615) & 63)))), 1, fi35}
                _ = fi37
                _ = v38
            }
        case "hello":
            v3[(v4 | (^(v4)))] = ("a" + "hello")
        }
        _ = fi35
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 float32 = (((-2.25 / 1.0) - 3.14159) * float32(((0.1 * 1e10) - (1e10 - 0.0))))
        _ = fi39
        _ = v40
    }
    v41 := []int{483}
    v2 += v2
    for fi42 := 0; fi42 < 15; fi42++ {
        var v43 uint8 = v2
        _ = fi42
        _ = v43
    }
    var v44 int32 = 0
    for fi45 := 0; fi45 < 12; fi45++ {
        v44 -= v44
        v46 := GenS3{F0: 255, F1: 0}
        v2 = (v2 / ((v2) | 1))
        _ = fi45
        _ = v46
    }
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    switch "a" {
    case "\n\t":
        for fi47 := 0; fi47 < 1; fi47++ {
            v48 := v1[2147483648]
            var v49 string = (("hello" + "a") + "\n\t")
            _ = fi47
            _ = v48
            _ = v49
        }
    case "a":
        if (!false) {
            var v50 string = "foo bar"
            _ = v50
        }
    default:
        var v51 float32 = 1.5
        _ = v51
    }
    v52 := GenS1{F0: (257 >= 42), F1: false, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v52
}
