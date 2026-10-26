package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 100
    var v2 string = "foo bar"
    v3 := make(map[uint32]int)
    var v4 int32 = -428
    var v5 uint64 = 0
    var v6 float32 = float32((0.0 - 3.14159))
    if true {
        if (255 == 2) {
            fmt.Println(v1, v4, v4)
            for fi7 := 0; fi7 < 1; fi7++ {
                fmt.Println(v1, v6, fi7)
                v1 = v1
                _ = fi7
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                delete(v3, 492)
                if (!false) {
                    if true {
                        for fi9 := 0; fi9 < 1; fi9++ {
                            v4 += v4
                            if (((^(v4)) != 504) && (false && (42 != 699))) {
                                fmt.Println(fi8)
                                v10 := GenS1{F0: fi8, F1: 65535}
                                v11 := GenS2{F0: fi9}
                                _ = v10
                                _ = v11
                            }
                            switch v2 {
                            case "foo bar":
                                v2 = v2
                            case "hello":
                                switch (fi9 ^ fi8) {
                                case -1:
                                    v1 = (^(v1))
                                case 905:
                                    v2 = v2
                                    v12 := []int{}
                                    _ = v12
                                default:
                                    v13 := GenS3{F0: 151}
                                    _ = v13
                                }
                            default:
                                v6 -= 3.14159
                            }
                            fmt.Println(v5, v6, v4)
                            v5 -= v5
                            _ = fi9
                        }
                        switch "" {
                        case "hello":
                            delete(v3, 401)
                        case "foo bar":
                            switch v2 {
                            case "\n\t":
                                fmt.Println(v5)
                            case "foo bar":
                                v14 := GenS2{F0: int(v1)}
                                _ = v14
                            case "a":
                                v15 := []int{}
                                _ = v15
                            default:
                                v1 = (^(v1))
                            }
                        case "":
                            v16 := []int{}
                            _ = v16
                        default:
                            var v17 uint16 = 42
                            _ = v17
                        }
                        if true {
                            var v18 float32 = 1e10
                            if true {
                                v18 = ((0.1 + float32(((0.1 + 0.1) / 1.0))) + v18)
                                v19 := GenS2{F0: -484}
                                var v20 int8 = (v1 << ((255) & 63))
                                _ = v19
                                _ = v20
                            } else {
                                v6 = v18
                                v21 := []int{-463, 42, (fi8 ^ 0), fi8}
                                v21 = append(v21, fi8)
                                _ = v21
                            }
                            var v22 int16 = -1
                            _ = v18
                            _ = v22
                        } else {
                            var v23 int = (-(fi8))
                            _ = v23
                        }
                    }
                    if (67 > 18446744073709551615) {
                        var v24 bool = false
                        var v25 string = (v2 + (v2 + "\n\t"))
                        v2 = "foo bar"
                        _ = v24
                        _ = v25
                    }
                    if true {
                        v2 = ("foo bar" + (v2 + v2))
                    }
                } else {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch (fi26 << ((255) & 63)) {
                        case 32767:
                            fmt.Println(v5)
                        }
                        v27 := GenS3{F0: v4}
                        _ = fi26
                        _ = v27
                    }
                }
                if true {
                    v28 := v3[2147483648]
                    if (((!false) || (!false)) && true) {
                        v6 -= v6
                    }
                    _ = v28
                }
                _ = fi8
            }
        }
        var v29 string = (((((v2 + v2) + "") + v2) + (("" + (("foo bar" + v2) + "hello")) + v2)) + "a")
        for fi30 := 0; fi30 < 5; fi30++ {
            v31 := []int{(fi30 & fi30), (fi30 - -34), (fi30 / ((fi30) | 1)), fi30}
            for fi32 := 0; fi32 < 1; fi32++ {
                v2 = (v2 + ((v29 + (("foo bar" + (v29 + ("hello" + (v2 + "\n\t")))) + v29)) + v29))
                v33 := GenS4{F0: 255}
                _ = fi32
                _ = v33
            }
            if len(v31) > 0 {
                v31[0] = fi30
            }
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        fmt.Println(v4)
    }
    v34 := GenS3{F0: (v4 ^ (v4 + v4))}
    if false {
        var v35 uint8 = 164
        _ = v35
    } else {
        v4 = (v4 << ((0) & 63))
        var v36 string = "foo bar"
        _ = v36
    }
    switch "" {
    case "hello":
        _ = len(v3)
    case "a":
        v5 = 367
    }
    v37 := []int{}
    v1 -= (v1 / (((v1 | (v1 * (v1 ^ -128)))) | 1))
    for fi38 := 0; fi38 < 15; fi38++ {
        v5 = (^(v5))
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS2{F0: (fi39 % ((fi39) | 1))}
            fmt.Println(v5, v2)
            fmt.Println(fi38, v2)
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    var v41 float32 = v6
    if false {
        v37 = append(v37, 4294967295)
    } else {
        v42 := GenS5{F0: 1, F1: 42, F2: 42}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
