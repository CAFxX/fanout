package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32 }
func genFn0() bool {
    var v7 int8 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(fi8, v7)
            _ = fi9
        }
        v7 -= v7
        v10 := []int{fi8, fi8}
        _ = fi8
        _ = v10
    }
    _ = v7
    return false
}
func genFn1(p0 int8) int32 {
    var v11 int8 = p0
    v12 := GenS1{F0: 42, F1: 311}
    switch "\n\t" {
    case "":
        v13 := GenS2{F0: "", F1: 0, F2: 1}
        _ = v13
    default:
        for fi14 := 0; fi14 < 15; fi14++ {
            if (0.1 < (float32(-2.25) * ((0.1 * -2.25) * 0.1))) {
                v15 := GenS2{F0: ("hello" + "foo bar"), F1: 4294967295, F2: 0}
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (true && (true || true)) {
                        var v17 uint8 = 23
                        fmt.Println(fi14)
                        p0 = v11
                        _ = v17
                    } else {
                        v11 = (v11 | (v11 | v11))
                        if (2 < 255) {
                            switch (fi14 << ((0) & 63)) {
                            case 128:
                                var v18 float64 = 1.5
                                _ = v18
                            default:
                                v19 := []int{(fi16 * fi14)}
                                _ = v19
                            }
                            if false {
                                v20 := []int{(fi16 << ((2) & 63)), fi16}
                                if true {
                                    fmt.Println(fi16, v11)
                                    p0 = p0
                                    var v21 bool = ((!false) && true)
                                    _ = v21
                                }
                                _ = v20
                            } else {
                                var v22 int32 = 0
                                _ = v22
                            }
                        }
                    }
                    _ = fi16
                }
                _ = v15
            }
            _ = fi14
        }
    }
    p0 = (p0 - (p0 / ((42) | 1)))
    v11 = 0
    _ = p0
    _ = v11
    _ = v12
    return 195
}

func main() {
    var v1 bool = false
    var v2 uint32 = 65535
    var v3 uint16 = 0
    v4 := []int{}
    var v5 uint8 = 128
    var v6 int = -685
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v1 {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 uint = 1
            _ = fi23
            _ = v24
        }
    } else {
        if len(v4) > 0 {
            v4[0] = (^(v6))
        }
        v25 := GenS3{F0: v2, F1: ""}
        v5 = (^(v5))
        _ = v25
    }
    v26 := GenS4{F0: v2}
    if true {
        fmt.Println(v1)
        _ = len(v4)
        var v27 uint32 = v2
        _ = v27
    }
    _ = len(v4)
    var v28 int8 = 61
    switch -895 {
    case 32767:
        if (-221 > 190) {
            if ((v28 + v28) <= v28) {
                fmt.Println(v2, v3, v3)
            } else {
                var v29 uint = 18446744073709551615
                for fi30 := 0; fi30 < 1; fi30++ {
                    v1 = ((v6 + (fi30 ^ -576)) == -9223372036854775808)
                    v2 = v2
                    _ = fi30
                }
                v1 = v1
                _ = v29
            }
            v1 = false
            switch (v6 / ((32768) | 1)) {
            case 82:
                if true {
                    v3 -= v3
                } else {
                    var v31 uint16 = (v3 - (^(v3)))
                    _ = v31
                }
            default:
                v2 = v2
            }
        }
    case 9223372036854775807:
        var v32 uint32 = v2
        _ = v32
    case 32768:
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 uint64 = 2
            v4 = append(v4, v6)
            v4 = append(v4, (v6 % ((v6) | 1)))
            _ = fi33
            _ = v34
        }
    default:
        v35 := GenS3{F0: v2, F1: ""}
        _ = v35
    }
    if len(v4) > 0 {
        v4[0] = (^(v6))
    }
    v4 = append(v4, (v6 - (v6 & 0)))
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 12; fi37++ {
            fmt.Println(v3)
            var v38 uint64 = 2
            if (int32(v6) == 42) {
                _ = len(v4)
            }
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 uint16 = v3
        var v41 uint16 = v3
        _ = fi39
        _ = v40
        _ = v41
    }
    v6 = -1
    var v42 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
