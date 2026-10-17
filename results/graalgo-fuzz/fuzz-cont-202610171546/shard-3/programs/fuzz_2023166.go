package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8 }
func genFn0() int64 {
    v4 := []int{2147483647}
    if len(v4) > 0 {
        v4[0] = -610
    }
    _ = v4
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 int16 = 261
    v3 := make(map[int]uint32)
    var v5 int32 = v1
    if true {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 8; fi7++ {
                v2 += v2
                _ = fi7
            }
            v1 = (v5 ^ (v1 % (((v1 ^ (^(v1)))) | 1)))
            _ = fi6
        }
        fmt.Println(v1, v2, v1)
    } else {
        if (42 != 40) {
            for fi8 := 0; fi8 < 9; fi8++ {
                switch -2147483648 {
                case -869:
                    var v9 string = "foo bar"
                    _ = v9
                case 127:
                    v5 = 685
                case 1:
                    v10 := GenS1{F0: (v1 / (((v5 ^ v1)) | 1)), F1: 1}
                    _ = v10
                }
                if false {
                    v1 = (v1 / ((-59) | 1))
                } else {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        if true {
                            v12 := GenS2{F0: 2, F1: 1}
                            fmt.Println(fi8, fi8, fi8)
                            if (v2 >= 1000) {
                                if false {
                                    var v13 int64 = 9223372036854775807
                                    v2 += (v2 + 239)
                                    v14 := []int{9223372036854775807}
                                    _ = v13
                                    _ = v14
                                } else {
                                    if true {
                                        v15 := GenS3{F0: 1}
                                        _ = v15
                                    }
                                    switch fi8 {
                                    case 42:
                                        switch ("" + ("" + "hello")) {
                                        case "a":
                                            v16 := []int{2147483647, (^(fi11))}
                                            _ = v16
                                        case "foo bar":
                                            var v17 uint8 = 165
                                            _ = v17
                                        case "\n\t":
                                            v1 -= 2147483647
                                        }
                                    default:
                                        fmt.Println(fi11, v5)
                                    }
                                }
                                var v18 int64 = -391
                                _ = v18
                            } else {
                                var v19 float32 = 0.0
                                _ = v19
                            }
                            v20 := []int{fi8, fi11, (fi11 & (fi11 ^ (fi11 << ((0) & 63))))}
                            var v21 uint8 = 128
                            _ = v12
                            _ = v20
                            _ = v21
                        }
                        var v22 uint16 = 2
                        _ = fi11
                        _ = v22
                    }
                }
                _ = fi8
            }
            var v23 int64 = 9223372036854775807
            _ = v23
        }
        if true {
            v24 := []int{}
            var v25 int16 = (-(v2))
            _ = v24
            _ = v25
        }
        var v26 int8 = -93
        _ = v26
    }
    fmt.Println(v5, v2, v5)
    fmt.Println(v5)
    v1 = (^(v1))
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS2{F0: 159, F1: 255}
        v29 := GenS2{F0: 42, F1: 378}
        v5 = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: -40}
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int8 = -128
            if (!(float64(((1.5 / (1.5 + 1.5)) - float32(-2.25))) > 0.0)) {
                var v34 int = 32767
                v1 += (v1 / ((v1) | 1))
                _ = v34
            } else {
                var v35 bool = false
                _ = v35
            }
            v2 -= int16(fi31)
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v5
}
