package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }

func main() {
    v1 := []int{4294967295, -240, -590}
    var v2 float32 = float32(0.0)
    var v3 int8 = 100
    v2 = v2
    switch "\n\t" {
    case "hello":
        v4 := GenS1{F0: 141, F1: -1, F2: true}
        _ = v4
    case "\n\t":
        if true {
            var v5 uint8 = 255
            _ = v5
        } else {
            if ((true || (false && ((v3 % ((-128) | 1)) <= (v3 * -128)))) && false) {
                v1 = append(v1, 81)
            } else {
                if (!(2 <= 255)) {
                    for fi6 := 0; fi6 < 19; fi6++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "a" {
                        case "a":
                            for fi7 := 0; fi7 < 10; fi7++ {
                                switch fi7 {
                                case -399:
                                    fmt.Println(v2, fi6, v2)
                                case 4294967295:
                                    var v8 int8 = (v3 % (((v3 - 109)) | 1))
                                    _ = v8
                                case 1:
                                    v2 += -2.25
                                }
                                _ = fi7
                            }
                        }
                        _ = fi6
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 -= float32(-2.25)
            }
            if ((true || (255 < 1)) || (10 <= 145)) {
                v9 := GenS1{F0: 135, F1: 42, F2: (2 < 69)}
                _ = v9
            } else {
                _ = len(v1)
                var v10 uint64 = 42
                _ = v10
            }
        }
    case "a":
        var v11 int8 = 0
        _ = v11
    default:
        if (-2.25 <= v2) {
            v3 = 1
            var v12 int32 = 2147483647
            if true {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v2 = v2
                    _ = fi13
                }
                if (!(285 == 1)) {
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                }
            }
            _ = v12
        }
    }
    fmt.Println(v3, v2)
    fmt.Println(v3)
    for fi14 := 0; fi14 < 7; fi14++ {
        if (v3 == v3) {
            v15 := GenS1{F0: 2, F1: 1, F2: true}
            switch "" {
            case "hello":
                v16 := GenS2{F0: 255, F1: 0, F2: 4294967295}
                _ = v16
            case "foo bar":
                fmt.Println(v3)
            }
            _ = v15
        }
        var v17 int8 = (v3 ^ (v3 + (v3 / ((0) | 1))))
        _ = fi14
        _ = v17
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    if true {
        var v18 uint32 = 42
        v1 = append(v1, 445)
        _ = v18
    }
    _ = len(v1)
    v3 -= (v3 | 100)
    if true {
        if false {
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v2)
                v2 -= 3.14159
                _ = fi19
            }
        }
    }
    switch "foo bar" {
    case "hello":
        var v20 uint16 = 1
        _ = v20
    case "a":
        var v21 int32 = 212
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        _ = len(v1)
        v2 = 1e10
        v3 = v3
        _ = fi22
    }
    if true {
        _ = len(v1)
        v3 += v3
        v3 -= -1
    } else {
        v2 = v2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = 42
    v3 = (v3 & (v3 >> ((0) & 63)))
    if (330 > 1) {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int = (fi23 >> ((4294967295) & 63))
            v2 = (v2 / float32(float64(v2)))
            var v25 string = ("" + "a")
            _ = fi23
            _ = v24
            _ = v25
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v26 float64 = (0.1 / (((0.1 * 1.5) + 0.1) + 1e10))
        switch -409 {
        case 127:
            var v27 int = 32768
            _ = v27
        }
        var v28 uint64 = 0
        _ = v26
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        fmt.Println(v2, v2)
        v30 := GenS1{F0: 255, F1: 32767, F2: (1.5 == float64(((float32(0.1) + float32(0.1)) + (float32((1e10 - float64(v2))) - 1e10))))}
        _ = fi29
        _ = v30
    }
    v3 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
