package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 string) int {
    for fi5 := 0; fi5 < 19; fi5++ {
        switch p0 {
        case "":
            if true {
                switch fi5 {
                case -771:
                    p0 = (p0 + "a")
                case 2147483647:
                    v6 := []int{}
                    _ = v6
                default:
                    p0 = (p0 + p0)
                }
                v7 := GenS1{F0: 255}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 string = "foo bar"
                    var v10 uint64 = 4294967295
                    v11 := []int{}
                    switch (fi8 & 127) {
                    case 519:
                        v12 := GenS2{F0: -183, F1: 54, F2: 837}
                        _ = v12
                    default:
                        p0 = "a"
                        v13 := GenS1{F0: 0}
                        fmt.Println(p0)
                        v14 := GenS1{F0: 2}
                        _ = v13
                        _ = v14
                    }
                    var v15 uint16 = 65535
                    if true {
                        v9 = p0
                    } else {
                        var v16 uint16 = (v15 + (v15 ^ (v15 / (((v15 ^ 65535)) | 1))))
                        var v17 bool = (!(-127 != -2147483648))
                        _ = v16
                        _ = v17
                    }
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v15
                    _ = v9
                }
                var v18 uint32 = 972
                _ = v18
            }
        case "hello":
            v19 := GenS3{F0: (!(-938 == 224))}
            _ = v19
        default:
            var v20 float64 = (float64((((0.1 * -2.25) - (1.5 / (float32(-2.25) - ((1e10 + 0.1) / ((1.5 + float32(0.0)) - (0.1 * (-2.25 * float32(0.0)))))))) / 1e10)) / 1.0)
            _ = v20
        }
        p0 = "hello"
        v21 := []int{(-(fi5)), fi5, 32767}
        _ = fi5
        _ = v21
    }
    v22 := GenS4{F0: 2147483648, F1: 2, F2: 1000}
    fmt.Println(p0)
    _ = p0
    _ = v22
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    var v3 int = 0
    var v4 uint8 = 223
    if false {
        fmt.Println(v3, v1)
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint64 = 18446744073709551615
            _ = fi23
            _ = v24
        }
        var v25 int8 = 31
        _ = v25
    } else {
        var v26 uint16 = v1
        _ = v26
    }
    v27 := GenS2{F0: 160, F1: 1, F2: (v1 % ((v1) | 1))}
    v28 := []int{}
    for fi29 := 0; fi29 < 9; fi29++ {
        _ = len(v2)
        _ = fi29
    }
    v30 := GenS2{F0: 0, F1: -6, F2: (v1 - 255)}
    v4 = (v4 / ((127) | 1))
    v31 := GenS1{F0: (v1 & v1)}
    for fi32 := 0; fi32 < 14; fi32++ {
        v1 += v1
        for fi33 := 0; fi33 < 17; fi33++ {
            switch "hello" {
            case "hello":
                var v34 int64 = 42
                _ = v34
            case "\n\t":
                var v35 float32 = 1e10
                _ = v35
            case "":
                v4 -= (v4 * 176)
            }
            fmt.Println(fi32, v3, v4)
            var v36 uint32 = 580
            _ = fi33
            _ = v36
        }
        _ = fi32
    }
    var v37 uint8 = 42
    for fi38 := 0; fi38 < 15; fi38++ {
        v1 += (v1 / ((v1) | 1))
        _ = fi38
    }
    fmt.Println(v3)
    var v39 uint = 18446744073709551615
    var v40 int8 = 1
    v28 = append(v28, 975)
    switch "\n\t" {
    case "\n\t":
        var v41 int16 = 42
        _ = v41
    }
    if (42 != int64(v40)) {
        v42 := GenS4{F0: 2147483648, F1: (v1 & 32767), F2: 1000}
        var v43 int16 = 0
        _ = v42
        _ = v43
    }
    var v44 float64 = 0.1
    var v45 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
}
