package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int; F1 int16; F2 int8 }
func genFn0() uint8 {
    var v6 string = "hello"
    fmt.Println(v6)
    v7 := []int{-868, 1, 32767}
    var v8 uint = 649
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 0, F1: -2147483648, F2: 261}
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool) string {
    p0 = (-116 >= 42)
    v11 := []int{2147483647}
    if p0 {
        fmt.Println(p0)
        switch 611 {
        case 2147483647:
            p0 = ((0.0 + 3.14159) <= 0.1)
        }
    } else {
        p0 = false
    }
    _ = p0
    _ = v11
    return ("" + (("hello" + (("hello" + "") + "hello")) + ("foo bar" + "foo bar")))
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 float32 = float32((((0.0 / (0.0 + (((-2.25 - 1.5) - 1.5) / 1e10))) + 1e10) - (-2.25 + float64(3.14159))))
    var v3 uint = 2
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    v4["\n\t"] = 123
    for fi12 := 0; fi12 < 2; fi12++ {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := []int{(fi13 | (fi12 + fi13)), (fi12 % ((fi13) | 1)), 9223372036854775807}
                switch "" {
                case "foo bar":
                    if (false && ((1 == 100) || true)) {
                        if true {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            var v15 bool = (!(-1 > 1))
                            switch fi12 {
                            case 32767:
                                v3 -= (v3 - (^(v3)))
                            case 9223372036854775807:
                                var v16 uint32 = 374
                                var v17 uint64 = 4294967295
                                _ = v16
                                _ = v17
                            }
                            _ = v15
                        } else {
                            if true {
                                if true {
                                    if true {
                                        switch v1 {
                                        case "":
                                            v18 := GenS2{F0: (true || true)}
                                            _ = v18
                                        }
                                        v2 -= v2
                                        if false {
                                            switch (fi12 << (((^(v3))) & 63)) {
                                            case 34:
                                                var v19 uint16 = 255
                                                _ = v19
                                            }
                                        }
                                    } else {
                                        var v20 float32 = float32(((float64(1.5) / 3.14159) / 1.5))
                                        _ = v20
                                    }
                                    var v21 float64 = ((0.0 / 1.0) - float64(v2))
                                    _ = v21
                                }
                            }
                        }
                        v3 += v3
                    }
                default:
                    var v22 uint32 = 65535
                    _ = v22
                }
                if (v3 < v3) {
                    v23 := v5[-1]
                    v2 -= v2
                    v24 := GenS3{F0: 34, F1: 240}
                    _ = v23
                    _ = v24
                }
                fmt.Println(v2, v1)
                v25 := GenS4{F0: 0, F1: 2147483647}
                _ = fi13
                _ = v14
                _ = v25
            }
            if false {
                v26 := GenS5{F0: fi12, F1: 42, F2: -1}
                _ = v26
            }
            v27 := []int{fi12, 0}
            _ = v27
        } else {
            v2 += v2
            var v28 float64 = (3.14159 / 1.0)
            _ = v28
        }
        var v29 int8 = 42
        fmt.Println(fi12, v3, v1)
        _ = fi12
        _ = v29
    }
    _ = len(v5)
    switch 32767 {
    case -9223372036854775808:
        var v30 int = -284
        _ = v30
    case -55:
        if false {
            var v31 int64 = -1
            var v32 string = ""
            v1 = v1
            _ = v31
            _ = v32
        }
    default:
        v1 = ""
    }
    var v33 uint64 = 1
    v33 += (v33 << ((v3) & 63))
    for fi34 := 0; fi34 < 1; fi34++ {
        switch (fi34 >> ((uint(v33)) & 63)) {
        case 0:
            v1 = v1
        case -277:
            v35 := GenS3{F0: -1, F1: 247}
            _ = v35
        case 520:
            switch fi34 {
            case 127:
                var v36 int32 = 42
                _ = v36
            default:
                if (!(1e10 < 1.5)) {
                    fmt.Println(v1)
                    for fi37 := 0; fi37 < 13; fi37++ {
                        var v38 uint64 = v33
                        _ = fi37
                        _ = v38
                    }
                }
            }
        default:
            var v39 uint8 = 128
            _ = v39
        }
        _ = fi34
    }
    if true {
        var v40 float32 = v2
        for fi41 := 0; fi41 < 13; fi41++ {
            v42 := []int{128}
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    var v43 int64 = -738
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v5
}
