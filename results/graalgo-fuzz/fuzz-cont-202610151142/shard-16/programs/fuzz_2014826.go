package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int16, p2 uint16) int64 {
    v4 := GenS1{F0: (p0 | p0), F1: 0, F2: 576}
    p0 += p0
    fmt.Println(p0, p1)
    v5 := GenS1{F0: p0, F1: 42, F2: 386}
    p2 -= 2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]string)
    var v6 int16 = -373
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = 1e10
        var v9 int8 = -46
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 255, F1: 42, F2: 18446744073709551615}
    delete(v3, 583)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int32 = -2147483648
        var v13 uint64 = 182
        switch "" {
        case "\n\t":
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                var v14 uint32 = 955
                _ = v14
            case "\n\t":
                v13 += (v13 % ((1) | 1))
            case "a":
                if ("" <= ("a" + "a")) {
                    _ = len(v3)
                    switch fi11 {
                    case -9223372036854775808:
                        if false {
                            switch fi11 {
                            case 0:
                                v2 += v2
                            }
                        }
                    case 127:
                        switch 362 {
                        case 211:
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                v15 := GenS2{F0: genFn0(255, v6, 42), F1: 42, F2: true}
                                _ = v15
                            }
                        case -916:
                            var v16 int64 = -394
                            _ = v16
                        case -558:
                            v17 := GenS1{F0: 68, F1: 127, F2: 401}
                            _ = v17
                        default:
                            delete(v3, 504)
                        }
                    case -2147483648:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v2 += (v2 / ((fi11) | 1))
                }
            default:
                if true {
                    fmt.Println(v2, fi11)
                    v18 := GenS3{F0: 0, F1: 128}
                    _ = v18
                } else {
                    var v19 int8 = -128
                    fmt.Println(v12)
                    v2 -= (^(fi11))
                    _ = v19
                }
            }
        case "foo bar":
            v13 -= (v13 / ((v13) | 1))
        case "":
            switch 4294967295 {
            case 2147483647:
                var v20 uint = 720
                _ = v20
            case 1:
                v1 = append(v1, "\n\t")
            case -174:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            default:
                for fi21 := 0; fi21 < 12; fi21++ {
                    _ = len(v3)
                    if (true && false) {
                        v12 += (v12 & 509)
                        if true {
                            v13 = 1
                            if len(v1) > 0 {
                                v1[0] = ("a" + "foo bar")
                            }
                            switch 4294967295 {
                            case 366:
                                v2 = v2
                                v12 += (v12 - (v12 >> ((18446744073709551615) & 63)))
                            case 180:
                                fmt.Println(v12)
                            default:
                                v22 := GenS3{F0: 42, F1: 58}
                                _ = v22
                            }
                        } else {
                            fmt.Println(fi21, v6, fi11)
                        }
                        if (255 > 20) {
                            v6 -= (-(v6))
                            fmt.Println(v12, v6, v6)
                        } else {
                            v1 = append(v1, ((("hello" + "a") + "\n\t") + "foo bar"))
                            fmt.Println(v6, v6, v2)
                            if true {
                                var v23 float32 = (0.0 - ((1.5 - 3.14159) - (1.5 - 0.1)))
                                _ = v23
                            } else {
                                fmt.Println(fi11)
                                v13 = (v13 >> ((18446744073709551615) & 63))
                            }
                        }
                    }
                    var v24 uint64 = (v13 / ((18446744073709551615) | 1))
                    _ = fi21
                    _ = v24
                }
            }
        default:
            var v25 bool = (v13 < (v13 | (^(v13))))
            _ = v25
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2, v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 = v6
    v26 := GenS4{F0: v6}
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v6
}
