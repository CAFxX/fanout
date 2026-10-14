package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 int; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 uint, p1 uint8, p2 uint16) uint32 {
    v5 := []int{2147483647}
    fmt.Println(p2, p0)
    var v6 uint8 = p1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v7 := GenS1{F0: 127, F1: 323}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 uint) uint64 {
    fmt.Println(p1, p0)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS2{F0: 255, F1: (p0 * (^(p0))), F2: ""}
        var v10 int16 = 525
        p1 += 379
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    p0 = p0
    _ = p0
    _ = p1
    _ = v11
    return (^(v11))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 int = 2147483647
    var v12 uint16 = 2
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            switch fi13 {
            case 9223372036854775807:
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS1{F0: (fi13 - (-(v4))), F1: genFn0(0, 0, (v12 >> ((0) & 63)))}
                    var v16 uint32 = genFn0(681, 0, (v12 << ((42) & 63)))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            v4 -= (^(v4))
            _ = fi13
        }
        var v17 int64 = 0
        _ = v17
    }
    var v18 int16 = 1000
    if (!false) {
        if (1000 >= (v18 & (-(v18)))) {
            var v19 int32 = 0
            _ = v19
        }
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "")
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS3{F0: (v4 % ((v4) | 1)), F1: "foo bar", F2: 100}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v4, v12)
        v4 -= int(v12)
        _ = fi22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v18 -= (-(v18))
        v24 := GenS2{F0: 128, F1: (v12 + v12), F2: "foo bar"}
        if (!((-(fi23)) <= (fi23 << ((247) & 63)))) {
            if (true || false) {
                for fi25 := 0; fi25 < 15; fi25++ {
                    switch "a" {
                    case "a":
                        v4 = (fi23 - 128)
                    case "hello":
                        v2 = append(v2, "")
                    case "foo bar":
                        if (!(46 == 851)) {
                            if false {
                                if ((false || (154 != 856)) || true) {
                                    v26 := GenS4{F0: 85, F1: (v12 - v12)}
                                    v18 = (^(v18))
                                    var v27 int64 = 1
                                    _ = v26
                                    _ = v27
                                }
                                fmt.Println(fi25, fi23, fi25)
                            } else {
                                v18 -= v18
                                v18 = v18
                            }
                            switch fi23 {
                            case 2147483647:
                                v28 := v1[2147483647]
                                _ = v28
                            case 906:
                                delete(v3, v4)
                            case -2147483648:
                                if (v18 >= (v18 ^ (-(v18)))) {
                                    v12 = 32767
                                    fmt.Println(fi23)
                                    fmt.Println(fi23)
                                } else {
                                    if true {
                                        v4 -= fi25
                                        v29 := GenS3{F0: (fi25 % ((fi23) | 1)), F1: ("hello" + (("hello" + "") + "\n\t")), F2: 0}
                                        var v30 uint16 = (v12 >> ((2) & 63))
                                        _ = v29
                                        _ = v30
                                    } else {
                                        if (!false) {
                                            var v31 string = "foo bar"
                                            if (!false) {
                                                if (753 >= 2147483647) {
                                                    if (true && true) {
                                                        fmt.Println(fi23, v31)
                                                        var v32 float32 = (float32(0.0) / 1e10)
                                                        _ = v32
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                }
                                                if true {
                                                    fmt.Println(v31, fi23)
                                                    var v33 float64 = (3.14159 / ((((float64(1.5) - -2.25) - float64((0.1 - 3.14159))) + 0.0) / 1.5))
                                                    _ = v33
                                                } else {
                                                    v18 -= -1
                                                    var v34 float32 = float32(3.14159)
                                                    var v35 uint32 = 2
                                                    if true {
                                                        var v36 uint16 = (v12 % (((v12 ^ v12)) | 1))
                                                        _ = v36
                                                    } else {
                                                        if (!false) {
                                                            var v37 float32 = ((0.1 + (v34 - v34)) * (((1.5 * (1e10 - float32(0.1))) - (float32(1.5) / v34)) / ((float32(-2.25) + (v34 * v34)) + float32(0.1))))
                                                            _ = v37
                                                        }
                                                        v38 := GenS2{F0: 255, F1: v12, F2: ((v31 + v31) + ("a" + ("" + "\n\t")))}
                                                        fmt.Println(v35)
                                                        _ = v38
                                                    }
                                                    _ = v34
                                                    _ = v35
                                                }
                                                var v39 uint16 = (v12 >> ((0) & 63))
                                                _ = v39
                                            }
                                            _ = v31
                                        } else {
                                            if true {
                                                var v40 uint = 459
                                                v41 := GenS5{F0: (v18 * v18), F1: uint64(v12), F2: v12}
                                                if false {
                                                    v4 += (fi25 / ((v4) | 1))
                                                } else {
                                                    _ = len(v3)
                                                    var v42 int32 = -1
                                                    v4 = fi25
                                                    _ = v42
                                                }
                                                _ = v40
                                                _ = v41
                                            } else {
                                                _ = len(v2)
                                            }
                                            v18 -= (v18 - v18)
                                        }
                                        delete(v3, -2147483648)
                                        delete(v1, 2147483647)
                                    }
                                }
                                v43 := GenS4{F0: -689, F1: (v12 ^ 1)}
                                _ = v43
                            }
                            v4 -= fi25
                        } else {
                            var v44 uint = 42
                            v45 := GenS6{F0: uint32(v12), F1: -2147483648, F2: -9223372036854775808}
                            _ = v44
                            _ = v45
                        }
                    default:
                        fmt.Println(fi25, v18)
                    }
                    var v46 float64 = 0.1
                    _ = fi25
                    _ = v46
                }
            }
        }
        _ = fi23
        _ = v24
    }
    v47 := GenS5{F0: (v18 >> ((788) & 63)), F1: 18446744073709551615, F2: 255}
    for fi48 := 0; fi48 < 16; fi48++ {
        v4 += 1
        _ = fi48
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("\n\t" + ("" + "")) {
    case "hello":
        if true {
            switch v4 {
            case 835:
                v4 -= (v4 << ((42) & 63))
            }
            v4 = v4
            v18 = v18
        }
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v49 string = ((("" + "a") + ("hello" + ("hello" + "a"))) + "\n\t")
    v12 = (v12 & (v12 >> ((950) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
}
