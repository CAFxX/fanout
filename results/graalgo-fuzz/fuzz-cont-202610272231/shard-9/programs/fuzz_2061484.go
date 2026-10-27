package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 int16) int {
    var v4 int32 = 573
    v5 := []int{4294967295, -645, 128, 32767}
    v4 = v4
    _ = p0
    _ = v4
    _ = v5
    return 296
}
func genFn1() int {
    var v6 uint32 = 0
    v7 := GenS1{F0: v6, F1: true}
    v8 := GenS1{F0: (v6 >> ((uint(v6)) & 63)), F1: false}
    var v9 uint16 = 42
    v10 := []int{42, -749, 1, 0}
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := []uint32{65535}
    var v2 string = "hello"
    v3 := make(map[uint32]string)
    if false {
        var v11 bool = (!(!((true || false) && false)))
        _ = v11
    } else {
        v2 = "hello"
    }
    _ = len(v3)
    v12 := GenS1{F0: 44, F1: true}
    var v13 float64 = (1.5 - (1e10 - float64(3.14159)))
    var v14 float32 = 3.14159
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS1{F0: 199, F1: (true && false)}
            _ = fi15
            _ = v16
        }
        var v17 float64 = (1e10 * v13)
        v17 -= (0.1 + ((v13 * (v13 / ((0.1 / ((0.1 / v13) - v13)) * v13))) * ((0.1 + -2.25) + (0.1 / (0.0 * v17)))))
        _ = v17
    } else {
        var v18 uint8 = 0
        _ = v18
    }
    if (true && true) {
        v14 = (float32((v13 * (v13 - v13))) * ((float32(-2.25) - -2.25) + float32(v13)))
        var v19 int8 = 100
        _ = v19
    } else {
        if true {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint64 = 2
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 int8 = int8(fi22)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi22
                    _ = v23
                }
                if false {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v24 bool = ((2 < 114) && (!(!(!(false && false)))))
                        _ = v24
                    } else {
                        v25 := GenS1{F0: 425, F1: (!(!true))}
                        v26 := GenS2{F0: 127, F1: -1}
                        for fi27 := 0; fi27 < 12; fi27++ {
                            var v28 float32 = v14
                            v21 = (v21 & (v21 % (((v21 * (v21 & v21))) | 1)))
                            _ = fi27
                            _ = v28
                        }
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = len(v1)
            switch ("hello" + "a") {
            case "hello":
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 uint8 = 2
                    var v31 uint = 758
                    if (213 >= -1) {
                        var v32 uint8 = (v30 - (v30 - (v30 >> (((v31 >> (((v31 % ((v31) | 1))) & 63))) & 63))))
                        _ = v32
                    } else {
                        v1 = append(v1, 65535)
                        if false {
                            v33 := GenS2{F0: 100, F1: 975}
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 3; fi34++ {
                                _ = len(v1)
                                v35 := GenS3{F0: 1, F1: 1}
                                fmt.Println(v30, fi34)
                                _ = fi34
                                _ = v35
                            }
                            switch (v2 + v2) {
                            case "\n\t":
                                var v36 int64 = 9223372036854775807
                                _ = v36
                            case "":
                                var v37 float32 = 3.14159
                                _ = v37
                            case "hello":
                                var v38 bool = false
                                _ = v38
                            }
                        }
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        } else {
            for fi39 := 0; fi39 < 8; fi39++ {
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS2{F0: -83, F1: 1000}
                    v42 := GenS1{F0: 42, F1: (!(((4294967295 <= 2) || true) && ((!true) || (!false))))}
                    var v43 int8 = -1
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                if (!true) {
                    var v44 int16 = 660
                    v45 := GenS1{F0: 0, F1: false}
                    _ = v44
                    _ = v45
                }
                _ = fi39
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v46 float32 = v14
    var v47 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
