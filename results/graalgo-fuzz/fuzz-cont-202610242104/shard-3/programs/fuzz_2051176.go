package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }
type GenS6 struct { F0 uint8; F1 uint32; F2 string }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 int16 }

func main() {
    v1 := []int{-1, 1}
    var v2 int32 = -1
    var v3 uint16 = 42
    var v4 uint8 = 128
    switch "" {
    case "hello":
        switch 32767 {
        case -138:
            var v5 int16 = -1
            _ = v5
        case 9223372036854775807:
            for fi6 := 0; fi6 < 20; fi6++ {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 uint16 = (v3 ^ 32767)
                    if (1 >= -39) {
                        v9 := GenS1{F0: (v3 | uint16(v4))}
                        if (242 <= 651) {
                            var v10 int64 = -1
                            var v11 uint16 = (v8 * uint16(v4))
                            _ = v10
                            _ = v11
                        } else {
                            v12 := GenS1{F0: (v8 * v3)}
                            v8 = v3
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v12
                        }
                        _ = v9
                    }
                    v2 = (^(v2))
                    _ = fi7
                    _ = v8
                }
                var v13 int16 = 675
                if len(v1) > 0 {
                    v1[0] = 4
                }
                _ = fi6
                _ = v13
            }
        }
    case "\n\t":
        v4 = (^(v4))
    case "":
        v3 += 679
    }
    v14 := GenS2{F0: 46, F1: v3, F2: v4}
    switch "a" {
    case "a":
        switch 42 {
        case 42:
            v15 := GenS3{F0: -9223372036854775808, F1: 0}
            _ = v15
        }
    case "foo bar":
        switch "\n\t" {
        case "hello":
            v16 := GenS2{F0: -18, F1: (v3 & 1), F2: (v4 ^ 47)}
            _ = v16
        default:
            if false {
                for fi17 := 0; fi17 < 8; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS4{F0: 1}
                        _ = fi18
                        _ = v19
                    }
                    var v20 float64 = 1.5
                    _ = fi17
                    _ = v20
                }
            } else {
                switch "\n\t" {
                case "hello":
                    if ((0.0 + 1e10) < 1e10) {
                        v1 = append(v1, 999)
                        v2 -= (-(v2))
                    }
                case "":
                    _ = len(v1)
                }
            }
        }
    default:
        var v21 bool = false
        _ = v21
    }
    fmt.Println(v3, v3, v2)
    var v22 uint32 = 63
    var v23 int64 = 42
    v3 += (v3 ^ 776)
    var v24 string = "\n\t"
    var v25 int32 = v2
    v22 = 4294967295
    var v26 uint32 = (^(v22))
    var v27 int8 = 100
    if false {
        var v28 bool = (!true)
        v29 := GenS2{F0: 1, F1: v3, F2: (v4 & 127)}
        switch 0 {
        case 32767:
            fmt.Println(v26)
        case 9223372036854775807:
            v30 := GenS5{F0: 42, F1: v27, F2: v28}
            _ = v30
        case 32768:
            switch -99 {
            case 32767:
                var v31 bool = ((-(v27)) != (v27 / ((v27) | 1)))
                _ = v31
            case 127:
                fmt.Println(v23, v23, v23)
            case 0:
                var v32 int = 2147483647
                _ = v32
            default:
                v2 -= v25
            }
        default:
            if v28 {
                if (v28 || (v26 != (v26 ^ (v22 - 0)))) {
                    v24 = ""
                } else {
                    var v33 float32 = ((-2.25 - 0.1) + 0.0)
                    _ = v33
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    v35 := GenS6{F0: v4, F1: (v26 << ((4294967295) & 63)), F2: (("a" + ((v24 + "hello") + "")) + v24)}
                    switch -305 {
                    case 127:
                        if v28 {
                            v36 := GenS7{F0: 4294967295, F1: uint64(v3)}
                            fmt.Println(v27)
                            for fi37 := 0; fi37 < 1; fi37++ {
                                if v28 {
                                    var v38 uint = 42
                                    v39 := GenS2{F0: 0, F1: (v3 ^ (v3 & v3)), F2: v4}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v23 += (v23 / ((v23) | 1))
                                    if (!(0.1 <= 1e10)) {
                                        v40 := GenS8{F0: 960, F1: 1}
                                        var v41 float32 = 1.5
                                        if false {
                                            var v42 bool = v28
                                            var v43 bool = true
                                            v24 = v24
                                            _ = v42
                                            _ = v43
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        v22 = (v22 | (v22 / ((uint32(v3)) | 1)))
                                    case "":
                                        var v44 uint8 = (v4 / ((128) | 1))
                                        _ = v44
                                    case "foo bar":
                                        if len(v1) > 0 {
                                            v1[0] = (fi37 - 4294967295)
                                        }
                                    default:
                                        v1 = append(v1, fi37)
                                    }
                                }
                                var v45 int32 = (v25 ^ (v2 << ((817) & 63)))
                                _ = fi37
                                _ = v45
                            }
                            _ = v36
                        } else {
                            var v46 uint32 = (v22 * v26)
                            _ = v46
                        }
                    case 1:
                        v25 = -256
                    case 32768:
                        var v47 uint64 = 2
                        _ = v47
                    }
                    _ = fi34
                    _ = v35
                }
                var v48 bool = (!false)
                _ = v48
            }
        }
        _ = v28
        _ = v29
    }
    v49 := GenS9{F0: -836}
    fmt.Println(v24, v22, v23)
    fmt.Println(v3, v3, v27)
    switch "a" {
    case "foo bar":
        switch 128 {
        case -332:
            for fi50 := 0; fi50 < 14; fi50++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v51 := GenS3{F0: 312, F1: 0}
                _ = fi50
                _ = v51
            }
        case -96:
            var v52 float32 = 1e10
            _ = v52
        default:
            v53 := GenS8{F0: 32767, F1: 0}
            _ = v53
        }
    case "hello":
        fmt.Println(v4, v22, v26)
    case "":
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    default:
        var v54 float64 = 0.0
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v49
}
