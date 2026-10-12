package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 string; F2 bool }
type GenS5 struct { F0 uint; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    var v4 int32 = 1
    var v5 int32 = v4
    var v6 float32 = -2.25
    delete(v1, 46)
    if false {
        if (255 > 32767) {
            var v7 uint32 = 2147483647
            for fi8 := 0; fi8 < 12; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    fmt.Println(v4)
                    v10 := GenS1{F0: 255}
                    var v11 float32 = 1.5
                    v1[v7] = true
                    var v12 int16 = 42
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                if true {
                    v13 := GenS2{F0: -1, F1: 1}
                    v14 := []int{fi8, fi8, (fi8 % ((272) | 1))}
                    v1[(v7 << ((60) & 63))] = true
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            var v16 uint8 = 208
                            var v17 int32 = v5
                            fmt.Println(fi8)
                            v4 = v17
                            var v18 uint64 = 1
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 float64 = ((((0.0 + (0.1 - (1e10 * 1e10))) - 1.5) + -2.25) + 0.1)
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS3{F0: v2}
                        _ = v21
                    } else {
                        var v22 int16 = 0
                        var v23 string = v2
                        v24 := []int{(-(fi8)), (^(fi8))}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                v4 -= 1
                _ = fi8
            }
            v25 := GenS4{F0: 127, F1: v2, F2: (417 > 255)}
            _ = v25
            _ = v7
        }
        if ((!true) || (255 > 824)) {
            var v26 uint32 = 519
            switch v2 {
            case "a":
                var v27 uint8 = 29
                _ = v27
            case "hello":
                for fi28 := 0; fi28 < 8; fi28++ {
                    if (1 != 42) {
                        if ((((!(v26 < (v26 * (^(v26))))) && false) || false) && (v6 >= (v6 / ((-2.25 - float32(float64(-2.25))) / ((v6 - 0.0) / 1.5))))) {
                            v29 := GenS5{F0: 1, F1: 225, F2: 212}
                            _ = v29
                        } else {
                            var v30 uint8 = 1
                            _ = v30
                        }
                    } else {
                        v4 += (-(v4))
                        for fi31 := 0; fi31 < 4; fi31++ {
                            delete(v3, v26)
                            _ = fi31
                        }
                    }
                    _ = fi28
                }
            case "\n\t":
                var v32 int8 = 127
                _ = v32
            default:
                var v33 string = v2
                _ = v33
            }
            switch 32767 {
            case -994:
                v2 = v2
            default:
                v34 := GenS2{F0: int64(v26), F1: int8(v5)}
                _ = v34
            }
            _ = v26
        }
    } else {
        switch 892 {
        case 42:
            v6 -= 0.0
        case 32767:
            switch "\n\t" {
            case "\n\t":
                var v35 uint64 = 2
                _ = v35
            default:
                for fi36 := 0; fi36 < 9; fi36++ {
                    v4 = v5
                    v37 := GenS1{F0: 1}
                    _ = fi36
                    _ = v37
                }
            }
        case 0:
            v2 = ("hello" + (v2 + v2))
        default:
            v4 -= v4
        }
        if (!true) {
            var v38 float64 = -2.25
            var v39 float64 = v38
            _ = v38
            _ = v39
        }
        var v40 uint8 = 128
        _ = v40
    }
    v6 -= ((v6 + v6) + float32(3.14159))
    v41 := []int{1, 953, -902, 2147483647}
    for fi42 := 0; fi42 < 3; fi42++ {
        switch 32767 {
        case -2147483648:
            switch v2 {
            case "":
                fmt.Println(v4)
            default:
                var v43 float64 = (0.1 + 1e10)
                _ = v43
            }
        case 42:
            var v44 uint16 = 1
            _ = v44
        }
        _ = fi42
    }
    var v45 float32 = v6
    switch 760 {
    case 431:
        v46 := GenS6{F0: 513}
        _ = v46
    default:
        v47 := GenS1{F0: 1}
        _ = v47
    }
    v41 = append(v41, 2147483647)
    for fi48 := 0; fi48 < 20; fi48++ {
        delete(v1, 264)
        for fi49 := 0; fi49 < 1; fi49++ {
            var v50 int32 = v5
            v6 -= v45
            _ = fi49
            _ = v50
        }
        switch (v2 + (v2 + "\n\t")) {
        case "":
            v41 = append(v41, fi48)
        default:
            v2 = ((v2 + v2) + (v2 + v2))
        }
        _ = fi48
    }
    var v51 int16 = 843
    _ = len(v41)
    var v52 int64 = 9223372036854775807
    v4 -= -357
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
