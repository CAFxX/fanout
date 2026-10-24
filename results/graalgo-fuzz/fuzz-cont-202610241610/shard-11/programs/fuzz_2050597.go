package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint8 = 42
    v2 := []uint32{641}
    var v3 int32 = 2147483647
    v4 := make(map[int]string)
    v5 := GenS1{F0: 2147483648, F1: -462}
    fmt.Println(v3, v3)
    v4[-2147483648] = "hello"
    switch -342 {
    case -1:
        v3 = v3
    case 4294967295:
        v3 = (^(v3))
    case 128:
        var v6 int64 = -989
        _ = v6
    }
    if (0.0 != 1.5) {
        v1 += v1
    }
    v1 = 255
    delete(v4, 127)
    if false {
        v7 := GenS2{F0: -1, F1: 1, F2: 100}
        switch ("" + (("a" + "\n\t") + "\n\t")) {
        case "a":
            fmt.Println(v3, v1)
        case "\n\t":
            var v8 float32 = 3.14159
            _ = v8
        case "hello":
            v9 := GenS1{F0: uint32(v1), F1: 1}
            _ = v9
        default:
            v3 = v3
        }
        _ = v7
    }
    v3 += v3
    var v10 string = "\n\t"
    switch -74 {
    case 0:
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 65535, F1: 573}
            if false {
                _ = len(v2)
                v1 -= 0
            }
            fmt.Println(fi11, v1)
            _ = fi11
            _ = v12
        }
    case 2147483647:
        v10 = v10
    case -966:
        delete(v4, 9223372036854775807)
    default:
        fmt.Println(v1, v1, v1)
    }
    if true {
        var v13 uint32 = 713
        _ = v13
    }
    if false {
        if false {
            _ = len(v4)
            if (277 <= 0) {
                if true {
                    if true {
                        v1 += 86
                    } else {
                        for fi14 := 0; fi14 < 9; fi14++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi14
                        }
                    }
                    var v15 int32 = v3
                    switch v10 {
                    case "hello":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 uint16 = 1
                            if true {
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    v1 = 93
                                    if (false && (((v1 + v1) > (v1 * v1)) && (v1 >= (v1 * (^(v1)))))) {
                                        v10 = v10
                                    }
                                    _ = fi18
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                    case "":
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            fmt.Println(v10)
                        }
                    case "a":
                        v19 := GenS2{F0: -202, F1: 73, F2: -6}
                        _ = v19
                    }
                    _ = v15
                }
                var v20 int64 = -1
                _ = v20
            }
        } else {
            var v21 uint16 = uint16(v1)
            _ = v21
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            delete(v4, -908)
            if false {
                switch (fi22 - fi22) {
                case -1:
                    var v23 uint16 = 0
                    _ = v23
                case 88:
                    var v24 int = (-(fi22))
                    _ = v24
                case 345:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 int16 = 0
                        v1 = 220
                        _ = fi25
                        _ = v26
                    }
                default:
                    if (true || (150 != 42)) {
                        v27 := GenS1{F0: 184, F1: 42}
                        var v28 uint32 = 2
                        v29 := GenS1{F0: v28, F1: 1000}
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        delete(v4, 32768)
                    }
                }
            }
            if true {
                fmt.Println(v1)
                if len(v2) > 0 {
                    v2[0] = 717
                }
                v30 := GenS2{F0: 0, F1: 42, F2: 55}
                _ = v30
            } else {
                switch 2147483647 {
                case 127:
                    var v31 int8 = 56
                    _ = v31
                default:
                    _ = len(v2)
                }
                v1 += 0
                v1 -= v1
            }
            _ = fi22
        }
        var v32 uint16 = 255
        _ = v32
    } else {
        switch v10 {
        case "hello":
            if (false && ((1 < 0) && (((!(!true)) && (true || false)) && false))) {
                _ = len(v4)
                var v33 string = v10
                fmt.Println(v1, v10)
                _ = v33
            }
        case "a":
            v34 := GenS3{F0: 0, F1: 143}
            _ = v34
        case "":
            switch 2147483647 {
            case 42:
                for fi35 := 0; fi35 < 10; fi35++ {
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    v10 = ""
                    if (2147483647 != 2147483647) {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            var v37 int32 = (v3 ^ v3)
                            if ((float64((float32(1e10) / (1e10 + 0.0))) < 1.5) || ((fi36 ^ (fi35 / ((-9223372036854775808) | 1))) > (fi36 + 127))) {
                                v2 = append(v2, 2)
                            }
                            v10 = v10
                            _ = fi36
                            _ = v37
                        }
                    }
                    _ = fi35
                }
            }
        default:
            switch -2147483648 {
            case 893:
                if false {
                    var v38 int16 = 42
                    var v39 uint = 1
                    v39 += v39
                    _ = v38
                    _ = v39
                }
            case 128:
                var v40 bool = (65535 != 1)
                _ = v40
            case 0:
                v41 := v4[611]
                _ = v41
            default:
                v42 := GenS2{F0: 1000, F1: -756, F2: 43}
                _ = v42
            }
        }
    }
    v43 := v4[-2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
