package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int8 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    var v3 uint64 = 2
    v4 := make(map[string]bool)
    switch 32768 {
    case 32767:
        v1 = v1
    default:
        var v5 int64 = 1
        _ = v5
    }
    v6 := []int{604, 32768}
    if (1000 < -32768) {
        for fi7 := 0; fi7 < 6; fi7++ {
            switch (fi7 >> ((0) & 63)) {
            case 32768:
                v8 := GenS1{F0: 5, F1: 1, F2: 0}
                _ = v8
            }
            _ = fi7
        }
        v3 = 75
    }
    var v9 uint16 = 42
    var v10 uint64 = (v3 ^ (v3 / ((v3) | 1)))
    fmt.Println(v10)
    var v11 int16 = 997
    var v12 uint8 = 128
    var v13 float64 = ((1e10 / float64(1e10)) - 3.14159)
    var v14 float64 = 1.5
    v3 += (v3 << ((732) & 63))
    for fi15 := 0; fi15 < 1; fi15++ {
        if true {
            if false {
                switch v1 {
                case "":
                    v16 := GenS2{F0: (v1 + "\n\t"), F1: 2, F2: (42 >= (v11 / ((v11) | 1)))}
                    _ = v16
                case "a":
                    var v17 bool = (542 >= 0)
                    _ = v17
                default:
                    v1 = v1
                }
                v12 -= (^(v12))
            } else {
                if (float64(1.5) != ((v14 * v13) + (v13 / v14))) {
                    v9 -= (v9 % ((v9) | 1))
                    if false {
                        if true {
                            fmt.Println(v14, v12)
                            switch (fi15 & 32768) {
                            case 128:
                                v11 += (-(v11))
                            case 32767:
                                var v18 int = (fi15 * 4294967295)
                                _ = v18
                            }
                        }
                        var v19 int32 = 0
                        _ = v19
                    }
                } else {
                    switch v1 {
                    case "hello":
                        switch (("hello" + "") + "\n\t") {
                        case "a":
                            fmt.Println(v14, v3, v14)
                        case "foo bar":
                            v20 := GenS2{F0: (v1 + "hello"), F1: (v12 + v12), F2: (1.5 <= float32(float64(0.0)))}
                            _ = v20
                        default:
                            fmt.Println(v3, v1)
                        }
                    case "":
                        for fi21 := 0; fi21 < 20; fi21++ {
                            var v22 uint64 = v3
                            _ = fi21
                            _ = v22
                        }
                    case "foo bar":
                        var v23 float64 = v13
                        _ = v23
                    default:
                        v14 = (((-2.25 / (-2.25 - 3.14159)) - (((v14 - (v14 + v14)) / (v13 / v14)) * (((v13 - (float64(float32((((v13 - 1e10) + v14) - float64((3.14159 * 0.1))))) / ((v14 * (v13 + 0.1)) + v13))) - -2.25) / 1.0))) * v13)
                    }
                }
                var v24 uint16 = (v9 ^ 2)
                _ = v24
            }
            fmt.Println(v9, v3, fi15)
            var v25 int16 = v11
            _ = v25
        } else {
            var v26 int8 = 100
            v3 -= v10
            _ = v26
        }
        var v27 int32 = 2
        switch fi15 {
        case 202:
            v9 = v9
        case 42:
            v28 := GenS3{F0: 42}
            _ = v28
        default:
            v9 = v9
        }
        _ = fi15
        _ = v27
    }
    var v29 string = v1
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
