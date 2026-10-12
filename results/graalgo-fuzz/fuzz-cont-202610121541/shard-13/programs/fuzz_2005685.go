package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
func genFn0(p0 string) string {
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 143
    fmt.Println(v5, v6, p0)
    if (false && (!(138 != v6))) {
        fmt.Println(v5, p0)
        if true {
            var v7 float64 = ((0.0 - 0.0) / 1.0)
            _ = v7
        } else {
            var v8 int16 = 546
            p0 = (p0 + p0)
            _ = v8
        }
        if true {
            if true {
                v6 += 2
            }
        }
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{fi9, fi9}
        p0 = p0
        switch (fi9 >> ((4294967295) & 63)) {
        case 32768:
            v6 += v6
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v6
    return (("hello" + p0) + p0)
}
func genFn1() uint32 {
    var v11 float64 = 0.1
    var v12 uint8 = 0
    _ = v11
    _ = v12
    return 491
}

func main() {
    v1 := []int{0}
    v2 := make(map[string]string)
    var v3 bool = (!(42 >= -479))
    v4 := []uint32{979, 368, 1, 1}
    v13 := GenS1{F0: -411, F1: 1, F2: -328}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = -1
        for fi16 := 0; fi16 < 20; fi16++ {
            v15 -= (v15 % ((-1) | 1))
            if len(v4) > 0 {
                v4[0] = 65535
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v17 int32 = 42
    var v18 string = (("foo bar" + "foo bar") + "\n\t")
    v18 = ((v18 + "hello") + "")
    var v19 uint = 2
    fmt.Println(v3, v18, v17)
    v18 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
