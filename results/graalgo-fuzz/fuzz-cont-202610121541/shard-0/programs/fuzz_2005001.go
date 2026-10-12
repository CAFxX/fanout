package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
func genFn0() uint8 {
    var v4 uint16 = 32767
    if false {
        v4 -= (v4 & v4)
        v5 := []int{9223372036854775807, 32768}
        var v6 int16 = 42
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 127, F1: 314, F2: 2147483647}
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 bool = true
    v3 := make(map[uint32]int)
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{fi9, 32767}
            switch ("a" + (("foo bar" + "hello") + "hello")) {
            case "a":
                switch 4294967295 {
                case 32767:
                    v11 := GenS1{F0: 255, F1: 1, F2: (fi9 >> ((0) & 63))}
                    _ = v11
                case 32768:
                    v2 = (0.1 == 0.0)
                case -341:
                    var v12 float64 = ((0.1 / (-2.25 / -2.25)) - ((1.5 - 0.1) / 3.14159))
                    _ = v12
                }
            default:
                v2 = (v2 && false)
            }
            v13 := GenS2{F0: 0}
            _ = fi9
            _ = v10
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            v1 += (v1 >> ((65535) & 63))
            var v15 string = "a"
            v2 = (2147483647 > 42)
            _ = fi14
            _ = v15
        }
        _ = fi8
    }
    v16 := []int{-925, 9223372036854775807, -1}
    fmt.Println(v2, v1)
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS3{F0: 421}
            _ = fi18
            _ = v19
        }
        var v20 int16 = -803
        _ = fi17
        _ = v20
    }
    v21 := GenS1{F0: 128, F1: 966, F2: -1}
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                v24 := GenS2{F0: 18446744073709551615}
                _ = v24
            } else {
                if true {
                    v25 := GenS1{F0: 234, F1: 88, F2: (fi22 << ((1) & 63))}
                    v3[383] = fi22
                    if (!(("hello" == "a") || true)) {
                        v16 = append(v16, fi23)
                        fmt.Println(v2)
                    }
                    _ = v25
                }
                v1 += v1
                var v26 int32 = -1
                fmt.Println(v26, v26)
                if false {
                    switch "a" {
                    case "foo bar":
                        v26 += (v26 * -2147483648)
                        _ = len(v16)
                        switch (^(fi23)) {
                        case 32768:
                            var v27 int16 = -24
                            _ = v27
                        case 268:
                            v2 = v2
                        default:
                            switch int(v1) {
                            case 620:
                                _ = len(v3)
                            }
                        }
                        v1 -= (v1 & -1)
                    case "":
                        var v28 bool = v2
                        _ = v28
                    default:
                        v29 := GenS2{F0: 4294967295}
                        _ = v29
                    }
                    var v30 int64 = -338
                    _ = v30
                } else {
                    v31 := GenS1{F0: 2, F1: -61, F2: (-(fi22))}
                    switch fi22 {
                    case 9223372036854775807:
                        var v32 uint8 = genFn0()
                        v33 := GenS3{F0: 42}
                        v32 -= v32
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(fi22)
                    _ = v31
                }
                _ = v26
            }
            var v34 int32 = 42
            v35 := GenS3{F0: 2147483647}
            _ = fi23
            _ = v34
            _ = v35
        }
        _ = fi22
    }
    fmt.Println(v1, v1)
    var v36 uint8 = 127
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 uint64 = 1
        v38 = 42
        _ = fi37
        _ = v38
    }
    if true {
        if ((((genFn0() <= 0) || (((0 < uint32(v36)) && v2) || (1.5 == 0.1))) && v2) || false) {
            for fi39 := 0; fi39 < 7; fi39++ {
                for fi40 := 0; fi40 < 20; fi40++ {
                    var v41 int32 = 2147483647
                    v42 := GenS1{F0: (v36 & 137), F1: -52, F2: (fi40 - fi40)}
                    if ((fi39 - (-(fi40))) != 2147483647) {
                        v41 -= v41
                        fmt.Println(fi40, fi39)
                        var v43 bool = false
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi39
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        v1 += 127
    } else {
        var v44 uint16 = 169
        _ = v44
    }
    var v45 int32 = -2147483648
    v46 := GenS4{F0: 42, F1: (-661 != 811)}
    var v47 int = 127
    v45 -= (v45 / ((42) | 1))
    for fi48 := 0; fi48 < 6; fi48++ {
        for fi49 := 0; fi49 < 13; fi49++ {
            v50 := GenS5{F0: (v47 << ((4294967295) & 63)), F1: v45, F2: (v45 << ((65535) & 63))}
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    fmt.Println(v2)
    var v51 uint = 0
    switch "foo bar" {
    case "":
        var v52 uint16 = 365
        _ = v52
    case "a":
        v53 := GenS5{F0: v47, F1: v45, F2: v45}
        _ = v53
    default:
        delete(v3, 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v45
    _ = v46
    _ = v47
    _ = v51
}
