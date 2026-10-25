package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0() uint8 {
    v6 := GenS1{F0: 658, F1: -1}
    v7 := []int{128, 0}
    v7 = append(v7, 0)
    _ = len(v7)
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint32 {
    var v8 float64 = 1.5
    var v9 int8 = -1
    v8 = (v8 * (0.0 - v8))
    var v10 bool = ((false || true) && false)
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = ("hello" + "")
    var v2 int8 = -5
    v3 := make(map[uint32]int)
    var v4 uint16 = 65535
    v5 := make(map[uint32]int)
    switch 127 {
    case 461:
        v1 = v1
    case 128:
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                fmt.Println(v1, fi12)
                _ = fi12
            }
            v13 := []int{(fi11 >> ((1) & 63)), fi11, (fi11 / (((fi11 / (((fi11 >> ((181) & 63))) | 1))) | 1))}
            v1 = ("foo bar" + ((v1 + "") + (v1 + v1)))
            _ = fi11
            _ = v13
        }
    case 0:
        if true {
            v14 := GenS1{F0: 42, F1: -90}
            for fi15 := 0; fi15 < 4; fi15++ {
                v2 = (v2 % ((-120) | 1))
                _ = fi15
            }
            v2 -= (v2 - -1)
            _ = v14
        }
    default:
        switch -9223372036854775808 {
        case 128:
            v4 += (v4 | v4)
        case 1:
            for fi16 := 0; fi16 < 14; fi16++ {
                v4 += (v4 | v4)
                var v17 string = v1
                _ = fi16
                _ = v17
            }
        default:
            var v18 uint64 = 1
            _ = v18
        }
    }
    if true {
        if ((v1 + ((v1 + "hello") + "a")) >= v1) {
            v1 = ("a" + "foo bar")
        }
        v19 := v3[2147483648]
        for fi20 := 0; fi20 < 14; fi20++ {
            if false {
                fmt.Println(fi20, v1, v19)
            } else {
                v21 := []int{891, (-(fi20)), (-(fi20))}
                var v22 int8 = v2
                for fi23 := 0; fi23 < 10; fi23++ {
                    if false {
                        switch (^(fi23)) {
                        case 32767:
                            _ = len(v21)
                        case -1:
                            if false {
                                v4 = v4
                                switch (v1 + (v1 + (v1 + (v1 + "")))) {
                                case "foo bar":
                                    if (0 < uint32(v4)) {
                                        var v24 int64 = 9223372036854775807
                                        _ = v24
                                    } else {
                                        v22 = 42
                                        v25 := v5[580]
                                        if (3.14159 == 0.0) {
                                            var v26 uint32 = 607
                                            _ = v26
                                        } else {
                                            var v27 uint64 = 18446744073709551615
                                            _ = v27
                                        }
                                        switch fi20 {
                                        case 9223372036854775807:
                                            var v28 uint16 = v4
                                            _ = v28
                                        case 127:
                                            switch (-(fi20)) {
                                            case 0:
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            default:
                                                v3[2147483647] = (fi23 / ((v25) | 1))
                                            }
                                        case 32767:
                                            fmt.Println(v25)
                                        default:
                                            switch (v1 + "hello") {
                                            case "hello":
                                                switch (v1 + ((v1 + ("foo bar" + ((v1 + v1) + v1))) + (v1 + v1))) {
                                                case "":
                                                    v3[533] = (v19 | (fi20 >> ((0) & 63)))
                                                case "a":
                                                    _ = len(v21)
                                                }
                                            case "foo bar":
                                                v29 := GenS2{F0: 9223372036854775807, F1: 0, F2: -876}
                                                _ = v29
                                            default:
                                                v4 = v4
                                            }
                                        }
                                        _ = v25
                                    }
                                case "a":
                                    if (!false) {
                                        var v30 uint16 = 1
                                        var v31 int64 = 9223372036854775807
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    }
                                }
                                switch fi23 {
                                case 128:
                                    v22 = (v2 + v2)
                                }
                                v2 = (v22 * (v2 - v2))
                            } else {
                                v32 := GenS1{F0: 0, F1: 1}
                                switch (^(v19)) {
                                case 2147483647:
                                    v33 := GenS2{F0: -666, F1: 0, F2: 0}
                                    _ = v33
                                case 32768:
                                    v2 = (v22 - (v22 >> ((393) & 63)))
                                    v34 := GenS2{F0: 9223372036854775807, F1: 0, F2: 0}
                                    v4 = (v4 / (((v4 / ((v4) | 1))) | 1))
                                    _ = v34
                                case 0:
                                    v22 = v2
                                }
                                var v35 uint32 = 1
                                v36 := GenS3{F0: 18446744073709551615, F1: (-(v2))}
                                _ = v32
                                _ = v35
                                _ = v36
                            }
                        }
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
            v37 := GenS1{F0: 2, F1: (^(v2))}
            _ = fi20
            _ = v37
        }
        _ = v19
    } else {
        switch 63 {
        case -2147483648:
            for fi38 := 0; fi38 < 1; fi38++ {
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 float32 = (1.5 + (1e10 * ((0.0 + (-2.25 - (-2.25 * ((0.0 / float32(1e10)) / 0.1)))) / 1e10)))
                    fmt.Println(fi39, fi38)
                    fmt.Println(fi38, fi39)
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
        case 1:
            v41 := GenS1{F0: 419, F1: v2}
            _ = v41
        }
    }
    v1 = "foo bar"
    v1 = "hello"
    var v42 float32 = 1.5
    for fi43 := 0; fi43 < 14; fi43++ {
        v44 := v3[genFn1()]
        if false {
            v1 = v1
        }
        fmt.Println(fi43, v4, v42)
        _ = fi43
        _ = v44
    }
    v45 := GenS2{F0: -815, F1: -879, F2: 42}
    var v46 float32 = v42
    var v47 string = v1
    v48 := []int{-1, 1}
    switch 315 {
    case -7:
        var v49 uint64 = 1
        _ = v49
    }
    v1 = "hello"
    for fi50 := 0; fi50 < 11; fi50++ {
        switch 602 {
        case 127:
            v51 := GenS2{F0: 0, F1: -1, F2: -2147483648}
            _ = v51
        default:
            v4 += v4
        }
        _ = fi50
    }
    var v52 uint16 = (v4 + v4)
    fmt.Println(v47, v52, v1)
    var v53 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v52
    _ = v53
}
