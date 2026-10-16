package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }
func genFn0(p0 uint32) int32 {
    v6 := []int{9223372036854775807, 32767, 4294967295}
    v7 := GenS1{F0: 42, F1: 2}
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(fi8)
        var v9 int64 = 42
        for fi10 := 0; fi10 < 5; fi10++ {
            v6 = append(v6, (fi10 * 733))
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    p0 -= 478
    var v11 int16 = 42
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    return 506
}
func genFn1(p0 int64, p1 int8) int32 {
    p0 += (p0 / ((p0) | 1))
    for fi12 := 0; fi12 < 6; fi12++ {
        if (255 <= 223) {
            p0 = -177
        } else {
            switch "\n\t" {
            case "a":
                fmt.Println(p1)
            case "":
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := GenS2{F0: 0, F1: -110, F2: p1}
                    _ = fi13
                    _ = v14
                }
            }
        }
        var v15 uint32 = 978
        _ = fi12
        _ = v15
    }
    if (2 >= 1) {
        v16 := GenS2{F0: 244, F1: p1, F2: 42}
        var v17 uint = 42
        _ = v16
        _ = v17
    } else {
        var v18 int16 = 821
        _ = v18
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    var v3 int8 = 127
    v4 := make(map[int]string)
    var v5 int = -2147483648
    v2 += v2
    var v19 uint32 = 936
    var v20 uint8 = 2
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := []int{fi21}
        _ = fi21
        _ = v22
    }
    delete(v4, (v5 * (^(v5))))
    v23 := GenS1{F0: 292, F1: (v20 - v20)}
    v3 = v3
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(fi25, v5, v3)
            switch -9223372036854775808 {
            case 4294967295:
                v26 := GenS2{F0: 4294967295, F1: (v3 + v3), F2: v3}
                _ = v26
            case 32767:
                v5 = (v5 % (((v5 & (-(fi24)))) | 1))
            }
            var v27 string = "\n\t"
            _ = fi25
            _ = v27
        }
        switch (v5 * (v5 % (((v5 & (^(fi24)))) | 1))) {
        case 42:
            switch (fi24 / ((v5) | 1)) {
            case 2147483647:
                var v28 float64 = ((v2 * (((v2 + (((v2 + v2) - v2) / float64(float32(((((v2 + (((v2 - 0.1) * v2) * 3.14159)) + 0.1) / 0.1) + (v2 / (((-2.25 - (3.14159 / 1.0)) - (v2 + -2.25)) + v2))))))) * ((1e10 - 0.1) + v2)) / v2)) + v2)
                _ = v28
            case -797:
                v4[(^(fi24))] = "a"
            }
        case 4294967295:
            var v29 float64 = (float64(0.1) * float64(3.14159))
            _ = v29
        case 0:
            fmt.Println(v3)
        }
        _ = fi24
    }
    v2 = (v2 - (0.0 - float64(0.1)))
    if true {
        v19 += v19
        for fi30 := 0; fi30 < 8; fi30++ {
            v19 = (v19 | v19)
            _ = fi30
        }
        var v31 uint = 0
        _ = v31
    }
    v5 -= 0
    var v32 string = "a"
    var v33 int = v5
    v32 = "foo bar"
    fmt.Println(v3, v32, v33)
    v3 -= (v3 % ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
