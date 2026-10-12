package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 918
    v3 := make(map[uint32]uint32)
    var v4 uint = v2
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint64 = 600
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi5)
            v4 -= v2
            _ = fi7
        }
        v2 = (v2 * v4)
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    fmt.Println(v2, v4)
    if true {
        v2 = v4
        var v8 float64 = 3.14159
        _ = v8
    } else {
        var v9 float64 = 1.5
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = -3
            v9 = (((v9 + ((v9 / (v9 * 1.5)) / 1.0)) / v9) / (((v9 + v9) - v9) / 0.1))
            v12 := []int{(fi10 >> ((859) & 63)), 42}
            _ = fi10
            _ = v11
            _ = v12
        }
        v9 += v9
        _ = v9
    }
    v2 = (v4 - v4)
    if (!false) {
        v4 = (v2 & (v2 / ((1) | 1)))
        fmt.Println(v4)
        if (3.14159 == -2.25) {
            v13 := []int{-1, 127}
            if (!false) {
                v2 -= v2
                var v14 string = "\n\t"
                for fi15 := 0; fi15 < 19; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v17 := GenS1{F0: (!true), F1: 1, F2: -119}
                        switch v14 {
                        case "foo bar":
                            v14 = ""
                        default:
                            v18 := GenS1{F0: (0.0 == (1e10 - 0.1)), F1: -1, F2: -83}
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
                _ = v14
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                _ = len(v1)
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = fi19
            }
            _ = v13
        }
    }
    fmt.Println(v4, v2)
    v4 = v4
    var v20 uint = (v4 << ((65535) & 63))
    v21 := []int{2147483647, 407, 32767}
    v20 = v20
    var v22 uint = v20
    var v23 float64 = float64(-2.25)
    var v24 uint64 = 0
    if (1 > 2147483647) {
        var v25 bool = (((false || true) && (!true)) || true)
        fmt.Println(v24)
        _ = v25
    } else {
        var v26 uint32 = 0
        _ = v26
    }
    fmt.Println(v4)
    v27 := GenS2{F0: -661}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
