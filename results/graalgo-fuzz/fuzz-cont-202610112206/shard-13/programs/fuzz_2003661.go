package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    var v3 string = "hello"
    var v4 int8 = -128
    v5 := GenS1{F0: 201}
    for fi6 := 0; fi6 < 20; fi6++ {
        if (true || true) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: (-(fi7))}
                _ = fi7
                _ = v8
            }
            v2 = v2
            var v9 uint32 = 920
            _ = v9
        } else {
            var v10 int8 = (v4 / (((v4 - (v4 + v4))) | 1))
            _ = v10
        }
        _ = fi6
    }
    v11 := GenS1{F0: -2147483648}
    if (false && (false && (v2 <= v2))) {
        var v12 uint64 = 0
        var v13 uint32 = 293
        _ = v12
        _ = v13
    } else {
        if (false && false) {
            v4 -= 127
            var v14 float32 = ((v2 * float32(float64(v2))) / float32(-2.25))
            switch 1 {
            case 127:
                if true {
                    if false {
                        var v15 float32 = v14
                        v2 = 1e10
                        var v16 uint64 = 596
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v3, v2)
                } else {
                    v14 += v14
                    switch 114 {
                    case 9223372036854775807:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 uint8 = 0
                            if (42 > 1) {
                                fmt.Println(fi17, v2, v4)
                            }
                            for fi19 := 0; fi19 < 1; fi19++ {
                                v20 := GenS2{F0: 1, F1: 1}
                                switch (fi19 / ((fi17) | 1)) {
                                case 32767:
                                    v21 := GenS2{F0: v18, F1: -9223372036854775808}
                                    _ = v21
                                default:
                                    fmt.Println(v2)
                                }
                                if (162 != 18446744073709551615) {
                                    var v22 float64 = 1e10
                                    v23 := GenS3{F0: v3, F1: (v18 ^ (v18 ^ v18)), F2: 648}
                                    _ = v22
                                    _ = v23
                                } else {
                                    if true {
                                        if ((23 >= (v18 | 162)) || true) {
                                            v14 -= float32((float64(1e10) * (float64((v2 * 0.1)) / (float64(1e10) - 3.14159))))
                                        }
                                        fmt.Println(v14)
                                    }
                                    var v24 float64 = 0.0
                                    if (((-2.25 + float64(v2)) + (((v24 / 1.0) / v24) * float64(1e10))) > v24) {
                                        if false {
                                            var v25 float32 = -2.25
                                            v26 := []int{}
                                            v1[v3] = 42
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    _ = v24
                                }
                                var v27 int8 = (v4 & v4)
                                _ = fi19
                                _ = v20
                                _ = v27
                            }
                            _ = fi17
                            _ = v18
                        }
                    }
                }
            }
            _ = v14
        } else {
            var v28 uint = 255
            _ = v28
        }
    }
    fmt.Println(v4, v3)
    var v29 int = 1
    if true {
        v30 := GenS3{F0: v3, F1: 2, F2: 270}
        var v31 uint = 0
        delete(v1, (v3 + ("" + v3)))
        _ = v30
        _ = v31
    } else {
        switch v3 {
        case "\n\t":
            v32 := GenS4{F0: v4}
            _ = v32
        case "":
            if true {
                v33 := GenS5{F0: 32767}
                if false {
                    if true {
                        for fi34 := 0; fi34 < 11; fi34++ {
                            v4 = (v4 << ((255) & 63))
                            _ = fi34
                        }
                    } else {
                        fmt.Println(v3, v4, v4)
                        v2 = v2
                        fmt.Println(v2, v29, v2)
                    }
                    fmt.Println(v3)
                    v4 = (v4 >> ((446) & 63))
                }
                _ = v33
            } else {
                var v35 uint64 = 4294967295
                var v36 float64 = 1e10
                _ = v35
                _ = v36
            }
        case "a":
            fmt.Println(v4, v29)
        }
    }
    var v37 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
