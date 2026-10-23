package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
func genFn0() int8 {
    var v6 uint8 = 117
    v7 := []int{230, 590}
    fmt.Println(v6)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 int64 = 9223372036854775807
    _ = v6
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint32 = 1
    var v3 int8 = -128
    var v4 uint8 = 0
    v5 := []int{(v1 >> ((4294967295) & 63)), 753, v1, v1}
    fmt.Println(v4, v3)
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v1, v2, fi9)
        switch v1 {
        case -2147483648:
            if true {
                v4 = (^(v4))
                switch "" {
                case "\n\t":
                    switch fi9 {
                    case -392:
                        var v10 uint8 = v4
                        _ = v10
                    case 0:
                        v11 := GenS1{F0: 0, F1: fi9, F2: ("hello" < "")}
                        _ = v11
                    case 9223372036854775807:
                        switch (fi9 | fi9) {
                        case -9223372036854775808:
                            fmt.Println(v1, v2, v3)
                        default:
                            fmt.Println(v1)
                        }
                    }
                }
            } else {
                v2 += (v2 - v2)
                var v12 int64 = -636
                var v13 uint32 = 0
                _ = v12
                _ = v13
            }
        default:
            switch (fi9 ^ (v1 / (((fi9 | 496)) | 1))) {
            case -9223372036854775808:
                var v14 float64 = 0.1
                _ = v14
            case -584:
                for fi15 := 0; fi15 < 3; fi15++ {
                    if false {
                        switch fi15 {
                        case 100:
                            var v16 int64 = -9223372036854775808
                            _ = v16
                        case 128:
                            fmt.Println(fi15, v2)
                        default:
                            v5 = append(v5, fi9)
                        }
                        if false {
                            fmt.Println(v1, fi9, v2)
                            var v17 int32 = -2147483648
                            v3 = v3
                            _ = v17
                        } else {
                            var v18 uint32 = uint32(v4)
                            v19 := GenS2{F0: 127, F1: 153, F2: (v4 | v4)}
                            var v20 int8 = (v3 + (v3 / ((v3) | 1)))
                            fmt.Println(v18)
                            v2 += (^(v18))
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        fmt.Println(v4, v1, fi15)
                        switch (fi15 << ((132) & 63)) {
                        case -1:
                            var v21 int8 = v3
                            var v22 bool = (((("" + ("a" + "")) + "foo bar") + "\n\t") > "foo bar")
                            fmt.Println(fi9)
                            _ = v21
                            _ = v22
                        default:
                            v1 = (fi9 << ((255) & 63))
                        }
                        v4 = (v4 & v4)
                    }
                    if false {
                        if true {
                            var v23 float32 = (float32((1.5 * ((1e10 - 1e10) + (float64((1e10 + 3.14159)) / 1.5)))) / (1e10 - (0.0 / 1.5)))
                            if false {
                                fmt.Println(fi15, v23, v1)
                                v24 := GenS2{F0: (v3 & v3), F1: 513, F2: (^(v4))}
                                if len(v5) > 0 {
                                    v5[0] = v1
                                }
                                _ = v24
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v23
                        } else {
                            v2 = v2
                            fmt.Println(fi9)
                            if (39 != 42) {
                                v4 = 255
                                var v25 float32 = 1e10
                                v3 += genFn0()
                                _ = v25
                            }
                        }
                        if (1000 == 32767) {
                            v26 := GenS2{F0: v3, F1: 42, F2: 40}
                            fmt.Println(v1, v4)
                            _ = v26
                        }
                    }
                    _ = len(v5)
                    fmt.Println(fi15, v4, fi9)
                    var v27 int = 9223372036854775807
                    _ = fi15
                    _ = v27
                }
            }
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v3, fi9, v1)
            var v29 uint = 2
            _ = fi28
            _ = v29
        }
        _ = fi9
    }
    _ = len(v5)
    var v30 uint32 = (v2 + v2)
    v5 = append(v5, 32768)
    var v31 uint32 = (v2 | (v30 + 65535))
    if (!(4294967295 < 683)) {
        v5 = append(v5, 127)
        var v32 uint16 = 1
        var v33 int8 = (v3 % (((v3 & (v3 >> ((2) & 63)))) | 1))
        _ = v32
        _ = v33
    } else {
        v1 += (v1 - v1)
        v2 += (v30 / ((v31) | 1))
        v31 -= (v2 / (((v31 + (^(v31)))) | 1))
    }
    if false {
        var v34 uint8 = v4
        var v35 float32 = 0.1
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: 65535, F1: 9223372036854775807, F2: 19}
    switch "a" {
    case "foo bar":
        var v37 int8 = 1
        _ = v37
    }
    fmt.Println(v31, v3, v3)
    var v38 int16 = 496
    var v39 float32 = 0.1
    v40 := GenS1{F0: 0, F1: 42, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
