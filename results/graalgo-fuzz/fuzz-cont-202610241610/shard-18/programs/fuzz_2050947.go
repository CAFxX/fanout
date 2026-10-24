package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int; F1 int32 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]string)
    var v3 float64 = 3.14159
    var v4 uint16 = (v1 / ((v1) | 1))
    v5 := GenS1{F0: 93, F1: 164, F2: -1}
    if (!(v1 != v4)) {
        switch ("foo bar" + "") {
        case "hello":
            v4 = v1
        case "":
            if (!false) {
                v3 -= (v3 - (float64(0.0) / v3))
                delete(v2, "a")
            } else {
                v4 = (^(v4))
            }
        case "\n\t":
            for fi6 := 0; fi6 < 16; fi6++ {
                switch (fi6 * -9223372036854775808) {
                case 9223372036854775807:
                    v1 += (v1 - 65535)
                case -1:
                    fmt.Println(fi6, fi6)
                default:
                    if (false && true) {
                        v1 = v4
                        var v7 uint64 = 2
                        var v8 uint64 = v7
                        _ = v7
                        _ = v8
                    }
                }
                v3 = 0.0
                _ = fi6
            }
        default:
            if (-119 == 117) {
                var v9 uint64 = 2
                v4 = v1
                v3 = v3
                _ = v9
            } else {
                var v10 float32 = float32((v3 - (v3 + v3)))
                var v11 int = 1
                _ = v10
                _ = v11
            }
        }
        var v12 int8 = -92
        if (386 == 0) {
            if (float32(1.5) >= 0.0) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    if (!(true || (true || true))) {
                        var v14 string = "hello"
                        if (127 > 247) {
                            v15 := GenS1{F0: 255, F1: 216, F2: v12}
                            v16 := GenS1{F0: 255, F1: 255, F2: -81}
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    } else {
                        switch (-(fi13)) {
                        case -1:
                            var v17 bool = true
                            _ = v17
                        case 127:
                            var v18 int32 = 2147483647
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 12; fi19++ {
                            if (true && true) {
                                var v20 uint64 = 0
                                _ = v20
                            } else {
                                v3 -= (((float64(1e10) * 0.0) * v3) / 1.0)
                                v21 := GenS2{F0: -1}
                                v3 -= (v3 / ((v3 + 3.14159) + (1.5 / 1.5)))
                                _ = v21
                            }
                            _ = fi19
                        }
                    }
                    _ = fi13
                }
                var v22 float32 = -2.25
                if (!((42 > 0) && (!true))) {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        if true {
                            if false {
                                var v24 bool = (-1 == -203)
                                var v25 uint16 = v1
                                v3 = v3
                                _ = v24
                                _ = v25
                            } else {
                                fmt.Println(v1, v22, v3)
                                var v26 int = 127
                                _ = v26
                            }
                            v22 = v22
                            v27 := GenS3{F0: 0, F1: -32768, F2: false}
                            _ = v27
                        }
                        if (!false) {
                            switch "hello" {
                            case "\n\t":
                                v28 := []int{1, fi23}
                                _ = v28
                            case "":
                                var v29 uint = 897
                                _ = v29
                            case "foo bar":
                                v3 += (float64(float32(((v3 / float64(v22)) + float64(3.14159)))) * v3)
                            }
                        }
                        _ = fi23
                    }
                }
                _ = v22
            }
        } else {
            var v30 int64 = -354
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 uint16 = 255
                switch 32767 {
                case 127:
                    fmt.Println(v1)
                default:
                    fmt.Println(v32)
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    if true {
                        v34 := GenS4{F0: 42}
                        if ((!true) || true) {
                            var v35 int = (fi33 | (fi31 | fi31))
                            _ = v35
                        } else {
                            v1 = (v32 - (v32 / ((v32) | 1)))
                        }
                        _ = v34
                    }
                    v36 := []int{(fi33 % ((-9223372036854775808) | 1)), (fi33 - int(v12)), (fi33 >> ((42) & 63)), 128}
                    _ = fi33
                    _ = v36
                }
                _ = fi31
                _ = v32
            }
            v37 := []int{55, 4294967295, -934, 739}
            _ = v30
            _ = v37
        }
        _ = v12
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v39 := []int{(-(fi38))}
        var v40 float64 = (((-2.25 * v3) / float64(1e10)) / 1.5)
        var v41 int8 = 127
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 uint8 = 2
    v43 := GenS5{F0: 582, F1: 1}
    switch 380 {
    case 649:
        v3 = v3
    }
    delete(v2, ("\n\t" + ""))
    fmt.Println(v4, v1, v4)
    if (((0 != 2147483648) && true) && true) {
        for fi44 := 0; fi44 < 13; fi44++ {
            v45 := GenS6{F0: (fi44 / (((^(fi44))) | 1)), F1: -1}
            switch "" {
            case "\n\t":
                var v46 int8 = 127
                _ = v46
            default:
                switch "a" {
                case "a":
                    v47 := GenS1{F0: 127, F1: v42, F2: 1}
                    _ = v47
                }
            }
            delete(v2, ("\n\t" + "foo bar"))
            _ = fi44
            _ = v45
        }
        for fi48 := 0; fi48 < 5; fi48++ {
            if false {
                v4 = (v4 % ((1) | 1))
            }
            _ = fi48
        }
        switch "" {
        case "foo bar":
            v49 := []int{127, 32768, 9223372036854775807}
            _ = v49
        case "a":
            for fi50 := 0; fi50 < 1; fi50++ {
                delete(v2, (("a" + ("hello" + "")) + ""))
                for fi51 := 0; fi51 < 2; fi51++ {
                    var v52 int = fi50
                    _ = fi51
                    _ = v52
                }
                _ = fi50
            }
        }
    }
    if true {
        v42 = v42
        var v53 int8 = 0
        _ = v53
    }
    fmt.Println(v1, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
