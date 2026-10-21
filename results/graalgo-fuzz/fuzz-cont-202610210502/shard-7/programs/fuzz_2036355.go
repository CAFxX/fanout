package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int = 953
    v2 := []int{(v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))}
    var v3 string = "\n\t"
    var v4 uint16 = 255
    var v5 uint64 = 139
    var v6 string = v3
    var v7 int32 = 230
    v8 := GenS1{F0: 0, F1: (("\n\t" + "a") + ""), F2: 2}
    var v9 uint64 = 164
    v9 -= v9
    v5 -= v5
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 = v6
        _ = fi10
    }
    if (!(-2.25 > 1e10)) {
        if (9223372036854775807 >= v1) {
            v7 = (v7 * v7)
            if true {
                if ((true && (false && false)) || true) {
                    if (22 < 1) {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            v2 = append(v2, fi11)
                            if len(v2) > 0 {
                                v2[0] = -247
                            }
                            switch fi11 {
                            case 4294967295:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 42:
                                v4 = 226
                            case 2147483647:
                                var v12 int8 = int8(v4)
                                _ = v12
                            default:
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    if true {
                                        fmt.Println(v9)
                                        var v14 uint8 = 128
                                        _ = v14
                                    } else {
                                        fmt.Println(v9, v1, v3)
                                    }
                                    if false {
                                        var v15 int32 = 1
                                        var v16 uint64 = 0
                                        var v17 uint = 646
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = fi13
                                }
                            }
                            _ = fi11
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    v18 := GenS2{F0: true, F1: 1, F2: 471}
                    _ = v18
                } else {
                    var v19 uint16 = uint16(v9)
                    switch v1 {
                    case 128:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 uint8 = 252
                            _ = fi20
                            _ = v21
                        }
                    case 998:
                        v1 = (v1 * 1)
                    default:
                        v22 := GenS3{F0: 147, F1: (^(v9))}
                        _ = v22
                    }
                    _ = v19
                }
                fmt.Println(v9, v6)
                v1 = -2147483648
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (("a" + v3) + "hello") {
                case "a":
                    if true {
                        if len(v2) > 0 {
                            v2[0] = (-(v1))
                        }
                    }
                case "hello":
                    v7 = (-(v7))
                case "foo bar":
                    v7 += v7
                }
                v5 -= (v9 / (((v5 % ((v9) | 1))) | 1))
            }
        }
        var v23 uint64 = v9
        var v24 uint = 42
        _ = v23
        _ = v24
    } else {
        v25 := GenS2{F0: false, F1: 42, F2: 1}
        _ = v25
    }
    _ = len(v2)
    var v26 uint64 = 2
    v7 -= (-(v7))
    v27 := GenS3{F0: 255, F1: v9}
    switch ((("hello" + v3) + "hello") + ("\n\t" + v6)) {
    case "\n\t":
        v28 := GenS3{F0: 194, F1: (v9 | (v26 + (v9 << ((798) & 63))))}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
