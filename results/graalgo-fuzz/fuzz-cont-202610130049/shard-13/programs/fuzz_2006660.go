package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int; F2 uint }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
func genFn0(p0 int8) uint32 {
    p0 -= p0
    var v5 float64 = 0.1
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 int16, p1 uint32, p2 string) int32 {
    var v6 uint = 0
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    if (!(244 <= 127)) {
        v9 := GenS1{F0: 127, F1: 4294967295, F2: 2}
        v10 := []int{32768, -2147483648, -409}
        _ = v10
        _ = v9
    } else {
        if true {
            var v11 uint32 = p1
            v12 := []int{}
            switch ("hello" + "foo bar") {
            case "":
                switch p2 {
                case "":
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                case "a":
                    var v13 int8 = -128
                    _ = v13
                default:
                    var v14 uint = v6
                    _ = v14
                }
            case "hello":
                switch p2 {
                case "\n\t":
                    v11 = genFn0(0)
                case "":
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v16 := GenS2{F0: 2, F1: 255}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = fi15
                        _ = v16
                    }
                default:
                    switch -442 {
                    case -747:
                        v12 = append(v12, -9223372036854775808)
                    }
                }
            default:
                switch p2 {
                case "\n\t":
                    var v17 uint8 = 105
                    _ = v17
                default:
                    fmt.Println(p2, p0)
                }
            }
            _ = v11
            _ = v12
        } else {
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 bool = true
    var v2 float32 = 0.1
    v3 := make(map[int]string)
    var v4 float32 = 1e10
    fmt.Println(v1, v4, v4)
    switch "hello" {
    case "a":
        v18 := GenS3{F0: -760, F1: 127, F2: 2}
        _ = v18
    case "":
        v19 := []int{}
        _ = v19
    case "\n\t":
        var v20 int16 = 781
        _ = v20
    }
    v21 := GenS2{F0: 763, F1: 255}
    v22 := v3[9223372036854775807]
    v2 = 3.14159
    switch 32768 {
    case 0:
        for fi23 := 0; fi23 < 18; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 int32 = 42
                v26 := []int{fi23, (fi24 / (((fi24 + 42)) | 1))}
                _ = fi24
                _ = v25
                _ = v26
            }
            delete(v3, fi23)
            fmt.Println(v22)
            _ = fi23
        }
    default:
        v22 = (v22 + v22)
    }
    fmt.Println(v4)
    if (1000 < -1) {
        var v27 string = v22
        v2 -= float32(0.0)
        v28 := GenS1{F0: -1, F1: 38, F2: 1}
        _ = v27
        _ = v28
    } else {
        delete(v3, -662)
        if (((-2.25 - (1.5 / (0.0 * ((float32((0.1 - (float64(((float32(1e10) - -2.25) + v4)) + (3.14159 - 1e10)))) + v2) / ((1e10 - v4) * (v2 * v4)))))) >= (v2 / v4)) || (0 == 9223372036854775807)) {
            v29 := []int{-9223372036854775808, 4294967295, 2147483647, 32768}
            v1 = (false && v1)
            v30 := GenS4{F0: 127, F1: 0}
            _ = v29
            _ = v30
        }
        v4 = (0.0 + v2)
    }
    v31 := GenS5{F0: 113}
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 uint8 = 253
        v34 := GenS4{F0: v33, F1: 0}
        if (!((v2 * v4) >= (v2 - v4))) {
            v35 := GenS6{F0: v22}
            switch ((((v22 + (v22 + v22)) + v22) + (("\n\t" + ("foo bar" + "hello")) + "")) + ((("hello" + ("foo bar" + v22)) + "\n\t") + "")) {
            case "a":
                if (!true) {
                    switch v22 {
                    case "":
                        v22 = v22
                    case "foo bar":
                        var v36 uint = 42
                        _ = v36
                    default:
                        var v37 int = fi32
                        _ = v37
                    }
                    var v38 uint32 = 883
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 10; fi39++ {
                        var v40 uint = 42
                        _ = fi39
                        _ = v40
                    }
                    if ((946 != 1) && (((((42 <= 9223372036854775807) || false) || v1) || false) && true)) {
                        var v41 int = (-(fi32))
                        if (0 > 651) {
                            for fi42 := 0; fi42 < 17; fi42++ {
                                v43 := GenS6{F0: v22}
                                var v44 float32 = float32(0.1)
                                var v45 uint32 = 980
                                _ = fi42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        }
                        var v46 string = v22
                        _ = v41
                        _ = v46
                    } else {
                        var v47 uint8 = v33
                        var v48 int = fi32
                        v33 -= 255
                        _ = v47
                        _ = v48
                    }
                }
            }
            _ = v35
        } else {
            v49 := v3[128]
            _ = len(v3)
            v50 := []int{fi32, (fi32 / ((fi32) | 1)), fi32, 129}
            _ = v49
            _ = v50
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    v22 = "\n\t"
    for fi51 := 0; fi51 < 17; fi51++ {
        if (((v1 || true) && v1) || false) {
            v22 = (v22 + v22)
            var v52 float32 = v4
            var v53 uint32 = 65535
            _ = v52
            _ = v53
        } else {
            v2 += (-2.25 + v2)
        }
        for fi54 := 0; fi54 < 2; fi54++ {
            v55 := []int{}
            var v56 float64 = (0.0 - 1.5)
            _ = fi54
            _ = v55
            _ = v56
        }
        v1 = (!v1)
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
}
