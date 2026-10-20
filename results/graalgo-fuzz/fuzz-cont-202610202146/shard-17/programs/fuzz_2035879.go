package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int64) int16 {
    var v7 uint = 65535
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int8 = -128
        p0 = (p0 + p0)
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 32767:
                var v11 string = "foo bar"
                _ = v11
            default:
                switch (fi8 - -9223372036854775808) {
                case 931:
                    v9 -= 1
                case 0:
                    if (v9 <= v9) {
                        v9 -= v9
                    }
                case 800:
                    if true {
                        v12 := GenS1{F0: 2}
                        fmt.Println(v9)
                        _ = v12
                    }
                }
            }
            v9 = (^(v9))
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    if (125 >= 127) {
        v7 = (v7 >> (((v7 * 255)) & 63))
        v13 := []int{695, -955, -358, 1}
        _ = v13
    } else {
        var v14 float64 = ((0.1 + 1e10) + (-2.25 * (1e10 / 1.0)))
        if (3.14159 <= 1e10) {
            v15 := GenS2{F0: 42, F1: 255}
            _ = v15
        }
        p0 += p0
        _ = v14
    }
    _ = p0
    _ = v7
    return 614
}

func main() {
    v1 := []uint32{153, 42, 158, 533}
    v2 := []uint32{42}
    var v3 int32 = -922
    var v4 int16 = -541
    v5 := make(map[string]bool)
    v6 := make(map[string]uint32)
    switch 0 {
    case 1:
        v16 := GenS3{F0: 42}
        _ = v16
    default:
        v17 := v5["\n\t"]
        _ = v17
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v4 += genFn0(-1)
        if (29 == 502) {
            v18 := GenS3{F0: 1}
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            if true {
                var v20 uint32 = 1
                _ = v20
            } else {
                v21 := GenS3{F0: 1}
                _ = v21
            }
            _ = fi19
        }
    }
    v22 := GenS2{F0: 18446744073709551615, F1: 2}
    var v23 uint = 794
    v24 := GenS2{F0: 888, F1: 25}
    v4 += -389
    var v25 int64 = 9223372036854775807
    v26 := GenS4{F0: 127, F1: 722}
    v3 -= (v3 ^ (^(v3)))
    var v27 uint = v23
    switch 105 {
    case -576:
        v28 := GenS1{F0: 42}
        _ = v28
    case 128:
        switch "foo bar" {
        case "foo bar":
            var v29 float64 = (0.1 / 1e10)
            _ = v29
        case "\n\t":
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 bool = true
                switch fi30 {
                case 9223372036854775807:
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v33 := GenS5{F0: 4294967295, F1: (^(fi30)), F2: "hello"}
                        _ = fi32
                        _ = v33
                    }
                case 687:
                    v34 := GenS2{F0: uint64(v27), F1: 2}
                    _ = v34
                case 32768:
                    v35 := GenS6{F0: 42}
                    _ = v35
                default:
                    if (!((v4 % ((-32768) | 1)) != genFn0((v25 << (((v27 % ((2) | 1))) & 63))))) {
                        if (42 < 2) {
                            v31 = (3.14159 < 3.14159)
                            v36 := GenS5{F0: 1, F1: fi30, F2: ((("hello" + "a") + "a") + ("hello" + "foo bar"))}
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v1 = append(v1, 4294967295)
                                _ = fi37
                            }
                            _ = v36
                        }
                        var v38 int32 = 0
                        var v39 string = "hello"
                        _ = v38
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 7; fi40++ {
                            if v31 {
                                v27 = v23
                            } else {
                                v4 = int16(fi30)
                            }
                            var v41 int = fi30
                            var v42 uint16 = 0
                            v42 += (v42 - v42)
                            fmt.Println(v25, v27)
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    }
                }
                _ = len(v6)
                _ = fi30
                _ = v31
            }
        case "":
            var v43 uint16 = 255
            _ = v43
        default:
            v27 += (v27 % (((v23 << ((v23) & 63))) | 1))
        }
    default:
        _ = len(v5)
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 uint8 = 127
        _ = fi44
        _ = v45
    }
    if ((^(v27)) != (v27 & v27)) {
        var v46 int16 = -32768
        var v47 uint64 = 421
        _ = v46
        _ = v47
    }
    v23 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
