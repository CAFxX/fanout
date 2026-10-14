package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 1000
    var v3 float32 = 1.5
    var v4 int32 = 904
    var v5 int16 = (v2 + (v2 & v2))
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 4; fi7++ {
        v5 = -1
        v2 += (v5 * v5)
        _ = fi7
    }
    if false {
        var v8 uint8 = 255
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int = (fi9 * (fi9 % (((fi10 % ((-9223372036854775808) | 1))) | 1)))
                _ = fi10
                _ = v11
            }
            v8 = 255
            _ = fi9
        }
        var v12 int64 = -943
        _ = v12
        _ = v8
    }
    fmt.Println(v2, v3, v2)
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            var v14 int64 = -922
            for fi15 := 0; fi15 < 5; fi15++ {
                switch (("foo bar" + "hello") + (("hello" + ("hello" + "foo bar")) + "")) {
                case "\n\t":
                    if true {
                        fmt.Println(v5)
                    } else {
                        v6[fi13] = 858
                        switch "a" {
                        case "hello":
                            if ((!false) && (("\n\t" < "\n\t") || true)) {
                                var v16 int64 = v14
                                var v17 int64 = (-(v14))
                                switch "\n\t" {
                                case "a":
                                    v16 += -9223372036854775808
                                case "hello":
                                    var v18 int = (fi13 | (^(fi13)))
                                    _ = v18
                                case "foo bar":
                                    var v19 string = "hello"
                                    if true {
                                        var v20 uint8 = 172
                                        v21 := []int{(^(fi13)), (fi15 * fi15)}
                                        v2 -= 237
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (!(!((v5 / (((v5 ^ v5)) | 1)) >= v2))) {
                                            fmt.Println(v5, fi13)
                                            v4 += (v4 & (v4 / ((0) | 1)))
                                        }
                                        var v22 int8 = -128
                                        _ = v22
                                    }
                                    v23 := GenS1{F0: 127, F1: 128}
                                    _ = v19
                                    _ = v23
                                default:
                                    var v24 int64 = v14
                                    _ = v24
                                }
                                _ = v16
                                _ = v17
                            }
                        case "":
                            fmt.Println(v4)
                        }
                        if (!true) {
                            switch -1 {
                            case 4294967295:
                                if true {
                                    fmt.Println(v2)
                                    var v25 int16 = v5
                                    _ = v25
                                }
                            }
                        } else {
                            v2 = -1
                            fmt.Println(fi15)
                            v5 -= (v5 >> ((255) & 63))
                        }
                    }
                }
                _ = fi15
            }
            switch (fi13 | 32768) {
            case 9223372036854775807:
                var v26 float64 = float64((v3 * v3))
                _ = v26
            case -979:
                for fi27 := 0; fi27 < 9; fi27++ {
                    v28 := []int{(fi13 << ((42) & 63)), fi13, 28, (fi13 << ((919) & 63))}
                    _ = fi27
                    _ = v28
                }
            default:
                v29 := GenS2{F0: 255, F1: 805}
                _ = v29
            }
            _ = v14
        }
        _ = fi13
    }
    var v30 uint64 = 284
    v2 -= (v5 & (v5 | (v5 + v5)))
    v30 -= (v30 << ((0) & 63))
    v3 = v3
    switch "" {
    case "a":
        fmt.Println(v3, v5, v30)
    }
    v31 := GenS2{F0: 2, F1: 0}
    for fi32 := 0; fi32 < 5; fi32++ {
        switch (fi32 >> ((4294967295) & 63)) {
        case 9223372036854775807:
            if (0 == 1) {
                v33 := GenS2{F0: 255, F1: 0}
                var v34 int64 = 247
                _ = v33
                _ = v34
            } else {
                fmt.Println(v3, v30)
                switch fi32 {
                case 32768:
                    v35 := GenS3{F0: -1}
                    _ = v35
                }
                switch (fi32 ^ fi32) {
                case -747:
                    switch ((("foo bar" + "") + "hello") + "\n\t") {
                    case "a":
                        var v36 uint16 = 42
                        _ = v36
                    case "hello":
                        switch fi32 {
                        case 32768:
                            v4 = v4
                        }
                    case "":
                        v37 := []int{int(v30), fi32, 127}
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 6; fi38++ {
                            v5 = (v5 << ((65535) & 63))
                            if (true || (2 < (v30 >> ((0) & 63)))) {
                                if true {
                                    delete(v6, (fi32 - -2147483648))
                                    v5 = v5
                                    fmt.Println(fi32)
                                } else {
                                    v39 := GenS4{F0: -128, F1: 27}
                                    _ = v39
                                }
                            } else {
                                if false {
                                    v40 := GenS4{F0: -101, F1: 2}
                                    v30 += v30
                                    _ = v40
                                } else {
                                    v2 = int16(v30)
                                    v41 := []int{}
                                    _ = v41
                                }
                                if (-2.25 >= 1.5) {
                                    fmt.Println(v5)
                                } else {
                                    if false {
                                        fmt.Println(fi38, v3)
                                        switch "foo bar" {
                                        case "\n\t":
                                            v2 = (v5 * (v5 ^ v2))
                                        case "":
                                            var v42 uint32 = 42
                                            _ = v42
                                        case "a":
                                            v3 += (v3 / v3)
                                        default:
                                            var v43 uint = 159
                                            _ = v43
                                        }
                                        v44 := v6[128]
                                        _ = v44
                                    }
                                }
                                v45 := GenS5{F0: -32768}
                                _ = v45
                            }
                            if true {
                                var v46 uint8 = 128
                                _ = v46
                            }
                            var v47 int = 408
                            var v48 int = v47
                            _ = fi38
                            _ = v47
                            _ = v48
                        }
                    }
                }
            }
        case 32768:
            fmt.Println(v5)
        case 4294967295:
            var v49 string = "foo bar"
            _ = v49
        default:
            v50 := GenS6{F0: int32(v2), F1: 205}
            _ = v50
        }
        _ = fi32
    }
    var v51 float32 = 0.0
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
