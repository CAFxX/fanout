package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 bool) int {
    if p0 {
        for fi6 := 0; fi6 < 15; fi6++ {
            if true {
                p0 = p0
                var v7 float64 = 1e10
                var v8 int16 = 1
                _ = v7
                _ = v8
            } else {
                var v9 uint8 = 128
                for fi10 := 0; fi10 < 1; fi10++ {
                    v11 := []int{(fi6 >> ((65535) & 63)), (fi6 - fi10), fi6}
                    _ = fi10
                    _ = v11
                }
                v12 := []int{42, 0}
                _ = v12
                _ = v9
            }
            if (p0 && false) {
                var v13 int32 = 261
                var v14 int = fi6
                switch fi6 {
                case 518:
                    var v15 float32 = 1e10
                    _ = v15
                case -9223372036854775808:
                    v16 := GenS1{F0: (^(v13))}
                    _ = v16
                case 9223372036854775807:
                    if p0 {
                        v14 = fi6
                        fmt.Println(v14, v13, v13)
                    }
                default:
                    v14 = -1
                }
                _ = v13
                _ = v14
            }
            _ = fi6
        }
    }
    v17 := GenS1{F0: 42}
    v18 := []int{}
    _ = p0
    _ = v17
    _ = v18
    return -146
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 181
    var v3 float64 = float64(((0.0 * 0.0) - 1e10))
    v4 := make(map[int]uint32)
    var v5 float64 = 0.0
    v19 := []int{genFn0((42 < v2)), 32768}
    fmt.Println(v2)
    if false {
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(fi20, v3, v5)
                var v22 float64 = float64(0.1)
                _ = fi21
                _ = v22
            }
            delete(v4, (fi20 & fi20))
            _ = fi20
        }
        fmt.Println(v1, v1, v2)
        switch -208 {
        case -2147483648:
            var v23 int16 = 42
            _ = v23
        case 32768:
            if (false && (1e10 >= ((float32(float64(1e10)) * 1e10) / 3.14159))) {
                fmt.Println(v2, v2, v5)
                for fi24 := 0; fi24 < 16; fi24++ {
                    switch (fi24 / (((fi24 - -2147483648)) | 1)) {
                    case 42:
                        fmt.Println(fi24, v1)
                    }
                    _ = len(v4)
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v26 := GenS1{F0: v1}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            } else {
                switch "hello" {
                case "a":
                    v5 -= v5
                }
                v4[32768] = 2
                v27 := GenS1{F0: (^(v1))}
                _ = v27
            }
        case 42:
            var v28 int16 = 1000
            _ = v28
        }
    }
    v2 = 2
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS1{F0: int32(fi29)}
        v31 := v4[fi29]
        _ = fi29
        _ = v30
        _ = v31
    }
    v3 -= 1.5
    _ = len(v19)
    fmt.Println(v3, v1, v1)
    v2 = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
