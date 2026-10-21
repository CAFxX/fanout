package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 string; F2 uint16 }

func main() {
    v1 := []int{-934, 441, 32768, 42}
    var v2 uint16 = 65535
    var v3 uint8 = 42
    var v4 uint32 = 2
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 int64 = 9223372036854775807
                var v8 uint8 = 186
                var v9 int8 = -128
                v1 = append(v1, (fi6 % (((fi6 >> ((255) & 63))) | 1)))
                var v10 uint = uint(v4)
                _ = fi6
                _ = v10
                _ = v7
                _ = v8
                _ = v9
            }
            switch "hello" {
            case "\n\t":
                v11 := GenS1{F0: (v2 | 564)}
                _ = v11
            case "hello":
                if (v2 != uint16(v3)) {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        _ = len(v1)
                        if true {
                            var v13 float32 = 0.1
                            var v14 uint32 = (v4 - v4)
                            if (998 == 9223372036854775807) {
                                v15 := GenS2{F0: int8(fi12)}
                                switch ("\n\t" + "") {
                                case "a":
                                    var v16 uint16 = (v2 - v2)
                                    _ = v16
                                case "foo bar":
                                    if (-1 >= -1) {
                                        var v17 int64 = 1
                                        switch (fi5 >> ((260) & 63)) {
                                        case 32767:
                                            fmt.Println(v2)
                                        default:
                                            var v18 uint16 = (v2 * 42)
                                            _ = v18
                                        }
                                        _ = v17
                                    } else {
                                        var v19 string = "foo bar"
                                        v2 = v2
                                        var v20 uint32 = (v4 << ((0) & 63))
                                        _ = v19
                                        _ = v20
                                    }
                                default:
                                    v13 -= 0.0
                                }
                                fmt.Println(v2, v3)
                                _ = v15
                            }
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(v2, v3)
                            fmt.Println(v3, v4, fi5)
                            switch ("" + "") {
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "foo bar":
                                fmt.Println(v3, fi5)
                            case "a":
                                if ("foo bar" == "\n\t") {
                                    v1 = append(v1, fi12)
                                } else {
                                    fmt.Println(v2)
                                    v21 := GenS1{F0: (v2 & (v2 % ((v2) | 1)))}
                                    var v22 uint = 42
                                    switch "" {
                                    case "a":
                                        fmt.Println(v2, v2, fi5)
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi12))
                                        }
                                        var v23 int8 = 56
                                        _ = v23
                                    case "foo bar":
                                        if (false && (1e10 > 0.0)) {
                                            var v24 int32 = -1
                                            _ = v24
                                        }
                                    case "":
                                        if true {
                                            v3 = (v3 / (((v3 >> ((0) & 63))) | 1))
                                            v3 = uint8(v22)
                                            v1 = append(v1, fi5)
                                        } else {
                                            v25 := GenS2{F0: -128}
                                            var v26 uint16 = 0
                                            fmt.Println(v2)
                                            _ = v25
                                            _ = v26
                                        }
                                        switch fi12 {
                                        case -2147483648:
                                            v3 = 55
                                        case -1:
                                            switch "foo bar" {
                                            case "a":
                                                var v27 float64 = ((float64(0.0) - 1e10) * 1.5)
                                                var v28 int = fi12
                                                _ = v27
                                                _ = v28
                                            default:
                                                if (!true) {
                                                    var v29 uint = (^(v22))
                                                    var v30 uint8 = (v3 % (((v3 + 128)) | 1))
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                        default:
                                            v31 := GenS3{F0: 216, F1: 42, F2: v3}
                                            _ = v31
                                        }
                                    }
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                if (127 <= 1) {
                                    if (((1e10 * -2.25) >= 0.0) || false) {
                                        fmt.Println(fi5)
                                        if true {
                                            v3 = (^(v3))
                                            var v32 uint16 = v2
                                            _ = v32
                                        }
                                        if true {
                                            switch (("\n\t" + "hello") + "a") {
                                            case "hello":
                                                var v33 float64 = 0.1
                                                _ = v33
                                            }
                                            v4 = v4
                                        }
                                    } else {
                                        var v34 bool = true
                                        v35 := GenS1{F0: v2}
                                        var v36 float32 = 0.1
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    v2 = 42
                                    switch fi12 {
                                    case 32767:
                                        v2 -= uint16(v4)
                                    }
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = -2147483648
                        }
                        v37 := GenS4{F0: 0}
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                var v38 int = int(v3)
                                var v39 float32 = 1e10
                                var v40 uint = 1
                                if len(v1) > 0 {
                                    v1[0] = fi12
                                }
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                var v41 uint64 = 0
                                v2 = v2
                                fmt.Println(fi12)
                                _ = v41
                            }
                        default:
                            var v42 bool = (!true)
                            _ = v42
                        }
                        switch "a" {
                        case "foo bar":
                            v43 := GenS5{F0: fi5}
                            _ = v43
                        case "\n\t":
                            var v44 int = (fi5 >> ((65535) & 63))
                            _ = v44
                        }
                        _ = fi12
                        _ = v37
                    }
                } else {
                    var v45 int16 = -1
                    _ = v45
                }
            default:
                fmt.Println(v3, v3)
            }
            _ = fi5
        }
    }
    v1 = append(v1, 9223372036854775807)
    v2 += 855
    var v46 bool = false
    v47 := GenS4{F0: 75}
    if v46 {
        var v48 uint8 = 197
        v49 := GenS6{F0: 0, F1: 688}
        v50 := GenS7{F0: true, F1: 1, F2: v3}
        _ = v48
        _ = v49
        _ = v50
    } else {
        switch -371 {
        case 2147483647:
            switch 705 {
            case 42:
                v4 += v4
            case -887:
                if len(v1) > 0 {
                    v1[0] = -2147483648
                }
            }
        default:
            v51 := GenS8{F0: 32767, F1: 42}
            _ = v51
        }
        switch 32767 {
        case 2147483647:
            _ = len(v1)
        case 32767:
            var v52 int16 = 1
            _ = v52
        case 4294967295:
            var v53 int64 = 9223372036854775807
            _ = v53
        default:
            v3 += v3
        }
        v54 := GenS9{F0: 991}
        _ = v54
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        for fi55 := 0; fi55 < 1; fi55++ {
            if len(v1) > 0 {
                v1[0] = int(v4)
            }
            v2 += v2
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 uint = 690
                var v58 bool = (1 > (v4 >> (((v57 ^ (^(v57)))) & 63)))
                var v59 uint64 = 1
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
            }
            _ = fi55
        }
    case "foo bar":
        for fi60 := 0; fi60 < 14; fi60++ {
            var v61 int8 = 0
            _ = fi60
            _ = v61
        }
    default:
        var v62 float32 = -2.25
        _ = v62
    }
    var v63 int32 = 1
    switch 880 {
    case 32767:
        fmt.Println(v4, v2)
    default:
        var v64 uint32 = uint32(v2)
        _ = v64
    }
    for fi65 := 0; fi65 < 13; fi65++ {
        v2 += v2
        _ = fi65
    }
    v66 := GenS10{F0: -32768, F1: ("foo bar" + ("\n\t" + "")), F2: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v63
    _ = v66
}
