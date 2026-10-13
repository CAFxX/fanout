package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float32 = 1.5
    v2 := []int{127}
    var v3 uint64 = 428
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    var v6 int16 = -40
    v3 = (v3 ^ v3)
    switch 4294967295 {
    case -972:
        v3 += 833
    case -2147483648:
        delete(v5, 469)
    }
    if (true && ((v1 - (v1 / (v1 + 1e10))) >= (v1 * v1))) {
        for fi7 := 0; fi7 < 20; fi7++ {
            switch fi7 {
            case -436:
                v3 = (^(v3))
            case 9223372036854775807:
                v3 -= (v3 * 42)
            case -879:
                switch "\n\t" {
                case "foo bar":
                    v8 := GenS1{F0: -75, F1: "a"}
                    _ = v8
                case "a":
                    v9 := GenS2{F0: (fi7 < (fi7 ^ fi7)), F1: -1}
                    _ = v9
                }
            default:
                var v10 int32 = 0
                _ = v10
            }
            _ = fi7
        }
        fmt.Println(v1)
        switch 4294967295 {
        case -9223372036854775808:
            if ((42 <= 499) && (1e10 != ((0.0 - 1.5) / 1e10))) {
                v11 := GenS2{F0: (false || true), F1: 1}
                v12 := v5[60]
                var v13 string = ((((v12 + (("a" + "") + v12)) + (v12 + v12)) + "\n\t") + "\n\t")
                _ = v11
                _ = v12
                _ = v13
            } else {
                fmt.Println(v6, v6, v1)
                v1 += (v1 - ((v1 / 1.5) + 3.14159))
                fmt.Println(v6, v1)
            }
        case 261:
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS2{F0: false, F1: 0}
                _ = fi14
                _ = v15
            }
        default:
            switch "\n\t" {
            case "a":
                if false {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        var v17 int64 = -875
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS1{F0: 127, F1: ("\n\t" + "hello")}
                    switch "" {
                    case "\n\t":
                        if (!(uint32(v3) <= 2)) {
                            v1 += float32(0.1)
                            v19 := GenS1{F0: 127, F1: "foo bar"}
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 8; fi20++ {
                                var v21 bool = (false || false)
                                _ = fi20
                                _ = v21
                            }
                            v6 = v6
                        }
                    case "":
                        v22 := GenS3{F0: 9223372036854775807}
                        _ = v22
                    case "foo bar":
                        var v23 uint8 = 43
                        _ = v23
                    }
                    _ = v18
                } else {
                    v1 -= (3.14159 * (v1 * 1e10))
                }
            }
        }
    } else {
        var v24 uint16 = 255
        v24 = 716
        if (true || (-1 == 0)) {
            if true {
                if (953 <= 1) {
                    if (((v1 * 1e10) != (((v1 - ((-2.25 - (v1 + 0.1)) / 3.14159)) * (v1 / (0.0 - 0.1))) / v1)) || (false && true)) {
                        v24 = (v24 | (v24 % ((255) | 1)))
                    }
                }
                switch ("hello" + "a") {
                case "\n\t":
                    var v25 int64 = -683
                    _ = v25
                case "":
                    var v26 int8 = 0
                    _ = v26
                case "hello":
                    var v27 uint64 = (v3 >> ((248) & 63))
                    _ = v27
                default:
                    var v28 uint8 = 128
                    _ = v28
                }
                var v29 int = 127
                _ = v29
            }
            v24 += 255
            var v30 uint = 2
            _ = v30
        }
        _ = v24
    }
    var v31 uint16 = 255
    var v32 uint = 4294967295
    v1 -= (v1 / -2.25)
    v32 = v32
    switch ("hello" + "foo bar") {
    case "hello":
        if (0 <= 233) {
            var v33 int8 = 114
            v6 = v6
            _ = v33
        }
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
