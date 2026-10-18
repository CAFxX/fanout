package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int; F2 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint = 656
    v4 := []string{"", ("foo bar" + "foo bar"), ("hello" + ""), ("foo bar" + ("" + "hello"))}
    v5 := []int{-2147483648}
    switch -2147483648 {
    case 573:
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi6, fi6)
            v3 = v3
            _ = fi6
        }
    case 42:
        v7 := GenS1{F0: 42, F1: 1000}
        _ = v7
    case 32768:
        if (-2.25 <= 0.0) {
            v3 = v3
            var v8 uint16 = 2
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (2 < 128) {
                    switch 128 {
                    case -886:
                        v11 := GenS2{F0: -32768, F1: 2147483647, F2: -1}
                        _ = v11
                    case 1:
                        var v12 uint32 = 179
                        _ = v12
                    default:
                        v13 := GenS2{F0: -1, F1: 2147483647, F2: -450}
                        _ = v13
                    }
                    if (0 == 1) {
                        if (!false) {
                            fmt.Println(fi10, fi10)
                        }
                        if (-1 <= 42) {
                            var v14 uint8 = 255
                            v14 -= (v14 >> (((v3 << (((v3 / (((v3 | 830)) | 1))) & 63))) & 63))
                            var v15 int32 = 42
                            _ = v14
                            _ = v15
                        } else {
                            v3 = 1
                            v3 += 216
                            v3 -= (v3 + v3)
                            fmt.Println(fi9, fi9, fi10)
                            var v16 int8 = 127
                            _ = v16
                        }
                        var v17 float32 = float32(((((float64((1.5 * (1e10 / 1.5))) + 0.0) + 0.1) - (-2.25 * 0.0)) / 3.14159))
                        _ = v17
                    } else {
                        var v18 uint = (v3 ^ 255)
                        v4 = append(v4, "foo bar")
                        v4 = append(v4, "\n\t")
                        _ = v18
                    }
                    if (false || false) {
                        var v19 int16 = 0
                        _ = v19
                    } else {
                        if false {
                            fmt.Println(fi9, fi9)
                            switch fi10 {
                            case 699:
                                if len(v5) > 0 {
                                    v5[0] = (fi10 >> ((0) & 63))
                                }
                            case -9223372036854775808:
                                v1[fi10] = "a"
                            case -1:
                                v3 += v3
                            }
                        }
                        var v20 uint16 = 0
                        v2[("hello" + (("" + ("\n\t" + (("" + ("hello" + "")) + "hello"))) + "\n\t"))] = true
                        _ = v20
                    }
                }
                delete(v1, -289)
                var v21 float64 = ((((1e10 * float64(-2.25)) * -2.25) - 1e10) * 0.0)
                _ = fi10
                _ = v21
            }
            v3 = (v3 >> ((v3) & 63))
            _ = fi9
        }
    }
    var v22 uint32 = 2147483648
    var v23 float64 = 0.1
    var v24 uint8 = 0
    v22 = v22
    var v25 uint32 = v22
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v26 uint64 = 42
    fmt.Println(v25, v23, v26)
    var v27 float32 = ((1.5 * 0.0) - 1e10)
    v28 := GenS3{F0: 42}
    _ = len(v1)
    switch -203 {
    case 4294967295:
        fmt.Println(v26, v23, v22)
    case 32767:
        var v29 uint64 = (v26 | 42)
        _ = v29
    case 128:
        if (false && (!false)) {
            var v30 int = -1
            v22 += uint32(v26)
            var v31 uint8 = (v24 + (v24 >> ((937) & 63)))
            _ = v30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 9; fi32++ {
            v26 = 2
            v33 := GenS4{F0: ((v25 & (v22 % ((105) | 1))) <= (v22 | 2)), F1: fi32, F2: -1}
            v34 := GenS4{F0: true, F1: 0, F2: 0}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    v25 = 2
    var v35 uint8 = (v24 + (v24 % (((v24 / (((v24 % ((127) | 1))) | 1))) | 1)))
    v36 := GenS4{F0: true, F1: 4294967295, F2: 586}
    var v37 uint8 = v35
    if true {
        _ = len(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
