package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS9 struct { F0 int32 }

func main() {
    v1 := []string{"\n\t", "a"}
    var v2 uint = 65535
    var v3 uint64 = 145
    var v4 int8 = 127
    var v5 int32 = 2147483647
    var v6 int64 = 911
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(v4, v6)
            _ = fi8
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            v1 = append(v1, "\n\t")
            v1 = append(v1, "a")
            _ = fi9
        }
        if (931 < 42) {
            var v10 int = (fi7 >> ((v2) & 63))
            v2 = 94
            _ = v10
        }
        _ = fi7
    }
    var v11 int32 = v5
    v2 += 2
    if (1000 > 32767) {
        v4 += v4
        var v12 int8 = 1
        switch 829 {
        case -2147483648:
            _ = len(v1)
        default:
            v5 += (v11 << ((779) & 63))
        }
        _ = v12
    } else {
        var v13 int = 128
        _ = v13
    }
    if true {
        var v14 float32 = float32((-2.25 - 0.1))
        for fi15 := 0; fi15 < 12; fi15++ {
            if true {
                v2 += 919
                v16 := GenS1{F0: 521}
                _ = len(v1)
                _ = v16
            }
            _ = fi15
        }
        _ = v14
    }
    v3 = v3
    var v17 float32 = 3.14159
    v3 = (v3 ^ v3)
    _ = len(v1)
    var v18 bool = true
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 float32 = float32(-2.25)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS1{F0: 2}
        for fi23 := 0; fi23 < 11; fi23++ {
            if v18 {
                v11 -= 1
                if len(v1) > 0 {
                    v1[0] = ("foo bar" + ((("hello" + "a") + "\n\t") + "hello"))
                }
            } else {
                switch "hello" {
                case "hello":
                    v1 = append(v1, "\n\t")
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    v3 += (v3 ^ (v3 & (v3 - v3)))
                }
                v24 := GenS1{F0: 828}
                _ = v24
            }
            fmt.Println(v6, v17, v5)
            _ = fi23
        }
        v25 := GenS2{F0: 731}
        _ = fi21
        _ = v22
        _ = v25
    }
    v6 = -745
    switch "\n\t" {
    case "\n\t":
        if true {
            if len(v1) > 0 {
                v1[0] = ""
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v2, v3)
            default:
                switch (("" + "hello") + "foo bar") {
                case "\n\t":
                    v26 := GenS1{F0: 255}
                    _ = v26
                case "hello":
                    v27 := GenS3{F0: v11}
                    _ = v27
                case "foo bar":
                    switch ("hello" + "") {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    case "":
                        v6 += (^(v6))
                    default:
                        var v28 uint16 = 1
                        _ = v28
                    }
                default:
                    fmt.Println(v4, v6, v6)
                }
            }
        } else {
            if ((float64(v17) * -2.25) < (0.0 + (0.0 / (0.1 * ((float64(1.5) * (1e10 / (0.1 + float64(3.14159)))) + 1.5))))) {
                var v29 float64 = 0.1
                _ = len(v1)
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := GenS3{F0: (-(v5))}
                    var v32 float64 = v29
                    var v33 uint16 = 1
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            } else {
                v34 := GenS4{F0: 2147483647, F1: 1}
                for fi35 := 0; fi35 < 20; fi35++ {
                    switch "" {
                    case "foo bar":
                        v36 := GenS2{F0: (v2 | v2)}
                        _ = v36
                    }
                    switch 0 {
                    case -1:
                        v5 = (-(v11))
                    case 9223372036854775807:
                        for fi37 := 0; fi37 < 10; fi37++ {
                            switch (("\n\t" + ("" + ("foo bar" + "\n\t"))) + ("\n\t" + "")) {
                            case "a":
                                v11 = (v11 >> (((v2 % ((255) | 1))) & 63))
                            case "foo bar":
                                v1 = append(v1, "")
                            }
                            switch fi35 {
                            case -9223372036854775808:
                                var v38 int32 = v5
                                _ = v38
                            case 42:
                                var v39 int64 = v6
                                _ = v39
                            default:
                                if v18 {
                                    if true {
                                        v4 = (v4 >> (((v2 * (v2 & v2))) & 63))
                                        v40 := GenS3{F0: (v11 | (v11 * (v11 / ((v5) | 1))))}
                                        var v41 float64 = 0.1
                                        fmt.Println(fi37, v18, v4)
                                        var v42 int16 = 942
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 uint16 = 42
                                        v44 := GenS1{F0: (v43 + uint16(v2))}
                                        v1 = append(v1, ("\n\t" + "hello"))
                                        v45 := GenS5{F0: (v18 || false), F1: 620}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                } else {
                                    var v46 uint32 = 0
                                    var v47 uint16 = 42
                                    var v48 int32 = v5
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                            }
                            v17 = v17
                            _ = fi37
                        }
                    case -27:
                        v49 := GenS6{F0: -128, F1: 193}
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 int64 = (v6 & v6)
                        v52 := GenS7{F0: 2147483647}
                        if false {
                            fmt.Println(fi35, fi50)
                            v53 := GenS8{F0: v3, F1: v11, F2: 72}
                            _ = v53
                        } else {
                            var v54 uint8 = uint8(v3)
                            _ = v54
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = fi35
                }
                _ = v34
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                var v56 int = (^(fi55))
                v6 = (^(v6))
                _ = fi55
                _ = v56
            }
            if v18 {
                v57 := GenS9{F0: (v5 ^ 957)}
                fmt.Println(v6, v4)
                for fi58 := 0; fi58 < 9; fi58++ {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    var v59 bool = ((fi58 * fi58) <= (fi58 % ((-1) | 1)))
                    _ = fi58
                    _ = v59
                }
                _ = v57
            } else {
                var v60 float32 = float32(1e10)
                fmt.Println(v6, v11, v5)
                var v61 int = 272
                _ = v60
                _ = v61
            }
        }
    case "hello":
        v62 := GenS4{F0: 629, F1: 284}
        _ = v62
    case "":
        if ((v4 + v4) > (v4 >> (((v2 >> (((^(v2))) & 63))) & 63))) {
            fmt.Println(v3, v4)
        }
    }
    var v63 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
