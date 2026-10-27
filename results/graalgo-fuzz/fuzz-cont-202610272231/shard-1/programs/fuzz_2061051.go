package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int64; F1 string }
func genFn0() int64 {
    var v6 int8 = 1
    v7 := GenS1{F0: 670, F1: -1, F2: -2147483648}
    if (122 == v6) {
        switch "a" {
        case "":
            var v8 uint64 = 1
            _ = v8
        case "\n\t":
            if true {
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 uint = 1
                    v11 := GenS1{F0: 32767, F1: 42, F2: -423}
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 uint8 = 255
                        v14 := []int{fi9, (-(fi12))}
                        var v15 bool = true
                        v16 := GenS1{F0: 32767, F1: 1, F2: 42}
                        v14 = append(v14, -723)
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int32 = 0
                    switch (-(fi17)) {
                    case 1:
                        if true {
                            v19 := GenS2{F0: 255, F1: false}
                            if (255 > 22) {
                                fmt.Println(fi17, fi17)
                                v20 := GenS3{F0: 127, F1: (fi17 ^ -2147483648)}
                                _ = v20
                            } else {
                                v21 := []int{-23, fi17, fi17}
                                v21 = append(v21, fi17)
                                v22 := GenS1{F0: 829, F1: 0, F2: (v18 & v18)}
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        }
                    case -146:
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := []int{(fi23 - fi17)}
                            var v25 string = ("\n\t" + "hello")
                            var v26 string = "hello"
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                    v27 := []int{fi17, (fi17 * (-(fi17))), fi17, (fi17 - fi17)}
                    _ = fi17
                    _ = v18
                    _ = v27
                }
                if (1 <= 0) {
                    switch 511 {
                    case 4294967295:
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                v28 := []int{-1, 127}
                                var v29 int16 = -773
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                } else {
                    var v30 uint64 = 0
                    for fi31 := 0; fi31 < 12; fi31++ {
                        switch 32767 {
                        case 32767:
                            fmt.Println(v6, v30, v6)
                        case 0:
                            if (true && ((-2.25 / (0.0 - (0.1 / -2.25))) == 1.5)) {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    var v33 int64 = 0
                                    _ = fi32
                                    _ = v33
                                }
                                fmt.Println(fi31)
                                v6 = (v6 / (((^(v6))) | 1))
                            } else {
                                var v34 int32 = int32(v6)
                                v30 -= (v30 & 42)
                                _ = v34
                            }
                        default:
                            switch 0 {
                            case 42:
                                var v35 uint64 = v30
                                _ = v35
                            }
                        }
                        var v36 uint32 = 42
                        if true {
                            fmt.Println(v30)
                            v30 = 0
                            if false {
                                v37 := GenS1{F0: -312, F1: -2147483648, F2: 1}
                                var v38 uint64 = v30
                                _ = v37
                                _ = v38
                            }
                        } else {
                            var v39 int32 = -320
                            _ = v39
                        }
                        _ = fi31
                        _ = v36
                    }
                    v40 := []int{2147483647, 329, -962}
                    _ = v30
                    _ = v40
                }
            } else {
                v6 = (v6 & v6)
            }
        default:
            fmt.Println(v6)
        }
    } else {
        for fi41 := 0; fi41 < 15; fi41++ {
            v6 = v6
            var v42 bool = (true || false)
            _ = fi41
            _ = v42
        }
        var v43 bool = (285 <= 0)
        _ = v43
    }
    _ = v6
    _ = v7
    return -761
}

func main() {
    v1 := []uint32{193, 0, 2147483648, 65535}
    var v2 int8 = -71
    v3 := make(map[uint32]string)
    v4 := []int{2147483647, 2147483647, -482, 2147483647}
    v5 := []uint32{42}
    for fi44 := 0; fi44 < 6; fi44++ {
        v45 := GenS4{F0: genFn0()}
        v46 := v3[4294967295]
        for fi47 := 0; fi47 < 15; fi47++ {
            v46 = (("hello" + v46) + ((v46 + (v46 + v46)) + v46))
            switch 32767 {
            case 0:
                v2 -= v2
            case 32767:
                if true {
                    v46 = v46
                } else {
                    var v48 uint64 = 0
                    v5 = append(v5, 42)
                    _ = v48
                }
            default:
                var v49 uint32 = 2147483647
                _ = v49
            }
            _ = fi47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    v2 = v2
    var v50 int64 = genFn0()
    if true {
        fmt.Println(v50, v50)
        fmt.Println(v2, v2)
    } else {
        for fi51 := 0; fi51 < 13; fi51++ {
            var v52 uint64 = 1
            _ = fi51
            _ = v52
        }
        for fi53 := 0; fi53 < 10; fi53++ {
            v50 = (v50 ^ -1)
            _ = fi53
        }
        if len(v1) > 0 {
            v1[0] = 998
        }
    }
    v5 = append(v5, 1)
    if ((2147483647 == 2147483647) && true) {
        var v54 float32 = 0.1
        var v55 bool = (233 <= 1)
        _ = v54
        _ = v55
    }
    fmt.Println(v50, v50)
    for fi56 := 0; fi56 < 15; fi56++ {
        for fi57 := 0; fi57 < 9; fi57++ {
            v58 := GenS1{F0: 448, F1: 42, F2: -88}
            if false {
                v50 = 305
            } else {
                if len(v5) > 0 {
                    v5[0] = 0
                }
                if (2 <= 2147483647) {
                    v59 := GenS5{F0: 1, F1: fi57}
                    v2 = (v2 & 83)
                    _ = v59
                }
            }
            v50 = genFn0()
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    v60 := GenS6{F0: v50, F1: ""}
    v61 := GenS5{F0: 42, F1: 9223372036854775807}
    v62 := GenS6{F0: v50, F1: ""}
    v63 := GenS6{F0: 570, F1: "foo bar"}
    switch 4294967295 {
    case 0:
        v3[0] = ("" + "a")
    case -120:
        var v64 int16 = 0
        _ = v64
    case -2147483648:
        var v65 int8 = v2
        _ = v65
    default:
        var v66 float32 = ((0.0 * 0.0) - ((3.14159 - 3.14159) - 1e10))
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
