package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (-215 < 0)
    var v3 float32 = float32(-2.25)
    var v4 string = "foo bar"
    v5 := make(map[int]int)
    switch v4 {
    case "foo bar":
        switch -1 {
        case 42:
            v4 = v4
        case 128:
            v5[42] = -9223372036854775808
        }
    case "a":
        v4 = v4
    default:
        fmt.Println(v1, v4)
    }
    var v6 string = ""
    v7 := GenS1{F0: 2}
    for fi8 := 0; fi8 < 14; fi8++ {
        if (v2 && (60 != 0)) {
            switch (("a" + (v6 + (("hello" + v6) + "a"))) + (((((v4 + ("a" + v4)) + "") + "foo bar") + (v4 + v4)) + v6)) {
            case "\n\t":
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v1)
                    fmt.Println(v1, v4, v3)
                    switch v6 {
                    case "\n\t":
                        if (!v2) {
                            v10 := GenS2{F0: 2}
                            var v11 int32 = 2147483647
                            v12 := []int{fi8, (fi8 / (((fi8 - int(v11))) | 1)), 1}
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int8 = 42
                            v14 := GenS2{F0: 0}
                            if (-1 < 1) {
                                v2 = v2
                            } else {
                                fmt.Println(fi8, fi9, v3)
                                v6 = "hello"
                                var v15 uint32 = 65535
                                _ = v15
                            }
                            v16 := GenS2{F0: 140}
                            fmt.Println(v1)
                            _ = v13
                            _ = v14
                            _ = v16
                        }
                    case "":
                        switch (-(fi9)) {
                        case -1:
                            switch (v4 + "\n\t") {
                            case "hello":
                                v17 := GenS3{F0: 611, F1: (2147483648 >= 0), F2: 0}
                                _ = v17
                            default:
                                if false {
                                    v18 := GenS1{F0: 255}
                                    var v19 uint32 = 2147483648
                                    _ = v18
                                    _ = v19
                                } else {
                                    switch ((("a" + v6) + v6) + (v4 + ("hello" + "hello"))) {
                                    case "a":
                                        v1 = 3.14159
                                    case "":
                                        var v20 int16 = -32768
                                        _ = v20
                                    case "foo bar":
                                        var v21 uint32 = 352
                                        _ = v21
                                    }
                                    if v2 {
                                        v6 = (("a" + (("hello" + (((v4 + "\n\t") + "") + ("foo bar" + v4))) + "\n\t")) + v6)
                                        v22 := v5[fi9]
                                        v3 = 0.0
                                        _ = v22
                                    } else {
                                        v5[fi8] = (fi9 % ((fi8) | 1))
                                    }
                                    fmt.Println(v1, fi9)
                                    v2 = (-2147483648 <= 575)
                                }
                            }
                        case -372:
                            v23 := GenS3{F0: 42, F1: (182 == 42), F2: 2}
                            _ = v23
                        default:
                            switch fi8 {
                            case 32767:
                                v24 := v5[fi9]
                                _ = v24
                            }
                        }
                    case "hello":
                        v3 = v3
                    default:
                        v25 := GenS2{F0: 42}
                        _ = v25
                    }
                    _ = fi9
                }
            }
        } else {
            fmt.Println(fi8, v2, v3)
            v2 = v2
            v1 = ((v1 / (3.14159 / (v1 - v1))) - v1)
        }
        if false {
            v26 := []int{}
            for fi27 := 0; fi27 < 14; fi27++ {
                fmt.Println(fi27, v6, fi27)
                if (2147483647 != (fi8 * 32768)) {
                    if (-1 != 42) {
                        switch fi27 {
                        case 42:
                            v2 = (65535 > 940)
                        case 32767:
                            v2 = (((-90 == 9) && (112 <= 127)) || (42 <= 1))
                        }
                        v3 = v3
                    } else {
                        switch v4 {
                        case "\n\t":
                            delete(v5, fi8)
                        }
                        v2 = ((fi8 & 32768) > fi27)
                        if v2 {
                            v6 = (v6 + v4)
                        } else {
                            switch fi27 {
                            case -1:
                                fmt.Println(v2, v6)
                            default:
                                var v28 float32 = (v3 - v3)
                                _ = v28
                            }
                        }
                    }
                }
                var v29 bool = (!(true || ("" >= (((((v4 + "") + "a") + v4) + v4) + "\n\t"))))
                _ = fi27
                _ = v29
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                if v2 {
                    v31 := GenS1{F0: 42}
                    _ = v31
                } else {
                    fmt.Println(v3, fi30)
                    v1 -= float64(v3)
                }
                fmt.Println(fi30, v3, fi30)
                _ = fi30
            }
            _ = v26
        } else {
            v2 = v2
            _ = len(v5)
        }
        _ = fi8
    }
    var v32 int16 = -51
    var v33 uint32 = 2147483648
    v34 := []int{-740, -2147483648}
    var v35 int32 = -1
    var v36 string = "foo bar"
    if true {
        v36 = (v36 + v36)
    }
    v37 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
