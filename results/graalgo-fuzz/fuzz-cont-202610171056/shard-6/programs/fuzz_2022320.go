package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint) bool {
    var v5 int8 = 100
    v5 = -13
    _ = p0
    _ = v5
    return true
}
func genFn1(p0 uint16, p1 uint8, p2 string) bool {
    p0 -= p0
    v6 := []int{-9223372036854775808, 9223372036854775807}
    var v7 int32 = 333
    v7 -= (-(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 uint32 = 577
    var v2 float64 = 0.0
    var v3 int32 = -1
    v4 := make(map[int]bool)
    if ((v3 / ((v3) | 1)) >= (v3 - v3)) {
        var v8 float32 = float32(float64(0.1))
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                for fi10 := 0; fi10 < 18; fi10++ {
                    fmt.Println(fi10, v8)
                    v8 += v8
                    var v11 float64 = 0.0
                    var v12 bool = (!true)
                    v13 := []int{}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if false {
                    v1 = v1
                    v3 += v3
                    v2 = v2
                }
                v14 := []int{fi9}
                _ = v14
            }
            v15 := GenS1{F0: 255}
            v16 := GenS1{F0: 4294967295}
            _ = fi9
            _ = v15
            _ = v16
        }
        v17 := v4[-636]
        _ = v17
        _ = v8
    }
    if (-9223372036854775808 != -13) {
        v18 := []int{2147483647, 14}
        _ = v18
    } else {
        var v19 int8 = -93
        v20 := GenS1{F0: 2}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if false {
            for fi22 := 0; fi22 < 5; fi22++ {
                fmt.Println(v2, fi22, v1)
                fmt.Println(fi21)
                v4[fi21] = (v3 <= v3)
                _ = fi22
            }
            var v23 float32 = -2.25
            _ = v23
        } else {
            fmt.Println(v2)
            v3 = (v3 & v3)
            for fi24 := 0; fi24 < 5; fi24++ {
                v3 += v3
                _ = fi24
            }
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 bool = (!true)
        _ = fi25
        _ = v26
    }
    var v27 uint64 = 1
    fmt.Println(v27, v2)
    fmt.Println(v3, v1)
    for fi28 := 0; fi28 < 3; fi28++ {
        v4[fi28] = false
        _ = fi28
    }
    v29 := GenS2{F0: 2, F1: 128}
    if ((!true) || (!(!((!(false && true)) && true)))) {
        if true {
            v27 = v27
            v30 := GenS3{F0: (false || (220 != 128))}
            for fi31 := 0; fi31 < 18; fi31++ {
                v32 := GenS3{F0: true}
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        fmt.Println(v3)
        if false {
            v2 -= (v2 - 0.1)
        } else {
            if (false || (!true)) {
                var v33 uint = 42
                _ = v33
            }
            v27 = v27
            var v34 int32 = (v3 * -1)
            _ = v34
        }
    }
    v35 := GenS4{F0: 1, F1: (v27 - v27)}
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
