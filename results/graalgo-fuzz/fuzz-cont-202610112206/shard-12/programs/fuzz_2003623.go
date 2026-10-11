package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
func genFn0() bool {
    v6 := []int{1, 32768, -333, 1}
    for fi7 := 0; fi7 < 4; fi7++ {
        switch 1 {
        case -971:
            if len(v6) > 0 {
                v6[0] = (fi7 | fi7)
            }
        default:
            var v8 uint = 18446744073709551615
            _ = v8
        }
        _ = fi7
    }
    var v9 int64 = 9223372036854775807
    var v10 uint8 = 255
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 42}
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v6
    _ = v9
    return (65535 > 42)
}
func genFn1(p0 bool, p1 int16, p2 bool) int32 {
    var v13 string = "hello"
    if false {
        var v14 float64 = (((-2.25 * (0.1 + 1.5)) / ((3.14159 - 1e10) / -2.25)) / (1e10 / -2.25))
        _ = v14
    }
    v15 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    v5 := []uint32{0}
    var v16 int32 = 2147483647
    _ = len(v5)
    var v17 float32 = 1.5
    var v18 string = "\n\t"
    _ = len(v1)
    v4[2147483647] = 42
    v16 -= -833
    switch 407 {
    case 32768:
        var v19 uint = (^(v2))
        _ = v19
    }
    v20 := GenS1{F0: 1}
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS1{F0: 222}
        switch fi21 {
        case 227:
            switch (fi21 % ((fi21) | 1)) {
            case -9223372036854775808:
                v2 = (v2 % (((v2 * v2)) | 1))
            default:
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
            }
        case 32768:
            v23 := GenS2{F0: ""}
            _ = v23
        }
        var v24 float32 = v17
        _ = fi21
        _ = v22
        _ = v24
    }
    var v25 int32 = 0
    fmt.Println(v25, v25, v25)
    var v26 int8 = 0
    if false {
        v25 = (v25 << ((1) & 63))
        v27 := GenS2{F0: (v18 + v18)}
        fmt.Println(v16, v16, v17)
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
