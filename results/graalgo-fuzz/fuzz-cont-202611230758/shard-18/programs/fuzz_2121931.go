package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32) uint16 {
    fmt.Println(p0)
    v7 := []int{-2147483648, -318}
    p0 = p0
    if false {
        p0 = (p0 ^ p0)
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 bool = false
    var v2 int64 = -593
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 1
    v6 := make(map[uint32]int)
    for fi8 := 0; fi8 < 20; fi8++ {
        if (false || (!false)) {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: 356, F1: 1}
                v2 += v2
                fmt.Println(v1)
                _ = fi9
                _ = v10
            }
            v3 = (v3 + ((("foo bar" + ("\n\t" + (v3 + v3))) + v3) + "foo bar"))
        } else {
            v11 := GenS1{F0: 465, F1: 56}
            var v12 string = v3
            v4[0] = v5
            _ = v11
            _ = v12
        }
        switch ("hello" + ("hello" + "\n\t")) {
        case "hello":
            v3 = v3
        default:
            switch v3 {
            case "":
                var v13 uint32 = 959
                _ = v13
            case "foo bar":
                if (1 == 118) {
                    delete(v6, 0)
                    v5 += (^(v5))
                    switch fi8 {
                    case 750:
                        v14 := []int{}
                        _ = v14
                    default:
                        fmt.Println(fi8, fi8)
                    }
                } else {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 uint16 = 32767
                        if (!false) {
                            var v17 int16 = -32768
                            _ = v17
                        } else {
                            v1 = v1
                            if ((0 == 478) || ((fi8 >> ((305) & 63)) < fi8)) {
                                var v18 string = v3
                                fmt.Println(fi15)
                                _ = v18
                            }
                        }
                        var v19 uint8 = 149
                        _ = fi15
                        _ = v16
                        _ = v19
                    }
                }
            default:
                v20 := GenS1{F0: genFn0(42), F1: 2}
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := v4[(v5 | 192)]
            v2 -= (v2 * v2)
            _ = fi21
            _ = v22
        }
        _ = fi8
    }
    fmt.Println(v1, v5, v1)
    for fi23 := 0; fi23 < 12; fi23++ {
        v1 = v1
        _ = fi23
    }
    var v24 uint16 = 2
    for fi25 := 0; fi25 < 5; fi25++ {
        if v1 {
            fmt.Println(v1, v5, v1)
        }
        v2 = (-(v2))
        _ = fi25
    }
    v3 = v3
    v3 = v3
    var v26 float32 = (float32(0.1) * (1.5 / 0.1))
    var v27 int8 = -128
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 int32 = 1
        fmt.Println(v2, v27, v5)
        if (v26 < v26) {
            if true {
                v24 = v24
            } else {
                v27 += (v27 - 100)
                switch fi28 {
                case 2147483647:
                    v30 := []int{(fi28 ^ 32768)}
                    _ = v30
                case 9223372036854775807:
                    var v31 bool = (((0.0 + (v26 + v26)) * v26) == 1.5)
                    _ = v31
                case -9223372036854775808:
                    var v32 float64 = 0.0
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := GenS1{F0: 2, F1: 2}
                        var v35 int = fi28
                        switch v35 {
                        case -560:
                            v26 = v26
                        case 9223372036854775807:
                            v36 := GenS2{F0: 1, F1: (v5 & v5), F2: (v2 >> ((18446744073709551615) & 63))}
                            _ = v36
                        case 128:
                            switch 1 {
                            case -307:
                                if true {
                                    var v37 int32 = (^(v29))
                                    fmt.Println(v1, v37)
                                    v26 -= v26
                                    _ = v37
                                } else {
                                    switch (v35 - (fi33 & fi33)) {
                                    case 32768:
                                        v38 := v4[(v5 / ((v5) | 1))]
                                        _ = v38
                                    case 609:
                                        v26 = 3.14159
                                    }
                                    v39 := GenS3{F0: (v5 >> ((255) & 63))}
                                    v40 := GenS2{F0: 97, F1: 4294967295, F2: v2}
                                    _ = v39
                                    _ = v40
                                }
                                var v41 string = v3
                                _ = v41
                            }
                        default:
                            v3 = v3
                        }
                        v6[629] = 127
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            }
            v24 -= 646
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
