package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint; F2 int16 }
type GenS5 struct { F0 uint16; F1 int; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", "a", "a", "\n\t"}
    var v3 uint16 = 255
    var v4 uint8 = 128
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: -809}
        switch "a" {
        case "\n\t":
            v3 += v3
        default:
            fmt.Println(v4, fi7)
        }
        v9 := GenS2{F0: fi7, F1: 127, F2: (false && (!(!true)))}
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v10 uint = 255
    v10 = v10
    if (-2.25 < (((float64(1.5) / 1.5) * (float64(-2.25) / 1.0)) * -2.25)) {
        var v11 uint8 = v4
        switch -9223372036854775808 {
        case 127:
            var v12 uint8 = (^(v4))
            _ = v12
        case 874:
            switch -1 {
            case 128:
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v2) > 0 {
                        v2[0] = (("\n\t" + "foo bar") + "\n\t")
                    }
                    _ = fi13
                }
            case 0:
                switch 425 {
                case 9223372036854775807:
                    if false {
                        v14 := GenS3{F0: (v11 * uint8(v10))}
                        var v15 float32 = 0.1
                        v15 += 1e10
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint = v10
                        var v17 int = 1
                        _ = v16
                        _ = v17
                    }
                default:
                    v10 += v10
                }
            case 9223372036854775807:
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := GenS2{F0: (fi18 * (fi18 - 667)), F1: 1, F2: (true || (true || (-643 <= -1)))}
                    _ = fi18
                    _ = v19
                }
            }
        }
        fmt.Println(v11, v3, v10)
        _ = v11
    } else {
        v3 = v3
        v20 := GenS4{F0: 100, F1: v10, F2: int16(v10)}
        _ = v20
    }
    var v21 int = 0
    switch v21 {
    case -718:
        v21 = (v21 - (v21 * 127))
    case 4294967295:
        if true {
            if (!true) {
                if false {
                    if true {
                        if false {
                            v21 = v21
                            for fi22 := 0; fi22 < 3; fi22++ {
                                v10 = v10
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    v3 += (v3 % (((v3 % ((v3) | 1))) | 1))
                                    _ = fi23
                                }
                                switch (v21 & 32767) {
                                case 9223372036854775807:
                                    var v24 int = fi22
                                    _ = v24
                                case -158:
                                    v25 := GenS3{F0: uint8(v3)}
                                    _ = v25
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = ("\n\t" + ("hello" + ("hello" + "")))
                                    }
                                }
                                _ = fi22
                            }
                        } else {
                            v26 := GenS5{F0: v3, F1: v21, F2: 1}
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v27 float32 = 3.14159
                                _ = v27
                            default:
                                if (((0.1 / ((1.5 + 3.14159) + float64(0.0))) - 3.14159) != (3.14159 * ((-2.25 - float64(0.0)) * (-2.25 * (1e10 - 1.5))))) {
                                    var v28 float32 = (3.14159 - 1e10)
                                    v29 := GenS5{F0: v3, F1: 42, F2: 42}
                                    fmt.Println(v21, v3)
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        v30 := GenS6{F0: true}
                                        _ = v30
                                    default:
                                        if true {
                                            var v31 int = 221
                                            _ = v31
                                        } else {
                                            switch int(v3) {
                                            case 32767:
                                                v32 := GenS6{F0: false}
                                                _ = v32
                                            case -2147483648:
                                                v3 = v3
                                            case 127:
                                                var v33 int8 = -40
                                                _ = v33
                                            }
                                            v34 := GenS2{F0: v21, F1: -59, F2: true}
                                            _ = v34
                                        }
                                    }
                                    if (!true) {
                                        var v35 int64 = -721
                                        var v36 uint64 = uint64(v10)
                                        v4 += (v4 >> (((v10 + 255)) & 63))
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                            if (true || (false || (!false))) {
                                _ = len(v5)
                                fmt.Println(v21, v10)
                                v4 += (v4 >> ((18446744073709551615) & 63))
                            }
                            _ = v26
                        }
                    }
                }
            }
            if true {
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 string = ("hello" + (("" + ("\n\t" + "\n\t")) + "foo bar"))
                    _ = fi37
                    _ = v38
                }
                var v39 uint64 = uint64(v10)
                _ = v39
            }
            var v40 uint16 = (^(v3))
            _ = v40
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v41 := GenS7{F0: ("" + "")}
            delete(v1, uint32(v4))
            _ = v41
        }
    case 450:
        v42 := v6["a"]
        _ = v42
    default:
        v10 += (v10 + (v10 << (((v10 / (((^(v10))) | 1))) & 63)))
    }
    var v43 uint16 = v3
    var v44 int8 = 100
    var v45 uint64 = 42
    var v46 int32 = -1
    if ((^(v10)) <= v10) {
        var v47 int32 = (-(v46))
        _ = v47
    } else {
        _ = len(v6)
        switch v21 {
        case 4294967295:
            v45 -= v45
        default:
            if false {
                var v48 int32 = -2147483648
                v2 = append(v2, ((("" + "hello") + "hello") + "hello"))
                var v49 string = ""
                _ = v48
                _ = v49
            }
        }
        _ = len(v2)
    }
    v45 += 124
    v46 -= v46
    v46 -= (v46 / ((v46) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
