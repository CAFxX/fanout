package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint16; F1 string }

func main() {
    var v1 uint32 = 434
    v2 := []uint32{v1}
    v3 := make(map[string]uint32)
    if false {
        v1 = 65535
        for fi4 := 0; fi4 < 4; fi4++ {
            for fi5 := 0; fi5 < 19; fi5++ {
                var v6 uint64 = 42
                v6 = v6
                _ = fi5
                _ = v6
            }
            v1 += 2
            _ = fi4
        }
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        switch fi7 {
        case -2147483648:
            v1 = 376
        }
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint8 = 42
                var v10 float64 = (float64(float32(float64(-2.25))) - 3.14159)
                _ = fi8
                _ = v10
                _ = v9
            }
            v1 -= (v1 * (v1 ^ (^(v1))))
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                delete(v3, "\n\t")
                switch fi7 {
                case 9223372036854775807:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v1, v1)
                    case "hello":
                        var v12 uint64 = 0
                        _ = v12
                    case "":
                        v1 += (v1 / (((v1 % (((v1 ^ (v1 | (v1 + 0)))) | 1))) | 1))
                    }
                case 32767:
                    fmt.Println(v1, fi11, v1)
                case -2147483648:
                    v1 -= v1
                default:
                    if (!false) {
                        v13 := GenS1{F0: 1, F1: 0}
                        var v14 int16 = 1000
                        _ = v13
                        _ = v14
                    }
                }
                delete(v3, "\n\t")
                _ = fi11
            }
            switch "a" {
            case "hello":
                v1 -= 915
            case "a":
                fmt.Println(v1)
            case "foo bar":
                v15 := GenS2{F0: 1, F1: 9223372036854775807, F2: -464}
                _ = v15
            default:
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint16 = 0
                    _ = fi16
                    _ = v17
                }
            }
            if (!true) {
                var v18 int32 = -420
                v19 := GenS1{F0: v18, F1: 264}
                _ = v18
                _ = v19
            }
        }
        _ = fi7
    }
    v20 := GenS1{F0: 645, F1: 1}
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(v1)
            v23 := GenS3{F0: v1, F1: 2, F2: 255}
            _ = fi22
            _ = v23
        }
        var v24 float64 = (0.1 / 1.5)
        _ = fi21
        _ = v24
    }
    var v25 uint = 65535
    if ("" == "") {
        if (true || (2 > 42)) {
            v26 := GenS2{F0: 0, F1: -1, F2: -267}
            v27 := GenS3{F0: v1, F1: 255, F2: v25}
            if (true || (-2.25 < ((0.1 / 3.14159) + 0.0))) {
                var v28 int64 = -835
                _ = v28
            }
            _ = v26
            _ = v27
        }
    } else {
        if (65535 > v1) {
            v1 = v1
        } else {
            v29 := GenS4{F0: 428, F1: ""}
            var v30 int16 = 1
            v30 -= v30
            _ = v29
            _ = v30
        }
    }
    v31 := GenS2{F0: 208, F1: 0, F2: 32767}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v25 = (^(v25))
        if true {
            v1 += (v1 % ((2147483647) | 1))
        }
        fmt.Println(v25, v1, fi32)
        _ = fi32
    }
    v1 = (^(v1))
    delete(v3, ("" + ("\n\t" + ("hello" + "a"))))
    var v33 int8 = -52
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v31
    _ = v33
}
