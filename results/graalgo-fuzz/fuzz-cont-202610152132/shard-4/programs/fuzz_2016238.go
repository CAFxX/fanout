package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []string{"\n\t", "foo bar"}
    var v2 int16 = 42
    var v3 uint8 = 255
    var v4 string = "hello"
    var v5 uint16 = 841
    var v6 bool = false
    if v6 {
        v5 = v5
        _ = len(v1)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v1 = append(v1, "foo bar")
    var v7 int = 4294967295
    var v8 uint16 = (v5 | v5)
    for fi9 := 0; fi9 < 11; fi9++ {
        _ = len(v1)
        _ = len(v1)
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 0, F1: (^(v8)), F2: 2147483648}
            if len(v1) > 0 {
                v1[0] = ((v4 + (v4 + v4)) + "foo bar")
            }
            fmt.Println(v2, v8)
            v12 := GenS2{F0: (v5 * (^(v5))), F1: v4}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if (4294967295 == uint(v3)) {
        var v13 float32 = 1.5
        _ = v13
    } else {
        v8 = v5
    }
    v7 = 1
    var v14 uint16 = v8
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: (v5 >> ((566) & 63)), F1: "a"}
        fmt.Println(v5, fi15, v14)
        v17 := GenS2{F0: (v14 / (((v8 + v14)) | 1)), F1: (v4 + "a")}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint = 0
    var v19 int16 = 0
    v14 += 15
    v4 = (v4 + "a")
    if (v6 && v6) {
        if (v8 >= v5) {
            v20 := GenS1{F0: (v7 & (v7 / (((v7 | v7)) | 1))), F1: (^(v14)), F2: 2}
            v19 = v19
            _ = v20
        }
        switch v7 {
        case 128:
            switch v7 {
            case 127:
                _ = len(v1)
            case -2147483648:
                var v21 int16 = (v19 | 1)
                _ = v21
            }
        case -9223372036854775808:
            v22 := GenS3{F0: v18}
            _ = v22
        case 1:
            var v23 int8 = int8(v18)
            _ = v23
        default:
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := GenS3{F0: v18}
                if (v6 && v6) {
                    v6 = v6
                } else {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        if ((v6 && v6) || (!true)) {
                            v27 := GenS4{F0: 1, F1: (27 > 42), F2: 255}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v27
                        } else {
                            var v28 string = v4
                            _ = v28
                        }
                        fmt.Println(v14)
                        _ = fi26
                    }
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := GenS5{F0: -5}
                        v19 -= (v2 << ((v18) & 63))
                        _ = fi29
                        _ = v30
                    }
                }
                v18 = v18
                _ = fi24
                _ = v25
            }
        }
        switch "a" {
        case "hello":
            v2 -= v19
        case "":
            v1 = append(v1, v4)
        default:
            fmt.Println(v2, v4)
        }
    } else {
        v31 := GenS5{F0: 42}
        var v32 float64 = (1e10 + 0.1)
        fmt.Println(v6)
        _ = v31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
