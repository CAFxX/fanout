package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float64 = float64((1e10 + -2.25))
    v2 := make(map[uint32]bool)
    v3 := []int{-326, 371, 4294967295}
    var v4 bool = true
    if (!v4) {
        var v5 float32 = 3.14159
        var v6 int8 = 127
        _ = v5
        _ = v6
    } else {
        if (!false) {
            fmt.Println(v1, v1)
            if len(v3) > 0 {
                v3[0] = 610
            }
            v2[4294967295] = (!(!(true || true)))
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint = 255
            v4 = (false || v4)
            _ = fi7
            _ = v8
        }
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4, v4)
    case "":
        if ("" >= ("a" + "\n\t")) {
            var v9 uint = 0
            _ = v9
        }
    case "\n\t":
        if false {
            v1 = v1
            switch 1 {
            case 0:
                switch "foo bar" {
                case "a":
                    if v4 {
                        v10 := GenS1{F0: 1, F1: 0}
                        _ = len(v2)
                        fmt.Println(v1, v4)
                        _ = v10
                    } else {
                        if (337 >= -32768) {
                            v11 := GenS1{F0: 1, F1: 127}
                            fmt.Println(v1)
                            _ = v11
                        }
                        v12 := GenS1{F0: 42, F1: -92}
                        var v13 uint8 = 1
                        _ = v12
                        _ = v13
                    }
                case "foo bar":
                    v1 += v1
                case "":
                    if (v1 != (0.1 + v1)) {
                        v2[596] = (!(2147483648 < 623))
                    }
                default:
                    v14 := GenS2{F0: 1}
                    _ = v14
                }
            case -910:
                fmt.Println(v4, v4)
            case 357:
                delete(v2, 2147483647)
            default:
                var v15 float64 = (v1 * v1)
                _ = v15
            }
            var v16 int32 = 2147483647
            _ = v16
        } else {
            delete(v2, 773)
            v4 = (!v4)
        }
    default:
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS2{F0: 0}
            if (!(42 > 779)) {
                var v19 bool = v4
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    }
    switch (("foo bar" + "") + "a") {
    case "\n\t":
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint16 = 42
            _ = len(v3)
            switch "" {
            case "foo bar":
                fmt.Println(fi20)
            case "":
                for fi22 := 0; fi22 < 18; fi22++ {
                    v4 = false
                    switch "a" {
                    case "\n\t":
                        var v23 bool = v4
                        _ = v23
                    case "a":
                        _ = len(v2)
                    case "hello":
                        v1 = (1e10 * v1)
                    }
                    if (v1 <= float64(-2.25)) {
                        v21 = (v21 % (((v21 | v21)) | 1))
                    } else {
                        var v24 bool = ((v1 < v1) || (-1 < -2147483648))
                        v2[282] = v4
                        _ = v24
                    }
                    _ = fi22
                }
            case "\n\t":
                v21 = v21
            default:
                v25 := GenS3{F0: 1}
                _ = v25
            }
            _ = fi20
            _ = v21
        }
    }
    var v26 uint64 = 0
    v1 = (v1 - v1)
    v27 := GenS2{F0: 127}
    fmt.Println(v26)
    v1 = v1
    if ((137 >= 2) && (-2147483648 != -216)) {
        var v28 uint64 = 1
        _ = v28
    } else {
        var v29 float32 = -2.25
        v2[1] = (true && true)
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        if (127 < 1) {
            var v31 float32 = 3.14159
            for fi32 := 0; fi32 < 16; fi32++ {
                if v4 {
                    var v33 int64 = -571
                    var v34 uint = 255
                    v2[2] = true
                    _ = v33
                    _ = v34
                }
                v3 = append(v3, (fi30 ^ (fi30 | fi32)))
                _ = fi32
            }
            var v35 uint16 = 2
            _ = v31
            _ = v35
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
