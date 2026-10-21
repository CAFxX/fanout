package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 int32) uint {
    v4 := GenS1{F0: 289, F1: 2147483647}
    v5 := GenS1{F0: 502, F1: 65535}
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    v3 := make(map[int]string)
    v6 := GenS1{F0: 985, F1: 67}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    var v9 int8 = 18
    v2 = v2
    var v10 uint8 = v2
    var v11 float64 = 0.1
    for fi12 := 0; fi12 < 2; fi12++ {
        v1[861] = fi12
        switch (fi12 / ((fi12) | 1)) {
        case 484:
            var v13 int16 = 0
            _ = v13
        }
        _ = fi12
    }
    fmt.Println(v9, v11, v9)
    if ((float64((((float32((((float64(((1e10 * 0.1) * (3.14159 / 1.0))) * 3.14159) / v11) + 1e10)) - ((1e10 * 1e10) + 0.0)) + -2.25) - (1.5 * -2.25))) - 1e10) <= 0.1) {
        for fi14 := 0; fi14 < 11; fi14++ {
            if true {
                if true {
                    if (-1 > -1) {
                        v15 := GenS1{F0: (fi14 - fi14), F1: 2147483647}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        switch fi14 {
                        case -641:
                            var v17 uint = 255
                            _ = v17
                        case 32767:
                            v10 = v2
                        case -352:
                            delete(v1, 2147483647)
                        }
                        v11 = -2.25
                        v10 += (^(v10))
                        _ = fi16
                    }
                }
            }
            fmt.Println(fi14)
            _ = fi14
        }
        delete(v3, -9223372036854775808)
    } else {
        v18 := GenS1{F0: 194, F1: 2147483648}
        var v19 uint = 368
        _ = v18
        _ = v19
    }
    v2 -= 0
    var v20 int16 = -1
    v9 -= 0
    var v21 int64 = -28
    if true {
        var v22 string = "a"
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 uint8 = v2
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 int = -2147483648
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        var v27 string = v22
        _ = v22
        _ = v27
    } else {
        v2 -= v10
        v2 -= v2
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := []int{(fi28 << ((255) & 63))}
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v6
    _ = v9
}
