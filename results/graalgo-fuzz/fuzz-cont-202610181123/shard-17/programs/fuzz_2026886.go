package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint8, p1 int64, p2 int64) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            p1 -= p2
            var v9 uint8 = p0
            v10 := GenS1{F0: ("" + "\n\t"), F1: 1, F2: 42}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: ("\n\t" + (("" + "hello") + "a")), F1: 1, F2: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 738
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 80
    var v3 int8 = 73
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    v3 = (v3 ^ v3)
    v1["a"] = -2147483648
    if ((^(v3)) < (v3 << ((2) & 63))) {
        v12 := GenS1{F0: ((("a" + "a") + "hello") + "hello"), F1: 852, F2: 42}
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS2{F0: 0, F1: v3}
                var v16 float32 = 0.1
                if false {
                    v17 := v4[2147483647]
                    if (true || ((2 == 123) && true)) {
                        var v18 uint16 = 1
                        _ = v18
                    } else {
                        if (v2 == (v2 * (v2 | v2))) {
                            fmt.Println(v3)
                        }
                    }
                    if (1 == 0) {
                        delete(v5, ("\n\t" + ("\n\t" + "")))
                    }
                    _ = v17
                }
                v2 = v2
                switch (fi14 & 32768) {
                case -464:
                    v16 = (float32(1e10) + (v16 / v16))
                default:
                    v19 := []int{fi13}
                    _ = v19
                }
                var v20 int16 = 298
                v16 = v16
                _ = fi14
                _ = v15
                _ = v16
                _ = v20
            }
            _ = fi13
        }
        _ = v12
    } else {
        v3 += v3
        v21 := GenS1{F0: "", F1: (v2 >> ((v2) & 63)), F2: -164}
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: 941, F1: (v3 << ((255) & 63))}
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 392
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int16 = 805
        _ = fi25
        _ = v26
    }
    if false {
        var v27 uint32 = 1
        _ = v27
    }
    var v28 uint8 = 40
    var v29 int = 134
    v30 := GenS2{F0: 1, F1: v3}
    v31 := GenS1{F0: "\n\t", F1: (v2 ^ (v2 << (((v2 ^ 42)) & 63))), F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
