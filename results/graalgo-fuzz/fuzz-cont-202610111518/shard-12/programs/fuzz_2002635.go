package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 uint16 }
func genFn0(p0 uint16) int64 {
    v7 := GenS1{F0: -1}
    switch -76 {
    case 9223372036854775807:
        p0 = p0
    case -1:
        if false {
            v8 := GenS2{F0: 2, F1: 0, F2: 0}
            var v9 int8 = 6
            var v10 uint8 = 2
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        fmt.Println(p0)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int32 = -2147483648
        switch (-(fi11)) {
        case 127:
            var v13 float32 = 1.5
            _ = v13
        case 42:
            p0 -= (p0 * p0)
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v7
    return 682
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    var v3 bool = true
    var v4 int8 = -128
    var v5 float64 = 3.14159
    v6 := make(map[int]uint32)
    var v14 float32 = (-2.25 - 1e10)
    var v15 int16 = -32768
    var v16 int64 = 9223372036854775807
    delete(v1, "a")
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS1{F0: 0}
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := []int{(fi17 - (fi19 + fi19)), fi17, fi19, fi19}
            if v3 {
                if v3 {
                    v21 := GenS1{F0: -91}
                    fmt.Println(fi17, v5, fi17)
                    _ = v21
                } else {
                    var v22 int8 = v4
                    _ = v22
                }
                v23 := v6[fi17]
                v16 = genFn0(421)
                var v24 string = "hello"
                _ = v23
                _ = v24
            } else {
                _ = len(v6)
                var v25 int16 = v15
                _ = v25
            }
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v6[842] = 554
    var v26 uint = 0
    var v27 bool = false
    v5 -= (0.0 / v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
