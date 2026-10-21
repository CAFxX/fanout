package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 int, p2 int8) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p1, p0)
        switch (fi6 ^ -489) {
        case 2147483647:
            p0 = p0
        case -1:
            v7 := []int{p1, (p1 & (^(p1)))}
            _ = v7
        case 1:
            switch 32768 {
            case 128:
                if true {
                    v8 := []int{(p1 / (((p1 ^ fi6)) | 1)), fi6, p1}
                    for fi9 := 0; fi9 < 10; fi9++ {
                        switch "a" {
                        case "foo bar":
                            _ = len(v8)
                        }
                        v10 := GenS1{F0: -1, F1: 42, F2: 1}
                        var v11 float32 = -2.25
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    p2 += (p2 ^ 127)
                    _ = v8
                }
            case 32768:
                v12 := []int{(p1 & 573), (fi6 / ((-1) | 1)), 9223372036854775807, (p1 & (-(p1)))}
                _ = v12
            default:
                var v13 int8 = 80
                _ = v13
            }
        }
        _ = fi6
    }
    var v14 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p0
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("foo bar" + ("\n\t" + "foo bar"))}
    var v3 float32 = 0.1
    v4 := []uint32{0, 249, 2147483647}
    v5 := make(map[string]int)
    v3 = 0.0
    v4 = append(v4, 2)
    v3 = ((v3 * float32(1e10)) / v3)
    var v15 string = "\n\t"
    v15 = "a"
    fmt.Println(v3)
    if true {
        var v16 string = "a"
        _ = v16
    }
    fmt.Println(v15, v3)
    v17 := GenS1{F0: -359, F1: 4294967295, F2: 127}
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint16 = 0
        _ = fi18
        _ = v19
    }
    if true {
        switch (v15 + "a") {
        case "a":
            v20 := GenS1{F0: 1, F1: 1, F2: 42}
            _ = v20
        case "\n\t":
            v21 := GenS1{F0: 0, F1: 0, F2: 22}
            _ = v21
        default:
            v22 := GenS1{F0: 0, F1: 70, F2: 255}
            _ = v22
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS1{F0: 1, F1: 773, F2: 128}
                v15 = v15
                var v26 uint32 = 0
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v3, v3, v3)
            var v27 int64 = 365
            _ = fi23
            _ = v27
        }
    } else {
        var v28 uint32 = 1
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        switch -9223372036854775808 {
        case 0:
            var v30 uint8 = 0
            _ = v30
        default:
            var v31 int16 = 1000
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint = 1
            if (true || (genFn0(397, 42, 0) <= 65535)) {
                switch v15 {
                case "foo bar":
                    v34 := GenS2{F0: 83, F1: 18446744073709551615, F2: 66}
                    _ = v34
                default:
                    switch fi29 {
                    case -2147483648:
                        _ = len(v2)
                    case 1:
                        var v35 bool = (((float64((3.14159 + ((1e10 * -2.25) / v3))) * -2.25) == float64(3.14159)) || (2147483648 == 42))
                        _ = v35
                    case 42:
                        if ((-9223372036854775808 == -562) && ((-(fi29)) <= fi29)) {
                            _ = len(v2)
                        } else {
                            var v36 uint16 = 838
                            switch v15 {
                            case "a":
                                if (((((float32(0.1) * float32(0.0)) / v3) + v3) / v3) != 1e10) {
                                    switch 2147483647 {
                                    case 334:
                                        var v37 uint32 = 2147483647
                                        _ = v37
                                    default:
                                        var v38 float64 = 1e10
                                        _ = v38
                                    }
                                    v3 += v3
                                    if len(v4) > 0 {
                                        v4[0] = 2147483647
                                    }
                                }
                            case "\n\t":
                                v15 = ((v15 + "foo bar") + v15)
                            case "hello":
                                v39 := GenS3{F0: 31}
                                _ = v39
                            }
                            _ = v36
                        }
                    }
                }
            } else {
                if (false || false) {
                    fmt.Println(v33)
                    switch "" {
                    case "foo bar":
                        v40 := GenS4{F0: 32768}
                        _ = v40
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                var v41 uint16 = 560
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v41
            }
            _ = fi32
            _ = v33
        }
        v42 := GenS3{F0: 255}
        _ = fi29
        _ = v42
    }
    var v43 int8 = -122
    var v44 uint = 1
    var v45 int64 = -704
    v44 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
