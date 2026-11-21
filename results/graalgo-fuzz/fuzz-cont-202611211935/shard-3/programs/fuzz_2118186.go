package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 uint32) int32 {
    var v4 string = "foo bar"
    switch "\n\t" {
    case "":
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := []int{fi5, 734, (fi5 ^ fi5), fi5}
            _ = fi5
            _ = v6
        }
    case "hello":
        switch v4 {
        case "\n\t":
            v4 = (v4 + "hello")
        case "a":
            var v7 float32 = 0.0
            _ = v7
        case "foo bar":
            for fi8 := 0; fi8 < 6; fi8++ {
                fmt.Println(v4, p0, v4)
                var v9 uint8 = 1
                v10 := []int{-9223372036854775808}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    default:
        fmt.Println(p0, p0)
    }
    if (1 >= 1) {
        var v11 int8 = 127
        p0 = p0
        _ = v11
    }
    var v12 int64 = 0
    _ = p0
    _ = v12
    _ = v4
    return 0
}
func genFn1(p0 bool) int {
    v13 := []int{-2147483648, 799, 127, -9223372036854775808}
    v14 := GenS1{F0: 199, F1: 42, F2: 32767}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: 441, F1: 0, F2: 0}
        if (!false) {
            fmt.Println(p0)
            switch fi15 {
            case -2147483648:
                p0 = (!p0)
            default:
                if p0 {
                    if (71 != -366) {
                        if ((!((0.0 + 3.14159) <= 3.14159)) || (p0 || (-816 == 363))) {
                            fmt.Println(p0)
                            if len(v13) > 0 {
                                v13[0] = (fi15 - fi15)
                            }
                        } else {
                            p0 = true
                        }
                        if false {
                            p0 = (!false)
                            if p0 {
                                fmt.Println(fi15)
                                var v17 uint64 = 1
                                _ = v17
                            }
                        }
                        v18 := GenS1{F0: 692, F1: genFn0(333), F2: 42}
                        _ = v18
                    } else {
                        fmt.Println(p0, p0)
                    }
                    var v19 uint64 = 2
                    var v20 uint8 = 42
                    _ = v19
                    _ = v20
                } else {
                    var v21 bool = p0
                    var v22 bool = v21
                    _ = v21
                    _ = v22
                }
            }
        }
        var v23 string = ""
        _ = fi15
        _ = v16
        _ = v23
    }
    _ = p0
    _ = v13
    _ = v14
    return -947
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    v3 := make(map[string]bool)
    v2 -= v2
    if true {
        v2 -= v2
    } else {
        v3[("" + ("" + "hello"))] = false
        switch -9223372036854775808 {
        case -1:
            v2 = (^(v2))
        case 642:
            v24 := GenS3{F0: 128, F1: (v2 % ((v2) | 1))}
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v2 = 705
        v2 = (v2 ^ (v2 + (^(v2))))
        var v26 int64 = 921
        _ = fi25
        _ = v26
    }
    switch -1 {
    case -2147483648:
        v2 += (^(v2))
    case 1:
        v27 := GenS4{F0: ((-2.25 >= ((((3.14159 * (0.0 + 0.1)) / 0.1) / float64(1e10)) - (((0.1 / 1.0) - 0.1) + float64(((((((3.14159 / (3.14159 * ((1e10 + 3.14159) + -2.25))) - 0.0) / 1.0) / 1.5) + 0.1) + float32((-2.25 + 0.0))))))) || (!("foo bar" >= "foo bar"))), F1: 128, F2: 84}
        _ = v27
    default:
        v2 = (^(v2))
    }
    switch "foo bar" {
    case "foo bar":
        v28 := v1[65535]
        _ = v28
    case "":
        fmt.Println(v2)
    }
    if false {
        for fi29 := 0; fi29 < 14; fi29++ {
            v2 -= (v2 + 1)
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS1{F0: v2, F1: -1, F2: 32767}
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    }
    v32 := []int{-2147483648, 566}
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 int64 = 695
        _ = fi33
        _ = v34
    }
    fmt.Println(v2)
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    var v35 bool = false
    var v36 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
}
