package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint32 }
func genFn0() uint16 {
    switch 32767 {
    case 127:
        var v7 uint = 42
        _ = v7
    }
    switch "hello" {
    case "a":
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: 584, F1: 127, F2: -13}
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 int8 = 0
                _ = fi11
                _ = v12
            }
            _ = fi8
        }
    case "\n\t":
        if false {
            if false {
                switch 4294967295 {
                case 166:
                    switch "hello" {
                    case "":
                        if false {
                            var v13 bool = true
                            fmt.Println(v13)
                            v13 = ((18446744073709551615 >= 1) || (!(617 != 18446744073709551615)))
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 20; fi14++ {
                                var v15 float32 = 3.14159
                                _ = fi14
                                _ = v15
                            }
                            switch "a" {
                            case "":
                                var v16 string = ""
                                _ = v16
                            }
                        }
                    default:
                        v17 := GenS1{F0: 2147483647, F1: 1, F2: 100}
                        _ = v17
                    }
                case 42:
                    if false {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            for fi19 := 0; fi19 < 14; fi19++ {
                                var v20 uint16 = 255
                                v20 -= v20
                                _ = fi19
                                _ = v20
                            }
                            _ = fi18
                        }
                    } else {
                        var v21 float64 = ((((-2.25 / ((3.14159 - (float64(0.0) - float64(0.1))) + 3.14159)) + (0.1 - 3.14159)) / ((0.1 + ((-2.25 / (1e10 * 0.1)) - ((1e10 / -2.25) + 1.5))) / (float64(float32(-2.25)) + 0.0))) + (3.14159 / -2.25))
                        v21 = v21
                        _ = v21
                    }
                case 128:
                    var v22 int32 = 541
                    _ = v22
                default:
                    var v23 uint64 = 386
                    _ = v23
                }
                var v24 uint32 = 65535
                _ = v24
            }
        } else {
            var v25 bool = true
            fmt.Println(v25)
            _ = v25
        }
    case "":
        var v26 uint8 = 128
        _ = v26
    default:
        v27 := GenS1{F0: 0, F1: 96, F2: 100}
        _ = v27
    }
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    v4 := []int{-998, 128, 9223372036854775807, 9223372036854775807}
    var v5 uint64 = 42
    v6 := []int{4294967295, -324, 42, -9223372036854775808}
    var v28 uint16 = 1
    v5 += v5
    switch 128 {
    case 9223372036854775807:
        fmt.Println(v28)
    case 32768:
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint32 = 4294967295
            v28 += v28
            _ = fi29
            _ = v30
        }
    default:
        v31 := GenS1{F0: 65535, F1: 127, F2: 42}
        _ = v31
    }
    v5 = (v5 + 1)
    var v32 uint32 = 2
    var v33 uint16 = genFn0()
    v34 := GenS1{F0: (v32 + v32), F1: 255, F2: 0}
    if (v33 >= (v33 | genFn0())) {
        for fi35 := 0; fi35 < 13; fi35++ {
            switch "" {
            case "":
                v36 := GenS2{F0: "hello", F1: 42}
                _ = v36
            case "hello":
                v32 -= uint32(v5)
            case "foo bar":
                v28 += 1
            }
            var v37 uint = 68
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi35
            _ = v37
        }
    } else {
        v33 -= (v28 + v28)
        v38 := GenS3{F0: v32}
        fmt.Println(v33, v32, v28)
        _ = v38
    }
    v33 -= genFn0()
    fmt.Println(v33)
    v33 = (v28 % ((uint16(v32)) | 1))
    v3[873] = "hello"
    v1["\n\t"] = -638
    fmt.Println(v5, v32, v32)
    if false {
        if true {
            v2["\n\t"] = -879
            var v39 uint = 42
            v32 += (v32 % (((v32 >> ((42) & 63))) | 1))
            _ = v39
        }
    } else {
        var v40 int32 = -2147483648
        v41 := GenS2{F0: "", F1: (v40 % ((1) | 1))}
        switch 4294967295 {
        case -86:
            var v42 bool = false
            _ = v42
        }
        _ = v40
        _ = v41
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
