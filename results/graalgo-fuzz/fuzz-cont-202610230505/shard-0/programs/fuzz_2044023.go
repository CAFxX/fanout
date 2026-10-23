package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 uint8, p2 uint64) uint32 {
    var v4 int32 = 83
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 2; fi6++ {
        p0 -= p0
        _ = fi6
    }
    if (true || ((1.5 - 0.0) == (float32(((-2.25 - 3.14159) / 0.1)) + (((1e10 * ((float32(1e10) - 3.14159) - (3.14159 + 1.5))) / (3.14159 * -2.25)) / 0.1)))) {
        var v7 uint32 = 42
        switch "\n\t" {
        case "\n\t":
            v8 := GenS1{F0: 328, F1: int64(v4), F2: 1000}
            _ = v8
        }
        v9 := GenS1{F0: 1000, F1: 1, F2: -762}
        _ = v7
        _ = v9
    } else {
        var v10 bool = (69 > int8(v5))
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 5
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 int8 = 100
    switch "a" {
    case "foo bar":
        v3 = v3
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: 42, F1: 1000, F2: -97}
            v3 -= v3
            var v13 int32 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int64 = -1
    var v15 float32 = 1.5
    for fi16 := 0; fi16 < 20; fi16++ {
        v14 = -533
        var v17 uint8 = 42
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 -= (v3 / ((v3) | 1))
            v3 = 42
            fmt.Println(v15, v17)
            var v19 int64 = 9223372036854775807
            v15 -= float32(0.0)
            var v20 float64 = ((3.14159 * float64(float32((((1e10 / (-2.25 - (0.1 + 3.14159))) * 1e10) / float64((v15 * (v15 / (0.0 * v15)))))))) + (0.1 * ((1e10 / ((1.5 - 0.1) + (((1e10 + (1e10 / -2.25)) / 1.0) * 1e10))) / -2.25)))
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int = 32767
        if false {
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int8 = 127
                _ = fi23
                _ = v24
            }
        } else {
            v22 = -1
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := GenS2{F0: 255, F1: -690, F2: (v3 | int8(v14))}
                v27 := v1[(v22 ^ fi25)]
                v15 += v15
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := []int{32768, (v22 << ((0) & 63)), int(v14)}
            _ = v28
        }
        _ = fi21
        _ = v22
    }
    v29 := []int{42, 2147483647}
    var v30 float32 = ((-2.25 + v15) - ((1.5 * (v15 * v15)) - ((float32((float64((1.5 - v15)) * float64(v15))) / v15) / (float32(((0.1 * ((1.5 / (3.14159 / (1e10 / 1.5))) - 0.0)) / (((1.5 + float64(3.14159)) / ((float64(v15) + (float64(v15) + 0.0)) - (((float64(v15) - 1.5) - 1.5) + 3.14159))) - 0.1))) / 0.1))))
    v29 = append(v29, 1)
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 int64 = -9223372036854775808
        switch 736 {
        case 0:
            v15 -= (float32(1e10) + ((1e10 + v15) * (v15 - 1.5)))
        case 127:
            fmt.Println(fi31, fi31, v30)
        }
        v33 := v2[genFn0((-(v32)), 88, 513)]
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float32 = (float32(1.5) / 3.14159)
    if false {
        v35 := GenS3{F0: 100}
        fmt.Println(v30, v3, v3)
        _ = v35
    }
    if false {
        var v36 uint8 = 1
        _ = v36
    }
    var v37 string = ("hello" + "foo bar")
    for fi38 := 0; fi38 < 8; fi38++ {
        v29 = append(v29, fi38)
        _ = fi38
    }
    v30 = (float32((((1e10 * 0.0) * (((3.14159 - -2.25) + 0.0) - (-2.25 + (float64(float32(1.5)) / (-2.25 - float64(v15)))))) + -2.25)) + v30)
    v3 = -128
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 bool = false
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v37
}
