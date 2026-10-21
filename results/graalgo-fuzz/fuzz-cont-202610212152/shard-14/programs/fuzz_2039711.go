package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
func genFn0() uint {
    v7 := []int{827, 831, 0}
    var v8 uint64 = 42
    v7 = append(v7, 42)
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 bool, p2 int32) int16 {
    if ((!(638 >= 83)) || (((((3.14159 + (3.14159 + 1e10)) + float32(1e10)) / (1.5 - 0.1)) / float32(3.14159)) != 1e10)) {
        v9 := GenS1{F0: 4294967295, F1: 0, F2: 128}
        p2 -= p2
        _ = v9
    }
    var v10 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1000
}

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    var v4 float64 = (1e10 * (1.5 + 1.5))
    v5 := []int{158, 32768, 516}
    var v6 bool = ((v1 && (686 >= 0)) || false)
    fmt.Println(v4, v1, v3)
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (v4 / v4)
    v6 = false
    var v11 int64 = 374
    var v12 bool = (!v6)
    if (127 < uint8(v3)) {
        v5 = append(v5, 2147483647)
        switch 1 {
        case -2147483648:
            fmt.Println(v4, v12)
        default:
            switch "a" {
            case "a":
                if v6 {
                    _ = len(v2)
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(v11, fi13)
                        var v14 int8 = 100
                        var v15 string = (("\n\t" + "hello") + ("hello" + ("" + "")))
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            case "hello":
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 int32 = -1
                    v4 = v4
                    for fi18 := 0; fi18 < 5; fi18++ {
                        fmt.Println(v3)
                        if len(v5) > 0 {
                            v5[0] = (fi18 ^ (fi18 + (fi16 + (^(fi18)))))
                        }
                        _ = fi18
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
        v19 := GenS1{F0: (v3 / ((v3) | 1)), F1: genFn1(1, v12, -2147483648), F2: 330}
        _ = v19
    } else {
        var v20 uint32 = 1
        _ = v20
    }
    v1 = (2147483648 != 250)
    var v21 int64 = (v11 % (((-(v11))) | 1))
    fmt.Println(v1, v12)
    delete(v2, "hello")
    var v22 int64 = -965
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v5)
    for fi23 := 0; fi23 < 14; fi23++ {
        _ = len(v2)
        v24 := GenS1{F0: 963, F1: 0, F2: (fi23 + 285)}
        v25 := GenS1{F0: 2, F1: genFn1(32767, v1, 0), F2: (^(fi23))}
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: v3, F1: 1000, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
