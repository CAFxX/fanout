package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int16 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float64 = -2.25
    var v3 int32 = 2147483647
    v4 := []int{127, -557, -184}
    v5 := []int{}
    v2 = ((1e10 * (float64((1e10 * 3.14159)) - 0.1)) / 1.5)
    if (false || false) {
        for fi6 := 0; fi6 < 16; fi6++ {
            if (true && false) {
                fmt.Println(fi6)
                v7 := GenS1{F0: 0, F1: 32767}
                v2 = (v2 * (1e10 - 1e10))
                _ = v7
            }
            if len(v4) > 0 {
                v4[0] = fi6
            }
            var v8 uint = 2
            _ = fi6
            _ = v8
        }
    } else {
        v5 = append(v5, 42)
        switch "\n\t" {
        case "\n\t":
            for fi9 := 0; fi9 < 17; fi9++ {
                switch ("hello" + "") {
                case "a":
                    v1 += 780
                default:
                    v10 := GenS2{F0: 2, F1: false, F2: 0}
                    _ = v10
                }
                _ = fi9
            }
        }
        switch -1 {
        case 32767:
            var v11 int32 = -200
            _ = v11
        default:
            v1 = v1
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v4)
    if true {
        fmt.Println(v2, v1, v3)
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS3{F0: v1, F1: v3, F2: (v1 / (((v1 << ((65535) & 63))) | 1))}
            v14 := GenS2{F0: 2, F1: ((-2.25 != (1.5 - 1.5)) && (0.0 > (((1e10 + (-2.25 * (((0.0 / 1e10) * ((-2.25 + 1.5) - 1.5)) * -2.25))) - ((1.5 / 0.1) / 0.1)) - 0.1))), F2: 42}
            v1 = (v1 * (v1 >> ((4294967295) & 63)))
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = 578
    switch ("\n\t" + "a") {
    case "a":
        var v16 int8 = 0
        _ = v16
    case "hello":
        fmt.Println(v15)
    }
    v17 := GenS1{F0: 42, F1: 2}
    switch 2147483647 {
    case -2147483648:
        for fi18 := 0; fi18 < 9; fi18++ {
            switch (^(fi18)) {
            case 1:
                if false {
                    _ = len(v4)
                    v1 = (^(v1))
                } else {
                    if ("foo bar" != "a") {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(v15, v15)
                            switch (^(fi18)) {
                            case 408:
                                var v20 uint32 = v1
                                _ = v20
                            default:
                                v21 := GenS3{F0: 260, F1: 0, F2: 4294967295}
                                _ = v21
                            }
                            switch fi19 {
                            case -252:
                                switch ("a" + "") {
                                case "hello":
                                    if false {
                                        _ = len(v5)
                                        v22 := GenS2{F0: 0, F1: true, F2: 0}
                                        fmt.Println(v3)
                                        _ = v22
                                    }
                                }
                            case -471:
                                v23 := GenS1{F0: 42, F1: 2}
                                _ = v23
                            case -205:
                                switch fi18 {
                                case 1:
                                    var v24 uint32 = v1
                                    v1 += (v1 ^ v1)
                                    _ = v24
                                case 0:
                                    fmt.Println(v1)
                                default:
                                    v25 := GenS3{F0: v1, F1: 42, F2: (v1 * v1)}
                                    _ = v25
                                }
                            }
                            _ = fi19
                        }
                    } else {
                        if true {
                            v15 = v15
                            fmt.Println(v1)
                            if len(v4) > 0 {
                                v4[0] = fi18
                            }
                        }
                        var v26 uint8 = 57
                        _ = v26
                    }
                    v5 = append(v5, 4294967295)
                    switch ("" + "foo bar") {
                    case "":
                        v5 = append(v5, fi18)
                    default:
                        var v27 int8 = -1
                        _ = v27
                    }
                }
            case 2147483647:
                v28 := GenS2{F0: 2, F1: false, F2: 9223372036854775807}
                _ = v28
            case 32768:
                switch (fi18 >> ((4294967295) & 63)) {
                case 0:
                    var v29 uint32 = v1
                    _ = v29
                case -415:
                    fmt.Println(v2)
                case -448:
                    var v30 bool = false
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 19; fi31++ {
                    v15 = v3
                    _ = fi31
                }
            }
            v32 := GenS4{F0: v1, F1: 100, F2: 253}
            v15 -= (v15 & (v3 ^ (v15 << ((255) & 63))))
            _ = fi18
            _ = v32
        }
    case 1:
        if (-137 == 9223372036854775807) {
            switch "foo bar" {
            case "\n\t":
                if (585 != 255) {
                    var v33 uint32 = v1
                    _ = v33
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        switch "" {
                        case "\n\t":
                            switch 9223372036854775807 {
                            case 0:
                                v2 += float64((0.0 + (((1.5 * ((float32(float64((-2.25 - (0.1 + 1.5)))) + (-2.25 + (float32(1.5) / -2.25))) - (float32(float64(0.0)) - 0.0))) / 1.0) - (-2.25 + (1e10 / float32(0.1))))))
                            case -2147483648:
                                switch 0 {
                                case 0:
                                    v2 = (((v2 / v2) + float64(0.1)) / ((0.0 - v2) * float64((1.5 * 1e10))))
                                case 2147483647:
                                    var v34 int = 699
                                    _ = v34
                                case 9223372036854775807:
                                    v1 += v1
                                default:
                                    v2 += float64(float32((v2 * float64(float32(((1e10 + 1e10) / ((float64((-2.25 - 3.14159)) + 0.1) + (v2 * v2))))))))
                                }
                            default:
                                fmt.Println(v1, v1, v1)
                            }
                        }
                    }
                    var v35 float64 = 3.14159
                    _ = v35
                }
            case "a":
                v36 := GenS3{F0: v1, F1: 705, F2: v1}
                _ = v36
            }
            switch 127 {
            case -1:
                for fi37 := 0; fi37 < 1; fi37++ {
                    v38 := GenS3{F0: (v1 << ((18446744073709551615) & 63)), F1: int32(fi37), F2: (v1 - (v1 ^ 2))}
                    _ = fi37
                    _ = v38
                }
            case -129:
                var v39 int8 = 32
                _ = v39
            case 0:
                var v40 uint16 = 237
                _ = v40
            }
        } else {
            v4 = append(v4, 301)
        }
    default:
        var v41 uint64 = 2
        _ = v41
    }
    if (v2 > (v2 + (((3.14159 * (float64(-2.25) + 0.1)) / ((v2 - -2.25) + ((v2 / float64((0.1 + (1.5 + float32(v2))))) + v2))) * (3.14159 + 0.0)))) {
        var v42 int64 = 9223372036854775807
        v43 := GenS5{F0: -1}
        v2 = 1e10
        _ = v42
        _ = v43
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v3)
    }
    fmt.Println(v15, v15, v1)
    if (false && (2 < 255)) {
        for fi44 := 0; fi44 < 16; fi44++ {
            v1 = 0
            _ = fi44
        }
    }
    v1 -= v1
    switch "\n\t" {
    case "":
        var v45 uint16 = 65535
        _ = v45
    default:
        var v46 int32 = v15
        _ = v46
    }
    fmt.Println(v2, v1)
    v47 := GenS4{F0: (v1 ^ (v1 >> ((980) & 63))), F1: -128, F2: 255}
    switch -443 {
    case 2147483647:
        v4 = append(v4, 512)
    case -485:
        v48 := GenS6{F0: 44, F1: 960}
        _ = v48
    case 32768:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
