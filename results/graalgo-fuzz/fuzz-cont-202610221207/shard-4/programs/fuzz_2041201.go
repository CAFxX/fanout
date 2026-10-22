package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) int32 {
    v6 := GenS1{F0: 42, F1: 127, F2: 631}
    p0 -= p0
    switch 2147483647 {
    case 980:
        fmt.Println(p0)
    case 32767:
        var v7 int16 = -32768
        _ = v7
    case 896:
        if false {
            switch "a" {
            case "\n\t":
                fmt.Println(p0)
            case "":
                for fi8 := 0; fi8 < 8; fi8++ {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        v10 := GenS1{F0: 32767, F1: 1, F2: -9223372036854775808}
                        _ = fi9
                        _ = v10
                    }
                    fmt.Println(p0)
                    v11 := GenS2{F0: fi8}
                    _ = fi8
                    _ = v11
                }
            case "hello":
                if (false && true) {
                    v12 := []int{1, 4294967295, 32767, 127}
                    var v13 int64 = 42
                    _ = v12
                    _ = v13
                } else {
                    switch 4294967295 {
                    case 32767:
                        var v14 uint8 = 1
                        _ = v14
                    case -820:
                        var v15 bool = (!true)
                        _ = v15
                    case -1:
                        v16 := []int{-1, 0}
                        _ = v16
                    }
                    fmt.Println(p0)
                    fmt.Println(p0)
                }
            }
            v17 := GenS2{F0: 128}
            _ = v17
        } else {
            p0 -= (p0 * 4294967295)
        }
    }
    var v18 int = 127
    v18 -= v18
    _ = p0
    _ = v18
    _ = v6
    return -2147483648
}
func genFn1(p0 bool, p1 int) int16 {
    v19 := []int{128}
    fmt.Println(p0, p0)
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            v19 = append(v19, 127)
            p0 = p0
            _ = fi21
        }
        var v22 float32 = float32(float64(0.0))
        v19 = append(v19, p1)
        _ = fi20
        _ = v22
    }
    v23 := GenS3{F0: 76}
    var v24 uint16 = 65535
    _ = p0
    _ = p1
    _ = v19
    _ = v23
    _ = v24
    return 42
}

func main() {
    var v1 uint16 = 281
    var v2 int64 = 0
    var v3 string = "hello"
    var v4 uint8 = 100
    v5 := make(map[string]int)
    v25 := []int{-1, -1}
    fmt.Println(v4, v1)
    for fi26 := 0; fi26 < 9; fi26++ {
        fmt.Println(v1, v1)
        _ = fi26
    }
    v1 = 578
    var v27 float32 = 1e10
    var v28 uint8 = 128
    if false {
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 float64 = 0.1
            var v31 float32 = v27
            _ = fi29
            _ = v30
            _ = v31
        }
        v5[((((v3 + ((v3 + v3) + ("\n\t" + ""))) + (v3 + "foo bar")) + v3) + v3)] = 389
    }
    if false {
        v4 -= v4
    }
    v5[(v3 + "hello")] = 0
    var v32 bool = (((^(v1)) <= v1) && (1000 <= -549))
    v32 = v32
    var v33 int8 = -123
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 bool = v32
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
