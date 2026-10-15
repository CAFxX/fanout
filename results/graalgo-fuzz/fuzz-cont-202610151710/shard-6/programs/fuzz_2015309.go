package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int = -77
    var v2 uint16 = 255
    var v3 float64 = 1e10
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: false}
        v1 = (v1 % ((v1) | 1))
        _ = fi4
        _ = v5
    }
    v6 := GenS2{F0: 2, F1: 0, F2: 2}
    var v7 uint16 = v2
    v3 = 3.14159
    v8 := GenS1{F0: false}
    if (true && (!(38 <= 1))) {
        for fi9 := 0; fi9 < 4; fi9++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(v3, v7, v1)
            }
            _ = fi9
        }
        if ((239 != 127) || (1 < -32768)) {
            for fi10 := 0; fi10 < 2; fi10++ {
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        v11 := GenS3{F0: (v7 - (^(v2))), F1: 608}
                        _ = v11
                    case "hello":
                        v12 := []int{(^(v1)), (fi10 | fi10)}
                        _ = v12
                    case "foo bar":
                        for fi13 := 0; fi13 < 13; fi13++ {
                            if false {
                                v1 += fi13
                                var v14 uint32 = 42
                                _ = v14
                            }
                            switch fi13 {
                            case 2147483647:
                                fmt.Println(v7, fi13)
                            case 0:
                                v15 := GenS4{F0: fi10, F1: 65535}
                                _ = v15
                            case 9223372036854775807:
                                if (1 == 42) {
                                    if (!(!(65535 > 0))) {
                                        var v16 uint16 = (v2 ^ 32767)
                                        v3 = (1e10 * (1e10 * ((float64((0.0 - (1e10 - (0.0 * 3.14159)))) - float64(float32(float64((-2.25 - 1e10))))) - v3)))
                                        _ = v16
                                    }
                                    v3 += 1.5
                                    v17 := []int{}
                                    _ = v17
                                }
                            default:
                                v3 = v3
                            }
                            v2 -= 65535
                            v18 := []int{int(v2), 64, fi10, 42}
                            _ = fi13
                            _ = v18
                        }
                    }
                    if true {
                        var v19 uint64 = 42
                        var v20 uint = 255
                        v21 := []int{(fi10 | (^(v1)))}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS4{F0: (fi10 / ((-1) | 1)), F1: 42}
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := []int{623, v1, (^(fi23))}
                        if ((0 > 100) && (!true)) {
                            v25 := GenS4{F0: (^(fi23)), F1: 0}
                            var v26 uint8 = 6
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS3{F0: (v7 | 0), F1: 990}
                            v28 := GenS2{F0: 18446744073709551615, F1: 42, F2: (v7 ^ (v7 << ((2) & 63)))}
                            v2 = 290
                            _ = v27
                            _ = v28
                        }
                        v3 = 0.0
                        v24 = append(v24, v1)
                        switch ("" + ("hello" + "")) {
                        case "":
                            v24 = append(v24, fi10)
                        case "a":
                            if (v3 == v3) {
                                var v29 bool = (float64((float32(v3) - 3.14159)) <= 0.0)
                                if true {
                                    v30 := GenS3{F0: (v2 % ((v2) | 1)), F1: 537}
                                    var v31 float32 = 0.0
                                    _ = v30
                                    _ = v31
                                } else {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                }
                                if (2 > 88) {
                                    if ("foo bar" != (("" + "hello") + "hello")) {
                                        var v32 uint32 = 1
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: ((^(fi23)) != fi23)}
                                        _ = v33
                                    }
                                    _ = len(v24)
                                    if v29 {
                                        fmt.Println(v29, fi23, v7)
                                    }
                                } else {
                                    v34 := GenS5{F0: 2147483647, F1: -5}
                                    fmt.Println(v7)
                                    v35 := GenS1{F0: ((255 == 65535) || (346 > 42))}
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS4{F0: -1, F1: 144}
                                _ = v29
                                _ = v36
                            } else {
                                v7 = v7
                                var v37 int32 = -2147483648
                                fmt.Println(v1, v7, v3)
                                _ = v37
                            }
                            fmt.Println(v1)
                            var v38 string = "foo bar"
                            _ = v38
                        case "hello":
                            if (398 > 0) {
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                            } else {
                                fmt.Println(v7, fi10, v1)
                                v2 -= v2
                            }
                        default:
                            fmt.Println(v2, fi23, fi23)
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi10
            }
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 float32 = float32((v3 / v3))
                _ = fi39
                _ = v40
            }
        }
    } else {
        fmt.Println(v3)
        if false {
            var v41 float32 = 0.0
            _ = v41
        } else {
            v7 = v2
            var v42 uint64 = 4294967295
            _ = v42
        }
        var v43 int16 = -32768
        _ = v43
    }
    fmt.Println(v1, v3)
    v3 -= 1e10
    for fi44 := 0; fi44 < 12; fi44++ {
        v2 = 32767
        _ = fi44
    }
    v3 += 3.14159
    if (3.14159 < ((float64(1.5) / 0.1) - (-2.25 * float64(1e10)))) {
        for fi45 := 0; fi45 < 19; fi45++ {
            fmt.Println(v7, fi45)
            v46 := []int{}
            _ = fi45
            _ = v46
        }
    }
    switch v1 {
    case 2147483647:
        if (-128 <= 10) {
            v47 := GenS6{F0: 0, F1: 1}
            fmt.Println(v1, v7)
            var v48 int64 = 895
            _ = v47
            _ = v48
        } else {
            if false {
                var v49 int32 = int32(v1)
                var v50 string = "\n\t"
                _ = v49
                _ = v50
            }
        }
    case 0:
        fmt.Println(v3)
    }
    v7 -= 189
    var v51 uint64 = 4294967295
    for fi52 := 0; fi52 < 17; fi52++ {
        var v53 int16 = 1
        _ = fi52
        _ = v53
    }
    var v54 int = (v1 % (((v1 * (v1 + v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v54
    _ = v6
    _ = v7
    _ = v8
}
