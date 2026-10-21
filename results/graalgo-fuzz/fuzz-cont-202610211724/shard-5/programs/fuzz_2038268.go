package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }

func main() {
    var v1 float64 = (-2.25 / -2.25)
    var v2 string = ("\n\t" + "a")
    v3 := []string{v2}
    var v4 uint = 2
    var v5 int32 = 522
    v6 := GenS1{F0: 9223372036854775807}
    v1 -= v1
    v1 -= v1
    v4 -= (v4 % (((v4 >> ((v4) & 63))) | 1))
    fmt.Println(v2)
    var v7 uint8 = 12
    for fi8 := 0; fi8 < 20; fi8++ {
        if ((v1 + v1) <= 0.0) {
            fmt.Println(v7, v7, v1)
            for fi9 := 0; fi9 < 2; fi9++ {
                if (-2.25 >= float32(3.14159)) {
                    if (((3.14159 + ((((1.5 - v1) - v1) / v1) - -2.25)) == (0.0 * float64(float32(v1)))) && (!(1 >= 63))) {
                        if false {
                            var v10 int16 = 1000
                            if (!(2147483647 <= 0)) {
                                v3 = append(v3, v2)
                            }
                            var v11 uint = 1
                            _ = v10
                            _ = v11
                        }
                        if (!(true || false)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v3 = append(v3, ((v2 + "") + "foo bar"))
                            if false {
                                switch (v2 + ((v2 + v2) + v2)) {
                                case "hello":
                                    v7 -= 127
                                    v12 := GenS2{F0: 1, F1: -633}
                                    _ = v12
                                case "":
                                    var v13 uint8 = 92
                                    _ = v13
                                }
                                var v14 int8 = -92
                                var v15 int64 = -934
                                _ = v14
                                _ = v15
                            }
                        }
                    } else {
                        switch 648 {
                        case 297:
                            var v16 string = (v2 + v2)
                            _ = v16
                        case 127:
                            var v17 int8 = -86
                            _ = v17
                        case -9223372036854775808:
                            var v18 string = ("a" + v2)
                            _ = v18
                        default:
                            fmt.Println(fi9, v4)
                        }
                        var v19 int8 = int8(v4)
                        if len(v3) > 0 {
                            v3[0] = v2
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v19
                    }
                    v3 = append(v3, v2)
                }
                _ = fi9
            }
        } else {
            _ = len(v3)
        }
        var v20 uint = (v4 << ((v4) & 63))
        for fi21 := 0; fi21 < 10; fi21++ {
            v5 -= 480
            v22 := GenS2{F0: (v7 | 127), F1: v5}
            _ = fi21
            _ = v22
        }
        _ = fi8
        _ = v20
    }
    fmt.Println(v2)
    var v23 float32 = 1.5
    if (0 >= 32767) {
        for fi24 := 0; fi24 < 4; fi24++ {
            fmt.Println(v5)
            v25 := GenS2{F0: 18, F1: 42}
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS1{F0: fi26}
                switch "" {
                case "":
                    fmt.Println(v23, v23, v5)
                case "a":
                    v28 := GenS3{F0: -30, F1: 42}
                    _ = v28
                case "hello":
                    _ = len(v3)
                default:
                    fmt.Println(v2, v2)
                }
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v7)
    var v29 int64 = 0
    v30 := GenS4{F0: v2, F1: v5}
    var v31 int64 = (v29 % ((v29) | 1))
    switch 9223372036854775807 {
    case -1:
        switch "hello" {
        case "hello":
            if false {
                v4 += (v4 & v4)
            } else {
                v32 := GenS5{F0: -2147483648, F1: (v5 >> (((v4 | (v4 % ((v4) | 1)))) & 63)), F2: 319}
                _ = v32
            }
        }
    default:
        fmt.Println(v4)
    }
    switch v2 {
    case "hello":
        v33 := GenS4{F0: (v2 + v2), F1: v5}
        _ = v33
    default:
        v3 = append(v3, v2)
    }
    switch (v2 + v2) {
    case "hello":
        v1 -= ((v1 / (-2.25 / v1)) / 0.1)
    case "a":
        if (32767 != int16(v31)) {
            var v34 bool = true
            _ = v34
        }
    default:
        var v35 uint64 = 132
        _ = v35
    }
    if len(v3) > 0 {
        v3[0] = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
