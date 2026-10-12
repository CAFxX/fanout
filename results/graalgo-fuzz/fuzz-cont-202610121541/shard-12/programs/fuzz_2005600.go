package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 bool }
type GenS3 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 int32 = 161
    v2 := []int{-9223372036854775808, 1, 128}
    var v3 int16 = -144
    v4 := GenS1{F0: (^(v1))}
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if false {
                var v7 uint8 = 2
                var v8 uint = 144
                fmt.Println(v7)
                if false {
                    var v9 float64 = 1.5
                    _ = v9
                } else {
                    v8 = v8
                    switch (fi5 * fi5) {
                    case -1:
                        var v10 int64 = 1
                        _ = v10
                    case -782:
                        v1 += v1
                    case 538:
                        switch fi6 {
                        case -9223372036854775808:
                            v7 = (v7 / (((v7 + (v7 * 128))) | 1))
                        case -449:
                            _ = len(v2)
                        case 348:
                            v7 = v7
                        }
                        var v11 int = (fi5 % ((fi6) | 1))
                        _ = v11
                    default:
                        v2 = append(v2, fi6)
                    }
                    switch "hello" {
                    case "foo bar":
                        if (-2147483648 < -2147483648) {
                            var v12 int8 = 42
                            v7 += v7
                            fmt.Println(v1, v3)
                            _ = v12
                        } else {
                            v8 -= v8
                            v13 := GenS1{F0: v1}
                            var v14 int8 = 127
                            _ = v13
                            _ = v14
                        }
                    case "":
                        var v15 bool = true
                        _ = v15
                    case "\n\t":
                        var v16 uint16 = 255
                        _ = v16
                    default:
                        fmt.Println(v8, v8)
                    }
                    var v17 bool = true
                    v3 += (v3 % ((v3) | 1))
                    _ = v17
                }
                switch "foo bar" {
                case "\n\t":
                    var v18 int8 = 0
                    _ = v18
                case "hello":
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(fi6, fi5, v8)
                    case "":
                        v8 -= 767
                    default:
                        v8 -= 255
                    }
                case "":
                    var v19 uint = v8
                    _ = v19
                }
                var v20 uint64 = 0
                v1 -= v1
                _ = v20
                _ = v7
                _ = v8
            } else {
                var v21 float32 = 3.14159
                var v22 uint16 = 2
                v1 -= 2147483647
                _ = v21
                _ = v22
            }
            _ = fi6
        }
        _ = fi5
    }
    if true {
        fmt.Println(v1, v1)
        var v23 uint64 = 46
        _ = v23
    }
    v3 += (v3 / (((v3 >> ((2) & 63))) | 1))
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    switch -9223372036854775808 {
    case 42:
        if (1e10 <= 3.14159) {
            if (!(!(false || true))) {
                var v24 int32 = -1
                if len(v2) > 0 {
                    v2[0] = 216
                }
                _ = v24
            } else {
                for fi25 := 0; fi25 < 6; fi25++ {
                    v26 := GenS1{F0: 2147483647}
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi27 := 0; fi27 < 20; fi27++ {
                            v28 := GenS2{F0: 18446744073709551615, F1: 419, F2: (false && (!(705 > 2)))}
                            v29 := GenS2{F0: 192, F1: 0, F2: ("hello" >= "a")}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    v1 -= 358
                    _ = fi25
                    _ = v26
                }
                v3 += (-(v3))
                if true {
                    var v30 uint8 = 128
                    fmt.Println(v3, v1)
                    _ = v30
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch 128 {
            case -237:
                var v31 int8 = -128
                _ = v31
            }
        }
    case 916:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 942:
        if false {
            switch "foo bar" {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                if (false && (!(!false))) {
                    v1 += (v1 | -2147483648)
                    v3 += (v3 | (v3 | 1000))
                }
            case "a":
                fmt.Println(v3, v3)
            }
        } else {
            var v32 uint8 = 90
            switch 867 {
            case 32767:
                v33 := GenS3{F0: 4294967295, F1: "foo bar", F2: v3}
                _ = v33
            case 127:
                v3 = v3
            }
            _ = v32
        }
    default:
        v3 -= (v3 >> ((2) & 63))
    }
    v34 := GenS1{F0: v1}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
