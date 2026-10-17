package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 int8 = 0
    var v5 bool = (!false)
    if (1 < 2147483648) {
        v1 = append(v1, "")
    } else {
        var v6 float32 = -2.25
        for fi7 := 0; fi7 < 9; fi7++ {
            if v5 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi7
        }
        _ = v6
    }
    v8 := v2["hello"]
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 bool = (0 >= -720)
        var v11 uint16 = 798
        switch 32767 {
        case 42:
            var v12 float64 = (1e10 / -2.25)
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 128 {
    case -1:
        var v13 bool = true
        _ = v13
    default:
        if true {
            var v14 int16 = 32767
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v4, v4, v8)
            _ = v14
        }
    }
    var v15 uint32 = 1
    var v16 uint32 = (v15 / (((v15 % (((v15 & 68)) | 1))) | 1))
    v5 = v5
    if (945 == -1) {
        var v17 float32 = (float32(0.0) - 0.1)
        _ = v17
    } else {
        if false {
            if ((585 >= 1) && (v16 <= 1)) {
                var v18 int = 127
                var v19 int = (v18 | (v18 ^ v18))
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 int32 = 2147483647
                    v22 := GenS1{F0: ((v8 + "\n\t") + "a"), F1: false}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := GenS2{F0: 1}
                        v5 = v5
                        _ = fi23
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v18
                _ = v19
            }
            switch (((v8 + (v8 + v8)) + "\n\t") + v8) {
            case "":
                if false {
                    v4 += v4
                }
            }
            v8 = "\n\t"
        }
        v15 = 241
    }
    v16 -= 308
    v16 = (v16 - v16)
    if ((1000 < -105) && v5) {
        fmt.Println(v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
