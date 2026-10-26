package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }

func main() {
    var v1 float32 = ((1e10 + 0.0) - (1.5 / (((((0.1 + (-2.25 - 1e10)) / -2.25) - (1.5 * ((((float32(((1.5 / 3.14159) * -2.25)) + -2.25) / 1e10) * 1.5) / 0.1))) - (3.14159 * (((((1e10 + 0.0) + (1.5 * 0.1)) * (1e10 - (1e10 * (0.1 - ((3.14159 - float32(-2.25)) * float32(float64(0.1))))))) + (float32((3.14159 / float64(1e10))) - 1.5)) / ((3.14159 - 1e10) / (1e10 * 3.14159))))) / (3.14159 * 0.1))))
    var v2 uint32 = 2147483647
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    v6 := []int{127}
    if true {
        if false {
            if len(v6) > 0 {
                v6[0] = 32767
            }
        } else {
            v1 -= 0.1
            v7 := GenS1{F0: 0, F1: 596}
            _ = v7
        }
        if false {
            v1 = v1
            fmt.Println(v2, v2)
            v1 -= 1.5
        }
    } else {
        switch ("foo bar" + "hello") {
        case "foo bar":
            if true {
                for fi8 := 0; fi8 < 20; fi8++ {
                    if (42 < int32(v2)) {
                        v2 -= v2
                        fmt.Println(fi8, fi8)
                    } else {
                        if false {
                            switch 2147483647 {
                            case -981:
                                v9 := GenS1{F0: 0, F1: 1}
                                _ = v9
                            default:
                                for fi10 := 0; fi10 < 1; fi10++ {
                                    switch fi10 {
                                    case 873:
                                        fmt.Println(fi8, fi10, fi8)
                                    }
                                    _ = fi10
                                }
                            }
                            _ = len(v3)
                        } else {
                            var v11 uint8 = 127
                            var v12 uint16 = 255
                            v13 := GenS2{F0: 115, F1: (1 < 42)}
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: 895, F1: 801}
                        var v15 bool = (586 >= 2)
                        _ = v14
                        _ = v15
                    }
                    _ = fi8
                }
                v2 = v2
                var v16 int32 = 0
                _ = v16
            }
        case "":
            for fi17 := 0; fi17 < 10; fi17++ {
                if false {
                    switch fi17 {
                    case 42:
                        for fi18 := 0; fi18 < 4; fi18++ {
                            if false {
                                v19 := GenS1{F0: 255, F1: -2147483648}
                                _ = v19
                            } else {
                                var v20 int16 = 32767
                                _ = v20
                            }
                            var v21 int = -132
                            _ = fi18
                            _ = v21
                        }
                    case 969:
                        var v22 int16 = -1
                        _ = v22
                    }
                    var v23 uint64 = 18446744073709551615
                    v23 -= (v23 | v23)
                    _ = v23
                }
                _ = fi17
            }
        case "a":
            _ = len(v3)
        default:
            v2 = (v2 + v2)
        }
        switch (("hello" + ("a" + (("" + "hello") + "a"))) + "\n\t") {
        case "":
            for fi24 := 0; fi24 < 10; fi24++ {
                v1 = v1
                v25 := GenS2{F0: 18446744073709551615, F1: (!true)}
                _ = fi24
                _ = v25
            }
        case "a":
            var v26 float64 = (-2.25 / (3.14159 + 1.5))
            _ = v26
        case "\n\t":
            v2 = 65535
        default:
            if false {
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 int = (fi27 % ((fi27) | 1))
                    _ = fi27
                    _ = v28
                }
                v29 := GenS3{F0: 1, F1: 626, F2: v2}
                _ = v29
            }
        }
    }
    var v30 uint32 = v2
    v1 += v1
    var v31 float64 = 0.1
    v2 += (v30 | (v30 | v2))
    v6 = append(v6, 32768)
    var v32 int8 = -1
    var v33 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
