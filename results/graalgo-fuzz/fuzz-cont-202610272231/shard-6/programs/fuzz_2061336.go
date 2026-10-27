package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8; F2 int }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint8) uint8 {
    var v4 int32 = -283
    if false {
        var v5 string = ""
        var v6 uint8 = p0
        _ = v5
        _ = v6
    } else {
        var v7 int64 = 376
        v4 = (^(v4))
        var v8 uint32 = 1
        _ = v7
        _ = v8
    }
    if (!false) {
        if (true || false) {
            var v9 uint32 = 1
            v4 += v4
            _ = v9
        }
    }
    var v10 uint8 = 255
    p0 -= 255
    _ = p0
    _ = v10
    _ = v4
    return p0
}

func main() {
    var v1 uint8 = 39
    var v2 uint16 = 985
    var v3 float64 = (-2.25 * 3.14159)
    v3 = ((v3 - v3) + 3.14159)
    v2 += v2
    var v11 int = 128
    v12 := []int{869, v11, -9223372036854775808}
    var v13 int8 = 1
    v14 := GenS1{F0: ("a" + "hello"), F1: uint32(v2), F2: 32767}
    if ((-2.25 + ((((1.5 + (((v3 + (float64(0.0) - v3)) / 1.0) + 1.5)) - (((float64(1.5) / 1.5) + v3) - 1.5)) - v3) * v3)) >= float64(0.1)) {
        if false {
            var v15 int16 = -24
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            v13 = -15
            v2 = (v2 | 255)
            v17 := GenS2{F0: false}
            _ = fi16
            _ = v17
        }
        if true {
            v18 := GenS1{F0: "hello", F1: 2147483647, F2: 32767}
            _ = v18
        } else {
            v19 := GenS1{F0: (("a" + "a") + "a"), F1: 2147483648, F2: int16(v2)}
            v20 := GenS2{F0: ((0 != -9223372036854775808) || true)}
            _ = v19
            _ = v20
        }
    } else {
        v3 = float64(1.5)
    }
    v21 := GenS3{F0: 2, F1: 42, F2: v11}
    if ("hello" < ("foo bar" + ("\n\t" + ""))) {
        var v22 float32 = 0.1
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint8 = (v1 << ((4294967295) & 63))
            _ = fi23
            _ = v24
        }
        var v25 uint8 = (v1 / (((v1 + v1)) | 1))
        _ = v22
        _ = v25
    }
    v26 := GenS3{F0: 18446744073709551615, F1: -107, F2: v11}
    if false {
        if ((v2 % ((0) | 1)) != v2) {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 int8 = (^(v13))
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 19; fi29++ {
                switch -490 {
                case 42:
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v11 += (fi29 >> ((42) & 63))
                        var v31 int64 = 9223372036854775807
                        _ = fi30
                        _ = v31
                    }
                case 0:
                    if ((1 < v1) || false) {
                        v11 = fi29
                        v32 := GenS4{F0: 478, F1: 190}
                        _ = v32
                    } else {
                        switch "a" {
                        case "":
                            var v33 int16 = 32767
                            _ = v33
                        case "a":
                            var v34 int = fi29
                            _ = v34
                        case "hello":
                            v3 = v3
                        default:
                            var v35 float32 = 0.1
                            _ = v35
                        }
                        var v36 float64 = (0.0 - (((v3 - (v3 + float64((1e10 + 1.5)))) + 1.5) / v3))
                        _ = v36
                    }
                case -306:
                    var v37 float32 = 0.1
                    _ = v37
                }
                for fi38 := 0; fi38 < 19; fi38++ {
                    if true {
                        switch "" {
                        case "\n\t":
                            v39 := GenS4{F0: 42, F1: uint16(v1)}
                            _ = v39
                        case "a":
                            v11 -= fi29
                        case "foo bar":
                            switch 145 {
                            case -568:
                                var v40 int32 = 2147483647
                                _ = v40
                            case -31:
                                var v41 uint16 = (v2 - (v2 << ((18446744073709551615) & 63)))
                                _ = v41
                            }
                        default:
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                    } else {
                        v11 = (fi29 / ((127) | 1))
                        v13 = v13
                        switch (fi38 % ((v11) | 1)) {
                        case 914:
                            var v42 uint64 = 18446744073709551615
                            _ = v42
                        case 2147483647:
                            v11 -= fi38
                        case 0:
                            if false {
                                v13 -= 0
                                v2 = (^(v2))
                                var v43 uint32 = 42
                                _ = v43
                            } else {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                if (!(!true)) {
                                    v13 = v13
                                }
                                var v44 int32 = 761
                                v2 = v2
                                _ = v44
                            }
                        default:
                            v2 = v2
                        }
                        v45 := GenS3{F0: 627, F1: v13, F2: fi38}
                        _ = v45
                    }
                    switch -1 {
                    case -406:
                        v13 += -1
                    }
                    v11 = (fi38 & (fi29 + -50))
                    _ = fi38
                }
                _ = fi29
            }
        }
    } else {
        v13 -= (v13 ^ (v13 ^ (v13 + v13)))
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        var v47 int = (v11 | int(v2))
        switch fi46 {
        case 4294967295:
            v11 -= 9223372036854775807
        }
        v3 = (1.5 / v3)
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 float64 = (v3 * v3)
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v26
    _ = v3
}
