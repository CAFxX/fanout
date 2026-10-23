package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        if ((((1e10 + (0.1 + 3.14159)) * ((-2.25 - 1.5) - 3.14159)) > 1.5) || true) {
            v7 := GenS1{F0: 4294967295, F1: fi6, F2: 42}
            v8 := []int{568, 32767}
            _ = v7
            _ = v8
        }
        var v9 int8 = 42
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 string = "hello"
            var v12 int8 = v9
            if false {
                v12 = v12
                if (42 <= -32768) {
                    v12 = (^(v12))
                }
                v13 := []int{-351, fi10, fi6}
                fmt.Println(v9)
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
        _ = v9
    }
    var v14 float64 = (0.1 + (0.1 * 3.14159))
    var v15 uint8 = 225
    _ = v14
    _ = v15
    return (v15 * 42)
}

func main() {
    var v1 float32 = 1.5
    var v2 int64 = -43
    var v3 uint32 = 2147483647
    var v4 uint16 = 616
    v5 := make(map[int]string)
    v16 := []int{}
    var v17 uint64 = 42
    if (1 > 1) {
        v1 = 1e10
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            if false {
                _ = len(v5)
            }
            _ = fi18
        }
        v2 = v2
    }
    var v19 bool = true
    v3 -= (v3 / (((v3 + v3)) | 1))
    v20 := v5[599]
    v21 := GenS1{F0: 42, F1: 32768, F2: 0}
    switch -9223372036854775808 {
    case -579:
        switch v20 {
        case "a":
            v20 = (v20 + (v20 + "foo bar"))
        }
    }
    if (-128 <= 22) {
        var v22 uint16 = v4
        var v23 uint8 = 61
        _ = v22
        _ = v23
    }
    var v24 int64 = -28
    _ = len(v16)
    var v25 int = 9223372036854775807
    if v19 {
        v26 := GenS1{F0: v3, F1: v25, F2: -3}
        _ = v26
    }
    v24 += v24
    if (("\n\t" != v20) && (1.5 <= (1.5 - 3.14159))) {
        var v27 int32 = 42
        switch (v25 ^ (v25 >> ((1) & 63))) {
        case -9223372036854775808:
            fmt.Println(v4)
        default:
            if (v19 || false) {
                fmt.Println(v17, v27)
                fmt.Println(v3)
                if v19 {
                    var v28 uint64 = 1
                    v17 -= v28
                    v28 = (v17 & v28)
                    _ = v28
                }
            }
        }
        if v19 {
            if (((369 == v27) || (-2.25 != 1e10)) && (((v1 * (0.1 / 0.1)) == ((float32(1e10) * ((v1 - (v1 + -2.25)) + v1)) + v1)) && (v19 && (2 == 42)))) {
                switch 127 {
                case 4294967295:
                    if v19 {
                        v29 := GenS2{F0: -212, F1: -81}
                        for fi30 := 0; fi30 < 18; fi30++ {
                            fmt.Println(v19)
                            switch 42 {
                            case -165:
                                var v31 int64 = -1
                                _ = v31
                            case 9223372036854775807:
                                v27 = 0
                            case 128:
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    fmt.Println(fi30, fi30, v3)
                                    _ = fi32
                                }
                            default:
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    var v34 bool = v19
                                    v35 := GenS2{F0: int16(v27), F1: -1}
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            _ = fi30
                        }
                        _ = v29
                    }
                }
                v3 = v3
            } else {
                fmt.Println(v4)
                v1 -= (v1 / float32(0.1))
                v36 := GenS3{F0: true, F1: -128}
                _ = v36
            }
        }
        _ = v27
    }
    fmt.Println(v4, v2, v1)
    fmt.Println(v20, v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
