package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint }

func main() {
    var v1 uint8 = 41
    v2 := []string{"hello"}
    var v3 float64 = 1.5
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    var v6 int16 = 1000
    var v7 float64 = v3
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: 1, F1: (v6 / ((32767) | 1))}
        v2 = append(v2, "foo bar")
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int = (fi11 + -72)
        var v13 int32 = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v3)
    v6 += 42
    switch 557 {
    case 128:
        var v14 int16 = v6
        _ = v14
    case 599:
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: 18446744073709551615, F1: v6}
            v17 := GenS1{F0: -81}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v1, v7)
    for fi18 := 0; fi18 < 4; fi18++ {
        v1 += 128
        _ = fi18
    }
    switch "" {
    case "hello":
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                switch fi20 {
                case 128:
                    var v21 string = ("foo bar" + "hello")
                    _ = v21
                case 9223372036854775807:
                    v1 = 42
                case -1:
                    if (!(!false)) {
                        fmt.Println(v6)
                        v22 := GenS1{F0: 100}
                        switch -6 {
                        case 651:
                            var v23 float32 = 1e10
                            _ = v23
                        case 32767:
                            v24 := GenS2{F0: 1, F1: (v6 / ((v6) | 1))}
                            _ = v24
                        }
                        if true {
                            var v25 float64 = v3
                            _ = v25
                        }
                        fmt.Println(v1, v6, v7)
                        _ = v22
                    }
                    var v26 uint16 = 80
                    _ = v26
                default:
                    var v27 int16 = v6
                    _ = v27
                }
                _ = fi20
            }
            v28 := v5[(fi19 | fi19)]
            _ = fi19
            _ = v28
        }
    default:
        v29 := GenS3{F0: v1, F1: 32767, F2: 255}
        _ = v29
    }
    if (255 <= 376) {
        var v30 int32 = 0
        var v31 uint64 = 42
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 12; fi32++ {
            v3 = v7
            v3 = v7
            fmt.Println(fi32, v1, v3)
            _ = fi32
        }
        v1 = v1
        var v33 uint = 586
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
