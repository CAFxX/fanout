package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int32 }
func genFn0(p0 uint64, p1 string) uint {
    fmt.Println(p1)
    v7 := GenS1{F0: (p1 + "hello"), F1: 0}
    var v8 uint8 = 255
    v9 := GenS2{F0: -509}
    if true {
        if true {
            v10 := []int{-1, 2147483647}
            var v11 int = 9223372036854775807
            _ = v10
            _ = v11
        } else {
            var v12 uint64 = p0
            p0 += v12
            if false {
                fmt.Println(v12, p1)
                v13 := []int{-1, 32768, -483, 0}
                _ = v13
            }
            _ = v12
        }
        var v14 uint8 = 128
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int32, p1 bool, p2 uint) bool {
    p0 += 660
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS3{F0: (p0 & (p0 & p0))}
        v17 := GenS4{F0: 128, F1: 2, F2: -497}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        p0 -= p0
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := []int{479, (-(fi19)), (^(fi19))}
            v21 := GenS3{F0: p0}
            p0 -= p0
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    return (127 == 55)
}

func main() {
    v1 := []string{"a", ("foo bar" + ""), ("hello" + "hello")}
    var v2 float64 = ((3.14159 / ((3.14159 / 0.1) + (1e10 / -2.25))) * (float64(float32(1e10)) * (-2.25 + ((-2.25 / 0.1) - ((3.14159 + 1e10) - -2.25)))))
    var v3 float64 = (0.1 - ((v2 / (v2 * v2)) * (v2 - float64((-2.25 / 3.14159)))))
    var v4 uint8 = 123
    var v5 int64 = -685
    v6 := make(map[int]int)
    v22 := GenS2{F0: 32767}
    fmt.Println(v3)
    for fi23 := 0; fi23 < 6; fi23++ {
        _ = len(v6)
        fmt.Println(fi23, v5, fi23)
        _ = fi23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 bool = (false && ((2 < 0) && true))
        v3 -= -2.25
        _ = fi24
        _ = v25
    }
    v2 = v2
    if (-128 >= -128) {
        v4 = (v4 & 127)
    } else {
        v2 = float64(float32(v3))
        v4 += 255
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v26 int32 = 0
    v2 -= 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
