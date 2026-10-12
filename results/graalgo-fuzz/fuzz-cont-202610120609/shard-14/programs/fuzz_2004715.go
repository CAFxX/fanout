package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    var v1 int16 = 0
    var v2 string = "hello"
    var v3 float64 = (((((1e10 + 1e10) * 1e10) / (((1e10 / 1.5) * 3.14159) / float64((1e10 * 0.1)))) / 0.1) - 1.5)
    switch ((("a" + (("a" + "hello") + v2)) + v2) + v2) {
    case "hello":
        v1 = (v1 + (v1 >> ((65535) & 63)))
    case "a":
        v4 := GenS1{F0: 344}
        _ = v4
    default:
        var v5 int16 = (-(v1))
        _ = v5
    }
    var v6 uint16 = 0
    v2 = "\n\t"
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 18
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS2{F0: 100}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            switch "a" {
            case "":
                if (!((!true) || true)) {
                    v6 += 255
                    var v12 int8 = 1
                    if true {
                        fmt.Println(v3, v8)
                        switch v2 {
                        case "":
                            v13 := []int{1, fi7, fi11, 2147483647}
                            _ = v13
                        case "\n\t":
                            v12 = (v12 + v12)
                        case "a":
                            var v14 int32 = 1
                            _ = v14
                        }
                    }
                    v15 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: (v8 ^ (v8 ^ (v8 & v8)))}
                    if (true || true) {
                        v16 := []int{}
                        v12 = 80
                        _ = v16
                    } else {
                        var v17 uint = 65535
                        var v18 int16 = (v1 / (((v1 % ((42) | 1))) | 1))
                        _ = v17
                        _ = v18
                    }
                    _ = v12
                    _ = v15
                }
            default:
                fmt.Println(v3)
            }
            _ = fi11
        }
        _ = fi7
        _ = v8
    }
    var v19 float32 = float32(3.14159)
    v20 := GenS2{F0: 0}
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 int = fi21
        if (!(v19 == v19)) {
            if (!false) {
                switch v22 {
                case -9223372036854775808:
                    fmt.Println(v1, v22)
                case -738:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if false {
                            if (!true) {
                                v24 := []int{fi23, 1, -815, 734}
                                _ = len(v24)
                                v25 := GenS2{F0: 0}
                                fmt.Println(v1, v1)
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                if false {
                                    var v26 uint8 = 127
                                    v26 -= (v26 & v26)
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    _ = v26
                                } else {
                                    v27 := GenS3{F0: 42, F1: 4294967295, F2: 1}
                                    switch "" {
                                    case "a":
                                        v28 := GenS3{F0: 133, F1: 42, F2: 737}
                                        _ = v28
                                    case "hello":
                                        var v29 string = v2
                                        _ = v29
                                    }
                                    v2 = v2
                                    _ = v27
                                }
                                var v30 uint64 = 472
                                _ = v24
                                _ = v25
                                _ = v30
                            } else {
                                switch (v22 ^ (v22 - fi23)) {
                                case 0:
                                    v1 -= v1
                                case 32767:
                                    v3 += float64(v19)
                                default:
                                    var v31 string = (((v2 + v2) + v2) + "hello")
                                    _ = v31
                                }
                                fmt.Println(v22)
                                if (fi21 < v22) {
                                    fmt.Println(v3, v19, v22)
                                    v32 := GenS4{F0: v6, F1: 255}
                                    _ = v32
                                }
                            }
                        } else {
                            v19 = ((v19 / (v19 - (3.14159 / v19))) * (-2.25 - ((0.0 / v19) / v19)))
                            var v33 uint8 = 0
                            v33 = (v33 * v33)
                            _ = v33
                        }
                        v34 := GenS4{F0: v6, F1: 18446744073709551615}
                        fmt.Println(v6)
                        _ = fi23
                        _ = v34
                    }
                case -183:
                    v35 := GenS1{F0: 1}
                    _ = v35
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    v22 = 1
                    fmt.Println(v1, fi36, v19)
                    _ = fi36
                }
                var v37 int64 = -9223372036854775808
                _ = v37
            }
            v3 = ((v3 / v3) / (v3 / v3))
            var v38 int64 = 1
            _ = v38
        } else {
            v6 = (^(v6))
            v39 := []int{}
            _ = v39
        }
        v3 = -2.25
        _ = fi21
        _ = v22
    }
    fmt.Println(v1)
    v2 = (v2 + "a")
    var v40 float32 = (((((v19 - v19) - v19) / v19) + v19) + float32((v3 - 0.1)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v40
    _ = v6
}
