package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
func genFn0() int16 {
    v5 := GenS1{F0: 127, F1: 177, F2: 32767}
    var v6 int32 = -1
    _ = v5
    _ = v6
    return -117
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    var v3 int8 = 43
    var v4 int64 = (v1 - 408)
    var v7 uint32 = 2
    for fi8 := 0; fi8 < 8; fi8++ {
        v3 = v3
        v1 = (v1 << ((255) & 63))
        _ = fi8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            var v10 float64 = 0.0
            var v11 float64 = 3.14159
            v2[(v7 >> ((56) & 63))] = false
            _ = v10
            _ = v11
        }
        v3 += v3
        _ = fi9
    }
    switch "a" {
    case "hello":
        if true {
            if true {
                fmt.Println(v4)
                if false {
                    if false {
                        v1 += -209
                    } else {
                        if true {
                            var v12 int64 = v1
                            _ = v12
                        } else {
                            v13 := []int{0, -2147483648, -1}
                            var v14 string = "hello"
                            v1 = (v1 >> ((18446744073709551615) & 63))
                            _ = v13
                            _ = v14
                        }
                        var v15 int8 = int8(v7)
                        _ = v15
                    }
                    v16 := GenS1{F0: 36, F1: 1, F2: 42}
                    v17 := v2[0]
                    _ = v16
                    _ = v17
                }
            }
        } else {
            v4 = 0
        }
    }
    var v18 string = ("foo bar" + (("\n\t" + ("a" + "a")) + ((("a" + "\n\t") + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t"))))
    if true {
        v4 += (v1 % ((1) | 1))
    }
    if true {
        if false {
            delete(v2, v7)
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS1{F0: 0, F1: 255, F2: 255}
                for fi21 := 0; fi21 < 16; fi21++ {
                    v22 := []int{}
                    delete(v2, (v7 * 42))
                    var v23 uint32 = 228
                    if (65535 < 65535) {
                        _ = len(v22)
                        fmt.Println(v3, v4)
                    }
                    switch "foo bar" {
                    case "\n\t":
                        switch 127 {
                        case -745:
                            v3 += v3
                        case -9223372036854775808:
                            _ = len(v22)
                        case 32767:
                            switch (fi21 - (fi19 * fi21)) {
                            case 380:
                                v3 = v3
                            case 691:
                                v23 = (v23 >> ((2) & 63))
                                fmt.Println(fi19)
                                var v24 uint64 = 42
                                _ = v24
                            case 1:
                                var v25 uint16 = 65535
                                _ = v25
                            }
                        default:
                            v1 -= v1
                        }
                    case "hello":
                        switch (fi21 + fi19) {
                        case -877:
                            fmt.Println(v1)
                        case 9223372036854775807:
                            var v26 uint32 = (v23 >> ((1) & 63))
                            _ = v26
                        case 127:
                            if false {
                                var v27 uint64 = 146
                                v3 = 42
                                v7 = (v7 / ((v7) | 1))
                                _ = v27
                            }
                        }
                    case "foo bar":
                        v28 := GenS2{F0: 2}
                        _ = v28
                    default:
                        v4 = v1
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v1 = (v1 % (((v4 - (v4 | v4))) | 1))
                _ = fi19
                _ = v20
            }
        } else {
            var v29 uint8 = 121
            if ((0.0 != (0.1 * 1e10)) || false) {
                if ((v29 % (((v29 - v29)) | 1)) < (v29 / ((v29) | 1))) {
                    fmt.Println(v1)
                    delete(v2, (v7 + v7))
                } else {
                    if true {
                        switch ((((v18 + v18) + v18) + (v18 + v18)) + (v18 + ((("\n\t" + v18) + v18) + ((v18 + "\n\t") + "hello")))) {
                        case "foo bar":
                            v7 = (v7 << ((1) & 63))
                        case "":
                            fmt.Println(v3)
                        case "\n\t":
                            var v30 uint = 1
                            _ = v30
                        default:
                            v7 -= v7
                        }
                        var v31 uint32 = v7
                        _ = v31
                    } else {
                        if ((v29 <= 128) || false) {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                v33 := []int{0}
                                fmt.Println(v3)
                                v29 = v29
                                _ = fi32
                                _ = v33
                            }
                            fmt.Println(v4)
                            for fi34 := 0; fi34 < 16; fi34++ {
                                var v35 float32 = -2.25
                                _ = fi34
                                _ = v35
                            }
                        }
                        v36 := []int{-9223372036854775808, 1, -2147483648}
                        var v37 int8 = v3
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS3{F0: v7, F1: 1, F2: (v4 << ((65535) & 63))}
                    _ = v38
                }
                if false {
                    delete(v2, (v7 / ((2147483647) | 1)))
                    var v39 string = (((v18 + (v18 + v18)) + ((v18 + v18) + "hello")) + ("hello" + v18))
                    _ = v39
                }
            }
            for fi40 := 0; fi40 < 15; fi40++ {
                switch v18 {
                case "foo bar":
                    v41 := GenS1{F0: v29, F1: v29, F2: 32767}
                    _ = v41
                }
                _ = fi40
            }
            _ = v29
        }
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 12; fi43++ {
            if true {
                v44 := GenS3{F0: (v7 ^ (v7 << ((1) & 63))), F1: 255, F2: (v1 - (^(v4)))}
                _ = v44
            } else {
                if (42 <= 0) {
                    v7 = v7
                }
            }
            var v45 int32 = -313
            _ = fi43
            _ = v45
        }
        v7 = 308
        _ = fi42
    }
    switch -1 {
    case -9223372036854775808:
        v3 = 42
    }
    v46 := []int{-355, -951}
    for fi47 := 0; fi47 < 3; fi47++ {
        _ = len(v46)
        if true {
            var v48 uint32 = v7
            v49 := GenS2{F0: 65535}
            _ = v48
            _ = v49
        } else {
            var v50 int8 = 1
            fmt.Println(v18, v3)
            v51 := GenS1{F0: 127, F1: 1, F2: 382}
            _ = v50
            _ = v51
        }
        v4 += v1
        _ = fi47
    }
    v1 = v4
    var v52 uint64 = 848
    v53 := GenS2{F0: 0}
    if true {
        v1 += v4
    }
    if true {
        v54 := GenS4{F0: -636, F1: v7, F2: 255}
        if (682 <= 1) {
            for fi55 := 0; fi55 < 2; fi55++ {
                v56 := GenS3{F0: (v7 & (v7 - v7)), F1: 129, F2: v4}
                _ = fi55
                _ = v56
            }
            v57 := GenS4{F0: genFn0(), F1: (^(v7)), F2: 57}
            for fi58 := 0; fi58 < 2; fi58++ {
                if (v4 != 236) {
                    var v59 uint8 = 38
                    _ = len(v2)
                    fmt.Println(v59, v4, v1)
                    _ = v59
                }
                _ = fi58
            }
            _ = v57
        } else {
            v46 = append(v46, -1)
        }
        _ = v54
    }
    v18 = (v18 + v18)
    var v60 int = 329
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v52:", v52)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v52
    _ = v53
    _ = v60
    _ = v7
}
