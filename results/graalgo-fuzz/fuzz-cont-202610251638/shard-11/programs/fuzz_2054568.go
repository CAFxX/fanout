package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "hello"
    v3 := []int{579, 9223372036854775807, 32768}
    var v4 uint16 = 255
    var v5 uint8 = 255
    fmt.Println(v4)
    v6 := GenS1{F0: 1, F1: -1}
    var v7 int64 = 1
    if ((v2 + "") != v2) {
        fmt.Println(v2, v4)
        delete(v1, 903)
    }
    switch "foo bar" {
    case "\n\t":
        v4 -= (v4 / ((2) | 1))
    default:
        v8 := GenS1{F0: 42, F1: 1}
        _ = v8
    }
    if true {
        v4 += v4
        for fi9 := 0; fi9 < 3; fi9++ {
            if (false && true) {
                v10 := GenS2{F0: (2147483648 <= 2147483647)}
                v5 += v5
                if (!true) {
                    var v11 int64 = v7
                    switch v2 {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        var v12 bool = true
                        _ = v12
                    }
                    if false {
                        v7 = (v7 + 372)
                        for fi13 := 0; fi13 < 16; fi13++ {
                            switch (fi13 ^ 555) {
                            case -89:
                                v4 -= (v4 ^ v4)
                            case -151:
                                fmt.Println(v5, fi9)
                            case 32768:
                                var v14 uint16 = (v4 / (((v4 / ((v4) | 1))) | 1))
                                _ = v14
                            }
                            var v15 int = fi13
                            v16 := GenS3{F0: fi13, F1: 18446744073709551615, F2: v5}
                            switch v2 {
                            case "\n\t":
                                fmt.Println(v2, v15, fi13)
                            }
                            if false {
                                v4 -= 42
                                _ = len(v1)
                                v5 = (v5 | 1)
                                if (2147483647 >= -342) {
                                    v11 = v11
                                    switch 42 {
                                    case -2147483648:
                                        v17 := GenS2{F0: ((!true) || ((((v4 - (v4 % (((v4 - v4)) | 1))) < (v4 * 0)) && ((-(v11)) <= 42)) && true))}
                                        _ = v17
                                    }
                                    fmt.Println(fi9, v5, v15)
                                }
                            }
                            _ = fi13
                            _ = v15
                            _ = v16
                        }
                        fmt.Println(v2, v5, fi9)
                    }
                    _ = v11
                } else {
                    var v18 float32 = 1e10
                    v2 = v2
                    _ = v18
                }
                _ = v10
            } else {
                if ((-2.25 / -2.25) < (float32(1.5) * 0.1)) {
                    v4 += (v4 & v4)
                    var v19 bool = ((v4 / ((32767) | 1)) >= v4)
                    v5 = v5
                    _ = v19
                } else {
                    var v20 string = (v2 + v2)
                    var v21 uint16 = 1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v20
                    _ = v21
                }
                v2 = v2
                v22 := GenS3{F0: -2147483648, F1: 902, F2: (^(v5))}
                _ = v22
            }
            _ = fi9
        }
        fmt.Println(v2, v5, v7)
    } else {
        v7 = (v7 + v7)
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            var v24 uint8 = (v5 << ((309) & 63))
            if (-2147483648 == 42) {
                if true {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 uint32 = 4294967295
                        v27 := v1[fi23]
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                } else {
                    _ = len(v3)
                }
            } else {
                var v28 int64 = 480
                _ = v28
            }
            _ = v24
        }
        if (v5 == v5) {
            fmt.Println(v4, v2)
        }
        _ = fi23
    }
    var v29 int16 = -32768
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := GenS1{F0: (v5 % ((v5) | 1)), F1: v29}
        _ = fi30
        _ = v31
    }
    switch (v2 + v2) {
    case "a":
        fmt.Println(v4, v29)
    default:
        var v32 uint32 = 4294967295
        _ = v32
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS3{F0: fi33, F1: 2, F2: (^(v5))}
        if false {
            fmt.Println(v29)
            v5 = v5
        } else {
            var v35 int64 = v7
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    if ("\n\t" <= v2) {
        switch -917 {
        case 849:
            switch -2147483648 {
            case -2147483648:
                var v36 uint64 = 0
                _ = v36
            case 42:
                switch 128 {
                case 9223372036854775807:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32767:
                    v4 = v4
                default:
                    var v37 int64 = (v7 >> ((581) & 63))
                    _ = v37
                }
            case 0:
                v38 := GenS2{F0: true}
                _ = v38
            }
        }
        v5 = v5
        v4 += 65535
    }
    if (2 > 42) {
        v29 = v29
        v4 = (v4 - v4)
        v39 := GenS2{F0: true}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
