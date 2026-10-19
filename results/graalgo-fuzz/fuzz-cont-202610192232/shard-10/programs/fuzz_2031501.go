package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 string; F1 int8; F2 bool }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
type GenS4 struct { F0 string }
func genFn0(p0 uint, p1 bool) int {
    if p1 {
        v5 := GenS1{F0: 2, F1: 55, F2: 42}
        switch "a" {
        case "a":
            var v6 bool = (!p1)
            _ = v6
        }
        _ = v5
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint16 = 23
            p1 = true
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 4294967295, F1: -65, F2: -82}
                fmt.Println(fi9, fi7, p0)
                p1 = (342 > 823)
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        v11 := GenS1{F0: 1, F1: 0, F2: 42}
        p1 = (975 != 298)
        _ = v11
    }
    if true {
        fmt.Println(p0)
        v12 := []int{883}
        _ = v12
    } else {
        var v13 bool = (2 > 1)
        p1 = true
        _ = v13
    }
    _ = p0
    _ = p1
    return 16
}

func main() {
    var v1 int = -2147483648
    v2 := []uint32{}
    v3 := []int{v1, v1, (v1 << ((255) & 63))}
    v4 := make(map[uint32]uint32)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int32 = -1
        v15 = -1
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        if ((v1 % ((genFn0(65535, (genFn0(241, false) != fi16))) | 1)) != v1) {
            var v17 float64 = 0.0
            var v18 int32 = 0
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS2{F0: ("a" + ("a" + ("a" + ("hello" + ("a" + "hello"))))), F1: 14, F2: true}
    v1 = v1
    v20 := GenS2{F0: ("hello" + ("hello" + "a")), F1: 100, F2: false}
    if false {
        if false {
            var v21 uint16 = 255
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v1, v21, v21)
                _ = fi22
            }
            _ = v21
        } else {
            switch "\n\t" {
            case "a":
                var v23 uint16 = 601
                _ = v23
            case "":
                if true {
                    v24 := GenS2{F0: "", F1: 100, F2: (v1 >= (v1 * v1))}
                    if true {
                        switch -286 {
                        case 128:
                            v4[65535] = 0
                        case 973:
                            if false {
                                var v25 float32 = -2.25
                                _ = v25
                            }
                        default:
                            for fi26 := 0; fi26 < 2; fi26++ {
                                v1 -= (v1 - v1)
                                v1 = v1
                                fmt.Println(fi26, v1)
                                _ = fi26
                            }
                        }
                        var v27 float64 = ((float64((0.1 / (-2.25 * float32(3.14159)))) * 0.1) + 0.1)
                        _ = v27
                    } else {
                        var v28 int32 = -1
                        _ = v28
                    }
                    _ = v24
                }
            case "foo bar":
                var v29 float64 = (((0.0 + (0.1 / ((1e10 * float64(3.14159)) - ((((0.1 / 0.1) - float64(1.5)) * 0.0) - ((float64((((0.1 - 0.1) - float32(0.1)) * (((((-2.25 * (float32(1e10) - (float32(float64(float32((float64(1e10) - (0.1 + -2.25))))) * 0.1))) + 1e10) / (0.1 / 1.0)) + -2.25) + 1.5))) - 0.0) - 1e10))))) + ((3.14159 - -2.25) + 0.1)) * (-2.25 * 1e10))
                _ = v29
            default:
                _ = len(v2)
            }
            v1 -= v1
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            v1 += -911
            v1 -= fi30
            var v31 uint8 = 251
            _ = fi30
            _ = v31
        }
    }
    var v32 int8 = -94
    var v33 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v32)
    switch "\n\t" {
    case "\n\t":
        v33 = v33
    default:
        if (true && ((0.1 / 1e10) != 3.14159)) {
            if (2 <= 255) {
                var v34 uint = 42
                v35 := GenS1{F0: 1, F1: v32, F2: (^(v32))}
                var v36 int32 = 36
                _ = v34
                _ = v35
                _ = v36
            } else {
                v37 := GenS3{F0: 93, F1: 9223372036854775807, F2: 42}
                v38 := GenS4{F0: "hello"}
                v4[2] = 426
                _ = v37
                _ = v38
            }
            v39 := GenS3{F0: v32, F1: -990, F2: 4294967295}
            _ = v39
        } else {
            switch (-(v1)) {
            case 128:
                v40 := GenS2{F0: "\n\t", F1: v32, F2: (0.1 != -2.25)}
                _ = v40
            case -9223372036854775808:
                v32 += (v32 & 1)
            case -1:
                _ = len(v4)
            }
            var v41 int16 = 0
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        v43 := GenS4{F0: "a"}
        v4[2147483648] = 1
        v33 = (((!v33) && (-1 == 42)) && (true && (-1 == -798)))
        _ = fi42
        _ = v43
    }
    if true {
        switch "hello" {
        case "foo bar":
            var v44 int16 = -899
            _ = v44
        case "\n\t":
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 uint16 = 42
                _ = fi45
                _ = v46
            }
        case "":
            v1 += 32768
        }
        var v47 float32 = float32(-2.25)
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
