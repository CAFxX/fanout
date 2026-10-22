package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 string; F1 uint32; F2 string }
type GenS5 struct { F0 string; F1 uint8; F2 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 int }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int }

func main() {
    var v1 uint = 776
    var v2 int16 = 42
    v3 := []int{4294967295, -1, 32768, 594}
    v4 := make(map[string]uint32)
    var v5 uint16 = 43
    v6 := make(map[string]uint32)
    fmt.Println(v2, v1, v5)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 128, F1: (v5 - (v5 / ((v5) | 1))), F2: (false || ((^(v1)) == (^(v1))))}
        v1 += (v1 + (v1 * (v1 % ((v1) | 1))))
        for fi9 := 0; fi9 < 1; fi9++ {
            v1 -= 255
            _ = len(v3)
            v2 += (v2 - (^(v2)))
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch -272 {
    case 0:
        var v10 int64 = 42
        _ = v10
    case 32768:
        if false {
            for fi11 := 0; fi11 < 9; fi11++ {
                v1 = 4294967295
                var v12 int64 = -1
                if true {
                    if false {
                        v2 += (-(v2))
                        v5 = 255
                    } else {
                        fmt.Println(v2, v1)
                        switch "a" {
                        case "hello":
                            v2 = v2
                        }
                        var v13 bool = false
                        _ = v13
                    }
                    var v14 uint8 = 42
                    v15 := GenS2{F0: -1, F1: (-(fi11)), F2: 694}
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        if (-2.25 < ((float32((0.1 + 1e10)) - -2.25) / 1.5)) {
                            var v17 uint16 = v5
                            _ = v17
                        }
                        _ = fi16
                    }
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v19 := GenS1{F0: uint8(v5), F1: v5, F2: true}
                        v2 += 32767
                        fmt.Println(fi18, v12)
                        v20 := GenS2{F0: 1, F1: -389, F2: 42}
                        var v21 uint32 = 2147483648
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi11
                _ = v12
            }
            fmt.Println(v2, v5)
        } else {
            v5 = (v5 / ((v5) | 1))
        }
    case -9223372036854775808:
        if true {
            var v22 int8 = -1
            switch 4294967295 {
            case 32768:
                var v23 string = "foo bar"
                _ = v23
            case 9223372036854775807:
                if (-1 < -9223372036854775808) {
                    v2 = (v2 | v2)
                    if (v1 < (v1 ^ (v1 | (^(v1))))) {
                        var v24 uint = v1
                        switch ("a" + ("a" + "foo bar")) {
                        case "\n\t":
                            var v25 bool = (0 > 128)
                            _ = v25
                        case "a":
                            v5 = 91
                        case "foo bar":
                            v26 := GenS3{F0: 0, F1: v2}
                            _ = v26
                        default:
                            fmt.Println(v24, v24, v22)
                        }
                        var v27 uint16 = v5
                        _ = v24
                        _ = v27
                    } else {
                        v28 := GenS4{F0: (("hello" + "foo bar") + "hello"), F1: 361, F2: "foo bar"}
                        v29 := GenS5{F0: "foo bar", F1: uint8(v1), F2: v22}
                        switch "a" {
                        case "a":
                            for fi30 := 0; fi30 < 19; fi30++ {
                                switch -320 {
                                case 32768:
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        if (false && (1 >= 156)) {
                                            switch -1 {
                                            case 2147483647:
                                                fmt.Println(fi31)
                                            }
                                        } else {
                                            v5 = (v5 | 42)
                                            v32 := GenS4{F0: ("" + ""), F1: 2147483647, F2: ""}
                                            _ = v32
                                        }
                                        switch fi31 {
                                        case -739:
                                            var v33 int64 = 322
                                            _ = v33
                                        }
                                        fmt.Println(fi30, fi30, v22)
                                        _ = fi31
                                    }
                                case 4294967295:
                                    _ = len(v4)
                                case 1:
                                    v34 := v4[("a" + ("\n\t" + "hello"))]
                                    _ = v34
                                }
                                _ = fi30
                            }
                        case "foo bar":
                            v35 := GenS6{F0: 697, F1: v5, F2: -1}
                            _ = v35
                        case "hello":
                            v22 = (v22 - 42)
                        }
                        _ = v28
                        _ = v29
                    }
                } else {
                    v36 := GenS7{F0: (v22 | 42), F1: "\n\t"}
                    _ = v36
                }
            case -9223372036854775808:
                if (false && (1 > -2147483648)) {
                    delete(v6, "\n\t")
                }
            }
            _ = v22
        } else {
            v37 := GenS2{F0: 2147483647, F1: 0, F2: 859}
            _ = v37
        }
    }
    fmt.Println(v2, v2)
    fmt.Println(v5, v5)
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 float64 = (-2.25 * float64(1.5))
            var v41 string = "foo bar"
            var v42 int32 = 1
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
        v43 := GenS8{F0: (^(fi38))}
        _ = fi38
        _ = v43
    }
    _ = len(v3)
    var v44 uint8 = 8
    fmt.Println(v44, v44, v44)
    var v45 int8 = 42
    var v46 bool = false
    if v46 {
        v44 = uint8(v1)
        if true {
            v3 = append(v3, 42)
        } else {
            switch -1 {
            case 32767:
                v45 += 119
            }
            v46 = (true && true)
            var v47 int = 4294967295
            _ = v47
        }
    } else {
        v44 += 255
    }
    v48 := GenS6{F0: 1, F1: v5, F2: -9223372036854775808}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 uint32 = 2147483647
            var v52 int8 = (-(v45))
            var v53 string = "\n\t"
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi49
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
}
