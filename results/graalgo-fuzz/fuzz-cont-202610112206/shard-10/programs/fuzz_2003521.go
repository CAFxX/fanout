package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int16, p1 int32) int {
    v7 := GenS1{F0: 65535}
    var v8 float64 = 3.14159
    var v9 int = -908
    switch "hello" {
    case "foo bar":
        var v10 int64 = -9223372036854775808
        _ = v10
    case "":
        v11 := []int{v9, v9}
        _ = v11
    default:
        var v12 float64 = v8
        _ = v12
    }
    v13 := GenS2{F0: 382, F1: 65535}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    _ = v9
    return (v9 - (v9 | v9))
}
func genFn1() uint16 {
    for fi14 := 0; fi14 < 6; fi14++ {
        if (false && (!true)) {
            var v15 string = ("" + "hello")
            v16 := GenS1{F0: 65535}
            if (2 < 127) {
                var v17 int8 = 1
                v17 = (^(v17))
                _ = v17
            }
            _ = v15
            _ = v16
        }
        var v18 string = "hello"
        v19 := []int{4294967295}
        _ = fi14
        _ = v18
        _ = v19
    }
    var v20 float64 = 3.14159
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := []int{(fi21 ^ 32768), (fi21 & (^(fi21))), fi21}
        fmt.Println(v20)
        _ = fi21
        _ = v22
    }
    v20 = v20
    if true {
        var v23 int8 = 127
        _ = v23
    } else {
        if (!true) {
            v20 += (v20 + float64(0.1))
        }
    }
    _ = v20
    return 1
}

func main() {
    v1 := []string{"", ("foo bar" + "foo bar"), ("a" + (("\n\t" + "a") + "\n\t"))}
    var v2 int16 = 344
    var v3 uint8 = 1
    var v4 uint16 = uint16(v3)
    var v5 float32 = float32((1.5 * (1e10 * (float64(-2.25) * -2.25))))
    var v6 int32 = -159
    var v24 int16 = v2
    if false {
        v6 = -463
    } else {
        var v25 int64 = -181
        fmt.Println(v3)
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 string = ("foo bar" + "foo bar")
            var v28 float32 = (1e10 - (v5 + float32(3.14159)))
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 float64 = (((1.5 - (3.14159 - (3.14159 - 0.0))) * -2.25) * (0.1 - -2.25))
            v31 := GenS1{F0: 2147483648}
            if (v5 > 1.5) {
                switch fi29 {
                case 320:
                    var v32 int = (fi29 + 42)
                    _ = v32
                case -1:
                    var v33 string = ("a" + "hello")
                    _ = v33
                case 4294967295:
                    v6 = 1
                default:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        var v35 int32 = (v6 | v6)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch (fi34 >> ((700) & 63)) {
                        case -403:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v36 bool = (!(-678 <= -9223372036854775808))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v36
                        case 669:
                            fmt.Println(v30, v6)
                        case -542:
                            var v37 int16 = (-(v24))
                            _ = v37
                        default:
                            fmt.Println(v35, fi34)
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                v30 -= v30
            }
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        v1 = append(v1, "hello")
        for fi38 := 0; fi38 < 19; fi38++ {
            fmt.Println(v24, v5, v4)
            _ = fi38
        }
        if (128 < int(v6)) {
            fmt.Println(v24, v3)
            fmt.Println(v4, v5, v24)
        } else {
            if (1 > 0) {
                var v39 int8 = 100
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v41 float64 = (0.1 / float64(((3.14159 + v5) * (((v5 * 3.14159) / (float32(0.1) + ((v5 + v5) + 0.1))) - (v5 * 0.1)))))
        _ = fi40
        _ = v41
    }
    var v42 int16 = (-(v2))
    switch 32768 {
    case 4294967295:
        v4 = (v4 | v4)
    case 941:
        if true {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
    case -9223372036854775808:
        v24 = 40
    }
    v6 = (v6 ^ 32)
    v43 := GenS2{F0: (v4 + genFn1()), F1: 770}
    v6 = (v6 + (v6 >> ((255) & 63)))
    v4 = genFn1()
    v44 := GenS2{F0: 2, F1: 750}
    var v45 int = -93
    v5 = v5
    v6 += (^(v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
