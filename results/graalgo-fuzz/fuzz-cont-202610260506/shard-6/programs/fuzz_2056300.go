package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
type GenS4 struct { F0 int; F1 int64; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 string; F2 int16 }
type GenS8 struct { F0 string; F1 uint64; F2 bool }

func main() {
    v1 := []uint32{787, 2147483648, 2147483647}
    v2 := []int{-9223372036854775808, -9223372036854775808, 908}
    v3 := []int{109, 32767, 127}
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: false}
        v6 := GenS1{F0: true}
        var v7 int64 = 1
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 128
    v3 = append(v3, 443)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: false}
    var v10 bool = true
    for fi11 := 0; fi11 < 7; fi11++ {
        v2 = append(v2, fi11)
        if v10 {
            switch ("a" + ("" + ("foo bar" + "foo bar"))) {
            case "":
                _ = len(v2)
            case "\n\t":
                var v12 uint8 = (^(v8))
                _ = v12
            default:
                var v13 int8 = 34
                _ = v13
            }
            var v14 uint16 = 1
            fmt.Println(v8, v14)
            _ = v14
        }
        switch "hello" {
        case "":
            switch "foo bar" {
            case "a":
                v10 = v10
            }
        case "foo bar":
            if (1e10 > ((float32((float64(-2.25) / ((0.0 + (0.1 - ((1e10 * (((((1e10 * (-2.25 * (float64(3.14159) - (3.14159 - 1.5)))) + (-2.25 - 1e10)) + 0.0) / ((1e10 - 1.5) * (3.14159 / 1.5))) - float64(1.5))) - 0.0))) / 3.14159))) / (1.5 + 3.14159)) * (3.14159 - 3.14159))) {
                _ = len(v2)
            }
        case "a":
            if v10 {
                v8 -= (^(v8))
            }
        default:
            if (0.1 >= (1.5 * float64(((0.1 - 0.0) / (float32(0.1) * 1.5))))) {
                v2 = append(v2, (fi11 / (((fi11 >> ((65535) & 63))) | 1)))
            } else {
                switch "hello" {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = (fi11 % ((fi11) | 1))
                    }
                case "foo bar":
                    v8 = v8
                case "hello":
                    if v10 {
                        if (772 < uint32(v8)) {
                            if (((v10 || (!("foo bar" != (("a" + "hello") + "a")))) && v10) && ((720 > 512) || (3.14159 == -2.25))) {
                                v15 := GenS1{F0: ("hello" < ((("foo bar" + "foo bar") + ("\n\t" + "")) + ("foo bar" + "a")))}
                                fmt.Println(v8, v10, v8)
                                var v16 int8 = -98
                                _ = v15
                                _ = v16
                            } else {
                                fmt.Println(fi11, fi11)
                                _ = len(v3)
                            }
                        }
                        for fi17 := 0; fi17 < 1; fi17++ {
                            fmt.Println(fi11, fi11, fi11)
                            fmt.Println(v8)
                            _ = fi17
                        }
                    } else {
                        var v18 uint32 = 4294967295
                        if true {
                            var v19 bool = (v10 && ((-2.25 - 0.0) >= 1e10))
                            _ = v19
                        } else {
                            if v10 {
                                v8 += (v8 % (((v8 * v8)) | 1))
                            }
                            v18 += 2147483648
                        }
                        fmt.Println(v18)
                        _ = v18
                    }
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if false {
                            var v21 int = -295
                            var v22 uint8 = v8
                            _ = v21
                            _ = v22
                        } else {
                            if true {
                                v2 = append(v2, fi20)
                                var v23 uint32 = 315
                                v2 = append(v2, (^(fi11)))
                                _ = v23
                            }
                            v10 = (v10 || ((!(0.1 == -2.25)) || (-1 <= -1)))
                            var v24 uint64 = 1
                            switch "" {
                            case "":
                                var v25 int32 = -512
                                _ = v25
                            case "a":
                                v26 := GenS1{F0: (1e10 != (-2.25 + 1.5))}
                                _ = v26
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = (fi11 >> ((2) & 63))
                                }
                            default:
                                var v27 float64 = float64(1e10)
                                v24 -= v24
                                v8 = (v8 | (^(v8)))
                                _ = v27
                            }
                            _ = v24
                        }
                        v28 := GenS1{F0: (!true)}
                        var v29 int32 = -2147483648
                        _ = fi20
                        _ = v28
                        _ = v29
                    }
                }
                v30 := GenS1{F0: v10}
                _ = v30
            }
        }
        _ = fi11
    }
    var v31 float64 = (1e10 * 3.14159)
    if len(v3) > 0 {
        v3[0] = -1
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 uint32 = 2147483648
            v35 := GenS1{F0: (65535 != 42)}
            if false {
                fmt.Println(v31)
                v36 := GenS2{F0: 18446744073709551615, F1: (v34 - 4294967295), F2: 4294967295}
                v1 = append(v1, v34)
                _ = v36
            } else {
                if v10 {
                    if v10 {
                        v8 = 128
                        fmt.Println(v10, v34, fi32)
                    } else {
                        v37 := GenS3{F0: 32767, F1: v34, F2: ("a" + "\n\t")}
                        if false {
                            v38 := GenS4{F0: (fi33 >> ((0) & 63)), F1: 42, F2: (v34 << ((42) & 63))}
                            switch fi32 {
                            case 2147483647:
                                _ = len(v3)
                            case -885:
                                v39 := GenS5{F0: 42}
                                _ = v39
                            }
                            var v40 int32 = 551
                            _ = v38
                            _ = v40
                        } else {
                            v41 := GenS6{F0: 127}
                            fmt.Println(fi33, v34)
                            _ = v41
                        }
                        v42 := GenS7{F0: 42, F1: (("a" + "hello") + "a"), F2: -987}
                        _ = v37
                        _ = v42
                    }
                    var v43 bool = true
                    _ = v43
                } else {
                    v31 = (0.0 - 3.14159)
                }
                v44 := GenS8{F0: "", F1: 4294967295, F2: false}
                _ = v44
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    var v45 float64 = ((float64(float32((1e10 + 0.1))) - (v31 / 1e10)) / v31)
    var v46 string = "\n\t"
    fmt.Println(v31, v45, v45)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v45
    _ = v46
    _ = v8
    _ = v9
}
