package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 bool; F2 int8 }

func main() {
    v1 := []int{}
    var v2 int8 = -128
    var v3 uint8 = 42
    v4 := make(map[int]int)
    switch ("" + ("foo bar" + "foo bar")) {
    case "foo bar":
        var v5 uint32 = 508
        _ = v5
    }
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 11; fi7++ {
        switch (fi7 & 128) {
        case 128:
            var v8 string = ((("foo bar" + ("\n\t" + "hello")) + "a") + "")
            _ = v8
        case 1:
            switch "\n\t" {
            case "foo bar":
                _ = len(v4)
            case "":
                v2 += -53
            }
        default:
            switch "hello" {
            case "\n\t":
                switch "" {
                case "hello":
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 int32 = 1
                        _ = fi9
                        _ = v10
                    }
                case "foo bar":
                    delete(v4, fi7)
                case "":
                    v11 := GenS1{F0: 42, F1: 1000}
                    _ = v11
                }
            case "hello":
                switch "foo bar" {
                case "a":
                    var v12 bool = true
                    _ = v12
                }
            case "a":
                if (55 <= 0) {
                    v13 := GenS1{F0: 652, F1: -1}
                    _ = v13
                } else {
                    v14 := GenS1{F0: 1, F1: -32768}
                    for fi15 := 0; fi15 < 15; fi15++ {
                        if (42 > 720) {
                            var v16 int16 = 42
                            _ = v16
                        } else {
                            v3 += v3
                            switch (fi7 >> ((255) & 63)) {
                            case 53:
                                fmt.Println(v6, v2, v3)
                            }
                            if (fi7 <= fi7) {
                                var v17 int64 = 0
                                if false {
                                    v2 -= 127
                                    v4[(^(fi15))] = (-(fi15))
                                }
                                _ = v17
                            }
                        }
                        if true {
                            v1 = append(v1, fi15)
                            var v18 float64 = (0.0 + (((3.14159 * 3.14159) / float64(0.1)) + 1e10))
                            _ = v18
                        }
                        if len(v1) > 0 {
                            v1[0] = fi7
                        }
                        _ = fi15
                    }
                    v19 := GenS2{F0: 0, F1: v6, F2: "a"}
                    _ = v14
                    _ = v19
                }
            default:
                v6 = v6
            }
        }
        _ = fi7
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    var v20 int8 = (v2 + 60)
    var v21 uint = 18446744073709551615
    if (false || (2 < (v6 % (((v6 + v6)) | 1)))) {
        v22 := GenS2{F0: 1, F1: 999, F2: "\n\t"}
        v23 := GenS3{F0: 32767, F1: (4294967295 >= 313), F2: -1}
        var v24 uint32 = (v6 % (((^(v6))) | 1))
        _ = v22
        _ = v23
        _ = v24
    }
    _ = len(v4)
    for fi25 := 0; fi25 < 20; fi25++ {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = fi25
    }
    var v26 uint8 = v3
    switch 9223372036854775807 {
    case 0:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 32767:
        var v27 uint32 = (v6 + v6)
        _ = v27
    case 168:
        var v28 float32 = 1.5
        _ = v28
    }
    v20 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
