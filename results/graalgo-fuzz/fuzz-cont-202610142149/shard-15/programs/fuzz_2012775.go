package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint16, p1 int8) string {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 32767}
        _ = fi4
        _ = v5
    }
    p1 = (^(p1))
    var v6 uint64 = 0
    if true {
        var v7 int8 = p1
        _ = v7
    } else {
        var v8 float32 = float32((1e10 / (float64(float32(0.1)) / (1e10 * 0.1))))
        p1 = (-(p1))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return "a"
}
func genFn1(p0 string) uint64 {
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 string = "foo bar"
        v11 := []int{}
        v12 := GenS1{F0: 1}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if (!true) {
        p0 = p0
    }
    _ = p0
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]int)
    v13 := []int{0, 4294967295, 42, -590}
    var v14 uint8 = 1
    var v15 bool = true
    v16 := GenS1{F0: -32768}
    v14 = v14
    v13 = append(v13, -9223372036854775808)
    v14 = (^(v14))
    v17 := v3[""]
    v15 = v15
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: -462}
        if v15 {
            v17 += (fi18 / ((32767) | 1))
        } else {
            var v20 uint8 = v14
            if true {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v2 += v2
                v15 = (!(42 > 42))
            } else {
                var v21 int = (v17 / (((fi18 % ((128) | 1))) | 1))
                fmt.Println(v17)
                _ = v21
            }
            fmt.Println(v20, v20)
            _ = v20
        }
        v17 -= (^(fi18))
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v2)
    var v22 uint8 = (v14 + v14)
    var v23 uint8 = (v22 | (v14 & (v14 % ((v14) | 1))))
    var v24 int = (^(v17))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
