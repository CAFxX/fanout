package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{108}
    var v2 int = 9223372036854775807
    var v3 int = (-(v2))
    v4 := []uint32{2, 0}
    var v5 int8 = -56
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = (v5 & (v5 * 102))
        _ = fi6
    }
    v7 := GenS1{F0: 2147483647, F1: "\n\t", F2: 1}
    for fi8 := 0; fi8 < 7; fi8++ {
        if (0 == 42) {
            fmt.Println(v5, v5, v5)
            v5 = 0
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                var v10 int16 = 32767
                if true {
                    v3 += (fi8 >> ((1) & 63))
                    if (42 >= 1) {
                        var v11 uint32 = 2147483648
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v11
                    } else {
                        var v12 uint16 = 610
                        _ = v12
                    }
                    switch fi8 {
                    case 32768:
                        v1 = append(v1, 0)
                    }
                    switch "hello" {
                    case "":
                        _ = len(v4)
                    case "a":
                        v2 -= v2
                    default:
                        v13 := GenS1{F0: 1, F1: "", F2: 772}
                        v1 = append(v1, 140)
                        _ = v13
                    }
                    var v14 uint8 = 164
                    var v15 uint64 = 0
                    v14 += v14
                    switch (("\n\t" + "a") + (("foo bar" + ("" + "foo bar")) + "a")) {
                    case "\n\t":
                        fmt.Println(v2, v5, fi8)
                    case "a":
                        _ = len(v1)
                    case "hello":
                        var v16 uint = 847
                        _ = v16
                    }
                    if true {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        if false {
                            var v17 uint8 = (v14 + v14)
                            _ = v17
                        } else {
                            v2 -= 0
                            fmt.Println(v14)
                        }
                        var v18 int64 = int64(v14)
                        var v19 uint64 = (v15 | v15)
                        v14 = 128
                        var v20 int8 = int8(v18)
                        v4 = append(v4, 1)
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi8)
                    }
                    _ = v14
                    _ = v15
                } else {
                    switch ("" + "\n\t") {
                    case "foo bar":
                        v21 := GenS2{F0: 127, F1: (((v10 < (^(v10))) || false) && true)}
                        _ = v21
                    case "\n\t":
                        fmt.Println(v3, fi9)
                    case "a":
                        var v22 uint8 = 0
                        _ = v22
                    default:
                        v5 += (-(v5))
                    }
                    var v23 uint8 = 2
                    var v24 int = fi9
                    _ = v23
                    _ = v24
                }
                _ = v10
            }
            switch "" {
            case "\n\t":
                _ = len(v1)
            case "hello":
                v3 -= fi9
            case "a":
                switch fi9 {
                case 127:
                    v3 += (fi9 % (((fi8 + 4294967295)) | 1))
                default:
                    fmt.Println(v5, fi8)
                }
            }
            v4 = append(v4, 65535)
            _ = fi9
        }
        if false {
            var v25 uint32 = 918
            var v26 int = fi8
            _ = v25
            _ = v26
        }
        _ = fi8
    }
    fmt.Println(v3, v2, v3)
    v5 = (v5 % (((v5 >> ((42) & 63))) | 1))
    fmt.Println(v5, v2)
    v3 = (-(v2))
    v3 += 0
    v2 = (v3 + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
