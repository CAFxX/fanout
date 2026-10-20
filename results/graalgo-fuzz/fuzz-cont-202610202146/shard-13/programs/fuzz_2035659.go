package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint = 42
    var v2 uint = v1
    var v3 float32 = 1.5
    v3 = 0.0
    v3 -= v3
    if true {
        if (!false) {
            v4 := GenS1{F0: 0, F1: 169, F2: 2147483647}
            var v5 uint = v1
            _ = v4
            _ = v5
        } else {
            v6 := GenS1{F0: 127, F1: 32768, F2: -1}
            var v7 string = "\n\t"
            var v8 uint32 = 0
            _ = v6
            _ = v7
            _ = v8
        }
        if (!false) {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9, v2)
                for fi10 := 0; fi10 < 5; fi10++ {
                    if false {
                        var v11 float32 = v3
                        fmt.Println(v3, fi9, fi9)
                        _ = v11
                    } else {
                        v3 += ((v3 - (v3 + v3)) + 1.5)
                        switch 4294967295 {
                        case 32767:
                            if (false && true) {
                                if ("hello" == "hello") {
                                    fmt.Println(fi10, v3)
                                    if (0 <= 690) {
                                        v12 := []int{fi9, 0}
                                        var v13 int = fi9
                                        v13 += fi9
                                        _ = len(v12)
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v14 := GenS2{F0: 2147483647, F1: 0, F2: -1}
                                        var v15 bool = true
                                        if (-86 >= 0) {
                                            var v16 string = "a"
                                            fmt.Println(v2, fi9)
                                            v3 -= float32((((0.1 / 1e10) - 3.14159) - -2.25))
                                            _ = v16
                                        }
                                        v2 -= (v1 / ((v1) | 1))
                                        var v17 uint64 = 0
                                        if v15 {
                                            fmt.Println(v17, fi10, fi9)
                                            v2 -= 798
                                        }
                                        _ = v14
                                        _ = v15
                                        _ = v17
                                    }
                                }
                            } else {
                                fmt.Println(v2)
                                if (false && false) {
                                    v1 += 0
                                    fmt.Println(fi10)
                                    fmt.Println(fi9, fi10)
                                    var v18 int = fi10
                                    _ = v18
                                }
                                var v19 int8 = 42
                                _ = v19
                            }
                        case -500:
                            var v20 uint8 = 160
                            _ = v20
                        default:
                            var v21 int = (fi10 + fi10)
                            _ = v21
                        }
                        v3 = float32(float64(v3))
                    }
                    var v22 bool = (-32768 == 309)
                    var v23 int16 = -738
                    _ = fi10
                    _ = v22
                    _ = v23
                }
                _ = fi9
            }
        } else {
            var v24 float64 = 3.14159
            _ = v24
        }
        v25 := GenS2{F0: 537, F1: 697, F2: -698}
        _ = v25
    } else {
        v26 := []int{-266}
        v3 -= 0.1
        _ = v26
    }
    v27 := []int{-9223372036854775808, -640, -330}
    var v28 uint = (v1 | v2)
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int = 2147483647
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 151, F1: -2147483648, F2: 9223372036854775807}
    v32 := GenS1{F0: 127, F1: -992, F2: 1}
    var v33 int = 2147483647
    var v34 uint16 = 1
    var v35 float64 = -2.25
    v36 := GenS1{F0: 207, F1: 0, F2: -576}
    if false {
        v37 := GenS4{F0: ("" + (("a" + "hello") + "\n\t"))}
        v27 = append(v27, (v33 & v33))
        _ = len(v27)
        _ = v37
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
