package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0() uint16 {
    v7 := []int{-265, 127}
    if ((0.0 > (1.5 * 1e10)) || false) {
        if (4294967295 < 9223372036854775807) {
            for fi8 := 0; fi8 < 1; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    if true {
                        var v10 uint32 = 665
                        fmt.Println(v10, v10, fi9)
                        _ = v10
                    }
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v7)
                    }
                    v7 = append(v7, 814)
                    _ = fi9
                }
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 uint64 = 18446744073709551615
                    fmt.Println(fi8, fi8)
                    _ = fi11
                    _ = v12
                }
                v7 = append(v7, 1)
                _ = fi8
            }
            v13 := GenS1{F0: 1, F1: 0}
            for fi14 := 0; fi14 < 2; fi14++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi14
            }
            _ = v13
        } else {
            var v15 bool = true
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v15
        }
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 261
    v3 := make(map[string]string)
    var v4 int8 = 1
    var v5 string = "foo bar"
    var v6 uint8 = 0
    fmt.Println(v2, v2)
    v5 = "a"
    var v16 int64 = 9223372036854775807
    var v17 int16 = 32767
    v18 := []int{42, 42}
    if ((v6 << ((1) & 63)) > (^(v6))) {
        if len(v18) > 0 {
            v18[0] = -42
        }
    }
    v19 := GenS1{F0: v16, F1: (v4 + -1)}
    for fi20 := 0; fi20 < 14; fi20++ {
        if ((v5 + v5) >= "a") {
            switch "" {
            case "foo bar":
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 bool = (!(true && (42 == 1)))
                    var v23 uint32 = v2
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                var v24 int16 = 1000
                _ = v24
            }
            var v25 string = v5
            v25 = "\n\t"
            _ = v25
        } else {
            v16 += (v16 / (((v16 + v16)) | 1))
            switch fi20 {
            case -568:
                fmt.Println(v4, v17, fi20)
            }
        }
        _ = fi20
    }
    v5 = ("hello" + v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
