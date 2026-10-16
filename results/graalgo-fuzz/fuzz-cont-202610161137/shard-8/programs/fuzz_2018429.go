package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    v3 := []int{128}
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    if false {
        var v4 uint16 = 32767
        v5 := GenS1{F0: 4294967295, F1: ("a" + "a"), F2: 1}
        _ = v4
        _ = v5
    } else {
        var v6 float64 = float64(0.1)
        _ = v6
    }
    v1 -= v1
    var v7 int16 = int16(v1)
    if false {
        var v8 bool = false
        _ = v8
    }
    var v9 uint32 = 2
    var v10 float32 = 1.5
    switch 42 {
    case 2147483647:
        v11 := GenS2{F0: 544, F1: -128, F2: (v1 % ((v1) | 1))}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v9 = v9
        switch "hello" {
        case "a":
            if false {
                var v13 uint64 = 26
                var v14 uint = 42
                if ("hello" <= "\n\t") {
                    var v15 uint8 = 128
                    if true {
                        var v16 int32 = -291
                        _ = len(v3)
                        _ = v16
                    } else {
                        fmt.Println(v1, v14)
                        var v17 int8 = 127
                        _ = v17
                    }
                    _ = v15
                } else {
                    var v18 uint8 = 105
                    var v19 bool = ((true && true) && (v7 >= v7))
                    _ = v18
                    _ = v19
                }
                _ = v13
                _ = v14
            }
        case "hello":
            for fi20 := 0; fi20 < 3; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi20
            }
        case "\n\t":
            var v21 float32 = v10
            _ = v21
        default:
            switch fi12 {
            case -221:
                v22 := GenS1{F0: 42, F1: "", F2: 100}
                _ = v22
            case -2147483648:
                v1 += -1
            default:
                v23 := GenS3{F0: (^(v9)), F1: -323}
                _ = v23
            }
        }
        if false {
            if (255 == 757) {
                var v24 int = 4294967295
                _ = v24
            }
            for fi25 := 0; fi25 < 17; fi25++ {
                v10 = -2.25
                _ = fi25
            }
            switch "" {
            case "":
                v26 := GenS3{F0: (v9 | v9), F1: v7}
                _ = v26
            case "a":
                v7 += (v7 & -370)
            default:
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 uint16 = 2
                    if false {
                        var v29 int64 = (v1 << ((705) & 63))
                        _ = v29
                    } else {
                        v3 = append(v3, fi12)
                        var v30 int8 = -90
                        _ = v30
                    }
                    if ((v1 ^ 42) > v1) {
                        if ((42 == 127) && ((v7 != (v7 ^ 1)) && false)) {
                            if false {
                                fmt.Println(fi27, v9)
                                v31 := GenS3{F0: (v9 << ((4294967295) & 63)), F1: 0}
                                if true {
                                    v7 = (^(v7))
                                    v1 = (v1 + (v1 + int64(v28)))
                                }
                                v9 += 23
                                v10 += v10
                                v7 -= (v7 % ((156) | 1))
                                _ = v31
                            }
                            v32 := GenS1{F0: (^(v9)), F1: "a", F2: -33}
                            var v33 int32 = 1
                            _ = v32
                            _ = v33
                        }
                        var v34 float64 = 3.14159
                        _ = v34
                    } else {
                        v28 -= v28
                        var v35 uint16 = 1
                        _ = v35
                    }
                    _ = fi27
                    _ = v28
                }
            }
        }
        _ = fi12
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v9, v9)
    }
    var v36 int16 = 32767
    _ = len(v2)
    for fi37 := 0; fi37 < 20; fi37++ {
        v38 := GenS4{F0: v1, F1: 1, F2: (v9 / ((v9) | 1))}
        _ = fi37
        _ = v38
    }
    if (!(426 < 18446744073709551615)) {
        if (65535 > 42) {
            var v39 int16 = -470
            _ = len(v3)
            v1 += 9223372036854775807
            _ = v39
        } else {
            if false {
                for fi40 := 0; fi40 < 4; fi40++ {
                    if false {
                        v41 := GenS2{F0: 4294967295, F1: -3, F2: v1}
                        switch (fi40 >> ((18446744073709551615) & 63)) {
                        case -9223372036854775808:
                            for fi42 := 0; fi42 < 15; fi42++ {
                                v7 = -1
                                var v43 int8 = -128
                                _ = fi42
                                _ = v43
                            }
                        }
                        var v44 uint16 = 1
                        _ = v41
                        _ = v44
                    }
                    v36 -= (v7 % (((v36 % (((v36 + v7)) | 1))) | 1))
                    _ = fi40
                }
            } else {
                v36 += int16(v9)
            }
            v45 := GenS2{F0: 42, F1: 69, F2: (-(v1))}
            _ = v45
        }
    }
    v10 = v10
    if (-1 > -618) {
        if len(v3) > 0 {
            v3[0] = -520
        }
        var v46 int8 = -73
        v47 := v2[-2147483648]
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v7
    _ = v9
}
