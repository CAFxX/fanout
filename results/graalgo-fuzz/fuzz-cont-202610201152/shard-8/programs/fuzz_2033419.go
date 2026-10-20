package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int16, p1 uint64, p2 uint) bool {
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 1; fi8++ {
        if (!(42 >= 99)) {
            var v9 uint16 = 620
            if (883 == 2147483648) {
                p0 = p0
                if true {
                    fmt.Println(v9)
                    switch fi8 {
                    case -1:
                        for fi10 := 0; fi10 < 4; fi10++ {
                            v9 -= (^(v9))
                            v11 := []int{fi10}
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = fi10
                            _ = v11
                        }
                    default:
                        v12 := []int{fi8}
                        _ = v12
                    }
                }
            }
            _ = v9
        } else {
            var v13 string = v7
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 bool = false
            _ = fi14
            _ = v15
        }
        if ((42 == 76) && true) {
            p0 -= 0
            v16 := GenS1{F0: 1, F1: (p2 % ((p2) | 1))}
            v17 := GenS1{F0: int8(p1), F1: p2}
            _ = v16
            _ = v17
        }
        _ = fi8
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if false {
            var v19 uint64 = p1
            _ = v19
        } else {
            if false {
                fmt.Println(v7, v7, p1)
            }
        }
        var v20 float64 = ((3.14159 * -2.25) - 1e10)
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int8 = 42
            var v23 uint64 = 844
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi18
        _ = v20
    }
    var v24 uint32 = 4294967295
    v25 := GenS1{F0: 100, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v7
    return (1 < 65535)
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := []int{-2147483648, 42, -809, 32767}
    var v5 uint32 = 2147483648
    v6 := []int{-777}
    v4 = append(v4, -204)
    var v26 float64 = -2.25
    var v27 float32 = 1.5
    fmt.Println(v1, v1, v26)
    v28 := GenS1{F0: 1, F1: 0}
    fmt.Println(v5)
    if true {
        v1 -= (^(v1))
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 uint = 1
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v26)
    v31 := GenS1{F0: -63, F1: 545}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
