package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 uint) uint16 {
    v6 := []int{128, -1}
    if true {
        var v7 float64 = (1.5 - (0.0 + ((1e10 * 1.5) * 0.0)))
        var v8 int = 127
        fmt.Println(p1, p2, v8)
        _ = v7
        _ = v8
    }
    var v9 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 0
}
func genFn1(p0 string, p1 string, p2 string) int32 {
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := []int{fi10}
        _ = fi10
        _ = v11
    }
    v12 := []int{2147483647, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 1.5
    v3 := make(map[uint32]int)
    var v4 uint64 = 0
    v5 := make(map[uint32]int)
    var v13 float64 = ((v2 / v2) / 3.14159)
    fmt.Println(v2)
    var v14 string = ""
    if true {
        var v15 int8 = 0
        var v16 int32 = 384
        var v17 float32 = (-2.25 * -2.25)
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: ((float64(0.0) != ((v13 - v2) * (3.14159 / (v13 - v2)))) && true), F1: 18446744073709551615, F2: 937}
    fmt.Println(v2)
    var v19 uint8 = 24
    var v20 int16 = 282
    v21 := []int{486, 127, 840}
    _ = len(v3)
    var v22 uint8 = v19
    var v23 float32 = -2.25
    v22 = v19
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int32 = -983
        if false {
            if (float32(3.14159) >= v23) {
                switch "" {
                case "foo bar":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                case "\n\t":
                    v2 = float64(float32(v13))
                case "hello":
                    if false {
                        fmt.Println(v25)
                        _ = len(v5)
                    }
                }
            } else {
                if true {
                    v26 := GenS1{F0: (1 <= 9223372036854775807), F1: 1, F2: 0}
                    _ = v26
                } else {
                    var v27 uint32 = 240
                    _ = v27
                }
            }
            v1 += (v20 / ((v20) | 1))
            v28 := GenS1{F0: (!(true && true)), F1: 849, F2: 2147483647}
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    switch "hello" {
    case "":
        v21 = append(v21, -901)
    case "\n\t":
        for fi29 := 0; fi29 < 16; fi29++ {
            v23 = v23
            _ = fi29
        }
    case "a":
        if (v4 < 459) {
            if (42 < 804) {
                v23 = 1e10
                v22 -= (^(v19))
            } else {
                v30 := GenS1{F0: (!(!true)), F1: 1, F2: 802}
                v14 = v14
                v31 := GenS2{F0: (v4 | 387)}
                _ = v30
                _ = v31
            }
        }
    default:
        switch "" {
        case "foo bar":
            v1 = v1
        case "hello":
            v21 = append(v21, -9223372036854775808)
        case "a":
            v32 := GenS1{F0: false, F1: 42, F2: 774}
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 bool = ((v19 | (v19 >> ((42) & 63))) >= 0)
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
