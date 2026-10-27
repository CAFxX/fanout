package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 int32) uint16 {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 float32 = 0.1
        if false {
            var v7 int8 = 1
            v8 := GenS1{F0: v7}
            _ = v7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    if false {
        if (32767 >= 2) {
            var v9 uint32 = 4294967295
            p0 -= p0
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 string = "\n\t"
                for fi12 := 0; fi12 < 13; fi12++ {
                    if (4294967295 != 42) {
                        var v13 uint64 = 0
                        var v14 int = -678
                        _ = v13
                        _ = v14
                    }
                    switch "" {
                    case "hello":
                        if (float64(1e10) == ((3.14159 * (((float64(1.5) / 0.1) / float64(0.1)) / float64(((1.5 / (1e10 - (3.14159 / 1e10))) / (0.1 / (3.14159 + ((float32((-2.25 + 3.14159)) / 3.14159) - 3.14159))))))) - -2.25)) {
                            fmt.Println(p1, fi10)
                        } else {
                            v15 := []int{(fi12 << (((p0 - 2)) & 63)), (-(fi12)), 1, fi10}
                            _ = len(v15)
                            fmt.Println(v9, fi10, v11)
                            _ = v15
                        }
                    case "\n\t":
                        v16 := []int{9223372036854775807, fi12, 42, (-(fi12))}
                        _ = v16
                    case "a":
                        fmt.Println(v9, fi12)
                    }
                    _ = fi12
                }
                v17 := GenS2{F0: 128}
                _ = fi10
                _ = v11
                _ = v17
            }
            _ = v9
        }
        v18 := []int{1}
        var v19 int16 = -32768
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: uint8(p0)}
        fmt.Println(p0, p0)
        for fi21 := 0; fi21 < 17; fi21++ {
            p0 = (p0 + 705)
            v22 := GenS1{F0: -90}
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 297
}
func genFn1(p0 uint8, p1 int32) int32 {
    var v23 float64 = (3.14159 + (-2.25 + (float64((float32(float64(0.0)) + 0.0)) + float64(0.0))))
    v24 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 2147483647
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = -43
    v3 := make(map[string]uint32)
    var v4 float64 = float64(1.5)
    if false {
        var v25 int = -739
        if (42 != 4294967295) {
            v26 := GenS1{F0: (v2 - 1)}
            _ = v26
        }
        _ = v25
    } else {
        v27 := []int{}
        _ = v27
    }
    if (false && ((((!((float32(v4) * -2.25) != (3.14159 - 1.5))) || (false || false)) || true) || (255 < 2))) {
        v28 := []int{-845, 42, 42, 2147483647}
        v1 = v1
        _ = v28
    } else {
        var v29 int32 = -1
        _ = v29
    }
    v4 = v4
    v4 = v4
    var v30 string = "hello"
    v31 := GenS2{F0: 133}
    switch 386 {
    case -881:
        v1 = v1
    default:
        v4 = ((((-2.25 * ((v4 - (v4 / v4)) - (3.14159 + 3.14159))) - 3.14159) - v4) / ((0.1 - v4) / float64(0.1)))
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        if (42 < 0) {
            v1 -= genFn0(0, -1)
            v33 := GenS1{F0: v2}
            _ = v33
        } else {
            v2 = (v2 % (((v2 >> ((255) & 63))) | 1))
            v30 = v30
            v2 = v2
        }
        v34 := GenS2{F0: 209}
        _ = fi32
        _ = v34
    }
    var v35 int64 = 0
    var v36 int16 = 55
    v1 -= 1
    v37 := []int{-458}
    var v38 uint64 = 165
    if true {
        var v39 bool = ((v36 ^ (v36 * (v36 ^ (^(v36))))) <= v36)
        v40 := GenS3{F0: v38, F1: -279, F2: 4294967295}
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
