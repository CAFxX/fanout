package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0() bool {
    var v4 uint16 = 2
    if (1.5 > (float32(0.1) * 0.0)) {
        var v5 float64 = -2.25
        _ = v5
    }
    fmt.Println(v4)
    _ = v4
    return false
}

func main() {
    var v1 string = ("a" + "a")
    var v2 int8 = -128
    var v3 bool = true
    v6 := GenS1{F0: v3, F1: 9223372036854775807}
    v7 := []int{0, 1}
    fmt.Println(v3, v3)
    fmt.Println(v2, v1)
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS2{F0: 475, F1: 42, F2: 65535}
            v2 += (-(v2))
            _ = fi8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(v2, v3)
            if true {
                for fi11 := 0; fi11 < 4; fi11++ {
                    if (1 != fi10) {
                        var v12 int8 = v2
                        v13 := GenS1{F0: (!(((1.5 + (0.0 - (1e10 + 1e10))) / ((3.14159 + float64(1e10)) / 0.1)) != float64(float32(((0.0 / 1.0) / 1.5))))), F1: 964}
                        fmt.Println(fi10, v1)
                        _ = v12
                        _ = v13
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = fi11
                }
                var v14 uint16 = 65535
                _ = v14
            }
            if v3 {
                for fi15 := 0; fi15 < 14; fi15++ {
                    if (false || ((1.5 * (0.1 - 0.1)) < (3.14159 / 1.0))) {
                        var v16 uint16 = 29
                        _ = v16
                    } else {
                        var v17 int = (fi15 - 140)
                        _ = v17
                    }
                    _ = fi15
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(fi10, fi18)
                    v19 := GenS3{F0: 127, F1: 35}
                    var v20 int64 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v1)
            }
            _ = fi10
        }
    }
    var v21 int64 = 1
    if (false || (!(v3 && (1000 < 1000)))) {
        v3 = ((true && (2147483647 <= 9223372036854775807)) && true)
        v21 = (v21 >> ((255) & 63))
    }
    var v22 uint = 308
    var v23 int32 = -730
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 uint16 = 65535
        var v26 uint = 971
        if (false || ((!false) || v3)) {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 uint32 = 651
                _ = fi27
                _ = v28
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v2 = v2
    var v29 int64 = (^(v21))
    if len(v7) > 0 {
        v7[0] = 0
    }
    var v30 int64 = 9223372036854775807
    v7 = append(v7, -277)
    v22 = 361
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 3; fi32++ {
            fmt.Println(fi32)
            v33 := GenS3{F0: fi32, F1: 75}
            _ = fi32
            _ = v33
        }
        v21 -= v21
        if ((-(v23)) <= (v23 ^ (v23 & v23))) {
            v34 := GenS3{F0: fi31, F1: 127}
            var v35 uint64 = 18446744073709551615
            fmt.Println(v21, v2, v1)
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v6
    _ = v7
}
