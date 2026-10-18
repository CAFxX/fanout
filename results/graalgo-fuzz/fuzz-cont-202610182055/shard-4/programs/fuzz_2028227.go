package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 float32 = 1.5
    v2 := []int{42, -499}
    var v3 float64 = 0.0
    var v4 int16 = 1
    v5 := make(map[int]string)
    var v6 bool = true
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := v5[fi7]
        if (v6 || true) {
            if (!true) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch v8 {
                    case "\n\t":
                        switch ("hello" + "hello") {
                        case "hello":
                            switch v8 {
                            case "a":
                                v2 = append(v2, fi9)
                            default:
                                switch v8 {
                                case "":
                                    v10 := GenS1{F0: 42, F1: -640}
                                    _ = v10
                                case "foo bar":
                                    var v11 string = (v8 + ("" + "\n\t"))
                                    _ = v11
                                default:
                                    var v12 int32 = -1
                                    _ = v12
                                }
                            }
                        case "":
                            if (42 < 2147483647) {
                                switch v8 {
                                case "foo bar":
                                    fmt.Println(fi7)
                                case "\n\t":
                                    _ = len(v5)
                                default:
                                    v8 = ((((v8 + "a") + v8) + (v8 + "")) + "a")
                                }
                                var v13 uint8 = 128
                                _ = v13
                            }
                        }
                    default:
                        var v14 int = 506
                        _ = v14
                    }
                    v4 -= (v4 % ((1) | 1))
                    v8 = v8
                    _ = fi9
                }
                for fi15 := 0; fi15 < 17; fi15++ {
                    if false {
                        if (!(339 > 32767)) {
                            if false {
                                var v16 uint16 = 0
                                _ = v16
                            }
                            var v17 int32 = 2147483647
                            _ = v17
                        } else {
                            switch (fi15 - fi15) {
                            case -1:
                                v5[fi7] = "\n\t"
                            case 1:
                                if v6 {
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v3, v4)
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = fi7
                                    }
                                    var v18 int16 = v4
                                    v19 := GenS1{F0: 42, F1: -930}
                                    if v6 {
                                        var v20 int16 = 42
                                        _ = v20
                                    }
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(fi7, v6)
                                    v21 := GenS2{F0: -2147483648}
                                    _ = v21
                                }
                            }
                            v6 = v6
                            var v22 uint8 = 42
                            v1 += v1
                            _ = v22
                        }
                        v4 = (^(v4))
                        fmt.Println(fi7, v3, v3)
                    } else {
                        v23 := v5[fi7]
                        var v24 bool = (355 >= 42)
                        _ = v23
                        _ = v24
                    }
                    _ = fi15
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v4 += (v4 ^ (v4 | v4))
    v4 = (v4 ^ (v4 | (-(v4))))
    v25 := GenS2{F0: -2147483648}
    var v26 int = 2147483647
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v1, v4, v3)
        v4 -= -130
        _ = fi27
    }
    var v28 string = ("" + "")
    if len(v2) > 0 {
        v2[0] = v26
    }
    var v29 int64 = -9223372036854775808
    v1 += float32(float64(v1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v30 uint = 580
    v31 := GenS3{F0: (v26 ^ (v26 - v26)), F1: (v29 / (((v29 % ((9223372036854775807) | 1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
