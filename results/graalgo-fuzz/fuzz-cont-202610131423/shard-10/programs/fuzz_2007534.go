package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 int, p1 int32, p2 int) int32 {
    v5 := GenS1{F0: 1, F1: 42, F2: 127}
    var v6 uint8 = 35
    for fi7 := 0; fi7 < 19; fi7++ {
        if true {
            v8 := GenS1{F0: 97, F1: 199, F2: v6}
            var v9 int64 = int64(v6)
            _ = v8
            _ = v9
        }
        p2 = -956
        _ = fi7
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := []int{(^(p0)), p0, (p2 + 127)}
        var v12 float64 = -2.25
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{(^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    _ = v6
    return (^(p1))
}

func main() {
    v1 := []string{("" + "foo bar")}
    var v2 string = "foo bar"
    v3 := make(map[string]uint32)
    var v4 uint16 = 0
    var v14 int = int(v4)
    if ((v2 + v2) >= "") {
        v14 = -208
        if false {
            if true {
                var v15 string = v2
                v16 := GenS1{F0: -1, F1: 128, F2: 0}
                var v17 int = 128
                _ = v15
                _ = v16
                _ = v17
            } else {
                if (0.0 == 0.1) {
                    v18 := GenS1{F0: -128, F1: 203, F2: 2}
                    var v19 int = 197
                    _ = v18
                    _ = v19
                }
                v4 -= (v4 + v4)
                if true {
                    v1 = append(v1, v2)
                } else {
                    switch v2 {
                    case "foo bar":
                        v2 = v2
                    }
                }
            }
        }
        v20 := GenS1{F0: 42, F1: 0, F2: 1}
        _ = v20
    }
    var v21 float64 = 1e10
    v21 += float64(-2.25)
    if (!false) {
        v14 = 0
        var v22 int32 = -1
        fmt.Println(v22, v2)
        _ = v22
    }
    fmt.Println(v14, v21, v21)
    var v23 bool = false
    v23 = v23
    var v24 uint32 = 4294967295
    v25 := GenS2{F0: v4}
    v4 = 255
    v23 = v23
    for fi26 := 0; fi26 < 9; fi26++ {
        _ = len(v1)
        delete(v3, v2)
        delete(v3, (((v2 + (v2 + "a")) + v2) + ("a" + "\n\t")))
        _ = fi26
    }
    switch (v14 * (v14 % (((v14 ^ (v14 % (((v14 ^ (v14 / ((v14) | 1)))) | 1)))) | 1))) {
    case 4294967295:
        fmt.Println(v4, v2, v2)
    case 451:
        v2 = (v2 + v2)
    case 128:
        if len(v1) > 0 {
            v1[0] = v2
        }
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 11; fi28++ {
                v29 := v3[v2]
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: -1, F1: 42, F2: 255}
            _ = fi27
            _ = v30
        }
    }
    v31 := GenS3{F0: -1, F1: 4294967295}
    _ = len(v1)
    v21 += float64(float32(v21))
    var v32 uint16 = 40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
