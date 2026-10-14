package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int16 }
func genFn0() int8 {
    v5 := GenS1{F0: (0.0 != float64(1e10)), F1: 0}
    v6 := GenS2{F0: 81}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint16 = 42
        if false {
            v9 := GenS1{F0: true, F1: fi7}
            var v10 bool = false
            _ = v10
            _ = v9
        } else {
            if false {
                var v11 int8 = 127
                _ = v11
            } else {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v8 = (v8 % (((v8 + v8)) | 1))
                    _ = fi12
                }
            }
            v13 := []int{}
            _ = len(v13)
            _ = v13
        }
        if ((1 < 42) || (true || (32767 <= 1))) {
            if true {
                switch 32768 {
                case -895:
                    var v14 uint8 = 1
                    _ = v14
                case 32768:
                    var v15 uint16 = 2
                    _ = v15
                case 0:
                    v16 := GenS3{F0: 0, F1: -433}
                    _ = v16
                default:
                    if false {
                        var v17 int64 = 16
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v19 := GenS4{F0: 42}
                            if (false && (true || (true || ((fi18 >> ((255) & 63)) <= fi7)))) {
                                fmt.Println(v17, v8, v8)
                                var v20 uint32 = 1
                                _ = v20
                            }
                            v8 -= (v8 >> ((255) & 63))
                            v21 := []int{}
                            if (!true) {
                                var v22 uint64 = 0
                                fmt.Println(fi18)
                                _ = v22
                            } else {
                                var v23 float32 = 0.0
                                v24 := GenS5{F0: 2, F1: 717}
                                var v25 uint8 = 1
                                fmt.Println(fi7, v23, fi18)
                                if false {
                                    if false {
                                        v26 := GenS3{F0: 18446744073709551615, F1: 42}
                                        fmt.Println(fi18, v23, v23)
                                        var v27 string = "foo bar"
                                        v17 = (^(v17))
                                        v23 = 0.1
                                        _ = v26
                                        _ = v27
                                    } else {
                                        _ = len(v21)
                                        v28 := GenS2{F0: 2}
                                        _ = v28
                                    }
                                    if (!false) {
                                        v23 = (v23 + (v23 - ((0.0 / 0.1) - v23)))
                                        if ((0.1 / float32(0.1)) == v23) {
                                            var v29 float64 = float64(float32(0.1))
                                            var v30 int32 = int32(v8)
                                            if (-1 != (-(v30))) {
                                                v23 -= v23
                                                var v31 uint32 = 2147483648
                                                fmt.Println(v23, v8, v17)
                                                v32 := GenS6{F0: 824}
                                                v21 = append(v21, (-(fi7)))
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                    v33 := GenS5{F0: 2, F1: v8}
                                    _ = v33
                                }
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(v8, fi7)
                            _ = fi18
                            _ = v19
                            _ = v21
                        }
                        _ = v17
                    }
                }
                fmt.Println(v8, fi7)
            } else {
                var v34 int8 = 0
                var v35 float64 = float64(float32((0.1 * 1.5)))
                v34 = (v34 | v34)
                _ = v34
                _ = v35
            }
            v8 = 42
            var v36 uint16 = (v8 & 1)
            _ = v36
        }
        _ = fi7
        _ = v8
    }
    var v37 int8 = 43
    _ = v37
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []string{}
    var v2 int16 = 954
    v3 := []int{4294967295, -9223372036854775808, 127}
    var v4 int32 = 42
    _ = len(v1)
    if len(v3) > 0 {
        v3[0] = -1
    }
    v38 := GenS2{F0: 42}
    var v39 int32 = (v4 % (((v4 + -743)) | 1))
    var v40 float64 = 3.14159
    v41 := GenS2{F0: 2}
    if true {
        var v42 int = -9223372036854775808
        _ = v42
    } else {
        v2 = -46
    }
    fmt.Println(v2, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
