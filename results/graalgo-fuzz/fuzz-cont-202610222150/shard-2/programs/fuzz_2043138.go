package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint) uint {
    var v7 float32 = 3.14159
    var v8 float32 = 1.5
    p0 = p0
    v9 := GenS1{F0: (!false), F1: 1, F2: (!(true && (1000 <= 42)))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p1 & p1)
}
func genFn1() uint16 {
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: (2 >= 65535), F1: 0, F2: false}
        var v12 uint8 = 128
        fmt.Println(fi10, fi10)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = -16
    _ = v13
    return 2
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 0
    v3 := make(map[string]int)
    var v4 float32 = 1.5
    var v5 uint16 = 255
    var v6 int = 32767
    v14 := GenS1{F0: (v2 != (v2 / ((v2) | 1))), F1: 42, F2: (true || ((((v4 * v4) * (0.0 * float32(v1))) / ((v4 + 1e10) / v4)) == v4))}
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint16 = (v5 ^ 2)
        switch ("a" + "\n\t") {
        case "\n\t":
            for fi17 := 0; fi17 < 4; fi17++ {
                if true {
                    v1 += float64((((v4 + 1.5) + 3.14159) - v4))
                    v18 := GenS2{F0: (v6 ^ 737)}
                    _ = v18
                } else {
                    v5 -= (v16 | (v5 % ((genFn1()) | 1)))
                }
                var v19 int16 = 1000
                if true {
                    var v20 int16 = (v19 ^ v19)
                    _ = v20
                }
                _ = fi17
                _ = v19
            }
        case "hello":
            var v21 uint = 38
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    v4 += v4
    delete(v3, "hello")
    v22 := []int{}
    var v23 uint16 = 2
    if false {
        var v24 uint8 = v2
        var v25 int16 = 42
        fmt.Println(v5)
        _ = v24
        _ = v25
    }
    fmt.Println(v5, v6, v6)
    _ = len(v3)
    if (false || false) {
        if false {
            var v26 float64 = ((float64(1.5) * (v1 - 1.5)) - (v1 + (1.5 + float64((v4 - (v4 / v4))))))
            _ = v26
        } else {
            v1 -= (((v1 / 3.14159) - (1e10 * 3.14159)) - v1)
            var v27 uint8 = (^(v2))
            _ = v27
        }
    } else {
        if true {
            if (42 > (v2 ^ (^(v2)))) {
                v28 := GenS3{F0: (v6 | (v6 >> ((42) & 63))), F1: int16(v6)}
                var v29 bool = true
                _ = v28
                _ = v29
            }
            var v30 uint8 = 128
            var v31 bool = (!true)
            _ = v30
            _ = v31
        } else {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        }
    }
    if true {
        if false {
            switch v6 {
            case 2147483647:
                fmt.Println(v23, v1, v6)
            case -1:
                if (true || false) {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    if (32767 < 1) {
                        v2 -= 0
                        v32 := GenS4{F0: 42}
                        _ = v32
                    } else {
                        if false {
                            switch v6 {
                            case -520:
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            case 32768:
                                v33 := GenS5{F0: 2147483647, F1: 371}
                                _ = v33
                            case -9223372036854775808:
                                v6 -= 32767
                            default:
                                v4 = (v4 / 0.1)
                            }
                        } else {
                            var v34 float32 = (v4 + 0.0)
                            _ = v34
                        }
                        var v35 int = 127
                        for fi36 := 0; fi36 < 12; fi36++ {
                            fmt.Println(v2, fi36)
                            _ = fi36
                        }
                        _ = v35
                    }
                    if (0 == 0) {
                        fmt.Println(v23, v4)
                    } else {
                        v23 += (v23 - (v5 ^ v23))
                        switch 4294967295 {
                        case -19:
                            for fi37 := 0; fi37 < 4; fi37++ {
                                if (127 <= 1) {
                                    if true {
                                        if true {
                                            v38 := GenS6{F0: (!(65535 <= 42)), F1: 4294967295, F2: 42}
                                            v2 -= v2
                                            v6 = (v6 >> ((546) & 63))
                                            _ = v38
                                        }
                                        v23 = v23
                                    } else {
                                        var v39 uint16 = (v5 ^ (v23 & (^(v23))))
                                        for fi40 := 0; fi40 < 12; fi40++ {
                                            v41 := GenS4{F0: -32768}
                                            if (!((!true) || (0 >= 18446744073709551615))) {
                                                switch "a" {
                                                case "\n\t":
                                                    v6 = fi40
                                                case "foo bar":
                                                    var v42 bool = false
                                                    _ = v42
                                                case "":
                                                    switch 128 {
                                                    case 9223372036854775807:
                                                        var v43 bool = false
                                                        _ = v43
                                                    case 919:
                                                        v44 := GenS2{F0: (-(v6))}
                                                        _ = v44
                                                    }
                                                }
                                                v45 := GenS3{F0: (v6 | v6), F1: 32767}
                                                _ = v45
                                            }
                                            var v46 uint = 2
                                            _ = fi40
                                            _ = v41
                                            _ = v46
                                        }
                                        _ = v39
                                    }
                                    v2 += v2
                                    if false {
                                        v2 -= 65
                                    }
                                } else {
                                    if true {
                                        v2 = 0
                                        fmt.Println(v6, v1)
                                        for fi47 := 0; fi47 < 14; fi47++ {
                                            var v48 uint8 = (v2 / ((128) | 1))
                                            _ = fi47
                                            _ = v48
                                        }
                                    } else {
                                        v22 = append(v22, 1)
                                    }
                                }
                                _ = fi37
                            }
                        case -458:
                            if true {
                                fmt.Println(v2, v23, v5)
                                v6 = -373
                            }
                        case 9223372036854775807:
                            if ((v23 ^ v5) == v23) {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                if false {
                                    fmt.Println(v6, v6)
                                }
                                fmt.Println(v1)
                            }
                        default:
                            for fi49 := 0; fi49 < 11; fi49++ {
                                if false {
                                    v4 += 1e10
                                }
                                var v50 uint16 = v23
                                v51 := GenS5{F0: 794, F1: 22}
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                        }
                    }
                }
            case -846:
                var v52 uint = 0
                _ = v52
            default:
                _ = len(v3)
            }
            fmt.Println(v6, v5, v2)
        }
        for fi53 := 0; fi53 < 10; fi53++ {
            if true {
                v54 := GenS7{F0: 42}
                v55 := GenS5{F0: 2147483647, F1: 18446744073709551615}
                _ = v54
                _ = v55
            } else {
                fmt.Println(v23)
            }
            v5 -= 255
            v56 := GenS8{F0: 0}
            _ = fi53
            _ = v56
        }
    }
    delete(v3, "a")
    var v57 uint64 = 4294967295
    switch v6 {
    case 9223372036854775807:
        v4 -= v4
    case 2147483647:
        fmt.Println(v6, v2)
    default:
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
