package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32; F1 uint32; F2 string }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 int8, p2 uint32) int64 {
    var v7 int32 = 1
    if true {
        var v8 uint16 = 32767
        if true {
            if (float32((-2.25 + 1.5)) > -2.25) {
                var v9 uint8 = 127
                _ = v9
            } else {
                var v10 uint64 = 746
                _ = v10
            }
        }
        if (1000 >= -1) {
            var v11 uint32 = (p0 & p0)
            _ = v11
        } else {
            v12 := GenS1{F0: ((-1 >= 42) || false), F1: -803, F2: ("\n\t" + ("" + "\n\t"))}
            _ = v12
        }
        _ = v8
    }
    var v13 uint32 = 2147483648
    v14 := GenS1{F0: true, F1: 1000, F2: ("a" + "hello")}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]int)
    var v3 uint32 = 701
    var v4 uint16 = 294
    v5 := make(map[string]string)
    var v6 uint8 = 128
    var v15 uint64 = 0
    var v16 string = ""
    v17 := v2[(v1 / ((v1) | 1))]
    var v18 uint64 = (v15 / ((v15) | 1))
    v19 := GenS2{F0: 686, F1: -88}
    v16 = (((v16 + "\n\t") + "") + (v16 + ((v16 + "hello") + v16)))
    var v20 int32 = 42
    v21 := GenS3{F0: (v3 - (v3 & v3)), F1: v3, F2: "hello"}
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            if (((^(v4)) < (v4 << ((0) & 63))) || ((1e10 * (float64(0.0) + float64(3.14159))) > ((0.1 - 1.5) + ((float64((0.1 / 0.1)) / 1.0) * ((1.5 / 1e10) + (0.0 * (0.1 / 0.1))))))) {
                fmt.Println(v1, v6)
            }
            _ = fi23
        }
        v4 += 0
        _ = fi22
    }
    var v24 bool = true
    v25 := GenS4{F0: ((^(v20)) != v20)}
    switch "a" {
    case "foo bar":
        if v24 {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := []int{}
                v4 = 1
                _ = fi26
                _ = v27
            }
            v28 := GenS1{F0: v24, F1: -32768, F2: v16}
            for fi29 := 0; fi29 < 15; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    var v31 string = (v16 + ("" + "\n\t"))
                    _ = fi30
                    _ = v31
                }
                v20 -= 1
                _ = fi29
            }
            _ = v28
        }
    case "a":
        fmt.Println(v4)
    case "":
        v32 := GenS3{F0: (v3 >> ((2) & 63)), F1: 454, F2: v16}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
