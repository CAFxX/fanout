package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (-2.25 / 1.0)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 int64 = 42
    if true {
        if true {
            switch 127 {
            case 714:
                if (898 < 4294967295) {
                    v2 = v2
                    for fi6 := 0; fi6 < 14; fi6++ {
                        var v7 uint32 = 0
                        if (-285 <= -680) {
                            for fi8 := 0; fi8 < 20; fi8++ {
                                if false {
                                    fmt.Println(v5, v2)
                                    v9 := GenS1{F0: 0, F1: 9223372036854775807}
                                    fmt.Println(v5, fi8)
                                    switch -1 {
                                    case 1:
                                        fmt.Println(fi6, v5)
                                    case -9223372036854775808:
                                        fmt.Println(fi6, v7, v5)
                                    case 128:
                                        v10 := v3[(^(fi8))]
                                        _ = v10
                                    }
                                    _ = v9
                                }
                                switch (fi6 / ((fi6) | 1)) {
                                case 0:
                                    v11 := GenS2{F0: 255, F1: 42}
                                    _ = v11
                                case 4294967295:
                                    v12 := GenS1{F0: 42, F1: (fi6 >> ((1) & 63))}
                                    _ = v12
                                default:
                                    if false {
                                        if ("\n\t" < "") {
                                            v13 := GenS1{F0: 716, F1: fi6}
                                            v7 += v7
                                            _ = v13
                                        } else {
                                            var v14 bool = false
                                            var v15 uint32 = v7
                                            _ = v14
                                            _ = v15
                                        }
                                    }
                                }
                                if true {
                                    v16 := GenS2{F0: 255, F1: 100}
                                    _ = v16
                                }
                                var v17 uint8 = 128
                                var v18 int64 = v5
                                _ = fi8
                                _ = v17
                                _ = v18
                            }
                            v5 = (v5 / ((int64(v7)) | 1))
                        }
                        v19 := GenS3{F0: 0}
                        _ = fi6
                        _ = v19
                        _ = v7
                    }
                } else {
                    v2 += 1e10
                    switch -9223372036854775808 {
                    case -736:
                        v5 += v5
                    case -740:
                        v20 := []int{-24, 4294967295, 1}
                        _ = v20
                    case 0:
                        switch ((("foo bar" + "") + "hello") + ((("a" + ("" + "")) + "") + "")) {
                        case "foo bar":
                            var v21 int16 = 460
                            _ = v21
                        case "":
                            v22 := GenS1{F0: 65535, F1: -9223372036854775808}
                            _ = v22
                        default:
                            switch 289 {
                            case 322:
                                var v23 int8 = 1
                                _ = v23
                            }
                        }
                    default:
                        fmt.Println(v2)
                    }
                    fmt.Println(v2, v5)
                }
            }
            v5 = 0
            var v24 int = 2147483647
            _ = v24
        }
        fmt.Println(v2, v2)
    }
    fmt.Println(v2, v2)
    v2 = v2
    var v25 int8 = -1
    fmt.Println(v2, v25, v5)
    var v26 int8 = v25
    v25 -= v26
    v5 -= v5
    v2 = (((0.0 + -2.25) / -2.25) / v2)
    var v27 uint8 = 124
    var v28 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
