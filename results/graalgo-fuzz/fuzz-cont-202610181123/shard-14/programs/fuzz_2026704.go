package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint; F2 int8 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 int16) uint32 {
    p0 += (p0 | (-(p0)))
    v4 := GenS1{F0: 227, F1: 2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 string = ""
    v2 := []uint32{}
    var v3 float32 = 1e10
    v2 = append(v2, 6)
    v2 = append(v2, 2147483647)
    v5 := GenS1{F0: 4294967295, F1: 944}
    var v6 bool = true
    v7 := GenS2{F0: -1}
    v8 := GenS3{F0: 3, F1: 42}
    switch 127 {
    case 32768:
        v1 = v1
    case 4294967295:
        var v9 int = 32768
        _ = v9
    case -96:
        v10 := GenS1{F0: 42, F1: genFn0(714)}
        _ = v10
    default:
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    v3 -= 0.1
    for fi11 := 0; fi11 < 2; fi11++ {
        if v6 {
            v3 = (3.14159 + v3)
        }
        v12 := GenS1{F0: 129, F1: 978}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi11
        _ = v12
    }
    v3 += ((v3 * v3) * ((v3 - v3) / 1.0))
    if v6 {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v13 := GenS4{F0: 0, F1: 255, F2: 42}
        _ = v13
    } else {
        if (!false) {
            fmt.Println(v6, v6, v1)
            v3 = ((float32(float64(v3)) / v3) + v3)
            var v14 bool = v6
            _ = v14
        } else {
            switch (v1 + "") {
            case "a":
                v15 := GenS3{F0: 2, F1: -1}
                _ = v15
            default:
                switch -800 {
                case -1:
                    var v16 uint32 = 544
                    _ = v16
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 uint8 = 128
                        v2 = append(v2, 202)
                        _ = fi17
                        _ = v18
                    }
                case 995:
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v2 = append(v2, 65535)
                        for fi20 := 0; fi20 < 11; fi20++ {
                            switch 42 {
                            case -500:
                                var v21 int16 = -32768
                                _ = v21
                            }
                            if ((false || false) || false) {
                                var v22 uint64 = 4294967295
                                _ = v22
                            }
                            fmt.Println(v1, v6)
                            _ = fi20
                        }
                        var v23 bool = (!(812 > 829))
                        _ = fi19
                        _ = v23
                    }
                default:
                    v3 -= float32((0.0 - 3.14159))
                }
            }
            var v24 int32 = 275
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS2{F0: -1}
                for fi27 := 0; fi27 < 11; fi27++ {
                    fmt.Println(fi25, v6)
                    _ = fi27
                }
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    }
    if (2147483647 != 310) {
        if (475 <= 1000) {
            v28 := GenS4{F0: 731, F1: 481, F2: 100}
            for fi29 := 0; fi29 < 14; fi29++ {
                v6 = v6
                _ = fi29
            }
            _ = v28
        } else {
            var v30 bool = (-492 <= 134)
            _ = v30
        }
        fmt.Println(v1)
        _ = len(v2)
    } else {
        if (32767 == 1) {
            var v31 int8 = 0
            var v32 uint8 = 2
            _ = v31
            _ = v32
        } else {
            var v33 string = (("\n\t" + "a") + v1)
            v6 = (42 >= 71)
            _ = v33
        }
    }
    if (float64(v3) != 1e10) {
        var v34 int16 = -32768
        _ = v34
    } else {
        v35 := GenS5{F0: 882, F1: ""}
        _ = v35
    }
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = 42
    }
    v1 = v1
    v2 = append(v2, genFn0(577))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
