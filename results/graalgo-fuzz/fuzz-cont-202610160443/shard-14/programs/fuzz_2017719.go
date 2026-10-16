package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int32 }
func genFn0() int32 {
    var v7 uint16 = 255
    fmt.Println(v7)
    v8 := GenS1{F0: 104}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = (!((1 == 42) && true))
    var v2 string = ("hello" + ("" + ("\n\t" + "hello")))
    var v3 float32 = 1e10
    var v4 uint16 = 0
    v5 := make(map[int]int)
    var v6 uint16 = v4
    var v9 float32 = (v3 + v3)
    var v10 string = (((v2 + (v2 + v2)) + v2) + "")
    var v11 float32 = (3.14159 + v9)
    v12 := GenS1{F0: -128}
    switch 4294967295 {
    case -2147483648:
        if v1 {
            var v13 bool = (0.1 < float64(1.5))
            var v14 int = -746
            _ = v13
            _ = v14
        } else {
            var v15 uint = 42
            _ = v15
        }
    default:
        var v16 int32 = 245
        _ = v16
    }
    if false {
        v17 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 2}
        if true {
            for fi18 := 0; fi18 < 3; fi18++ {
                if (0 < -32768) {
                    if v1 {
                        v10 = "hello"
                        if (-2147483648 < 42) {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                _ = len(v5)
                                var v20 int = fi18
                                v21 := GenS3{F0: 1, F1: 2147483647}
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            switch ("a" + v2) {
                            case "\n\t":
                                v22 := []int{2147483647, -1, fi18}
                                _ = v22
                            default:
                                v23 := GenS4{F0: (fi18 / ((fi18) | 1))}
                                _ = v23
                            }
                            v24 := GenS5{F0: ("" + (v10 + v2)), F1: 257}
                            _ = v24
                        } else {
                            var v25 uint32 = 2147483647
                            _ = v25
                        }
                        if (42 > 0) {
                            var v26 int = (fi18 | fi18)
                            _ = v26
                        }
                    } else {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := GenS1{F0: -128}
                            v4 -= v6
                            _ = fi27
                            _ = v28
                        }
                        if v1 {
                            fmt.Println(v4, v11, v4)
                        }
                        delete(v5, 0)
                    }
                    for fi29 := 0; fi29 < 19; fi29++ {
                        if (1 < 2147483648) {
                            v1 = true
                            switch fi29 {
                            case -995:
                                v1 = v1
                            case 32767:
                                v30 := GenS3{F0: 57, F1: 2147483647}
                                _ = v30
                            }
                            fmt.Println(v9, v11, v3)
                        }
                        var v31 uint16 = v4
                        v32 := GenS6{F0: 2, F1: 0}
                        _ = fi29
                        _ = v31
                        _ = v32
                    }
                }
                switch v2 {
                case "":
                    fmt.Println(v2)
                case "a":
                    if v1 {
                        v5[154] = fi18
                        v33 := GenS2{F0: 1, F1: 1, F2: 2147483648}
                        if false {
                            switch int(v4) {
                            case 479:
                                if true {
                                    if v1 {
                                        fmt.Println(v9, v1, v4)
                                    } else {
                                        v10 = ("hello" + v10)
                                        var v34 int32 = 558
                                        _ = v34
                                    }
                                    v9 -= ((1.5 * (v11 * v3)) * float32(3.14159))
                                } else {
                                    v3 = v11
                                    var v35 float64 = float64((v3 - (float32(-2.25) - v9)))
                                    v1 = v1
                                    _ = v35
                                }
                            case 128:
                                var v36 int8 = 45
                                _ = v36
                            }
                        }
                        _ = v33
                    }
                default:
                    var v37 int64 = -1
                    _ = v37
                }
                fmt.Println(v4, v4, v6)
                _ = fi18
            }
        }
        fmt.Println(v4)
        _ = v17
    } else {
        v9 = float32(0.1)
        fmt.Println(v4, v11)
        var v38 float32 = float32(1.5)
        _ = v38
    }
    var v39 int = 32768
    var v40 int64 = 566
    switch ((("hello" + v10) + v10) + "hello") {
    case "":
        var v41 int = (v39 / (((v39 * 2147483647)) | 1))
        _ = v41
    case "foo bar":
        v5[v39] = 42
    case "a":
        if (1.5 != (3.14159 + ((0.0 / (float64(v3) * float64(-2.25))) + 3.14159))) {
            var v42 int32 = 1
            _ = v42
        } else {
            switch v2 {
            case "\n\t":
                for fi43 := 0; fi43 < 14; fi43++ {
                    if v1 {
                        v39 += v39
                        for fi44 := 0; fi44 < 18; fi44++ {
                            v39 = v39
                            v45 := GenS4{F0: v39}
                            _ = fi44
                            _ = v45
                        }
                    }
                    _ = fi43
                }
            }
            v46 := v5[-2147483648]
            _ = v46
        }
    }
    v1 = v1
    if true {
        v47 := []int{v39, -1}
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS7{F0: -1}
            _ = len(v5)
            _ = fi48
            _ = v49
        }
        var v50 int = int(v4)
        _ = v47
        _ = v50
    }
    v40 = v40
    if (51 == 1000) {
        v51 := GenS5{F0: v2, F1: int16(v40)}
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
