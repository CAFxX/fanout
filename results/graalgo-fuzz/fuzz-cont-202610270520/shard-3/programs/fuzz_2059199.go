package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    var v5 float64 = 3.14159
    fmt.Println(v2, v2, v4)
    v2 = v2
    v6 := GenS1{F0: 2147483647}
    v5 = 3.14159
    var v7 int8 = 0
    switch "foo bar" {
    case "\n\t":
        v2 += 0.0
    }
    switch -2147483648 {
    case 1:
        v8 := GenS1{F0: 1}
        _ = v8
    case 32768:
        if (!false) {
            fmt.Println(v2, v5, v4)
        } else {
            v3[0] = true
            var v9 int8 = v7
            _ = v9
        }
    case 127:
        var v10 uint = 18446744073709551615
        _ = v10
    }
    switch 128 {
    case 32768:
        if (false && (591 >= -38)) {
            switch "" {
            case "a":
                v7 = (^(v7))
            default:
                var v11 uint32 = 0
                _ = v11
            }
            v5 = float64(float32((((v5 - (v5 * v5)) / (3.14159 - -2.25)) + -2.25)))
            v12 := v3[-64]
            _ = v12
        }
    case -757:
        var v13 int8 = v7
        _ = v13
    case 32767:
        v7 -= -128
    default:
        var v14 uint16 = 957
        _ = v14
    }
    var v15 int32 = 2147483647
    fmt.Println(v2, v15)
    switch -1 {
    case 42:
        v5 += float64((v2 * 0.1))
    case 862:
        switch 32768 {
        case 9223372036854775807:
            var v16 int32 = v15
            _ = v16
        case -1:
            if true {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    if false {
                        v2 += v2
                        var v17 string = (("a" + "a") + "a")
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 uint16 = 354
                            var v20 bool = true
                            for fi21 := 0; fi21 < 19; fi21++ {
                                var v22 uint32 = 42
                                v23 := GenS2{F0: 9223372036854775807, F1: v20, F2: true}
                                v24 := v3[int(v22)]
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    } else {
                        var v25 int = 32768
                        _ = v25
                    }
                }
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            } else {
                v15 -= v15
                if ((((v7 & 127) != (^(v7))) && false) && (v4 != (v4 + v4))) {
                    var v26 uint = 88
                    var v27 bool = false
                    _ = v26
                    _ = v27
                }
                switch int(v15) {
                case 9223372036854775807:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 42:
                    var v28 float32 = (v2 / v2)
                    _ = v28
                }
            }
        case 430:
            if (32767 != 0) {
                if true {
                    if (-1 > 32767) {
                        var v29 int = 32768
                        _ = v29
                    }
                    fmt.Println(v5)
                    fmt.Println(v4, v7)
                } else {
                    v1 = append(v1, "hello")
                    var v30 int64 = 379
                    _ = v30
                }
                var v31 uint16 = 255
                _ = v31
            }
        default:
            v1 = append(v1, ("" + "a"))
        }
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println(v2, v5, v15)
    var v32 uint = 0
    v15 = -414
    v7 = v7
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := GenS2{F0: 141, F1: false, F2: true}
        var v35 int = -1
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
