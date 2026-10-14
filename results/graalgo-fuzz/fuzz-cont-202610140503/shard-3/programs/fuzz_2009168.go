package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int16; F1 bool; F2 uint }
type GenS7 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 int32) bool {
    var v5 int = -1
    if true {
        var v6 int32 = p0
        if (!((v5 / ((42) | 1)) < v5)) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 bool = ((127 <= -48) || true)
                fmt.Println(v6, p0)
                var v9 uint32 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(p0, v6, v6)
                _ = fi10
            }
            fmt.Println(p0, p0)
        } else {
            var v11 uint8 = 166
            if ((uint32(v11) > 676) || (!true)) {
                if (!false) {
                    fmt.Println(v11)
                    var v12 float64 = ((1e10 / -2.25) / 1e10)
                    _ = v12
                }
                fmt.Println(v11, v5, v11)
            } else {
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint32 = 0
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS1{F0: 0, F1: (-(v6))}
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                if true {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS2{F0: 2}
                        var v19 bool = false
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    switch "a" {
                    case "hello":
                        v11 = v11
                    case "foo bar":
                        v20 := GenS1{F0: 127, F1: v6}
                        _ = v20
                    case "a":
                        var v21 int16 = -369
                        _ = v21
                    }
                    v6 = (-(p0))
                }
                var v22 int16 = 42
                _ = v22
            }
            p0 += p0
            _ = v11
        }
        switch (v5 + (v5 + v5)) {
        case 633:
            v5 += v5
        case -1:
            v23 := GenS1{F0: 127, F1: (v6 / ((p0) | 1))}
            _ = v23
        case -9223372036854775808:
            if (101 >= 0) {
                if false {
                    if ((42 <= 47) || true) {
                        v24 := []int{127}
                        if true {
                            fmt.Println(v6)
                            var v25 string = "hello"
                            v26 := GenS3{F0: 0, F1: 0, F2: 1}
                            _ = v25
                            _ = v26
                        } else {
                            var v27 int8 = -128
                            var v28 string = (((((("\n\t" + "\n\t") + "hello") + "") + "") + "foo bar") + "")
                            _ = v27
                            _ = v28
                        }
                        _ = v24
                    } else {
                        var v29 int8 = 42
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 uint32 = 236
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                }
            }
        default:
            if (139 >= 1) {
                v5 = v5
                v5 -= (v5 % ((v5) | 1))
            } else {
                v32 := []int{-418, (v5 << ((18446744073709551615) & 63))}
                _ = v32
            }
        }
        _ = v6
    }
    _ = p0
    _ = v5
    return (!((3.14159 * (float64((((3.14159 - 1e10) / (0.1 / (3.14159 / 1.0))) + (0.1 / 1.0))) * 0.1)) <= 0.1))
}
func genFn1(p0 uint64, p1 uint, p2 string) bool {
    var v33 uint16 = 65535
    if (-232 >= -1) {
        fmt.Println(p2, p2)
    }
    if false {
        p2 = p2
        for fi34 := 0; fi34 < 12; fi34++ {
            switch fi34 {
            case -2147483648:
                if true {
                    v35 := GenS4{F0: -677, F1: "a", F2: 168}
                    fmt.Println(v33, v33)
                    var v36 int64 = 0
                    _ = v35
                    _ = v36
                } else {
                    v37 := []int{fi34, 42, (fi34 ^ fi34), (fi34 * (fi34 + fi34))}
                    p2 = "\n\t"
                    _ = v37
                }
            }
            v38 := GenS5{F0: (int16(p1) >= 1), F1: p0, F2: 0}
            var v39 float64 = 0.0
            _ = fi34
            _ = v38
            _ = v39
        }
    } else {
        fmt.Println(p1)
        if (false && false) {
            fmt.Println(v33, p1, p0)
            v40 := GenS6{F0: 0, F1: false, F2: p1}
            if (1 <= 820) {
                var v41 int8 = -1
                _ = v41
            } else {
                var v42 float32 = 1e10
                var v43 uint8 = 255
                p1 = (p1 << ((p1) & 63))
                _ = v42
                _ = v43
            }
            _ = v40
        }
        switch 127 {
        case 42:
            p0 = (p0 | p0)
        }
    }
    var v44 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v44
    return (!false)
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{0}
    var v3 float64 = 1.5
    var v4 uint8 = 42
    v45 := GenS7{F0: 2147483648, F1: 2147483647, F2: v1}
    var v46 float32 = ((((float32(((v3 / (v3 / v3)) / -2.25)) + 0.0) * 0.1) * (3.14159 * 0.0)) - 0.1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v47 float64 = v3
    v1 = (v1 | v1)
    v1 = v1
    if (("hello" + "") <= "a") {
        v48 := GenS8{F0: 2, F1: 42, F2: 1}
        _ = v48
    }
    var v49 uint32 = 2
    var v50 uint8 = (^(v4))
    fmt.Println(v47, v3, v49)
    var v51 int64 = 0
    var v52 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
