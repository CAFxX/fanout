package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0() string {
    var v4 string = "a"
    v4 = v4
    v4 = ((v4 + v4) + v4)
    _ = v4
    return v4
}
func genFn1(p0 int32, p1 uint16) int16 {
    var v5 int16 = 833
    var v6 float32 = 0.1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []int{9223372036854775807, -405}
    var v2 uint8 = 255
    v3 := make(map[int]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v8 int32 = -754
    v2 = (v2 % ((v2) | 1))
    for fi9 := 0; fi9 < 14; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
    }
    switch "" {
    case "\n\t":
        if ((4294967295 == 4294967295) && false) {
            var v10 uint32 = 1
            v11 := GenS1{F0: 0}
            _ = v10
            _ = v11
        }
    default:
        var v12 uint16 = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 0}
        v8 = 0
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v8 -= -2147483648
            if (0 > -32768) {
                switch "a" {
                case "\n\t":
                    v2 = v2
                case "a":
                    v8 = 254
                case "hello":
                    fmt.Println(v2, v2)
                    switch fi16 {
                    case -1:
                        var v17 int64 = 139
                        _ = v17
                    case 127:
                        var v18 float32 = -2.25
                        _ = v18
                    case -25:
                        v8 -= -1
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v8 += (-(v8))
            }
            _ = fi16
        }
        if false {
            var v19 bool = (!(498 > 410))
            fmt.Println(v8, fi15)
            _ = v19
        }
        _ = fi15
    }
    var v20 int8 = 100
    v8 = (v8 / ((v8) | 1))
    v20 = (-(v20))
    switch 1 {
    case 128:
        var v21 int32 = v8
        _ = v21
    }
    v20 += -128
    if false {
        var v22 int32 = (v8 - v8)
        var v23 uint32 = 1
        for fi24 := 0; fi24 < 14; fi24++ {
            switch "foo bar" {
            case "hello":
                v25 := GenS1{F0: v22}
                _ = v25
            }
            _ = fi24
        }
        _ = v22
        _ = v23
    } else {
        var v26 string = ("" + (("\n\t" + "\n\t") + "foo bar"))
        if true {
            v20 -= (^(v20))
            var v27 float64 = (1e10 * float64((0.1 + 3.14159)))
            switch 42 {
            case -1:
                v28 := GenS2{F0: 244, F1: -32768, F2: -837}
                _ = v28
            default:
                v29 := GenS3{F0: (v8 | (^(v8))), F1: 18446744073709551615}
                _ = v29
            }
            _ = v27
        } else {
            if len(v1) > 0 {
                v1[0] = 128
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 20; fi31++ {
                    delete(v3, 4294967295)
                    v2 -= (^(v2))
                    if true {
                        v26 = "\n\t"
                    } else {
                        v3[fi30] = v26
                    }
                    v20 = v20
                    _ = fi31
                }
                for fi32 := 0; fi32 < 20; fi32++ {
                    if (true && true) {
                        if false {
                            switch fi32 {
                            case 0:
                                var v33 uint8 = (^(v2))
                                _ = v33
                            case -902:
                                var v34 int = -878
                                _ = v34
                            default:
                                if true {
                                    var v35 uint8 = 1
                                    _ = v35
                                } else {
                                    fmt.Println(fi32, v26)
                                    v8 = v8
                                    v26 = ""
                                    v2 = (v2 + v2)
                                }
                            }
                            var v36 uint32 = 1
                            v20 += (v20 * v20)
                            _ = v36
                        } else {
                            if false {
                                var v37 float32 = 3.14159
                                var v38 bool = (!false)
                                var v39 uint16 = 1
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = fi30
                                }
                            }
                        }
                    } else {
                        var v40 int32 = v8
                        v2 = v2
                        _ = v40
                    }
                    v41 := GenS4{F0: false}
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    _ = fi32
                    _ = v41
                }
                switch 32768 {
                case 127:
                    v42 := v3[(-(fi30))]
                    _ = v42
                default:
                    v26 = "hello"
                }
                _ = fi30
            }
        }
        for fi43 := 0; fi43 < 11; fi43++ {
            fmt.Println(v20, v2, v2)
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 int32 = v8
                switch (fi44 * -29) {
                case 179:
                    v46 := GenS1{F0: (-(v8))}
                    _ = v46
                }
                switch -1 {
                case 374:
                    v47 := GenS4{F0: false}
                    _ = v47
                case -9223372036854775808:
                    var v48 int32 = v45
                    _ = v48
                default:
                    v1 = append(v1, (fi44 - fi44))
                }
                _ = fi44
                _ = v45
            }
            v49 := GenS1{F0: 1}
            _ = fi43
            _ = v49
        }
        _ = v26
    }
    var v50 uint64 = 18446744073709551615
    switch ("foo bar" + (("hello" + "foo bar") + "foo bar")) {
    case "foo bar":
        if (v2 <= 42) {
            var v51 uint16 = 161
            _ = v51
        }
    case "hello":
        for fi52 := 0; fi52 < 16; fi52++ {
            fmt.Println(v20)
            var v53 int16 = 0
            _ = fi52
            _ = v53
        }
    case "\n\t":
        var v54 float64 = 0.0
        _ = v54
    }
    if (-90 >= v20) {
        fmt.Println(v8, v50)
    } else {
        v2 += v2
        v55 := GenS2{F0: 0, F1: 175, F2: 859}
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v50
    _ = v8
}
