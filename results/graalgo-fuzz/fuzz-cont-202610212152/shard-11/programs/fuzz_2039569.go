package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint, p1 uint) int8 {
    p0 += 0
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := []int{}
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 += p1
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1() string {
    var v9 uint8 = 2
    v9 += (v9 | v9)
    var v10 int8 = 1
    _ = v10
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{4294967295, 42, 42, 1}
    var v3 uint8 = 161
    var v4 uint32 = 2147483648
    var v5 int = 267
    if true {
        var v11 uint64 = 1
        _ = v11
    } else {
        var v12 uint = 255
        v3 = v3
        v13 := GenS1{F0: 2}
        _ = v12
        _ = v13
    }
    v4 += (v4 | (v4 << ((2) & 63)))
    if ((!true) || (-32768 <= 1000)) {
        var v14 int64 = 0
        for fi15 := 0; fi15 < 1; fi15++ {
            if ((v3 <= v3) && false) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    v3 = v3
                    if false {
                        v17 := GenS1{F0: 42}
                        if true {
                            v14 = (-(v14))
                            var v18 uint32 = 42
                            switch (v5 % (((fi16 << ((65535) & 63))) | 1)) {
                            case -237:
                                var v19 int8 = 100
                                _ = v19
                            case 943:
                                if (65535 < 255) {
                                    v18 -= 593
                                    var v20 int16 = 42
                                    var v21 uint64 = 1
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 uint = 898
                                    var v23 float32 = (float32(1e10) / 1.0)
                                    var v24 int64 = (v14 >> ((34) & 63))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            }
                            if (((127 != int8(v14)) && true) || ((v14 + v14) == (v14 / ((v14) | 1)))) {
                                _ = len(v1)
                                if (uint(v4) == 460) {
                                    v25 := v1["a"]
                                    fmt.Println(v18, v5)
                                    _ = v25
                                } else {
                                    switch (fi16 >> ((18446744073709551615) & 63)) {
                                    case 207:
                                        var v26 string = "\n\t"
                                        _ = v26
                                    case 0:
                                        _ = len(v2)
                                    case 128:
                                        var v27 float64 = (1e10 + (3.14159 * (0.0 + -2.25)))
                                        _ = v27
                                    default:
                                        v5 = 32767
                                        v18 += (v18 % (((v4 ^ (v18 ^ v4))) | 1))
                                        v28 := GenS1{F0: (v18 | (v4 * v18))}
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: 421}
                                    _ = v29
                                }
                            } else {
                                fmt.Println(fi16, v4, fi16)
                            }
                            _ = v18
                        }
                        _ = v17
                    }
                    _ = fi16
                }
                v30 := GenS3{F0: v5, F1: 18446744073709551615, F2: 255}
                _ = v30
            }
            if false {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 int8 = 42
                    fmt.Println(v14, fi31)
                    _ = fi31
                    _ = v32
                }
            } else {
                v5 += 832
                for fi33 := 0; fi33 < 20; fi33++ {
                    v5 += 0
                    var v34 float32 = -2.25
                    if len(v2) > 0 {
                        v2[0] = (v4 - (v4 | (^(v4))))
                    }
                    _ = fi33
                    _ = v34
                }
            }
            _ = fi15
        }
        delete(v1, "foo bar")
        _ = v14
    }
    if false {
        v35 := v1["hello"]
        var v36 uint = 65535
        _ = v35
        _ = v36
    }
    v5 = v5
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 bool = (-9223372036854775808 != -1)
        v39 := GenS4{F0: 154, F1: 0}
        for fi40 := 0; fi40 < 13; fi40++ {
            if (0 < 42) {
                v41 := GenS4{F0: 556, F1: 1}
                _ = v41
            } else {
                switch fi40 {
                case 127:
                    v42 := GenS4{F0: 353, F1: 96}
                    _ = v42
                case -2147483648:
                    v5 = fi40
                }
                var v43 uint8 = 16
                v44 := GenS5{F0: 758}
                _ = v43
                _ = v44
            }
            var v45 int = fi40
            _ = fi40
            _ = v45
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    var v46 uint8 = v3
    if false {
        v4 = v4
    } else {
        var v47 uint16 = 255
        var v48 int8 = 1
        for fi49 := 0; fi49 < 7; fi49++ {
            if false {
                var v50 uint64 = 155
                v47 -= 934
                _ = v50
            }
            v51 := GenS6{F0: v4, F1: v3}
            var v52 float64 = (1.5 + float64(0.0))
            _ = fi49
            _ = v51
            _ = v52
        }
        _ = v47
        _ = v48
    }
    v46 = v3
    if ((661 == -9223372036854775808) || true) {
        for fi53 := 0; fi53 < 11; fi53++ {
            fmt.Println(v4)
            _ = fi53
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
