package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 uint8 = 42
    var v2 float64 = (((0.1 + 1e10) * 0.1) - 3.14159)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 19; fi4++ {
        v2 -= ((((float64((1e10 / 1e10)) + v2) * float64(-2.25)) * -2.25) * v2)
        _ = fi4
    }
    v3[2] = "a"
    if (false && (true && false)) {
        v1 = 189
        if true {
            switch "foo bar" {
            case "":
                v2 = v2
            }
        } else {
            for fi5 := 0; fi5 < 7; fi5++ {
                var v6 int = (fi5 % (((fi5 % ((int(v1)) | 1))) | 1))
                if false {
                    delete(v3, 4294967295)
                }
                var v7 uint = 42
                _ = fi5
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := GenS1{F0: -389}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            if (true || true) {
                var v11 int32 = -2147483648
                var v12 int8 = 0
                switch "" {
                case "hello":
                    v13 := GenS2{F0: 0, F1: int32(v1), F2: 0}
                    _ = v13
                }
                _ = v11
                _ = v12
            } else {
                var v14 int8 = 100
                if (2 > 2) {
                    _ = len(v3)
                    var v15 float32 = 1.5
                    _ = v15
                } else {
                    v16 := GenS1{F0: 23}
                    fmt.Println(v1, v1)
                    var v17 uint32 = 42
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
        }
        switch -634 {
        case 127:
            v18 := GenS1{F0: 721}
            _ = v18
        default:
            v19 := GenS3{F0: true, F1: -1}
            _ = v19
        }
    } else {
        v2 = (v2 - (v2 - v2))
        fmt.Println(v2, v1)
        if false {
            switch -394 {
            case 32767:
                if (("" + "foo bar") <= ("foo bar" + ("\n\t" + ("\n\t" + "")))) {
                    v20 := v3[65535]
                    for fi21 := 0; fi21 < 6; fi21++ {
                        v20 = "a"
                        v22 := []int{(-(fi21)), (fi21 ^ fi21)}
                        _ = fi21
                        _ = v22
                    }
                    if true {
                        var v23 uint16 = 484
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 9; fi24++ {
                            var v25 uint8 = 42
                            v26 := GenS4{F0: 2147483647}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v20
                }
            case 9223372036854775807:
                fmt.Println(v2)
            case 147:
                v27 := []int{127, 0, -137}
                _ = v27
            }
        }
    }
    switch 18 {
    case -656:
        var v28 float64 = 0.1
        _ = v28
    case -2147483648:
        v29 := GenS5{F0: 42}
        _ = v29
    default:
        v30 := v3[455]
        _ = v30
    }
    v2 += 0.1
    fmt.Println(v2)
    v31 := GenS2{F0: -22, F1: -939, F2: 9223372036854775807}
    var v32 int = 42
    v32 = (v32 + v32)
    var v33 int16 = int16(v32)
    v3[1] = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
