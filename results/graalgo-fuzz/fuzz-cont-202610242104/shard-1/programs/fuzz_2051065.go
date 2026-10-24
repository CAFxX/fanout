package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 string; F2 int }
func genFn0(p0 uint64, p1 uint64) uint32 {
    if false {
        var v7 int32 = 2147483647
        v8 := GenS1{F0: true, F1: 2, F2: 1}
        _ = v7
        _ = v8
    }
    var v9 int32 = 2147483647
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v9
    return 495
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 42
    v3 := []uint32{}
    var v4 float64 = float64(0.1)
    var v5 int32 = 641
    var v6 bool = true
    if v6 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1[0] = "a"
        v5 = v5
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int32 = v5
        switch "foo bar" {
        case "\n\t":
            var v12 float64 = v4
            _ = v12
        }
        var v13 int = (fi10 + (fi10 >> ((65535) & 63)))
        _ = fi10
        _ = v11
        _ = v13
    }
    if (v6 && (-9223372036854775808 == -9223372036854775808)) {
        switch "\n\t" {
        case "a":
            var v14 string = "foo bar"
            _ = v14
        case "foo bar":
            var v15 uint = 0
            _ = v15
        case "\n\t":
            switch ("foo bar" + (("foo bar" + "") + "\n\t")) {
            case "\n\t":
                v6 = (102 <= 42)
            case "a":
                v2 += (v2 & v2)
            case "hello":
                fmt.Println(v2)
            }
        default:
            v6 = (523 < v2)
        }
    }
    if v6 {
        switch 595 {
        case 826:
            v16 := GenS1{F0: (!v6), F1: (v2 % (((v2 >> ((42) & 63))) | 1)), F2: 255}
            _ = v16
        case 9223372036854775807:
            for fi17 := 0; fi17 < 9; fi17++ {
                v5 = (v5 * (v5 ^ -484))
                v18 := GenS1{F0: v6, F1: 339, F2: 255}
                _ = fi17
                _ = v18
            }
        case 128:
            _ = len(v3)
        default:
            v4 = 1e10
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v19 float64 = v4
        delete(v1, -9223372036854775808)
        for fi20 := 0; fi20 < 16; fi20++ {
            if (-2.25 < 0.1) {
                v6 = (9223372036854775807 >= -9223372036854775808)
                v3 = append(v3, 2147483647)
                if (true && v6) {
                    if (!v6) {
                        v21 := GenS2{F0: 1, F1: ("a" + (("foo bar" + "foo bar") + "hello")), F2: fi20}
                        switch fi20 {
                        case 1:
                            v2 = (v2 % (((v2 ^ 65535)) | 1))
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                switch (fi20 - fi20) {
                                case -1:
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        if true {
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(fi20, v5)
                                            fmt.Println(v6)
                                            v6 = v6
                                        }
                                        if (v2 >= (v2 << ((65535) & 63))) {
                                            var v23 int16 = int16(v5)
                                            var v24 float64 = (1e10 - float64(float32(((-2.25 / ((((v19 * ((float64(-2.25) + 0.1) + ((float64((((1e10 + (3.14159 - -2.25)) - ((0.1 - -2.25) / 1e10)) + 3.14159)) * v19) * float64((((1.5 / float32((v4 + v19))) - -2.25) - -2.25))))) * (v19 * (((((((v4 / 3.14159) * 0.1) * (v19 - (-2.25 + v19))) - (v19 + (1e10 / (float64(1.5) / (3.14159 - (float64(0.1) * v4)))))) - 1.5) - -2.25) * ((1e10 * -2.25) + (float64(0.1) / float64(0.1)))))) / 1e10) - v19)) * (0.1 + (0.0 - v19))))))
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi22
                                    }
                                case 42:
                                    var v25 string = ""
                                    _ = v25
                                case -2147483648:
                                    v26 := v1[(-(fi20))]
                                    _ = v26
                                default:
                                    v2 -= v2
                                }
                            case "foo bar":
                                var v27 float64 = v19
                                _ = v27
                            }
                        }
                        _ = v21
                    }
                    var v28 uint32 = v2
                    _ = v28
                }
            } else {
                fmt.Println(v19)
                var v29 uint32 = 779
                v4 = v4
                _ = v29
            }
            _ = fi20
        }
        _ = v19
    }
    v2 = v2
    v2 = (v2 | v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v30 int8 = 0
    v31 := v1[42]
    v1[127] = v31
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v32 uint = 255
    v33 := GenS1{F0: ((v32 / (((v32 >> (((v32 * (^(v32)))) & 63))) | 1)) != (v32 + 1)), F1: (v2 + 2), F2: (v32 % ((431) | 1))}
    if (v6 && ((!false) && (((v32 * (v32 ^ (^(v32)))) < (^(v32))) && v6))) {
        v5 += (v5 ^ -2147483648)
    } else {
        var v34 uint32 = genFn0(819, 2)
        var v35 uint64 = 42
        _ = v34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
