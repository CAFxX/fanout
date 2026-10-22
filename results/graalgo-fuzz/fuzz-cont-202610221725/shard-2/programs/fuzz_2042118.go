package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0(p0 int64) uint16 {
    if (!false) {
        fmt.Println(p0)
        switch "foo bar" {
        case "":
            var v7 uint8 = 18
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS1{F0: 255, F1: 4294967295}
                v10 := GenS1{F0: 0, F1: 0}
                var v11 int8 = 1
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        var v12 string = "hello"
        _ = v12
    } else {
        fmt.Println(p0)
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 int = fi13
            v15 := GenS2{F0: 0, F1: 1}
            v16 := GenS3{F0: "a"}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 uint64 = 0
        _ = v17
    }
    var v18 float32 = ((0.0 * float32(((0.0 * float64(((-2.25 + float32(0.0)) / 1e10))) + ((1e10 + (float64(-2.25) / float64(1.5))) / (-2.25 + -2.25))))) + -2.25)
    v19 := []int{42, -747, 42, 2147483647}
    if true {
        switch "hello" {
        case "a":
            p0 -= p0
        }
        v18 -= v18
        for fi20 := 0; fi20 < 20; fi20++ {
            p0 = -58
            _ = fi20
        }
    }
    switch 32768 {
    case 2147483647:
        v18 = (1e10 - v18)
    case 128:
        v18 = 0.0
    case 4294967295:
        var v21 bool = false
        _ = v21
    default:
        for fi22 := 0; fi22 < 10; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        if len(v19) > 0 {
                            v19[0] = (-(fi23))
                        }
                    case "":
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    }
                } else {
                    if ((false || true) && ((fi23 >> ((2) & 63)) != (fi23 << ((1) & 63)))) {
                        var v24 uint8 = 206
                        _ = v24
                    }
                    if false {
                        fmt.Println(fi22, v18)
                        v19 = append(v19, fi23)
                    }
                    v25 := GenS4{F0: (fi23 % ((fi23) | 1)), F1: 4294967295}
                    _ = v25
                }
                if (!true) {
                    var v26 uint = 0
                    var v27 int64 = p0
                    _ = v26
                    _ = v27
                } else {
                    if true {
                        var v28 string = (("a" + "a") + ("foo bar" + ("\n\t" + ("hello" + ""))))
                        fmt.Println(v28, p0, v28)
                        _ = v28
                    }
                }
                p0 = 1
                fmt.Println(fi23, fi22)
                _ = fi23
            }
            _ = fi22
        }
    }
    _ = p0
    _ = v18
    _ = v19
    return 255
}

func main() {
    var v1 int32 = 0
    v2 := []string{}
    var v3 int16 = -1
    v4 := make(map[int]int)
    var v5 bool = true
    v6 := make(map[int]string)
    v29 := GenS5{F0: 127, F1: 2, F2: 0}
    v1 += v1
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 = 42
        fmt.Println(v5, fi30)
        fmt.Println(v3, v3)
        _ = fi30
    }
    v31 := GenS4{F0: 32767, F1: 4294967295}
    fmt.Println(v3, v5, v5)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v3 += v3
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
