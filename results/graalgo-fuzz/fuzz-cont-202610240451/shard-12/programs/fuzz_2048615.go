package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 string = "foo bar"
    v3 := []string{}
    var v4 int = 910
    v2 = v2
    switch (v4 & v4) {
    case -2147483648:
        var v5 uint8 = 128
        _ = v5
    case -9223372036854775808:
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 18446744073709551615}
            switch -599 {
            case -788:
                fmt.Println(v4, fi6, fi6)
            case 127:
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int64 = 313
                    v2 = v2
                    _ = fi8
                    _ = v9
                }
            }
            switch ("a" + v2) {
            case "\n\t":
                fmt.Println(fi6, fi6, v2)
            default:
                v10 := GenS1{F0: 18446744073709551615}
                _ = v10
            }
            _ = fi6
            _ = v7
        }
    case 1:
        v1[(v4 % ((v4) | 1))] = ""
    default:
        switch v4 {
        case 127:
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 float32 = 0.1
                var v13 float32 = 0.0
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    v14 := GenS1{F0: 4294967295}
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if true {
            v4 = (v4 << ((0) & 63))
            v4 -= v4
        }
        v15 := GenS2{F0: 819, F1: 18446744073709551615}
        _ = v15
    } else {
        v16 := GenS3{F0: 984}
        var v17 bool = (false || (42 >= 234))
        _ = v16
        _ = v17
    }
    fmt.Println(v2, v4)
    _ = len(v3)
    var v18 float32 = 0.0
    v4 = v4
    if true {
        switch v2 {
        case "foo bar":
            v19 := GenS3{F0: (^(v4))}
            _ = v19
        case "\n\t":
            for fi20 := 0; fi20 < 4; fi20++ {
                if false {
                    fmt.Println(v2, fi20)
                } else {
                    switch v2 {
                    case "a":
                        if false {
                            var v21 int8 = 56
                            _ = v21
                        } else {
                            var v22 uint16 = 255
                            _ = v22
                        }
                    case "":
                        for fi23 := 0; fi23 < 13; fi23++ {
                            fmt.Println(fi23)
                            switch v2 {
                            case "a":
                                v18 += ((((((1e10 + (0.1 - (((v18 / v18) + ((1e10 / (v18 - v18)) * -2.25)) - v18))) + 0.0) + v18) + float32((0.1 * 3.14159))) / v18) / 0.1)
                            case "\n\t":
                                v4 -= -556
                            }
                            _ = fi23
                        }
                    default:
                        var v24 int32 = -1
                        _ = v24
                    }
                    var v25 string = v2
                    _ = v25
                }
                v26 := GenS3{F0: -135}
                _ = fi20
                _ = v26
            }
        case "a":
            var v27 bool = false
            _ = v27
        }
        if true {
            var v28 int64 = 42
            _ = v28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
