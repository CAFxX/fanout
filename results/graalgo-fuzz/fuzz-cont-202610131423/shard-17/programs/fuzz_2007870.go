package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint32; F1 bool }
func genFn0(p0 uint8, p1 string, p2 uint16) uint64 {
    var v7 int32 = 42
    var v8 int32 = 20
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int8, p1 int, p2 int) int64 {
    fmt.Println(p1, p2, p1)
    v9 := GenS1{F0: p2, F1: p1, F2: p0}
    switch ("a" + ("\n\t" + "hello")) {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p1, p2)
        case "hello":
            var v10 int32 = 557
            _ = v10
        case "":
            var v11 float64 = 0.1
            _ = v11
        }
    case "\n\t":
        if (!false) {
            p2 += p1
        } else {
            p0 = -128
            var v12 int64 = -1
            v13 := []int{-555, p2, (p2 ^ p1)}
            _ = v12
            _ = v13
        }
    case "":
        v14 := GenS1{F0: p2, F1: 32767, F2: p0}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 float32 = (1.5 / ((1.5 - (((float32(0.0) / (1.5 + 0.1)) + 1e10) + 0.0)) - 3.14159))
    var v2 bool = false
    var v3 string = "hello"
    v4 := make(map[uint32]bool)
    v5 := []uint32{0, 971, 0, 65535}
    v6 := make(map[int]string)
    v15 := GenS1{F0: 32768, F1: 446, F2: -15}
    fmt.Println(v3)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float64 = 0.1
        v4[2] = false
        var v18 uint8 = 127
        _ = fi16
        _ = v17
        _ = v18
    }
    switch "a" {
    case "foo bar":
        var v19 int8 = 100
        _ = v19
    }
    v2 = (-890 == -529)
    v20 := GenS1{F0: 929, F1: -2147483648, F2: 0}
    v21 := v4[42]
    switch 42 {
    case 247:
        var v22 uint8 = 128
        _ = v22
    default:
        var v23 int8 = 0
        _ = v23
    }
    if true {
        fmt.Println(v1)
    }
    switch -9223372036854775808 {
    case 128:
        if false {
            switch (("a" + v3) + (v3 + ("\n\t" + v3))) {
            case "foo bar":
                fmt.Println(v21, v21)
            case "a":
                for fi24 := 0; fi24 < 15; fi24++ {
                    if v2 {
                        v2 = (!v21)
                        v3 = v3
                        switch v3 {
                        case "hello":
                            var v25 uint32 = 65535
                            _ = v25
                        case "":
                            v3 = "\n\t"
                        case "\n\t":
                            var v26 int8 = 100
                            _ = v26
                        default:
                            var v27 float32 = ((((0.0 - ((3.14159 * v1) * v1)) * 1e10) + (v1 * v1)) - 0.0)
                            _ = v27
                        }
                    }
                    if v2 {
                        v2 = v2
                        var v28 uint64 = 4294967295
                        _ = v28
                    }
                    if v21 {
                        if false {
                            if true {
                                v29 := GenS2{F0: 1, F1: 1, F2: 1}
                                var v30 int = (fi24 / (((fi24 | 1)) | 1))
                                _ = v29
                                _ = v30
                            }
                            v3 = ""
                        }
                        var v31 int16 = -32768
                        var v32 string = v3
                        _ = v31
                        _ = v32
                    }
                    _ = fi24
                }
            case "hello":
                switch 1 {
                case -9223372036854775808:
                    if true {
                        _ = len(v6)
                    }
                case 127:
                    v33 := GenS3{F0: 1}
                    _ = v33
                case 2147483647:
                    if v2 {
                        if (42 >= 760) {
                            v34 := GenS4{F0: v3, F1: 42}
                            var v35 uint16 = 652
                            _ = v34
                            _ = v35
                        } else {
                            v2 = (1 != 42)
                            v36 := GenS2{F0: 127, F1: 217, F2: 793}
                            var v37 uint8 = 0
                            _ = v36
                            _ = v37
                        }
                        fmt.Println(v3, v3)
                        v38 := GenS1{F0: 0, F1: 2147483647, F2: -1}
                        _ = v38
                    } else {
                        v39 := GenS5{F0: 0, F1: 4294967295}
                        _ = v39
                    }
                }
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                v41 := GenS6{F0: 0, F1: (!v2)}
                v21 = (1.5 > (0.0 - 1e10))
                _ = fi40
                _ = v41
            }
        }
    case 9223372036854775807:
        v21 = true
    case 464:
        var v42 int32 = 2147483647
        _ = v42
    default:
        if v21 {
            switch v3 {
            case "hello":
                v2 = false
            case "a":
                var v43 int8 = -128
                _ = v43
            default:
                _ = len(v5)
            }
            for fi44 := 0; fi44 < 1; fi44++ {
                var v45 uint8 = 173
                var v46 bool = true
                _ = fi44
                _ = v45
                _ = v46
            }
        }
    }
    var v47 int8 = 100
    v48 := v4[2147483648]
    fmt.Println(v1, v48, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
