package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    v5 := make(map[string]uint32)
    v6 := make(map[int]int)
    if (503 >= 1000) {
        var v7 uint64 = 18446744073709551615
        fmt.Println(v4)
        _ = v7
    }
    var v8 uint64 = 1
    if v1 {
        v8 = (v8 - 4294967295)
        var v9 int = 1
        _ = v9
    }
    if ((65535 == 2147483648) && ((!true) && (false && (0 <= 752)))) {
        var v10 bool = ((((v4 - 0.0) * v4) - float64(0.0)) <= float64((0.0 / (3.14159 * 1e10))))
        v11 := []int{-9223372036854775808, 55}
        _ = v10
        _ = v11
    }
    v8 = 616
    v4 = v4
    fmt.Println(v1)
    v12 := GenS1{F0: 0, F1: 2, F2: (v1 && (!v1))}
    v13 := GenS2{F0: (v8 & v8), F1: 9223372036854775807}
    var v14 int8 = 1
    v4 -= v4
    v1 = true
    v1 = v1
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            delete(v5, "foo bar")
            fmt.Println(v4, fi15, v1)
        }
        var v16 int8 = 127
        _ = fi15
        _ = v16
    }
    if v1 {
        if v1 {
            switch "foo bar" {
            case "foo bar":
                if (((3.14159 / 1.0) <= 3.14159) && v1) {
                    fmt.Println(v8, v1, v4)
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v18 := GenS1{F0: v8, F1: 2, F2: v1}
                        v14 = (v14 - -128)
                        _ = fi17
                        _ = v18
                    }
                } else {
                    switch 4294967295 {
                    case 191:
                        fmt.Println(v8, v14)
                    case -381:
                        var v19 uint16 = 32767
                        _ = v19
                    case 18:
                        var v20 uint = 65535
                        _ = v20
                    default:
                        if (255 <= 2) {
                            v1 = v1
                        } else {
                            v8 -= v8
                        }
                    }
                    v1 = true
                    var v21 float64 = 1e10
                    _ = v21
                }
            case "a":
                v4 += -2.25
            default:
                for fi22 := 0; fi22 < 11; fi22++ {
                    delete(v2, "")
                    _ = fi22
                }
            }
        }
        var v23 int = 4294967295
        v24 := []int{(v23 * v23), v23, v23, v23}
        _ = v23
        _ = v24
    }
    var v25 float32 = -2.25
    var v26 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
