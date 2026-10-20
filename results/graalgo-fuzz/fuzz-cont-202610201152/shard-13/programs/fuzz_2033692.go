package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []int{42, 2147483647, 128}
    var v2 float64 = ((0.0 - 0.0) / float64(1.5))
    var v3 uint16 = 779
    var v4 bool = true
    var v5 string = "foo bar"
    switch -939 {
    case -201:
        v4 = ((192 == 42) || ((v5 + v5) <= v5))
    case -1:
        var v6 uint64 = 0
        _ = v6
    case 2147483647:
        v3 = v3
    default:
        switch 32767 {
        case 326:
            var v7 int32 = 0
            _ = v7
        case 1:
            if len(v1) > 0 {
                v1[0] = 127
            }
        case -1:
            var v8 uint32 = 65535
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            if v4 {
                fmt.Println(v5, v3)
            }
            _ = fi10
        }
        fmt.Println(v4, v4)
        _ = fi9
    }
    v11 := GenS1{F0: v3}
    v12 := GenS1{F0: v3}
    var v13 uint = 18446744073709551615
    var v14 int16 = 273
    v4 = v4
    if false {
        if ((((0.1 + 0.1) - ((((1.5 * 0.0) * (1.5 * (0.1 + (1.5 / -2.25)))) * 1e10) / 1.0)) - 1e10) <= (1.5 - 0.1)) {
            fmt.Println(v2)
            v2 = v2
            var v15 int8 = 127
            _ = v15
        }
    } else {
        if false {
            v16 := GenS1{F0: (v3 % (((v3 + v3)) | 1))}
            for fi17 := 0; fi17 < 7; fi17++ {
                if v4 {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if v4 {
                            v14 += v14
                        }
                        v19 := GenS1{F0: v3}
                        v14 = -32768
                        _ = fi18
                        _ = v19
                    }
                } else {
                    switch fi17 {
                    case -604:
                        if (0 < 1) {
                            var v20 uint16 = v3
                            _ = v20
                        }
                    case 624:
                        v1 = append(v1, fi17)
                    case 127:
                        v4 = v4
                    }
                    v2 = ((v2 - 1.5) + 3.14159)
                }
                _ = fi17
            }
            v21 := GenS2{F0: v14, F1: 100, F2: 369}
            _ = v16
            _ = v21
        } else {
            fmt.Println(v2, v3)
            var v22 int = 9223372036854775807
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: 165}
                var v25 float64 = (1e10 - v2)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v22
        }
        var v26 bool = v4
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        if true {
            switch v5 {
            case "":
                v28 := GenS3{F0: v13}
                _ = v28
            case "\n\t":
                _ = len(v1)
            default:
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 float32 = (0.1 / 1.0)
                    v31 := GenS4{F0: 4294967295, F1: 42}
                    if (true || (!false)) {
                        if v4 {
                            v2 = 3.14159
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        var v32 int8 = -128
                        v13 -= 430
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            fmt.Println(v2)
        }
        _ = fi27
    }
    var v33 bool = (537 != 42)
    var v34 int8 = 40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
