package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 bool }
func genFn0(p0 int, p1 int8, p2 int) bool {
    p0 = p0
    var v5 int8 = p1
    v6 := GenS1{F0: 4294967295, F1: -1}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS2{F0: 867}
        v9 := []int{fi7, (fi7 & (p0 << ((4294967295) & 63))), -331, -9223372036854775808}
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 bool, p1 uint16) int32 {
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS2{F0: 4294967295}
            _ = fi11
            _ = v12
        }
        var v13 uint = 42
        _ = fi10
        _ = v13
    }
    if (32767 >= 1) {
        var v14 float32 = 1.5
        v15 := []int{42, 0}
        var v16 float32 = 0.0
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 string = "a"
        if true {
            v18 := GenS2{F0: 2147483648}
            if ((p0 || true) || p0) {
                v19 := GenS2{F0: 0}
                var v20 float64 = ((float64((-2.25 * (1.5 / 1.5))) / (-2.25 / 1.5)) * (1e10 / 1e10))
                fmt.Println(v20, p0, v20)
                _ = v19
                _ = v20
            } else {
                var v21 int8 = 64
                _ = v21
            }
            v22 := GenS3{F0: 48, F1: 0}
            _ = v18
            _ = v22
        } else {
            for fi23 := 0; fi23 < 6; fi23++ {
                switch 9223372036854775807 {
                case -2147483648:
                    v17 = v17
                case 42:
                    if (42 < 597) {
                        v24 := []int{-821, 0, fi23}
                        _ = v24
                    } else {
                        fmt.Println(v17)
                    }
                case 127:
                    v25 := []int{fi23}
                    _ = v25
                }
                _ = fi23
            }
            fmt.Println(p1, p1)
        }
        fmt.Println(p0, p1)
        _ = v17
    }
    fmt.Println(p0)
    p0 = ((746 == 42) && false)
    _ = p0
    _ = p1
    return 497
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 string = "hello"
    var v4 uint16 = 2
    v26 := []int{128, 32767}
    fmt.Println(v4, v3, v1)
    v1 = false
    v1 = (false && v1)
    var v27 int16 = 42
    v1 = true
    var v28 int8 = -44
    var v29 int16 = v27
    v4 = v4
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 int = 910
        var v32 int16 = (v29 * (v29 + 32767))
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int32 = -645
    v34 := v2[0]
    v4 = 255
    v35 := GenS4{F0: (true && v1), F1: v28, F2: ((0.1 / 0.1) == 3.14159)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
