package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint16) int32 {
    var v6 uint32 = 42
    v7 := GenS1{F0: p1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 926
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 int64 = -325
    var v4 float64 = float64(((1e10 / 1e10) - -2.25))
    var v5 uint32 = 533
    var v8 int = 127
    if (!(!false)) {
        v9 := []int{(-(v8)), (v8 & v8), v8}
        _ = v9
    }
    switch (("a" + "\n\t") + "hello") {
    case "\n\t":
        if (v4 > -2.25) {
            var v10 uint32 = (^(v5))
            _ = v10
        } else {
            v11 := []int{v8, 545}
            v11 = append(v11, (v8 ^ v8))
            if (!false) {
                v1[(v5 << ((65535) & 63))] = true
            } else {
                var v12 bool = false
                _ = v12
            }
            _ = v11
        }
    case "foo bar":
        var v13 uint16 = 2
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v5 += v5
        _ = fi14
    }
    var v15 uint64 = 987
    var v16 int8 = 125
    v3 += (v3 << ((461) & 63))
    if (int16(v15) < 1) {
        fmt.Println(v16, v8, v4)
    }
    fmt.Println(v5, v5)
    var v17 int64 = (v3 | v3)
    if ((float64(1.5) - 0.1) < ((v4 + 3.14159) / 1.0)) {
        v5 -= (v5 >> ((671) & 63))
    }
    var v18 int = (v8 & (-(v8)))
    var v19 int64 = -182
    var v20 int32 = 42
    fmt.Println(v17, v5)
    switch 18 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 float64 = float64(3.14159)
            switch "hello" {
            case "hello":
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint32 = v5
                    switch int(v24) {
                    case 32767:
                        v25 := []int{}
                        _ = v25
                    default:
                        switch "foo bar" {
                        case "a":
                            if true {
                                if (493 <= 255) {
                                    v8 = v18
                                    var v26 int = 4294967295
                                    _ = len(v2)
                                    var v27 int16 = 42
                                    v18 += int(v27)
                                    _ = v26
                                    _ = v27
                                }
                                v28 := []int{-9223372036854775808, (fi23 << ((4294967295) & 63)), 0, (fi23 ^ v18)}
                                v3 = -765
                                _ = v28
                            }
                        case "hello":
                            v18 += int(v19)
                        default:
                            var v29 uint64 = 417
                            _ = v29
                        }
                    }
                    switch "foo bar" {
                    case "hello":
                        var v30 float32 = (-2.25 + 3.14159)
                        _ = v30
                    case "foo bar":
                        var v31 uint16 = 65535
                        _ = v31
                    default:
                        v32 := []int{-439, (^(v8)), (v8 << ((790) & 63)), 257}
                        _ = v32
                    }
                    _ = fi23
                    _ = v24
                }
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
