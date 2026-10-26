package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint32) int64 {
    var v7 int64 = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            var v9 bool = true
            var v10 uint = 126
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                if ((0.0 >= float64(0.0)) && ((v7 >> ((42) & 63)) > (v7 >> ((255) & 63)))) {
                    p0 += p0
                    if false {
                        v12 := []int{fi8, fi11}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v12
                    }
                    v13 := []int{}
                    _ = len(v13)
                    _ = v13
                }
                _ = fi11
            }
        }
        _ = fi8
    }
    p0 += (^(p0))
    var v14 int16 = 32767
    _ = p0
    _ = v14
    _ = v7
    return v7
}
func genFn1(p0 int32) uint32 {
    fmt.Println(p0)
    v15 := []int{-2147483648, 744, 842}
    var v16 int16 = 0
    _ = p0
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 int32 = -95
    v2 := make(map[uint32]uint32)
    var v3 int16 = 817
    var v4 uint8 = 1
    v5 := []string{}
    v6 := make(map[string]uint32)
    var v17 string = "a"
    fmt.Println(v1, v17, v3)
    if len(v5) > 0 {
        v5[0] = v17
    }
    v4 = 42
    var v18 uint64 = 1
    if false {
        if false {
            v18 += v18
            if (9223372036854775807 != -374) {
                v3 += v3
                for fi19 := 0; fi19 < 6; fi19++ {
                    v20 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: 688}
                    _ = fi19
                    _ = v20
                }
                switch v17 {
                case "":
                    fmt.Println(v3, v18)
                }
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(v3)
                _ = fi21
            }
        }
        v3 -= v3
    } else {
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(v18, v4)
            var v23 int = fi22
            _ = fi22
            _ = v23
        }
    }
    v5 = append(v5, "foo bar")
    v24 := v2[4294967295]
    v25 := GenS2{F0: (v4 ^ uint8(v24)), F1: 4294967295, F2: 881}
    if (true || true) {
        if true {
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 uint64 = v18
                _ = fi26
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float64 = (1.5 / 1.0)
                v30 := GenS1{F0: (v1 / (((v1 << ((763) & 63))) | 1)), F1: 0}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        v24 -= genFn1((v1 * -775))
        if (((v1 | v1) == 2147483647) && (255 >= 42)) {
            var v31 uint = 65535
            _ = v31
        }
    }
    v4 += (v4 - (v4 - 148))
    for fi32 := 0; fi32 < 19; fi32++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            if true {
                v24 += (v24 >> ((4294967295) & 63))
                v1 -= v1
            } else {
                var v33 bool = (float64(0.0) == 1.5)
                fmt.Println(v1, v1)
                _ = v33
            }
            var v34 string = v17
            if len(v5) > 0 {
                v5[0] = (("foo bar" + "foo bar") + "foo bar")
            }
            _ = v34
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch fi32 {
            case 128:
                var v35 int = fi32
                _ = v35
            case 557:
                v2[(v24 / ((65535) | 1))] = genFn1(1)
            }
            var v36 int16 = v3
            _ = v36
        }
        v37 := GenS2{F0: 255, F1: (fi32 + -939), F2: 4294967295}
        _ = fi32
        _ = v37
    }
    var v38 uint64 = 141
    _ = len(v2)
    if false {
        var v39 int64 = 9223372036854775807
        for fi40 := 0; fi40 < 16; fi40++ {
            v41 := GenS3{F0: 255, F1: ((false || true) || (((true && true) && (((v24 >> ((42) & 63)) >= (v24 & v24)) && true)) && true)), F2: false}
            var v42 uint32 = v24
            switch int(v1) {
            case -2147483648:
                v38 += v18
            case 0:
                var v43 uint64 = v38
                _ = v43
            case 2147483647:
                delete(v6, v17)
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        fmt.Println(v4, v17)
        _ = v39
    } else {
        var v44 string = ""
        var v45 bool = ((genFn0(569) >= 1) && (-1 >= 100))
        var v46 uint8 = v4
        _ = v44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        fmt.Println(v18, v18)
        switch (fi47 * fi47) {
        case 32768:
            switch (v17 + "foo bar") {
            case "a":
                if (((^(v24)) > v24) && (false && true)) {
                    switch "\n\t" {
                    case "a":
                        switch v17 {
                        case "foo bar":
                            var v48 int16 = int16(v24)
                            _ = v48
                        default:
                            if false {
                                for fi49 := 0; fi49 < 5; fi49++ {
                                    v4 -= v4
                                    _ = fi49
                                }
                                fmt.Println(v18, v1, v18)
                                var v50 int32 = -2147483648
                                _ = v50
                            }
                        }
                    case "":
                        var v51 uint64 = 4294967295
                        _ = v51
                    case "foo bar":
                        v38 = (v38 ^ v38)
                    }
                    if (873 < uint(v24)) {
                        for fi52 := 0; fi52 < 14; fi52++ {
                            var v53 uint64 = v38
                            switch "foo bar" {
                            case "hello":
                                if true {
                                    delete(v2, 2147483648)
                                    v53 -= 1
                                } else {
                                    delete(v6, (("a" + v17) + v17))
                                    v54 := GenS2{F0: (v4 ^ uint8(v24)), F1: (fi47 ^ fi52), F2: 839}
                                    v38 -= (v18 >> ((uint(v24)) & 63))
                                    _ = v54
                                }
                            case "":
                                switch "" {
                                case "hello":
                                    var v55 int16 = (-(v3))
                                    _ = v55
                                case "\n\t":
                                    v4 = (v4 | v4)
                                }
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            if false {
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                delete(v2, v24)
                                v38 += (v18 % ((v53) | 1))
                            }
                            _ = fi52
                            _ = v53
                        }
                        v56 := GenS2{F0: v4, F1: fi47, F2: uint(v4)}
                        _ = v56
                    }
                } else {
                    var v57 bool = false
                    for fi58 := 0; fi58 < 10; fi58++ {
                        var v59 int = -567
                        _ = fi58
                        _ = v59
                    }
                    v60 := GenS4{F0: 100}
                    _ = v57
                    _ = v60
                }
            case "\n\t":
                v61 := GenS4{F0: 43}
                _ = v61
            case "foo bar":
                switch fi47 {
                case 128:
                    var v62 int8 = 0
                    _ = v62
                case 1:
                    var v63 uint8 = v4
                    _ = v63
                case -2147483648:
                    var v64 uint32 = v24
                    _ = v64
                }
            }
        case 32767:
            v65 := GenS4{F0: -128}
            _ = v65
        case 789:
            v1 = v1
        default:
            v66 := GenS5{F0: -9223372036854775808}
            _ = v66
        }
        for fi67 := 0; fi67 < 15; fi67++ {
            switch (^(fi47)) {
            case 9223372036854775807:
                var v68 int32 = v1
                _ = v68
            case -2147483648:
                switch (((v17 + ((v17 + (v17 + v17)) + v17)) + v17) + v17) {
                case "":
                    v6[(v17 + v17)] = v24
                }
            case 42:
                var v69 int8 = -128
                _ = v69
            default:
                var v70 int32 = (v1 & -364)
                _ = v70
            }
            v3 += 32767
            fmt.Println(v4)
            v1 -= (v1 | 42)
            v38 = v38
            if (false && ((^(v4)) < v4)) {
                v71 := GenS2{F0: (v4 / ((154) | 1)), F1: (fi67 * fi67), F2: 255}
                v72 := v2[(v24 >> ((65535) & 63))]
                _ = v71
                _ = v72
            }
            _ = fi67
        }
        _ = fi47
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
