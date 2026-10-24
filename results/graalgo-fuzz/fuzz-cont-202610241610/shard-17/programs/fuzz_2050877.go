package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint64 }

func main() {
    var v1 bool = ((true || ((false || (true || false)) && true)) || (false && false))
    var v2 float64 = (float64(-2.25) - 3.14159)
    var v3 uint = 18446744073709551615
    v4 := GenS1{F0: 121}
    fmt.Println(v2, v2, v3)
    v2 -= ((v2 - (float64(((0.0 - 1.5) / 0.1)) / (float64(1.5) * 0.1))) * v2)
    if (1000 >= -32768) {
        for fi5 := 0; fi5 < 1; fi5++ {
            if v1 {
                v2 += (1.5 * (3.14159 / v2))
                fmt.Println(v2, v1)
            } else {
                fmt.Println(v1)
                v1 = (v1 || v1)
            }
            v3 += 4294967295
            _ = fi5
        }
        if ((float64(0.1) / 1.0) <= v2) {
            if v1 {
                for fi6 := 0; fi6 < 16; fi6++ {
                    v3 -= v3
                    _ = fi6
                }
                v1 = (1 >= 75)
                fmt.Println(v1)
            }
            var v7 uint16 = 257
            _ = v7
        }
    } else {
        v2 -= float64((-2.25 + (0.1 - (1e10 * -2.25))))
        fmt.Println(v1, v3)
    }
    var v8 bool = (!(true || false))
    if v8 {
        fmt.Println(v8, v3, v2)
    } else {
        v1 = (!(!(274 > 620)))
        v9 := GenS1{F0: -128}
        _ = v9
    }
    fmt.Println(v1, v3)
    fmt.Println(v3)
    fmt.Println(v2, v2)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch (("foo bar" + "foo bar") + ("a" + "hello")) {
        case "\n\t":
            v11 := []int{fi10}
            _ = v11
        case "foo bar":
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int64 = 908
                v14 := GenS1{F0: 42}
                v1 = v8
                switch fi12 {
                case -85:
                    if (!false) {
                        if (1 >= 127) {
                            switch "\n\t" {
                            case "a":
                                var v15 string = ("hello" + ("" + "foo bar"))
                                _ = v15
                            case "\n\t":
                                v16 := GenS1{F0: -18}
                                _ = v16
                            }
                        } else {
                            fmt.Println(v2)
                            v17 := GenS2{F0: 42, F1: 564, F2: -521}
                            if (v13 < v13) {
                                fmt.Println(fi10, v8)
                                v18 := GenS1{F0: 100}
                                if v8 {
                                    v3 += v3
                                    v8 = v8
                                    v19 := GenS2{F0: 0, F1: 44, F2: v13}
                                    _ = v19
                                } else {
                                    var v20 float32 = (0.0 / 0.1)
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                v13 += -1
                            }
                            var v21 int16 = 0
                            switch fi10 {
                            case 32767:
                                var v22 int32 = -1
                                _ = v22
                            case 787:
                                var v23 float64 = 1.5
                                _ = v23
                            default:
                                var v24 float32 = float32(1.5)
                                _ = v24
                            }
                            _ = v17
                            _ = v21
                        }
                    } else {
                        var v25 uint = v3
                        switch "a" {
                        case "a":
                            fmt.Println(fi10, v2, v25)
                        case "\n\t":
                            v26 := []int{(-(fi12)), (fi12 * fi12), -1}
                            _ = v26
                        default:
                            var v27 uint8 = 2
                            _ = v27
                        }
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v1, v25, v25)
                        default:
                            var v28 uint32 = 0
                            _ = v28
                        }
                        _ = v25
                    }
                default:
                    var v29 string = ("foo bar" + "")
                    _ = v29
                }
                var v30 uint32 = 42
                _ = fi12
                _ = v13
                _ = v14
                _ = v30
            }
        case "a":
            v31 := []int{}
            _ = v31
        default:
            v32 := GenS1{F0: 127}
            _ = v32
        }
        var v33 float32 = 0.1
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := []int{(fi10 >> ((879) & 63)), 128, (fi34 / (((fi34 >> (((^(v3))) & 63))) | 1))}
            if len(v35) > 0 {
                v35[0] = fi34
            }
            _ = fi34
            _ = v35
        }
        _ = fi10
        _ = v33
    }
    switch -50 {
    case 32768:
        v36 := []int{-2147483648, 1}
        _ = v36
    case 1:
        var v37 float64 = (1e10 - v2)
        _ = v37
    case 272:
        v8 = false
    }
    v38 := []int{}
    switch "hello" {
    case "a":
        v39 := GenS3{F0: -2147483648}
        _ = v39
    case "\n\t":
        v8 = v8
    }
    var v40 int16 = 32767
    switch 128 {
    case 32768:
        if len(v38) >= 2 {
            v38 = v38[1:len(v38)]
        }
    case 556:
        var v41 uint8 = 2
        _ = v41
    case 771:
        v1 = (v1 && true)
    default:
        var v42 int64 = 0
        _ = v42
    }
    if false {
        for fi43 := 0; fi43 < 9; fi43++ {
            v44 := GenS4{F0: 4294967295, F1: -42}
            for fi45 := 0; fi45 < 3; fi45++ {
                v8 = v1
                _ = fi45
            }
            _ = fi43
            _ = v44
        }
        v46 := GenS5{F0: -186, F1: 100, F2: 42}
        _ = v46
    } else {
        fmt.Println(v1, v2, v40)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v8
}
