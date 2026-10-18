package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 int64 = 876
    v2 := []string{"\n\t", "\n\t", "hello"}
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    var v6 int16 = -864
    var v7 uint32 = 163
    if (!true) {
        v3 = (v3 % ((42) | 1))
        var v8 float32 = float32((0.0 + float64((((1e10 * 1.5) - ((0.1 - float32(1e10)) - float32(float64(1.5)))) - 0.1))))
        if false {
            v9 := GenS1{F0: 1, F1: 1, F2: int8(v6)}
            for fi10 := 0; fi10 < 18; fi10++ {
                v2 = append(v2, "\n\t")
                if false {
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "\n\t")
                    }
                    if false {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            v12 := GenS1{F0: -23, F1: 2, F2: -1}
                            v13 := GenS2{F0: (fi11 >> ((354) & 63)), F1: fi10, F2: "foo bar"}
                            var v14 uint = 0
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        switch "\n\t" {
                        case "a":
                            fmt.Println(fi10)
                        default:
                            v5[("" + "foo bar")] = int(v1)
                        }
                        var v15 uint = 18446744073709551615
                        _ = v15
                    }
                } else {
                    if true {
                        v3 += v3
                    }
                }
                if true {
                    var v16 int32 = -2147483648
                    _ = v16
                }
                _ = fi10
            }
            var v17 int8 = 100
            _ = v17
            _ = v9
        }
        _ = v8
    }
    if false {
        v1 += v1
    }
    var v18 uint8 = 2
    var v19 int = 9223372036854775807
    v20 := GenS1{F0: 120, F1: 572, F2: 42}
    fmt.Println(v18, v7)
    if true {
        if true {
            var v21 bool = true
            var v22 float64 = 1e10
            var v23 uint16 = 280
            _ = v21
            _ = v22
            _ = v23
        }
        v24 := GenS3{F0: v1, F1: v3}
        _ = len(v2)
        _ = v24
    }
    v3 = (v3 * (v3 + v3))
    if false {
        var v25 uint64 = 545
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        var v26 int32 = (v3 % ((int32(v6)) | 1))
        _ = v25
        _ = v26
    }
    var v27 uint32 = v7
    v7 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
