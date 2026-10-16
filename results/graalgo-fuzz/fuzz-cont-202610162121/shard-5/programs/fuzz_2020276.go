package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    var v2 int16 = 1000
    v3 := []string{("a" + ("" + "")), ""}
    fmt.Println(v2)
    if ((255 <= 255) || true) {
        fmt.Println(v2)
    }
    _ = len(v3)
    var v4 uint32 = 42
    fmt.Println(v2)
    var v5 uint8 = 42
    fmt.Println(v4, v5)
    v4 -= (^(v4))
    fmt.Println(v2)
    switch "\n\t" {
    case "\n\t":
        var v6 uint16 = 2
        _ = v6
    }
    if false {
        if (283 > 4294967295) {
            var v7 uint64 = 0
            _ = v7
        } else {
            v1 = append(v1, (v4 ^ (^(v4))))
            if false {
                if true {
                    v4 = 2147483647
                    var v8 float64 = 1e10
                    _ = v8
                }
                var v9 float32 = 0.0
                v10 := GenS1{F0: 2}
                _ = v10
                _ = v9
            } else {
                var v11 uint64 = 18446744073709551615
                v11 += (v11 | (v11 & v11))
                _ = v11
            }
        }
        if (42 <= -733) {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 127}
                var v14 int64 = -1
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                var v16 int16 = -289
                fmt.Println(v5)
                _ = v16
            } else {
                switch (fi15 & int(v2)) {
                case 811:
                    fmt.Println(fi15, v4, v5)
                default:
                    var v17 int16 = (v2 & (v2 & (-(v2))))
                    _ = v17
                }
                if true {
                    switch "a" {
                    case "foo bar":
                        if false {
                            var v18 uint8 = 42
                            _ = v18
                        } else {
                            v19 := GenS1{F0: 127}
                            _ = v19
                        }
                    case "\n\t":
                        v5 += v5
                    case "hello":
                        if ("foo bar" == ("foo bar" + "hello")) {
                            var v20 int = (fi15 ^ fi15)
                            _ = v20
                        } else {
                            _ = len(v3)
                        }
                    default:
                        v2 += v2
                    }
                    var v21 int64 = 9223372036854775807
                    _ = v21
                } else {
                    v4 -= v4
                }
            }
            if true {
                for fi22 := 0; fi22 < 9; fi22++ {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v23 := GenS2{F0: 651}
                    v4 -= v4
                    _ = fi22
                    _ = v23
                }
            } else {
                v24 := GenS1{F0: 1}
                fmt.Println(fi15, v5, v4)
                _ = v24
            }
            var v25 int32 = -2147483648
            _ = fi15
            _ = v25
        }
    } else {
        var v26 int16 = 42
        if true {
            if true {
                v4 -= (v4 | 2)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v4 -= (^(v4))
        } else {
            var v27 int32 = -2147483648
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            v29 := GenS2{F0: 0}
            var v30 string = "\n\t"
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v26
    }
    var v31 float32 = -2.25
    _ = len(v1)
    if (false || (-2147483648 != 1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            if (((("hello" + "foo bar") + "foo bar") + (("foo bar" + ("\n\t" + "a")) + "foo bar")) != "foo bar") {
                v33 := GenS2{F0: 42}
                switch fi32 {
                case 4294967295:
                    v2 -= v2
                case -963:
                    v31 = (-2.25 - (v31 / v31))
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                switch "hello" {
                case "hello":
                    v3 = append(v3, "\n\t")
                }
                _ = v33
            } else {
                v31 += (v31 / float32((1.5 / 3.14159)))
                fmt.Println(fi32)
            }
            v1 = append(v1, v4)
            for fi34 := 0; fi34 < 13; fi34++ {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    fmt.Println(v31, v5)
                case "":
                    if (9223372036854775807 >= 887) {
                        fmt.Println(v31, v5)
                    } else {
                        var v35 int32 = 1
                        v5 += v5
                        _ = v35
                    }
                case "hello":
                    fmt.Println(fi34)
                default:
                    if (!true) {
                        v1 = append(v1, 2147483647)
                        if (1 != -9223372036854775808) {
                            v3 = append(v3, (("\n\t" + "") + "hello"))
                            if ((!false) && false) {
                                var v36 int32 = -844
                                var v37 uint = 1
                                if false {
                                    fmt.Println(v5, fi32, v5)
                                }
                                if ("foo bar" <= (("hello" + "\n\t") + ("a" + "a"))) {
                                    _ = len(v1)
                                    v38 := GenS2{F0: 1}
                                    var v39 uint = (v37 * v37)
                                    _ = v38
                                    _ = v39
                                }
                                var v40 uint8 = uint8(v37)
                                var v41 float32 = 0.1
                                if true {
                                    if (v5 <= 127) {
                                        v37 += v37
                                        var v42 int64 = 1
                                        _ = v42
                                    } else {
                                        v43 := GenS2{F0: 18446744073709551615}
                                        var v44 int64 = -1
                                        v45 := GenS1{F0: 42}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    v40 -= v40
                                    switch "a" {
                                    case "hello":
                                        switch ("hello" + ((("hello" + "foo bar") + ("a" + ("foo bar" + "hello"))) + "\n\t")) {
                                        case "a":
                                            fmt.Println(v41, v41)
                                        case "":
                                            if false {
                                                v40 = v40
                                            }
                                        case "hello":
                                            v4 = (^(v4))
                                        default:
                                            v40 = (v40 * uint8(v37))
                                        }
                                    default:
                                        var v46 int64 = 9223372036854775807
                                        _ = v46
                                    }
                                } else {
                                    var v47 int16 = v2
                                    v5 += (v5 * v40)
                                    _ = v47
                                }
                                v4 += v4
                                _ = v36
                                _ = v37
                                _ = v40
                                _ = v41
                            }
                            if false {
                                var v48 int8 = -1
                                var v49 int32 = 0
                                _ = v48
                                _ = v49
                            } else {
                                if (true && (v31 >= float32((-2.25 - 3.14159)))) {
                                    v31 = (v31 / float32((0.1 - 1.5)))
                                    v50 := GenS3{F0: v5}
                                    _ = v50
                                }
                                v51 := GenS4{F0: v4, F1: 1, F2: v4}
                                v52 := GenS5{F0: 9223372036854775807, F1: -9223372036854775808}
                                _ = v51
                                _ = v52
                            }
                        }
                        v53 := GenS6{F0: "a", F1: v5}
                        _ = v53
                    } else {
                        var v54 int32 = 351
                        v4 = (v4 + (^(v4)))
                        fmt.Println(fi32)
                        var v55 bool = false
                        _ = v54
                        _ = v55
                    }
                }
                var v56 int = (fi34 + -640)
                _ = fi34
                _ = v56
            }
            _ = fi32
        }
        v2 = v2
    } else {
        v57 := GenS5{F0: 4294967295, F1: 128}
        v3 = append(v3, ("" + ("" + "foo bar")))
        _ = v57
    }
    switch 855 {
    case 32768:
        if (1 != 73) {
            v31 = 3.14159
            for fi58 := 0; fi58 < 11; fi58++ {
                fmt.Println(v5)
                _ = fi58
            }
            if true {
                fmt.Println(v5, v4, v4)
                v59 := GenS7{F0: v4}
                _ = v59
            }
        }
    case 1:
        for fi60 := 0; fi60 < 19; fi60++ {
            v5 = 127
            v31 = float32(-2.25)
            _ = fi60
        }
    }
    v31 -= ((v31 / v31) * v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
