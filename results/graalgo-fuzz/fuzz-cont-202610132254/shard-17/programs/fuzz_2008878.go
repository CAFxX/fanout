package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 int64, p2 int) bool {
    v5 := []int{(p2 ^ (p2 ^ -2147483648))}
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(p1, fi6)
        var v7 uint32 = 65535
        _ = len(v5)
        _ = fi6
        _ = v7
    }
    if len(v5) > 0 {
        v5[0] = -60
    }
    v8 := GenS1{F0: 32767, F1: (int32(p1) < 42)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := []string{(("foo bar" + ("a" + (("hello" + (("foo bar" + ("a" + "a")) + "")) + ("hello" + "")))) + ("" + "")), "foo bar", "\n\t"}
    var v4 string = "hello"
    var v9 uint32 = 4294967295
    var v10 int64 = -9223372036854775808
    v11 := GenS1{F0: 295, F1: (!(2 <= 769))}
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: 1, F1: true}
        v14 := GenS2{F0: 32767, F1: ("hello" + (v4 + v4)), F2: 42}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = "hello"
    v16 := GenS3{F0: -24, F1: v9}
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS4{F0: fi17}
        var v19 int16 = 32767
        _ = fi17
        _ = v18
        _ = v19
    }
    v2[0] = -1
    v4 = "hello"
    switch "hello" {
    case "hello":
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float32 = (0.0 / 1.0)
            _ = fi20
            _ = v21
        }
    case "\n\t":
        v4 = (v15 + "foo bar")
    case "foo bar":
        for fi22 := 0; fi22 < 13; fi22++ {
            v4 = ("\n\t" + "")
            _ = fi22
        }
    default:
        switch 1 {
        case -501:
            switch v4 {
            case "\n\t":
                for fi23 := 0; fi23 < 14; fi23++ {
                    v9 = v9
                    if true {
                        switch fi23 {
                        case -110:
                            switch v4 {
                            case "hello":
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    v1[((v15 + ("foo bar" + ("\n\t" + ""))) + v4)] = v4
                                    var v25 uint16 = 857
                                    v10 -= -9223372036854775808
                                    _ = len(v3)
                                    _ = fi24
                                    _ = v25
                                }
                            case "":
                                var v26 float64 = 1e10
                                _ = v26
                            default:
                                v15 = v4
                            }
                        case 9223372036854775807:
                            for fi27 := 0; fi27 < 9; fi27++ {
                                v10 = 42
                                _ = fi27
                            }
                        case -13:
                            var v28 uint32 = v9
                            _ = v28
                        default:
                            v29 := GenS1{F0: 255, F1: true}
                            _ = v29
                        }
                        fmt.Println(fi23, v10)
                    }
                    for fi30 := 0; fi30 < 17; fi30++ {
                        _ = len(v1)
                        if len(v3) > 0 {
                            v3[0] = (((((((("foo bar" + v15) + (((v4 + ("hello" + (v4 + "a"))) + v15) + (v15 + "hello"))) + (("hello" + "foo bar") + "")) + v4) + ((v15 + ("" + "")) + ("foo bar" + v4))) + v4) + (v4 + ((v15 + ((v15 + "") + v4)) + "hello"))) + "\n\t")
                        }
                        var v31 uint16 = 255
                        _ = fi30
                        _ = v31
                    }
                    _ = fi23
                }
            }
        case 277:
            fmt.Println(v15, v15, v10)
        default:
            v32 := GenS3{F0: -724, F1: (v9 & (v9 / (((v9 ^ (v9 & (v9 | v9)))) | 1)))}
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := v1["hello"]
        var v35 uint32 = 2
        v15 = ""
        _ = fi33
        _ = v34
        _ = v35
    }
    v10 += -9223372036854775808
    var v36 int8 = 114
    if (-128 <= v36) {
        _ = len(v1)
        v4 = (v15 + (v4 + v4))
        v37 := GenS2{F0: 42, F1: (((v4 + (((v15 + (("a" + "a") + "a")) + (v4 + (v15 + "hello"))) + "foo bar")) + v4) + "hello"), F2: 42}
        _ = v37
    } else {
        v38 := GenS5{F0: 1, F1: 913}
        v9 = 266
        _ = v38
    }
    v4 = (((("foo bar" + v4) + v15) + "foo bar") + (((("" + ((v15 + (v4 + (("" + v4) + v4))) + "")) + "hello") + (((v4 + v15) + ("hello" + (((v4 + "hello") + "\n\t") + "a"))) + v15)) + ("foo bar" + v15)))
    _ = len(v3)
    if false {
        v15 = ((("\n\t" + "") + "hello") + "a")
    } else {
        fmt.Println(v9, v9, v36)
        if (false && true) {
            v36 -= (v36 << ((1) & 63))
        } else {
            v36 += (v36 / (((v36 * v36)) | 1))
        }
        v39 := GenS6{F0: v36, F1: v15}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v9
}
