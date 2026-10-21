package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 bool) uint64 {
    var v5 int16 = 32767
    if (4294967295 == 2147483647) {
        if p0 {
            var v6 string = "foo bar"
            v7 := []int{-9223372036854775808, -2147483648}
            p0 = true
            _ = v6
            _ = v7
        }
        p0 = ("foo bar" != "\n\t")
        p0 = true
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            if (42 <= 2) {
                fmt.Println(p0, fi8, v5)
                v9 := GenS1{F0: 9223372036854775807}
                for fi10 := 0; fi10 < 20; fi10++ {
                    p0 = (0 >= 41)
                    p0 = p0
                    v11 := GenS2{F0: 255}
                    p0 = false
                    var v12 int64 = 42
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            } else {
                var v13 int16 = (v5 * v5)
                _ = v13
            }
            p0 = true
            fmt.Println(p0)
            _ = fi8
        }
        v14 := []int{32768}
        _ = v14
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 int8) uint64 {
    var v15 int16 = 1000
    var v16 int8 = p0
    var v17 int16 = (v15 / ((v15) | 1))
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 152
}

func main() {
    var v1 uint8 = 128
    var v2 int64 = -1
    var v3 float64 = (1.5 - 0.1)
    v4 := make(map[int]int)
    var v18 uint = 128
    for fi19 := 0; fi19 < 7; fi19++ {
        v3 -= 1e10
        _ = fi19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(v3)
        fmt.Println(v1)
        v21 := GenS2{F0: 772}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint16 = 255
        v24 := GenS3{F0: genFn0((-47 <= 1)), F1: v2}
        v1 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v3 = v3
        _ = fi25
    }
    var v26 int64 = (v2 ^ v2)
    fmt.Println(v3)
    _ = len(v4)
    v3 = (v3 * v3)
    if false {
        v3 += -2.25
        fmt.Println(v26, v3, v18)
    } else {
        var v27 bool = false
        _ = v27
    }
    var v28 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
