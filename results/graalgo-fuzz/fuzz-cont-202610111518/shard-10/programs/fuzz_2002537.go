package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int }
func genFn0(p0 bool) uint {
    v5 := GenS1{F0: 65535, F1: -2147483648, F2: 2}
    switch "foo bar" {
    case "\n\t":
        v6 := []int{4294967295, 1, -348}
        _ = v6
    case "foo bar":
        if false {
            var v7 int64 = -9223372036854775808
            p0 = p0
            _ = v7
        }
    case "hello":
        var v8 float32 = -2.25
        _ = v8
    }
    _ = p0
    _ = v5
    return 266
}
func genFn1(p0 int16) uint8 {
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := []int{(fi9 & (fi9 ^ (fi9 + fi9)))}
        _ = fi9
        _ = v10
    }
    p0 = (^(p0))
    v11 := []int{-410, 4294967295, 9223372036854775807}
    v12 := GenS1{F0: 2, F1: -65, F2: 4294967295}
    switch "a" {
    case "\n\t":
        p0 -= 0
    case "a":
        switch -50 {
        case 0:
            p0 += (^(p0))
        default:
            v11 = append(v11, 127)
        }
    case "hello":
        p0 = p0
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS1{F0: 286, F1: -710, F2: 18446744073709551615}
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 string = "a"
    var v2 int64 = 42
    v3 := []uint32{2147483647, 65535}
    v4 := make(map[int]uint32)
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            if (true && (-2147483648 >= -2147483648)) {
                v17 := GenS2{F0: 42, F1: fi16, F2: -129}
                _ = v17
            } else {
                switch v1 {
                case "a":
                    var v18 string = ((v1 + v1) + "")
                    _ = v18
                }
                fmt.Println(v2, fi15)
                v19 := GenS1{F0: 255, F1: -1, F2: 0}
                var v20 bool = (!((false || (492 != 2147483648)) || (18446744073709551615 == 4294967295)))
                _ = len(v3)
                _ = v19
                _ = v20
            }
            _ = fi16
        }
        _ = fi15
    }
    if false {
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS2{F0: 1, F1: (fi21 - fi21), F2: (fi21 << ((255) & 63))}
            _ = fi21
            _ = v22
        }
        v2 = 503
    }
    v1 = (v1 + v1)
    if false {
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                switch fi24 {
                case 2147483647:
                    var v25 bool = (false && (-1 != -989))
                    _ = v25
                case -2147483648:
                    switch v1 {
                    case "foo bar":
                        if (fi24 >= (fi24 >> ((2) & 63))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v26 := GenS2{F0: 2147483648, F1: fi24, F2: (^(fi23))}
                            _ = len(v3)
                            _ = v26
                        }
                    default:
                        var v27 float32 = float32((-2.25 - 0.1))
                        _ = v27
                    }
                case 32767:
                    var v28 bool = true
                    _ = v28
                default:
                    var v29 int8 = 116
                    _ = v29
                }
                _ = fi24
            }
            var v30 float32 = float32(1e10)
            v1 = v1
            _ = fi23
            _ = v30
        }
    }
    switch v1 {
    case "hello":
        v31 := GenS1{F0: 0, F1: -1000, F2: 2}
        _ = v31
    case "foo bar":
        _ = len(v4)
    }
    fmt.Println(v2)
    switch (((v1 + (((v1 + v1) + v1) + v1)) + (v1 + "")) + v1) {
    case "":
        var v32 uint32 = 0
        _ = v32
    case "hello":
        switch 0 {
        case 356:
            var v33 uint8 = 0
            _ = v33
        }
    default:
        switch "a" {
        case "":
            var v34 int8 = 1
            _ = v34
        default:
            v1 = ""
        }
    }
    fmt.Println(v1, v1)
    if (!false) {
        fmt.Println(v1, v1)
        v2 -= 614
    }
    var v35 bool = true
    var v36 int16 = 1
    var v37 string = v1
    if (0.0 > (0.0 / (1.5 * 3.14159))) {
        v37 = v1
        var v38 uint16 = 255
        v37 = (v1 + v1)
        _ = v38
    } else {
        var v39 int16 = v36
        _ = len(v3)
        _ = v39
    }
    var v40 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
}
