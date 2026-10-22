package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint32; F2 int16 }
func genFn0() int64 {
    if (true && (false && true)) {
        var v4 float64 = 1e10
        _ = v4
    }
    if true {
        if true {
            for fi5 := 0; fi5 < 17; fi5++ {
                var v6 uint = 581
                for fi7 := 0; fi7 < 7; fi7++ {
                    var v8 uint = 791
                    _ = fi7
                    _ = v8
                }
                _ = fi5
                _ = v6
            }
        }
        var v9 int8 = -128
        _ = v9
    } else {
        if (true && (!false)) {
            var v10 int = -2147483648
            var v11 uint64 = 537
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: 1, F2: 1}
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        } else {
            var v14 uint16 = 152
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v14 += v14
                    v17 := []int{-989, (^(fi15)), fi15}
                    var v18 float32 = (-2.25 + 3.14159)
                    var v19 int = fi16
                    if false {
                        if len(v17) > 0 {
                            v17[0] = (fi16 & (v19 / ((fi15) | 1)))
                        }
                        v20 := GenS1{F0: 1, F1: 1, F2: -2147483648}
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi15
            }
            _ = v14
        }
        v21 := []int{302, -730, 32767}
        _ = v21
    }
    v22 := GenS2{F0: 243}
    _ = v22
    return 42
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v23 := []int{0, 1, 127}
    for fi24 := 0; fi24 < 10; fi24++ {
        _ = len(v23)
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi24
    }
    var v25 uint = 194
    fmt.Println(v1, v25)
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 bool = (v1 > (v1 | v1))
        _ = fi26
        _ = v27
    }
    var v28 uint16 = 32767
    for fi29 := 0; fi29 < 14; fi29++ {
        if (42 < 1) {
            v25 = v25
            v30 := GenS3{F0: -813, F1: 353, F2: 1000}
            v28 -= (^(v28))
            _ = v30
        }
        delete(v2, "hello")
        _ = fi29
    }
    delete(v2, "\n\t")
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v28
    _ = v3
}
