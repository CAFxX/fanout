package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]uint32)
    var v3 uint = 289
    var v4 bool = true
    v5 := []int{(v1 % ((32767) | 1))}
    if ("\n\t" < "a") {
        var v6 int64 = -9223372036854775808
        if len(v5) > 0 {
            v5[0] = (v1 >> (((v3 | (v3 | (^(v3))))) & 63))
        }
        switch ("\n\t" + "foo bar") {
        case "a":
            if true {
                _ = len(v2)
            } else {
                switch (("" + "hello") + "hello") {
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        var v7 int = (v1 & (v1 & v1))
                        _ = v7
                    case "a":
                        v5 = append(v5, 1)
                    default:
                        if (!true) {
                            v1 = (v1 - v1)
                            v8 := GenS1{F0: 2, F1: 1, F2: 0}
                            _ = len(v2)
                            _ = v8
                        }
                    }
                case "":
                    for fi9 := 0; fi9 < 14; fi9++ {
                        for fi10 := 0; fi10 < 20; fi10++ {
                            v11 := GenS1{F0: 2147483647, F1: 2147483647, F2: 18446744073709551615}
                            _ = fi10
                            _ = v11
                        }
                        if v4 {
                            v6 = v6
                            for fi12 := 0; fi12 < 4; fi12++ {
                                v4 = (v4 && v4)
                                v13 := GenS1{F0: 4294967295, F1: 1, F2: 1}
                                v14 := GenS2{F0: -2147483648, F1: (63 <= 903)}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            if false {
                                v3 = 0
                            }
                        } else {
                            var v15 string = ("foo bar" + ("\n\t" + "foo bar"))
                            _ = v15
                        }
                        switch "hello" {
                        case "\n\t":
                            for fi16 := 0; fi16 < 13; fi16++ {
                                v17 := GenS3{F0: 32767}
                                var v18 int16 = -632
                                v6 = v6
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = fi9
                    }
                default:
                    v19 := GenS4{F0: (v3 >> ((v3) & 63))}
                    _ = v19
                }
                v2["\n\t"] = 65535
            }
        case "hello":
            var v20 int = (v1 / (((v1 % (((-(v1))) | 1))) | 1))
            _ = v20
        default:
            if len(v5) > 0 {
                v5[0] = v1
            }
        }
        _ = v6
    }
    fmt.Println(v1, v3)
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS3{F0: 104}
        if ((^(v3)) < 597) {
            for fi23 := 0; fi23 < 9; fi23++ {
                switch "hello" {
                case "\n\t":
                    var v24 float32 = (float32(-2.25) + 3.14159)
                    if ((0.1 / float32(1e10)) > v24) {
                        v24 = (0.0 + v24)
                        delete(v2, "foo bar")
                        if (65535 > 2147483647) {
                            var v25 int64 = -940
                            v2["hello"] = 1
                            _ = v25
                        }
                    }
                    var v26 uint64 = 0
                    _ = v24
                    _ = v26
                }
                delete(v2, "")
                _ = fi23
            }
        } else {
            var v27 int64 = 42
            _ = v27
        }
        _ = fi21
        _ = v22
    }
    var v28 float32 = float32(1e10)
    var v29 float64 = 3.14159
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 uint32 = 2147483648
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS5{F0: 219}
            v29 += float64(v28)
            _ = fi32
            _ = v33
        }
        v3 = 1
        _ = fi30
        _ = v31
    }
    if true {
        v1 += 110
        for fi34 := 0; fi34 < 3; fi34++ {
            v5 = append(v5, v1)
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
