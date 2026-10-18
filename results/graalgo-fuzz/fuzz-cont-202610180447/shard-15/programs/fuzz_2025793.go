package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 int8 = -1
    v2 := []uint32{1, 1, 4294967295, 760}
    v3 := []uint32{65535, 417, 2}
    var v4 int8 = v1
    if (4294967295 <= 2147483648) {
        v4 += v4
    } else {
        v5 := GenS1{F0: 65535, F1: 105}
        var v6 bool = ((2147483647 < 0) || true)
        if (!(65535 <= 4294967295)) {
            switch "foo bar" {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                if (0 < 42) {
                    var v7 int = 18
                    switch v7 {
                    case 32768:
                        var v8 uint16 = 255
                        _ = v8
                    case 75:
                        v7 += v7
                    }
                    var v9 uint8 = 199
                    _ = v7
                    _ = v9
                }
            case "a":
                if ((true && (0 <= -32768)) || v6) {
                    _ = len(v2)
                }
            default:
                for fi10 := 0; fi10 < 14; fi10++ {
                    switch ("foo bar" + "") {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 639
                        }
                    case "":
                        v6 = v6
                    default:
                        switch fi10 {
                        case 128:
                            v1 -= (v1 + -1)
                        case 713:
                            switch 32768 {
                            case -1:
                                v11 := GenS2{F0: int8(fi10), F1: 42, F2: 1}
                                _ = v11
                            case -2147483648:
                                var v12 int8 = 42
                                _ = v12
                            }
                        case -9223372036854775808:
                            for fi13 := 0; fi13 < 15; fi13++ {
                                switch ("" + "a") {
                                case "":
                                    var v14 int16 = 0
                                    _ = v14
                                }
                                _ = fi13
                            }
                        }
                    }
                    var v15 int32 = -93
                    _ = len(v2)
                    _ = fi10
                    _ = v15
                }
            }
            if (0 > 301) {
                if true {
                    var v16 string = "a"
                    if v6 {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            var v18 uint8 = 255
                            _ = fi17
                            _ = v18
                        }
                        var v19 uint32 = 1
                        if (!(31 > 42)) {
                            v2 = append(v2, (v19 >> ((313) & 63)))
                            if true {
                                var v20 float32 = 1e10
                                v21 := GenS1{F0: 364, F1: 18446744073709551615}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 int64 = -268
                                if false {
                                    v23 := GenS3{F0: (2147483647 == -2147483648), F1: 220, F2: 94}
                                    v24 := GenS1{F0: 2, F1: 954}
                                    fmt.Println(v4)
                                    _ = v23
                                    _ = v24
                                }
                                var v25 float32 = -2.25
                                _ = v22
                                _ = v25
                            }
                            v26 := GenS2{F0: (v4 >> ((65535) & 63)), F1: v19, F2: (v4 >> ((971) & 63))}
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 18; fi27++ {
                                var v28 uint32 = (v19 / (((^(v19))) | 1))
                                v3 = append(v3, (v28 & 2))
                                _ = fi27
                                _ = v28
                            }
                            var v29 string = (v16 + v16)
                            v2 = append(v2, 2147483647)
                            _ = v29
                        }
                        _ = v19
                    } else {
                        v16 = v16
                    }
                    v4 += v4
                    _ = v16
                } else {
                    v4 -= 0
                }
                if true {
                    v6 = (((554 > 18446744073709551615) && v6) || (32767 < 32767))
                } else {
                    var v30 uint8 = 2
                    if v6 {
                        v31 := GenS3{F0: (!v6), F1: v30, F2: v30}
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        }
                        _ = v31
                    }
                    _ = v30
                }
                for fi32 := 0; fi32 < 3; fi32++ {
                    v33 := GenS4{F0: ("" + "\n\t")}
                    _ = fi32
                    _ = v33
                }
            }
        } else {
            v4 += v4
            fmt.Println(v4, v4)
            var v34 uint8 = 42
            _ = v34
        }
        _ = v5
        _ = v6
    }
    v1 = (-(v1))
    var v35 uint32 = 278
    if true {
        for fi36 := 0; fi36 < 3; fi36++ {
            v35 += (^(v35))
            if (668 >= 1) {
                if true {
                    var v37 int32 = -1
                    var v38 uint = 72
                    _ = v37
                    _ = v38
                }
            } else {
                fmt.Println(v4, v4)
                v4 = v1
            }
            _ = fi36
        }
        var v39 int = -1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v39
    }
    v4 = v1
    var v40 uint = 65535
    v40 = v40
    fmt.Println(v40, v40)
    var v41 uint8 = 128
    for fi42 := 0; fi42 < 18; fi42++ {
        v41 = (^(v41))
        if ((v35 - 0) != (v35 >> ((v40) & 63))) {
            var v43 int16 = 1000
            v3 = append(v3, v35)
            fmt.Println(v40, v41, fi42)
            _ = v43
        }
        _ = fi42
    }
    v40 = (^(v40))
    switch 1 {
    case 128:
        for fi44 := 0; fi44 < 2; fi44++ {
            v45 := GenS5{F0: fi44, F1: 1}
            var v46 string = "hello"
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    var v47 bool = true
    switch "a" {
    case "\n\t":
        fmt.Println(v4)
    }
    fmt.Println(v1, v40, v35)
    v48 := GenS6{F0: (^(v41)), F1: 42, F2: v41}
    var v49 float64 = ((1e10 / -2.25) * float64(float32(((1.5 / 1.0) + 0.0))))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v48
    _ = v49
}
