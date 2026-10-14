package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 int16; F2 int }

func main() {
    v1 := []int{9223372036854775807, -453, 9223372036854775807}
    var v2 float32 = (((0.0 + (1e10 + ((3.14159 + 1e10) / -2.25))) * (1e10 / 1.0)) / -2.25)
    var v3 uint = 18446744073709551615
    v4 := []int{}
    var v5 float32 = (0.0 / (1e10 + float32(0.1)))
    var v6 uint16 = 65535
    if true {
        v5 += v2
        v2 = ((v2 - v2) - (v5 * (3.14159 - v2)))
    } else {
        switch "a" {
        case "\n\t":
            v3 = v3
        case "":
            var v7 uint16 = 0
            _ = v7
        case "a":
            if true {
                v8 := GenS1{F0: 0, F1: 594, F2: false}
                for fi9 := 0; fi9 < 5; fi9++ {
                    if (!((1 > 1) || false)) {
                        var v10 uint16 = (v6 % (((v6 ^ (^(v6)))) | 1))
                        var v11 uint = 891
                        _ = v10
                        _ = v11
                    } else {
                        var v12 uint32 = 42
                        var v13 int = (fi9 - (fi9 >> (((^(v3))) & 63)))
                        if false {
                            fmt.Println(v12, v3)
                            v5 -= (float32((1.5 + 0.1)) - float32((0.0 - 1.5)))
                            if (v13 <= fi9) {
                                _ = len(v1)
                            }
                        }
                        _ = v12
                        _ = v13
                    }
                    if true {
                        switch "a" {
                        case "hello":
                            v2 = v2
                        }
                        v14 := GenS2{F0: (v6 != (v6 - (v6 + v6)))}
                        var v15 float32 = v5
                        _ = v14
                        _ = v15
                    } else {
                        v16 := GenS2{F0: false}
                        for fi17 := 0; fi17 < 15; fi17++ {
                            if true {
                                if false {
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (18446744073709551615 < 786) {
                                                fmt.Println(v6)
                                                switch "hello" {
                                                case "foo bar":
                                                    _ = len(v4)
                                                case "hello":
                                                    if false {
                                                        v1 = append(v1, (fi17 >> (((v3 & 0)) & 63)))
                                                        if ((^(fi17)) < fi17) {
                                                            v18 := GenS3{F0: fi9, F1: 2}
                                                            switch -1 {
                                                            case 127:
                                                                if len(v1) > 0 {
                                                                    v1[0] = (fi9 / ((0) | 1))
                                                                }
                                                            default:
                                                                v19 := GenS1{F0: 0, F1: -1, F2: true}
                                                                _ = v19
                                                            }
                                                            v4 = append(v4, fi9)
                                                            var v20 float64 = 0.0
                                                            v5 -= float32(v20)
                                                            _ = v18
                                                            _ = v20
                                                        } else {
                                                            var v21 int8 = 104
                                                            if (42 >= 255) {
                                                                v5 = float32((3.14159 + 0.0))
                                                                var v22 int16 = -1
                                                                v21 = -128
                                                                var v23 float32 = 3.14159
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            _ = v21
                                                        }
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = fi17
                                                        }
                                                        v2 += ((v5 + (3.14159 + v2)) - 0.1)
                                                        switch (("a" + "") + "") {
                                                        case "foo bar":
                                                            var v24 int32 = -104
                                                            _ = v24
                                                        case "\n\t":
                                                            var v25 float32 = float32(1e10)
                                                            _ = v25
                                                        case "":
                                                            var v26 float64 = 1e10
                                                            _ = v26
                                                        }
                                                    }
                                                }
                                            }
                                        case "hello":
                                            if false {
                                                v27 := GenS4{F0: 42, F1: 0}
                                                var v28 uint8 = 1
                                                switch "a" {
                                                case "a":
                                                    if (("hello" + ("" + "")) == "") {
                                                        var v29 int = (fi9 * fi9)
                                                        _ = v29
                                                    }
                                                case "foo bar":
                                                    switch fi9 {
                                                    case 42:
                                                        v30 := GenS2{F0: true}
                                                        _ = v30
                                                    case -2147483648:
                                                        v31 := GenS5{F0: v3}
                                                        _ = v31
                                                    case 32768:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v5, fi17)
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                            var v32 int = fi9
                                            v33 := GenS6{F0: 2147483647, F1: -207, F2: fi17}
                                            _ = v32
                                            _ = v33
                                        default:
                                            if true {
                                                if false {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    fmt.Println(v3, v5, v2)
                                                    fmt.Println(fi17, v6)
                                                }
                                                var v34 bool = false
                                                _ = v34
                                            }
                                        }
                                        v3 = (v3 + v3)
                                        fmt.Println(fi9, fi9, fi17)
                                    } else {
                                        v3 -= 18446744073709551615
                                        switch fi9 {
                                        case 2147483647:
                                            var v35 uint8 = 0
                                            _ = v35
                                        case 1:
                                            switch "hello" {
                                            case "a":
                                                if len(v1) > 0 {
                                                    v1[0] = (fi17 ^ fi17)
                                                }
                                            case "hello":
                                                v2 = v2
                                            case "\n\t":
                                                fmt.Println(v6, v6, fi9)
                                            default:
                                                var v36 uint16 = (v6 ^ v6)
                                                _ = v36
                                            }
                                        }
                                    }
                                    fmt.Println(v2)
                                }
                            }
                            _ = fi17
                        }
                        v37 := GenS6{F0: -102, F1: 42, F2: 42}
                        _ = v16
                        _ = v37
                    }
                    _ = fi9
                }
                _ = v8
            }
        default:
            for fi38 := 0; fi38 < 6; fi38++ {
                fmt.Println(fi38, v2)
                fmt.Println(v5, v3, v2)
                v6 -= v6
                _ = fi38
            }
        }
        v6 = (v6 | (v6 & 832))
        for fi39 := 0; fi39 < 1; fi39++ {
            v6 -= 2
            v6 = (v6 * (^(v6)))
            var v40 uint32 = 982
            _ = fi39
            _ = v40
        }
    }
    var v41 uint16 = 255
    for fi42 := 0; fi42 < 2; fi42++ {
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := GenS6{F0: 2147483647, F1: -660, F2: 127}
            _ = fi43
            _ = v44
        }
        v5 += v2
        _ = fi42
    }
    v6 += (v41 + (v6 ^ (v6 % (((^(v6))) | 1))))
    var v45 uint32 = 89
    v3 -= v3
    if (18446744073709551615 >= 0) {
        v3 = 42
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v46 uint64 = 2
        _ = v46
    }
    fmt.Println(v3, v2)
    if len(v4) > 0 {
        v4[0] = -979
    }
    if ((v45 - v45) != (v45 - 65535)) {
        v5 -= v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v5
    _ = v6
}
