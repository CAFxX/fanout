package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 int16, p1 uint64) int {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint = 4294967295
        if false {
            var v6 int8 = 100
            v6 += (v6 % ((v6) | 1))
            _ = v6
        }
        var v7 uint64 = (^(p1))
        _ = fi4
        _ = v5
        _ = v7
    }
    p1 = 2
    var v8 float64 = 1e10
    _ = p0
    _ = p1
    _ = v8
    return 2147483647
}
func genFn1(p0 bool) int16 {
    p0 = (1 < 2)
    v9 := GenS1{F0: 717, F1: -665}
    p0 = (42 < 0)
    var v10 uint = 255
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 225
    var v2 bool = false
    var v3 string = (("foo bar" + ((("hello" + "") + "foo bar") + "hello")) + (("foo bar" + ("hello" + ((("foo bar" + "\n\t") + "") + "\n\t"))) + "hello"))
    v3 = (v3 + (v3 + (v3 + "a")))
    v11 := []int{}
    if len(v11) > 0 {
        v11[0] = 9223372036854775807
    }
    var v12 float64 = ((1.5 / 1.5) + (float64(1e10) * 1e10))
    for fi13 := 0; fi13 < 2; fi13++ {
        if v2 {
            if true {
                if false {
                    v11 = append(v11, 9223372036854775807)
                }
            }
            if v2 {
                v2 = v2
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: 2147483647, F1: genFn1(v2)}
                _ = fi14
                _ = v15
            }
        } else {
            var v16 uint = 255
            _ = v16
        }
        v12 -= v12
        v17 := GenS1{F0: 42, F1: 0}
        _ = fi13
        _ = v17
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if v2 {
        var v18 int64 = -9223372036854775808
        _ = v18
    } else {
        v1 -= 0
        var v19 int64 = -1
        var v20 uint16 = 133
        _ = v19
        _ = v20
    }
    v3 = ((((v3 + v3) + "hello") + "a") + ("" + v3))
    v11 = append(v11, 543)
    fmt.Println(v3, v2)
    var v21 int64 = -818
    if (true || ((1e10 - 0.0) == float32((1e10 + float64(1e10))))) {
        v22 := GenS1{F0: -1, F1: 1}
        v21 = v21
        for fi23 := 0; fi23 < 17; fi23++ {
            if v2 {
                v11 = append(v11, (^(fi23)))
                var v24 uint64 = 42
                var v25 bool = (!((false || false) || v2))
                _ = v24
                _ = v25
            }
            var v26 int32 = 2147483647
            _ = fi23
            _ = v26
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
}
