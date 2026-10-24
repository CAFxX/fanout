package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 bool) int64 {
    var v4 float32 = 0.0
    var v5 int = 32768
    fmt.Println(v4, p1)
    v6 := GenS1{F0: 42, F1: 65535}
    for fi7 := 0; fi7 < 9; fi7++ {
        if ((p0 % (((p0 | p0)) | 1)) > 0) {
            if (174 != 18446744073709551615) {
                if (v4 <= v4) {
                    if true {
                        p1 = p1
                        fmt.Println(v4, p1, fi7)
                    } else {
                        fmt.Println(p1, v4, p0)
                        p1 = (!p1)
                        v8 := GenS1{F0: 2, F1: 270}
                        _ = v8
                    }
                    v9 := GenS1{F0: 18446744073709551615, F1: 345}
                    _ = v9
                }
                var v10 int = fi7
                v11 := []int{-9223372036854775808, 370}
                _ = v10
                _ = v11
            }
        }
        var v12 uint = 87
        _ = fi7
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 19
    var v3 uint64 = 18446744073709551615
    var v13 string = ("hello" + "a")
    var v14 uint16 = 265
    var v15 int16 = -1
    switch -2147483648 {
    case 167:
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 float32 = ((0.1 - ((1e10 - 0.0) - (-2.25 * -2.25))) - 0.0)
            var v18 uint = 42
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v15, v15, fi19)
            _ = fi19
        }
    }
    if (!true) {
        v2 = (v2 << ((703) & 63))
        v20 := v1[v13]
        fmt.Println(v13, v20)
        _ = v20
    }
    v21 := []int{32768, -729, 4294967295}
    if (!false) {
        var v22 uint16 = v14
        _ = v22
    } else {
        fmt.Println(v15, v3, v13)
        switch 128 {
        case 341:
            _ = len(v1)
        case 731:
            var v23 uint8 = (v2 | (^(v2)))
            _ = v23
        case 4294967295:
            v3 = v3
        }
        v2 = v2
    }
    var v24 float64 = -2.25
    if true {
        v14 = (^(v14))
        v2 -= 1
        var v25 uint = 193
        _ = v25
    } else {
        var v26 uint8 = (^(v2))
        v27 := GenS2{F0: 1}
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
}
