package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
func genFn0(p0 uint) uint32 {
    v6 := GenS1{F0: -1, F1: 0}
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v6
    return 326
}

func main() {
    var v1 int16 = -353
    var v2 uint8 = 127
    v3 := make(map[int]uint32)
    v4 := []int{127, 9223372036854775807}
    var v5 uint32 = 4294967295
    v4 = append(v4, 435)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS2{F0: 2147483648, F1: -1, F2: 0}
        if ((!(30 != -1)) || ((false || (42 >= (v1 - v1))) && false)) {
            v5 = (v5 | (^(v5)))
            v1 = (v1 % ((-32768) | 1))
        } else {
            v5 = (v5 * 872)
            v5 -= (^(v5))
        }
        if true {
            v1 = -1
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS2{F0: v5, F1: -1, F2: -1}
                v2 = (v2 * 128)
                if false {
                    var v11 int8 = 0
                    v12 := v3[fi9]
                    _ = v11
                    _ = v12
                } else {
                    if false {
                        v3[fi9] = 180
                        var v13 uint16 = 42
                        v2 += (v2 >> ((398) & 63))
                        _ = v13
                    } else {
                        var v14 int8 = 82
                        switch (^(fi9)) {
                        case 4294967295:
                            if ((v2 % (((^(v2))) | 1)) != v2) {
                                v5 = v5
                            } else {
                                var v15 uint32 = uint32(v2)
                                v5 += v5
                                v16 := GenS3{F0: 2, F1: (v1 - (v1 | v1)), F2: v1}
                                v15 = v5
                                var v17 uint32 = v15
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            v5 -= v5
                            var v18 float32 = 3.14159
                            _ = v18
                        default:
                            if true {
                                delete(v3, (^(fi7)))
                                fmt.Println(v1)
                            }
                        }
                        switch ("" + ("foo bar" + "hello")) {
                        case "a":
                            v14 = int8(v1)
                        case "":
                            if true {
                                v19 := v3[fi7]
                                fmt.Println(fi9, v19, v1)
                                _ = v19
                            }
                        case "hello":
                            if true {
                                v20 := GenS1{F0: -32768, F1: v2}
                                _ = v20
                            } else {
                                v21 := GenS2{F0: genFn0(42), F1: 97, F2: int32(v1)}
                                v1 += (v1 & (-(v1)))
                                _ = v21
                            }
                        default:
                            var v22 uint32 = v5
                            _ = v22
                        }
                        _ = v14
                    }
                    v23 := GenS4{F0: 2147483647, F1: 1, F2: 255}
                    if ((51 > 32767) || false) {
                        var v24 float32 = 0.0
                        var v25 uint16 = 65535
                        _ = v24
                        _ = v25
                    } else {
                        v2 = v2
                        v1 = v1
                        if len(v4) > 0 {
                            v4[0] = -920
                        }
                        _ = len(v4)
                    }
                    _ = v23
                }
                _ = fi9
                _ = v10
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 uint8 = v2
                v28 := GenS1{F0: v1, F1: v2}
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            switch fi7 {
            case 128:
                for fi29 := 0; fi29 < 4; fi29++ {
                    fmt.Println(v2)
                    switch "a" {
                    case "\n\t":
                        var v30 uint = 65535
                        _ = v30
                    case "":
                        var v31 uint16 = 596
                        _ = v31
                    default:
                        var v32 int16 = v1
                        _ = v32
                    }
                    _ = fi29
                }
            case 36:
                switch fi7 {
                case 42:
                    var v33 int16 = (v1 << ((42) & 63))
                    _ = v33
                case -152:
                    var v34 int16 = v1
                    _ = v34
                }
            case 4294967295:
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 int = fi35
                    _ = fi35
                    _ = v36
                }
            }
        }
        _ = fi7
        _ = v8
    }
    var v37 uint32 = genFn0(4294967295)
    var v38 bool = ((v5 << ((255) & 63)) != (v5 + 0))
    switch "" {
    case "foo bar":
        v39 := GenS1{F0: (v1 * (v1 & (v1 << ((1) & 63)))), F1: (v2 * (v2 * (v2 >> ((0) & 63))))}
        _ = v39
    case "\n\t":
        v37 += uint32(v2)
    case "hello":
        v5 = v5
    }
    var v40 float32 = (float32((3.14159 + 0.0)) * 3.14159)
    for fi41 := 0; fi41 < 2; fi41++ {
        v42 := v3[fi41]
        _ = fi41
        _ = v42
    }
    v1 = (v1 - 1000)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
