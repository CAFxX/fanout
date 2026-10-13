package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int64 }
func genFn0(p0 string) int16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = (0.1 + 1.5)
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    switch p0 {
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(p0, p0)
            _ = fi9
        }
    case "\n\t":
        p0 = "hello"
    default:
        var v10 int8 = -128
        _ = v10
    }
    _ = p0
    return -460
}

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 1
    var v3 uint16 = 138
    v4 := make(map[int]uint32)
    var v5 uint16 = (v3 * (v3 + (^(v3))))
    var v6 uint = 2
    v6 -= v6
    if true {
        v11 := v4[1]
        _ = v11
    } else {
        if (((-1 > 458) && false) || (true && false)) {
            fmt.Println(v1)
            var v12 int64 = 1
            var v13 uint64 = v2
            _ = v12
            _ = v13
        } else {
            v3 = 1
            if (!false) {
                if false {
                    fmt.Println(v5, v5, v5)
                    v3 = v3
                    fmt.Println(v1, v6)
                }
            }
        }
    }
    v1 = v1
    var v14 int16 = -1
    v15 := []int{4294967295}
    for fi16 := 0; fi16 < 13; fi16++ {
        v4[fi16] = 793
        v17 := GenS1{F0: 1}
        _ = fi16
        _ = v17
    }
    fmt.Println(v2)
    var v18 string = ((v1 + ("\n\t" + "\n\t")) + v1)
    var v19 uint16 = 65535
    fmt.Println(v14, v3)
    if false {
        switch v18 {
        case "\n\t":
            fmt.Println(v1, v19, v19)
        case "":
            var v20 int = 514
            _ = v20
        case "hello":
            if (float32((float64((1e10 / (1e10 + float32(((1.5 + 1.5) * (1e10 / 1.5)))))) + 0.0)) < (3.14159 / (1.5 / 1.5))) {
                var v21 uint32 = 0
                if false {
                    var v22 int8 = 100
                    if false {
                        v23 := GenS2{F0: 42}
                        fmt.Println(v1, v6, v2)
                        v2 += (v2 + (v2 ^ (^(v2))))
                        _ = v23
                    }
                    v24 := GenS3{F0: false, F1: -9223372036854775808}
                    _ = v22
                    _ = v24
                } else {
                    v1 = v18
                    switch v18 {
                    case "foo bar":
                        for fi25 := 0; fi25 < 11; fi25++ {
                            v26 := GenS1{F0: 545}
                            _ = fi25
                            _ = v26
                        }
                    case "a":
                        fmt.Println(v18, v5, v18)
                    case "":
                        if len(v15) > 0 {
                            v15[0] = -573
                        }
                    }
                    if ((!(!true)) || (-61 != 81)) {
                        v2 = 4294967295
                    }
                }
                _ = v21
            } else {
                v27 := GenS4{F0: (v6 | v6), F1: (v14 + 1000)}
                v28 := GenS2{F0: uint8(v2)}
                v29 := GenS1{F0: 509}
                _ = v27
                _ = v28
                _ = v29
            }
        default:
            v3 += (v19 & (^(v3)))
        }
    } else {
        if false {
            for fi30 := 0; fi30 < 18; fi30++ {
                fmt.Println(v1, v14, v14)
                for fi31 := 0; fi31 < 5; fi31++ {
                    if (!false) {
                        v18 = v18
                    } else {
                        var v32 float32 = 3.14159
                        var v33 int32 = -2147483648
                        _ = v32
                        _ = v33
                    }
                    _ = fi31
                }
                switch v18 {
                case "a":
                    if (((1e10 / 1e10) / 1.5) == 3.14159) {
                        v4[fi30] = 1
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 int32 = 1
                            fmt.Println(v14, v2, v35)
                            v36 := GenS5{F0: 4294967295, F1: "\n\t"}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        if (-1 <= 108) {
                            v37 := GenS3{F0: false, F1: fi30}
                            v38 := GenS6{F0: 57}
                            _ = v37
                            _ = v38
                        }
                    }
                case "hello":
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v40 := GenS7{F0: (v19 | (v19 % (((v3 % (((v5 & 54)) | 1))) | 1))), F1: 0, F2: -9223372036854775808}
                        var v41 int32 = 1
                        _ = len(v15)
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                case "\n\t":
                    v3 -= 65535
                default:
                    var v42 uint8 = 1
                    _ = v42
                }
                _ = fi30
            }
            var v43 uint64 = uint64(v5)
            v44 := GenS5{F0: 2, F1: ("hello" + v1)}
            _ = v43
            _ = v44
        }
        fmt.Println(v19, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
