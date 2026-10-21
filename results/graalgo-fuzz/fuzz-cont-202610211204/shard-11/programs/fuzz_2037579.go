package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint64 }
func genFn0(p0 int8, p1 int) uint16 {
    switch (("\n\t" + "foo bar") + "hello") {
    case "a":
        if (!false) {
            for fi5 := 0; fi5 < 15; fi5++ {
                if (!(true || (!false))) {
                    var v6 uint = 2
                    v7 := GenS1{F0: 42, F1: 42, F2: (p1 >> (((v6 << ((89) & 63))) & 63))}
                    var v8 int8 = p0
                    _ = v6
                    _ = v7
                    _ = v8
                }
                p1 = 9223372036854775807
                _ = fi5
            }
        }
    case "":
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: ((42 == 0) && false), F1: -2147483648}
            _ = fi9
            _ = v10
        }
    case "foo bar":
        v11 := []int{269, p1, 32768, p1}
        _ = v11
    }
    var v12 int64 = 42
    var v13 uint32 = 85
    fmt.Println(p1, v12)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 int = 127
    v2 := make(map[string]int)
    var v3 uint64 = 886
    v4 := make(map[int]bool)
    var v14 float64 = (1e10 * 1.5)
    if false {
        var v15 string = (("" + "foo bar") + ("hello" + "hello"))
        _ = v15
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            v14 = (v14 / 1.0)
            _ = fi16
        }
        if true {
            var v17 float32 = (0.1 / 1.0)
            _ = v17
        }
    }
    v1 = (v1 & v1)
    v14 = v14
    v18 := []int{(v1 | (-(v1)))}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if len(v18) > 0 {
        v18[0] = (v1 * (v1 | (v1 << ((4294967295) & 63))))
    }
    fmt.Println(v1)
    _ = len(v4)
    if (32767 <= -32768) {
        var v19 uint64 = v3
        var v20 int = (v1 + v1)
        _ = v19
        _ = v20
    }
    var v21 uint = 329
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v3, fi22)
        for fi23 := 0; fi23 < 8; fi23++ {
            v1 = -9223372036854775808
            var v24 int = 32767
            var v25 int8 = 100
            v25 = -120
            switch "foo bar" {
            case "hello":
                v14 = float64(0.1)
            case "":
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            case "a":
                v26 := GenS3{F0: 42, F1: (v21 | (v21 * 65535)), F2: v3}
                _ = v26
            default:
                var v27 int8 = -128
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    fmt.Println(v21, v3, v21)
    var v28 uint16 = 513
    v3 += 1
    v18 = append(v18, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}
