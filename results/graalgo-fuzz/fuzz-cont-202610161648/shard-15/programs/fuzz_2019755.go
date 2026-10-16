package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }
type GenS4 struct { F0 int16; F1 int; F2 int }

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 0.1
    var v3 uint16 = 255
    v4 := make(map[int]uint32)
    var v5 bool = (167 >= 214)
    var v6 int = 470
    switch v1 {
    case "\n\t":
        for fi7 := 0; fi7 < 5; fi7++ {
            switch 954 {
            case 2147483647:
                v8 := GenS1{F0: ""}
                _ = v8
            case -2147483648:
                _ = len(v4)
            case -965:
                var v9 uint8 = 128
                _ = v9
            default:
                v3 = (v3 - (v3 >> ((65535) & 63)))
            }
            v10 := GenS1{F0: v1}
            v2 = ((-2.25 + v2) + (1.5 + (0.1 - (1e10 - float64((-2.25 + 1e10))))))
            _ = fi7
            _ = v10
        }
    default:
        v2 -= 3.14159
    }
    if (v5 || (13 == 109)) {
        fmt.Println(v2, v3, v3)
        v4[v6] = 1
    }
    v5 = v5
    v2 = -2.25
    v11 := GenS2{F0: 222, F1: (v3 ^ (v3 % (((^(v3))) | 1))), F2: 2}
    fmt.Println(v2)
    if v5 {
        var v12 uint64 = uint64(v3)
        _ = v12
    }
    var v13 int64 = 1
    v2 = v2
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            delete(v4, 42)
            if false {
                var v16 uint8 = 0
                _ = v16
            } else {
                v5 = v5
                fmt.Println(v6, v13)
            }
            var v17 int = (fi15 >> ((220) & 63))
            switch (-(v17)) {
            case 0:
                var v18 uint8 = 128
                _ = v18
            case 9223372036854775807:
                var v19 uint = 0
                _ = v19
            default:
                var v20 int32 = -2147483648
                _ = v20
            }
            _ = fi15
            _ = v17
        }
        v21 := GenS3{F0: 1, F1: 4294967295, F2: 1}
        for fi22 := 0; fi22 < 3; fi22++ {
            if (true || (false && v5)) {
                var v23 uint = 42
                _ = v23
            } else {
                var v24 uint = 0
                v25 := GenS4{F0: -32768, F1: (fi22 ^ int(v13)), F2: fi22}
                if (((fi22 | (^(fi14))) < (fi14 - 2147483647)) && v5) {
                    if (0.0 != (v2 - v2)) {
                        if (((v2 / (((((v2 / 0.1) * ((0.1 / v2) / v2)) - 0.1) * 3.14159) + v2)) / (float64(0.1) * v2)) < v2) {
                            var v26 int8 = 0
                            _ = v26
                        } else {
                            fmt.Println(v13, v13)
                            var v27 uint = (v24 | v24)
                            if v5 {
                                var v28 int16 = 32767
                                _ = v28
                            } else {
                                switch v6 {
                                case 4294967295:
                                    switch v1 {
                                    case "a":
                                        var v29 int8 = 42
                                        _ = v29
                                    case "":
                                        var v30 int8 = -82
                                        _ = v30
                                    case "hello":
                                        v27 = 65535
                                    default:
                                        v27 = 0
                                    }
                                case 128:
                                    if v5 {
                                        fmt.Println(v6, v5, v24)
                                        fmt.Println(fi14, v3, v27)
                                    }
                                case -2147483648:
                                    v31 := GenS2{F0: 42, F1: 1, F2: (v27 ^ 707)}
                                    _ = v31
                                }
                                var v32 uint = v24
                                _ = v32
                            }
                            _ = v27
                        }
                    } else {
                        switch v1 {
                        case "\n\t":
                            v33 := []int{fi14}
                            _ = v33
                        case "":
                            v13 = (^(v13))
                        }
                    }
                    var v34 int = (fi14 | fi22)
                    v2 -= -2.25
                    _ = v34
                }
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        _ = fi14
        _ = v21
    }
    v35 := v4[v6]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
