package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 int32; F1 uint64 }
func genFn0() uint {
    var v4 int = -359
    v4 -= 894
    fmt.Println(v4)
    switch (v4 & 32768) {
    case -9223372036854775808:
        fmt.Println(v4)
    case 345:
        var v5 uint64 = 42
        _ = v5
    }
    _ = v4
    return 255
}
func genFn1(p0 int64) int8 {
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 16; fi7++ {
        p0 += -1
        v6 = v6
        _ = fi7
    }
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = v6
    _ = v8
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (1e10 * 0.0)
    var v3 int64 = 42
    v9 := GenS1{F0: -1}
    v10 := GenS2{F0: "hello", F1: 255, F2: -32768}
    v11 := GenS1{F0: 128}
    switch 1 {
    case -2147483648:
        v3 = (-(v3))
    case 2147483647:
        _ = len(v1)
    case 9223372036854775807:
        v3 = 9223372036854775807
    }
    delete(v1, 2147483647)
    switch "a" {
    case "":
        var v12 int16 = -1
        _ = v12
    default:
        if false {
            v13 := []int{-9223372036854775808, 573, -451, 32767}
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 string = "\n\t"
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        if true {
            switch fi16 {
            case 4294967295:
                switch "hello" {
                case "hello":
                    if true {
                        var v17 int16 = 590
                        if true {
                            var v18 int = fi16
                            v19 := GenS2{F0: ("" + ""), F1: 1, F2: 165}
                            _ = v18
                            _ = v19
                        }
                        v17 -= v17
                        _ = v17
                    } else {
                        if false {
                            v1[128] = true
                        } else {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                if ((255 < 217) && false) {
                                    var v21 int64 = (v3 | 1)
                                    switch -1 {
                                    case 32768:
                                        v22 := GenS3{F0: (v3 & (v3 / (((v21 >> ((1) & 63))) | 1))), F1: 2}
                                        _ = v22
                                    case 808:
                                        v23 := GenS4{F0: genFn1((-(v3))), F1: 241, F2: 53}
                                        _ = v23
                                    default:
                                        v21 = (v3 << ((255) & 63))
                                    }
                                    v3 -= v21
                                    _ = v21
                                } else {
                                    fmt.Println(v3, v2)
                                    var v24 float32 = v2
                                    if false {
                                        delete(v1, (fi20 % (((^(fi20))) | 1)))
                                        var v25 float32 = (v24 * v24)
                                        var v26 int64 = (v3 ^ v3)
                                        v3 = v3
                                        delete(v1, (fi20 % ((fi20) | 1)))
                                        _ = v25
                                        _ = v26
                                    } else {
                                        fmt.Println(v3, v2, fi20)
                                        fmt.Println(fi16, fi16)
                                        v2 -= (v24 / v24)
                                        if true {
                                            if (65535 == 2147483647) {
                                                if (956 == 2) {
                                                    v1[fi16] = (258 != 32767)
                                                    v2 = ((0.1 / v24) - (float32(1e10) - v2))
                                                    var v27 int16 = 0
                                                    _ = v27
                                                } else {
                                                    var v28 int64 = v3
                                                    _ = v28
                                                }
                                                v29 := GenS5{F0: 1000, F1: 175}
                                                _ = v29
                                            }
                                        }
                                    }
                                    v30 := GenS4{F0: 100, F1: 128, F2: 32767}
                                    if (!false) {
                                        v3 += v3
                                    }
                                    _ = v24
                                    _ = v30
                                }
                                var v31 float32 = v2
                                _ = fi20
                                _ = v31
                            }
                            fmt.Println(fi16)
                        }
                        v3 = (v3 / ((v3) | 1))
                        if ("" != "\n\t") {
                            fmt.Println(v2)
                            v1[fi16] = true
                            if false {
                                v32 := []int{(^(fi16)), (fi16 + -1), (fi16 >> ((679) & 63)), (fi16 & (fi16 / ((-2147483648) | 1)))}
                                _ = v32
                            }
                        }
                    }
                case "a":
                    for fi33 := 0; fi33 < 10; fi33++ {
                        if true {
                            v3 = (v3 << ((758) & 63))
                            var v34 uint64 = 42
                            var v35 uint8 = 2
                            _ = v34
                            _ = v35
                        }
                        var v36 string = "a"
                        v2 += float32(-2.25)
                        _ = fi33
                        _ = v36
                    }
                case "foo bar":
                    _ = len(v1)
                }
            case 670:
                v2 -= (1e10 * (float32(0.1) * v2))
            }
            v37 := GenS6{F0: 7, F1: 0, F2: 42}
            v2 -= (float32(((-2.25 - (float64(v2) * 3.14159)) - 1.5)) * 0.0)
            _ = v37
        } else {
            if (true && false) {
                var v38 int8 = 1
                delete(v1, 2147483647)
                _ = v38
            }
            var v39 float64 = (((3.14159 * -2.25) + (0.1 - (1e10 * 0.1))) * 1e10)
            v40 := GenS7{F0: 255, F1: 1000, F2: 65535}
            _ = v39
            _ = v40
        }
        switch "hello" {
        case "\n\t":
            if (-2.25 < -2.25) {
                var v41 bool = (((1.5 * -2.25) * float64(v2)) >= -2.25)
                var v42 string = (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")
                _ = len(v1)
                _ = v41
                _ = v42
            }
        case "a":
            fmt.Println(fi16, v2)
        case "hello":
            v43 := []int{fi16, -1, 9223372036854775807}
            _ = v43
        }
        var v44 int16 = 664
        _ = fi16
        _ = v44
    }
    v3 -= v3
    for fi45 := 0; fi45 < 12; fi45++ {
        if (1 <= 65535) {
            fmt.Println(v2)
            if false {
                var v46 uint16 = 1
                _ = v46
            }
        }
        fmt.Println(v3)
        _ = fi45
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        v3 -= (-(v3))
        v48 := GenS6{F0: 611, F1: 70, F2: 4294967295}
        v49 := GenS8{F0: (v3 << ((genFn0()) & 63))}
        _ = fi47
        _ = v48
        _ = v49
    }
    v2 = 3.14159
    v50 := GenS9{F0: 760, F1: 1}
    v51 := GenS10{F0: -2147483648, F1: 0}
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 float64 = 3.14159
        _ = fi52
        _ = v53
    }
    var v54 uint32 = 65535
    var v55 float64 = (-2.25 / ((0.0 * 3.14159) + -2.25))
    for fi56 := 0; fi56 < 11; fi56++ {
        v57 := []int{(fi56 & 32767)}
        _ = fi56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v54
    _ = v55
    _ = v9
}
