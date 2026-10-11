package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }

func main() {
    v1 := []int{2147483647}
    var v2 bool = ("\n\t" > "hello")
    var v3 float32 = -2.25
    var v4 uint = 4294967295
    var v5 string = ((("" + "foo bar") + "") + "a")
    var v6 uint = v4
    var v7 bool = (428 <= 620)
    switch ((v5 + v5) + (v5 + v5)) {
    case "a":
        var v8 float64 = (1e10 - 0.0)
        _ = v8
    default:
        switch "" {
        case "hello":
            if (2147483648 < 2147483647) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: 31, F1: 128, F2: 2}
                    for fi11 := 0; fi11 < 15; fi11++ {
                        if (!(!false)) {
                            var v12 uint32 = 0
                            _ = len(v1)
                            _ = v12
                        } else {
                            if false {
                                var v13 uint32 = 715
                                v3 -= (float32(1.5) + v3)
                                _ = v13
                            }
                        }
                        _ = fi11
                    }
                    _ = fi9
                    _ = v10
                }
                v5 = v5
            }
        default:
            switch v5 {
            case "foo bar":
                v2 = (((((float32(1.5) - v3) - v3) - v3) >= v3) || (float32(((0.0 + 0.0) - -2.25)) < v3))
            default:
                fmt.Println(v6, v7)
            }
        }
    }
    var v14 uint = v4
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float64 = 0.1
        if true {
            fmt.Println(v5)
            if v2 {
                var v17 uint64 = 11
                if v2 {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 int = fi15
                        _ = fi18
                        _ = v19
                    }
                    var v20 int64 = -1
                    v1 = append(v1, fi15)
                    _ = v20
                }
                _ = v17
            }
        } else {
            fmt.Println(v6, v16, v5)
            var v21 uint8 = 1
            _ = v21
        }
        var v22 bool = (!true)
        _ = fi15
        _ = v16
        _ = v22
    }
    fmt.Println(v6, v3, v2)
    fmt.Println(v7, v14)
    for fi23 := 0; fi23 < 19; fi23++ {
        switch fi23 {
        case 117:
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint = 1
                _ = fi24
                _ = v25
            }
        default:
            v3 += ((-2.25 + (0.1 + (3.14159 - ((1e10 / 1.0) + (v3 - ((0.1 / v3) - v3)))))) * v3)
        }
        var v26 uint16 = 65535
        _ = fi23
        _ = v26
    }
    var v27 uint16 = 285
    if true {
        var v28 float64 = 0.0
        _ = len(v1)
        _ = v28
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v4 += v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
