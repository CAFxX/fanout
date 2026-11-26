package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 int32; F1 int8; F2 int32 }
type GenS6 struct { F0 int8; F1 int; F2 string }
func genFn0(p0 int8, p1 int8) int16 {
    v6 := []int{4294967295, -11, -790, -741}
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    var v7 int = 2147483647
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64) uint32 {
    var v9 uint8 = 68
    var v10 uint16 = 55
    if false {
        var v11 int16 = 1000
        v12 := GenS1{F0: 1, F1: ((("a" + "foo bar") + ("" + (("" + "\n\t") + "\n\t"))) + "foo bar")}
        v13 := GenS1{F0: 42, F1: ("" + "a")}
        _ = v11
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 float64 = ((1.5 / 0.1) * -2.25)
                switch fi14 {
                case 122:
                    var v17 uint16 = 32767
                    _ = v17
                default:
                    var v18 int32 = -774
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint64 = 42
            v9 = (v9 >> ((4294967295) & 63))
            _ = fi19
            _ = v20
        }
    }
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 0
    v3 := make(map[int]bool)
    var v4 int8 = 100
    var v5 int32 = 1
    delete(v3, 374)
    var v21 float32 = 1e10
    v2 -= (v2 & (v2 % (((v2 + v2)) | 1)))
    for fi22 := 0; fi22 < 19; fi22++ {
        v21 = (((3.14159 - 1.5) + v21) + v21)
        delete(v3, fi22)
        v2 = v2
        _ = fi22
    }
    v1 -= v1
    var v23 string = ""
    if true {
        var v24 int64 = -56
        _ = v24
    }
    fmt.Println(v21, v4)
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 uint32 = 1
        v27 := GenS2{F0: v26}
        v28 := []int{-2147483648, -9223372036854775808, 0}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    switch "\n\t" {
    case "hello":
        v5 = 4
    case "foo bar":
        v4 -= -1
    case "a":
        v29 := GenS1{F0: (v4 * -63), F1: ("\n\t" + "")}
        _ = v29
    }
    switch v23 {
    case "foo bar":
        v30 := v3[128]
        _ = v30
    case "a":
        if true {
            for fi31 := 0; fi31 < 14; fi31++ {
                v5 = v5
                _ = fi31
            }
        }
    case "hello":
        switch (v23 + ((v23 + "hello") + v23)) {
        case "foo bar":
            if false {
                if (!(1e10 > (-2.25 - 3.14159))) {
                    v1 = (v1 >> ((1) & 63))
                }
            } else {
                var v32 uint8 = 127
                if true {
                    switch 904 {
                    case 128:
                        switch 505 {
                        case 9223372036854775807:
                            var v33 uint64 = v2
                            _ = v33
                        case 165:
                            switch 4294967295 {
                            case 423:
                                fmt.Println(v2, v23, v2)
                            case 42:
                                if false {
                                    v34 := GenS2{F0: 42}
                                    v35 := GenS3{F0: "a"}
                                    _ = v34
                                    _ = v35
                                }
                            case 982:
                                for fi36 := 0; fi36 < 11; fi36++ {
                                    for fi37 := 0; fi37 < 8; fi37++ {
                                        if (v21 == 1.5) {
                                            v38 := []int{(fi36 << ((0) & 63)), (fi36 | (fi37 | 42)), fi37, int(v2)}
                                            switch (fi37 - (^(fi36))) {
                                            case -212:
                                                var v39 float64 = 1e10
                                                _ = v39
                                            default:
                                                fmt.Println(v5, fi37, v2)
                                            }
                                            delete(v3, -734)
                                            switch v23 {
                                            case "a":
                                                v40 := GenS2{F0: uint32(v32)}
                                                _ = v40
                                            case "":
                                                v41 := GenS4{F0: fi36, F1: v23, F2: 65535}
                                                _ = v41
                                            default:
                                                v4 += v4
                                            }
                                            _ = v38
                                        }
                                        _ = fi37
                                    }
                                    v42 := GenS3{F0: v23}
                                    v23 = ((v23 + v23) + (v23 + "foo bar"))
                                    _ = fi36
                                    _ = v42
                                }
                            }
                        case 656:
                            var v43 float64 = (1.5 * (float64(((v21 - -2.25) - v21)) - 1e10))
                            _ = v43
                        default:
                            v4 = -128
                        }
                    case 631:
                        fmt.Println(v4)
                    case -9223372036854775808:
                        var v44 string = "hello"
                        _ = v44
                    }
                    var v45 int8 = 27
                    v46 := GenS2{F0: 4294967295}
                    _ = v45
                    _ = v46
                } else {
                    fmt.Println(v2)
                }
                v4 += v4
                _ = v32
            }
        case "":
            v47 := GenS3{F0: (((v23 + v23) + v23) + "\n\t")}
            _ = v47
        default:
            v21 += (-2.25 + float32(((0.1 * (-2.25 * float64((v21 + v21)))) - 3.14159)))
        }
    }
    switch v23 {
    case "hello":
        v1 = v1
    case "\n\t":
        var v48 int32 = (v5 | v5)
        _ = v48
    default:
        if (genFn1(-9223372036854775808) <= 2) {
            if (float32((float64(((3.14159 / (3.14159 + (v21 - v21))) - (float32(0.1) / (v21 * (0.1 * ((v21 + ((1e10 * v21) - 1e10)) / ((v21 * (v21 * (v21 - (v21 / float32(0.1))))) + ((v21 * v21) - float32(1e10))))))))) + (float64((v21 * float32(1e10))) + float64(((0.0 - (3.14159 + float32(1e10))) * (1e10 * 0.1)))))) >= v21) {
                fmt.Println(v1, v2)
                for fi49 := 0; fi49 < 13; fi49++ {
                    var v50 bool = (false && (false && false))
                    _ = fi49
                    _ = v50
                }
            }
            for fi51 := 0; fi51 < 2; fi51++ {
                for fi52 := 0; fi52 < 13; fi52++ {
                    v53 := v3[(fi51 / (((^(fi52))) | 1))]
                    fmt.Println(v23, fi52)
                    v54 := GenS5{F0: int32(v1), F1: 42, F2: -1}
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi51
            }
            v55 := GenS6{F0: 42, F1: 2147483647, F2: v23}
            _ = v55
        }
    }
    fmt.Println(v23, v2)
    v5 -= v5
    var v56 int16 = -1
    for fi57 := 0; fi57 < 13; fi57++ {
        v1 = (v1 + genFn0((v4 / ((v4) | 1)), (v4 << ((752) & 63))))
        _ = fi57
    }
    var v58 float32 = (v21 + v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
}
