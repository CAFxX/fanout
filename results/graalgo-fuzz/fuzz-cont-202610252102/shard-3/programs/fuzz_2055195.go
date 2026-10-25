package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 int32, p2 int32) uint {
    if (1 < 255) {
        if false {
            v5 := GenS1{F0: 83, F1: -23, F2: false}
            v6 := []int{1, 42}
            _ = v5
            _ = v6
        }
        var v7 int8 = 100
        _ = v7
    }
    if ((p2 | (p1 - (p1 >> ((1) & 63)))) <= -426) {
        var v8 uint64 = 485
        var v9 int8 = 59
        switch 0 {
        case 9223372036854775807:
            v9 = 42
        case -9223372036854775808:
            switch "foo bar" {
            case "":
                var v10 string = "foo bar"
                _ = v10
            case "a":
                v8 += v8
            default:
                var v11 uint64 = (v8 - (^(v8)))
                _ = v11
            }
        }
        _ = v8
        _ = v9
    }
    var v12 string = ("\n\t" + "\n\t")
    var v13 string = v12
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 18446744073709551615
}
func genFn1(p0 int) uint8 {
    if false {
        p0 = (p0 & (p0 * p0))
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint64 = 2
        var v16 float64 = 1e10
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            p0 += (fi18 << ((65535) & 63))
            _ = fi18
        }
        v19 := GenS2{F0: 219, F1: 1, F2: 42}
        _ = fi17
        _ = v19
    }
    fmt.Println(p0)
    v20 := GenS3{F0: -32768, F1: 1, F2: 559}
    _ = p0
    _ = v20
    return 145
}

func main() {
    v1 := []int{-1, -1, -289, 1}
    var v2 float64 = ((float64(1e10) / 1e10) + 0.0)
    v3 := []string{(("hello" + "\n\t") + "foo bar")}
    var v4 string = "hello"
    v21 := GenS2{F0: 0, F1: 651, F2: 1}
    switch (v4 + (v4 + ((((v4 + "foo bar") + (v4 + "")) + "hello") + v4))) {
    case "a":
        if false {
            v3 = append(v3, ((v4 + v4) + (v4 + "")))
            fmt.Println(v4)
        }
    case "":
        if false {
            v4 = v4
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int32 = 833
                switch fi22 {
                case -614:
                    if (false || false) {
                        var v24 uint64 = 664
                        if true {
                            v2 += -2.25
                            v2 += (v2 - v2)
                        } else {
                            switch (fi22 << ((2) & 63)) {
                            case 121:
                                for fi25 := 0; fi25 < 4; fi25++ {
                                    v3 = append(v3, "\n\t")
                                    _ = fi25
                                }
                            case 32767:
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    fmt.Println(v4, fi26, fi22)
                                    v4 = ((v4 + (((v4 + "a") + "foo bar") + v4)) + v4)
                                    _ = fi26
                                }
                            case -968:
                                var v27 int8 = -1
                                _ = v27
                            }
                        }
                        for fi28 := 0; fi28 < 19; fi28++ {
                            var v29 uint32 = 4294967295
                            if len(v1) > 0 {
                                v1[0] = (-(fi22))
                            }
                            var v30 int8 = 74
                            v31 := GenS4{F0: 1}
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v24
                    } else {
                        for fi32 := 0; fi32 < 2; fi32++ {
                            _ = len(v3)
                            _ = fi32
                        }
                    }
                case 32768:
                    var v33 float64 = v2
                    _ = v33
                default:
                    v4 = v4
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    if true {
                        _ = len(v3)
                    } else {
                        fmt.Println(fi34)
                    }
                    _ = fi34
                }
                _ = fi22
                _ = v23
            }
        }
    }
    v35 := GenS3{F0: 42, F1: -470, F2: 389}
    fmt.Println(v4, v4)
    var v36 bool = (!(false && (0 <= -2147483648)))
    v2 = v2
    var v37 bool = (42 == 18446744073709551615)
    for fi38 := 0; fi38 < 13; fi38++ {
        if (!v36) {
            switch (v4 + ("hello" + "hello")) {
            case "a":
                switch (fi38 % (((fi38 << ((genFn0(65535, -231, 1)) & 63))) | 1)) {
                case -270:
                    v37 = (1 >= -1)
                case 2147483647:
                    v39 := GenS1{F0: -1, F1: 42, F2: v37}
                    _ = v39
                case 176:
                    if (!(2 < 42)) {
                        var v40 int64 = 9223372036854775807
                        for fi41 := 0; fi41 < 16; fi41++ {
                            v42 := GenS5{F0: ((0.0 * ((0.1 - 1e10) / 0.1)) > ((1e10 - float32((v2 + v2))) + float32(-2.25)))}
                            var v43 int64 = (v40 - 81)
                            fmt.Println(fi41)
                            v4 = ("hello" + v4)
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        _ = v40
                    }
                default:
                    v44 := GenS1{F0: -1, F1: 127, F2: v37}
                    _ = v44
                }
            }
            v37 = true
            v36 = v37
        }
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
