package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int64; F1 uint16 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint64 = 270
        v6 := []int{}
        v5 += v5
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{0, 128}
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 0
    var v3 int = 333
    switch ("hello" + "foo bar") {
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint16 = (v2 & (^(v2)))
            v10 := GenS1{F0: v3, F1: -197}
            fmt.Println(v9, v3)
            _ = fi8
            _ = v10
            _ = v9
        }
    case "foo bar":
        v11 := GenS1{F0: v3, F1: -148}
        _ = v11
    }
    v12 := GenS1{F0: v3, F1: -2147483648}
    if ((v2 + v2) != v2) {
        switch v3 {
        case 4294967295:
            var v13 uint32 = 4294967295
            _ = v13
        case 32768:
            v14 := GenS1{F0: (v3 / ((v3) | 1)), F1: 1}
            _ = v14
        case 2147483647:
            fmt.Println(v3, v1)
        default:
            v15 := []int{(v3 / ((v3) | 1)), int(v1), (v3 + 4294967295)}
            _ = v15
        }
        var v16 int16 = 42
        _ = v16
    }
    switch "a" {
    case "a":
        v17 := GenS1{F0: v3, F1: -513}
        _ = v17
    case "":
        var v18 bool = (true && false)
        _ = v18
    default:
        v19 := []int{v3}
        _ = v19
    }
    v20 := []int{v3, v3}
    for fi21 := 0; fi21 < 13; fi21++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        fmt.Println(v1, v2)
        v22 := GenS2{F0: (-112 > -791)}
        _ = fi21
        _ = v22
    }
    v20 = append(v20, v3)
    var v23 int16 = -1
    v2 = (^(v2))
    var v24 float32 = -2.25
    if len(v20) > 0 {
        v20[0] = 844
    }
    fmt.Println(v2, v2, v23)
    switch (("foo bar" + "") + "a") {
    case "\n\t":
        v25 := GenS3{F0: (^(v2))}
        _ = v25
    case "foo bar":
        _ = len(v20)
    default:
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 int8 = 0
            switch "hello" {
            case "hello":
                if false {
                    _ = len(v20)
                    if true {
                        v28 := GenS4{F0: -1}
                        for fi29 := 0; fi29 < 13; fi29++ {
                            fmt.Println(v23)
                            v30 := GenS1{F0: (fi26 & (-(v3))), F1: 1}
                            v31 := GenS2{F0: false}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        if len(v20) > 0 {
                            v20[0] = (v3 + -2147483648)
                        }
                        _ = v28
                    } else {
                        fmt.Println(fi26, v24)
                        v2 += (v2 + v2)
                        var v32 uint8 = 255
                        _ = v32
                    }
                } else {
                    if true {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        if (2 != 2) {
                            switch "" {
                            case "\n\t":
                                v23 -= genFn0()
                            case "hello":
                                v24 -= (0.1 + v24)
                            default:
                                switch "a" {
                                case "a":
                                    var v33 uint = 255
                                    _ = v33
                                default:
                                    var v34 float64 = 0.0
                                    _ = v34
                                }
                            }
                            for fi35 := 0; fi35 < 5; fi35++ {
                                switch "foo bar" {
                                case "a":
                                    if false {
                                        v23 = v23
                                        fmt.Println(fi35, v1)
                                    } else {
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    }
                                case "foo bar":
                                    if true {
                                        v36 := GenS5{F0: 432, F1: 0, F2: true}
                                        v2 = (v2 & v2)
                                        _ = v36
                                    }
                                case "hello":
                                    v37 := GenS1{F0: -617, F1: 2147483647}
                                    if ((-(fi26)) <= (v3 / (((fi26 - fi26)) | 1))) {
                                        if (("a" + (("hello" + "foo bar") + "foo bar")) != "foo bar") {
                                            var v38 int16 = v23
                                            fmt.Println(v2)
                                            v39 := GenS6{F0: 2147483647}
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    switch v3 {
                                    case 32767:
                                        v2 += (v2 << ((4294967295) & 63))
                                    case 2147483647:
                                        var v40 string = (("" + "\n\t") + "\n\t")
                                        _ = v40
                                    }
                                    v24 += v24
                                    _ = v37
                                default:
                                    var v41 uint16 = v2
                                    var v42 int64 = 9223372036854775807
                                    var v43 float64 = (0.0 + (1.5 - 0.0))
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                switch (^(v3)) {
                                case 128:
                                    var v44 uint = 42
                                    _ = v44
                                case -2147483648:
                                    switch 740 {
                                    case 32768:
                                        if (!true) {
                                            var v45 float64 = 1.5
                                            var v46 string = ((("a" + "foo bar") + "hello") + "")
                                            switch (fi35 & v3) {
                                            case 9223372036854775807:
                                                if false {
                                                    v47 := GenS7{F0: ((v46 + v46) + v46), F1: 42}
                                                    v48 := GenS4{F0: -882}
                                                    v45 = 3.14159
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    switch 387 {
                                                    case -9223372036854775808:
                                                        var v49 int64 = 1
                                                        _ = v49
                                                    }
                                                    v50 := GenS2{F0: (1 > 18446744073709551615)}
                                                    _ = v50
                                                }
                                            case -9223372036854775808:
                                                var v51 uint8 = 127
                                                _ = v51
                                            case 32767:
                                                if (false || true) {
                                                    if true {
                                                        switch (fi26 | 42) {
                                                        case 2147483647:
                                                            v46 = v46
                                                        case -1:
                                                            v23 += (v23 - int16(v27))
                                                        case -98:
                                                            fmt.Println(v3, v27)
                                                        default:
                                                            v2 += 32767
                                                        }
                                                    }
                                                }
                                            }
                                            v23 = (v23 ^ genFn0())
                                            _ = v45
                                            _ = v46
                                        }
                                        v2 = 835
                                        var v52 float64 = 3.14159
                                        _ = v52
                                    case 42:
                                        switch fi35 {
                                        case 764:
                                            v2 += (v2 ^ (v2 - (^(v2))))
                                        case -9223372036854775808:
                                            if false {
                                                var v53 int16 = 644
                                                v23 = v23
                                                v54 := GenS7{F0: "foo bar", F1: 42}
                                                if true {
                                                    v20 = append(v20, fi26)
                                                }
                                                switch fi35 {
                                                case -566:
                                                    var v55 float64 = 0.0
                                                    _ = v55
                                                case 1:
                                                    v27 = v27
                                                case -9223372036854775808:
                                                    v20 = append(v20, 42)
                                                }
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                    case -505:
                                        if len(v20) > 0 {
                                            v20[0] = (v3 / (((fi26 ^ 784)) | 1))
                                        }
                                        _ = len(v20)
                                    }
                                case 537:
                                    v23 = v23
                                }
                                v23 -= (-(v23))
                                _ = fi35
                            }
                        } else {
                            if true {
                                v1 = v1
                            }
                        }
                    }
                    _ = len(v20)
                }
            case "foo bar":
                var v56 int32 = 0
                _ = v56
            case "a":
                v2 = v2
            }
            _ = fi26
            _ = v27
        }
    }
    v57 := GenS8{F0: 9223372036854775807, F1: v2}
    if true {
        v24 -= v24
    }
    if (790 != 42) {
        for fi58 := 0; fi58 < 7; fi58++ {
            if len(v20) > 0 {
                v20[0] = 127
            }
            _ = fi58
        }
        var v59 int16 = (-(v23))
        fmt.Println(v2)
        _ = v59
    }
    _ = len(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v57
}
