package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 string; F2 int }

func main() {
    var v1 int64 = -1
    var v2 int64 = 1
    v3 := make(map[int]bool)
    var v4 uint8 = 62
    var v5 string = "a"
    if (true || true) {
        v6 := GenS1{F0: v1, F1: 621, F2: (v2 & (v1 - (v1 + 0)))}
        switch 705 {
        case 622:
            for fi7 := 0; fi7 < 16; fi7++ {
                v4 = (^(v4))
                v1 -= (v1 / ((1) | 1))
                v1 -= (v1 & v1)
                _ = fi7
            }
        case -1:
            v8 := []int{-2147483648, 32768, 32767}
            _ = v8
        }
        _ = v6
    } else {
        v9 := []int{}
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        if true {
            if (728 == 677) {
                v4 = (v4 * 255)
                var v10 uint8 = 210
                switch 9223372036854775807 {
                case 127:
                    switch 32767 {
                    case 94:
                        v1 = v1
                    case 2147483647:
                        if (0.0 != float32(3.14159)) {
                            var v11 uint8 = v10
                            switch 32768 {
                            case -2147483648:
                                switch 0 {
                                case 32768:
                                    for fi12 := 0; fi12 < 3; fi12++ {
                                        var v13 uint8 = (v11 / ((v10) | 1))
                                        v14 := GenS2{F0: 273}
                                        _ = fi12
                                        _ = v13
                                        _ = v14
                                    }
                                default:
                                    if (-477 >= 867) {
                                        for fi15 := 0; fi15 < 8; fi15++ {
                                            if true {
                                                switch v5 {
                                                case "":
                                                    v16 := GenS2{F0: 749}
                                                    _ = v16
                                                default:
                                                    var v17 int16 = 0
                                                    _ = v17
                                                }
                                                v5 = ""
                                            } else {
                                                fmt.Println(v2, v11)
                                                v11 = v4
                                                v4 = 243
                                            }
                                            _ = fi15
                                        }
                                    } else {
                                        v18 := []int{128, 1}
                                        for fi19 := 0; fi19 < 7; fi19++ {
                                            v20 := GenS3{F0: v11}
                                            _ = fi19
                                            _ = v20
                                        }
                                        _ = v18
                                    }
                                }
                            }
                            var v21 int = 42
                            _ = v11
                            _ = v21
                        }
                    case 774:
                        var v22 int16 = 1
                        _ = v22
                    }
                case 0:
                    if false {
                        var v23 int8 = 50
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 int32 = 0
                            v10 = v4
                            for fi26 := 0; fi26 < 10; fi26++ {
                                var v27 uint = uint(v10)
                                v28 := GenS3{F0: v4}
                                v23 = 1
                                var v29 uint64 = uint64(v27)
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = fi24
                            _ = v25
                        }
                        switch -26 {
                        case -260:
                            fmt.Println(v5, v23, v5)
                        case 60:
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v2 -= 0
                                _ = fi30
                            }
                        }
                        _ = v23
                    } else {
                        if true {
                            var v31 int32 = -278
                            _ = v31
                        }
                    }
                case -254:
                    fmt.Println(v10, v1, v5)
                }
                _ = v10
            }
            var v32 uint16 = uint16(v4)
            _ = v32
        } else {
            for fi33 := 0; fi33 < 12; fi33++ {
                if (true || (-9223372036854775808 == int64(fi33))) {
                    v34 := GenS3{F0: 12}
                    v4 += v4
                    v5 = v5
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 19; fi35++ {
                        var v36 uint32 = 2
                        _ = fi35
                        _ = v36
                    }
                    v4 += (v4 << ((4294967295) & 63))
                    v37 := GenS4{F0: 923, F1: 998}
                    _ = v37
                }
                for fi38 := 0; fi38 < 14; fi38++ {
                    var v39 int8 = -1
                    v2 = v1
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v2, v1)
                _ = fi33
            }
        }
    default:
        var v40 int32 = 0
        _ = v40
    }
    if ((^(v1)) >= v2) {
        v41 := []int{741}
        var v42 uint32 = 2147483647
        v43 := GenS5{F0: (((v5 + "hello") + v5) + "foo bar"), F1: 42}
        _ = v41
        _ = v42
        _ = v43
    } else {
        var v44 int64 = 1
        _ = v44
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        var v46 int8 = -6
        switch 127 {
        case 42:
            v47 := GenS6{F0: int8(fi45)}
            _ = v47
        case 32767:
            v48 := GenS7{F0: (v4 * (v4 % (((v4 ^ v4)) | 1))), F1: (("hello" + "") + (v5 + "foo bar")), F2: fi45}
            _ = v48
        case 32768:
            v4 = v4
        }
        v3[994] = (!true)
        _ = fi45
        _ = v46
    }
    var v49 float32 = (1.5 / 1.5)
    var v50 int = -557
    var v51 uint8 = (v4 % ((v4) | 1))
    fmt.Println(v1, v50, v49)
    v52 := []int{}
    v50 += -1
    if true {
        fmt.Println(v49, v49)
    } else {
        _ = len(v52)
    }
    if ((v4 + v51) > 197) {
        for fi53 := 0; fi53 < 16; fi53++ {
            for fi54 := 0; fi54 < 18; fi54++ {
                var v55 int16 = -512
                var v56 int = 31
                if false {
                    _ = len(v3)
                    v56 = v56
                    switch v50 {
                    case 2147483647:
                        v57 := GenS1{F0: (v2 % (((v2 ^ (-(v1)))) | 1)), F1: 42, F2: 42}
                        _ = v57
                    case 32767:
                        var v58 int32 = 42
                        _ = v58
                    }
                }
                _ = fi54
                _ = v55
                _ = v56
            }
            var v59 uint8 = 223
            var v60 int32 = -2147483648
            _ = fi53
            _ = v59
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
