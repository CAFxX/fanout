package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0() bool {
    v5 := GenS1{F0: 1}
    if false {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{(^(fi6)), 0, fi6}
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS2{F0: -27, F1: 1, F2: 0}
    _ = v5
    _ = v8
    return true
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 2
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    if true {
        v9 := v3[42]
        _ = v9
    } else {
        v10 := []int{32767, 9223372036854775807}
        v11 := GenS1{F0: -9223372036854775808}
        if len(v10) > 0 {
            v10[0] = 32767
        }
        _ = v10
        _ = v11
    }
    if true {
        v12 := []int{32768, 478, 128}
        var v13 int64 = 3
        for fi14 := 0; fi14 < 3; fi14++ {
            v13 = 1
            v15 := GenS2{F0: 42, F1: -128, F2: -966}
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    switch -9223372036854775808 {
    case -981:
        v16 := GenS3{F0: 42}
        _ = v16
    case 911:
        switch -241 {
        case -9223372036854775808:
            v17 := GenS3{F0: 100}
            _ = v17
        case -221:
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 float64 = (0.0 - 0.0)
                _ = fi18
                _ = v19
            }
        case 2147483647:
            switch 4294967295 {
            case -884:
                fmt.Println(v2, v1)
            default:
                var v20 uint = (v1 % ((v1) | 1))
                _ = v20
            }
        }
    case 2147483647:
        if (int8(v2) < 0) {
            v21 := GenS4{F0: 100}
            v1 -= v1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 5; fi22++ {
                if true {
                    v23 := GenS5{F0: 349, F1: 18446744073709551615}
                    var v24 int = (fi22 + (fi22 >> (((^(v1))) & 63)))
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        if ((-2.25 + float32(float64((1e10 - -2.25)))) <= 0.1) {
                            if (150 <= 0) {
                                var v26 string = ("" + ("" + ""))
                                _ = v26
                            } else {
                                v27 := v4[v2]
                                var v28 uint64 = 232
                                v29 := v3[uint32(v1)]
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = len(v4)
                            switch "hello" {
                            case "":
                                v2 = (^(v2))
                            default:
                                v30 := v3[v2]
                                _ = v30
                            }
                            if (!false) {
                                v31 := []int{fi22, fi25}
                                v32 := GenS3{F0: 127}
                                v2 = (^(v2))
                                fmt.Println(fi22, v2)
                                v2 = 2147483648
                                var v33 int8 = 1
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = fi25
                    }
                    _ = len(v4)
                    var v34 float64 = 1.5
                    _ = v34
                }
                _ = fi22
            }
            v2 = (v2 - (v2 | 42))
        }
    default:
        for fi35 := 0; fi35 < 13; fi35++ {
            v36 := v3[65535]
            for fi37 := 0; fi37 < 10; fi37++ {
                var v38 uint16 = 933
                fmt.Println(fi35, v36)
                v4[42] = "foo bar"
                v1 = 327
                var v39 float64 = 3.14159
                _ = fi37
                _ = v38
                _ = v39
            }
            v2 = uint32(v1)
            _ = fi35
            _ = v36
        }
    }
    v40 := GenS1{F0: -1}
    var v41 int32 = -2147483648
    v1 += 255
    var v42 uint8 = 154
    var v43 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
