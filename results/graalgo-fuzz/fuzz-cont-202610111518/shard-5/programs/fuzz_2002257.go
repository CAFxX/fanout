package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int16; F2 uint }
type GenS4 struct { F0 int64; F1 string; F2 int8 }
func genFn0(p0 string, p1 uint8) int16 {
    p0 = p0
    switch 798 {
    case 42:
        if (42 >= 1000) {
            switch 0 {
            case -9223372036854775808:
                for fi5 := 0; fi5 < 6; fi5++ {
                    p1 = p1
                    _ = fi5
                }
            case 127:
                if false {
                    for fi6 := 0; fi6 < 12; fi6++ {
                        var v7 uint16 = 2
                        v8 := GenS1{F0: 42}
                        v9 := []int{fi6, (fi6 << ((255) & 63)), fi6, fi6}
                        _ = fi6
                        _ = v7
                        _ = v8
                        _ = v9
                    }
                }
            case 395:
                fmt.Println(p1, p0)
            default:
                p0 = (p0 + ("hello" + p0))
            }
            v10 := []int{128, 1, 4294967295, 92}
            _ = v10
        } else {
            if ("foo bar" < "") {
                v11 := GenS1{F0: 205}
                var v12 int16 = 32767
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: -1}
            p0 = ""
            _ = v13
        }
    case 4294967295:
        var v14 uint64 = 18446744073709551615
        _ = v14
    default:
        v15 := []int{}
        _ = v15
    }
    var v16 uint64 = 2
    _ = p0
    _ = p1
    _ = v16
    return 1
}
func genFn1(p0 uint32, p1 string) int32 {
    var v17 float64 = -2.25
    v18 := []int{412, 32767}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 2147483647
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    if true {
        var v19 int16 = -32768
        _ = v19
    }
    var v20 uint = 4294967295
    var v21 uint32 = 2147483647
    v1 = v1
    if (42 > 2) {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = 0.1
            var v24 bool = true
            _ = fi22
            _ = v23
            _ = v24
        }
        if (false || false) {
            v20 = (v20 ^ (v20 + 660))
        }
    } else {
        v21 = v21
        var v25 bool = true
        _ = v25
    }
    v26 := GenS2{F0: 0, F1: 127}
    if false {
        switch 32767 {
        case 532:
            fmt.Println(v20, v21, v1)
        case 0:
            for fi27 := 0; fi27 < 4; fi27++ {
                v21 -= 2
                if true {
                    if false {
                        if ((v1 * float32(1e10)) == (v1 + (1.5 - -2.25))) {
                            var v28 uint64 = 42
                            _ = v28
                        }
                        _ = len(v3)
                        v29 := GenS2{F0: -1, F1: 1}
                        _ = v29
                    }
                } else {
                    if (169 > 9223372036854775807) {
                        var v30 int8 = 42
                        fmt.Println(v20, fi27)
                        _ = v30
                    }
                    var v31 float32 = ((v1 + ((((v1 - (0.1 / 0.1)) / (float32(-2.25) - ((v1 + v1) * ((v1 - v1) / 0.1)))) + -2.25) * 0.1)) + v1)
                    _ = v31
                }
                _ = fi27
            }
        case 128:
            v32 := GenS3{F0: 9223372036854775807, F1: 1, F2: 42}
            _ = v32
        }
    } else {
        v2[(("\n\t" + "a") + "hello")] = 57
    }
    var v33 float32 = (-2.25 - 0.1)
    _ = len(v3)
    v34 := GenS1{F0: 42}
    v20 = (v20 / (((v20 << (((^(v20))) & 63))) | 1))
    v35 := []int{2147483647, 848, 632, 0}
    v36 := GenS4{F0: 0, F1: "", F2: 100}
    var v37 bool = (-870 == 1)
    for fi38 := 0; fi38 < 3; fi38++ {
        fmt.Println(v20)
        var v39 bool = (-1 == -158)
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
