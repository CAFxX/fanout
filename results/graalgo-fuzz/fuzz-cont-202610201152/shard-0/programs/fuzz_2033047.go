package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint64 }

func main() {
    var v1 uint64 = 604
    v2 := make(map[int]uint32)
    var v3 uint = 355
    switch 4294967295 {
    case 128:
        v4 := GenS1{F0: -1, F1: 108, F2: 957}
        _ = v4
    default:
        fmt.Println(v3, v1)
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        switch (fi5 | fi5) {
        case 127:
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS2{F0: 42, F1: 308}
                v1 = v1
                v1 = 785
                _ = fi6
                _ = v7
            }
        case -899:
            v2[(fi5 & (-(fi5)))] = 503
        default:
            v8 := GenS3{F0: v1, F1: 100, F2: (v1 >> ((v3) & 63))}
            _ = v8
        }
        if (4294967295 == 2147483648) {
            v1 += v1
            var v9 uint = v3
            _ = v9
        }
        _ = fi5
    }
    var v10 int64 = -547
    fmt.Println(v3, v10)
    v3 = v3
    v10 -= -596
    var v11 uint64 = v1
    var v12 uint8 = 128
    var v13 uint64 = 2
    fmt.Println(v13, v13, v13)
    if (("" + "") > "") {
        if (-128 > -80) {
            var v14 int64 = v10
            v15 := GenS3{F0: (v1 + v11), F1: -114, F2: v1}
            _ = v14
            _ = v15
        } else {
            var v16 uint16 = 42
            v17 := GenS1{F0: int32(v12), F1: -108, F2: 77}
            _ = v16
            _ = v17
        }
        var v18 uint16 = 255
        fmt.Println(v11, v13, v10)
        _ = v18
    } else {
        fmt.Println(v11)
        for fi19 := 0; fi19 < 9; fi19++ {
            switch ("foo bar" + (("hello" + "a") + "hello")) {
            case "foo bar":
                v20 := []int{fi19, fi19, -1}
                _ = v20
            case "a":
                v1 = (v13 & v11)
            case "":
                switch 542 {
                case 127:
                    for fi21 := 0; fi21 < 4; fi21++ {
                        fmt.Println(fi21, v11, v13)
                        v3 += v3
                        _ = fi21
                    }
                case -968:
                    var v22 float64 = (float64(float32(0.1)) * 0.0)
                    _ = v22
                case 867:
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 float64 = float64(0.1)
                        if true {
                            var v25 float32 = 0.1
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                v12 = v12
                            default:
                                v13 = (v13 << (((^(v3))) & 63))
                            }
                            if ((32767 >= 32767) && (v12 != v12)) {
                                var v26 float64 = (float64(float32(((((float64((v25 + float32((v24 - v24)))) + ((v24 * v24) - v24)) + (v24 - v24)) / v24) + v24))) / v24)
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v10, v1)
                                }
                                _ = v26
                            } else {
                                v1 = v11
                                switch "" {
                                case "\n\t":
                                    var v27 string = ("a" + "hello")
                                    _ = v27
                                default:
                                    v24 = v24
                                }
                                if (!(!(int64(fi19) <= (v10 + v10)))) {
                                    v28 := GenS1{F0: -791, F1: -128, F2: (fi23 + (-(fi19)))}
                                    _ = v28
                                }
                            }
                            _ = v25
                        } else {
                            v10 = (v10 ^ 9223372036854775807)
                            var v29 uint64 = 4294967295
                            var v30 uint = (v3 & 859)
                            _ = v29
                            _ = v30
                        }
                        var v31 int = fi23
                        var v32 uint16 = 255
                        v24 += 3.14159
                        v24 += v24
                        _ = fi23
                        _ = v24
                        _ = v31
                        _ = v32
                    }
                }
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
