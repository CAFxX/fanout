package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int = 32768
    var v2 int8 = 42
    var v3 int32 = 42
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    delete(v5, 4294967295)
    v1 -= (v1 ^ v1)
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            switch v1 {
            case -354:
                var v7 float64 = 1.5
                _ = v7
            case 459:
                v8 := GenS1{F0: 65535, F1: 18446744073709551615}
                _ = v8
            case 0:
                var v9 uint = 4294967295
                _ = v9
            default:
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 string = "hello"
                    var v12 uint16 = 255
                    v13 := GenS2{F0: v11, F1: (v12 / ((255) | 1))}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            switch 128 {
            case -2147483648:
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS2{F0: "\n\t", F1: 2}
                    if true {
                        v16 := GenS1{F0: 255, F1: 65535}
                        v2 += v2
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
            case 127:
                switch (v1 << ((18446744073709551615) & 63)) {
                case -2147483648:
                    v2 = 42
                }
            case 94:
                if (446 != 255) {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v1 -= fi6
                        var v18 uint = 1
                        v19 := GenS3{F0: 0, F1: -1, F2: 18446744073709551615}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if true {
                        var v20 bool = true
                        v21 := v5[65535]
                        _ = v20
                        _ = v21
                    } else {
                        delete(v4, "\n\t")
                        var v22 uint64 = 701
                        _ = v22
                    }
                    v23 := GenS3{F0: (v3 * v3), F1: 42, F2: 139}
                    _ = v23
                }
            }
            v2 = v2
        } else {
            if true {
                var v24 float32 = (((0.0 + 0.0) / 0.1) + 0.1)
                _ = v24
            } else {
                for fi25 := 0; fi25 < 3; fi25++ {
                    v26 := GenS4{F0: 744, F1: 2147483648}
                    v1 = (fi25 / (((^(fi6))) | 1))
                    v2 += -1
                    if ((((0.0 - 3.14159) - 1e10) - 1.5) >= 0.0) {
                        v2 = (v2 | 127)
                        switch ((("a" + "") + "") + "\n\t") {
                        case "\n\t":
                            v27 := []int{}
                            _ = v27
                        case "foo bar":
                            fmt.Println(v3, v3, v3)
                        case "":
                            var v28 uint64 = 18446744073709551615
                            _ = v28
                        default:
                            var v29 int32 = (^(v3))
                            _ = v29
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                var v30 string = ""
                v31 := GenS2{F0: ("" + v30), F1: 42}
                _ = v30
                _ = v31
            }
        }
        _ = fi6
    }
    var v32 string = ("" + ("\n\t" + "foo bar"))
    var v33 uint32 = 2147483647
    if false {
        var v34 string = v32
        _ = v34
    } else {
        fmt.Println(v32)
    }
    v2 = -74
    v35 := GenS1{F0: 32767, F1: 112}
    fmt.Println(v1)
    if (v33 != v33) {
        v3 += (v3 ^ v3)
        var v36 uint32 = v33
        _ = v36
    }
    v37 := []int{(^(v1)), (v1 % ((v1) | 1)), v1}
    var v38 int64 = -16
    var v39 bool = false
    for fi40 := 0; fi40 < 18; fi40++ {
        v32 = ("a" + v32)
        _ = fi40
    }
    var v41 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
