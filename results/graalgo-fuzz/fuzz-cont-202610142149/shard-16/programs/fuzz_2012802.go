package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
func genFn0(p0 int16, p1 uint32, p2 bool) string {
    v5 := GenS1{F0: 128, F1: 255, F2: 11}
    var v6 int8 = 0
    p0 = p0
    p2 = (true || ((1e10 / ((1.5 * (float32((3.14159 / 3.14159)) + float32(3.14159))) * (1.5 / 3.14159))) < -2.25))
    v7 := GenS1{F0: -734, F1: 708, F2: (v6 & (v6 | (v6 & -122)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("a" + "")
}
func genFn1(p0 bool, p1 uint16, p2 int32) uint8 {
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = ((0.0 + 0.0) * (1e10 / 1.0))
        if p0 {
            var v10 uint16 = p1
            p1 = (p1 + v10)
            _ = v10
        } else {
            fmt.Println(p2, v9, fi8)
            if (!p0) {
                if false {
                    p0 = (true || p0)
                    v11 := []int{fi8}
                    _ = v11
                } else {
                    if p0 {
                        var v12 uint32 = 2147483648
                        var v13 uint8 = 147
                        _ = v12
                        _ = v13
                    } else {
                        var v14 float64 = v9
                        _ = v14
                    }
                    var v15 string = ((("a" + "\n\t") + "\n\t") + "foo bar")
                    var v16 uint32 = 4294967295
                    _ = v15
                    _ = v16
                }
            }
            var v17 bool = (p0 && p0)
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    v18 := []int{}
    var v19 uint32 = 42
    if (!(42 > v19)) {
        fmt.Println(v19)
        if p0 {
            var v20 int = -916
            _ = v20
        } else {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if true {
            if false {
                p1 -= (p1 ^ 854)
                fmt.Println(p2, fi21)
                p1 = (p1 + p1)
            }
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 float32 = float32(-2.25)
                v24 := GenS1{F0: fi21, F1: 42, F2: -128}
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(fi21, p2, p1)
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 246
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    v2 = false
    v25 := v4[("a" + "foo bar")]
    v26 := GenS1{F0: 32767, F1: 65535, F2: 69}
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 bool = (42 < 0)
        v25 = v25
        _ = fi27
        _ = v28
    }
    v2 = true
    for fi29 := 0; fi29 < 18; fi29++ {
        v25 = (((("" == ("foo bar" + "foo bar")) || (v3 >= 255)) || (-2.25 < 0.0)) && (32767 != 1000))
        v25 = true
        var v30 int16 = 0
        _ = fi29
        _ = v30
    }
    v31 := []int{1, 42, 537}
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
