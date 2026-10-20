package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 string }

func main() {
    var v1 string = "a"
    var v2 int = -771
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := []int{}
    for fi6 := 0; fi6 < 3; fi6++ {
        if (true || true) {
            var v7 int8 = 100
            if false {
                v8 := GenS1{F0: true, F1: -32768, F2: 241}
                var v9 float64 = 1e10
                _ = v8
                _ = v9
            } else {
                var v10 float32 = -2.25
                _ = v10
            }
            _ = v7
        }
        _ = fi6
    }
    if (-1 > 1000) {
        switch 9223372036854775807 {
        case 127:
            v5 = append(v5, v2)
        }
        v5 = append(v5, v2)
    }
    v11 := GenS2{F0: -128, F1: 0, F2: 255}
    v1 = ("a" + "foo bar")
    if (9223372036854775807 == 9223372036854775807) {
        if (false || (216 >= 128)) {
            fmt.Println(v1, v1)
            var v12 int16 = 1000
            switch "" {
            case "hello":
                var v13 bool = (3.14159 > 1.5)
                _ = v13
            case "":
                v14 := GenS1{F0: (!(false && true)), F1: (-(v12)), F2: 219}
                _ = v14
            default:
                delete(v4, (^(v2)))
            }
            _ = v12
        }
    }
    v2 += v2
    switch (v2 ^ -9223372036854775808) {
    case 4294967295:
        var v15 int8 = 58
        _ = v15
    default:
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint32 = 10
            switch v1 {
            case "\n\t":
                v18 := GenS1{F0: false, F1: -883, F2: 4294967295}
                _ = v18
            case "hello":
                v3["hello"] = true
            default:
                v19 := GenS3{F0: 128, F1: 33}
                _ = v19
            }
            _ = fi16
            _ = v17
        }
    }
    if (!(255 != 1)) {
        var v20 int8 = 103
        _ = v20
    }
    fmt.Println(v1)
    var v21 string = "foo bar"
    v22 := GenS1{F0: true, F1: int16(v2), F2: 18446744073709551615}
    v1 = v21
    if true {
        var v23 uint16 = 820
        for fi24 := 0; fi24 < 6; fi24++ {
            if ((^(v2)) <= v2) {
                var v25 uint = 744
                v23 = v23
                if (!((v25 ^ (v25 + (v25 + 2))) > (^(v25)))) {
                    var v26 int32 = -2147483648
                    var v27 float32 = float32((3.14159 / 1.5))
                    for fi28 := 0; fi28 < 10; fi28++ {
                        if (0 >= 417) {
                            v29 := GenS4{F0: ("\n\t" + v21)}
                            v2 += (^(v2))
                            switch v2 {
                            case 492:
                                _ = len(v5)
                            case 9223372036854775807:
                                if true {
                                    v30 := GenS2{F0: 100, F1: v23, F2: v23}
                                    _ = v30
                                }
                            }
                            _ = v29
                        }
                        var v31 int32 = (v26 | -2147483648)
                        var v32 int = fi24
                        _ = fi28
                        _ = v31
                        _ = v32
                    }
                    _ = v26
                    _ = v27
                } else {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 int8 = 1
                        v21 = v1
                        _ = fi33
                        _ = v34
                    }
                }
                _ = v25
            } else {
                v23 += (^(v23))
            }
            delete(v4, (v2 / (((v2 * (fi24 << ((93) & 63)))) | 1)))
            var v35 int32 = -2147483648
            _ = fi24
            _ = v35
        }
        var v36 int8 = 100
        _ = v23
        _ = v36
    } else {
        v2 += v2
    }
    v2 -= -9223372036854775808
    var v37 float64 = (float64(float32(3.14159)) + (((1.5 / 3.14159) + (((1.5 * ((float64(0.1) - (0.0 - float64(0.1))) - 1e10)) * 0.0) * 3.14159)) / 1.5))
    var v38 int8 = 127
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
