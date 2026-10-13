package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint32 }
func genFn0() string {
    v6 := GenS1{F0: -9223372036854775808, F1: 15, F2: 1}
    v7 := GenS2{F0: 2147483647}
    var v8 uint = 27
    if (0.0 >= 3.14159) {
        v9 := []int{}
        _ = v9
    }
    fmt.Println(v8)
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 uint = 512
    v2 := []uint32{671, 663}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    v5 := []int{32767, 1, 127, 149}
    fmt.Println(v4, v1)
    if (true && true) {
        switch -687 {
        case 0:
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int32 = -960
                var v12 int8 = int8(v1)
                _ = fi10
                _ = v11
                _ = v12
            }
        case 128:
            delete(v3, 361)
        case -1:
            _ = len(v3)
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 string = "a"
                switch v14 {
                case "hello":
                    v5 = append(v5, (fi13 & fi13))
                case "a":
                    switch fi13 {
                    case 42:
                        v15 := GenS3{F0: -2147483648, F1: 0}
                        _ = v15
                    case 0:
                        switch (v14 + (v14 + ((v14 + v14) + (((("\n\t" + v14) + v14) + ("\n\t" + v14)) + "hello")))) {
                        case "foo bar":
                            delete(v3, 2147483647)
                        case "\n\t":
                            fmt.Println(v4)
                        default:
                            if true {
                                for fi16 := 0; fi16 < 6; fi16++ {
                                    v14 = "hello"
                                    _ = fi16
                                }
                                v1 -= (v1 * v1)
                                v17 := GenS1{F0: (-(fi13)), F1: 26, F2: 2147483647}
                                _ = v17
                            }
                        }
                    case -95:
                        if true {
                            fmt.Println(v4, v14, fi13)
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                var v18 int = (fi13 | 306)
                                _ = v18
                            case "":
                                v19 := GenS1{F0: (fi13 << ((269) & 63)), F1: uint8(v4), F2: 32768}
                                _ = v19
                            case "foo bar":
                                v14 = (v14 + v14)
                            default:
                                var v20 int = int(v4)
                                _ = v20
                            }
                            _ = len(v3)
                        }
                    }
                case "\n\t":
                    switch (v14 + "foo bar") {
                    case "hello":
                        v1 = 634
                    case "a":
                        v4 = (v4 + v4)
                    default:
                        v1 = 1
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS1{F0: -635, F1: 42, F2: int(v4)}
            for fi23 := 0; fi23 < 8; fi23++ {
                v1 += (v1 & 0)
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v4, v1)
    } else {
        v24 := GenS2{F0: 2147483648}
        _ = v24
    }
    var v25 uint8 = uint8(v4)
    var v26 int = 32767
    var v27 string = "\n\t"
    v28 := GenS4{F0: (^(v25))}
    var v29 int64 = 1
    for fi30 := 0; fi30 < 18; fi30++ {
        if false {
            if len(v5) > 0 {
                v5[0] = v26
            }
            var v31 uint32 = 0
            _ = v31
        }
        _ = fi30
    }
    v4 = uint64(v25)
    for fi32 := 0; fi32 < 4; fi32++ {
        switch v27 {
        case "":
            var v33 int32 = 2147483647
            _ = v33
        case "foo bar":
            fmt.Println(v27)
        case "hello":
            for fi34 := 0; fi34 < 13; fi34++ {
                v4 += (v4 - (v4 | (v4 ^ v4)))
                v1 = (^(v1))
                _ = fi34
            }
        default:
            if (42 > 1) {
                v25 = v25
                for fi35 := 0; fi35 < 19; fi35++ {
                    fmt.Println(v29, v1, v4)
                    v36 := GenS5{F0: -119}
                    _ = fi35
                    _ = v36
                }
                var v37 string = v27
                _ = v37
            }
        }
        _ = fi32
    }
    switch v26 {
    case 0:
        switch 32768 {
        case 361:
            fmt.Println(v29, v27)
        case 333:
            var v38 bool = (!false)
            _ = v38
        default:
            var v39 uint = v1
            _ = v39
        }
    case -2147483648:
        var v40 int8 = -128
        _ = v40
    case 9223372036854775807:
        v26 -= v26
    }
    v41 := GenS4{F0: 255}
    fmt.Println(v4, v29)
    if true {
        var v42 uint64 = (v4 % ((uint64(v25)) | 1))
        _ = v42
    } else {
        v43 := GenS6{F0: -416, F1: uint32(v1)}
        _ = v43
    }
    var v44 float32 = (-2.25 + 3.14159)
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v5
}
