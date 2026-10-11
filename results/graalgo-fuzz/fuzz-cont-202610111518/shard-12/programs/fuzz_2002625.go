package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool }
func genFn0() string {
    v5 := GenS1{F0: false}
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint32 = 2
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := []int{fi9}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 string, p1 uint32) int8 {
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: (!(!(!false)))}
        p1 = (p1 & (p1 ^ p1))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 bool = true
        p1 += p1
        var v16 int32 = -334
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{-2147483648, -240, 128}
    if false {
        fmt.Println(p0, p0)
        if false {
            v11 = p0
            fmt.Println(p1, p0, v11)
            v18 := GenS1{F0: true}
            _ = v18
        } else {
            var v19 float64 = 0.1
            var v20 uint64 = 534
            _ = v19
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return 100
}

func main() {
    v1 := []uint32{65535, 0}
    v2 := make(map[int]int)
    v3 := []uint32{65535}
    var v4 int = 2147483647
    v4 = (^(v4))
    var v21 int32 = 979
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(fi22, fi22)
        if ((v21 - (v21 ^ v21)) != (v21 - (v21 ^ v21))) {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint8 = 42
                _ = fi23
                _ = v24
            }
        } else {
            if (("\n\t" + "") != "foo bar") {
                v21 = v21
                var v25 uint16 = 2
                if len(v3) > 0 {
                    v3[0] = 2
                }
                _ = v25
            }
        }
        var v26 int = -966
        _ = fi22
        _ = v26
    }
    var v27 int16 = 0
    var v28 int32 = v21
    v29 := GenS2{F0: 2}
    if true {
        var v30 float32 = float32(1e10)
        _ = v30
    } else {
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 uint8 = 127
            _ = fi31
            _ = v32
        }
        if (-2147483648 >= (v21 << ((255) & 63))) {
            v33 := GenS2{F0: 18446744073709551615}
            if (42 <= 681) {
                var v34 int = -2147483648
                v2[(v4 >> ((42) & 63))] = v4
                _ = v34
            } else {
                var v35 int64 = 655
                _ = v35
            }
            if (v28 != (v28 + 2147483647)) {
                var v36 uint16 = 65535
                v4 = (v4 - 4294967295)
                _ = v36
            }
            _ = v33
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 uint = 1
            var v39 string = "foo bar"
            switch fi37 {
            case 42:
                v40 := GenS2{F0: 18446744073709551615}
                _ = v40
            case 127:
                var v41 uint8 = 172
                _ = v41
            case 32767:
                var v42 int32 = v21
                _ = v42
            default:
                for fi43 := 0; fi43 < 3; fi43++ {
                    switch (fi37 | -9223372036854775808) {
                    case 128:
                        v44 := GenS3{F0: -966, F1: 18446744073709551615, F2: v21}
                        _ = v44
                    case 499:
                        v45 := GenS2{F0: 0}
                        _ = v45
                    default:
                        var v46 int64 = 0
                        _ = v46
                    }
                    var v47 float32 = (1e10 * 1e10)
                    var v48 uint = (v38 % ((v38) | 1))
                    var v49 bool = (v21 != v21)
                    _ = fi43
                    _ = v47
                    _ = v48
                    _ = v49
                }
            }
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    fmt.Println(v28, v21, v27)
    v50 := GenS4{F0: 988, F1: "foo bar"}
    var v51 uint8 = 127
    _ = len(v2)
    fmt.Println(v27, v21, v28)
    var v52 uint16 = 65535
    var v53 int32 = v21
    v53 = (-(v53))
    if (768 == 781) {
        v54 := GenS4{F0: 65535, F1: "hello"}
        var v55 bool = (255 == v52)
        for fi56 := 0; fi56 < 12; fi56++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi56
        }
        _ = v54
        _ = v55
    } else {
        _ = len(v1)
        if true {
            if false {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v57 float64 = (1.5 + -2.25)
                if (!(v53 >= (v28 & v28))) {
                    if (!(0 != 821)) {
                        for fi58 := 0; fi58 < 9; fi58++ {
                            v57 += ((((float64(1e10) / v57) * (v57 - (((-2.25 * v57) / 3.14159) + v57))) / -2.25) * 0.0)
                            _ = fi58
                        }
                        var v59 bool = false
                        if (0 < 42) {
                            var v60 int32 = int32(v4)
                            for fi61 := 0; fi61 < 1; fi61++ {
                                v52 -= 425
                                for fi62 := 0; fi62 < 10; fi62++ {
                                    v27 += (v27 - -811)
                                    v57 = (v57 + (v57 * -2.25))
                                    fmt.Println(v53, v51, fi62)
                                    _ = fi62
                                }
                                v63 := GenS5{F0: v59}
                                _ = fi61
                                _ = v63
                            }
                            var v64 uint16 = v52
                            _ = v60
                            _ = v64
                        } else {
                            v21 -= (^(v28))
                        }
                        _ = v59
                    }
                    for fi65 := 0; fi65 < 1; fi65++ {
                        var v66 float64 = 0.1
                        v27 = v27
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        _ = fi65
                        _ = v66
                    }
                    v52 = (v52 | v52)
                } else {
                    delete(v2, (v4 ^ v4))
                    if true {
                        v67 := GenS1{F0: (1 <= 1)}
                        for fi68 := 0; fi68 < 16; fi68++ {
                            v52 = v52
                            if true {
                                var v69 float64 = 0.0
                                var v70 uint8 = (v51 - (v51 ^ 2))
                                _ = v69
                                _ = v70
                            } else {
                                var v71 uint32 = 434
                                var v72 uint = 113
                                _ = v71
                                _ = v72
                            }
                            v73 := GenS5{F0: (true || (42 > -30))}
                            _ = fi68
                            _ = v73
                        }
                        _ = v67
                    }
                }
                _ = v57
            }
            fmt.Println(v28, v21)
        }
    }
    v74 := GenS1{F0: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v74
}
