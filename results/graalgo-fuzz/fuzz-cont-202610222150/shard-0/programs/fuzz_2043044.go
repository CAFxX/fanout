package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 int; F1 int8; F2 bool }
type GenS6 struct { F0 uint8; F1 int }
func genFn0() uint32 {
    v4 := GenS1{F0: 1000, F1: 42, F2: 2}
    v5 := GenS1{F0: -1, F1: -513, F2: 0}
    var v6 uint16 = 1
    if (int32(v6) > -656) {
        v7 := GenS1{F0: -796, F1: 678, F2: 1}
        _ = v7
    } else {
        fmt.Println(v6)
    }
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() bool {
    var v8 uint8 = 255
    var v9 int32 = int32(v8)
    v10 := GenS2{F0: (v8 << ((18446744073709551615) & 63)), F1: 127}
    var v11 uint64 = 4294967295
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []uint32{}
    v2 := []string{"a", ("" + "\n\t"), ("a" + "hello")}
    var v3 uint = 392
    v12 := GenS2{F0: 0, F1: 127}
    var v13 uint8 = 255
    v14 := GenS1{F0: 1000, F1: 1, F2: 255}
    for fi15 := 0; fi15 < 6; fi15++ {
        v1 = append(v1, 2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v16 := GenS2{F0: v13, F1: fi15}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint32 = 0
        var v19 uint8 = (v13 + (v13 + 255))
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 float64 = 1e10
    v21 := GenS1{F0: -1, F1: 0, F2: 999}
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS1{F0: -32768, F1: -9223372036854775808, F2: 2}
        if false {
            switch (fi22 >> (((^(v3))) & 63)) {
            case 334:
                v13 = v13
            default:
                fmt.Println(v3, v20, v3)
            }
            for fi24 := 0; fi24 < 12; fi24++ {
                if len(v1) > 0 {
                    v1[0] = 42
                }
                v3 -= v3
                _ = fi24
            }
            v25 := GenS2{F0: v13, F1: (fi22 >> ((308) & 63))}
            _ = v25
        } else {
            v3 -= 1
        }
        v20 += 1e10
        _ = fi22
        _ = v23
    }
    if true {
        _ = len(v1)
    } else {
        v26 := GenS1{F0: -32768, F1: 9223372036854775807, F2: 42}
        switch -1 {
        case 2147483647:
            var v27 float64 = float64(1e10)
            _ = v27
        case -69:
            v28 := GenS3{F0: -1}
            _ = v28
        default:
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 float64 = v20
                for fi31 := 0; fi31 < 14; fi31++ {
                    if false {
                        fmt.Println(fi29, v13)
                    } else {
                        v32 := GenS4{F0: 1000, F1: (v13 * 2), F2: 1}
                        _ = v32
                    }
                    v3 = (v3 & v3)
                    v33 := GenS4{F0: 709, F1: (^(v13)), F2: 0}
                    _ = fi31
                    _ = v33
                }
                if false {
                    if true {
                        var v34 bool = false
                        _ = v34
                    } else {
                        var v35 uint16 = 771
                        var v36 uint32 = genFn0()
                        _ = v35
                        _ = v36
                    }
                } else {
                    v13 = v13
                    v37 := GenS3{F0: 42}
                    v38 := GenS5{F0: fi29, F1: 42, F2: true}
                    _ = v37
                    _ = v38
                }
                _ = fi29
                _ = v30
            }
        }
        fmt.Println(v20, v13)
        _ = v26
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        if (2 <= 18446744073709551615) {
            for fi40 := 0; fi40 < 18; fi40++ {
                v41 := GenS3{F0: -1}
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = fi40
                _ = v41
            }
            if (!false) {
                var v42 uint = v3
                _ = v42
            }
            var v43 uint16 = 870
            _ = v43
        } else {
            if len(v1) > 0 {
                v1[0] = 2
            }
            v20 = v20
        }
        if ((0.0 * v20) <= (v20 / v20)) {
            var v44 uint64 = 145
            _ = v44
        }
        var v45 uint8 = v13
        _ = fi39
        _ = v45
    }
    if len(v2) > 0 {
        v2[0] = ("" + "a")
    }
    if true {
        if (1e10 == ((-2.25 - 0.0) / 3.14159)) {
            for fi46 := 0; fi46 < 12; fi46++ {
                v13 -= v13
                if (true || false) {
                    fmt.Println(v13, v3)
                    if true {
                        var v47 uint = (v3 >> ((v3) & 63))
                        v13 += 1
                        for fi48 := 0; fi48 < 19; fi48++ {
                            v49 := GenS6{F0: v13, F1: 32767}
                            if (423 >= 4294967295) {
                                v47 = (v47 * (v47 - 2))
                            }
                            _ = fi48
                            _ = v49
                        }
                        _ = v47
                    }
                } else {
                    v20 -= (v20 * v20)
                }
                var v50 bool = (!false)
                _ = fi46
                _ = v50
            }
        }
    }
    fmt.Println(v13)
    v3 -= 909
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
