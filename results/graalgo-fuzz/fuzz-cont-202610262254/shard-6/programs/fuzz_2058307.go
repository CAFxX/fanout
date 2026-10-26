package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
func genFn0(p0 uint16, p1 bool, p2 uint8) int16 {
    p2 = 13
    var v5 uint16 = (p0 - p0)
    p1 = (true && (false && false))
    if (42 > 193) {
        v6 := []int{127, 939}
        switch "" {
        case "a":
            switch "\n\t" {
            case "a":
                v7 := GenS1{F0: 0, F1: "", F2: 714}
                _ = v7
            default:
                if p1 {
                    v8 := GenS2{F0: 0, F1: (^(p2)), F2: 0}
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(p0, p1)
                        for fi10 := 0; fi10 < 20; fi10++ {
                            if p1 {
                                p1 = (true && false)
                            } else {
                                var v11 uint32 = 615
                                _ = v11
                            }
                            p1 = (!(!p1))
                            if true {
                                var v12 int16 = 1000
                                _ = v12
                            } else {
                                v13 := GenS2{F0: 127, F1: (p2 + 0), F2: 443}
                                v6 = append(v6, 127)
                                _ = v13
                            }
                            if (0 <= 1) {
                                fmt.Println(p2)
                            } else {
                                var v14 uint32 = 743
                                var v15 int = (fi9 % ((int(p2)) | 1))
                                _ = v14
                                _ = v15
                            }
                            p1 = p1
                            _ = fi10
                        }
                        _ = fi9
                    }
                    _ = v8
                }
            }
        case "hello":
            var v16 uint32 = 279
            _ = v16
        case "":
            switch 127 {
            case 583:
                for fi17 := 0; fi17 < 18; fi17++ {
                    p2 = p2
                    p0 += (p0 | v5)
                    _ = fi17
                }
            default:
                var v18 uint64 = 0
                _ = v18
            }
        default:
            switch 9223372036854775807 {
            case 9223372036854775807:
                v19 := GenS1{F0: -6, F1: "", F2: 2147483647}
                _ = v19
            case 1:
                if p1 {
                    p1 = p1
                    switch int(p0) {
                    case -377:
                        fmt.Println(v5)
                    default:
                        v6 = append(v6, 32768)
                    }
                }
            }
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    var v1 string = ("" + "a")
    var v2 uint16 = 45
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]string)
    fmt.Println(v1, v1, v2)
    var v20 int16 = -179
    for fi21 := 0; fi21 < 15; fi21++ {
        _ = len(v4)
        var v22 int8 = 42
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: -1, F1: 2, F2: 952}
    v24 := GenS1{F0: 100, F1: "foo bar", F2: 65535}
    switch (v1 + (v1 + (v1 + ("foo bar" + v1)))) {
    case "\n\t":
        var v25 int64 = -416
        _ = v25
    default:
        for fi26 := 0; fi26 < 10; fi26++ {
            v3 = 127
            v27 := GenS2{F0: -128, F1: 0, F2: 4294967295}
            v28 := GenS1{F0: 42, F1: v1, F2: 42}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v29 := []int{v3, v3, 313, (v3 * (v3 - (^(v3))))}
    _ = len(v29)
    if (false || false) {
        switch v1 {
        case "":
            fmt.Println(v2)
        case "a":
            switch v1 {
            case "foo bar":
                for fi30 := 0; fi30 < 5; fi30++ {
                    v3 -= (-(v3))
                    v31 := GenS3{F0: 911, F1: -24}
                    _ = fi30
                    _ = v31
                }
            default:
                v29 = append(v29, (v3 | (v3 & v3)))
            }
        default:
            var v32 int = 4294967295
            _ = v32
        }
        v2 = v2
        v3 = (v3 | -9223372036854775808)
    } else {
        var v33 float64 = (((1.5 - (0.1 + 1e10)) - ((0.1 + 1.5) + (1.5 - 0.0))) * -2.25)
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v3, fi34)
            fmt.Println(v33)
            v35 := GenS4{F0: 1}
            _ = fi34
            _ = v35
        }
        v36 := GenS5{F0: "a"}
        _ = v33
        _ = v36
    }
    var v37 int64 = -9223372036854775808
    if len(v29) > 0 {
        v29[0] = -893
    }
    if true {
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 float64 = 0.0
            _ = fi38
            _ = v39
        }
    } else {
        var v40 int32 = -51
        v29 = append(v29, (v3 % ((v3) | 1)))
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}
