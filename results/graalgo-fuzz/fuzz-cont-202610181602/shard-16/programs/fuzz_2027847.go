package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 string, p1 int16, p2 string) int8 {
    v6 := GenS1{F0: -767, F1: 849}
    v7 := GenS1{F0: 538, F1: 1}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float64 = (-2.25 + -2.25)
        fmt.Println(p2)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return -31
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = 32767
    var v3 float32 = 0.0
    v4 := make(map[string]bool)
    var v5 float64 = float64(v3)
    v11 := []int{-9223372036854775808, 32767, -847, 2147483647}
    if (uint(v1) >= 255) {
        v2 = v2
        fmt.Println(v1, v3, v1)
        for fi12 := 0; fi12 < 18; fi12++ {
            if true {
                v2 += v2
            } else {
                switch int(v1) {
                case -164:
                    if ((-9223372036854775808 >= 0) && (703 >= 2147483647)) {
                        var v13 string = ""
                        switch (((((v13 + ((v13 + v13) + v13)) + v13) + "\n\t") + (v13 + (v13 + ""))) + (v13 + v13)) {
                        case "foo bar":
                            var v14 int8 = 1
                            _ = v14
                        }
                        fmt.Println(fi12)
                        _ = v13
                    }
                case -9223372036854775808:
                    fmt.Println(v2)
                }
                if len(v11) > 0 {
                    v11[0] = 1
                }
            }
            fmt.Println(v2)
            _ = fi12
        }
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    switch 786 {
    case 32768:
        switch "" {
        case "a":
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint16 = 255
                for fi17 := 0; fi17 < 7; fi17++ {
                    if false {
                        v16 += (v16 & (^(v16)))
                    } else {
                        v5 -= v5
                        v2 += (v2 % (((v2 * 42)) | 1))
                    }
                    if true {
                        if false {
                            v1 = v1
                            var v18 uint16 = v16
                            v19 := GenS1{F0: int64(v1), F1: v1}
                            _ = v18
                            _ = v19
                        } else {
                            if (!false) {
                                var v20 uint64 = 42
                                v2 = 1
                                var v21 uint8 = 2
                                _ = v20
                                _ = v21
                            } else {
                                if (int16(fi17) == (-(v2))) {
                                    fmt.Println(v5)
                                    var v22 uint = 1
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    _ = v22
                                } else {
                                    v5 -= (v5 - 0.0)
                                    fmt.Println(fi17, fi17)
                                }
                            }
                        }
                    } else {
                        switch 32768 {
                        case 343:
                            v5 = ((v5 * 1e10) - float64(((v3 - (0.0 + v3)) * v3)))
                        case -2147483648:
                            var v23 string = "a"
                            _ = v23
                        case 4294967295:
                            v24 := v4["a"]
                            _ = v24
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v25 float64 = (v5 - 3.14159)
                        _ = v25
                    }
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
        default:
            fmt.Println(v3, v1, v5)
        }
    case 112:
        fmt.Println(v1)
    case 9223372036854775807:
        for fi26 := 0; fi26 < 11; fi26++ {
            v5 = v5
            _ = fi26
        }
    }
    var v27 int32 = -2147483648
    fmt.Println(v2)
    var v28 uint8 = 1
    if true {
        v29 := GenS2{F0: 4294967295, F1: 2}
        _ = v29
    }
    v30 := GenS3{F0: 127, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
