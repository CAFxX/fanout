package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint32; F2 int32 }
func genFn0() int8 {
    var v5 uint8 = 205
    var v6 bool = (!false)
    var v7 int32 = 665
    if (v6 || v6) {
        v6 = (!true)
        v8 := GenS1{F0: 42}
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS2{F0: v5, F1: -631}
            if (0 < 42) {
                v11 := GenS3{F0: 1, F1: 1, F2: (fi9 % (((fi9 ^ fi9)) | 1))}
                _ = v11
            } else {
                switch fi9 {
                case -9223372036854775808:
                    v7 += v7
                case 294:
                    switch 595 {
                    case 98:
                        var v12 string = ""
                        _ = v12
                    case 9223372036854775807:
                        var v13 uint16 = 0
                        _ = v13
                    case 0:
                        for fi14 := 0; fi14 < 11; fi14++ {
                            fmt.Println(fi14, v5)
                            fmt.Println(v5)
                            _ = fi14
                        }
                    default:
                        v15 := GenS4{F0: 255, F1: 918}
                        _ = v15
                    }
                }
                switch fi9 {
                case 127:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 float64 = (0.1 + 3.14159)
                        var v18 uint = 383
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    if (96 < (v5 % ((v5) | 1))) {
        switch "foo bar" {
        case "hello":
            v19 := GenS1{F0: -9223372036854775808}
            _ = v19
        case "a":
            switch "\n\t" {
            case "\n\t":
                v20 := GenS2{F0: v5, F1: 42}
                _ = v20
            case "a":
                for fi21 := 0; fi21 < 19; fi21++ {
                    if v6 {
                        switch (("\n\t" + ("foo bar" + "foo bar")) + ("\n\t" + "a")) {
                        case "\n\t":
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 int8 = 30
                                _ = fi22
                                _ = v23
                            }
                        case "":
                            fmt.Println(v7, fi21, v7)
                        default:
                            v24 := []int{127, (fi21 * fi21), fi21, -9223372036854775808}
                            _ = v24
                        }
                        var v25 float32 = 3.14159
                        _ = v25
                    } else {
                        if v6 {
                            var v26 uint8 = (^(v5))
                            var v27 uint16 = 2
                            _ = v26
                            _ = v27
                        } else {
                            var v28 float32 = float32(0.1)
                            _ = v28
                        }
                    }
                    _ = fi21
                }
            default:
                if ((912 > 1) && (1.5 != (float32((3.14159 + -2.25)) * (3.14159 - ((1e10 + 0.1) + 0.1))))) {
                    v29 := GenS4{F0: 0, F1: 321}
                    _ = v29
                }
            }
        }
        v30 := []int{9223372036854775807}
        _ = v30
    }
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 151
    var v4 uint32 = 42
    fmt.Println(v3, v2, v2)
    var v31 uint16 = 1
    var v32 uint64 = 406
    if (824 > -1) {
        v33 := []int{}
        _ = v33
    }
    v34 := GenS5{F0: 1, F1: 2147483648, F2: -45}
    var v35 int64 = -120
    var v36 uint32 = v4
    var v37 int16 = 823
    v31 -= 2
    v1 = "hello"
    v38 := []int{1, 592, -7}
    fmt.Println(v1, v3, v37)
    fmt.Println(v4, v31, v4)
    v4 -= 2
    fmt.Println(v35, v3, v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
