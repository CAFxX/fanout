package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint32; F2 int16 }
type GenS5 struct { F0 bool; F1 bool; F2 int8 }

func main() {
    var v1 uint = 42
    var v2 uint = v1
    v3 := []uint32{2147483648, 54}
    v4 := make(map[string]bool)
    var v5 uint16 = 2
    v1 -= 0
    fmt.Println(v2)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 597
        if true {
            v5 -= v5
            var v8 string = (("\n\t" + "\n\t") + "")
            v8 = (v8 + v8)
            _ = v8
        } else {
            v7 = 963
            switch 128 {
            case -2147483648:
                var v9 int16 = 0
                _ = v9
            case -695:
                var v10 int32 = -710
                _ = v10
            }
            if false {
                if (42 <= 2) {
                    v11 := GenS1{F0: 42}
                    _ = v11
                }
                switch "hello" {
                case "":
                    _ = len(v4)
                default:
                    v2 -= 0
                }
                fmt.Println(v7)
            } else {
                _ = len(v4)
                v12 := GenS2{F0: fi6, F1: 581, F2: uint64(v5)}
                _ = v12
            }
        }
        var v13 int8 = -59
        _ = fi6
        _ = v13
        _ = v7
    }
    v5 = v5
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v2, v1)
        v5 += uint16(v2)
        var v15 uint16 = v5
        _ = fi14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int8 = 36
            v1 -= (v1 - (v2 | v1))
            if true {
                switch (^(fi16)) {
                case 4294967295:
                    v18 := GenS1{F0: 778}
                    _ = v18
                case 9223372036854775807:
                    v2 += v1
                default:
                    if true {
                        var v19 string = ("" + "hello")
                        switch fi16 {
                        case -416:
                            for fi20 := 0; fi20 < 1; fi20++ {
                                var v21 uint = 255
                                v22 := GenS2{F0: (fi20 & (fi16 * fi20)), F1: 951, F2: 1}
                                v23 := GenS3{F0: 0, F1: v5}
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                        case 0:
                            v19 = ((v19 + "hello") + v19)
                        case -87:
                            for fi24 := 0; fi24 < 8; fi24++ {
                                v25 := GenS4{F0: 42, F1: 0, F2: -1}
                                _ = fi24
                                _ = v25
                            }
                        default:
                            v17 -= (v17 + 100)
                        }
                        _ = v19
                    }
                }
                var v26 float32 = 0.1
                _ = v26
            } else {
                for fi27 := 0; fi27 < 19; fi27++ {
                    v28 := GenS5{F0: false, F1: (1 != -796), F2: -128}
                    _ = fi27
                    _ = v28
                }
                var v29 int8 = (v17 * v17)
                _ = v29
            }
            _ = fi16
            _ = v17
        }
    } else {
        var v30 int64 = 0
        v5 = (v5 & uint16(v1))
        _ = v30
    }
    var v31 uint64 = 269
    var v32 uint16 = 32767
    switch 42 {
    case -9223372036854775808:
        v33 := GenS5{F0: true, F1: true, F2: 1}
        _ = v33
    case 32768:
        if (-1 == 42) {
            var v34 uint64 = (v31 | (v31 % ((1) | 1)))
            var v35 uint32 = uint32(v1)
            _ = v34
            _ = v35
        } else {
            switch 0 {
            case -2147483648:
                v5 = 0
            case 4294967295:
                var v36 float32 = float32(0.1)
                _ = v36
            case -9223372036854775808:
                v37 := v4[("hello" + (("" + "") + "hello"))]
                _ = v37
            }
        }
    }
    v5 = (v32 / (((v32 << (((v2 + v2)) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
