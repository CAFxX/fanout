package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0() int {
    v6 := GenS1{F0: -1, F1: 2}
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7)
        fmt.Println(fi7)
        var v8 int32 = -663
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 0, F1: -2147483648}
    _ = v6
    _ = v9
    return 128
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 uint = 4294967295
    if (955 != v5) {
        v5 = (v5 + (v5 ^ (v5 & v5)))
        var v10 uint32 = 2147483647
        v5 = 0
        _ = v10
    } else {
        fmt.Println(v5, v5)
    }
    if true {
        var v11 uint16 = 255
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                var v13 int8 = 70
                _ = v13
            } else {
                v1 = v1
                var v14 uint8 = 1
                _ = v14
            }
            _ = fi12
        }
        var v15 int16 = 307
        _ = v11
        _ = v15
    }
    var v16 int64 = 592
    v17 := []int{genFn0(), -9223372036854775808, 562, 127}
    _ = len(v3)
    if false {
        var v18 float32 = -2.25
        _ = v18
    } else {
        fmt.Println(v5)
        if true {
            var v19 int64 = 1
            _ = v19
        } else {
            v20 := GenS3{F0: -1, F1: "\n\t", F2: int16(v16)}
            _ = v20
        }
        v21 := GenS4{F0: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := v4[32767]
                v25 := GenS5{F0: 4294967295, F1: 464, F2: fi23}
                fmt.Println(v1, v16)
                _ = fi23
                _ = v24
                _ = v25
            }
            if false {
                var v26 uint16 = 255
                fmt.Println(fi22)
                _ = v26
            } else {
                var v27 bool = true
                _ = v27
            }
        } else {
            fmt.Println(v5)
            var v28 uint = v5
            _ = v28
        }
        var v29 uint16 = 42
        if true {
            v30 := GenS6{F0: (^(v29)), F1: v16}
            switch "a" {
            case "a":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case "":
                var v31 int = fi22
                _ = v31
            }
            _ = v30
        }
        _ = fi22
        _ = v29
    }
    v16 = 0
    v32 := GenS3{F0: -19, F1: "foo bar", F2: -32768}
    v1 += 1e10
    fmt.Println(v1, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
