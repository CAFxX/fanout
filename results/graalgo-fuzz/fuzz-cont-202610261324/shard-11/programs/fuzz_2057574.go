package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int16, p1 int8) uint32 {
    v4 := GenS1{F0: 42}
    var v5 float64 = (0.1 + 3.14159)
    fmt.Println(p0, v5)
    v6 := GenS1{F0: 9223372036854775807}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 729
}
func genFn1(p0 uint64, p1 string) int {
    p0 += (p0 & 0)
    var v8 bool = (!(-1 < 9223372036854775807))
    var v9 int32 = 1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 889
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v10 bool = false
    if v10 {
        v10 = true
    }
    if v10 {
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v10, v10)
            var v12 uint32 = 2147483647
            _ = fi11
            _ = v12
        }
        var v13 uint16 = 564
        var v14 int8 = 0
        _ = v13
        _ = v14
    } else {
        if v10 {
            v15 := []int{127, 42, 128, -208}
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint16 = 2
                for fi18 := 0; fi18 < 5; fi18++ {
                    v17 = v17
                    v19 := GenS2{F0: (!(65535 >= 506)), F1: 42}
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS2{F0: false, F1: -32768}
                switch -2147483648 {
                case 42:
                    var v22 string = ""
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = v15
        } else {
            v10 = v10
            v23 := GenS3{F0: 42}
            v10 = v10
            _ = v23
        }
        _ = len(v2)
        v24 := GenS4{F0: ("a" + (("" + "hello") + "")), F1: v10}
        _ = v24
    }
    fmt.Println(v10)
    var v25 float32 = (-2.25 * float32((3.14159 / (((3.14159 + 0.1) - -2.25) + 1e10))))
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int32 = -730
        fmt.Println(v27, v25)
        _ = fi26
        _ = v27
    }
    if v10 {
        v28 := GenS2{F0: (32767 > -1), F1: -617}
        _ = v28
    }
    v25 = (v25 - v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
}
