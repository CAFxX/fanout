package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 string; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int8; F1 int32; F2 uint32 }
type GenS8 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 uint) int64 {
    switch 344 {
    case -883:
        p0 += (p0 + 0)
    }
    if true {
        var v5 uint32 = 2
        _ = v5
    } else {
        v6 := GenS1{F0: 2147483647, F1: 100}
        _ = v6
    }
    v7 := GenS2{F0: -128}
    _ = p0
    _ = v7
    return 206
}

func main() {
    v1 := []int{}
    var v2 uint = 0
    var v3 int32 = 42
    v4 := make(map[string]int)
    switch 1 {
    case 4294967295:
        var v8 uint32 = 656
        _ = v8
    case 444:
        switch "hello" {
        case "hello":
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(v3)
                switch 42 {
                case 0:
                    v10 := GenS1{F0: (fi9 % ((-2147483648) | 1)), F1: 42}
                    _ = v10
                case 4294967295:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v2 = 255
                        _ = fi11
                    }
                default:
                    if (922 >= 32767) {
                        v12 := GenS2{F0: -128}
                        v2 -= (v2 / ((v2) | 1))
                        _ = v12
                    } else {
                        v4[("hello" + "a")] = 9223372036854775807
                        v13 := GenS1{F0: fi9, F1: 100}
                        for fi14 := 0; fi14 < 8; fi14++ {
                            fmt.Println(v3, fi9, fi9)
                            _ = fi14
                        }
                        _ = v13
                    }
                }
                var v15 bool = ((false && false) && (!true))
                _ = fi9
                _ = v15
            }
        default:
            if true {
                switch 4294967295 {
                case 149:
                    v16 := GenS3{F0: "hello", F1: 128}
                    _ = v16
                default:
                    if false {
                        if false {
                            if (1.5 != float64(0.1)) {
                                switch 389 {
                                case 1:
                                    var v17 int64 = -1
                                    _ = v17
                                case 0:
                                    v18 := GenS2{F0: -128}
                                    _ = v18
                                case -591:
                                    v2 += 406
                                default:
                                    v19 := GenS1{F0: 0, F1: -128}
                                    _ = v19
                                }
                                v2 = v2
                                if false {
                                    for fi20 := 0; fi20 < 5; fi20++ {
                                        v2 = (^(v2))
                                        v21 := GenS4{F0: ((((-2.25 + 0.1) / float64((float32(float64(1.5)) - 3.14159))) * 1e10) == (3.14159 / 1.0)), F1: 4294967295}
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                            } else {
                                var v22 int32 = v3
                                var v23 float32 = ((0.1 / 1.0) - 1e10)
                                if true {
                                    for fi24 := 0; fi24 < 19; fi24++ {
                                        var v25 uint = v2
                                        var v26 uint = v2
                                        if len(v1) > 0 {
                                            v1[0] = fi24
                                        }
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                    fmt.Println(v3, v2, v23)
                                    v27 := GenS5{F0: 599, F1: "foo bar", F2: 978}
                                    _ = v27
                                }
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v28 bool = false
                            var v29 uint = (v2 | v2)
                            _ = v28
                            _ = v29
                        }
                        v3 = v3
                        fmt.Println(v2, v3)
                    }
                }
                switch (("a" + "hello") + "hello") {
                case "a":
                    switch 9223372036854775807 {
                    case 128:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "":
                    v30 := GenS6{F0: "hello", F1: 0, F2: 970}
                    _ = v30
                case "\n\t":
                    var v31 float32 = 1e10
                    _ = v31
                default:
                    _ = len(v1)
                }
                if len(v1) > 0 {
                    v1[0] = -2147483648
                }
            } else {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := GenS7{F0: -99, F1: 1, F2: 2147483647}
                    var v34 int = (fi32 * (fi32 * (-(fi32))))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v35 uint64 = 465
                _ = v35
            }
        }
    case 233:
        var v36 string = "\n\t"
        _ = v36
    default:
        var v37 uint16 = 32767
        _ = v37
    }
    v3 -= v3
    if true {
        if true {
            if (true && true) {
                v3 = (v3 / ((v3) | 1))
            }
            fmt.Println(v3, v3)
        }
    } else {
        v38 := GenS8{F0: (v3 <= (^(v3))), F1: 697, F2: 1}
        _ = v38
    }
    v39 := v4[("a" + "hello")]
    var v40 uint64 = 1
    var v41 int8 = 85
    if ((v41 - v41) <= v41) {
        v4["\n\t"] = v39
    } else {
        v41 -= v41
        if true {
            var v42 bool = false
            _ = v42
        }
    }
    var v43 int64 = -482
    v39 = int(v43)
    var v44 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
}
