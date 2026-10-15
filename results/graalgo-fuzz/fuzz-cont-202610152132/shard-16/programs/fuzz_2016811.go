package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 uint32, p1 bool, p2 int32) uint64 {
    v5 := GenS1{F0: 2147483647, F1: 42}
    var v6 int8 = 26
    switch 1 {
    case 128:
        v7 := []int{4294967295, 32767}
        _ = v7
    case 766:
        v8 := GenS2{F0: ""}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 1
    var v3 int16 = 0
    var v4 float64 = ((0.0 + float64(((0.1 + ((0.0 - 1.5) * 0.1)) - (0.0 - 0.0)))) - -2.25)
    if (true || (-1 > 1)) {
        if ((false || (42 != (v2 >> ((2) & 63)))) || ("\n\t" >= ("foo bar" + "hello"))) {
            var v9 uint64 = 4294967295
            v10 := []int{9223372036854775807, -9223372036854775808, 2147483647, 32768}
            v10 = append(v10, 127)
            _ = v10
            _ = v9
        }
    } else {
        var v11 float64 = v4
        _ = v11
    }
    v4 -= v4
    if (0 < 618) {
        v12 := []int{2147483647, 785}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v12
    }
    v2 = (v2 & (v2 * 1))
    fmt.Println(v3, v3, v2)
    switch -264 {
    case -2147483648:
        var v13 float64 = 1e10
        _ = v13
    case 294:
        v14 := []int{-1, -645}
        _ = v14
    case -1:
        if (216 >= 1) {
            switch 4294967295 {
            case 9223372036854775807:
                v2 = (v2 + (v2 / (((v2 % (((v2 + (v2 ^ (v2 / (((v2 ^ 2)) | 1))))) | 1))) | 1)))
            case -1:
                var v15 uint64 = 934
                _ = v15
            case -286:
                for fi16 := 0; fi16 < 5; fi16++ {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if true {
                            if false {
                                v18 := GenS3{F0: v3, F1: -43, F2: false}
                                v19 := []int{32767, (fi16 << ((4294967295) & 63)), fi17}
                                if (!false) {
                                    v3 += 0
                                    fmt.Println(v3)
                                }
                                _ = v18
                                _ = v19
                            } else {
                                v20 := []int{(fi16 << (((v1 << ((v1) & 63))) & 63)), 825, (fi16 | 127), fi16}
                                var v21 int = fi16
                                v22 := GenS1{F0: -2147483648, F1: v3}
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            if false {
                                v23 := []int{-1, fi17}
                                switch fi17 {
                                case 779:
                                    v4 = (v4 / 0.1)
                                    v4 += v4
                                    var v24 uint64 = genFn0(549, false, -2147483648)
                                    _ = v24
                                case -1:
                                    v23 = append(v23, fi17)
                                default:
                                    v25 := GenS1{F0: -868, F1: -578}
                                    _ = v25
                                }
                                if false {
                                    var v26 float32 = (1.5 + ((3.14159 - ((1.5 + 0.1) - 1e10)) + (0.1 / 1e10)))
                                    _ = v26
                                } else {
                                    var v27 int = 127
                                    _ = v27
                                }
                                _ = v23
                            }
                        } else {
                            v28 := GenS3{F0: 449, F1: 42, F2: false}
                            var v29 uint16 = 255
                            v3 -= (v3 >> (((v1 % ((v1) | 1))) & 63))
                            var v30 uint16 = v29
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        var v31 uint64 = 42
                        var v32 uint = (v1 & (v1 + 640))
                        _ = fi17
                        _ = v31
                        _ = v32
                    }
                    v33 := []int{(fi16 / ((fi16) | 1)), fi16}
                    _ = fi16
                    _ = v33
                }
            }
            v34 := GenS2{F0: "a"}
            v1 = v1
            _ = v34
        } else {
            v3 -= -484
        }
    }
    fmt.Println(v4, v3, v1)
    var v35 uint64 = 891
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
