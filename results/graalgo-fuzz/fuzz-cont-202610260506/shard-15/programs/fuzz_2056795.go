package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0() int16 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v6 uint32 = 700
        _ = v6
    case 9223372036854775807:
        var v7 uint8 = 249
        _ = v7
    }
    var v8 int32 = 1
    for fi9 := 0; fi9 < 19; fi9++ {
        if (((float32(float64(1.5)) >= 1.5) || (v8 >= (v8 << ((42) & 63)))) || (-437 >= (^(v8)))) {
            switch "\n\t" {
            case "\n\t":
                var v10 uint32 = 644
                _ = v10
            case "a":
                if (4294967295 > 0) {
                    v8 = v8
                    fmt.Println(fi9, fi9)
                } else {
                    if true {
                        var v11 uint64 = 42
                        _ = v11
                    } else {
                        var v12 int32 = (v8 % ((v8) | 1))
                        v8 -= v12
                        _ = v12
                    }
                }
            case "foo bar":
                var v13 string = ("\n\t" + "a")
                _ = v13
            default:
                fmt.Println(v8)
            }
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                if false {
                    switch (fi9 - (fi14 % ((fi9) | 1))) {
                    case -9223372036854775808:
                        v15 := GenS1{F0: 4294967295, F1: 0}
                        _ = v15
                    }
                    var v16 int = fi9
                    _ = v16
                } else {
                    v8 = v8
                }
                _ = fi14
            }
        }
        _ = fi9
    }
    if (true && false) {
        v8 += v8
    }
    _ = v8
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 78
    var v3 uint8 = 43
    v4 := make(map[int]int)
    var v5 bool = false
    var v17 uint16 = 1
    if ((v2 / (((^(v2))) | 1)) >= 1) {
        for fi18 := 0; fi18 < 6; fi18++ {
            switch ("\n\t" + "hello") {
            case "":
                var v19 int = (fi18 << ((1) & 63))
                _ = v19
            case "hello":
                v20 := []int{fi18, (fi18 * 127), (fi18 - 32767)}
                _ = v20
            case "\n\t":
                var v21 uint64 = 4294967295
                _ = v21
            }
            v2 += (v2 * v2)
            _ = fi18
        }
    }
    fmt.Println(v1, v17, v2)
    v22 := []int{-9223372036854775808, 2147483647}
    v2 = (v2 / ((100) | 1))
    var v23 int64 = int64(v3)
    var v24 int = 0
    var v25 float32 = ((0.0 / 0.1) * (3.14159 / 3.14159))
    fmt.Println(v1, v24)
    v5 = v5
    v26 := v4[v24]
    if (true || false) {
        var v27 int = v26
        _ = v27
    }
    v2 += v2
    fmt.Println(v3, v25, v2)
    var v28 uint16 = (v17 + (v17 / (((v17 * v17)) | 1)))
    v26 -= (v24 ^ v24)
    v4[(v24 * 545)] = 536
    var v29 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
