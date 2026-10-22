package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32, p1 string) bool {
    p1 = "hello"
    v5 := GenS1{F0: "\n\t", F1: (3.14159 == (1.5 * (3.14159 / ((-2.25 - 3.14159) + float32(1.5))))), F2: 2}
    _ = p0
    _ = p1
    _ = v5
    return true
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    if (0 <= 1) {
        var v6 string = "foo bar"
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint32 = (^(v1))
            var v9 uint64 = uint64(v8)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    v10 := []int{-478, 2147483647, 32767}
    var v11 int16 = 0
    var v12 int16 = 32767
    var v13 uint64 = 1
    switch 9223372036854775807 {
    case -2147483648:
        var v14 int8 = 42
        _ = v14
    case 2147483647:
        if true {
            var v15 int64 = -740
            var v16 int8 = 1
            _ = v15
            _ = v16
        } else {
            fmt.Println(v2, v13, v13)
            var v17 string = "hello"
            if false {
                var v18 int64 = 0
                v19 := GenS2{F0: v12}
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 int8 = 100
                    v22 := GenS3{F0: 0}
                    _ = len(v10)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v18
                _ = v19
            } else {
                for fi23 := 0; fi23 < 5; fi23++ {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        delete(v4, v17)
                        var v25 uint16 = 65535
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if false {
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        }
                        v28 := GenS2{F0: v12}
                        _ = fi26
                        _ = v28
                    }
                    delete(v3, 545)
                    _ = fi23
                }
                v17 = "a"
                var v29 int32 = 2147483647
                _ = v29
            }
            _ = v17
        }
    case 240:
        var v30 uint64 = v13
        _ = v30
    default:
        var v31 int = 802
        _ = v31
    }
    fmt.Println(v2, v1)
    v32 := GenS3{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
