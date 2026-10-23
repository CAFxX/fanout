package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 bool, p1 int32, p2 uint) uint32 {
    v4 := []int{-564, -667, int(p1), -2147483648}
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            switch fi6 {
            case 9223372036854775807:
                var v7 int16 = 0
                _ = v7
            case 32768:
                p1 = (-(p1))
            case 32767:
                if (p0 || (true || (127 >= 0))) {
                    fmt.Println(fi5)
                }
            }
            p0 = (127 >= 255)
            var v8 int32 = p1
            _ = fi6
            _ = v8
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            p2 -= p2
            var v10 uint16 = 859
            _ = fi9
            _ = v10
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1(p0 uint8, p1 int64, p2 uint16) uint32 {
    fmt.Println(p2)
    if (0 >= 1000) {
        if true {
            p2 = 1
            p1 -= (p1 ^ 0)
            switch "a" {
            case "a":
                var v11 int8 = -128
                _ = v11
            case "foo bar":
                p0 = p0
            case "hello":
                var v12 uint32 = 286
                _ = v12
            }
        }
        fmt.Println(p0, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []int{-266, 414, 2147483647, 1}
    var v13 uint8 = 2
    v3 = append(v3, -120)
    _ = len(v3)
    if false {
        var v14 uint8 = v13
        var v15 uint8 = (v14 * (v13 * v13))
        _ = v14
        _ = v15
    }
    var v16 float64 = 3.14159
    if (3.14159 >= 0.1) {
        v17 := GenS1{F0: 18446744073709551615, F1: 762}
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 3; fi18++ {
                v13 -= v13
                var v19 int16 = 0
                v13 = (v13 & v13)
                _ = fi18
                _ = v19
            }
        case "foo bar":
            switch "hello" {
            case "\n\t":
                v16 -= v16
            case "foo bar":
                delete(v1, "hello")
            case "":
                _ = len(v3)
            default:
                for fi20 := 0; fi20 < 10; fi20++ {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS2{F0: 0, F1: 32767}
                        if true {
                            if true {
                                var v23 uint8 = v13
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v23
                            }
                            if ((1e10 * ((0.0 * (1e10 - (((0.0 + 0.0) + 1.5) + ((1e10 * -2.25) / ((3.14159 / ((0.1 / 1.5) * 3.14159)) - float32(float64(1.5))))))) + 0.0)) <= ((((0.0 * 0.0) + 0.1) - 0.1) * (1e10 + (1e10 - 0.0)))) {
                                v24 := v1[""]
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = (v24 ^ v24)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = fi21
                                    }
                                }
                                _ = v24
                            } else {
                                switch "" {
                                case "a":
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        v25 := GenS2{F0: 292, F1: 32767}
                                        _ = v25
                                    }
                                default:
                                    var v26 uint8 = (^(v13))
                                    _ = v26
                                }
                            }
                            if len(v3) > 0 {
                                v3[0] = (fi20 - (fi20 * (fi21 | (fi20 ^ fi20))))
                            }
                        } else {
                            if true {
                                var v27 int64 = 1
                                v13 = 42
                                _ = v27
                            }
                            _ = len(v2)
                        }
                        delete(v2, fi20)
                        _ = fi21
                        _ = v22
                    }
                    v28 := GenS3{F0: genFn0((0 != 32767), -1, 4294967295)}
                    _ = fi20
                    _ = v28
                }
            }
        case "":
            v13 = v13
        }
        switch 2147483647 {
        case 1:
            switch -262 {
            case -2147483648:
                v16 += 0.0
            }
        case 42:
            var v29 uint64 = uint64(v13)
            _ = v29
        }
        _ = v17
    } else {
        _ = len(v3)
        for fi30 := 0; fi30 < 8; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                v32 := GenS2{F0: 1, F1: 0}
                v16 += 0.0
                _ = fi31
                _ = v32
            }
            if false {
                var v33 uint16 = 32767
                _ = v33
            }
            _ = fi30
        }
        if false {
            var v34 bool = true
            _ = v34
        }
    }
    v13 += (v13 + v13)
    for fi35 := 0; fi35 < 5; fi35++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v3 = append(v3, (fi35 + 32768))
        _ = fi35
    }
    fmt.Println(v13, v13)
    fmt.Println(v16, v16)
    var v36 int32 = -583
    if false {
        delete(v2, 32767)
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        for fi38 := 0; fi38 < 8; fi38++ {
            switch "" {
            case "hello":
                fmt.Println(v13, v36, fi37)
            default:
                v39 := GenS1{F0: 42, F1: genFn1(v13, -1, 255)}
                _ = v39
            }
            _ = fi38
        }
        v36 += v36
        _ = fi37
    }
    v13 += v13
    var v40 bool = ((!("\n\t" <= "a")) && false)
    v41 := GenS3{F0: 42}
    var v42 float32 = (float32(v16) * ((1.5 + -2.25) - (1e10 * -2.25)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v40
    _ = v41
    _ = v42
}
