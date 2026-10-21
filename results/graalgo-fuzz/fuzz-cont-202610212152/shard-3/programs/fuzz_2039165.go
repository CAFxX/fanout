package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int; F1 bool }

func main() {
    var v1 int8 = -1
    var v2 float32 = (((0.0 / 1e10) / 1e10) / 1e10)
    var v3 uint8 = 2
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 float32 = float32(float64((float32(1e10) - 1e10)))
    if true {
        fmt.Println(v6, v6, v1)
        v6 = -2.25
    }
    if (817 > 4294967295) {
        if false {
            for fi7 := 0; fi7 < 16; fi7++ {
                v1 -= (^(v1))
                v8 := []int{fi7, (fi7 >> ((255) & 63)), fi7, 554}
                _ = fi7
                _ = v8
            }
            var v9 int8 = (v1 % (((v1 | (-(v1)))) | 1))
            _ = v9
        } else {
            switch 2147483647 {
            case 128:
                v10 := GenS1{F0: true, F1: 2}
                _ = v10
            }
            if false {
                fmt.Println(v1, v2)
            }
        }
        v11 := GenS1{F0: (0 > 918), F1: 776}
        v12 := GenS1{F0: false, F1: 2}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if (!true) {
            v14 := []int{}
            _ = v14
        }
        v1 = 42
        _ = fi13
    }
    if false {
        var v15 float32 = 0.0
        if (5 < 533) {
            for fi16 := 0; fi16 < 7; fi16++ {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v1 = v1
                    v18 := GenS1{F0: true, F1: 1}
                    v15 = (0.1 / v15)
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v3)
                for fi19 := 0; fi19 < 4; fi19++ {
                    v3 = v3
                    var v20 uint32 = 775
                    _ = fi19
                    _ = v20
                }
                _ = fi16
            }
        }
        _ = v15
    }
    var v21 uint = 18446744073709551615
    v22 := GenS2{F0: 695, F1: ((444 != 2) && (2 > 4294967295))}
    delete(v5, "hello")
    fmt.Println(v1, v1)
    v2 = ((float32(1.5) * v6) - (v6 - float32(0.0)))
    if true {
        if (-128 < v1) {
            fmt.Println(v3)
        } else {
            if (-9223372036854775808 == 443) {
                delete(v5, "")
            }
            switch 42 {
            case 32767:
                switch 127 {
                case 2147483647:
                    v1 = (v1 ^ (v1 * 0))
                }
            case 194:
                var v23 int64 = -620
                _ = v23
            case -480:
                v24 := []int{-319}
                _ = v24
            default:
                v21 -= (v21 ^ 42)
            }
            if true {
                var v25 float32 = (-2.25 - v6)
                _ = v25
            }
        }
        v26 := GenS2{F0: -278, F1: false}
        _ = v26
    }
    v27 := GenS2{F0: 1, F1: (false && (true && ((600 >= 2147483647) && true)))}
    v28 := []int{867}
    var v29 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
