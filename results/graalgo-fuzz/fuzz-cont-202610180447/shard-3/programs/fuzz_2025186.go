package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 bool }
type GenS5 struct { F0 int }
func genFn0(p0 uint) bool {
    p0 = (^(p0))
    p0 = p0
    v7 := []int{}
    var v8 int8 = 92
    _ = p0
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", "\n\t", "\n\t", "hello"}
    v3 := make(map[string]string)
    v4 := []string{}
    var v5 string = "hello"
    var v6 bool = false
    v9 := GenS1{F0: 224}
    v6 = v6
    v10 := GenS2{F0: 18446744073709551615, F1: 210, F2: 1}
    v5 = (v5 + (v5 + ("foo bar" + v5)))
    if v6 {
        if (916 > 2) {
            if (4294967295 != 4294967295) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 int32 = 951
                        v14 := GenS3{F0: -1, F1: 50, F2: -472}
                        var v15 int8 = -1
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS4{F0: 76, F1: v6, F2: (42 == 40)}
                    _ = fi11
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int32 = 2147483647
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            v6 = (0.0 <= float64(1e10))
            v6 = v6
            fmt.Println(v5, v5)
        }
    }
    if v6 {
        fmt.Println(v6, v6)
        v5 = v5
        var v19 float32 = -2.25
        _ = v19
    }
    v6 = false
    fmt.Println(v6, v5)
    v20 := GenS5{F0: 909}
    v21 := GenS5{F0: 127}
    var v22 uint64 = 887
    for fi23 := 0; fi23 < 1; fi23++ {
        switch (fi23 >> ((255) & 63)) {
        case 1:
            var v24 int8 = -10
            _ = v24
        case -9223372036854775808:
            if v6 {
                v4 = append(v4, (v5 + v5))
            }
        case 321:
            v6 = v6
        }
        _ = fi23
    }
    v3[(("" + "\n\t") + "")] = v5
    if (false && (127 > -1)) {
        fmt.Println(v22)
        if (!v6) {
            fmt.Println(v22, v6)
        }
        var v25 float32 = float32(0.1)
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if v6 {
            fmt.Println(v5)
        }
        fmt.Println(v6, v6)
        _ = fi26
    }
    var v27 string = v5
    v22 -= (v22 + v22)
    var v28 uint64 = (v22 ^ v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
