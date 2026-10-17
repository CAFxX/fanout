package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
func genFn0() uint8 {
    switch "a" {
    case "foo bar":
        var v6 bool = (((true || (-2147483648 > 475)) && false) && false)
        _ = v6
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                var v8 int = fi7
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: (fi9 % ((fi9) | 1)), F1: 255}
                    v8 = (fi9 + 128)
                    var v11 bool = (11 >= -32768)
                    v8 = (fi9 << ((4294967295) & 63))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            } else {
                var v12 uint = 42
                if (578 >= 9223372036854775807) {
                    fmt.Println(fi7, fi7)
                    if true {
                        var v13 uint64 = 88
                        fmt.Println(v13, fi7)
                        for fi14 := 0; fi14 < 15; fi14++ {
                            v13 = v13
                            if (true && true) {
                                if false {
                                    v12 = v12
                                } else {
                                    if false {
                                        var v15 float64 = (float64(float32(-2.25)) * (0.0 / 1.5))
                                        _ = v15
                                    } else {
                                        var v16 uint32 = 872
                                        v17 := GenS2{F0: 871, F1: v13, F2: 0}
                                        _ = v16
                                        _ = v17
                                    }
                                    v12 = v12
                                }
                                v13 = (v13 - 4294967295)
                                switch "\n\t" {
                                case "\n\t":
                                    if (false || ((100 == -1) && false)) {
                                        fmt.Println(v13, v13, fi14)
                                    } else {
                                        fmt.Println(fi14, fi7)
                                        v13 = (v13 ^ (^(v13)))
                                        v12 = (v12 - v12)
                                    }
                                default:
                                    v18 := GenS3{F0: -2147483648}
                                    _ = v18
                                }
                            } else {
                                v12 = v12
                            }
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        switch fi7 {
                        case 328:
                            v19 := GenS4{F0: 455}
                            _ = v19
                        }
                        if false {
                            for fi20 := 0; fi20 < 15; fi20++ {
                                v12 -= v12
                                v12 = (^(v12))
                                _ = fi20
                            }
                            var v21 bool = true
                            v21 = (v21 && (!true))
                            _ = v21
                        } else {
                            switch ("a" + "a") {
                            case "\n\t":
                                if (!((-(fi7)) < -1)) {
                                    v22 := []int{(fi7 << ((v12) & 63))}
                                    fmt.Println(v12, fi7)
                                    _ = v22
                                } else {
                                    v23 := GenS3{F0: 1}
                                    var v24 int8 = 127
                                    _ = v23
                                    _ = v24
                                }
                            }
                        }
                    }
                    v25 := GenS3{F0: -1}
                    _ = v25
                }
                var v26 bool = true
                _ = v12
                _ = v26
            }
            _ = fi7
        }
    }
    var v27 uint32 = 1
    _ = v27
    return 0
}
func genFn1(p0 int64) string {
    v28 := GenS2{F0: -2147483648, F1: 997, F2: p0}
    v29 := []int{218, 32768, 128, 708}
    v30 := GenS4{F0: 32767}
    _ = p0
    _ = v28
    _ = v29
    _ = v30
    return ("" + "a")
}

func main() {
    var v1 uint = 226
    var v2 bool = (1 == 410)
    var v3 string = "hello"
    v4 := []uint32{2147483648}
    var v5 int8 = 22
    v31 := GenS4{F0: -903}
    if v2 {
        v4 = append(v4, 2)
    } else {
        var v32 uint = v1
        var v33 uint = v1
        _ = v32
        _ = v33
    }
    _ = len(v4)
    v2 = (42 != -720)
    var v34 int = -2147483648
    fmt.Println(v2)
    fmt.Println(v1, v3, v5)
    v35 := GenS3{F0: 105}
    if v2 {
        if true {
            v36 := GenS3{F0: -1}
            fmt.Println(v5, v34, v2)
            _ = v36
        } else {
            v2 = false
            if (!v2) {
                for fi37 := 0; fi37 < 5; fi37++ {
                    var v38 uint64 = 0
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v39 uint8 = 99
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                fmt.Println(v34, v1)
            } else {
                v3 = "hello"
                v40 := GenS5{F0: 4294967295, F1: 0}
                _ = v40
            }
            switch "a" {
            case "hello":
                switch "a" {
                case "\n\t":
                    if false {
                        v41 := GenS6{F0: (v1 ^ v1)}
                        v3 = (v3 + ((v3 + v3) + "hello"))
                        _ = v41
                    }
                case "a":
                    for fi42 := 0; fi42 < 15; fi42++ {
                        v34 += (fi42 * 32767)
                        _ = fi42
                    }
                case "hello":
                    v43 := GenS3{F0: 42}
                    _ = v43
                }
            default:
                v2 = (!v2)
            }
        }
        v44 := GenS7{F0: 253}
        _ = v44
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        _ = len(v4)
        v34 += 4294967295
        _ = fi45
    }
    switch (^(v34)) {
    case 32:
        v34 += v34
    case 328:
        var v46 uint16 = 65535
        _ = v46
    case 32767:
        if false {
            switch v3 {
            case "\n\t":
                fmt.Println(v3, v5)
            case "foo bar":
                var v47 uint16 = 1
                _ = v47
            default:
                var v48 uint32 = 0
                _ = v48
            }
        } else {
            switch v3 {
            case "\n\t":
                v3 = v3
            case "a":
                fmt.Println(v3)
            default:
                fmt.Println(v1)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
