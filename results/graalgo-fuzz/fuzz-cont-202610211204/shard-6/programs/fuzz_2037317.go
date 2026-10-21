package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []uint32{}
    var v2 float64 = 0.0
    var v3 uint64 = 294
    var v4 uint64 = v3
    var v5 bool = (false && (1 < -1))
    if v5 {
        v6 := GenS1{F0: 4294967295, F1: 1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            switch ("\n\t" + "foo bar") {
            case "":
                var v8 bool = (-1 == 705)
                _ = v8
            }
            var v9 int64 = -277
            if (2147483648 >= 2) {
                v10 := GenS1{F0: 2, F1: -32}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if len(v1) > 0 {
                        v1[0] = 481
                    }
                    fmt.Println(v5, fi11, v2)
                    v5 = true
                    var v12 uint32 = 1
                    var v13 string = ""
                    v3 = (v4 ^ 4294967295)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 int16 = 32767
                for fi15 := 0; fi15 < 1; fi15++ {
                    if v5 {
                        var v16 float32 = (float32((v2 / v2)) * 1e10)
                        switch fi7 {
                        case -2147483648:
                            var v17 string = ("hello" + "a")
                            _ = v17
                        case 32767:
                            v18 := GenS1{F0: 516, F1: (v9 / (((^(v9))) | 1))}
                            _ = v18
                        default:
                            v19 := GenS2{F0: v4, F1: (v14 >> ((2) & 63))}
                            _ = v19
                        }
                        _ = v16
                    }
                    _ = fi15
                }
                _ = v14
            }
            _ = fi7
            _ = v9
        }
    }
    var v20 int32 = 2147483647
    v4 = (v4 >> ((161) & 63))
    var v21 int8 = -61
    v21 -= v21
    _ = len(v1)
    for fi22 := 0; fi22 < 11; fi22++ {
        _ = len(v1)
        if ("foo bar" <= "a") {
            var v23 int = (fi22 | fi22)
            fmt.Println(v20, v5)
            var v24 bool = false
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    v1 = append(v1, 1)
    fmt.Println(v20, v21)
    v5 = ((!true) || v5)
    var v25 string = ("foo bar" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
