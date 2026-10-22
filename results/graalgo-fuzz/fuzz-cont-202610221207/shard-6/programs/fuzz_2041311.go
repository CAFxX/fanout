package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int }

func main() {
    v1 := []int{-1, 127}
    v2 := []int{9223372036854775807, -1, 128, 32768}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint8 = 127
    v5 = (v5 - v5)
    if false {
        v6 := GenS1{F0: 42, F1: 772}
        if false {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int16 = 32767
                if true {
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                }
                var v9 uint16 = 42
                _ = fi7
                _ = v8
                _ = v9
            }
            fmt.Println(v5)
            v5 -= (^(v5))
        }
        _ = v6
    } else {
        switch "" {
        case "a":
            if false {
                var v10 uint32 = 676
                fmt.Println(v5, v5)
                _ = v10
            } else {
                if false {
                    var v11 int32 = 2147483647
                    for fi12 := 0; fi12 < 13; fi12++ {
                        fmt.Println(v11)
                        var v13 uint = 0
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                } else {
                    v5 += v5
                }
                v14 := GenS2{F0: 0, F1: (^(v5))}
                if false {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if true {
                            if true {
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    var v17 float32 = (-2.25 + (((3.14159 / (-2.25 - (1.5 / 0.1))) / 1.0) + ((3.14159 + float32(3.14159)) + (float32(float64(0.1)) / 1.0))))
                                    _ = fi16
                                    _ = v17
                                }
                                v5 -= v5
                                v5 = (v5 % ((v5) | 1))
                            }
                            v18 := GenS2{F0: 100, F1: v5}
                            v5 = v5
                            _ = v18
                        }
                        _ = fi15
                    }
                }
                _ = v14
            }
        default:
            v19 := GenS1{F0: 527, F1: 906}
            _ = v19
        }
    }
    fmt.Println(v5)
    if true {
        if false {
            delete(v4, "\n\t")
            var v20 int = -2147483648
            for fi21 := 0; fi21 < 6; fi21++ {
                if len(v1) > 0 {
                    v1[0] = (fi21 + (fi21 % ((4294967295) | 1)))
                }
                _ = fi21
            }
            _ = v20
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v5)
            case "":
                _ = len(v4)
            }
        }
        var v22 uint32 = 233
        var v23 bool = false
        _ = v22
        _ = v23
    }
    var v24 uint64 = 4294967295
    fmt.Println(v5)
    var v25 string = "a"
    for fi26 := 0; fi26 < 13; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            if (100 == -1) {
                if (50 != 127) {
                    v28 := GenS3{F0: -2147483648, F1: 128}
                    var v29 int64 = 9223372036854775807
                    _ = v28
                    _ = v29
                } else {
                    v24 = 1
                    var v30 int16 = 1000
                    _ = v30
                }
                var v31 float32 = (1.5 * (1e10 * ((-2.25 + float32(-2.25)) - 1e10)))
                _ = v31
            }
            v5 -= v5
            _ = fi27
        }
        _ = fi26
    }
    var v32 string = "a"
    v33 := GenS4{F0: 127}
    var v34 uint32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
