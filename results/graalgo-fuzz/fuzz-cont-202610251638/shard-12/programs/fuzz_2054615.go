package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 uint64, p2 uint64) uint {
    v6 := GenS1{F0: 621}
    v7 := []int{(p0 + (-(p0))), 127, p0, p0}
    if len(v7) > 0 {
        v7[0] = 881
    }
    v8 := GenS1{F0: 2}
    var v9 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (^(v9))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := []uint32{}
    v4 := []uint32{42}
    var v5 uint = 172
    v5 += (v5 << ((v5) & 63))
    var v10 float64 = 3.14159
    delete(v1, "")
    var v11 uint = (^(v5))
    v12 := v1["foo bar"]
    v13 := GenS1{F0: 2}
    if false {
        v14 := GenS2{F0: 42, F1: 65535}
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int64 = 577
            _ = fi15
            _ = v16
        }
        v17 := GenS3{F0: 1}
        _ = v14
        _ = v17
    }
    var v18 uint8 = 127
    if ((v5 > v11) || (!((0.1 / ((0.1 - float32((v10 - 3.14159))) + (0.1 - 0.1))) >= float32(0.1)))) {
        v19 := v1[("a" + ("" + "foo bar"))]
        _ = v19
    } else {
        fmt.Println(v10)
        if false {
            v20 := v1["a"]
            _ = v20
        } else {
            for fi21 := 0; fi21 < 3; fi21++ {
                v3 = append(v3, (v12 - v12))
                var v22 uint32 = 2147483648
                for fi23 := 0; fi23 < 18; fi23++ {
                    v11 += (v5 % (((v11 / ((v11) | 1))) | 1))
                    v12 += (v22 + (^(v12)))
                    v10 = v10
                    _ = fi23
                }
                _ = fi21
                _ = v22
            }
            v10 = (3.14159 + float64(-2.25))
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int32 = -2147483648
        v26 := GenS1{F0: 32767}
        v10 = v10
        _ = fi24
        _ = v25
        _ = v26
    }
    v12 = (v12 + (v12 / (((v12 << ((v5) & 63))) | 1)))
    var v27 int32 = 918
    fmt.Println(v11, v27, v18)
    var v28 float32 = (1.5 + 1e10)
    for fi29 := 0; fi29 < 8; fi29++ {
        v12 -= (v12 - v12)
        _ = fi29
    }
    if (!(0 < -914)) {
        fmt.Println(v10, v11)
        if true {
            if false {
                if false {
                    v28 = v28
                    var v30 int8 = 100
                    var v31 uint8 = (v18 + v18)
                    _ = v30
                    _ = v31
                } else {
                    var v32 int8 = -99
                    var v33 uint32 = v12
                    _ = v32
                    _ = v33
                }
            } else {
                _ = len(v4)
                var v34 int32 = -2147483648
                _ = v34
            }
            var v35 string = "hello"
            for fi36 := 0; fi36 < 9; fi36++ {
                v28 = v28
                fmt.Println(v12, v10)
                if true {
                    var v37 float64 = float64(v28)
                    _ = v37
                }
                _ = fi36
            }
            _ = v35
        }
    }
    v11 += 210
    var v38 float64 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
