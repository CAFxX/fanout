package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int64; F2 int }
func genFn0() int16 {
    switch 42 {
    case -996:
        v5 := GenS1{F0: 346}
        _ = v5
    case 0:
        var v6 uint16 = 255
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 int = fi7
                v10 := []int{128}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    default:
        if (!(!true)) {
            var v11 uint8 = 127
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 4294967295}
                v14 := GenS2{F0: 255, F1: 0}
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(v11)
            _ = v11
        }
    }
    var v15 uint64 = 18446744073709551615
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint8 = 200
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 bool = (2 <= (v15 & (v15 % (((v15 ^ (v15 ^ 316))) | 1))))
        v20 := []int{127, 1}
        fmt.Println(v19, v15, fi18)
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v15
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 42
    var v3 int = -9223372036854775808
    v4 := []int{(v3 + (v3 % (((v3 & 128)) | 1))), (-(v3)), (v3 / (((v3 & 42)) | 1))}
    v2 = (v2 + 733)
    v21 := GenS3{F0: -291}
    v2 -= v2
    v3 = (v3 / ((11) | 1))
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint8 = 128
            if (!(3.14159 == 0.0)) {
                v25 := GenS4{F0: -79}
                switch (fi22 | 2147483647) {
                case 128:
                    if false {
                        v24 -= v24
                    }
                default:
                    if false {
                        v24 = 128
                        switch 128 {
                        case 0:
                            switch fi23 {
                            case 410:
                                var v26 uint = 2
                                _ = v26
                            case 42:
                                v1["\n\t"] = (^(v2))
                                if (1 != 1) {
                                    v27 := GenS5{F0: "a", F1: -1, F2: (fi23 * fi23)}
                                    v24 = v24
                                    var v28 float64 = -2.25
                                    _ = v27
                                    _ = v28
                                }
                            case 4294967295:
                                var v29 uint16 = 0
                                _ = v29
                            default:
                                v24 = (v24 >> ((793) & 63))
                            }
                        case 128:
                            v2 -= v2
                        default:
                            if len(v4) > 0 {
                                v4[0] = int(v24)
                            }
                        }
                    } else {
                        _ = len(v1)
                        v24 = v24
                    }
                }
                _ = v25
            } else {
                fmt.Println(fi23)
                v3 = (^(v3))
                if (!(1.5 == (float32(0.1) / ((((float32(-2.25) + 0.0) + (0.1 - 1e10)) + (0.0 + float32(1e10))) + (((-2.25 + 0.1) * 0.0) + 0.1))))) {
                    var v30 int = fi23
                    _ = v30
                }
            }
            _ = fi23
            _ = v24
        }
        v3 = fi22
        v3 += fi22
        _ = fi22
    }
    v31 := GenS4{F0: 0}
    fmt.Println(v3, v3)
    if (!(("hello" != "") && false)) {
        switch 970 {
        case 32768:
            v32 := GenS1{F0: 42}
            _ = v32
        case 2147483647:
            switch "hello" {
            case "a":
                v33 := GenS1{F0: uint64(v2)}
                _ = v33
            case "\n\t":
                fmt.Println(v2, v2)
            default:
                for fi34 := 0; fi34 < 2; fi34++ {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        v36 := GenS4{F0: -1}
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
            }
        case 42:
            var v37 uint8 = 61
            _ = v37
        default:
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 uint32 = v2
                _ = fi38
                _ = v39
            }
        }
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            v3 -= (v3 + (^(v3)))
            if true {
                for fi41 := 0; fi41 < 16; fi41++ {
                    var v42 int16 = -782
                    _ = fi41
                    _ = v42
                }
                var v43 uint8 = 2
                _ = v43
            } else {
                var v44 bool = (1e10 < 0.1)
                var v45 uint = 65535
                v3 = v3
                _ = v44
                _ = v45
            }
            _ = fi40
        }
        v46 := v1["\n\t"]
        var v47 float32 = float32(-2.25)
        _ = v46
        _ = v47
    }
    v3 -= (v3 * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
