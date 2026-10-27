package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 bool }
func genFn0(p0 uint, p1 uint32, p2 uint32) int32 {
    var v4 uint64 = 344
    switch 869 {
    case 128:
        fmt.Println(p0, v4, p1)
    case 2147483647:
        var v5 float64 = 0.0
        _ = v5
    case 127:
        p0 += (p0 ^ (p0 * (^(p0))))
    default:
        if (42 >= -55) {
            var v6 int64 = -549
            _ = v6
        }
    }
    v7 := []int{127, 128, -116}
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 2147483647
}
func genFn1(p0 uint64, p1 int, p2 int8) uint64 {
    p2 -= p2
    var v8 int8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (^(p0))
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    var v3 int16 = 42
    v1 -= (v1 >> ((532) & 63))
    v9 := []int{383}
    for fi10 := 0; fi10 < 18; fi10++ {
        switch fi10 {
        case 9223372036854775807:
            if false {
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int8 = -86
                    _ = fi11
                    _ = v12
                }
                switch "" {
                case "foo bar":
                    v3 += (v3 ^ 0)
                case "\n\t":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v3 = int16(fi13)
                        v14 := GenS1{F0: 100, F1: 42}
                        v9 = append(v9, (fi13 * (fi13 << ((65535) & 63))))
                        _ = fi13
                        _ = v14
                    }
                case "":
                    var v15 float64 = 1.5
                    _ = v15
                }
                fmt.Println(v1, fi10, fi10)
            } else {
                if ((false || (0.0 < 0.1)) || false) {
                    v16 := GenS1{F0: 100, F1: -1}
                    v3 = v3
                    fmt.Println(fi10, v3)
                    _ = v16
                }
                v1 = -1
                fmt.Println(v1, v3, fi10)
            }
        case -867:
            var v17 string = "foo bar"
            _ = v17
        case -243:
            v1 = (v3 << ((255) & 63))
        default:
            if (!(0.0 > (1e10 + (((0.1 / -2.25) - (3.14159 + 0.1)) / -2.25)))) {
                if false {
                    if (!false) {
                        var v18 bool = (!true)
                        v18 = v18
                        _ = len(v9)
                        _ = v18
                    }
                } else {
                    var v19 int16 = 1000
                    v20 := GenS1{F0: -1, F1: 100}
                    _ = v19
                    _ = v20
                }
                v21 := v2[0]
                _ = v21
            } else {
                var v22 uint64 = 4294967295
                v22 = v22
                for fi23 := 0; fi23 < 4; fi23++ {
                    if ((!(v3 == v1)) || (true && (42 == 42))) {
                        v22 = (v22 / ((v22) | 1))
                        var v24 uint = 255
                        switch -9223372036854775808 {
                        case -807:
                            if true {
                                v3 = -198
                                v2[2147483647] = 152
                            } else {
                                v24 = (^(v24))
                                var v25 bool = false
                                _ = v25
                            }
                            v24 = 42
                        }
                        _ = v24
                    } else {
                        switch "" {
                        case "":
                            if len(v9) > 0 {
                                v9[0] = fi10
                            }
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                v22 += (^(v22))
                                var v26 bool = false
                                _ = v26
                            case "hello":
                                if (false && (true || false)) {
                                    var v27 uint32 = 1
                                    v1 += int16(v27)
                                    if true {
                                        _ = len(v2)
                                    } else {
                                        v1 = 42
                                        var v28 bool = (!(4294967295 == 42))
                                        _ = v28
                                    }
                                    _ = v27
                                }
                            default:
                                delete(v2, 2)
                            }
                        }
                        switch "" {
                        case "a":
                            v29 := GenS1{F0: -67, F1: 127}
                            _ = v29
                        case "\n\t":
                            v22 += v22
                        case "hello":
                            v30 := GenS2{F0: 42, F1: v22}
                            _ = v30
                        }
                    }
                    _ = fi23
                }
                _ = v22
            }
        }
        _ = fi10
    }
    v31 := GenS1{F0: 0, F1: 60}
    if ((3.14159 - ((3.14159 * float64(float32(-2.25))) + (1e10 - 0.1))) < ((1.5 * float64((1e10 + float32(0.0)))) + 0.1)) {
        v32 := GenS3{F0: 2, F1: 2, F2: (true && true)}
        _ = v32
    } else {
        v1 -= -1
        v1 += (v1 % (((v3 * v1)) | 1))
        var v33 float32 = 0.0
        _ = v33
    }
    var v34 int = -292
    if false {
        var v35 uint16 = 0
        _ = v35
    }
    v34 = -2147483648
    v1 = 234
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v9
}
