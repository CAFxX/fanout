package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 int) int64 {
    p0 += p0
    fmt.Println(p0)
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    if true {
        var v6 int8 = -116
        var v7 int64 = -1
        v8 := GenS1{F0: (v7 <= (v7 >> ((2) & 63))), F1: 461}
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 64
    v10 := GenS2{F0: 9223372036854775807, F1: -2147483648}
    var v11 float64 = -2.25
    if (v11 == (1e10 * (v11 + v11))) {
        if true {
            v5 = (v5 | v5)
            v9 = (^(v9))
            for fi12 := 0; fi12 < 15; fi12++ {
                v5 -= v5
                switch (((("hello" + "a") + "") + "") + "a") {
                case "a":
                    var v13 int64 = 42
                    _ = v13
                case "foo bar":
                    v9 = (v9 >> ((2) & 63))
                case "":
                    if (!false) {
                        switch fi12 {
                        case 32767:
                            var v14 int8 = -118
                            _ = v14
                        }
                        v5 += (v5 + (v5 + (v5 | v5)))
                        switch fi12 {
                        case 266:
                            if false {
                                var v15 uint8 = 1
                                var v16 float64 = -2.25
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint32 = 839
                                _ = v17
                            }
                        default:
                            var v18 bool = true
                            _ = v18
                        }
                    } else {
                        v11 -= 0.0
                        v19 := []int{}
                        _ = v19
                    }
                }
                _ = fi12
            }
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if true {
            switch "" {
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi20)
                }
            default:
                fmt.Println(v5)
            }
        } else {
            var v21 float64 = float64((0.0 - 0.1))
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint32 = 77
                v24 := []int{}
                if len(v24) > 0 {
                    v24[0] = 1
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v21
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            switch int(v9) {
            case 9223372036854775807:
                fmt.Println(fi25, v11)
            case 42:
                var v26 uint16 = v9
                _ = v26
            }
            fmt.Println(v9, v9)
            v27 := []int{(fi20 | 0), (-(fi20)), int(v5), (fi20 * fi20)}
            _ = fi25
            _ = v27
        }
        var v28 float32 = (-2.25 + 3.14159)
        _ = fi20
        _ = v28
    }
    var v29 bool = (1 >= 42)
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint64 = 0
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
