package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint, p1 int32, p2 int8) int8 {
    if true {
        p2 = -54
        switch "\n\t" {
        case "":
            v6 := []int{32768, 32768}
            _ = v6
        case "foo bar":
            var v7 uint8 = 42
            _ = v7
        }
        p2 = p2
    } else {
        fmt.Println(p2)
        var v8 uint64 = 4294967295
        _ = v8
    }
    if true {
        p1 = p1
        p0 -= (p0 ^ 2)
    } else {
        p1 = 0
        var v9 uint32 = 197
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{0, fi10, fi10}
            var v12 uint = (p0 % (((^(p0))) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    p1 -= int32(p0)
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 uint, p1 int, p2 string) uint8 {
    fmt.Println(p1, p2)
    fmt.Println(p2, p1)
    v13 := []int{-9223372036854775808, (p1 / (((p1 | 855)) | 1))}
    v14 := GenS1{F0: p1}
    var v15 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 72
}

func main() {
    var v1 int = 127
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 int = (-(v1))
    v5 := make(map[int]bool)
    var v16 string = "foo bar"
    var v17 bool = true
    for fi18 := 0; fi18 < 17; fi18++ {
        if v17 {
            v1 += v4
            var v19 int64 = 540
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS1{F0: -9223372036854775808}
                v17 = v17
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: fi18}
            _ = v22
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            if (v17 && ((2 < 526) && v17)) {
                if (3.14159 != (3.14159 + 0.0)) {
                    v24 := []int{(-(fi23)), (fi23 % (((v4 << ((42) & 63))) | 1)), v4, fi18}
                    _ = v24
                }
                v25 := GenS1{F0: 9223372036854775807}
                v2[226] = (("a" + v16) + v16)
                _ = v25
            }
            v26 := GenS1{F0: 80}
            var v27 uint = 4294967295
            v16 = "a"
            var v28 string = v16
            v29 := v2[4294967295]
            v27 += 889
            var v30 int = fi18
            v28 = "a"
            if true {
                v31 := GenS2{F0: 898, F1: (484 >= uint64(v27))}
                _ = v31
            }
            fmt.Println(v16, fi23, v27)
            _ = fi23
            _ = v26
            _ = v27
            _ = v28
            _ = v29
            _ = v30
        }
        var v32 uint32 = 2147483647
        _ = fi18
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v17 = v17
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 uint32 = 65535
            v36 := GenS1{F0: (fi34 * 1)}
            var v37 float64 = 0.0
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi33
    }
    fmt.Println(v4)
    v38 := GenS2{F0: 4294967295, F1: (-107 != 127)}
    v17 = v17
    var v39 float64 = float64(-2.25)
    for fi40 := 0; fi40 < 3; fi40++ {
        v39 = float64(-2.25)
        v41 := []int{(fi40 | (fi40 / ((v1) | 1))), (fi40 << ((2) & 63)), v1, fi40}
        _ = fi40
        _ = v41
    }
    var v42 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
