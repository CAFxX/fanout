package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := []uint32{4294967295}
    var v4 uint8 = 255
    v5 := GenS1{F0: 612, F1: 255}
    fmt.Println(v4)
    v4 = (v4 ^ (^(v4)))
    switch 32767 {
    case 285:
        var v6 uint64 = 461
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if (!(true && ((fi7 + (^(fi7))) < fi7))) {
            fmt.Println(fi7, fi7)
        }
        var v8 uint = 65535
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 2, F1: -316, F2: -1}
    var v10 string = "a"
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v11 bool = false
        _ = v11
    case 128:
        v4 = (v4 >> ((1) & 63))
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint8 = 128
            _ = fi12
            _ = v13
        }
    }
    if (427 <= 42) {
        v4 = v4
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int16 = 42
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16, fi14)
            if false {
                var v17 bool = ((!true) || (!(!(v4 > 238))))
                switch (fi14 | (^(fi16))) {
                case -558:
                    v10 = ((("a" + v10) + "foo bar") + (v10 + (v10 + ((v10 + v10) + v10))))
                case -2147483648:
                    if v17 {
                        v15 = (v15 / ((int16(fi14)) | 1))
                    } else {
                        _ = len(v3)
                        v18 := GenS3{F0: (v15 - -32768), F1: true}
                        var v19 int = 2147483647
                        if ((!v17) || (0 >= v4)) {
                            switch ((v10 + v10) + v10) {
                            case "":
                                fmt.Println(v15, v17, v19)
                            case "foo bar":
                                var v20 int64 = 737
                                _ = v20
                            }
                            if v17 {
                                var v21 int8 = 127
                                _ = v21
                            }
                            var v22 int = -2147483648
                            _ = v22
                        } else {
                            switch (fi14 % (((v19 + v19)) | 1)) {
                            case 127:
                                v4 -= v4
                            case 32768:
                                var v23 int32 = 1
                                _ = v23
                            case -815:
                                switch v19 {
                                case 2147483647:
                                    fmt.Println(v15)
                                    fmt.Println(v4, v4)
                                case -737:
                                    var v24 uint16 = 193
                                    _ = v24
                                default:
                                    switch 699 {
                                    case -388:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 32768:
                                        v4 = (^(v4))
                                    case 32767:
                                        var v25 uint8 = 0
                                        _ = v25
                                    }
                                }
                            default:
                                switch -231 {
                                case -480:
                                    var v26 float32 = 3.14159
                                    _ = v26
                                case 2147483647:
                                    delete(v2, -9223372036854775808)
                                }
                            }
                            fmt.Println(v17, v15, fi16)
                            v4 = v4
                        }
                        _ = v18
                        _ = v19
                    }
                default:
                    v15 += v15
                }
                _ = v17
            } else {
                if (((520 != 1) || false) || ((v4 - 127) < (^(v4)))) {
                    v4 -= v4
                } else {
                    fmt.Println(fi14, fi16, v10)
                    v15 -= int16(v4)
                    fmt.Println(fi16, v10, v4)
                }
                fmt.Println(v10)
                v15 += v15
            }
            _ = fi16
        }
        if (true || (562 > 438)) {
            v3 = append(v3, 65535)
            v27 := GenS4{F0: 2147483647, F1: 4294967295}
            _ = v27
        } else {
            v15 += int16(v4)
            switch (fi14 * -939) {
            case 42:
                fmt.Println(v15)
            default:
                for fi28 := 0; fi28 < 1; fi28++ {
                    if (((1e10 - 3.14159) >= float32((0.1 * float64(1e10)))) && ((0.0 == 1.5) || (!true))) {
                        v4 = (v4 * v4)
                        v4 = (v4 - (v4 % (((v4 >> ((731) & 63))) | 1)))
                    }
                    _ = fi28
                }
            }
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
