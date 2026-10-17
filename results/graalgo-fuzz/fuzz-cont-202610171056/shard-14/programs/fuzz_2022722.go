package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint8; F2 string }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 string; F1 int64 }
func genFn0(p0 uint64) uint8 {
    p0 = p0
    var v7 int16 = 1
    var v8 int16 = (v7 + v7)
    p0 -= (p0 ^ p0)
    _ = p0
    _ = v7
    _ = v8
    return 13
}
func genFn1(p0 int16, p1 uint64, p2 bool) int {
    var v9 int = 32768
    var v10 int8 = -1
    var v11 uint32 = 389
    if true {
        fmt.Println(v9, v9, v11)
        v12 := GenS1{F0: 336}
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                p2 = p2
                if (((3.14159 * ((((float32(1e10) - 3.14159) - float32((1e10 - float64(-2.25)))) * 1e10) * 3.14159)) / 1.0) < 0.0) {
                    var v15 int8 = 100
                    if p2 {
                        if false {
                            switch (fi14 % ((511) | 1)) {
                            case -2147483648:
                                v16 := []int{fi13, (^(fi14)), (^(fi14)), (v9 & v9)}
                                _ = v16
                            case 9223372036854775807:
                                var v17 int32 = 1
                                _ = v17
                            }
                            var v18 uint32 = (v11 >> ((1) & 63))
                            fmt.Println(p0)
                            _ = v18
                        }
                        v19 := GenS2{F0: 65535, F1: 9223372036854775807, F2: ((v11 | (v11 + v11)) < (v11 >> ((4294967295) & 63)))}
                        v20 := []int{594, v9, int(p1), -9223372036854775808}
                        _ = v19
                        _ = v20
                    }
                    _ = v15
                } else {
                    p0 += p0
                    if (fi14 <= 9223372036854775807) {
                        if ((v11 < (v11 - 857)) || p2) {
                            var v21 uint8 = 2
                            _ = v21
                        }
                        v9 = 32768
                    } else {
                        v22 := []int{2147483647}
                        _ = v22
                    }
                    fmt.Println(fi13)
                    var v23 uint16 = 886
                    _ = v23
                }
                if ((float64(0.1) + 1.5) < 1e10) {
                    if true {
                        v24 := GenS2{F0: 669, F1: 42, F2: (p2 || true)}
                        _ = v24
                    } else {
                        var v25 uint = 83
                        _ = v25
                    }
                    v26 := GenS3{F0: 4294967295}
                    _ = v26
                } else {
                    p0 = -32768
                    if false {
                        var v27 int8 = v10
                        v28 := GenS4{F0: 139, F1: uint8(p1), F2: ("a" + ("" + ("hello" + "hello")))}
                        var v29 float32 = ((float32(-2.25) + 0.0) - (((((3.14159 - (1e10 + ((0.1 * 1e10) * (1e10 - float32(1e10))))) + (0.1 / 0.1)) - ((0.0 * 0.1) + 0.0)) * (1e10 * float32(1e10))) - float32(0.1)))
                        var v30 float32 = v29
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS5{F0: -1, F1: p0}
                        _ = v31
                    }
                }
                _ = fi14
            }
            var v32 int8 = (v10 % ((v10) | 1))
            _ = fi13
            _ = v32
        }
        _ = v12
    } else {
        v33 := GenS6{F0: 144, F1: p0}
        v34 := []int{(v9 / ((v9) | 1))}
        _ = v33
        _ = v34
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return v9
}

func main() {
    v1 := []uint32{2147483647, 4294967295, 65535, 2}
    v2 := []string{(("a" + "a") + ("foo bar" + ((("\n\t" + "foo bar") + "a") + ("foo bar" + "foo bar")))), "", "foo bar"}
    var v3 int8 = -1
    v4 := make(map[string]string)
    var v5 int64 = 42
    v6 := make(map[string]string)
    switch "a" {
    case "":
        if true {
            v5 = (-(v5))
        } else {
            fmt.Println(v5)
            v3 = v3
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 uint8 = 128
                for fi37 := 0; fi37 < 17; fi37++ {
                    _ = len(v2)
                    var v38 bool = true
                    var v39 bool = (1 != 0)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi35
                _ = v36
            }
        }
    }
    v3 = (v3 / ((v3) | 1))
    for fi40 := 0; fi40 < 15; fi40++ {
        v5 -= v5
        if (!false) {
            fmt.Println(v5)
            fmt.Println(v5)
        }
        _ = fi40
    }
    var v41 int16 = 1
    var v42 uint = 335
    v42 -= (v42 | (v42 ^ 65535))
    var v43 int16 = int16(v3)
    v44 := GenS7{F0: "a", F1: (v5 << ((v42) & 63))}
    v42 -= (v42 * (^(v42)))
    switch "foo bar" {
    case "":
        v45 := GenS3{F0: v42}
        _ = v45
    }
    if false {
        var v46 bool = true
        v47 := GenS4{F0: int64(v3), F1: 44, F2: "\n\t"}
        _ = v46
        _ = v47
    } else {
        v48 := GenS7{F0: "a", F1: 0}
        _ = v48
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        var v50 int64 = int64(v3)
        _ = fi49
        _ = v50
    }
    var v51 float64 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v52 int64 = (-(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
