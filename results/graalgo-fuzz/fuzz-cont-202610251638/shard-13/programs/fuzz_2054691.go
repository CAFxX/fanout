package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 32767
    var v2 bool = true
    var v3 uint32 = 2
    v4 := []int{9223372036854775807}
    var v5 uint = 414
    for fi6 := 0; fi6 < 6; fi6++ {
        v4 = append(v4, (fi6 >> ((uint(v3)) & 63)))
        _ = fi6
    }
    var v7 uint = (v5 & 255)
    v2 = v2
    v5 -= v7
    var v8 uint16 = (v1 << (((v7 >> ((229) & 63))) & 63))
    v5 -= (v5 / (((^(v7))) | 1))
    var v9 int8 = -1
    v10 := GenS1{F0: 42, F1: 1000}
    v11 := GenS1{F0: -2147483648, F1: 32767}
    for fi12 := 0; fi12 < 9; fi12++ {
        switch fi12 {
        case -941:
            var v13 int16 = 32767
            _ = v13
        case 9223372036854775807:
            switch "foo bar" {
            case "hello":
                var v14 int64 = 1
                _ = v14
            }
        default:
            fmt.Println(v8, fi12, v5)
        }
        if true {
            v3 += v3
            v2 = v2
            if v2 {
                v8 -= v8
            } else {
                fmt.Println(v5, v1)
                switch "\n\t" {
                case "a":
                    var v15 int16 = 42
                    _ = v15
                case "\n\t":
                    for fi16 := 0; fi16 < 14; fi16++ {
                        switch (fi12 & (^(fi16))) {
                        case 259:
                            var v17 float64 = 3.14159
                            _ = v17
                        case 800:
                            fmt.Println(fi16, v3)
                        case -1:
                            switch "foo bar" {
                            case "foo bar":
                                v18 := GenS1{F0: -514, F1: 32767}
                                _ = v18
                            case "\n\t":
                                var v19 int64 = -9223372036854775808
                                _ = v19
                            default:
                                var v20 int32 = int32(fi12)
                                v8 = (v1 ^ v1)
                                var v21 uint64 = uint64(v7)
                                _ = v20
                                _ = v21
                            }
                        default:
                            if false {
                                v22 := GenS2{F0: 42, F1: (-(v9)), F2: 22}
                                v23 := GenS3{F0: v8}
                                if (!((v2 || (-2.25 >= (float32(0.0) * (3.14159 * ((3.14159 / (1.5 / 3.14159)) + 3.14159))))) && (0 < 128))) {
                                    fmt.Println(v1, v3)
                                }
                                v24 := GenS4{F0: 679}
                                _ = len(v4)
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                        switch (-(fi12)) {
                        case -1:
                            v1 -= 744
                        case 1:
                            v4 = append(v4, fi12)
                        }
                        _ = fi16
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        var v25 uint8 = 42
                        _ = v25
                    }
                default:
                    v7 = 1
                }
            }
        }
        _ = len(v4)
        _ = fi12
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if (0 < 819) {
            if false {
                var v27 int32 = -609
                var v28 uint = (v5 << ((2) & 63))
                var v29 string = ("foo bar" + (((("a" + "hello") + ("hello" + "hello")) + ((("\n\t" + "\n\t") + ("foo bar" + "\n\t")) + "\n\t")) + "foo bar"))
                _ = v27
                _ = v28
                _ = v29
            }
            v30 := GenS2{F0: 42, F1: (v9 + (v9 << ((v5) & 63))), F2: 0}
            _ = v30
        }
        var v31 int32 = 868
        fmt.Println(v31, fi26, v31)
        _ = fi26
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
