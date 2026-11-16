package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]string)
    var v3 float64 = 3.14159
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    var v6 string = ""
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        if (true && true) {
            switch v6 {
            case "a":
                v3 = (v3 / ((-2.25 * (v3 - v3)) * float64(1.5)))
            }
            fmt.Println(v3, fi7)
        } else {
            var v8 int64 = -9223372036854775808
            v9 := GenS1{F0: (v8 & v8)}
            var v10 uint = 18446744073709551615
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v6)
            var v12 float32 = 1.5
            _ = fi11
            _ = v12
        }
        var v13 int16 = 42
        _ = fi7
        _ = v13
    }
    v3 += v3
    var v14 float32 = (((3.14159 + 1e10) / 1.0) * 1.5)
    var v15 float32 = (v14 + ((1.5 - (((v14 + (0.1 - (-2.25 / 1.5))) - ((v14 - v14) - (0.1 - (v14 * 0.0)))) - v14)) - v14))
    switch "foo bar" {
    case "\n\t":
        v6 = (v6 + (v6 + v6))
    }
    v15 = 1e10
    fmt.Println(v6, v15, v14)
    v6 = v6
    v16 := GenS1{F0: -1}
    var v17 uint = 0
    if (!(1 == -375)) {
        switch "a" {
        case "hello":
            v18 := GenS2{F0: (2 == 2)}
            _ = v18
        case "foo bar":
            var v19 string = v6
            _ = v19
        case "":
            v6 = v6
        default:
            if true {
                v20 := GenS3{F0: 2, F1: -102}
                _ = v20
            }
        }
        if false {
            switch -9223372036854775808 {
            case 127:
                v15 += float32(0.1)
            case 4294967295:
                v17 = v17
            case 1:
                var v21 int = -9223372036854775808
                _ = v21
            default:
                if ((v3 < 1e10) && true) {
                    switch 240 {
                    case 1:
                        v22 := GenS1{F0: -1}
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 int8 = -1
                        fmt.Println(v14, fi23)
                        _ = fi23
                        _ = v24
                    }
                    v17 = 442
                } else {
                    if (127 < 1) {
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 int8 = 1
                            for fi27 := 0; fi27 < 4; fi27++ {
                                var v28 float32 = -2.25
                                if (!true) {
                                    if true {
                                        if ((v14 / (0.1 + v28)) <= v28) {
                                            switch fi27 {
                                            case 1:
                                                var v29 float64 = (0.0 * ((1.5 / 1e10) / 3.14159))
                                                _ = v29
                                            }
                                        } else {
                                            v15 = float32((3.14159 - float64(v15)))
                                            v14 -= v14
                                        }
                                    }
                                }
                                delete(v5, v6)
                                _ = fi27
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        var v30 string = v6
                        v17 = 42
                        _ = v30
                    }
                    var v31 int8 = 100
                    _ = v31
                }
            }
            v15 = ((-2.25 * v15) / (v15 - 1.5))
        } else {
            delete(v5, "")
            v15 += v15
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v6 = ""
            _ = fi32
        }
    }
    switch 1 {
    case 128:
        _ = len(v1)
    case -1:
        var v33 uint16 = 705
        _ = v33
    default:
        if (!false) {
            var v34 int = 42
            fmt.Println(v6)
            switch "foo bar" {
            case "a":
                v3 = v3
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = v34
                }
            case "":
                switch v6 {
                case "":
                    var v35 int64 = 42
                    _ = v35
                case "foo bar":
                    var v36 int16 = 0
                    _ = v36
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
