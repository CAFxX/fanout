package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64) int16 {
    v4 := []int{128, 0, -36}
    if true {
        var v5 int8 = 115
        _ = v5
    }
    _ = p0
    _ = v4
    return 328
}

func main() {
    var v1 int16 = 1
    var v2 bool = (766 == 9223372036854775807)
    v3 := make(map[uint32]int)
    fmt.Println(v1, v1)
    v6 := GenS1{F0: "\n\t", F1: 42, F2: -526}
    var v7 int16 = v1
    v7 -= v7
    v2 = (!(0 >= 471))
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 = int16(fi8)
        _ = fi8
    }
    v2 = (255 == 127)
    v2 = true
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: 1, F2: (fi9 >> ((1) & 63))}
        _ = fi9
        _ = v10
    }
    if (!((v1 <= v1) && (v2 || true))) {
        if (true || (150 > 255)) {
            v7 -= -811
            v11 := GenS1{F0: "foo bar", F1: 65535, F2: 607}
            fmt.Println(v2)
            _ = v11
        } else {
            var v12 uint8 = 1
            if (183 > 934) {
                if false {
                    v2 = false
                    var v13 int32 = -946
                    var v14 uint = 662
                    _ = v13
                    _ = v14
                } else {
                    switch 614 {
                    case 9223372036854775807:
                        v15 := GenS2{F0: 0}
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 4; fi16++ {
                            if false {
                                var v17 bool = false
                                _ = v17
                            } else {
                                if v2 {
                                    if v2 {
                                        for fi18 := 0; fi18 < 16; fi18++ {
                                            v19 := []int{(-(fi16)), fi16, (fi18 ^ 0)}
                                            _ = fi18
                                            _ = v19
                                        }
                                    } else {
                                        if v2 {
                                            if ((0.1 - ((1.5 * 0.0) / ((0.1 + 1e10) - 1.5))) < (0.1 / 1.5)) {
                                                v2 = ((-(fi16)) <= fi16)
                                            }
                                            var v20 bool = v2
                                            _ = v20
                                        }
                                        if (v12 <= (v12 ^ (v12 % (((v12 & (v12 << ((18446744073709551615) & 63)))) | 1)))) {
                                            if (("" + "foo bar") <= "a") {
                                                var v21 string = ""
                                                var v22 uint8 = (v12 / ((127) | 1))
                                                fmt.Println(v22, fi16, v12)
                                                _ = v21
                                                _ = v22
                                            }
                                        }
                                        v23 := GenS3{F0: v1}
                                        _ = v23
                                    }
                                    v7 = v1
                                }
                                v12 -= (^(v12))
                                var v24 int64 = int64(fi16)
                                _ = v24
                            }
                            _ = fi16
                        }
                    }
                    v1 = (v1 / (((v1 - -279)) | 1))
                    switch "foo bar" {
                    case "\n\t":
                        v12 = (v12 - (v12 | (^(v12))))
                    case "foo bar":
                        var v25 int = 32768
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            fmt.Println(v2, v2, v12)
                            v27 := []int{}
                            var v28 uint8 = v12
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                }
            }
            v2 = (v2 && (((float32((0.1 + float64((3.14159 - ((float32(3.14159) + ((float32(0.0) + 1e10) / 1.5)) + -2.25))))) / (((float32((((1e10 / (0.1 / 1.5)) / 1.0) * 1e10)) / float32(-2.25)) * 3.14159) + (3.14159 - 1.5))) / 1.5) == 0.1))
            _ = v12
        }
        v7 = 910
        var v29 int64 = 9223372036854775807
        _ = v29
    } else {
        v7 = v7
        v30 := []int{2147483647, -242}
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := GenS1{F0: "hello", F1: 288, F2: fi31}
            v1 += v7
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    v1 = (v1 % ((v1) | 1))
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint = 4294967295
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 string = "hello"
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    v2 = v2
    switch 42 {
    case 128:
        delete(v3, 0)
    default:
        v37 := GenS1{F0: "\n\t", F1: 32767, F2: 1}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
