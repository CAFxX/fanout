package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0(p0 uint64, p1 string) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int64 = 9223372036854775807
            if true {
                switch fi8 {
                case -9223372036854775808:
                    p1 = "a"
                case 0:
                    var v10 uint32 = 4294967295
                    _ = v10
                default:
                    if true {
                        var v11 int32 = 1
                        p0 = 18446744073709551615
                        v12 := []int{}
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(fi7, fi8)
                        var v13 uint64 = (p0 << ((65535) & 63))
                        _ = v13
                    }
                }
                p0 += p0
                var v14 float64 = (0.0 + 0.0)
                v15 := GenS1{F0: -1}
                _ = v14
                _ = v15
            } else {
                if false {
                    if (-1 != -128) {
                        v9 = (v9 | -9223372036854775808)
                        var v16 int8 = 100
                        var v17 int32 = -2147483648
                        _ = v16
                        _ = v17
                    }
                    var v18 int64 = 9223372036854775807
                    v19 := GenS2{F0: (fi7 << ((0) & 63)), F1: 0}
                    _ = v18
                    _ = v19
                }
                v9 += v9
                v9 += v9
                if true {
                    fmt.Println(p0, fi8)
                    v20 := GenS2{F0: 127, F1: 255}
                    _ = v20
                } else {
                    v21 := []int{(fi8 & (fi7 & (fi7 ^ fi7))), fi7}
                    var v22 float64 = 1e10
                    switch fi7 {
                    case -9223372036854775808:
                        v22 = float64((1e10 + (((1.5 + 0.0) - ((float32(1.5) + (0.1 * 1.5)) - -2.25)) - 0.0)))
                    }
                    var v23 uint8 = 0
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int = 2147483647
        p0 -= p0
        fmt.Println(p1, fi24)
        _ = fi24
        _ = v25
    }
    var v26 bool = (!(1 != 255))
    _ = p0
    _ = p1
    _ = v26
    return 76
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]string)
    v3 := []int{-9223372036854775808, -2147483648, -2147483648}
    var v4 uint16 = 42
    var v5 float64 = (-2.25 * 1e10)
    v6 := make(map[uint32]int)
    var v27 int32 = -2147483648
    if len(v3) > 0 {
        v3[0] = -932
    }
    var v28 uint8 = 1
    if (!false) {
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 uint64 = 4294967295
            if ((v30 & (v30 << ((0) & 63))) != (v30 - (v30 + uint64(v28)))) {
                v31 := v2[2]
                switch (fi29 % (((fi29 | fi29)) | 1)) {
                case -581:
                    if true {
                        var v32 int8 = 127
                        var v33 uint64 = v30
                        _ = v32
                        _ = v33
                    }
                case -626:
                    v4 = (v4 | (v4 >> ((31) & 63)))
                case 352:
                    var v34 uint = 2
                    _ = v34
                }
                _ = v31
            } else {
                fmt.Println(v1, v28, v5)
                v27 = (v27 - (v27 | -1))
            }
            v35 := GenS1{F0: 26}
            _ = fi29
            _ = v30
            _ = v35
        }
    } else {
        v1 -= (((v1 / (v1 * (v1 / (v1 / v1)))) - v1) * v1)
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 11; fi37++ {
                fmt.Println(v28, fi36, v1)
                v28 -= 137
                var v38 string = "\n\t"
                var v39 uint32 = 0
                var v40 uint16 = 32767
                switch v38 {
                case "":
                    v28 = 127
                case "hello":
                    if true {
                        v41 := GenS3{F0: (^(v39)), F1: ((((v1 / float32((v5 + v5))) + -2.25) - -2.25) <= ((1e10 / v1) * v1)), F2: ((v38 + v38) >= v38)}
                        if (("hello" + ((v38 + v38) + ("foo bar" + ("\n\t" + v38)))) < "hello") {
                            _ = len(v3)
                        } else {
                            if len(v3) > 0 {
                                v3[0] = -651
                            }
                        }
                        v42 := v2[(v39 * v39)]
                        _ = v41
                        _ = v42
                    } else {
                        if (!(v5 >= (3.14159 / -2.25))) {
                            if true {
                                var v43 uint64 = 75
                                v38 = v38
                                _ = v43
                            }
                            v44 := GenS4{F0: 1, F1: (v27 | -597)}
                            v1 -= 1.5
                            _ = v44
                        } else {
                            switch fi37 {
                            case 42:
                                v38 = v38
                            default:
                                v45 := GenS1{F0: 100}
                                _ = v45
                            }
                        }
                    }
                default:
                    v46 := GenS3{F0: (v39 << ((18446744073709551615) & 63)), F1: false, F2: false}
                    v3 = append(v3, (fi36 | fi37))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2[(^(v39))] = v38
                    var v47 uint8 = (v28 & (v28 >> ((1) & 63)))
                    v27 = (v27 - -1)
                    delete(v2, 674)
                    _ = v46
                    _ = v47
                }
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            if len(v3) > 0 {
                v3[0] = fi36
            }
            _ = fi36
        }
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        v49 := v2[0]
        var v50 int32 = 0
        _ = fi48
        _ = v49
        _ = v50
    }
    var v51 float32 = v1
    v2[uint32(v28)] = "a"
    v5 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
