package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0(p0 uint, p1 int8, p2 int16) int32 {
    if false {
        p0 += 668
        fmt.Println(p2, p1)
        p1 += (-(p1))
    }
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := []int{fi5, 9223372036854775807, fi5, -389}
            p1 -= (p1 % ((p1) | 1))
            _ = fi5
            _ = v6
        }
        var v7 uint64 = 42
        v8 := []int{}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    var v4 int16 = 1
    v3 = float64(0.0)
    v9 := v2[("" + ("" + "\n\t"))]
    var v10 float32 = (3.14159 - 1.5)
    v11 := v2["a"]
    for fi12 := 0; fi12 < 4; fi12++ {
        if v11 {
            v13 := GenS1{F0: 1, F1: fi12, F2: 555}
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{}
                v16 := v2[""]
                var v17 uint = 287
                var v18 uint16 = 688
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v13
        } else {
            var v19 uint64 = 1
            for fi20 := 0; fi20 < 9; fi20++ {
                v1 = v1
                _ = fi20
            }
            v2[""] = (553 >= uint32(v19))
            _ = v19
        }
        _ = fi12
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        switch ("hello" + "\n\t") {
        case "":
            fmt.Println(v4, v1)
        case "hello":
            for fi22 := 0; fi22 < 4; fi22++ {
                v4 -= 1
                _ = fi22
            }
        case "a":
            v3 = 3.14159
        }
        _ = fi21
    }
    v23 := GenS2{F0: 42, F1: 559, F2: -1}
    for fi24 := 0; fi24 < 14; fi24++ {
        v10 -= v10
        v25 := []int{(^(fi24)), fi24, (fi24 * 2147483647), (fi24 + (fi24 / ((fi24) | 1)))}
        _ = fi24
        _ = v25
    }
    fmt.Println(v10, v4, v10)
    var v26 int64 = 1
    v27 := []int{-946, 32767}
    var v28 uint64 = 4294967295
    v29 := GenS3{F0: 1, F1: 18446744073709551615, F2: 2}
    var v30 uint32 = 2147483647
    v28 += v28
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
