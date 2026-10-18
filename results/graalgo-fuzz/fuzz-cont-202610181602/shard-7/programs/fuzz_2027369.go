package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint64) uint {
    p0 += p0
    var v7 float32 = ((0.0 - (((0.0 + (3.14159 * (float32(-2.25) * 0.0))) / (1.5 + 0.1)) / 1.5)) * ((0.1 / 1.0) - 0.1))
    v8 := []int{-2147483648, 4294967295, 32767}
    var v9 float64 = -2.25
    var v10 uint64 = 704
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 float32 = 1.5
    var v4 string = "foo bar"
    var v5 string = "hello"
    var v6 float64 = (1e10 / 1e10)
    var v11 uint = 119
    if len(v2) > 0 {
        v2[0] = (v5 + (v4 + (("hello" + "\n\t") + v4)))
    }
    v12 := GenS1{F0: 42, F1: 4294967295}
    fmt.Println(v3, v3, v5)
    var v13 int8 = -128
    var v14 uint32 = 42
    if (0 <= 1000) {
        v2 = append(v2, ((v5 + v5) + v4))
        var v15 int16 = 0
        _ = v15
    }
    var v16 bool = ((42 == -1) || false)
    for fi17 := 0; fi17 < 1; fi17++ {
        v6 += float64(0.0)
        _ = fi17
    }
    if (genFn0(18446744073709551615) <= v11) {
        _ = len(v2)
    }
    var v18 int16 = 1000
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint8 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(v16, v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
    }
    if (float32((3.14159 - v6)) > (0.1 / (0.1 - (0.1 * v3)))) {
        v13 = v13
    } else {
        if false {
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 int16 = (v18 / (((-(v18))) | 1))
                    v14 = 1
                    var v25 uint32 = (^(v14))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        } else {
            var v26 uint64 = 18446744073709551615
            v27 := GenS1{F0: v13, F1: 4294967295}
            if ((v16 || (34 < 65535)) || (!v16)) {
                for fi28 := 0; fi28 < 18; fi28++ {
                    fmt.Println(v16)
                    if true {
                        v2 = append(v2, v5)
                    }
                    _ = fi28
                }
            }
            _ = v26
            _ = v27
        }
        var v29 float64 = 1.5
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 int32 = -2147483648
        _ = fi30
        _ = v31
    }
    v14 += (v14 >> ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
