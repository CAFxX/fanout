package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 int }
func genFn0() uint16 {
    var v7 uint64 = 2
    v7 += (v7 & (v7 + v7))
    var v8 int = 9223372036854775807
    _ = v7
    _ = v8
    return 697
}
func genFn1(p0 string, p1 string) uint32 {
    var v9 float32 = ((((((float32((((1e10 / (((-2.25 - float64(0.0)) / 1.0) + (((-2.25 + (-2.25 - 1e10)) + 1.5) * ((float64(float32(0.1)) - ((3.14159 / -2.25) + -2.25)) + (0.0 + 3.14159))))) / (-2.25 / 3.14159)) + (0.0 + 1.5))) - 3.14159) + float32((3.14159 * (-2.25 / 3.14159)))) + 1.5) / 0.1) * (0.1 - ((float32(1.5) / 0.1) / 1.5))) - -2.25)
    var v10 uint = 325
    switch ((("" + p0) + (p0 + "foo bar")) + p0) {
    case "foo bar":
        var v11 uint8 = 255
        _ = v11
    default:
        var v12 bool = (!true)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 94
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    var v3 int8 = -128
    var v4 int8 = (v3 / (((v3 + (-(v3)))) | 1))
    var v5 int8 = v3
    v6 := []int{4294967295, 2147483647, -623, 42}
    var v13 string = "a"
    v14 := GenS1{F0: 1, F1: 1}
    v6 = append(v6, 128)
    switch 0 {
    case -916:
        v15 := GenS1{F0: 32767, F1: 2147483648}
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: 81, F1: 42}
        _ = fi16
        _ = v17
    }
    v5 = v4
    fmt.Println(v5, v4)
    var v18 uint16 = 295
    var v19 uint16 = v18
    if true {
        fmt.Println(v1, v13, v18)
        fmt.Println(v4)
        var v20 int16 = 1
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            if true {
                if false {
                    v19 = (^(v18))
                    v19 = 255
                }
                if (false || false) {
                    fmt.Println(fi21, fi21, v18)
                    v13 = "a"
                } else {
                    switch int(v1) {
                    case -350:
                        var v22 float64 = 1.5
                        _ = v22
                    case -1:
                        v3 = v3
                    case 4294967295:
                        v3 = v5
                    }
                }
                v2[fi21] = (fi21 / ((fi21) | 1))
            }
            _ = fi21
        }
    }
    fmt.Println(v13, v1, v18)
    for fi23 := 0; fi23 < 8; fi23++ {
        if (v18 > (^(v19))) {
            v24 := GenS3{F0: -776, F1: fi23}
            _ = v24
        }
        if (127 != uint8(v19)) {
            fmt.Println(v4)
        }
        v25 := GenS3{F0: 1, F1: 4294967295}
        _ = fi23
        _ = v25
    }
    var v26 int64 = 42
    for fi27 := 0; fi27 < 6; fi27++ {
        fmt.Println(v1, v18, v18)
        var v28 uint64 = uint64(v19)
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
