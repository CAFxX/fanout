package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 string; F1 uint64; F2 int64 }
type GenS6 struct { F0 int; F1 bool; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint, p1 int32) uint {
    switch ("hello" + ((("" + (("" + "") + "hello")) + "\n\t") + ("a" + "hello"))) {
    case "\n\t":
        if false {
            for fi6 := 0; fi6 < 7; fi6++ {
                for fi7 := 0; fi7 < 9; fi7++ {
                    if ((false && (!(0 == -128))) && (100 != 100)) {
                        if (!false) {
                            if false {
                                v8 := GenS1{F0: 32767}
                                _ = v8
                            } else {
                                v9 := []int{(fi6 % ((1) | 1))}
                                _ = v9
                            }
                        }
                        if true {
                            var v10 uint64 = 682
                            var v11 int16 = -1
                            v12 := GenS1{F0: v11}
                            switch (-(fi7)) {
                            case -489:
                                v11 += (v11 << (((^(p0))) & 63))
                            case -2147483648:
                                var v13 uint16 = uint16(v10)
                                _ = v13
                            }
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    }
                    _ = fi7
                }
                _ = fi6
            }
            var v14 int16 = -871
            v15 := GenS2{F0: 207, F1: -1}
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := []int{32767, fi16}
                if false {
                    var v18 int8 = -79
                    _ = v18
                } else {
                    if ((905 < 2147483648) && false) {
                        var v19 uint = 0
                        _ = v19
                    }
                    if true {
                        var v20 int16 = 55
                        var v21 uint64 = 0
                        switch int(v21) {
                        case -242:
                            for fi22 := 0; fi22 < 1; fi22++ {
                                var v23 int8 = 100
                                v24 := GenS1{F0: v20}
                                fmt.Println(fi22)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case 1:
                            if false {
                                fmt.Println(p0)
                                if ((p1 << ((p0) & 63)) <= (p1 ^ 42)) {
                                    var v25 int8 = 0
                                    if false {
                                        for fi26 := 0; fi26 < 12; fi26++ {
                                            p1 = (p1 * (p1 >> (((p0 - (^(p0)))) & 63)))
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            _ = fi26
                                        }
                                        v27 := GenS1{F0: int16(fi16)}
                                        _ = v27
                                    }
                                    _ = v25
                                } else {
                                    var v28 uint = (p0 / ((p0) | 1))
                                    _ = v28
                                }
                            } else {
                                v21 = (^(v21))
                                v29 := GenS3{F0: 0}
                                _ = v29
                            }
                        default:
                            for fi30 := 0; fi30 < 12; fi30++ {
                                fmt.Println(fi30)
                                p1 = (-(p1))
                                _ = len(v17)
                                _ = fi30
                            }
                        }
                        _ = v20
                        _ = v21
                    }
                }
                v31 := GenS4{F0: 42, F1: 1, F2: 240}
                _ = fi16
                _ = v17
                _ = v31
            }
            var v32 uint16 = 32767
            v33 := []int{128, 42, 397}
            _ = v32
            _ = v33
        }
    case "a":
        v34 := []int{-260, -1}
        _ = v34
    default:
        v35 := GenS3{F0: (p1 % (((^(p1))) | 1))}
        _ = v35
    }
    var v36 uint64 = 1
    if (!(1 != 1)) {
        v37 := GenS4{F0: 141, F1: 0, F2: 1}
        _ = v37
    } else {
        fmt.Println(v36)
        v38 := []int{}
        p1 += -807
        _ = v38
    }
    v36 += uint64(p0)
    _ = p0
    _ = p1
    _ = v36
    return (^(p0))
}
func genFn1(p0 uint16, p1 uint32) int16 {
    v39 := []int{-9223372036854775808}
    var v40 float64 = 0.1
    _ = p0
    _ = p1
    _ = v39
    _ = v40
    return 32767
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    v3 := []int{1}
    var v4 string = ""
    v5 := make(map[int]uint32)
    var v41 string = v4
    var v42 int16 = 1
    var v43 string = v41
    v42 -= 42
    v4 = v4
    if false {
        switch v43 {
        case "":
            if true {
                v4 = v4
                fmt.Println(v42, v4, v43)
                v44 := GenS2{F0: 121, F1: -54}
                _ = v44
            } else {
                v45 := GenS5{F0: (((("a" + v41) + ("a" + v43)) + (v4 + v4)) + "hello"), F1: 0, F2: 0}
                var v46 int16 = (v42 / (((v42 | 42)) | 1))
                _ = v45
                _ = v46
            }
        case "\n\t":
            v47 := GenS6{F0: 128, F1: false, F2: 247}
            _ = v47
        case "a":
            if true {
                v48 := GenS7{F0: (-(v42)), F1: 0}
                _ = v48
            } else {
                v42 -= (v42 * v42)
                for fi49 := 0; fi49 < 7; fi49++ {
                    var v50 float32 = 3.14159
                    v50 -= v50
                    _ = fi49
                    _ = v50
                }
            }
        default:
            for fi51 := 0; fi51 < 4; fi51++ {
                fmt.Println(v41, v1)
                var v52 uint8 = 128
                _ = fi51
                _ = v52
            }
        }
        v1 = v1
    }
    if false {
        for fi53 := 0; fi53 < 6; fi53++ {
            v42 += v42
            _ = fi53
        }
    }
    v2[1] = 753
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 0)
    var v54 int64 = int64(v42)
    for fi55 := 0; fi55 < 15; fi55++ {
        fmt.Println(v54, v4, fi55)
        var v56 int32 = (v1 + (v1 >> ((42) & 63)))
        _ = fi55
        _ = v56
    }
    fmt.Println(v54, v41, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v54
}
