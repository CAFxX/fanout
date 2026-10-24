package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 int }

func main() {
    var v1 int8 = 90
    var v2 uint64 = 0
    var v3 uint32 = 2147483648
    v1 = v1
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 int16 = -862
        v6 := GenS1{F0: (^(v1)), F1: 1, F2: "a"}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float32 = float32(float64(-2.25))
    if false {
        var v8 int64 = -91
        if (0 >= -1) {
            fmt.Println(v1, v7)
        } else {
            for fi9 := 0; fi9 < 18; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 int32 = -1
                    v12 := []int{int(v2), fi9}
                    if false {
                        v13 := GenS1{F0: (v1 & v1), F1: (v3 - 2), F2: "\n\t"}
                        v14 := GenS1{F0: (v1 & (v1 / ((-23) | 1))), F1: 42, F2: ("foo bar" + "foo bar")}
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v15 uint = 42
                _ = fi9
                _ = v15
            }
            v16 := []int{}
            v16 = append(v16, 98)
            _ = v16
        }
        v8 = v8
        _ = v8
    }
    v17 := GenS2{F0: -61}
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int = fi18
        v3 -= v3
        _ = fi18
        _ = v19
    }
    v1 -= (v1 & 0)
    v20 := GenS1{F0: v1, F1: (v3 / (((v3 % (((^(v3))) | 1))) | 1)), F2: ""}
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := []int{(fi21 | fi21), 128}
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v22) > 0 {
                v22[0] = fi23
            }
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    if (238 == 89) {
        if false {
            fmt.Println(v2, v2)
        }
    } else {
        var v24 int32 = 2147483647
        _ = v24
    }
    var v25 int64 = -1
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v1, v3, v1)
        v27 := []int{fi26, fi26}
        v25 -= (^(v25))
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v7
}
