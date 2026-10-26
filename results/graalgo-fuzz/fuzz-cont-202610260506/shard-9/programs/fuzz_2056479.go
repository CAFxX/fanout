package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint32 = 302
    var v2 int16 = -426
    v3 := []uint32{v1, 0, (v1 << ((732) & 63))}
    var v4 uint = 65535
    v5 := make(map[uint32]int)
    v6 := make(map[string]bool)
    v2 += 1
    v2 -= (v2 | (^(v2)))
    var v7 float32 = 3.14159
    v8 := GenS1{F0: 0, F1: v4, F2: v2}
    var v9 int32 = -505
    if false {
        var v10 uint8 = 127
        v11 := GenS1{F0: 42, F1: 0, F2: (-(v2))}
        var v12 uint16 = 1
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v9)
    }
    v13 := GenS1{F0: 2, F1: v4, F2: v2}
    v14 := GenS1{F0: uint8(v4), F1: 255, F2: v2}
    fmt.Println(v2, v7, v2)
    v7 -= (float32(-2.25) + (v7 - float32((((0.0 / (((float64(v7) * (1e10 + (3.14159 / 0.1))) - 0.0) * (3.14159 - (3.14159 - 1e10)))) + 1e10) / 0.1))))
    v15 := GenS1{F0: 214, F1: (v4 & (v4 << (((v4 + (^(v4)))) & 63))), F2: (v2 & v2)}
    var v16 int = 127
    switch "" {
    case "a":
        if true {
            for fi17 := 0; fi17 < 6; fi17++ {
                switch "\n\t" {
                case "foo bar":
                    var v18 uint8 = 88
                    _ = v18
                }
                _ = fi17
            }
        }
    case "hello":
        fmt.Println(v7, v9, v4)
    default:
        v19 := GenS1{F0: 255, F1: (v4 & v4), F2: (v2 >> (((v4 | (v4 / ((v4) | 1)))) & 63))}
        _ = v19
    }
    var v20 float32 = 0.0
    _ = len(v5)
    if true {
        v21 := GenS1{F0: 128, F1: v4, F2: (v2 + -837)}
        switch v16 {
        case -137:
            v1 -= (^(v1))
        default:
            v22 := GenS1{F0: 13, F1: (v4 / (((v4 + v4)) | 1)), F2: v2}
            _ = v22
        }
        _ = v21
    } else {
        if (637 != 812) {
            switch (("a" + "\n\t") + "\n\t") {
            case "\n\t":
                fmt.Println(v9)
            default:
                var v23 uint8 = 36
                _ = v23
            }
            if false {
                var v24 uint32 = v1
                _ = len(v3)
                _ = v24
            } else {
                if (0.0 >= float32(0.0)) {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            if true {
                                var v26 float64 = 3.14159
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    fmt.Println(v2, v2)
                                    _ = fi27
                                }
                                _ = v26
                            }
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v28 int64 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi25
                        _ = v28
                    }
                } else {
                    var v29 uint16 = 1
                    _ = v29
                }
                var v30 uint = 4294967295
                _ = v30
            }
        } else {
            if ("" == "foo bar") {
                if (v2 < v2) {
                    fmt.Println(v20)
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    delete(v5, (^(v1)))
                    v2 = (v2 - -1)
                }
            }
            v31 := GenS2{F0: v16}
            _ = v31
        }
        var v32 int16 = 32767
        delete(v6, "a")
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
