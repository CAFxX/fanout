package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int64; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 uint }

func main() {
    var v1 int8 = 1
    var v2 string = "hello"
    v3 := []int{-2147483648, 127, 128, -910}
    v2 = v2
    var v4 int = 32768
    v3 = append(v3, (-(v4)))
    var v5 uint = 65535
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            v8 := GenS1{F0: 410, F1: 569, F2: v2}
            fmt.Println(v6, v6)
            fmt.Println(fi7, v5, v4)
            _ = v8
        } else {
            v4 = (-(v4))
        }
        var v9 int64 = 42
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS2{F0: (fi7 + 0), F1: v2, F2: 4294967295}
            switch "" {
            case "hello":
                var v12 float32 = -2.25
                _ = v12
            case "":
                var v13 uint = v5
                _ = v13
            case "foo bar":
                var v14 int64 = v9
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v9
    }
    v15 := GenS1{F0: 2147483647, F1: 2, F2: v2}
    var v16 float32 = (float32((-2.25 * 1e10)) - 1e10)
    v4 = (-(v4))
    v2 = "\n\t"
    fmt.Println(v1, v6)
    switch ("hello" + v2) {
    case "hello":
        switch "a" {
        case "\n\t":
            switch (v4 + (v4 / ((v4) | 1))) {
            case -714:
                v17 := GenS3{F0: 0, F1: 593}
                _ = v17
            case -255:
                if (42 > 9223372036854775807) {
                    if (!(!true)) {
                        v5 -= v5
                        var v18 int = 32767
                        var v19 uint = (v5 * 801)
                        _ = v18
                        _ = v19
                    } else {
                        if ("foo bar" > "hello") {
                            var v20 int16 = 734
                            _ = v20
                        }
                        var v21 string = v2
                        v6 = 127
                        _ = v21
                    }
                    v1 += v1
                }
            default:
                v22 := GenS3{F0: 42, F1: 1}
                _ = v22
            }
        default:
            v2 = (v2 + ("" + v2))
        }
    case "":
        if ((-2.25 / 1.5) < float64(1e10)) {
            v3 = append(v3, 1)
            var v23 bool = true
            _ = v23
        } else {
            if (0 == -261) {
                if (42 < -166) {
                    var v24 float32 = 0.0
                    fmt.Println(v6, v5)
                    if ((v1 & 100) < v1) {
                        var v25 uint64 = 570
                        fmt.Println(v24, v1, v24)
                        _ = v25
                    } else {
                        if (-351 != 0) {
                            v16 = v16
                            if len(v3) > 0 {
                                v3[0] = (v4 * v4)
                            }
                        } else {
                            v26 := GenS3{F0: -977, F1: 42}
                            if true {
                                var v27 int = 42
                                _ = v27
                            }
                            _ = v26
                        }
                        switch v2 {
                        case "\n\t":
                            for fi28 := 0; fi28 < 9; fi28++ {
                                fmt.Println(v1, v24)
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    switch ((v2 + "hello") + v2) {
                                    case "hello":
                                        v4 -= fi28
                                    case "":
                                        v1 = (v1 + v1)
                                    default:
                                        if false {
                                            v3 = append(v3, fi28)
                                        } else {
                                            fmt.Println(v6, v16, v1)
                                        }
                                    }
                                    var v30 int64 = -9223372036854775808
                                    v31 := GenS3{F0: 0, F1: -1}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                if true {
                                    fmt.Println(v16)
                                    var v32 int = v4
                                    _ = v32
                                } else {
                                    fmt.Println(v1)
                                }
                                _ = fi28
                            }
                        case "a":
                            v5 = v5
                        }
                        switch v2 {
                        case "\n\t":
                            if (!(13 < 717)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v33 int = (-(v4))
                                v34 := GenS4{F0: uint16(v6), F1: (v6 << ((4294967295) & 63)), F2: -858}
                                _ = v33
                                _ = v34
                            } else {
                                fmt.Println(v1)
                                var v35 int32 = -2147483648
                                if false {
                                    switch (v4 % ((9223372036854775807) | 1)) {
                                    case 9223372036854775807:
                                        switch (v4 + (v4 - v4)) {
                                        case 349:
                                            for fi36 := 0; fi36 < 20; fi36++ {
                                                var v37 float64 = (1.5 - (float64(float32(1e10)) / 1.0))
                                                if false {
                                                    for fi38 := 0; fi38 < 10; fi38++ {
                                                        if false {
                                                            v3 = append(v3, fi38)
                                                            v39 := GenS5{F0: 0, F1: 32767, F2: 550}
                                                            _ = v39
                                                        }
                                                        if (!false) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if (!false) {
                                                                v24 = v24
                                                            }
                                                            v3 = append(v3, 1)
                                                        } else {
                                                            var v40 float32 = v16
                                                            var v41 uint32 = 672
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        if false {
                                                            v3 = append(v3, fi36)
                                                        }
                                                        v42 := GenS6{F0: 192, F1: 127, F2: v2}
                                                        _ = fi38
                                                        _ = v42
                                                    }
                                                }
                                                _ = fi36
                                                _ = v37
                                            }
                                        default:
                                            v3 = append(v3, -140)
                                        }
                                    case 127:
                                        for fi43 := 0; fi43 < 2; fi43++ {
                                            if true {
                                                for fi44 := 0; fi44 < 20; fi44++ {
                                                    var v45 int16 = -1
                                                    v2 = ((v2 + "a") + v2)
                                                    var v46 int64 = -310
                                                    _ = fi44
                                                    _ = v45
                                                    _ = v46
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (fi43 - fi43)
                                                }
                                            }
                                            _ = fi43
                                        }
                                    default:
                                        fmt.Println(v5)
                                    }
                                    for fi47 := 0; fi47 < 19; fi47++ {
                                        if len(v3) > 0 {
                                            v3[0] = 232
                                        }
                                        var v48 float32 = v24
                                        switch v2 {
                                        case "foo bar":
                                            var v49 float64 = (0.0 * float64(v48))
                                            _ = v49
                                        default:
                                            fmt.Println(v2)
                                        }
                                        _ = fi47
                                        _ = v48
                                    }
                                }
                                _ = v35
                            }
                        case "a":
                            var v50 float64 = (((0.1 + 1.5) / 1.0) * 0.1)
                            _ = v50
                        case "foo bar":
                            switch v2 {
                            case "\n\t":
                                v16 -= 0.0
                            default:
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    fmt.Println(v6, v5)
                                    switch ("" + v2) {
                                    case "foo bar":
                                        fmt.Println(v1)
                                    case "a":
                                        _ = len(v3)
                                    case "hello":
                                        for fi52 := 0; fi52 < 14; fi52++ {
                                            fmt.Println(v16)
                                            fmt.Println(v2, v4)
                                            var v53 uint8 = (v6 | v6)
                                            switch (v2 + (v2 + v2)) {
                                            case "a":
                                                var v54 string = v2
                                                _ = v54
                                            case "foo bar":
                                                fmt.Println(v2, v53, v4)
                                            default:
                                                v55 := GenS5{F0: 32767, F1: 1, F2: 1}
                                                _ = v55
                                            }
                                            _ = fi52
                                            _ = v53
                                        }
                                    default:
                                        v56 := GenS7{F0: -32768, F1: 1}
                                        _ = v56
                                    }
                                    _ = fi51
                                }
                            }
                        default:
                            for fi57 := 0; fi57 < 19; fi57++ {
                                var v58 uint16 = 274
                                var v59 float32 = v24
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                    _ = v24
                }
            }
            _ = len(v3)
            for fi60 := 0; fi60 < 16; fi60++ {
                v1 -= (-(v1))
                _ = fi60
            }
        }
    case "\n\t":
        v16 += 1.5
    }
    v16 -= (v16 + (v16 * v16))
    if len(v3) > 0 {
        v3[0] = v4
    }
    v4 = (v4 & (v4 - (v4 / ((v4) | 1))))
    v5 = v5
    v4 -= 603
    var v61 float32 = float32(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
