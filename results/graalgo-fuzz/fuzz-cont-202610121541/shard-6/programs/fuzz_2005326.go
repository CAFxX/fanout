package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    var v3 bool = true
    v4 := []uint32{65535, 2, 0}
    var v5 float64 = (float64(1e10) - (float64((1.5 - (-2.25 / (0.0 - float32(1.5))))) / (3.14159 + (float64(0.0) + 1e10))))
    if v3 {
        switch 9223372036854775807 {
        case 2147483647:
            v6 := GenS1{F0: 95}
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 42}
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 string = "a"
                fmt.Println(fi8, v9, v2)
                var v10 bool = v3
                _ = fi8
                _ = v10
                _ = v9
            }
        case "hello":
            v5 -= 0.1
        case "":
            if v3 {
                fmt.Println(v3, v5)
            }
        default:
            v5 = v5
        }
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        if false {
            if v3 {
                var v11 uint64 = 1
                switch v2 {
                case 2147483647:
                    if v3 {
                        var v12 uint32 = 792
                        v12 = v12
                        _ = v12
                    } else {
                        v2 += 0
                        var v13 bool = (!true)
                        _ = v13
                    }
                default:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v1[v2] = v3
                        _ = fi14
                    }
                }
                _ = v11
            }
        } else {
            v1[v2] = (v3 && true)
            v15 := GenS1{F0: 23}
            if false {
                if v3 {
                    var v16 int8 = -121
                    _ = v16
                }
                var v17 uint16 = 42
                if v3 {
                    v18 := GenS1{F0: 128}
                    v2 -= -2147483648
                    _ = v18
                } else {
                    v19 := GenS2{F0: (^(v17)), F1: 0}
                    var v20 int16 = 277
                    _ = v19
                    _ = v20
                }
                _ = v17
            } else {
                v21 := GenS2{F0: 0, F1: 0}
                v3 = v3
                for fi22 := 0; fi22 < 1; fi22++ {
                    v3 = v3
                    var v23 int32 = 1
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            _ = v15
        }
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 uint8 = 1
            fmt.Println(fi24, v5, v2)
            var v26 uint64 = 2
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 int16 = -878
    v27 -= (-(v27))
    var v28 float64 = 3.14159
    var v29 float64 = float64((((3.14159 / -2.25) + 1e10) + (1.5 / float32((v28 - v5)))))
    if v3 {
        var v30 bool = (!false)
        if v3 {
            var v31 int32 = -900
            _ = v31
        } else {
            switch v2 {
            case -1:
                fmt.Println(v5, v2)
            default:
                var v32 bool = (1.5 != 1.5)
                _ = v32
            }
        }
        _ = v30
    }
    var v33 int64 = 1
    v29 += v28
    fmt.Println(v3, v33)
    var v34 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
