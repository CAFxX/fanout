package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = false
    var v3 uint = v1
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 uint16 = 255
    var v7 float32 = float32(float64(1e10))
    var v8 float32 = (v7 - v7)
    fmt.Println(v6, v2, v8)
    if v2 {
        switch -220 {
        case 145:
            var v9 int16 = 0
            _ = v9
        case 42:
            fmt.Println(v8)
        default:
            v7 = v7
        }
    }
    v10 := []int{-2147483648, 4294967295}
    v8 += (0.0 - v8)
    switch "hello" {
    case "foo bar":
        if (!(!(-2147483648 != -1))) {
            v6 -= (v6 / ((64) | 1))
        } else {
            var v11 int8 = 34
            _ = v11
        }
    }
    v12 := v5[4294967295]
    var v13 int8 = 1
    v13 -= 7
    if v2 {
        switch -9223372036854775808 {
        case 32767:
            var v14 uint16 = v6
            _ = v14
        case 309:
            v15 := GenS1{F0: 42, F1: ("a" + (v12 + v12)), F2: -9223372036854775808}
            _ = v15
        case 9223372036854775807:
            for fi16 := 0; fi16 < 7; fi16++ {
                _ = len(v10)
                _ = fi16
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 19; fi18++ {
                if false {
                    v2 = ((v7 / v7) >= v7)
                    switch fi17 {
                    case -9223372036854775808:
                        if (false || v2) {
                            v7 = (((v7 / (v7 - -2.25)) - float32(((1e10 * float64(v8)) - 0.0))) * v8)
                        }
                    case 42:
                        var v19 uint16 = (v6 | v6)
                        _ = v19
                    case -1:
                        if v2 {
                            var v20 int8 = v13
                            _ = v20
                        }
                    }
                    fmt.Println(v1, v6)
                } else {
                    if (!(!(!(42 > -352)))) {
                        fmt.Println(v13, fi17, fi17)
                        v21 := GenS1{F0: v1, F1: v12, F2: 978}
                        _ = v21
                    }
                }
                fmt.Println(v8)
                _ = fi18
            }
            var v22 int8 = 127
            v23 := v5[fi17]
            _ = fi17
            _ = v22
            _ = v23
        }
        v24 := v4[v12]
        _ = v24
    } else {
        for fi25 := 0; fi25 < 6; fi25++ {
            v6 += v6
            v12 = "a"
            _ = fi25
        }
        v5[42] = "\n\t"
    }
    v2 = (!(9 >= v13))
    switch v12 {
    case "":
        v2 = true
    case "foo bar":
        var v26 string = ((((v12 + ("foo bar" + v12)) + v12) + ((v12 + "") + v12)) + v12)
        _ = v26
    }
    fmt.Println(v13, v1)
    if v2 {
        v13 += (^(v13))
        v27 := GenS2{F0: (v3 ^ v1), F1: -2147483648, F2: 178}
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            v10 = append(v10, fi29)
            v12 = "\n\t"
            v30 := GenS2{F0: 2, F1: 726, F2: -84}
            _ = fi29
            _ = v30
        }
        fmt.Println(v6, v12, v1)
        if ((!false) || true) {
            var v31 int64 = -903
            var v32 bool = (((true && true) || (false && v2)) && v2)
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
