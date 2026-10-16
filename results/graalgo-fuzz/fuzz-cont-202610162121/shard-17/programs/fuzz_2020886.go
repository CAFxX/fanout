package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint8 = 159
    var v2 int8 = 42
    var v3 float64 = 0.1
    v4 := make(map[int]int)
    v5 := []string{}
    if (22 < 255) {
        v6 := GenS1{F0: (957 >= -1), F1: (v2 * (v2 & 44))}
        _ = v6
    } else {
        v7 := GenS1{F0: false, F1: 42}
        _ = v7
    }
    v2 = v2
    v1 += (v1 << ((0) & 63))
    v8 := GenS1{F0: (((!true) && (true && ((v1 - 42) == 127))) || (false || false)), F1: -66}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int32 = 862
        if (!(!(127 == 32))) {
            if false {
                var v11 float32 = 1e10
                fmt.Println(v11, v11)
                fmt.Println(v2)
                _ = v11
            } else {
                v3 = 0.1
                v3 -= v3
                v2 += 0
            }
            v12 := GenS1{F0: false, F1: v2}
            v5 = append(v5, "foo bar")
            _ = v12
        }
        if (42 >= 2147483647) {
            if false {
                var v13 bool = true
                var v14 uint8 = (v1 ^ (v1 >> ((4294967295) & 63)))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v13
                _ = v14
            }
            fmt.Println(fi9, v1, fi9)
        } else {
            var v15 string = "foo bar"
            var v16 bool = (9223372036854775807 != 519)
            v10 -= v10
            _ = v15
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    switch 42 {
    case 32768:
        if true {
            switch "foo bar" {
            case "hello":
                switch "\n\t" {
                case "hello":
                    if false {
                        v3 -= (v3 + v3)
                        if false {
                            v2 -= (v2 ^ v2)
                            if (0 >= 2) {
                                var v17 float32 = 1e10
                                if false {
                                    if false {
                                        var v18 uint = 1
                                        _ = v18
                                    }
                                    if true {
                                        if true {
                                            v1 = (v1 ^ (^(v1)))
                                        }
                                        var v19 int32 = 470
                                        fmt.Println(v1)
                                        _ = v19
                                    }
                                    var v20 int8 = (v2 & -128)
                                    _ = v20
                                } else {
                                    v21 := GenS1{F0: ((true && false) && false), F1: v2}
                                    var v22 int32 = 0
                                    _ = v21
                                    _ = v22
                                }
                                _ = v17
                            }
                        } else {
                            fmt.Println(v2, v2, v1)
                            if (32767 == -32768) {
                                v23 := GenS1{F0: true, F1: 127}
                                fmt.Println(v1, v2)
                                _ = v23
                            }
                        }
                        for fi24 := 0; fi24 < 18; fi24++ {
                            var v25 int64 = 0
                            if (!true) {
                                v3 -= (((0.1 + (float64(3.14159) * (0.0 / v3))) - 0.0) * v3)
                            }
                            switch "foo bar" {
                            case "":
                                v1 = 25
                            case "hello":
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    v27 := GenS1{F0: ((float64(-2.25) + 0.0) > float64(0.1)), F1: 127}
                                    var v28 int8 = 64
                                    v29 := GenS1{F0: true, F1: (-(v2))}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                case "foo bar":
                    fmt.Println(v1, v3)
                }
            case "foo bar":
                if false {
                    if true {
                        if false {
                            switch "\n\t" {
                            case "":
                                v30 := GenS1{F0: false, F1: 127}
                                _ = v30
                            case "\n\t":
                                if (false || (42 < -1)) {
                                    var v31 uint32 = 275
                                    _ = v31
                                }
                            }
                        } else {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                if len(v5) > 0 {
                                    v5[0] = "hello"
                                }
                                _ = fi32
                            }
                            v33 := GenS2{F0: ""}
                            fmt.Println(v3)
                            _ = v33
                        }
                    }
                    v4[int(v1)] = 4294967295
                    v34 := GenS1{F0: true, F1: v2}
                    _ = v34
                }
            case "\n\t":
                v35 := GenS1{F0: false, F1: (v2 - v2)}
                _ = v35
            }
            if false {
                fmt.Println(v2)
                v36 := GenS3{F0: 266, F1: 32767}
                _ = v36
            }
            switch "foo bar" {
            case "\n\t":
                for fi37 := 0; fi37 < 11; fi37++ {
                    v1 -= 2
                    v38 := v4[(^(fi37))]
                    _ = fi37
                    _ = v38
                }
            case "a":
                _ = len(v4)
            }
        } else {
            v39 := GenS4{F0: 32767, F1: 1}
            _ = v39
        }
    case 4294967295:
        for fi40 := 0; fi40 < 18; fi40++ {
            v5 = append(v5, "")
            var v41 bool = false
            _ = fi40
            _ = v41
        }
    case 127:
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    default:
        if len(v5) > 0 {
            v5[0] = ("hello" + "hello")
        }
    }
    fmt.Println(v1, v2, v1)
    var v42 int64 = -864
    var v43 int = 9223372036854775807
    v3 = 1.5
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := v4[4294967295]
        _ = fi44
        _ = v45
    }
    v42 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
