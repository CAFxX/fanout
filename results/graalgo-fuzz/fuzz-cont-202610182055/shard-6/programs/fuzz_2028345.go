package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int) int32 {
    var v6 float32 = (0.0 + 1e10)
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 = p0
        _ = fi7
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint64, p1 uint, p2 uint32) int {
    v8 := []int{-1000, 0}
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            p2 = (p2 - p2)
            _ = fi10
        }
        var v11 float32 = 0.1
        _ = fi9
        _ = v11
    }
    var v12 int16 = 1
    v8 = append(v8, 1)
    var v13 bool = (255 <= 65535)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return 127
}

func main() {
    var v1 string = ""
    var v2 int8 = -1
    var v3 uint64 = 4294967295
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v14 int16 = -32768
    fmt.Println(v3, v1, v2)
    if true {
        switch int(v3) {
        case 304:
            var v15 float64 = 1e10
            _ = v15
        case -580:
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 string = v1
                for fi18 := 0; fi18 < 14; fi18++ {
                    v14 -= (v14 / ((-32768) | 1))
                    switch (v17 + v1) {
                    case "\n\t":
                        v1 = (((v17 + ("a" + v1)) + "hello") + "hello")
                    default:
                        var v19 uint8 = 1
                        _ = v19
                    }
                    _ = fi18
                }
                fmt.Println(v3)
                _ = fi16
                _ = v17
            }
        default:
            var v20 int8 = 1
            _ = v20
        }
    } else {
        var v21 float32 = (-2.25 + (1.5 + 3.14159))
        _ = v21
    }
    v2 -= v2
    v22 := GenS1{F0: (!false), F1: 2}
    var v23 int = -1
    v24 := GenS2{F0: ((v2 % (((^(v2))) | 1)) <= v2)}
    var v25 uint8 = 2
    v26 := []int{1, v23, v23}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
