package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0(p0 uint8) uint {
    var v7 uint8 = 0
    var v8 uint = 994
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-566, -29}
    v3 := []int{128, 9223372036854775807, 32768, 128}
    var v4 float32 = 0.1
    var v5 string = "hello"
    v6 := make(map[uint32]int)
    v9 := GenS1{F0: 2}
    switch v5 {
    case "":
        switch ((((("" + v5) + "a") + "hello") + v5) + (("a" + (v5 + v5)) + ((v5 + v5) + v5))) {
        case "a":
            var v10 uint = 4294967295
            _ = v10
        default:
            if true {
                switch (v5 + v5) {
                case "\n\t":
                    v11 := GenS2{F0: -423}
                    _ = v11
                case "":
                    var v12 bool = false
                    _ = v12
                case "a":
                    fmt.Println(v4)
                default:
                    v13 := GenS3{F0: 1, F1: v5}
                    _ = v13
                }
                if ((((0.0 + (float32(((1e10 * 0.1) * 0.0)) + v4)) + v4) > (float32(0.1) - (0.0 * v4))) && false) {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        switch v5 {
                        case "hello":
                            v15 := GenS2{F0: -824}
                            _ = v15
                        case "":
                            if true {
                                v16 := GenS4{F0: false}
                                _ = v16
                            }
                        case "foo bar":
                            for fi17 := 0; fi17 < 16; fi17++ {
                                v4 -= (0.1 - v4)
                                var v18 int16 = -1
                                v4 += ((v4 + 0.0) - (v4 + (v4 + 3.14159)))
                                _ = fi17
                                _ = v18
                            }
                        default:
                            v19 := GenS5{F0: -32768, F1: 1}
                            _ = v19
                        }
                        _ = fi14
                    }
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) > 0 {
                            v2[0] = 32767
                        }
                    } else {
                        var v20 uint32 = 1
                        var v21 int64 = -14
                        for fi22 := 0; fi22 < 1; fi22++ {
                            v23 := GenS3{F0: 41, F1: v5}
                            _ = fi22
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                    }
                } else {
                    var v24 bool = false
                    v4 = 0.0
                    for fi25 := 0; fi25 < 11; fi25++ {
                        fmt.Println(fi25, fi25)
                        fmt.Println(v4)
                        _ = fi25
                    }
                    _ = v24
                }
                var v26 bool = true
                _ = v26
            }
        }
    case "a":
        v6[0] = 32768
    default:
        v27 := GenS1{F0: 1}
        _ = v27
    }
    var v28 uint32 = 0
    for fi29 := 0; fi29 < 4; fi29++ {
        if true {
            delete(v6, v28)
        }
        fmt.Println(v28, v5, v5)
        _ = fi29
    }
    v28 = v28
    switch 542 {
    case 42:
        v28 += (^(v28))
    }
    var v30 int16 = -1
    v5 = v5
    var v31 int64 = 0
    var v32 float32 = float32(0.0)
    v33 := GenS2{F0: (^(v31))}
    v5 = (v5 + v5)
    switch v5 {
    case "":
        if ((float64(float32(1e10)) + 1.5) < 1e10) {
            v34 := v6[422]
            v34 = 1
            var v35 uint16 = 440
            _ = v34
            _ = v35
        }
    case "\n\t":
        v36 := v1[(v28 >> ((768) & 63))]
        _ = v36
    case "a":
        v31 += v31
    }
    v37 := GenS6{F0: (v28 + v28), F1: true, F2: (v28 - 619)}
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
