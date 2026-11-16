package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := []int{128, -1, 9223372036854775807}
    var v2 uint32 = 2147483647
    v3 := make(map[int]string)
    var v4 bool = false
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 17; fi5++ {
        v2 -= (v2 % (((v2 & v2)) | 1))
        _ = fi5
    }
    v6 := GenS1{F0: -1, F1: 34, F2: -32768}
    v7 := GenS2{F0: 42, F1: 0, F2: 127}
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: -905, F1: 139, F2: 1}
            v2 = 1
            _ = fi9
            _ = v10
        }
        if ((((-110 > -56) || (((((((-2.25 / 3.14159) - 0.1) - float32((3.14159 / 3.14159))) + 1e10) == float32(1e10)) || true) || (v4 || v4))) || v4) && (v2 == 65535)) {
            v11 := GenS3{F0: 42}
            fmt.Println(v2, fi8, v4)
            var v12 int8 = 0
            _ = v11
            _ = v12
        }
        if v4 {
            v2 -= v2
        }
        _ = fi8
    }
    switch 32768 {
    case -1:
        if false {
            fmt.Println(v4, v2)
            if (!true) {
                v4 = false
            } else {
                if false {
                    v13 := GenS3{F0: 982}
                    for fi14 := 0; fi14 < 20; fi14++ {
                        if ((v2 == 1) && false) {
                            v4 = (1.5 > (3.14159 - 1.5))
                        } else {
                            v2 = 2
                            v15 := GenS3{F0: 2}
                            _ = v15
                        }
                        switch (fi14 + (fi14 ^ 4294967295)) {
                        case 1:
                            var v16 uint = 0
                            _ = v16
                        case 0:
                            v17 := GenS3{F0: 42}
                            _ = v17
                        case 2147483647:
                            if len(v1) > 0 {
                                v1[0] = (fi14 - 9223372036854775807)
                            }
                        default:
                            v4 = true
                        }
                        _ = fi14
                    }
                    _ = v13
                }
            }
            var v18 bool = (255 < 232)
            _ = v18
        } else {
            var v19 int8 = 42
            _ = v19
        }
    case 9223372036854775807:
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int64 = -967
            v2 = (^(v2))
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int32 = 278
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(fi22, v23, v23)
            v25 := GenS2{F0: v23, F1: 1, F2: 255}
            if false {
                v23 = v23
                if v4 {
                    v2 -= (v2 % ((1) | 1))
                    if true {
                        var v26 uint16 = 255
                        fmt.Println(v2, fi24)
                        var v27 int8 = 1
                        _ = v26
                        _ = v27
                    }
                } else {
                    fmt.Println(fi22)
                }
                switch fi22 {
                case 201:
                    v4 = v4
                case 127:
                    v28 := GenS1{F0: 1, F1: v2, F2: 994}
                    _ = v28
                case 4294967295:
                    var v29 int64 = 942
                    _ = v29
                }
            }
            _ = fi24
            _ = v25
        }
        v2 -= (v2 * (v2 * v2))
        _ = fi22
        _ = v23
    }
    if (true || (65535 != 1)) {
        v2 -= v2
        var v30 uint = 744
        _ = v30
    } else {
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 string = ""
            var v33 int16 = 1000
            v1 = append(v1, (fi31 & (fi31 % ((fi31) | 1))))
            _ = fi31
            _ = v32
            _ = v33
        }
        v34 := v3[42]
        for fi35 := 0; fi35 < 18; fi35++ {
            if v4 {
                for fi36 := 0; fi36 < 1; fi36++ {
                    v34 = v34
                    delete(v3, -9223372036854775808)
                    _ = fi36
                }
            } else {
                switch v34 {
                case "":
                    var v37 int8 = 100
                    _ = v37
                case "foo bar":
                    _ = len(v1)
                case "a":
                    v34 = (v34 + v34)
                default:
                    for fi38 := 0; fi38 < 5; fi38++ {
                        v34 = (v34 + (v34 + v34))
                        v34 = (v34 + v34)
                        _ = len(v3)
                        _ = fi38
                    }
                }
            }
            var v39 string = v34
            v40 := GenS1{F0: 1, F1: 2147483648, F2: -1}
            _ = fi35
            _ = v39
            _ = v40
        }
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
