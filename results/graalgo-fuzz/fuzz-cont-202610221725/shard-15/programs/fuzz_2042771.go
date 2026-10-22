package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    v1 := []uint32{89, 700, 2147483647, 1}
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    var v5 int16 = int16(v4)
    v4 = v4
    if true {
        var v6 string = "hello"
        v5 = -32768
        _ = v6
    } else {
        if false {
            if false {
                v7 := v2[276]
                _ = v7
            } else {
                var v8 uint64 = 10
                _ = v8
            }
            delete(v2, 110)
        } else {
            v9 := GenS1{F0: (-(v5)), F1: 117, F2: 65535}
            fmt.Println(v5)
            if true {
                var v10 int32 = -608
                _ = v10
            } else {
                if ((v4 & v4) == v4) {
                    fmt.Println(v5, v5)
                } else {
                    fmt.Println(v5, v5)
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v4 = v4
                        if (true && (42 >= (v4 - v4))) {
                            v5 -= v5
                        } else {
                            v4 = 0
                        }
                        v5 = v5
                        _ = fi11
                    }
                }
            }
            _ = v9
        }
        var v12 uint32 = 636
        var v13 int64 = -1
        _ = v12
        _ = v13
    }
    v1 = append(v1, 65535)
    if (42 >= 72) {
        var v14 bool = false
        var v15 int8 = 42
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint16 = v4
            var v18 int8 = (v15 % ((v15) | 1))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v14
        _ = v15
    } else {
        var v19 bool = true
        _ = v19
    }
    fmt.Println(v4, v4)
    v20 := v2[9223372036854775807]
    fmt.Println(v20)
    v21 := GenS2{F0: 100, F1: 1}
    var v22 uint = 4294967295
    for fi23 := 0; fi23 < 5; fi23++ {
        if false {
            v24 := GenS2{F0: 42, F1: (fi23 / ((32767) | 1))}
            _ = v24
        } else {
            switch v20 {
            case "\n\t":
                var v25 float32 = (((1.5 + (0.0 - (-2.25 / 3.14159))) + (-2.25 + 1e10)) * ((((((0.0 / 1.5) - (float32(float64((1.5 + 0.1))) - (0.1 / 3.14159))) * 0.0) + (-2.25 * (float32((0.0 * 3.14159)) - (((-2.25 - 0.0) / (0.1 + float32(1.5))) - 1e10)))) / 1.5) / (1.5 / 1.5)))
                _ = v25
            case "hello":
                for fi26 := 0; fi26 < 14; fi26++ {
                    v27 := GenS1{F0: v5, F1: 0, F2: 797}
                    if (-658 > 0) {
                        var v28 int = (^(fi23))
                        v20 = v20
                        _ = len(v1)
                        _ = v28
                    } else {
                        fmt.Println(v20, v4, v4)
                        switch "foo bar" {
                        case "foo bar":
                            if (2147483647 > 2147483648) {
                                v29 := GenS1{F0: 1000, F1: -128, F2: 0}
                                v5 = (^(v5))
                                fmt.Println(fi26, v5, fi26)
                                _ = v29
                            }
                        case "hello":
                            var v30 uint32 = 2
                            _ = v30
                        }
                    }
                    v2[fi23] = v20
                    if (v20 == ((((v20 + v20) + v20) + (("\n\t" + "hello") + (("hello" + v20) + "\n\t"))) + "")) {
                        if true {
                            var v31 int16 = v5
                            v32 := GenS3{F0: v4}
                            switch (fi26 - fi26) {
                            case 688:
                                if false {
                                    v20 = ((v20 + v20) + (("a" + "foo bar") + "hello"))
                                }
                            case 32768:
                                if (441 > 976) {
                                    v22 += (v22 / (((v22 % ((42) | 1))) | 1))
                                }
                            default:
                                v4 += 32767
                                var v33 uint8 = 42
                                v34 := GenS1{F0: v31, F1: 15, F2: 918}
                                v20 = "\n\t"
                                var v35 uint16 = 65535
                                if (1 <= 1) {
                                    var v36 bool = (!false)
                                    if v36 {
                                        fmt.Println(v22)
                                    }
                                    var v37 float64 = ((0.1 + -2.25) + ((0.1 * float64(0.0)) / (0.1 / (1.5 - 0.1))))
                                    _ = v36
                                    _ = v37
                                } else {
                                    v5 -= 1000
                                    v4 -= v35
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                }
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(v22, fi23)
                        }
                        v4 += v4
                    } else {
                        v38 := GenS4{F0: 4294967295, F1: (v4 / ((42) | 1)), F2: 703}
                        var v39 bool = (1 > -40)
                        var v40 bool = (0.1 < (1e10 / 1.5))
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    _ = fi26
                    _ = v27
                }
            case "foo bar":
                v22 += (v22 * v22)
            default:
                fmt.Println(v20, v4)
            }
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
