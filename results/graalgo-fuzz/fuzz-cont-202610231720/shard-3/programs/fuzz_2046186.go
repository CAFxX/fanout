package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }

func main() {
    var v1 string = "a"
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 252
    var v4 int16 = 840
    var v5 uint8 = 145
    var v6 uint = 18446744073709551615
    var v7 float32 = 0.0
    if true {
        fmt.Println(v1, v1)
    } else {
        if true {
            switch (v1 + v1) {
            case "\n\t":
                v8 := GenS1{F0: 42, F1: 2147483647, F2: -1}
                _ = v8
            case "":
                fmt.Println(v2, v5, v6)
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                if true {
                    if false {
                        v10 := []int{127, (fi9 | fi9), (fi9 & 2147483647), 127}
                        if (!(!(1 >= 0))) {
                            switch fi9 {
                            case 127:
                                for fi11 := 0; fi11 < 1; fi11++ {
                                    fmt.Println(fi9)
                                    fmt.Println(fi9, fi11)
                                    var v12 bool = (((3.14159 - (1.5 - (((0.1 / ((float64(v7) * 3.14159) * 1.5)) + ((((float64(v7) * 1e10) * float64(((v7 - v7) * v7))) * float64(float32(1e10))) - 3.14159)) / 1.5))) + 0.0) > -2.25)
                                    _ = fi11
                                    _ = v12
                                }
                            case -1:
                                _ = len(v10)
                            case 32768:
                                var v13 string = (v1 + "\n\t")
                                _ = v13
                            }
                            if len(v10) > 0 {
                                v10[0] = (-(fi9))
                            }
                            v7 -= v7
                        }
                        var v14 int8 = 1
                        _ = v10
                        _ = v14
                    }
                }
                _ = fi9
            }
        } else {
            v15 := []int{344, 9223372036854775807, int(v5)}
            switch "foo bar" {
            case "foo bar":
                v4 = v4
            default:
                for fi16 := 0; fi16 < 4; fi16++ {
                    fmt.Println(fi16, v1, v2)
                    _ = fi16
                }
            }
            _ = v15
        }
        v5 = (v5 % ((2) | 1))
    }
    switch 0 {
    case 32768:
        var v17 uint64 = 132
        _ = v17
    case -9223372036854775808:
        if (!false) {
            fmt.Println(v5, v2)
        }
    case 2147483647:
        v18 := GenS1{F0: 246, F1: 0, F2: -83}
        _ = v18
    default:
        v6 += (v6 % ((18446744073709551615) | 1))
    }
    v19 := GenS1{F0: 202, F1: 2147483647, F2: 1}
    fmt.Println(v3)
    var v20 uint = (v6 % ((uint(v5)) | 1))
    fmt.Println(v2)
    var v21 uint32 = 1
    v3 = 255
    var v22 string = v1
    fmt.Println(v22, v2, v21)
    if ((v20 & (v6 ^ (^(v20)))) >= (v20 & 0)) {
        if false {
            fmt.Println(v22, v7, v22)
            var v23 uint32 = uint32(v5)
            _ = v23
        } else {
            if (0.1 != ((0.0 / (1.5 * (float64(-2.25) + -2.25))) * 1.5)) {
                v24 := []int{}
                v2 = (v2 + 42)
                v25 := GenS1{F0: 2147483647, F1: -899, F2: 42}
                _ = v24
                _ = v25
            }
            v3 = (v3 >> (((^(v20))) & 63))
            if (!(v7 > (-2.25 * (((v7 * 0.0) - v7) - (v7 + (v7 - (-2.25 + v7))))))) {
                var v26 uint16 = v3
                _ = v26
            }
        }
        switch 128 {
        case -383:
            if (18446744073709551615 == (v20 - v20)) {
                var v27 int64 = -9223372036854775808
                v3 = (v3 / ((v3) | 1))
                for fi28 := 0; fi28 < 19; fi28++ {
                    if (true && true) {
                        v29 := []int{fi28, 4294967295, fi28, fi28}
                        var v30 uint16 = 662
                        fmt.Println(v6, v27, v4)
                        _ = v29
                        _ = v30
                    } else {
                        v27 -= (v27 >> (((v20 + uint(v3))) & 63))
                        var v31 float32 = v7
                        v2 -= v27
                        _ = v31
                    }
                    var v32 int64 = -836
                    var v33 uint64 = 4294967295
                    _ = fi28
                    _ = v32
                    _ = v33
                }
                _ = v27
            }
        case 0:
            v34 := []int{}
            _ = v34
        default:
            v35 := []int{}
            _ = v35
        }
        v36 := []int{128, 32, 128}
        _ = v36
    } else {
        v6 -= v6
    }
    switch "foo bar" {
    case "\n\t":
        var v37 bool = (-547 != 42)
        _ = v37
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        switch v1 {
        case "a":
            v2 += (v2 >> ((v20) & 63))
        default:
            var v39 int32 = 2147483647
            _ = v39
        }
        _ = fi38
    }
    v40 := []int{-2147483648, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
