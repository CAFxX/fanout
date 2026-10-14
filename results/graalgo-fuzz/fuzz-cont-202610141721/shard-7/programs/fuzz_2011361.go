package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int64; F1 int8; F2 string }
type GenS5 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint }
func genFn0() int64 {
    var v7 uint16 = 42
    v8 := GenS1{F0: 42, F1: 9223372036854775807, F2: 325}
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{1, 0}
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    var v5 uint32 = 42
    v6 := make(map[string]int)
    fmt.Println(v5, v5)
    if (!false) {
        switch "" {
        case "":
            _ = len(v4)
        case "hello":
            var v9 bool = true
            _ = v9
        case "\n\t":
            v10 := GenS2{F0: genFn0(), F1: 80}
            _ = v10
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v11 bool = false
        var v12 uint16 = 255
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 868, F1: 42}
    if true {
        switch "" {
        case "a":
            switch -2147483648 {
            case 32768:
                var v14 int = 9223372036854775807
                _ = v14
            default:
                v5 += 928
            }
        case "\n\t":
            v15 := GenS3{F0: 371, F1: (v1 - 293), F2: v5}
            _ = v15
        default:
            if true {
                v16 := GenS1{F0: -8, F1: 42, F2: v1}
                var v17 int8 = -1
                _ = len(v2)
                _ = v16
                _ = v17
            }
        }
        v5 = v5
        if true {
            for fi18 := 0; fi18 < 1; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    if ("" > "") {
                        if (-46 <= 947) {
                            var v20 uint = 65535
                            v21 := GenS4{F0: 42, F1: -1, F2: ""}
                            _ = v20
                            _ = v21
                        } else {
                            var v22 uint8 = 184
                            _ = v22
                        }
                        v5 = (v5 | v5)
                        var v23 uint8 = 249
                        _ = v23
                    } else {
                        v1 = 2
                        var v24 int64 = genFn0()
                        var v25 uint8 = 124
                        v1 = (v1 ^ uint64(v25))
                        fmt.Println(fi19, v1, fi19)
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(fi18, v1, v1)
                    _ = fi19
                }
                var v26 int = fi18
                _ = fi18
                _ = v26
            }
            var v27 int = -1
            v28 := GenS2{F0: int64(v27), F1: 2}
            _ = v27
            _ = v28
        }
    }
    v1 -= (v1 + (^(v1)))
    if false {
        v29 := GenS2{F0: 42, F1: 1}
        fmt.Println(v5, v5)
        _ = v29
    } else {
        v1 += v1
        v5 -= v5
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 float32 = 3.14159
        v32 := GenS5{F0: 127, F1: genFn0(), F2: v5}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v1, v5)
    v1 += 996
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS6{F0: 65535}
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
