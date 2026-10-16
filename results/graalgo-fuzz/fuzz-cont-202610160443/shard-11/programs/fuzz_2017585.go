package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []int{128, 635}
    var v2 uint = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483647
    var v5 int32 = 2147483647
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(v2, v4, v4)
        switch -1 {
        case 32767:
            v7 := GenS1{F0: (^(v2)), F1: -128}
            _ = v7
        case -350:
            if (fi6 <= fi6) {
                if true {
                    var v8 float32 = 1e10
                    _ = v8
                }
                var v9 int8 = 0
                _ = v9
            } else {
                var v10 float64 = (0.1 / (-2.25 - 0.0))
                if (1 < 1000) {
                    switch "a" {
                    case "hello":
                        for fi11 := 0; fi11 < 2; fi11++ {
                            var v12 int16 = 0
                            fmt.Println(fi11, v12)
                            switch "foo bar" {
                            case "":
                                v2 += v2
                            case "a":
                                v4 = v4
                            }
                            _ = fi11
                            _ = v12
                        }
                    default:
                        for fi13 := 0; fi13 < 10; fi13++ {
                            var v14 int = fi6
                            _ = fi13
                            _ = v14
                        }
                    }
                    if (true && false) {
                        if (v10 >= v10) {
                            fmt.Println(v5)
                            var v15 uint32 = (v4 | (v4 * v4))
                            _ = v15
                        }
                    } else {
                        var v16 uint32 = 0
                        _ = v16
                    }
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v17 int64 = 0
                        _ = v17
                    case "a":
                        _ = len(v1)
                    default:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 int = fi6
                            _ = fi18
                            _ = v19
                        }
                    }
                    var v20 uint = (v2 % (((v2 | 726)) | 1))
                    _ = v20
                }
                _ = v10
            }
        case -728:
            var v21 uint64 = 255
            _ = v21
        }
        if false {
            v2 = v2
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 uint16 = uint16(v2)
                _ = fi22
                _ = v23
            }
        }
        _ = fi6
    }
    v4 = v4
    var v24 float64 = (3.14159 / (((1.5 / 1.0) + float64(3.14159)) * 1.5))
    v25 := GenS1{F0: 890, F1: -1}
    for fi26 := 0; fi26 < 13; fi26++ {
        v4 += 850
        var v27 float64 = 3.14159
        _ = fi26
        _ = v27
    }
    fmt.Println(v2, v24)
    var v28 string = "hello"
    v24 -= v24
    v24 = v24
    for fi29 := 0; fi29 < 14; fi29++ {
        fmt.Println(fi29)
        _ = fi29
    }
    v30 := GenS1{F0: v2, F1: -33}
    switch ("foo bar" + v28) {
    case "a":
        var v31 int64 = -1
        _ = v31
    case "\n\t":
        var v32 uint16 = 32
        _ = v32
    case "foo bar":
        var v33 uint = 2
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS2{F0: (v5 | v5), F1: 65535, F2: v5}
        _ = len(v3)
        if len(v1) > 0 {
            v1[0] = (fi34 % ((int(v4)) | 1))
        }
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        fmt.Println(v5, v24)
        var v37 float64 = (v24 + -2.25)
        v38 := GenS3{F0: 237}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
