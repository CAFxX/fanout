package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint32 = 2
    var v3 float64 = (1.5 / (((float64(-2.25) / -2.25) * -2.25) / 1.0))
    var v4 uint = 18446744073709551615
    v5 := []string{}
    if (!(true && (true || (!true)))) {
        v6 := GenS1{F0: 1}
        _ = v6
    } else {
        v7 := GenS2{F0: 284, F1: 127}
        v1 += v1
        for fi8 := 0; fi8 < 17; fi8++ {
            v1 += 1e10
            switch "a" {
            case "":
                if false {
                    _ = len(v5)
                    fmt.Println(v4, v4, v3)
                    fmt.Println(v1, v4)
                } else {
                    var v9 int = fi8
                    _ = v9
                }
            case "hello":
                v2 = v2
            case "foo bar":
                var v10 bool = (274 >= 2)
                _ = v10
            default:
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint32 = (v2 & (^(v2)))
                    var v13 int32 = 42
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi14
            }
            _ = fi8
        }
        _ = v7
    }
    fmt.Println(v3, v2)
    var v15 uint64 = 0
    fmt.Println(v4, v3, v15)
    if true {
        switch 42 {
        case 32768:
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(v3)
                _ = fi16
            }
        default:
            if true {
                var v17 int64 = -9223372036854775808
                if (v2 > (v2 << (((v4 << (((^(v4))) & 63))) & 63))) {
                    v5 = append(v5, "a")
                } else {
                    switch 2147483647 {
                    case 127:
                        fmt.Println(v3, v1, v4)
                    case 42:
                        var v18 float32 = v1
                        _ = v18
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v1 += v1
                    if (255 == 1) {
                        if (!(((4294967295 < 32768) || false) && (0 != 42))) {
                            switch "a" {
                            case "":
                                v2 = (v2 + (v2 >> ((391) & 63)))
                            case "hello":
                                v3 = v3
                            case "\n\t":
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    fmt.Println(fi19, v2)
                                    var v20 uint64 = v15
                                    _ = fi19
                                    _ = v20
                                }
                            default:
                                var v21 string = "foo bar"
                                _ = v21
                            }
                        }
                    } else {
                        var v22 int32 = int32(v2)
                        var v23 uint = v4
                        _ = v22
                        _ = v23
                    }
                }
                v24 := GenS2{F0: v15, F1: 42}
                _ = v17
                _ = v24
            } else {
                _ = len(v5)
                v5 = append(v5, "a")
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        var v25 int16 = 707
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = "a"
        }
    case "":
        fmt.Println(v2, v2, v3)
    default:
        v15 -= (^(v15))
    }
    var v26 int64 = -207
    var v27 string = "hello"
    v28 := GenS3{F0: 255, F1: 980}
    var v29 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
