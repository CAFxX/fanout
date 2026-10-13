package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 uint64) uint64 {
    var v5 uint = 1
    var v6 int = 595
    fmt.Println(v5)
    if (2147483647 <= 2147483647) {
        if true {
            var v7 uint64 = p0
            if false {
                var v8 uint = v5
                var v9 float64 = -2.25
                _ = v8
                _ = v9
            } else {
                var v10 int8 = -128
                v5 += 65535
                _ = v10
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 uint64 = (p0 / ((18446744073709551615) | 1))
                if true {
                    fmt.Println(v5)
                    var v13 float32 = 0.1
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            v14 := GenS1{F0: 42, F1: 0, F2: 74}
            if true {
                v15 := []int{(v6 - (v6 * (-(v6))))}
                v5 = (v5 & (v5 * v5))
                if true {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 uint32 = 227
                        v18 := GenS1{F0: v17, F1: 32767, F2: 12}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    v6 -= v6
                }
                _ = v15
            } else {
                var v19 uint = (v5 / ((670) | 1))
                _ = v19
            }
            _ = v14
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := []uint32{2147483648}
    var v2 int32 = 37
    var v3 uint = 4294967295
    var v4 float64 = (1.5 / -2.25)
    v3 -= (v3 % ((0) | 1))
    var v20 float32 = 1.5
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if false {
                var v23 uint16 = 2
                fmt.Println(v23, v4, v4)
                _ = v23
            } else {
                v24 := GenS1{F0: 2, F1: 65535, F2: 0}
                _ = len(v1)
                _ = v24
            }
            _ = fi22
        }
        _ = fi21
    }
    v1 = append(v1, 240)
    var v25 float64 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += v2
    switch "hello" {
    case "\n\t":
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 20; fi27++ {
                if true {
                    var v28 uint64 = 4294967295
                    _ = v28
                }
                v29 := GenS1{F0: 694, F1: 718, F2: 42}
                var v30 int8 = 0
                var v31 uint64 = 182
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                _ = fi27
                _ = v29
                _ = v30
                _ = v31
            }
            fmt.Println(v25, v2, v4)
            var v32 uint32 = 4294967295
            _ = fi26
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
