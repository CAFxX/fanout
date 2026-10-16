package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint8; F2 uint32 }
func genFn0(p0 uint32) int32 {
    var v6 float64 = 3.14159
    p0 = (p0 << ((42) & 63))
    fmt.Println(p0, v6)
    v6 = (v6 / float64(3.14159))
    var v7 int64 = 1
    _ = p0
    _ = v6
    _ = v7
    return 440
}
func genFn1(p0 uint64, p1 uint32, p2 bool) int16 {
    var v8 float64 = 1e10
    v9 := GenS1{F0: (^(p1)), F1: "foo bar", F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 52
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    var v5 int = 1
    var v10 int32 = 2147483647
    var v11 bool = true
    v10 += (v10 >> ((44) & 63))
    fmt.Println(v10, v5)
    var v12 uint32 = 65535
    v13 := []int{v5, 37, v5, 32767}
    var v14 int = 0
    if v11 {
        v11 = v11
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int16 = -440
            var v17 bool = ((^(fi15)) > v14)
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                v11 = ((1e10 <= -2.25) || true)
                if true {
                    var v20 float64 = 1.5
                    var v21 string = (("foo bar" + "\n\t") + "")
                    _ = v20
                    _ = v21
                } else {
                    v14 = (v14 ^ -920)
                }
                v5 = v5
                _ = fi19
            }
            _ = fi18
        }
    }
    v14 += -9223372036854775808
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(fi22, fi22, v12)
        for fi23 := 0; fi23 < 9; fi23++ {
            v14 = (v5 + v14)
            var v24 float64 = -2.25
            _ = fi23
            _ = v24
        }
        fmt.Println(v5, v12)
        _ = fi22
    }
    if true {
        var v25 float32 = float32(((((-2.25 - 3.14159) * 1e10) + (float64(-2.25) * ((0.1 + -2.25) * 1e10))) + (float64(float32(((1.5 - 3.14159) * ((((1.5 + 3.14159) + 0.0) / 1.0) / (((float64(((((((float32(-2.25) + (1e10 - 3.14159)) - (((1e10 / 1.5) + 0.0) + float32(1e10))) * 0.0) / 1.5) - ((float32(0.1) / 1.0) * 0.1)) * -2.25)) / 3.14159) * ((1.5 / 0.1) + 0.0)) + (-2.25 / -2.25)))))) + ((0.0 + 0.1) * 1.5))))
        var v26 int64 = -9223372036854775808
        _ = v25
        _ = v26
    } else {
        v14 = 127
        var v27 uint = 255
        v11 = (-1 != 127)
        _ = v27
    }
    var v28 uint = 2
    var v29 uint8 = 40
    v30 := v2["a"]
    if (!(v10 <= genFn0(v30))) {
        var v31 int8 = 100
        _ = v31
    } else {
        if true {
            v10 -= (v10 ^ (-(v10)))
        } else {
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := GenS2{F0: uint(v12)}
                v34 := GenS3{F0: fi32, F1: v29, F2: (v30 * v12)}
                for fi35 := 0; fi35 < 14; fi35++ {
                    v10 = (v10 / ((2147483647) | 1))
                    _ = fi35
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        switch v5 {
        case 2147483647:
            v28 = v28
        }
        var v36 int32 = v10
        _ = v36
    }
    _ = len(v4)
    var v37 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
}
