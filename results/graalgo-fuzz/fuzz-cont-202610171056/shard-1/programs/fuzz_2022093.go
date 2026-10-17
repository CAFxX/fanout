package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int32) int32 {
    if (!((p0 / (((-(p0))) | 1)) != p0)) {
        var v4 string = "\n\t"
        v4 = v4
        var v5 uint8 = 120
        _ = v4
        _ = v5
    }
    p0 += p0
    _ = p0
    return (p0 | p0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 3.14159
    var v3 string = ("" + "\n\t")
    var v6 int64 = 42
    if false {
        var v7 uint = 65535
        _ = v7
    }
    switch (v3 + ("foo bar" + (v3 + v3))) {
    case "\n\t":
        v8 := GenS1{F0: 255, F1: 2, F2: -1}
        _ = v8
    default:
        v3 = "a"
    }
    if true {
        if true {
            if true {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := GenS2{F0: 42}
                    _ = fi9
                    _ = v10
                }
                v11 := []int{-9223372036854775808, 127, 127, 42}
                v12 := GenS3{F0: v3, F1: 42, F2: 42}
                _ = v11
                _ = v12
            } else {
                switch 2147483647 {
                case -2147483648:
                    switch 2147483647 {
                    case 9223372036854775807:
                        v13 := []int{881, -855, 32767, 0}
                        _ = v13
                    default:
                        var v14 uint8 = 164
                        _ = v14
                    }
                }
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                v3 = (v3 + ((v3 + (((v3 + v3) + v3) + (v3 + (v3 + "a")))) + ""))
                _ = fi15
            }
            v16 := GenS1{F0: 65535, F1: 42, F2: 1}
            _ = v16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 bool = (-1 <= 3)
                fmt.Println(v3)
                var v20 uint32 = 2
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v3, v3, v2)
            _ = fi17
        }
        v21 := GenS1{F0: 623, F1: 635, F2: 67}
        _ = v21
    }
    v2 = v2
    v6 -= -1
    v6 -= 408
    for fi22 := 0; fi22 < 17; fi22++ {
        if true {
            var v23 int16 = 0
            v24 := []int{(fi22 / ((-224) | 1)), fi22, (fi22 * (fi22 ^ (fi22 * fi22)))}
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    if true {
        v25 := GenS3{F0: ("hello" + v3), F1: -49, F2: -1}
        var v26 int32 = 2147483647
        for fi27 := 0; fi27 < 14; fi27++ {
            if true {
                v1[(fi27 + fi27)] = (genFn0(genFn0(v26)) <= (v26 | (v26 / ((v26) | 1))))
            } else {
                for fi28 := 0; fi28 < 17; fi28++ {
                    v6 -= v6
                    var v29 uint64 = 0
                    _ = fi28
                    _ = v29
                }
                if (int8(fi27) >= -128) {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v1[fi27] = (!true)
                        if (1 <= 127) {
                            if false {
                                v2 += v2
                            } else {
                                var v31 uint64 = 963
                                if (v6 == 0) {
                                    if false {
                                        v3 = v3
                                        v32 := GenS3{F0: ((v3 + ("a" + v3)) + v3), F1: -128, F2: 42}
                                        _ = v32
                                    } else {
                                        v3 = "foo bar"
                                        v33 := GenS4{F0: -805, F1: 255, F2: -32768}
                                        v34 := GenS5{F0: 42}
                                        _ = v33
                                        _ = v34
                                    }
                                    var v35 float32 = 3.14159
                                    _ = v35
                                } else {
                                    var v36 int16 = 42
                                    v3 = ("\n\t" + v3)
                                    _ = v36
                                }
                                if (v6 < v6) {
                                    var v37 uint32 = 538
                                    _ = v37
                                }
                                _ = v31
                            }
                            v38 := GenS6{F0: 0, F1: 0}
                            _ = v38
                        }
                        _ = len(v1)
                        _ = fi30
                    }
                    v39 := GenS7{F0: 42}
                    for fi40 := 0; fi40 < 19; fi40++ {
                        if true {
                            fmt.Println(fi27)
                            if true {
                                v6 += 9223372036854775807
                            }
                        }
                        v41 := GenS6{F0: 2147483647, F1: 72}
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
            }
            var v42 int8 = 2
            var v43 float32 = 3.14159
            _ = fi27
            _ = v42
            _ = v43
        }
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v3, v2)
    v44 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v6
}
