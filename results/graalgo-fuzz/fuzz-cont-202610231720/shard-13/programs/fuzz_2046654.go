package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{fi4, 218, (^(fi4))}
        if false {
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := GenS1{F0: 0, F1: (fi4 | 127)}
                v8 := GenS1{F0: 42, F1: -9223372036854775808}
                if len(v5) > 0 {
                    v5[0] = fi4
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            fmt.Println(fi4)
        }
        switch (fi4 / (((fi4 + (fi4 | 2147483647))) | 1)) {
        case 2147483647:
            var v9 bool = ((!true) || (!(!false)))
            _ = v9
        case -894:
            v10 := GenS1{F0: 255, F1: (fi4 >> ((4294967295) & 63))}
            _ = v10
        case -182:
            for fi11 := 0; fi11 < 7; fi11++ {
                if false {
                    if (13 < 2147483648) {
                        switch fi11 {
                        case -443:
                            if (!(2147483648 > 65535)) {
                                fmt.Println(fi4, fi11)
                            }
                        case 32768:
                            if true {
                                v5 = append(v5, 1)
                                _ = len(v5)
                                var v12 int = fi11
                                _ = v12
                            }
                        }
                    } else {
                        fmt.Println(fi11, fi11)
                        var v13 uint16 = 997
                        _ = v13
                    }
                } else {
                    fmt.Println(fi4, fi4)
                    switch (fi11 << ((2) & 63)) {
                    case 1:
                        if true {
                            var v14 int16 = 42
                            var v15 uint32 = 65535
                            v14 -= v14
                            v16 := GenS1{F0: 118, F1: (fi4 << ((42) & 63))}
                            if true {
                                v14 = v14
                                switch int(v14) {
                                case 32768:
                                    var v17 int = (fi11 >> ((4294967295) & 63))
                                    _ = v17
                                }
                                v14 -= v14
                                v15 += v15
                            } else {
                                v15 = 814
                                var v18 float64 = float64((1e10 + ((float32((0.0 - -2.25)) / 1e10) + 0.1)))
                                _ = v18
                            }
                            if (true || (1 > 42)) {
                                v15 += (v15 & 856)
                                var v19 uint16 = 774
                                _ = v19
                            }
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            var v20 string = ("\n\t" + "hello")
                            _ = v20
                        }
                    case 0:
                        if true {
                            var v21 bool = false
                            switch fi11 {
                            case -9223372036854775808:
                                var v22 uint64 = 1
                                _ = v22
                            default:
                                if len(v5) > 0 {
                                    v5[0] = (-(fi11))
                                }
                            }
                            v21 = true
                            _ = v21
                        } else {
                            if (false || (false || (!true))) {
                                var v23 float64 = ((3.14159 + 3.14159) - 3.14159)
                                v23 -= (v23 - (v23 + (((((1.5 / v23) - float64(0.0)) * (1.5 + v23)) / 1.0) * (float64(0.0) * float64(-2.25)))))
                                v24 := GenS2{F0: 2147483648, F1: 42, F2: -578}
                                fmt.Println(fi4, v23)
                                _ = v23
                                _ = v24
                            }
                        }
                    case -536:
                        var v25 float32 = ((((-2.25 * 3.14159) - 0.0) / 1.0) + ((3.14159 + ((((0.0 * -2.25) - 0.1) / (0.1 * 0.1)) - -2.25)) / (float32(-2.25) / 0.1)))
                        _ = v25
                    default:
                        v26 := GenS3{F0: 32767, F1: -808, F2: 2}
                        _ = v26
                    }
                    v5 = append(v5, (fi4 + -1))
                }
                if false {
                    var v27 uint8 = 115
                    fmt.Println(fi4)
                    fmt.Println(fi4, fi11, fi4)
                    _ = v27
                }
                if true {
                    if true {
                        v28 := GenS1{F0: 230, F1: fi11}
                        _ = v28
                    }
                }
                _ = fi11
            }
        }
        _ = fi4
        _ = v5
    }
    v29 := []int{-560}
    if len(v29) > 0 {
        v29[0] = 1
    }
    _ = v29
    return -2147483648
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    switch 0 {
    case -968:
        var v30 int16 = 0
        _ = v30
    default:
        v3 -= (v3 + v3)
    }
    v3 = genFn0()
    fmt.Println(v1)
    v31 := GenS4{F0: 378}
    if false {
        v3 -= v3
    } else {
        if true {
            v32 := GenS5{F0: 128}
            _ = v32
        } else {
            fmt.Println(v3, v3)
        }
    }
    fmt.Println(v3, v3)
    v3 -= -2147483648
    var v33 bool = (0 > 771)
    var v34 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
}
