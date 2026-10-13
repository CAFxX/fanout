package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := []uint32{2147483648, 2147483648, 261}
    v2 := make(map[string]uint32)
    var v3 int8 = -1
    v3 -= 42
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float64 = float64(float32(-2.25))
        v3 += 1
        switch "\n\t" {
        case "\n\t":
            if (42 > -340) {
                v6 += (v6 - ((float64(float32((0.0 - (3.14159 + v6)))) + v6) * v6))
                v3 = (-(v3))
            }
        default:
            if (fi5 >= -2147483648) {
                v7 := GenS1{F0: (v4 | v4), F1: 1}
                _ = v7
            } else {
                var v8 string = ""
                if true {
                    var v9 uint8 = 96
                    _ = v9
                } else {
                    _ = len(v1)
                    var v10 int8 = 127
                    _ = v10
                }
                _ = v8
            }
        }
        _ = fi5
        _ = v6
    }
    v4 += 587
    switch 128 {
    case 127:
        for fi11 := 0; fi11 < 4; fi11++ {
            v4 = 1
            v12 := GenS1{F0: v4, F1: 473}
            v4 += 868
            _ = fi11
            _ = v12
        }
    case -659:
        if true {
            var v13 bool = (!true)
            if v13 {
                if true {
                    if (!false) {
                        var v14 bool = (-386 >= -931)
                        var v15 int8 = (v3 / ((v3) | 1))
                        switch 1 {
                        case -801:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 bool = (("hello" + "foo bar") < "\n\t")
                    if (!v13) {
                        v4 = 42
                    } else {
                        var v17 int64 = -482
                        _ = v17
                    }
                    var v18 uint32 = 2147483648
                    _ = v16
                    _ = v18
                }
                v13 = (255 >= 42)
            }
            _ = v13
        }
    }
    if (1 >= 0) {
        switch "foo bar" {
        case "\n\t":
            var v19 int8 = v3
            _ = v19
        default:
            if (1e10 == 3.14159) {
                var v20 int = -1
                _ = v20
            }
        }
    } else {
        v3 = 100
    }
    var v21 uint32 = 1
    var v22 int64 = -216
    v23 := GenS1{F0: 2, F1: 0}
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
