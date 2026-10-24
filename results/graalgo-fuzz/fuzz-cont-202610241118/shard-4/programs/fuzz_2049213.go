package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint16 }
func genFn0() uint8 {
    var v5 uint16 = 32767
    switch "hello" {
    case "hello":
        v5 = 746
    case "a":
        switch ("hello" + "\n\t") {
        case "a":
            var v6 int8 = 100
            _ = v6
        case "hello":
            var v7 int32 = -2147483648
            _ = v7
        default:
            if true {
                v5 = (v5 % ((255) | 1))
                for fi8 := 0; fi8 < 13; fi8++ {
                    fmt.Println(v5)
                    fmt.Println(fi8, v5)
                    _ = fi8
                }
                v5 = (v5 * 1)
            }
        }
    }
    var v9 int64 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int64 = v9
        var v12 uint = 4294967295
        _ = fi10
        _ = v11
        _ = v12
    }
    v5 += v5
    _ = v5
    _ = v9
    return 128
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = false
    var v3 string = ""
    var v4 uint8 = 127
    v13 := []int{0}
    fmt.Println(v4)
    v14 := GenS1{F0: (v4 / ((2) | 1)), F1: 127}
    if v2 {
        var v15 int = -2147483648
        if false {
            fmt.Println(v4)
        }
        v16 := GenS2{F0: 32767, F1: 2}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v1, v2)
        _ = len(v13)
    }
    switch -2147483648 {
    case -1:
        if (42 == 174) {
            switch 9223372036854775807 {
            case -324:
                v17 := GenS3{F0: -32768}
                _ = v17
            default:
                var v18 int16 = 42
                _ = v18
            }
        }
    case 4294967295:
        var v19 int16 = 0
        _ = v19
    case -423:
        fmt.Println(v1)
    default:
        var v20 int8 = 100
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v1 += v1
        for fi22 := 0; fi22 < 4; fi22++ {
            if (2147483648 == 497) {
                v1 -= v1
                if (!v2) {
                    v4 -= v4
                } else {
                    var v23 uint = 0
                    v24 := GenS2{F0: 456, F1: 2}
                    v25 := GenS4{F0: -2147483648}
                    switch (v3 + "\n\t") {
                    case "foo bar":
                        var v26 float32 = (1e10 + float32(float64(0.0)))
                        _ = v26
                    }
                    v27 := GenS2{F0: 175, F1: 0}
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v27
                }
            } else {
                if len(v13) > 0 {
                    v13[0] = fi21
                }
                switch v3 {
                case "a":
                    if (!(!((2147483647 <= 269) && true))) {
                        v2 = (709 == -892)
                    } else {
                        if (((v3 + (v3 + ("a" + v3))) + "foo bar") == "\n\t") {
                            var v28 uint = 255
                            var v29 int16 = -1
                            _ = v28
                            _ = v29
                        }
                    }
                case "foo bar":
                    var v30 string = v3
                    _ = v30
                default:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                v2 = v2
            }
            v31 := GenS5{F0: 42, F1: 266, F2: 65535}
            _ = fi22
            _ = v31
        }
        v13 = append(v13, fi21)
        _ = fi21
    }
    fmt.Println(v4, v3)
    v4 = (v4 + genFn0())
    var v32 uint = 0
    v4 += 1
    var v33 string = "a"
    if true {
        var v34 int = -1
        var v35 int8 = -6
        _ = v34
        _ = v35
    } else {
        fmt.Println(v2, v1, v33)
        v13 = append(v13, int(v32))
        v36 := GenS1{F0: v4, F1: 424}
        _ = v36
    }
    var v37 string = v3
    fmt.Println(v3, v3, v3)
    v13 = append(v13, 9223372036854775807)
    v38 := GenS1{F0: 42, F1: -1}
    var v39 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
