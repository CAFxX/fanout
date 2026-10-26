package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int16; F1 string }

func main() {
    v1 := []int{127, -135, 184, 430}
    v2 := []string{("\n\t" + ""), ""}
    v3 := []int{-214, 9223372036854775807, 1}
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 = v4
        if len(v3) > 0 {
            v3[0] = (fi5 * (fi5 ^ (fi5 << ((42) & 63))))
        }
        _ = fi5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        if (42 >= -10) {
            v4 = v4
            v4 = (("foo bar" + v4) + v4)
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS1{F0: 1, F1: 1}
                fmt.Println(fi6, fi7, v4)
                switch v4 {
                case "a":
                    v4 = v4
                }
                var v9 bool = ((v4 + "a") != "a")
                fmt.Println(fi6, fi6)
                _ = fi7
                _ = v8
                _ = v9
            }
            v4 = ((v4 + ((v4 + ("hello" + v4)) + (v4 + "hello"))) + v4)
        }
        var v10 string = (v4 + v4)
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v10, v10)
            if len(v3) > 0 {
                v3[0] = (fi11 / (((fi6 ^ (fi11 + -2147483648))) | 1))
            }
            _ = fi11
        }
        _ = fi6
        _ = v10
    }
    v4 = v4
    v12 := GenS2{F0: 42, F1: 127, F2: 0}
    if (true || true) {
        v4 = v4
        v4 = "\n\t"
    }
    v13 := GenS3{F0: 75, F1: 411}
    fmt.Println(v4)
    var v14 uint16 = 2
    var v15 int16 = 1
    switch 128 {
    case -2147483648:
        var v16 float32 = float32(0.0)
        _ = v16
    case 874:
        var v17 int32 = -2147483648
        _ = v17
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            v14 = v14
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi18
        }
    }
    var v19 float32 = float32((float64(0.1) * 1.5))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 128 {
    case -520:
        var v20 uint32 = 4294967295
        _ = v20
    default:
        v14 += v14
    }
    if false {
        for fi21 := 0; fi21 < 7; fi21++ {
            switch fi21 {
            case -40:
                var v22 float64 = 0.1
                _ = v22
            case -111:
                switch ("\n\t" + v4) {
                case "hello":
                    _ = len(v1)
                case "":
                    var v23 int32 = 2147483647
                    _ = v23
                case "foo bar":
                    switch fi21 {
                    case 4294967295:
                        v24 := GenS4{F0: 42, F1: true}
                        _ = v24
                    case -513:
                        var v25 uint64 = 2
                        _ = v25
                    case 774:
                        switch v4 {
                        case "a":
                            for fi26 := 0; fi26 < 4; fi26++ {
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                                if len(v2) > 0 {
                                    v2[0] = ((v4 + v4) + "")
                                }
                                var v27 float64 = 0.0
                                _ = fi26
                                _ = v27
                            }
                        case "foo bar":
                            for fi28 := 0; fi28 < 4; fi28++ {
                                if (-1 < 845) {
                                    v14 += (v14 / ((v14) | 1))
                                    if ((uint32(v14) < 378) && (!true)) {
                                        var v29 int16 = (-(v15))
                                        var v30 int64 = -631
                                        _ = v29
                                        _ = v30
                                    }
                                    v4 = ("foo bar" + v4)
                                    v4 = ("foo bar" + "")
                                }
                                v1 = append(v1, fi21)
                                _ = fi28
                            }
                        }
                    }
                default:
                    v4 = "hello"
                }
            default:
                var v31 uint16 = v14
                _ = v31
            }
            _ = len(v2)
            _ = fi21
        }
        switch v4 {
        case "\n\t":
            var v32 uint64 = 284
            _ = v32
        }
        v33 := GenS5{F0: v15, F1: ("" + (v4 + (("foo bar" + ("foo bar" + "")) + v4)))}
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
