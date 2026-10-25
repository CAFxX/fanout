package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    fmt.Println(v1)
    switch 1 {
    case -1:
        if (("hello" + ("a" + "foo bar")) >= ("foo bar" + ("a" + "a"))) {
            fmt.Println(v1)
            switch 536 {
            case 32767:
                if false {
                    v4 := GenS1{F0: 294, F1: 2, F2: -2147483648}
                    v1 += (float32((1e10 * 0.0)) + (1e10 - v1))
                    _ = v4
                }
            case -2147483648:
                v5 := []int{359, 4294967295}
                _ = v5
            case 128:
                switch 32768 {
                case -9223372036854775808:
                    var v6 float32 = -2.25
                    _ = v6
                case 143:
                    v1 = 1.5
                case 128:
                    var v7 int64 = -603
                    _ = v7
                }
            }
        }
    case 4294967295:
        var v8 uint8 = 127
        _ = v8
    case -843:
        if (2 == 2) {
            for fi9 := 0; fi9 < 7; fi9++ {
                _ = len(v3)
                var v10 float32 = ((float32(1e10) - float32(float64(v1))) - v1)
                _ = fi9
                _ = v10
            }
            switch "" {
            case "":
                var v11 uint8 = 0
                _ = v11
            }
            v12 := GenS2{F0: 32768}
            _ = v12
        }
    }
    v3["hello"] = false
    if false {
        if ((168 != 1) && false) {
            v13 := []int{1, 0, -365}
            _ = v13
        } else {
            v14 := []int{-980}
            var v15 uint8 = 169
            if false {
                if false {
                    v16 := GenS2{F0: 2147483647}
                    v17 := GenS3{F0: 1}
                    _ = v16
                    _ = v17
                }
                switch -573 {
                case 230:
                    v15 += v15
                case -194:
                    switch -401 {
                    case 246:
                        var v18 int32 = 2147483647
                        _ = v18
                    case 4294967295:
                        v19 := GenS3{F0: 42}
                        _ = v19
                    case 128:
                        var v20 int32 = -3
                        _ = v20
                    default:
                        v15 += (v15 + v15)
                    }
                default:
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v1, v15)
                    case "":
                        var v21 uint32 = 4294967295
                        _ = v21
                    case "a":
                        v22 := GenS3{F0: -1}
                        _ = v22
                    }
                }
            } else {
                fmt.Println(v15)
                v15 = v15
            }
            _ = v14
            _ = v15
        }
    }
    v23 := GenS3{F0: 116}
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int32 = 1
        _ = fi24
        _ = v25
    }
    if (!(2147483647 <= 0)) {
        v26 := GenS3{F0: -128}
        if (65535 == 255) {
            _ = len(v2)
            if (false || true) {
                v1 += (1.5 * 1.5)
            } else {
                v1 -= 1.5
                switch "a" {
                case "\n\t":
                    delete(v2, -139)
                case "a":
                    var v27 string = "foo bar"
                    _ = v27
                default:
                    v28 := []int{386, 42, 542, 2147483647}
                    _ = v28
                }
                var v29 float32 = -2.25
                _ = v29
            }
        } else {
            v1 += v1
            v30 := GenS4{F0: -9223372036854775808}
            _ = v30
        }
        _ = v26
    } else {
        v3[("" + ("foo bar" + (("foo bar" + "") + ("foo bar" + "\n\t"))))] = false
    }
    var v31 uint64 = 18446744073709551615
    v31 = 1
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(fi32, v31)
        v31 += v31
        _ = fi32
    }
    v31 -= v31
    var v33 int = -436
    v34 := GenS1{F0: 753, F1: 65535, F2: -707}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v33
    _ = v34
}
