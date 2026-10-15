package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = ((false || (!true)) || (true || (false && (!false))))
    var v2 string = "foo bar"
    var v3 string = v2
    var v4 float32 = ((0.0 - -2.25) - (float32((float64(0.1) - 1e10)) + 3.14159))
    v5 := make(map[uint32]string)
    v6 := make(map[int]int)
    if v1 {
        switch 2147483647 {
        case 0:
            v7 := []int{128, 649, 127}
            _ = v7
        case 1:
            switch -1 {
            case 4294967295:
                var v8 uint16 = 0
                _ = v8
            default:
                v4 = float32(1e10)
            }
        case 447:
            _ = len(v5)
        }
        var v9 uint = 133
        v10 := GenS1{F0: -923, F1: 0, F2: (v2 + v3)}
        _ = v10
        _ = v9
    } else {
        v6[127] = -580
        var v11 float32 = (-2.25 / -2.25)
        v12 := v6[32768]
        _ = v11
        _ = v12
    }
    v2 = v3
    fmt.Println(v3, v3, v2)
    switch v2 {
    case "foo bar":
        var v13 uint16 = 51
        _ = v13
    default:
        v14 := []int{32767}
        _ = v14
    }
    if (0 != -1) {
        var v15 uint8 = 128
        _ = v15
    } else {
        var v16 int32 = -1
        _ = v16
    }
    var v17 uint = 18446744073709551615
    var v18 uint32 = 344
    if true {
        v18 -= (v18 * (v18 | (v18 / (((v18 | 1)) | 1))))
        v1 = v1
        fmt.Println(v18)
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            _ = len(v6)
            v17 = (^(v17))
            _ = fi19
        }
        var v20 bool = (100 < 127)
        var v21 float32 = 3.14159
        _ = v20
        _ = v21
    }
    v4 -= float32(1.5)
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint32 = (v18 * 2147483648)
        var v24 uint16 = 2
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 1, F1: 127, F2: 1}
    if true {
        v26 := GenS3{F0: 128}
        fmt.Println(v1, v4, v1)
        _ = v26
    } else {
        v27 := GenS1{F0: 0, F1: 1, F2: (v3 + (v2 + (("\n\t" + (v2 + ("foo bar" + (v3 + v3)))) + v3)))}
        _ = v27
    }
    v17 += (v17 ^ 635)
    if v1 {
        if (true || v1) {
            var v28 uint8 = 127
            switch "\n\t" {
            case "\n\t":
                delete(v5, v18)
            }
            _ = v28
        } else {
            switch -2147483648 {
            case -11:
                v1 = v1
            case 381:
                v17 = 18446744073709551615
            }
        }
        var v29 uint = 65535
        var v30 int8 = 79
        _ = v29
        _ = v30
    }
    switch 42 {
    case 127:
        if v1 {
            for fi31 := 0; fi31 < 6; fi31++ {
                fmt.Println(v17, v18, v17)
                if (937 <= 986) {
                    if (((v18 / ((v18) | 1)) > v18) && (!v1)) {
                        var v32 int64 = -1
                        v1 = (v4 > 0.0)
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            _ = len(v5)
                            var v34 uint32 = 2147483647
                            v35 := GenS1{F0: 886, F1: -796, F2: v2}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                    var v36 int8 = 127
                    _ = v36
                } else {
                    var v37 float64 = (0.1 * ((-2.25 - 1.5) - 3.14159))
                    if ((128 <= fi31) || ("\n\t" != v2)) {
                        v17 = (v17 ^ v17)
                        v3 = v2
                        for fi38 := 0; fi38 < 3; fi38++ {
                            v39 := []int{(fi31 << ((v17) & 63)), -9223372036854775808, fi38}
                            v2 = v2
                            v3 = "\n\t"
                            v39 = append(v39, fi31)
                            if true {
                                fmt.Println(v18, v18)
                                fmt.Println(v37)
                                if len(v39) >= 2 {
                                    v39 = v39[1:len(v39)]
                                }
                                if len(v39) > 0 {
                                    v39[0] = (fi38 ^ (fi31 - (fi38 - 32767)))
                                }
                            }
                            _ = fi38
                            _ = v39
                        }
                    }
                    for fi40 := 0; fi40 < 15; fi40++ {
                        var v41 int = 9223372036854775807
                        _ = fi40
                        _ = v41
                    }
                    _ = v37
                }
                for fi42 := 0; fi42 < 11; fi42++ {
                    if v1 {
                        var v43 uint8 = 9
                        delete(v6, int(v43))
                        _ = v43
                    }
                    v4 = ((3.14159 / (float32(((0.1 + 3.14159) + 3.14159)) / (float32((0.0 * 1e10)) + float32(((3.14159 * 1e10) / -2.25))))) * v4)
                    var v44 int8 = -128
                    _ = fi42
                    _ = v44
                }
                _ = fi31
            }
        }
    case -227:
        v45 := []int{128}
        _ = v45
    default:
        var v46 uint64 = 4294967295
        _ = v46
    }
    v47 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
