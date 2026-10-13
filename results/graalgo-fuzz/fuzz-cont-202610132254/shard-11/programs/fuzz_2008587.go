package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint; F1 int32 }

func main() {
    v1 := []int{32768}
    var v2 int16 = 32767
    var v3 bool = true
    var v4 int32 = 42
    var v5 uint32 = 42
    var v6 int16 = 42
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: (v4 | v4), F1: 640}
        if v3 {
            if v3 {
                if (v3 || (-1 == 0)) {
                    fmt.Println(fi7, v6, v4)
                    fmt.Println(v4)
                }
            }
        } else {
            var v9 uint32 = v5
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if v3 {
        for fi10 := 0; fi10 < 1; fi10++ {
            _ = len(v1)
            v3 = (v5 > v5)
            fmt.Println(v2)
            _ = fi10
        }
        v5 -= v5
    }
    switch "" {
    case "foo bar":
        if v3 {
            v5 -= (v5 - v5)
        }
    case "hello":
        v2 = (-(v2))
    default:
        if (!true) {
            switch ((("a" + "") + "a") + "foo bar") {
            case "":
                for fi11 := 0; fi11 < 16; fi11++ {
                    if (1 < 4294967295) {
                        if v3 {
                            var v12 int16 = v6
                            switch (fi11 >> ((179) & 63)) {
                            case 32768:
                                for fi13 := 0; fi13 < 5; fi13++ {
                                    v4 -= (v4 | -2147483648)
                                    var v14 int16 = v12
                                    var v15 float64 = 0.1
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            case 209:
                                var v16 uint64 = 42
                                _ = v16
                            }
                            v2 -= -66
                            _ = v12
                        } else {
                            v4 += v4
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi17 := 0; fi17 < 4; fi17++ {
                                if false {
                                    v5 = 235
                                    fmt.Println(fi11, v3, v2)
                                } else {
                                    v18 := GenS2{F0: (v6 & 1), F1: (v6 + (v2 >> ((255) & 63))), F2: 14}
                                    var v19 int8 = 0
                                    _ = v18
                                    _ = v19
                                }
                                v6 += v6
                                _ = fi17
                            }
                        }
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS3{F0: v6}
                            var v22 uint8 = 127
                            fmt.Println(v2, v3)
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi11
                }
            case "hello":
                v23 := GenS4{F0: (v6 | (v6 * v6)), F1: v5}
                _ = v23
            }
        } else {
            v24 := GenS5{F0: 0, F1: 0}
            _ = v24
        }
    }
    var v25 int16 = v6
    v2 += -999
    v26 := GenS6{F0: 97, F1: (v5 / (((v5 - 1)) | 1))}
    v27 := GenS1{F0: -2147483648, F1: -924}
    v25 += 32767
    fmt.Println(v3)
    var v28 uint = 255
    if v3 {
        if (v3 || ((v3 && v3) && false)) {
            var v29 int32 = v4
            for fi30 := 0; fi30 < 18; fi30++ {
                v31 := GenS1{F0: (v4 & v29), F1: 0}
                if v3 {
                    v5 = v5
                    var v32 uint8 = 0
                    if len(v1) > 0 {
                        v1[0] = fi30
                    }
                    _ = v32
                } else {
                    _ = len(v1)
                    fmt.Println(v29, v4, v2)
                }
                _ = fi30
                _ = v31
            }
            if v3 {
                var v33 int32 = 2147483647
                for fi34 := 0; fi34 < 7; fi34++ {
                    _ = len(v1)
                    if false {
                        if (-568 <= v4) {
                            fmt.Println(v4, fi34)
                            switch fi34 {
                            case 2147483647:
                                v3 = (-1 < 1)
                            case -2147483648:
                                fmt.Println(v28, v33)
                            default:
                                fmt.Println(v3, v6)
                            }
                        } else {
                            v35 := GenS7{F0: 728, F1: (v29 - (v33 & (^(v4))))}
                            _ = v35
                        }
                    } else {
                        v6 = -32768
                        for fi36 := 0; fi36 < 7; fi36++ {
                            if v3 {
                                v28 -= 905
                                var v37 uint = 0
                                fmt.Println(fi34)
                                _ = v37
                            } else {
                                var v38 string = "a"
                                if len(v1) > 0 {
                                    v1[0] = 32767
                                }
                                _ = v38
                            }
                            var v39 float64 = 0.0
                            _ = fi36
                            _ = v39
                        }
                    }
                    _ = fi34
                }
                var v40 int16 = (v6 ^ v2)
                _ = v33
                _ = v40
            } else {
                var v41 int32 = -1
                var v42 int = 32767
                _ = v41
                _ = v42
            }
            _ = v29
        } else {
            var v43 float32 = (0.1 * ((((-2.25 * (1.5 - ((-2.25 * -2.25) * 3.14159))) - ((0.1 + float32(0.0)) / -2.25)) * (0.1 + float32((1.5 / 3.14159)))) + 3.14159))
            v1 = append(v1, -1)
            _ = v43
        }
        v28 = 1
        v1 = append(v1, -2147483648)
    } else {
        var v44 float64 = (3.14159 / float64(3.14159))
        for fi45 := 0; fi45 < 9; fi45++ {
            if len(v1) > 0 {
                v1[0] = fi45
            }
            _ = fi45
        }
        fmt.Println(v25, v5, v44)
        _ = v44
    }
    var v46 uint = (v28 >> ((v28) & 63))
    var v47 int8 = 0
    for fi48 := 0; fi48 < 13; fi48++ {
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v49 int64 = 43
            _ = v49
        default:
            v50 := GenS3{F0: v2}
            _ = v50
        }
        _ = fi48
    }
    var v51 uint16 = 0
    switch -667 {
    case -348:
        fmt.Println(v6, v46, v2)
    case 127:
        switch ("a" + "\n\t") {
        case "":
            var v52 float64 = 3.14159
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        if ((!v3) || (uint8(v5) < 163)) {
            v4 += (^(v4))
            v3 = (1.5 > 0.1)
        } else {
            v51 = v51
            var v54 uint64 = 891
            _ = v54
        }
        _ = fi53
    }
    var v55 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v55
    _ = v6
}
