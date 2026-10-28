package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16) int64 {
    var v6 uint16 = 42
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = 1
            fmt.Println(v8, v8)
            fmt.Println(fi7, fi7, v6)
            _ = fi7
            _ = v8
        }
    case 128:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint8 = 105
            var v11 float64 = 3.14159
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        v12 := []int{-634, 459}
        _ = v12
    }
    _ = p0
    _ = v6
    return -184
}

func main() {
    var v1 int32 = 980
    v2 := []int{}
    v3 := make(map[int]uint32)
    var v4 uint = 65535
    v5 := make(map[int]int)
    fmt.Println(v1, v1)
    v1 -= (-(v1))
    var v13 float64 = 0.1
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint32 = 4294967295
        for fi17 := 0; fi17 < 8; fi17++ {
            switch "\n\t" {
            case "foo bar":
                switch (("" + "") + ("\n\t" + "a")) {
                case "":
                    fmt.Println(fi15, v4, fi15)
                default:
                    var v18 bool = false
                    _ = v18
                }
            default:
                v16 = uint32(v4)
            }
            fmt.Println(v4)
            delete(v5, (fi15 % (((fi15 << (((v4 / (((v4 - v4)) | 1))) & 63))) | 1)))
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    if false {
        v1 -= -1
    } else {
        v13 = v13
    }
    _ = len(v2)
    fmt.Println(v4, v1)
    v13 = v13
    fmt.Println(v13)
    if ((v4 / (((^(v4))) | 1)) < (v4 / (((v4 * v4)) | 1))) {
        _ = len(v5)
        if true {
            for fi19 := 0; fi19 < 16; fi19++ {
                switch "a" {
                case "hello":
                    v20 := GenS1{F0: -32768}
                    _ = v20
                default:
                    v14 -= -2.25
                }
                var v21 uint8 = 205
                var v22 uint16 = 2
                _ = fi19
                _ = v21
                _ = v22
            }
        } else {
            var v23 string = ("foo bar" + "hello")
            fmt.Println(v14, v13)
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int16 = 32767
                v26 := GenS2{F0: fi24, F1: int16(fi24)}
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int32 = v1
            fmt.Println(fi27, v4)
            _ = fi27
            _ = v28
        }
        switch ("" + "") {
        case "hello":
            v14 -= (-2.25 - v14)
        case "":
            v1 -= -988
        }
        if true {
            v29 := GenS3{F0: 20, F1: -1}
            if (!(!false)) {
                if false {
                    fmt.Println(v14)
                    switch 78 {
                    case 32767:
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v14 -= -2.25
                            _ = fi30
                        }
                    case -2147483648:
                        for fi31 := 0; fi31 < 3; fi31++ {
                            _ = len(v2)
                            if true {
                                switch fi31 {
                                case 0:
                                    var v32 uint64 = 683
                                    _ = v32
                                case 1:
                                    if len(v2) > 0 {
                                        v2[0] = fi31
                                    }
                                default:
                                    var v33 int32 = v1
                                    _ = v33
                                }
                                delete(v5, 708)
                            }
                            _ = fi31
                        }
                    default:
                        fmt.Println(v14)
                    }
                    v1 = v1
                }
            } else {
                for fi34 := 0; fi34 < 19; fi34++ {
                    delete(v5, (fi34 * fi34))
                    _ = fi34
                }
                var v35 float64 = 0.1
                v36 := GenS2{F0: -9223372036854775808, F1: -19}
                _ = v35
                _ = v36
            }
            _ = v29
        }
    }
    fmt.Println(v14, v1, v14)
    v37 := v3[-9223372036854775808]
    _ = len(v5)
    if (v4 == v4) {
        v38 := GenS4{F0: 1, F1: 247}
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS1{F0: -1}
            fmt.Println(v13, v1, v4)
            v14 -= -2.25
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    v4 = uint(v37)
    switch "hello" {
    case "foo bar":
        if (false && true) {
            var v41 bool = false
            var v42 int8 = 1
            if false {
                v1 = 373
                v43 := v5[42]
                var v44 uint8 = 0
                _ = v43
                _ = v44
            }
            _ = v41
            _ = v42
        } else {
            for fi45 := 0; fi45 < 18; fi45++ {
                v1 += v1
                _ = fi45
            }
            v46 := GenS5{F0: (-(v1))}
            _ = v46
        }
    case "\n\t":
        v14 -= 0.0
    default:
        var v47 bool = true
        _ = v47
    }
    var v48 float64 = (0.1 - (v13 - v13))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v48
    _ = v5
}
