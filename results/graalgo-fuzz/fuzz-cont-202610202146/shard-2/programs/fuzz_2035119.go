package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
func genFn0(p0 int64, p1 int32, p2 int) int32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if (933 <= -308) {
                v9 := GenS1{F0: fi8, F1: 42, F2: 0}
                v10 := []int{}
                var v11 string = ("foo bar" + "hello")
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi8
        }
        _ = fi7
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}
func genFn1(p0 uint64, p1 int16, p2 uint32) int8 {
    v13 := GenS2{F0: ("\n\t" + "a"), F1: 2, F2: 255}
    fmt.Println(p1, p1, p2)
    for fi14 := 0; fi14 < 18; fi14++ {
        p2 = 1
        switch 128 {
        case -9223372036854775808:
            fmt.Println(p1, p0)
        case 0:
            if (!false) {
                v15 := GenS3{F0: 0}
                if (!false) {
                    p2 = 933
                } else {
                    if (!(1 > -2147483648)) {
                        v16 := []int{(fi14 - 32767), -238, (fi14 & fi14)}
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if (p2 >= p2) {
                                switch fi17 {
                                case 426:
                                    if (!true) {
                                        var v18 bool = true
                                        var v19 int32 = 235
                                        _ = v18
                                        _ = v19
                                    } else {
                                        p2 = p2
                                        var v20 int64 = -1
                                        v21 := GenS2{F0: ("foo bar" + ("\n\t" + "a")), F1: 1, F2: 2}
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 uint16 = 0
                                    _ = v22
                                case -2147483648:
                                    p2 -= p2
                                }
                            }
                            _ = fi17
                        }
                        _ = v16
                    } else {
                        p0 -= p0
                    }
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 int32 = 2147483647
                        v25 := GenS1{F0: 4294967295, F1: -128, F2: 447}
                        v26 := []int{(fi14 ^ fi23), 32767, (fi14 + fi23), fi14}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 int = 0
                    _ = v27
                }
                if false {
                    v28 := GenS1{F0: fi14, F1: -92, F2: 2}
                    var v29 uint32 = p2
                    _ = v28
                    _ = v29
                }
                _ = v15
            }
        case 209:
            if (2147483647 < -1) {
                v30 := GenS4{F0: uint64(p2), F1: 127, F2: 2}
                _ = v30
            } else {
                p0 -= p0
            }
        }
        fmt.Println(p1)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 127
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = -2.25
    var v3 uint64 = v1
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v6 uint64 = 391
    var v31 uint64 = 4294967295
    var v32 uint32 = 42
    v4[v32] = 4294967295
    switch "" {
    case "\n\t":
        switch "a" {
        case "":
            var v33 uint32 = (v32 + (v32 << ((18446744073709551615) & 63)))
            _ = v33
        case "hello":
            fmt.Println(v2)
        case "a":
            var v34 bool = true
            _ = v34
        default:
            v31 = 548
        }
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        switch "\n\t" {
        case "foo bar":
            v6 += uint64(v32)
        case "":
            var v36 uint64 = (^(v6))
            _ = v36
        default:
            v37 := GenS5{F0: uint8(v32)}
            _ = v37
        }
        _ = fi35
    }
    v38 := GenS6{F0: 121}
    if (3.14159 > v2) {
        v39 := v5[("\n\t" + ("foo bar" + "a"))]
        var v40 uint32 = 2147483647
        _ = v39
        _ = v40
    }
    fmt.Println(v1, v31)
    v41 := GenS3{F0: 193}
    v42 := GenS7{F0: 45, F1: "a", F2: (6 <= 255)}
    var v43 int32 = 388
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
