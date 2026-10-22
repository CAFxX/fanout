package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16; F2 int16 }
func genFn0(p0 int8) uint64 {
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                switch "hello" {
                case "\n\t":
                    v9 := GenS1{F0: 444, F1: 0, F2: 425}
                    _ = v9
                case "foo bar":
                    var v10 uint16 = 32767
                    _ = v10
                case "a":
                    var v11 uint64 = 0
                    _ = v11
                }
                v12 := GenS2{F0: 4294967295, F1: true}
                _ = fi8
                _ = v12
            }
            var v13 int64 = -542
            _ = fi7
            _ = v13
        }
    case "foo bar":
        p0 = (p0 * p0)
    default:
        var v14 uint64 = 4294967295
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            if (65535 > 645) {
                var v16 int = fi15
                _ = v16
            }
            _ = fi15
        }
    }
    var v17 uint8 = 127
    fmt.Println(v17)
    var v18 uint64 = 2
    _ = p0
    _ = v17
    _ = v18
    return 18446744073709551615
}
func genFn1() bool {
    if ("\n\t" <= "") {
        var v19 uint = 703
        _ = v19
    } else {
        var v20 int8 = -1
        v21 := []int{0, -9223372036854775808, 1, 931}
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 313, F1: false}
    for fi23 := 0; fi23 < 15; fi23++ {
        switch (^(fi23)) {
        case -208:
            var v24 bool = (1000 <= 42)
            _ = v24
        }
        _ = fi23
    }
    var v25 uint8 = 2
    v26 := GenS3{F0: 0}
    _ = v22
    _ = v25
    _ = v26
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{2147483647}
    v3 := make(map[int]bool)
    var v4 uint16 = 42
    var v5 int = 32767
    v6 := []int{}
    v4 = v4
    fmt.Println(v5, v4)
    var v27 uint16 = (v4 / ((v4) | 1))
    fmt.Println(v4)
    _ = len(v1)
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS1{F0: 42, F1: 42, F2: 1}
        v30 := GenS4{F0: "\n\t", F1: 1, F2: 1}
        v27 -= v27
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = len(v6)
    var v31 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
