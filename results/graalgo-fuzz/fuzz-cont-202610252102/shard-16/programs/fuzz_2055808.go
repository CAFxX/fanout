package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 32767
    v3 := []uint32{2147483648}
    if len(v3) > 0 {
        v3[0] = 955
    }
    fmt.Println(v2)
    v4 := GenS1{F0: 42, F1: 175, F2: -129}
    if false {
        v5 := v1[v2]
        switch "a" {
        case "foo bar":
            fmt.Println(v5, v2)
        default:
            switch v2 {
            case 526:
                if len(v3) > 0 {
                    v3[0] = 2
                }
            default:
                v5 = "foo bar"
            }
        }
        for fi6 := 0; fi6 < 11; fi6++ {
            switch (v2 >> ((397) & 63)) {
            case 153:
                var v7 int16 = -32768
                _ = v7
            case -454:
                var v8 int8 = 78
                _ = v8
            case 128:
                if false {
                    var v9 int64 = 58
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                if true {
                    v2 -= (^(v2))
                    v11 := v1[(fi6 * (v2 - v2))]
                    var v12 bool = true
                    _ = v11
                    _ = v12
                }
                if ((42 >= 42) || (!false)) {
                    fmt.Println(fi10)
                } else {
                    v5 = v5
                }
                v13 := GenS1{F0: 42, F1: (-(fi6)), F2: 1000}
                _ = fi10
                _ = v13
            }
            v14 := GenS1{F0: 72, F1: (fi6 ^ v2), F2: 358}
            _ = fi6
            _ = v14
        }
        _ = v5
    } else {
        delete(v1, v2)
        v15 := GenS1{F0: 143, F1: v2, F2: 0}
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS1{F0: 5, F1: 2147483647, F2: 32767}
            v2 -= (-(fi16))
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: 127, F1: 915, F2: -1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v19 float64 = (1e10 * (3.14159 - float64((0.0 - 0.1))))
    var v20 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
