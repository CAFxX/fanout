package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint, p1 int32) uint {
    v6 := GenS1{F0: 18446744073709551615, F1: -1, F2: 42}
    p0 = p0
    v7 := []int{2147483647, 1, -2147483648}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: 42, F2: 1}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 int16) int32 {
    fmt.Println(p0)
    if false {
        var v10 float32 = float32(((-2.25 / (0.0 - (0.1 - 1e10))) / 3.14159))
        _ = v10
    } else {
        var v11 uint = 18
        _ = v11
    }
    v12 := GenS2{F0: 18446744073709551615, F1: 42}
    _ = p0
    _ = v12
    return 42
}

func main() {
    var v1 bool = (2 == 44)
    var v2 bool = v1
    v3 := make(map[string]string)
    v4 := []string{"hello", "\n\t", "hello", "foo bar"}
    v5 := make(map[uint32]int)
    fmt.Println(v1, v2)
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    v2 = v2
    var v13 float32 = 1e10
    v2 = false
    var v14 uint16 = 39
    v15 := v5[65535]
    var v16 float64 = (1e10 - (0.1 / (float64((float32(-2.25) / v13)) * 1.5)))
    v15 -= 128
    v13 += (0.0 - (float32(-2.25) + -2.25))
    var v17 float32 = (((v13 / v13) * v13) - v13)
    var v18 int32 = genFn1(42)
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v2)
        _ = fi19
    }
    var v20 uint16 = v14
    var v21 uint32 = 2147483647
    if v1 {
        switch 4294967295 {
        case 32768:
            if false {
                var v22 uint = uint(v21)
                for fi23 := 0; fi23 < 1; fi23++ {
                    delete(v5, v21)
                    v2 = (v2 || true)
                    _ = fi23
                }
                _ = v22
            } else {
                var v24 uint8 = 53
                _ = len(v4)
                _ = v24
            }
        case 42:
            for fi25 := 0; fi25 < 15; fi25++ {
                for fi26 := 0; fi26 < 4; fi26++ {
                    v1 = ((!(v20 <= (^(v20)))) && (100 == -1))
                    v27 := GenS3{F0: v18, F1: (v14 + 0)}
                    _ = fi26
                    _ = v27
                }
                v1 = v2
                var v28 uint64 = 18446744073709551615
                _ = fi25
                _ = v28
            }
        }
    }
    if false {
        v14 -= 65535
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 float64 = (3.14159 * v16)
            _ = fi29
            _ = v30
        }
    } else {
        if (805 != 251) {
            delete(v5, 2147483648)
        }
    }
    fmt.Println(v16, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
