package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32; F1 int; F2 int }
type GenS5 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int; F1 uint64 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint32 = 235
        v9 := GenS1{F0: 127, F1: "hello", F2: 504}
        _ = fi7
        _ = v8
        _ = v9
    }
    if (309 >= 42) {
        var v10 int16 = 111
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        var v11 uint = 4294967295
        _ = v11
    default:
        var v12 int8 = -1
        _ = v12
    }
    var v13 int = -9223372036854775808
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            if false {
                var v16 uint8 = 255
                v17 := GenS2{F0: 42, F1: 2}
                _ = v16
                _ = v17
            } else {
                v18 := []int{}
                fmt.Println(fi15)
                var v19 string = "\n\t"
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        fmt.Println(v13, v13)
        v20 := []int{127, fi14, v13}
        _ = fi14
        _ = v20
    }
    _ = v13
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 628
    var v3 int8 = 65
    v4 := make(map[int]bool)
    var v5 uint = 4294967295
    v6 := []string{"hello", "\n\t"}
    v2 = v2
    for fi21 := 0; fi21 < 3; fi21++ {
        fmt.Println(v2, v2, v5)
        fmt.Println(v5, v5)
        _ = fi21
    }
    fmt.Println(v5, v3)
    var v22 string = (("a" + (("hello" + ("a" + "foo bar")) + ((("foo bar" + ("a" + "\n\t")) + "foo bar") + ("hello" + "a")))) + "")
    v22 = ""
    v5 -= (v5 + 220)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 float64 = 0.1
        _ = fi23
        _ = v24
    }
    if true {
        v25 := GenS3{F0: (v5 * v5), F1: -1}
        v26 := GenS4{F0: 65535, F1: 1, F2: 0}
        if false {
            v27 := GenS1{F0: v3, F1: (v22 + (v22 + "\n\t")), F2: v2}
            var v28 float32 = float32(3.14159)
            if false {
                var v29 float32 = 1e10
                v30 := GenS3{F0: (v5 ^ (v5 | v5)), F1: -866}
                switch ("a" + v22) {
                case "\n\t":
                    for fi31 := 0; fi31 < 9; fi31++ {
                        switch fi31 {
                        case -876:
                            v2 = v2
                        case -1:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                var v33 uint16 = 32767
                                v34 := GenS5{F0: v2, F1: -651, F2: genFn0()}
                                var v35 bool = (false || (v33 == (v33 & (v33 ^ (^(v33))))))
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        default:
                            fmt.Println(v29, v3)
                        }
                        _ = fi31
                    }
                case "":
                    v4[-9223372036854775808] = true
                case "hello":
                    switch ((v22 + "") + ("a" + (("hello" + v22) + v22))) {
                    case "a":
                        v36 := GenS6{F0: 452, F1: 1, F2: 849}
                        _ = v36
                    case "\n\t":
                        v37 := GenS3{F0: 18446744073709551615, F1: 1000}
                        _ = v37
                    case "hello":
                        var v38 float32 = (((0.0 - (v29 + v29)) * v28) / v28)
                        _ = v38
                    default:
                        fmt.Println(v22, v29)
                    }
                default:
                    switch -929 {
                    case 329:
                        if true {
                            for fi39 := 0; fi39 < 2; fi39++ {
                                v4[-591] = true
                                _ = fi39
                            }
                            delete(v1, -479)
                            if true {
                                fmt.Println(v22)
                            } else {
                                v40 := GenS7{F0: 259}
                                _ = v40
                            }
                        }
                    default:
                        switch -356 {
                        case -562:
                            _ = len(v6)
                        }
                    }
                }
                _ = v29
                _ = v30
            }
            _ = v27
            _ = v28
        }
        _ = v25
        _ = v26
    }
    v41 := GenS8{F0: -1, F1: v2}
    v42 := GenS5{F0: 4294967295, F1: 2147483647, F2: 59}
    if false {
        fmt.Println(v5, v5)
        for fi43 := 0; fi43 < 2; fi43++ {
            v2 = (v2 - v2)
            v44 := GenS9{F0: fi43, F1: v2}
            switch v22 {
            case "hello":
                for fi45 := 0; fi45 < 5; fi45++ {
                    v3 -= (v3 / ((v3) | 1))
                    _ = fi45
                }
            case "\n\t":
                var v46 string = v22
                _ = v46
            case "":
                v47 := GenS9{F0: 4294967295, F1: v2}
                _ = v47
            }
            _ = fi43
            _ = v44
        }
    } else {
        delete(v4, 285)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
