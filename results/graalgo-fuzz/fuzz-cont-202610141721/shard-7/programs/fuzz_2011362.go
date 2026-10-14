package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := []int{-2147483648, -316, 32767, 42}
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint = 2
        _ = fi4
        _ = v5
    }
    _ = len(v3)
    var v6 bool = false
    if true {
        var v7 int = 2147483647
        _ = v7
    }
    v8 := GenS1{F0: 0}
    _ = len(v1)
    if (true || false) {
        delete(v2, "\n\t")
        for fi9 := 0; fi9 < 17; fi9++ {
            if (545 <= 255) {
                var v10 uint16 = 723
                _ = v10
            } else {
                var v11 float32 = float32(1e10)
                if v6 {
                    var v12 bool = (fi9 >= 4294967295)
                    if ((("hello" + "\n\t") != "hello") || v12) {
                        var v13 bool = (164 >= 255)
                        var v14 string = "hello"
                        v15 := v1[fi9]
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint16 = 2
                        if (18446744073709551615 < 2) {
                            if ((("a" + "hello") + "\n\t") < "foo bar") {
                                var v17 uint16 = (v16 ^ (^(v16)))
                                _ = v17
                            } else {
                                v16 = v16
                                if ((fi9 | (fi9 - 1)) <= (fi9 ^ (fi9 + (fi9 ^ (fi9 << ((2) & 63)))))) {
                                    for fi18 := 0; fi18 < 19; fi18++ {
                                        var v19 int8 = 0
                                        if (!(((-418 > 534) || (true || true)) && (v6 && (44 >= 1)))) {
                                            if v6 {
                                                v20 := v2["\n\t"]
                                                _ = v20
                                            }
                                            var v21 float32 = (v11 + v11)
                                            var v22 uint32 = 1
                                            _ = v21
                                            _ = v22
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (-(fi18))
                                        }
                                        var v23 string = "foo bar"
                                        var v24 int32 = 981
                                        _ = fi18
                                        _ = v19
                                        _ = v23
                                        _ = v24
                                    }
                                    switch ("a" + "a") {
                                    case "\n\t":
                                        v12 = v6
                                    case "foo bar":
                                        v25 := GenS1{F0: int32(v16)}
                                        _ = v25
                                    case "hello":
                                        for fi26 := 0; fi26 < 18; fi26++ {
                                            if (!((!(true && (true || (-52 >= 0)))) && ((!v6) && true))) {
                                                var v27 uint32 = 4294967295
                                                _ = v27
                                            } else {
                                                switch 32767 {
                                                case -842:
                                                    v28 := GenS2{F0: 2147483648, F1: (fi9 * -9223372036854775808), F2: -32768}
                                                    _ = v28
                                                }
                                                if true {
                                                    if v12 {
                                                        var v29 int32 = 0
                                                        _ = v29
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch -219 {
                                                            case -72:
                                                                v30 := GenS1{F0: 2147483647}
                                                                _ = v30
                                                            case 32767:
                                                                var v31 string = ("a" + ("foo bar" + "hello"))
                                                                _ = v31
                                                            }
                                                        case "hello":
                                                            var v32 int64 = 695
                                                            _ = v32
                                                        }
                                                        v33 := GenS2{F0: 4294967295, F1: (fi9 + 32767), F2: -1}
                                                        _ = v33
                                                    }
                                                    var v34 int8 = int8(v16)
                                                    fmt.Println(v34, v16, fi9)
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: -674, F1: -2147483648}
                                                _ = v35
                                            }
                                            v3 = append(v3, (fi9 / ((fi26) | 1)))
                                            _ = fi26
                                        }
                                    default:
                                        fmt.Println(v11, v11, v16)
                                    }
                                    switch ("hello" + ("foo bar" + ((("foo bar" + "foo bar") + "\n\t") + "\n\t"))) {
                                    case "foo bar":
                                        v11 = (v11 * v11)
                                    }
                                }
                                fmt.Println(v6, v11, v11)
                            }
                            fmt.Println(v6)
                        } else {
                            v36 := GenS4{F0: 0, F1: 0}
                            _ = v36
                        }
                        var v37 float32 = 0.0
                        _ = v16
                        _ = v37
                    }
                    switch 32767 {
                    case -1:
                        for fi38 := 0; fi38 < 14; fi38++ {
                            switch 127 {
                            case 127:
                                if (-1 <= -1) {
                                    fmt.Println(v6, v12)
                                    var v39 bool = v12
                                    _ = v39
                                }
                            case 32768:
                                fmt.Println(fi9)
                            case 4294967295:
                                var v40 bool = true
                                _ = v40
                            }
                            v11 += (v11 / ((float32(1e10) / v11) - (v11 * (v11 + v11))))
                            v11 = -2.25
                            _ = fi38
                        }
                    case -785:
                        switch ("\n\t" + ("\n\t" + "hello")) {
                        case "\n\t":
                            v41 := GenS2{F0: 65535, F1: (fi9 % (((-(fi9))) | 1)), F2: 1000}
                            _ = v41
                        case "":
                            var v42 uint16 = 203
                            _ = v42
                        }
                    case 127:
                        delete(v1, -9223372036854775808)
                    default:
                        v43 := GenS4{F0: 255, F1: 100}
                        _ = v43
                    }
                    _ = v12
                } else {
                    v6 = v6
                    v6 = v6
                }
                _ = v11
            }
            _ = fi9
        }
    } else {
        if v6 {
            v6 = (255 <= 0)
            if v6 {
                if (-2.25 < ((0.0 + (float32(1.5) / (1.5 / 0.1))) * (float32(1.5) / 0.1))) {
                    v6 = v6
                    v6 = true
                    if true {
                        v44 := GenS2{F0: 0, F1: 127, F2: 190}
                        _ = v44
                    }
                } else {
                    v45 := v2["\n\t"]
                    var v46 int = 9223372036854775807
                    _ = v45
                    _ = v46
                }
            }
            var v47 string = "a"
            _ = v47
        }
        v6 = (3.14159 < 1.5)
    }
    fmt.Println(v6)
    if v6 {
        fmt.Println(v6)
        var v48 int16 = 0
        for fi49 := 0; fi49 < 5; fi49++ {
            v48 = v48
            _ = fi49
        }
        _ = v48
    }
    var v50 uint = 2
    v51 := GenS5{F0: "a", F1: -358}
    var v52 float64 = -2.25
    var v53 float64 = v52
    var v54 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v8
}
