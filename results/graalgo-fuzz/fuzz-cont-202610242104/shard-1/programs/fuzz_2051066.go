package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int; F1 bool; F2 uint16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS7 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := []uint32{405, 1, 65535}
    var v4 uint32 = 850
    v5 := make(map[string]int)
    var v6 int8 = -128
    if ((128 > 0) || false) {
        v7 := v5[""]
        if (!true) {
            v8 := GenS1{F0: 695}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := v5["\n\t"]
                var v11 int16 = 1
                v7 -= (fi9 | (fi9 - (fi9 ^ (v7 >> ((255) & 63)))))
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(v6, v6, v6)
            var v12 int32 = 2147483647
            _ = v12
        }
        _ = v7
    }
    switch 15 {
    case 42:
        var v13 uint8 = 39
        _ = v13
    case 973:
        _ = len(v5)
    case 32767:
        for fi14 := 0; fi14 < 9; fi14++ {
            v6 = v6
            var v15 int64 = 486
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
    default:
        var v16 bool = (!(188 <= -9223372036854775808))
        _ = v16
    }
    fmt.Println(v4, v4)
    var v17 int32 = -1
    if false {
        var v18 uint8 = 255
        v19 := GenS1{F0: 1000}
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: 0, F1: -171}
        _ = v20
    }
    var v21 int = 32768
    delete(v5, "hello")
    var v22 bool = true
    fmt.Println(v6, v4)
    if (1.5 > 3.14159) {
        var v23 uint8 = 0
        for fi24 := 0; fi24 < 17; fi24++ {
            switch "hello" {
            case "foo bar":
                switch v21 {
                case 754:
                    switch int(v17) {
                    case 761:
                        for fi25 := 0; fi25 < 10; fi25++ {
                            switch "" {
                            case "foo bar":
                                v3 = append(v3, (v4 + v4))
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "hello":
                                v17 -= v17
                            default:
                                var v26 int = (fi24 + (v21 << ((42) & 63)))
                                _ = v26
                            }
                            _ = fi25
                        }
                    case 128:
                        _ = len(v3)
                    default:
                        v23 = (v23 ^ (^(v23)))
                    }
                case 32768:
                    if false {
                        switch "a" {
                        case "":
                            for fi27 := 0; fi27 < 4; fi27++ {
                                v28 := GenS1{F0: int16(v23)}
                                var v29 float64 = 0.1
                                delete(v5, ("" + "hello"))
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        case "foo bar":
                            v30 := GenS3{F0: v21, F1: 42}
                            _ = v30
                        case "a":
                            v21 = (-(v21))
                        default:
                            v17 = v17
                        }
                        v21 = (v21 / ((-410) | 1))
                        v31 := v2[(v4 & 974)]
                        _ = v31
                    }
                case 14:
                    var v32 float64 = (-2.25 + 1e10)
                    _ = v32
                default:
                    _ = len(v3)
                }
            case "a":
                if false {
                    var v33 int64 = -9223372036854775808
                    for fi34 := 0; fi34 < 17; fi34++ {
                        if (1.5 != float32(((1.5 - 0.1) / 0.1))) {
                            v35 := GenS2{F0: v23, F1: fi24}
                            _ = v35
                        }
                        if v22 {
                            v36 := GenS2{F0: (v23 | 1), F1: v21}
                            v4 += (v4 >> ((0) & 63))
                            _ = v36
                        } else {
                            var v37 float32 = -2.25
                            _ = v37
                        }
                        v6 = (v6 / ((int8(v33)) | 1))
                        _ = fi34
                    }
                    _ = v33
                } else {
                    var v38 float64 = float64((1e10 / (1.5 * -2.25)))
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 uint64 = 2
                        _ = fi39
                        _ = v40
                    }
                    switch "a" {
                    case "\n\t":
                        if ((^(v17)) < v17) {
                            for fi41 := 0; fi41 < 20; fi41++ {
                                v17 -= (-(v17))
                                var v42 int8 = 127
                                fmt.Println(v42, fi24, v4)
                                v4 = v4
                                _ = fi41
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 11; fi43++ {
                                v23 += (v23 & v23)
                                var v44 uint = 255
                                _ = fi43
                                _ = v44
                            }
                            v45 := GenS3{F0: (fi24 * v21), F1: 181}
                            _ = v45
                        } else {
                            v46 := GenS4{F0: fi24, F1: true, F2: 235}
                            if (v38 >= 1e10) {
                                if v22 {
                                    fmt.Println(fi24, v21, v22)
                                    for fi47 := 0; fi47 < 1; fi47++ {
                                        if v22 {
                                            var v48 float64 = 1e10
                                            v49 := GenS5{F0: int64(v6), F1: (-(v17))}
                                            v21 -= fi47
                                            _ = v48
                                            _ = v49
                                        }
                                        var v50 uint64 = 304
                                        _ = fi47
                                        _ = v50
                                    }
                                }
                                var v51 int = (v21 + 1)
                                _ = v51
                            } else {
                                v4 -= (v4 & v4)
                                v2[v4] = (v21 - (fi24 >> ((178) & 63)))
                                for fi52 := 0; fi52 < 13; fi52++ {
                                    v21 += (-(fi24))
                                    _ = fi52
                                }
                            }
                            v6 = 127
                            _ = v46
                        }
                    case "a":
                        v53 := GenS6{F0: (v23 * v23), F1: (^(v4)), F2: 0}
                        _ = v53
                    default:
                        var v54 uint8 = 1
                        _ = v54
                    }
                    _ = v38
                }
            default:
                v55 := GenS7{F0: ((v6 >> ((0) & 63)) == (v6 * 27)), F1: 1000, F2: 42}
                _ = v55
            }
            if v22 {
                var v56 uint64 = 919
                _ = v56
            } else {
                v23 = v23
                _ = len(v1)
                v57 := v1[("hello" + "\n\t")]
                _ = v57
            }
            _ = len(v3)
            _ = fi24
        }
        v58 := GenS7{F0: true, F1: 1000, F2: 0}
        _ = v23
        _ = v58
    } else {
        v17 -= 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
