package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int16; F2 int }
type GenS7 struct { F0 int8; F1 uint8; F2 uint }
func genFn0() int64 {
    var v4 string = "\n\t"
    fmt.Println(v4)
    fmt.Println(v4)
    fmt.Println(v4)
    v5 := GenS1{F0: 1}
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 int) uint32 {
    var v6 bool = (false || ((3.14159 / 3.14159) == -2.25))
    p1 -= (^(p1))
    fmt.Println(p0, v6, v6)
    _ = p0
    _ = p1
    _ = v6
    return 65535
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 331
    var v3 uint16 = 550
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            v8 := GenS1{F0: 1000}
            v9 := GenS2{F0: 528}
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    if true {
        var v10 int64 = 1
        v11 := []int{-372}
        if ((v1 / (((v1 >> ((255) & 63))) | 1)) <= (^(v1))) {
            switch "foo bar" {
            case "":
                var v12 uint = 4294967295
                _ = v12
            default:
                fmt.Println(v2, v2, v1)
            }
            var v13 bool = (0 < 486)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 1; fi14++ {
                if true {
                    if true {
                        switch fi14 {
                        case 128:
                            var v15 int16 = int16(v1)
                            _ = v15
                        case 42:
                            for fi16 := 0; fi16 < 11; fi16++ {
                                if true {
                                    v17 := GenS3{F0: 81}
                                    _ = v17
                                }
                                _ = fi16
                            }
                        case -478:
                            v18 := GenS4{F0: (v10 >> ((68) & 63))}
                            _ = v18
                        default:
                            var v19 float64 = 0.1
                            _ = v19
                        }
                        var v20 int = -314
                        switch ("foo bar" + ("\n\t" + "\n\t")) {
                        case "\n\t":
                            for fi21 := 0; fi21 < 4; fi21++ {
                                var v22 uint8 = 255
                                _ = fi21
                                _ = v22
                            }
                        case "foo bar":
                            for fi23 := 0; fi23 < 14; fi23++ {
                                if (v1 > (v1 << ((298) & 63))) {
                                    var v24 uint64 = 146
                                    var v25 uint64 = v24
                                    _ = v24
                                    _ = v25
                                }
                                switch 42 {
                                case 127:
                                    fmt.Println(v3, v2)
                                case -1:
                                    v1 += v1
                                }
                                _ = fi23
                            }
                        case "":
                            var v26 uint64 = 1
                            _ = v26
                        default:
                            v20 -= (fi14 + fi14)
                        }
                        _ = v20
                    }
                    var v27 int32 = (v2 * v2)
                    var v28 uint16 = (^(v3))
                    _ = v27
                    _ = v28
                }
                _ = fi14
            }
            v29 := GenS1{F0: -1}
            if false {
                v30 := GenS5{F0: 0, F1: 18446744073709551615}
                if (!false) {
                    v11 = append(v11, 0)
                    var v31 int16 = int16(v2)
                    _ = v31
                }
                switch 32768 {
                case 4294967295:
                    var v32 string = ("foo bar" + "\n\t")
                    _ = v32
                case 42:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 uint = 1
                        var v35 int16 = 1
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                case 127:
                    v36 := GenS6{F0: v3, F1: -199, F2: 2147483647}
                    _ = v36
                }
                _ = v30
            }
            _ = v29
        }
        _ = v10
        _ = v11
    }
    if false {
        v3 = 959
    } else {
        v37 := []int{1, 212}
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        if true {
            v39 := []int{fi38, (fi38 >> ((4294967295) & 63))}
            var v40 string = "a"
            fmt.Println(v3)
            _ = v39
            _ = v40
        } else {
            v41 := []int{fi38, (fi38 ^ 9223372036854775807), 2147483647}
            v1 -= 2
            _ = v41
        }
        v42 := GenS7{F0: -7, F1: 128, F2: 546}
        fmt.Println(v3, v3)
        _ = fi38
        _ = v42
    }
    v1 = genFn1(128, 4294967295)
    v43 := []int{-2147483648, 2147483647, 32767, 173}
    v1 += genFn1(128, -2147483648)
    v2 = v2
    var v44 uint64 = 2
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 float32 = (3.14159 * 0.0)
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
}
