package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 bool; F1 int; F2 int64 }
type GenS7 struct { F0 string; F1 int; F2 string }
type GenS8 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS9 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 float64 = (1e10 + 0.0)
    var v4 string = ""
    var v5 uint64 = 18446744073709551615
    var v6 float64 = (1.5 + (v3 + 0.0))
    fmt.Println(v3, v4, v4)
    if true {
        var v7 int = -2147483648
        var v8 int8 = 43
        v4 = (((v4 + "\n\t") + "hello") + ("a" + "hello"))
        _ = v7
        _ = v8
    }
    v9 := []int{9223372036854775807, 9223372036854775807}
    var v10 int16 = -32768
    v11 := GenS1{F0: -1}
    fmt.Println(v10)
    for fi12 := 0; fi12 < 18; fi12++ {
        v3 += v6
        for fi13 := 0; fi13 < 18; fi13++ {
            v5 = v5
            if true {
                if true {
                    fmt.Println(v3, v10)
                    var v14 uint = 42
                    _ = v14
                }
                v15 := GenS2{F0: v5, F1: true, F2: 2}
                var v16 float64 = 1e10
                _ = v15
                _ = v16
            } else {
                var v17 uint8 = 111
                fmt.Println(v6, v10, v4)
                switch fi12 {
                case 611:
                    if false {
                        if false {
                            var v18 int32 = 42
                            _ = v18
                        }
                        v4 = v4
                        v17 -= v17
                    }
                case 32768:
                    v19 := GenS3{F0: (v5 - 42), F1: -746}
                    _ = v19
                default:
                    var v20 int8 = -1
                    _ = v20
                }
                _ = v17
            }
            v21 := v2[(fi13 >> ((255) & 63))]
            _ = fi13
            _ = v21
        }
        _ = fi12
    }
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        if (!false) {
            if (!false) {
                v4 = v4
                v9 = append(v9, fi22)
            } else {
                for fi23 := 0; fi23 < 14; fi23++ {
                    v5 += 4294967295
                    _ = fi23
                }
            }
            v9 = append(v9, fi22)
            var v24 int64 = -9223372036854775808
            _ = v24
        } else {
            v25 := GenS3{F0: v5, F1: fi22}
            switch fi22 {
            case -1:
                v26 := GenS3{F0: (v5 >> ((42) & 63)), F1: fi22}
                _ = v26
            case -117:
                fmt.Println(v3, v3, v4)
            case 42:
                if false {
                    var v27 uint16 = 65535
                    v6 -= (v6 + ((float64((-2.25 + ((float32(v3) + (-2.25 - (float32(1.5) / 3.14159))) * 0.1))) - (((-2.25 - v3) + float64((1e10 * 0.0))) * ((((float64(0.1) + (v3 + (((v3 + float64(float32(float64(-2.25)))) / (((1e10 + 1e10) / v3) * v3)) / (v3 + (((((-2.25 + float64((((0.1 / (0.0 + (1e10 / 1.5))) + 3.14159) - 1.5))) * float64(3.14159)) + 0.1) / float64(3.14159)) / v6))))) * float64(0.0)) * 0.0) / 1.0))) + (((v3 + 1.5) * float64(float32(((((1e10 * float64((1e10 + (0.1 / -2.25)))) + (v6 + float64((-2.25 + 0.1)))) - v3) / v3)))) + 0.1)))
                    _ = v27
                }
            default:
                v28 := GenS1{F0: -810}
                _ = v28
            }
            if false {
                v29 := GenS4{F0: 2}
                var v30 float32 = (1.5 / (1.5 * 0.1))
                _ = v29
                _ = v30
            }
            _ = v25
        }
        v31 := GenS1{F0: 0}
        _ = fi22
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 int = -1
            switch fi33 {
            case 2147483647:
                var v35 int8 = 100
                _ = v35
            case -388:
                switch v4 {
                case "a":
                    v36 := GenS1{F0: -16}
                    _ = v36
                }
            }
            var v37 float32 = (1e10 / 1.5)
            switch v4 {
            case "":
                if false {
                    var v38 uint8 = 42
                    if true {
                        v39 := GenS5{F0: 5, F1: (v10 / (((v10 >> ((18446744073709551615) & 63))) | 1))}
                        if false {
                            if true {
                                v3 = 3.14159
                                switch "" {
                                case "a":
                                    if (0 == 686) {
                                        var v40 int8 = 100
                                        var v41 uint8 = v38
                                        v42 := GenS6{F0: true, F1: 1, F2: -60}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                case "\n\t":
                                    var v43 string = "foo bar"
                                    _ = v43
                                }
                            }
                            var v44 bool = (true || (!true))
                            switch "foo bar" {
                            case "a":
                                v37 += (float32((v6 - 0.1)) - 0.1)
                            }
                            _ = v44
                        }
                        fmt.Println(v4, v34, v37)
                        _ = v39
                    } else {
                        var v45 bool = false
                        switch "a" {
                        case "hello":
                            v46 := GenS7{F0: v4, F1: 32768, F2: "hello"}
                            _ = v46
                        case "":
                            if true {
                                if true {
                                    if v45 {
                                        fmt.Println(v34)
                                        var v47 int = 9223372036854775807
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            v6 -= -2.25
                                        case "foo bar":
                                            v1[((("foo bar" + "a") + v4) + "")] = 4294967295
                                        case "hello":
                                            var v48 int16 = 1000
                                            _ = v48
                                        default:
                                            v6 = ((1e10 + (float64(v37) + -2.25)) + (v6 + v3))
                                        }
                                        _ = v47
                                    }
                                    v49 := GenS4{F0: 0}
                                    v50 := GenS8{F0: 787, F1: (v5 >> ((255) & 63)), F2: 32767}
                                    var v51 float32 = (1e10 * v37)
                                    if (v51 <= -2.25) {
                                        var v52 uint32 = 117
                                        v1["hello"] = uint32(v5)
                                        _ = v52
                                    } else {
                                        v34 = fi33
                                        v38 = (v38 | 128)
                                        var v53 int8 = 1
                                        _ = v53
                                    }
                                    v9 = append(v9, (fi33 % ((668) | 1)))
                                    v54 := GenS4{F0: 2}
                                    if (!(1 < 65535)) {
                                        var v55 int = fi33
                                        _ = v55
                                    }
                                    v56 := GenS1{F0: 42}
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                    _ = v54
                                    _ = v56
                                }
                            }
                        default:
                            v3 = -2.25
                        }
                        _ = v45
                    }
                    _ = v38
                }
            case "\n\t":
                v10 = -1
            case "hello":
                v34 -= 42
            }
            _ = fi33
            _ = v34
            _ = v37
        }
        v57 := GenS9{F0: (((3.14159 + float32(v3)) + 3.14159) > (1.5 + 3.14159))}
        _ = fi32
        _ = v57
    }
    v58 := GenS6{F0: false, F1: 4294967295, F2: -1}
    for fi59 := 0; fi59 < 19; fi59++ {
        _ = len(v2)
        _ = fi59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v9
}
