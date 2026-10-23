package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0() int32 {
    var v5 bool = ((!true) || (!(1 <= 2147483647)))
    var v6 uint64 = 454
    v6 -= v6
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 int64) int {
    fmt.Println(p0)
    var v7 int32 = genFn0()
    if (!(!(42 == -1))) {
        if (1 > int32(p0)) {
            v8 := []int{305, 1, 559}
            _ = v8
        }
        var v9 string = ("hello" + "")
        _ = v9
    } else {
        v7 -= (^(v7))
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]int)
    v3 := []uint32{0, 386, 2}
    var v4 uint = 4294967295
    v4 = (v4 & v4)
    var v10 bool = (!false)
    if v10 {
        var v11 int32 = 1
        switch 42 {
        case 42:
            var v12 bool = (v10 && false)
            _ = v12
        case 32768:
            switch "foo bar" {
            case "hello":
                switch -2147483648 {
                case -1:
                    var v13 bool = (0.0 >= 1.5)
                    _ = v13
                case 42:
                    v11 -= v11
                case 1:
                    if v10 {
                        _ = len(v2)
                    } else {
                        v10 = (32 >= -101)
                        v4 -= uint(v1)
                        var v14 int = genFn1(-9223372036854775808)
                        _ = v14
                    }
                }
            }
        }
        _ = v11
    } else {
        v15 := GenS1{F0: 0, F1: 488, F2: 2}
        _ = v15
    }
    v10 = v10
    if v10 {
        var v16 float64 = 0.1
        _ = v16
    }
    v10 = (v10 && false)
    _ = len(v3)
    var v17 float64 = (((1e10 - 0.1) * 0.1) / float64(1e10))
    switch "a" {
    case "foo bar":
        var v18 float32 = (float32((v17 / 1.0)) + 1.5)
        _ = v18
    case "":
        var v19 int = -1
        _ = v19
    case "\n\t":
        v20 := GenS2{F0: 0, F1: 1}
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := GenS2{F0: -1, F1: -32768}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            if (!v10) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    fmt.Println(v1)
                    _ = fi24
                }
                var v25 uint64 = (v1 << (((v4 / (((v4 * (v4 + (v4 & (^(v4)))))) | 1))) & 63))
                _ = v25
            }
            v26 := GenS3{F0: (fi23 + fi23), F1: (v1 + 42), F2: (^(v1))}
            if true {
                _ = len(v3)
                var v27 string = "hello"
                _ = v27
            }
            _ = fi23
            _ = v26
        }
    }
    v3 = append(v3, 129)
    var v28 int8 = 116
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
