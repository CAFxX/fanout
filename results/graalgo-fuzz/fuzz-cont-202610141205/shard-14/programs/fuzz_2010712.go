package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 int32; F1 int64; F2 string }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 string, p1 int) bool {
    fmt.Println(p0, p1)
    var v4 uint8 = 2
    if true {
        p1 += 32767
        fmt.Println(v4, v4, p0)
    } else {
        switch "" {
        case "a":
            if false {
                switch 4294967295 {
                case 4294967295:
                    for fi5 := 0; fi5 < 16; fi5++ {
                        p1 += fi5
                        p0 = p0
                        v6 := []int{fi5, p1, fi5, fi5}
                        _ = fi5
                        _ = v6
                    }
                case 242:
                    fmt.Println(p0, v4)
                case 0:
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 int64 = 9223372036854775807
                        _ = fi7
                        _ = v8
                    }
                default:
                    p1 += (-(p1))
                }
                if (!(!(!(0 < 0)))) {
                    var v9 int8 = 0
                    switch 32768 {
                    case -513:
                        if false {
                            p0 = "hello"
                        }
                    default:
                        switch p0 {
                        case "\n\t":
                            v10 := GenS1{F0: ((^(v9)) < (^(v9)))}
                            _ = v10
                        default:
                            v9 = (v9 << ((18446744073709551615) & 63))
                        }
                    }
                    _ = v9
                }
                if (!true) {
                    fmt.Println(v4)
                    p0 = ((("" + (("hello" + (p0 + ((("a" + "foo bar") + p0) + p0))) + (("\n\t" + p0) + "foo bar"))) + ((p0 + "") + (p0 + p0))) + "a")
                    v11 := []int{-359, (p1 ^ p1)}
                    _ = v11
                }
            }
        }
        var v12 int64 = 9223372036854775807
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v4
    return ((p1 / (((p1 >> ((2) & 63))) | 1)) >= 127)
}
func genFn1(p0 string) uint32 {
    for fi13 := 0; fi13 < 14; fi13++ {
        if true {
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 uint16 = 255
                v16 := GenS2{F0: 255, F1: -1, F2: (1e10 == -2.25)}
                switch "a" {
                case "a":
                    var v17 uint32 = 860
                    _ = v17
                case "":
                    v18 := GenS2{F0: 690, F1: 42, F2: false}
                    _ = v18
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            var v19 uint = 255
            p0 = "\n\t"
            v19 = v19
            _ = v19
        }
        v20 := []int{fi13, (fi13 % (((^(fi13))) | 1)), fi13, (^(fi13))}
        _ = fi13
        _ = v20
    }
    v21 := GenS2{F0: 18446744073709551615, F1: 199, F2: false}
    _ = p0
    _ = v21
    return 994
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int32 = -1
        _ = fi22
        _ = v23
    }
    var v24 bool = true
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 uint16 = 65535
        v27 := GenS2{F0: 255, F1: -433, F2: ((true || (-2.25 <= (1.5 - (float32((0.0 / 1.0)) / (1.5 * -2.25))))) || false)}
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{42, 128, 150, 0}
    fmt.Println(v2)
    switch ("" + "foo bar") {
    case "":
        for fi29 := 0; fi29 < 19; fi29++ {
            v24 = v24
            v2 += (-(v2))
            if (-395 <= 0) {
                v1 = (v1 >> ((42) & 63))
            } else {
                var v30 uint32 = 4294967295
                if (v2 == v2) {
                    fmt.Println(v1)
                }
                v31 := GenS3{F0: 42, F1: 42, F2: "a"}
                _ = v30
                _ = v31
            }
            _ = fi29
        }
    case "hello":
        fmt.Println(v24, v24, v2)
    default:
        v1 -= v1
    }
    v2 -= (v2 - (v2 * -26))
    v32 := GenS4{F0: (v2 ^ -9), F1: 1}
    var v33 uint = 266
    v33 = (^(v33))
    for fi34 := 0; fi34 < 5; fi34++ {
        for fi35 := 0; fi35 < 14; fi35++ {
            fmt.Println(v2, v1, v2)
            switch "" {
            case "a":
                fmt.Println(v1, v1)
            }
            if v24 {
                v1 = 1
                var v36 int8 = v2
                _ = v36
            } else {
                var v37 uint8 = 127
                v1 = v1
                _ = v37
            }
            v24 = (!false)
            _ = fi35
        }
        v38 := v3[0]
        var v39 uint32 = 4294967295
        _ = fi34
        _ = v38
        _ = v39
    }
    fmt.Println(v1, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
