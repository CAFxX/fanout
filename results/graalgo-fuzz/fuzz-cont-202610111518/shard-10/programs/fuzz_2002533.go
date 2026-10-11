package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0() uint {
    var v4 uint8 = 0
    v4 -= (v4 - (v4 * v4))
    v5 := []int{1, 2147483647, 127}
    if (v4 != 1) {
        if (0 <= -1) {
            var v6 uint = 2
            _ = v6
        } else {
            fmt.Println(v4)
            _ = len(v5)
        }
        fmt.Println(v4)
        var v7 float64 = 1e10
        _ = v7
    } else {
        if len(v5) > 0 {
            v5[0] = 207
        }
        var v8 uint16 = 160
        switch 0 {
        case -1:
            v4 -= v4
        case -9223372036854775808:
            var v9 uint32 = 65535
            _ = v9
        case -2147483648:
            fmt.Println(v4)
        }
        _ = v8
    }
    var v10 string = "hello"
    _ = v10
    _ = v4
    _ = v5
    return 738
}
func genFn1(p0 int32, p1 uint16, p2 int32) int {
    p2 -= p2
    v11 := GenS1{F0: -643, F1: 83}
    v12 := GenS1{F0: 0, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 int32 = -803
    v3 = 2147483647
    v3 = (v3 >> ((18446744073709551615) & 63))
    var v13 int16 = 407
    switch 127 {
    case -9223372036854775808:
        fmt.Println(v13, v13)
    default:
        if (("a" > "foo bar") || (!(!true))) {
            v3 += v3
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 float32 = 0.1
                v3 = 1
                _ = fi14
                _ = v15
            }
            _ = len(v1)
        }
    }
    var v16 int = -497
    var v17 uint16 = 2
    if ((v3 << ((218) & 63)) > v3) {
        v18 := GenS1{F0: v13, F1: 0}
        v3 = v3
        v19 := []int{v16, (v16 ^ 42), (v16 << ((255) & 63))}
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: v13, F1: 42}
        _ = v20
    }
    v17 += v17
    var v21 int16 = (-(v13))
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            switch (v16 ^ (v16 ^ fi23)) {
            case 4294967295:
                var v24 bool = false
                _ = v24
            case 128:
                if false {
                    v25 := []int{(v16 + genFn1((-(v3)), (v17 | v17), 0)), 128, v16, 0}
                    switch "\n\t" {
                    case "\n\t":
                        var v26 int8 = 42
                        _ = v26
                    case "a":
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            var v27 int8 = -128
                            _ = v27
                        case "a":
                            v13 -= v21
                        case "":
                            v21 = v13
                        default:
                            delete(v1, 42)
                        }
                    case "":
                        v3 = 2147483647
                    }
                    var v28 uint8 = 1
                    _ = v25
                    _ = v28
                }
            }
            v29 := GenS1{F0: v13, F1: 255}
            var v30 string = "a"
            _ = fi23
            _ = v29
            _ = v30
        }
        var v31 uint8 = 131
        v3 = 1
        _ = fi22
        _ = v31
    }
    var v32 uint8 = 127
    v21 += 42
    v3 -= (^(v3))
    fmt.Println(v21, v32)
    var v33 uint64 = 395
    v34 := GenS1{F0: 137, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
