package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint64 }
func genFn0(p0 string, p1 uint, p2 int) int8 {
    switch p0 {
    case "\n\t":
        switch 32768 {
        case 42:
            var v5 string = (p0 + "foo bar")
            _ = v5
        }
    case "":
        switch 32767 {
        case 425:
            v6 := GenS1{F0: 525}
            _ = v6
        case 2147483647:
            v7 := GenS2{F0: 42}
            _ = v7
        case 205:
            v8 := GenS1{F0: 982}
            _ = v8
        default:
            if (!(p0 > p0)) {
                v9 := GenS3{F0: 133, F1: -32768}
                var v10 int64 = 9223372036854775807
                v11 := []int{-818, p2}
                _ = v10
                _ = v11
                _ = v9
            }
        }
    case "hello":
        for fi12 := 0; fi12 < 6; fi12++ {
            p0 = "foo bar"
            for fi13 := 0; fi13 < 4; fi13++ {
                if (fi12 < fi12) {
                    if false {
                        v14 := []int{}
                        var v15 float32 = 0.1
                        fmt.Println(p0, p1, p0)
                        var v16 int = 0
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        switch p0 {
                        case "\n\t":
                            v17 := []int{(fi13 - (fi13 >> (((^(p1))) & 63))), 128, fi13}
                            _ = v17
                        case "a":
                            v18 := []int{(^(fi12)), fi13, p2}
                            _ = v18
                        default:
                            v19 := GenS4{F0: 0, F1: 148}
                            _ = v19
                        }
                    }
                    v20 := []int{(fi12 | (p2 ^ (-(fi12)))), (p2 - fi12), (p2 ^ (fi13 + 128)), 0}
                    var v21 int32 = -551
                    _ = v20
                    _ = v21
                }
                v22 := GenS2{F0: 2}
                _ = fi13
                _ = v22
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS5{F0: 0, F1: 52}
                v25 := []int{42, (fi12 % ((-2147483648) | 1)), p2, -868}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi12
        }
    }
    var v26 float64 = (-2.25 + (1.5 - (-2.25 * 3.14159)))
    v27 := GenS6{F0: 60, F1: (p1 | (p1 ^ (p1 - p1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 1
}
func genFn1() uint64 {
    var v28 uint = 42
    if (v28 == v28) {
        var v29 float32 = 3.14159
        _ = v29
    } else {
        switch 128 {
        case 0:
            v30 := GenS4{F0: -1, F1: 157}
            _ = v30
        case 2147483647:
            if false {
                v31 := []int{}
                v28 += 18446744073709551615
                v28 -= 174
                _ = v31
            } else {
                var v32 float32 = (1e10 / 1e10)
                var v33 string = "hello"
                v28 = (v28 >> ((v28) & 63))
                _ = v32
                _ = v33
            }
        }
    }
    _ = v28
    return 447
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int = 32768
    v4 := []uint32{}
    if false {
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 uint = 65535
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch 1 {
            case 128:
                v36 := GenS3{F0: 42, F1: 0}
                _ = v36
            }
            _ = fi34
            _ = v35
        }
        if (42 < 2) {
            _ = len(v1)
            v3 = v3
            for fi37 := 0; fi37 < 7; fi37++ {
                if (720 > 155) {
                    v3 += v3
                    if (0.1 != -2.25) {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            var v39 uint16 = 32767
                            v3 += 128
                            switch 127 {
                            case -19:
                                v39 -= 65535
                            default:
                                var v40 int16 = -32768
                                _ = v40
                            }
                            switch ("a" + "foo bar") {
                            case "a":
                                var v41 uint = 255
                                _ = v41
                            default:
                                fmt.Println(v3)
                            }
                            _ = fi38
                            _ = v39
                        }
                        var v42 string = "\n\t"
                        _ = v42
                    } else {
                        switch 32768 {
                        case 1:
                            switch "" {
                            case "foo bar":
                                v3 = v3
                            case "":
                                v43 := v2["foo bar"]
                                _ = v43
                            }
                        case -9223372036854775808:
                            v1["\n\t"] = -229
                        case 4294967295:
                            var v44 float32 = 1.5
                            _ = v44
                        }
                        var v45 int64 = 42
                        v46 := v2["foo bar"]
                        _ = v45
                        _ = v46
                    }
                    v47 := GenS3{F0: 42, F1: 32767}
                    _ = v47
                } else {
                    fmt.Println(v3)
                }
                v48 := GenS6{F0: 74, F1: 2}
                _ = fi37
                _ = v48
            }
        }
    } else {
        var v49 float64 = ((1e10 - 1e10) * (0.0 + float64(float32((1.5 - 0.1)))))
        v3 += -470
        _ = v49
    }
    _ = len(v1)
    var v50 string = ((("foo bar" + "hello") + "\n\t") + "\n\t")
    var v51 float64 = 3.14159
    var v52 uint8 = 255
    var v53 bool = false
    if true {
        for fi54 := 0; fi54 < 2; fi54++ {
            v1["foo bar"] = (fi54 - fi54)
            _ = fi54
        }
    }
    v53 = v53
    v55 := GenS5{F0: (^(v52)), F1: 255}
    v3 -= 0
    fmt.Println(v3, v50)
    switch v50 {
    case "":
        for fi56 := 0; fi56 < 13; fi56++ {
            v50 = v50
            _ = fi56
        }
    case "\n\t":
        for fi57 := 0; fi57 < 12; fi57++ {
            for fi58 := 0; fi58 < 20; fi58++ {
                switch (fi57 + fi57) {
                case 32768:
                    var v59 int = (fi57 ^ 128)
                    _ = v59
                case 4294967295:
                    v4 = append(v4, 2147483647)
                case 1:
                    var v60 float32 = float32(0.0)
                    _ = v60
                }
                v52 = v52
                if (!(1 > fi57)) {
                    v61 := GenS5{F0: (v52 ^ v52), F1: 65535}
                    _ = v61
                } else {
                    switch fi58 {
                    case 127:
                        var v62 uint8 = (^(v52))
                        _ = v62
                    case -445:
                        switch 2147483647 {
                        case 32768:
                            var v63 uint64 = 0
                            _ = v63
                        default:
                            _ = len(v4)
                        }
                    default:
                        v51 = v51
                    }
                }
                _ = fi58
            }
            v3 -= -9223372036854775808
            var v64 int = v3
            _ = fi57
            _ = v64
        }
    }
    if v53 {
        v52 = (v52 / ((v52) | 1))
        switch v50 {
        case "\n\t":
            v65 := GenS7{F0: 1, F1: 65535, F2: 1}
            _ = v65
        case "hello":
            fmt.Println(v51)
        default:
            v53 = true
        }
        fmt.Println(v3, v3, v51)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v55
}
