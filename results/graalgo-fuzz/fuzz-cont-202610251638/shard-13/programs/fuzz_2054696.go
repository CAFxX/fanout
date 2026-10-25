package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string, p1 uint32) int8 {
    var v7 int16 = 752
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(p0, v7, p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -9223372036854775808
    var v3 float32 = 3.14159
    v4 := make(map[int]uint32)
    v5 := []int{-788}
    var v6 uint64 = 42
    switch "a" {
    case "":
        var v9 uint16 = 32767
        _ = v9
    }
    v3 += (((float32((0.1 + ((1e10 - -2.25) - float64(1e10)))) / v3) / v3) / (-2.25 + -2.25))
    var v10 float64 = 3.14159
    delete(v4, 218)
    if false {
        switch -2147483648 {
        case 42:
            v3 = v3
        case 389:
            var v11 int32 = 42
            _ = v11
        }
        var v12 float32 = 0.1
        _ = v12
    }
    fmt.Println(v3, v3, v10)
    v1[0] = "hello"
    fmt.Println(v6, v2, v2)
    v2 = v2
    if (v6 > v6) {
        fmt.Println(v2)
    } else {
        if (9223372036854775807 > -9223372036854775808) {
            for fi13 := 0; fi13 < 15; fi13++ {
                if true {
                    v3 = v3
                    v5 = append(v5, (fi13 | (fi13 % (((-(fi13))) | 1))))
                    v2 -= 1
                } else {
                    v14 := GenS1{F0: (255 < 0), F1: 446}
                    _ = v14
                }
                v10 = 3.14159
                if false {
                    if ((v3 != 0.0) || true) {
                        v15 := GenS1{F0: (v10 == float64(float32(v10))), F1: 9223372036854775807}
                        v6 -= v6
                        v16 := GenS2{F0: 42}
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int16 = 42
                        _ = v17
                    }
                    fmt.Println(v2, fi13, v10)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v5) > 0 {
                        v5[0] = fi13
                    }
                    v18 := GenS1{F0: false, F1: (fi13 | -227)}
                    _ = v18
                }
                _ = fi13
            }
            if (!true) {
                var v19 int = -2147483648
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch (fi20 << ((1) & 63)) {
                    case 290:
                        for fi21 := 0; fi21 < 8; fi21++ {
                            fmt.Println(fi21)
                            _ = fi21
                        }
                    }
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v3 = v3
                        v19 += -527
                        fmt.Println(v3)
                        _ = fi22
                    }
                    v23 := GenS1{F0: true, F1: 32768}
                    _ = fi20
                    _ = v23
                }
                _ = v19
            }
        }
        var v24 int16 = 491
        _ = v24
    }
    v6 = (v6 + 18446744073709551615)
    v10 -= 1.5
    var v25 int32 = 986
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
