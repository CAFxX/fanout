package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 2147483647
    var v3 int64 = 1
    v4 := make(map[int]uint32)
    var v5 string = "a"
    var v6 float32 = 1e10
    if false {
        v5 = v5
        if false {
            v7 := []int{}
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 int16 = 1000
                    _ = fi9
                    _ = v10
                }
                var v11 uint16 = 65535
                _ = fi8
                _ = v11
            }
            v2 += (^(v2))
            _ = v7
        }
        switch v5 {
        case "\n\t":
            if false {
                v2 -= v2
            } else {
                var v12 int64 = (v3 % ((1) | 1))
                _ = v12
            }
        default:
            v2 = v2
        }
    } else {
        var v13 uint32 = 132
        _ = v13
    }
    v5 = "\n\t"
    fmt.Println(v5)
    var v14 string = v5
    v15 := v1["foo bar"]
    if (212 <= 255) {
        if true {
            fmt.Println(v3)
            v16 := GenS1{F0: v14, F1: -1, F2: 822}
            v17 := []int{(^(v2)), v2, v2, (v2 << ((4294967295) & 63))}
            _ = v16
            _ = v17
        }
        if v15 {
            if (!v15) {
                var v18 uint16 = 65535
                v6 = 0.0
                var v19 uint16 = 2
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 int8 = 127
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v14)
            }
        } else {
            v15 = (811 <= 42)
            v22 := []int{(v2 << ((65535) & 63))}
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := GenS1{F0: "\n\t", F1: 2147483647, F2: -500}
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        fmt.Println(v5)
    } else {
        v25 := GenS1{F0: ((v5 + "\n\t") + v5), F1: 906, F2: -459}
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 float64 = 0.0
                var v29 int8 = 100
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        _ = v25
    }
    if false {
        fmt.Println(v15, v6, v3)
        v15 = v15
    }
    var v30 uint8 = 208
    for fi31 := 0; fi31 < 12; fi31++ {
        fmt.Println(v6, v15)
        var v32 float64 = 0.0
        v15 = (-2.25 != (1e10 - ((v6 / 1e10) * 1e10)))
        _ = fi31
        _ = v32
    }
    var v33 uint64 = 42
    if v15 {
        switch v2 {
        case 4294967295:
            var v34 int8 = 76
            _ = v34
        default:
            v15 = v15
        }
        v5 = "a"
        v30 += (v30 | (v30 - v30))
    } else {
        v15 = (!true)
    }
    v35 := []int{(v2 % ((v2) | 1)), (^(v2)), (v2 + (v2 & v2))}
    v36 := v1[(v5 + "foo bar")]
    var v37 int16 = 0
    v6 = 0.1
    var v38 uint64 = 188
    fmt.Println(v14, v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
