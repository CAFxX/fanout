package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int32 }
func genFn0(p0 uint16) int {
    for fi6 := 0; fi6 < 15; fi6++ {
        if true {
            v7 := GenS1{F0: 1, F1: 52}
            var v8 int64 = 42
            switch (("\n\t" + ("a" + "hello")) + ("\n\t" + "foo bar")) {
            case "\n\t":
                fmt.Println(v8, fi6, v8)
            default:
                if (1000 != -413) {
                    fmt.Println(v8, fi6)
                }
            }
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 42, F1: 87}
    p0 += 774
    _ = p0
    _ = v9
    return 793
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 string = "hello"
    v5 := []int{-107}
    var v10 uint32 = 0
    var v11 string = "a"
    var v12 int32 = -1
    v13 := GenS2{F0: 100}
    var v14 int8 = -128
    var v15 int16 = 1
    v12 -= (v12 >> ((v1) & 63))
    var v16 uint = v1
    switch 128 {
    case -861:
        if ((1.5 + 0.1) == 0.1) {
            v15 = 1
            var v17 uint16 = 1
            switch 1 {
            case 42:
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 bool = true
                    _ = fi18
                    _ = v19
                }
            case 127:
                if ((float64((1e10 / float32(1e10))) / 1.0) != 0.0) {
                    if true {
                        v12 -= (^(v12))
                    } else {
                        v16 += 4294967295
                    }
                } else {
                    v20 := GenS2{F0: v14}
                    v15 = v15
                    _ = v20
                }
            case 128:
                v10 += 2147483647
            }
            _ = v17
        } else {
            v1 += 1
            fmt.Println(v14, v14)
            if false {
                var v21 bool = false
                _ = v21
            } else {
                var v22 string = "\n\t"
                _ = v22
            }
        }
    default:
        if false {
            if true {
                v1 += v1
                v11 = v11
                _ = len(v5)
            } else {
                var v23 int32 = (v12 & (v12 * v12))
                var v24 bool = false
                var v25 float64 = 1e10
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            var v26 string = ""
            var v27 int64 = 1
            var v28 int8 = (-(v14))
            _ = v26
            _ = v27
            _ = v28
        }
    }
    v29 := GenS3{F0: 42, F1: -341, F2: (v12 | v12)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
