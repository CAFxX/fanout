package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 string, p1 bool, p2 int8) uint16 {
    switch 44 {
    case 1:
        var v4 int64 = 559
        _ = v4
    case 996:
        var v5 uint64 = 2
        _ = v5
    case -18:
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 32767, F1: 382}
            fmt.Println(fi6)
            v8 := []int{32767, (fi6 ^ (fi6 | 9223372036854775807))}
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        p1 = true
    }
    p1 = false
    v9 := []int{127, 517, 349}
    var v10 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return v10
}
func genFn1() int32 {
    var v11 int64 = 0
    if (9223372036854775807 <= -792) {
        var v12 uint16 = 255
        v11 = (v11 + v11)
        _ = v12
    } else {
        var v13 int32 = 65
        _ = v13
    }
    v14 := []int{-536}
    _ = v11
    _ = v14
    return 147
}

func main() {
    v1 := []uint32{69, 525, 2147483647, 377}
    var v2 uint16 = 919
    v3 := make(map[string]uint32)
    var v15 string = "\n\t"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 float32 = 3.14159
    if (!true) {
        var v17 uint32 = 2147483647
        var v18 int32 = -2147483648
        _ = v17
        _ = v18
    } else {
        var v19 float64 = 1e10
        var v20 string = ""
        v2 = v2
        _ = v19
        _ = v20
    }
    v16 += 0.1
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: (^(fi21)), F1: -9223372036854775808}
        _ = fi21
        _ = v22
    }
    var v23 float64 = 1e10
    v24 := GenS1{F0: -9223372036854775808, F1: 503}
    if false {
        v23 += v23
        v15 = v15
        switch ((v15 + "hello") + "foo bar") {
        case "hello":
            v25 := GenS1{F0: 42, F1: -562}
            _ = v25
        case "foo bar":
            switch 297 {
            case 459:
                if true {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 uint = 198
                        switch v15 {
                        case "a":
                            v28 := GenS1{F0: 1, F1: 42}
                            _ = v28
                        case "foo bar":
                            for fi29 := 0; fi29 < 18; fi29++ {
                                var v30 uint32 = 0
                                _ = fi29
                                _ = v30
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                }
            case 2147483647:
                switch -186 {
                case -471:
                    _ = len(v1)
                default:
                    if false {
                        v23 = (v23 - 1.5)
                    } else {
                        var v31 float32 = (((((v16 - float32(v23)) - v16) / v16) * ((v16 + v16) * 1.5)) * v16)
                        v32 := GenS1{F0: -952, F1: 0}
                        v2 += (v2 / (((v2 & v2)) | 1))
                        _ = v31
                        _ = v32
                    }
                }
            case 716:
                if true {
                    delete(v3, v15)
                    if true {
                        if (255 >= 0) {
                            for fi33 := 0; fi33 < 19; fi33++ {
                                if true {
                                    var v34 float32 = (float32(3.14159) / (v16 * ((v16 * v16) / (v16 / v16))))
                                    v35 := GenS2{F0: 0, F1: 118, F2: 32767}
                                    delete(v3, (((v15 + (v15 + "foo bar")) + "\n\t") + ((v15 + "a") + (v15 + "\n\t"))))
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS2{F0: 374, F1: 1, F2: 168}
                                    _ = v36
                                }
                                fmt.Println(fi33, fi33, v15)
                                v37 := GenS3{F0: false, F1: 0}
                                _ = fi33
                                _ = v37
                            }
                        }
                    }
                    fmt.Println(v23, v15)
                }
            }
        default:
            var v38 int8 = -128
            _ = v38
        }
    }
    v39 := GenS1{F0: 128, F1: -9223372036854775808}
    fmt.Println(v16, v16)
    var v40 bool = ((((("\n\t" + (v15 + v15)) + "foo bar") + v15) + "\n\t") > ("" + v15))
    var v41 int = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
