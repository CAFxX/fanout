package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 int32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int32, p1 uint16) int32 {
    p1 = p1
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = 465
        _ = fi6
        _ = v7
    }
    v8 := []int{127, -349, -236}
    _ = p0
    _ = p1
    _ = v8
    return (p0 + p0)
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]int)
    var v3 float64 = (0.1 + 1.5)
    var v4 int16 = 42
    v5 := make(map[int]uint32)
    fmt.Println(v3, v4)
    if true {
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v3, fi9, v4)
            _ = fi9
        }
    } else {
        var v10 uint32 = 2
        _ = v10
    }
    v4 = (^(v4))
    switch 9223372036854775807 {
    case -9223372036854775808:
        v11 := []int{0, 127}
        _ = v11
    case -874:
        switch (("hello" + "") + "hello") {
        case "a":
            if false {
                v3 = float64(0.1)
                if (!false) {
                    fmt.Println(v3, v4)
                    v4 += v4
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v13 := []int{(fi12 / (((fi12 - fi12)) | 1)), 0}
                        _ = fi12
                        _ = v13
                    }
                }
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS1{F0: (v1 >> ((929) & 63)), F1: 65535}
                    _ = fi14
                    _ = v15
                }
            }
        case "foo bar":
            v16 := GenS2{F0: 262}
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 15; fi17++ {
                for fi18 := 0; fi18 < 6; fi18++ {
                    switch (fi17 * 105) {
                    case 174:
                        v19 := GenS3{F0: 0}
                        _ = v19
                    case 846:
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                var v20 string = "\n\t"
                                _ = v20
                            }
                        }
                    case 42:
                        if true {
                            var v21 int64 = 684
                            switch -9223372036854775808 {
                            case 186:
                                if false {
                                    var v22 int8 = int8(v4)
                                    _ = v22
                                } else {
                                    var v23 float32 = ((-2.25 / 0.1) + ((-2.25 * float32(((1e10 * v3) * v3))) - -2.25))
                                    if (1 >= 0) {
                                        var v24 int16 = -1
                                        v24 = 1000
                                        _ = v24
                                    } else {
                                        v23 = 3.14159
                                        v25 := GenS2{F0: 74}
                                        if (false || false) {
                                            v4 = 1
                                        }
                                        fmt.Println(fi18, v23)
                                        _ = v25
                                    }
                                    if (true && (-9223372036854775808 < fi18)) {
                                        v4 = v4
                                        var v26 uint16 = 10
                                        _ = v26
                                    } else {
                                        v27 := v5[32767]
                                        _ = v27
                                    }
                                    _ = v23
                                }
                            default:
                                v28 := GenS4{F0: 65535, F1: 42, F2: 42}
                                _ = v28
                            }
                            v29 := GenS2{F0: 800}
                            _ = v21
                            _ = v29
                        }
                    default:
                        v1 += v1
                    }
                    _ = fi18
                }
                _ = fi17
            }
        default:
            v1 -= v1
        }
    default:
        v30 := []int{32767, 429, -717}
        _ = v30
    }
    v3 = ((0.0 * float64(-2.25)) / v3)
    v5[-568] = 42
    v3 -= -2.25
    v31 := []int{32768, 0, 2147483647}
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint64 = 1
        _ = fi32
        _ = v33
    }
    v34 := GenS5{F0: 11, F1: 0, F2: genFn0((v1 - -2147483648), 743)}
    v5[2147483647] = 468
    v35 := GenS6{F0: (v4 >> ((279) & 63))}
    if false {
        switch 4294967295 {
        case -2147483648:
            v36 := v2[""]
            _ = v36
        case -9223372036854775808:
            v4 = 32767
        }
    } else {
        v4 -= (v4 | (v4 - (v4 >> ((4294967295) & 63))))
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    var v37 float64 = v3
    v37 += v3
    for fi38 := 0; fi38 < 16; fi38++ {
        v31 = append(v31, 127)
        _ = fi38
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
