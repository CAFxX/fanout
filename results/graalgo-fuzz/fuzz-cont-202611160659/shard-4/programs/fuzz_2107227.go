package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 string }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 uint32) uint8 {
    var v4 float32 = 1.5
    v5 := GenS1{F0: false, F1: 642, F2: 4294967295}
    p1 = 1
    var v6 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 string = ("foo bar" + "a")
    v2 := make(map[uint32]bool)
    var v3 int16 = -32768
    v3 -= -32768
    v3 = (v3 ^ -911)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS2{F0: -1, F1: 42, F2: 1}
        _ = len(v2)
        if (!(42 < -22)) {
            v1 = ("hello" + ((v1 + v1) + v1))
            var v9 int = fi7
            _ = v9
        } else {
            v1 = "a"
            v3 -= 881
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: (!false), F1: 42, F2: 67}
    v1 = v1
    var v11 float32 = (0.1 * 1e10)
    v12 := GenS3{F0: v3, F1: 1, F2: (v1 + ("" + v1))}
    var v13 float64 = 3.14159
    for fi14 := 0; fi14 < 11; fi14++ {
        v3 += (v3 + (v3 / ((-907) | 1)))
        _ = fi14
    }
    v15 := []int{127, 0, 2147483647, 32768}
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 float32 = (0.1 - 1e10)
            switch (((v1 + ("\n\t" + v1)) + v1) + v1) {
            case "hello":
                var v19 string = v1
                _ = v19
            default:
                if (2 != 18446744073709551615) {
                    v2[888] = false
                }
                var v20 float32 = (float32(v13) * ((0.1 - 1e10) * 1.5))
                var v21 uint8 = 144
                _ = v20
                _ = v21
            }
            if len(v15) > 0 {
                v15[0] = (fi17 ^ 127)
            }
            v13 -= 3.14159
            v22 := GenS3{F0: -92, F1: fi17, F2: (v1 + (v1 + v1))}
            _ = fi17
            _ = v18
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint32 = 4294967295
            if true {
                var v25 float64 = 1e10
                if true {
                    v1 = "a"
                    _ = len(v2)
                    switch fi23 {
                    case -253:
                        var v26 uint32 = v24
                        _ = v26
                    case 127:
                        if true {
                            v11 = 0.0
                            v27 := v2[v24]
                            _ = v27
                        }
                    }
                    switch v1 {
                    case "a":
                        var v28 int64 = 42
                        _ = v28
                    case "":
                        v11 = v11
                    case "foo bar":
                        v29 := GenS4{F0: int8(v24), F1: 254}
                        _ = v29
                    default:
                        switch v1 {
                        case "hello":
                            v30 := GenS4{F0: 88, F1: 0}
                            _ = v30
                        case "":
                            if true {
                                v25 = float64(v11)
                            } else {
                                if false {
                                    v24 += (v24 << ((4294967295) & 63))
                                }
                                v31 := GenS5{F0: -128}
                                v25 = float64(1e10)
                                fmt.Println(fi23, v25)
                                var v32 uint16 = 1
                                _ = v31
                                _ = v32
                            }
                        case "a":
                            switch v1 {
                            case "a":
                                v11 = (float32(v25) / v11)
                            case "hello":
                                if true {
                                    v3 = v3
                                    v33 := GenS6{F0: 42, F1: -128, F2: 65535}
                                    v34 := GenS7{F0: 4294967295}
                                    _ = v33
                                    _ = v34
                                } else {
                                    fmt.Println(fi23)
                                }
                                var v35 uint16 = 32767
                                _ = v35
                            }
                        }
                    }
                } else {
                    if (true || false) {
                        v36 := GenS7{F0: 2}
                        _ = v36
                    }
                }
                v37 := v2[v24]
                v38 := GenS1{F0: ((1 != 58) || (!v37)), F1: 255, F2: 42}
                var v39 uint8 = genFn0(2, 1, 910)
                _ = v25
                _ = v37
                _ = v38
                _ = v39
            } else {
                if false {
                    switch (v1 + "a") {
                    case "\n\t":
                        if (3.14159 < (v11 + (float32(((v13 - 3.14159) - (v13 + v13))) + (3.14159 / v11)))) {
                            v40 := GenS4{F0: 1, F1: 128}
                            _ = v40
                        }
                    case "":
                        if (1e10 == float32(v13)) {
                            var v41 int8 = 1
                            if true {
                                v42 := GenS8{F0: (-(v41)), F1: -9223372036854775808, F2: (^(v24))}
                                _ = v42
                            }
                            _ = v41
                        }
                    case "hello":
                        v3 -= 1000
                    default:
                        switch fi16 {
                        case 229:
                            v13 = v13
                        default:
                            var v43 int16 = (v3 & (v3 ^ int16(v24)))
                            _ = v43
                        }
                    }
                    switch (fi16 + (-(fi16))) {
                    case 42:
                        var v44 bool = ((1.5 / (v13 + float64(v11))) < (v13 - (-2.25 - v13)))
                        _ = v44
                    case 32767:
                        var v45 float32 = v11
                        _ = v45
                    }
                }
                var v46 int = (fi16 | 32767)
                _ = v46
            }
            v1 = "a"
            _ = fi23
            _ = v24
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
