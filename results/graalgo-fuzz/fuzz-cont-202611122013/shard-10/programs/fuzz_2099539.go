package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() int {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{}
        fmt.Println(fi7)
        var v9 uint64 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    fmt.Println(v10)
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v10)
        _ = fi11
    }
    v10 = (32767 <= 42)
    _ = v10
    return 536
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := []int{}
    v4 := make(map[int]string)
    v5 := []int{}
    var v6 int = 4294967295
    switch (^(v6)) {
    case 362:
        var v12 int8 = 100
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v2)
        v6 = -620
        if true {
            var v14 uint16 = 531
            delete(v4, -1)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 int8 = 100
                _ = len(v4)
                v6 += v6
                v6 = (fi13 / ((-631) | 1))
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(fi17)
                fmt.Println(v6, v6, fi13)
                _ = fi17
            }
            switch v6 {
            case 1:
                var v18 float32 = (0.0 + (1.5 - (-2.25 * (((1.5 + (3.14159 / 1.5)) + 3.14159) - 0.0))))
                _ = v18
            case 32767:
                for fi19 := 0; fi19 < 12; fi19++ {
                    v3 = append(v3, v6)
                    var v20 uint64 = 725
                    v2 = v2
                    _ = fi19
                    _ = v20
                }
            case 9223372036854775807:
                if len(v3) > 0 {
                    v3[0] = (fi13 & 0)
                }
            default:
                var v21 int32 = -706
                _ = v21
            }
        }
        _ = fi13
    }
    v2 += (v2 - (v2 | v2))
    v22 := GenS1{F0: 42, F1: 1}
    v23 := GenS2{F0: (v6 - 1), F1: 332}
    v2 += v2
    for fi24 := 0; fi24 < 10; fi24++ {
        if false {
            if len(v3) > 0 {
                v3[0] = genFn0()
            }
            v6 -= v6
        }
        _ = fi24
    }
    if true {
        var v25 int64 = 9223372036854775807
        _ = v25
    }
    var v26 bool = (1 <= -2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
