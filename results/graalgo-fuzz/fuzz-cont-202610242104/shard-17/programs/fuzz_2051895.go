package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    var v3 uint = 493
    var v4 float32 = 1.5
    v3 = 255
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 = v4
        if ((fi5 / ((fi5) | 1)) <= fi5) {
            if (false && false) {
                if (9223372036854775807 > 1) {
                    for fi6 := 0; fi6 < 13; fi6++ {
                        v4 = ((v4 * 0.0) * 0.0)
                        if (-194 <= -2147483648) {
                            switch (-(fi6)) {
                            case 1:
                                fmt.Println(fi6, v4, fi6)
                            default:
                                switch 127 {
                                case -1:
                                    var v7 uint = (v3 % (((v3 >> (((v3 - 135)) & 63))) | 1))
                                    _ = v7
                                }
                            }
                            delete(v2, fi5)
                        }
                        var v8 uint = v3
                        _ = fi6
                        _ = v8
                    }
                }
            }
        } else {
            switch (fi5 * fi5) {
            case 100:
                v9 := GenS1{F0: v3, F1: fi5, F2: (42 < -1)}
                _ = v9
            case -9223372036854775808:
                var v10 int16 = int16(v3)
                _ = v10
            case -2147483648:
                switch (fi5 + (fi5 * fi5)) {
                case -965:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v3 += v3
                        v12 := v2[fi11]
                        if true {
                            _ = len(v2)
                            fmt.Println(v3)
                        }
                        switch 331 {
                        case -9223372036854775808:
                            var v13 uint8 = 128
                            _ = v13
                        case 1:
                            v4 += v4
                        case 42:
                            if true {
                                v14 := GenS1{F0: (^(v3)), F1: fi11, F2: true}
                                v15 := GenS2{F0: 1, F1: uint16(v3), F2: -1}
                                v12 += 267
                                _ = v14
                                _ = v15
                            } else {
                                _ = len(v2)
                            }
                        default:
                            var v16 int16 = 42
                            _ = v16
                        }
                        _ = len(v2)
                        if (!false) {
                            var v17 int8 = 42
                            switch fi11 {
                            case 1:
                                var v18 bool = false
                                _ = v18
                            case -1:
                                var v19 string = ("" + "\n\t")
                                _ = v19
                            }
                            var v20 float32 = float32(1e10)
                            _ = v17
                            _ = v20
                        }
                        v4 = v4
                        var v21 float64 = float64(((v4 + 0.1) / 1.0))
                        _ = fi11
                        _ = v12
                        _ = v21
                    }
                case 712:
                    v4 += -2.25
                case 128:
                    fmt.Println(v1)
                default:
                    _ = len(v2)
                }
            default:
                if (847 > 826) {
                    fmt.Println(v4, fi5)
                    var v22 float32 = float32(-2.25)
                    fmt.Println(v1, v4)
                    _ = v22
                } else {
                    v23 := GenS1{F0: 418, F1: fi5, F2: true}
                    _ = v23
                }
            }
        }
        _ = fi5
    }
    v24 := GenS2{F0: 42, F1: uint16(v3), F2: 9223372036854775807}
    v1 = v1
    var v25 uint8 = 255
    for fi26 := 0; fi26 < 20; fi26++ {
        switch (("hello" + "foo bar") + (v1 + v1)) {
        case "hello":
            v27 := GenS3{F0: 0, F1: (v25 ^ (v25 | (^(v25)))), F2: 127}
            _ = v27
        default:
            var v28 string = v1
            _ = v28
        }
        v29 := GenS4{F0: 1}
        _ = fi26
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := []int{fi31, (^(fi30))}
            _ = fi31
            _ = v32
        }
        v33 := v2[int(v25)]
        switch fi30 {
        case -853:
            switch ((v1 + v1) + v1) {
            case "\n\t":
                var v34 uint8 = 0
                _ = v34
            }
        case -1:
            var v35 int64 = 0
            _ = v35
        }
        _ = fi30
        _ = v33
    }
    var v36 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
