package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool; F1 int }
func genFn0() int32 {
    if false {
        switch "" {
        case "a":
            var v6 int8 = 42
            _ = v6
        case "\n\t":
            var v7 float64 = (float64(3.14159) * (float64(float32(-2.25)) * float64(-2.25)))
            _ = v7
        }
        v8 := []int{32767, 4294967295, -681, 32767}
        _ = v8
    }
    if false {
        var v9 uint = 42
        v10 := GenS1{F0: (v9 << (((v9 / (((v9 / ((255) | 1))) | 1))) & 63)), F1: 65535, F2: 198}
        if (650 != 65535) {
            if (32767 <= 1) {
                var v11 uint16 = 705
                _ = v11
            } else {
                v9 -= (^(v9))
            }
            var v12 uint16 = 2
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint32 = 875
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            var v15 int64 = -362
            var v16 int64 = (-(v15))
            fmt.Println(v16, v16)
            _ = v15
            _ = v16
        }
        _ = v10
        _ = v9
    } else {
        var v17 int32 = 2147483647
        var v18 uint32 = 2147483647
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(v17)
            v17 -= -236
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    switch ("a" + "a") {
    case "foo bar":
        var v20 uint8 = 238
        _ = v20
    case "\n\t":
        if (((0.1 * -2.25) == 3.14159) && (307 != 0)) {
            var v21 int8 = 42
            _ = v21
        } else {
            v22 := []int{-912, 42}
            _ = v22
        }
    default:
        var v23 uint8 = 0
        _ = v23
    }
    var v24 int16 = 0
    _ = v24
    return 42
}
func genFn1() uint {
    var v25 int16 = 0
    fmt.Println(v25)
    _ = v25
    return 18446744073709551615
}

func main() {
    var v1 int32 = 1
    v2 := []int{609, 9223372036854775807, -2147483648}
    var v3 uint = 452
    var v4 uint64 = 362
    v5 := make(map[uint32]bool)
    switch 127 {
    case 0:
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                v1 += v1
                if (127 <= 1) {
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                } else {
                    v1 -= 2147483647
                    if true {
                        v28 := GenS1{F0: genFn1(), F1: 680, F2: 127}
                        _ = v28
                    } else {
                        delete(v5, 2)
                        v29 := GenS2{F0: 127, F1: 255, F2: 1}
                        if ((true || (0 < -2147483648)) && true) {
                            var v30 uint = 2
                            if (1 != -761) {
                                v31 := GenS3{F0: 42, F1: (("foo bar" + "hello") + "foo bar"), F2: -9223372036854775808}
                                _ = v31
                            }
                            _ = v30
                        }
                        _ = v29
                    }
                }
                if (float64(1e10) > 0.0) {
                    v32 := v5[969]
                    if (-32768 == -32) {
                        if (false && (!(v32 || v32))) {
                            var v33 int8 = -125
                            var v34 uint64 = v4
                            v35 := GenS4{F0: (v33 | v33), F1: -32768}
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS5{F0: v32, F1: (fi26 * 99)}
                            if v32 {
                                if (!v32) {
                                    var v37 uint8 = 42
                                    v32 = v32
                                    v5[0] = (!(!v32))
                                    _ = v37
                                } else {
                                    fmt.Println(fi26, fi27)
                                }
                            } else {
                                v3 = v3
                            }
                            _ = v36
                        }
                        v32 = ((0 == 42) && (18446744073709551615 != (v4 / ((4294967295) | 1))))
                        v32 = (v4 == (v4 * (v4 >> ((genFn1()) & 63))))
                    }
                    v1 = v1
                    _ = v32
                } else {
                    var v38 bool = false
                    _ = v38
                }
                _ = fi27
            }
            _ = fi26
        }
    case 1:
        v39 := GenS4{F0: int8(v3), F1: 1000}
        _ = v39
    case 216:
        _ = len(v2)
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        v3 = genFn1()
        if false {
            v4 += (v4 ^ (v4 >> ((42) & 63)))
        } else {
            var v41 bool = false
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 int8 = 42
                v3 = (v3 / (((^(v3))) | 1))
                v2 = append(v2, fi42)
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        fmt.Println(v3, fi40)
        _ = fi40
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = v4
    }
    v4 -= (v4 * v4)
    fmt.Println(v3, v1)
    v44 := GenS5{F0: false, F1: 0}
    if true {
        var v45 uint = 1
        var v46 float32 = 3.14159
        _ = v45
        _ = v46
    } else {
        switch ("a" + "\n\t") {
        case "\n\t":
            var v47 int8 = -1
            _ = v47
        case "a":
            var v48 int = 127
            _ = v48
        default:
            var v49 bool = true
            _ = v49
        }
        fmt.Println(v4, v1, v4)
        for fi50 := 0; fi50 < 2; fi50++ {
            switch fi50 {
            case 42:
                var v51 uint64 = 1
                _ = v51
            }
            v4 += v4
            _ = fi50
        }
    }
    v1 += v1
    v52 := GenS2{F0: 2, F1: 128, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v52
}
