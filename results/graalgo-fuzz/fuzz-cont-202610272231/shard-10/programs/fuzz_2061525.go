package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint) int64 {
    v6 := GenS1{F0: 2147483647, F1: false}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: 539, F1: (1.5 == 3.14159)}
        switch -837 {
        case 0:
            switch ("" + "\n\t") {
            case "a":
                v9 := []int{(fi7 ^ 127)}
                _ = v9
            case "\n\t":
                v10 := []int{(fi7 >> ((p0) & 63)), (fi7 ^ fi7), (^(fi7))}
                _ = v10
            default:
                for fi11 := 0; fi11 < 16; fi11++ {
                    if true {
                        p0 -= (p0 % (((p0 >> ((p0) & 63))) | 1))
                    } else {
                        p0 = 42
                        v12 := []int{(fi11 << ((p0) & 63))}
                        p0 = 255
                        _ = v12
                    }
                    _ = fi11
                }
            }
        case 42:
            var v13 uint32 = 4
            _ = v13
        case -402:
            var v14 uint64 = 1
            _ = v14
        }
        p0 -= p0
        _ = fi7
        _ = v8
    }
    var v15 bool = (true && false)
    _ = p0
    _ = v15
    _ = v6
    return -811
}

func main() {
    var v1 int64 = 42
    var v2 int64 = 9223372036854775807
    var v3 bool = (true || false)
    var v4 float64 = 3.14159
    var v5 uint16 = 65535
    v1 = genFn0(122)
    if v3 {
        for fi16 := 0; fi16 < 18; fi16++ {
            switch (fi16 | (fi16 - -1)) {
            case 634:
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 string = "\n\t"
                    _ = fi17
                    _ = v18
                }
            default:
                if (0.0 < -2.25) {
                    var v19 float32 = float32(3.14159)
                    _ = v19
                } else {
                    fmt.Println(fi16, v2, v3)
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int16 = 1
                        var v22 int16 = (-(v21))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            }
            _ = fi16
        }
        var v23 uint64 = 0
        var v24 uint16 = (^(v5))
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 554, F1: (0.0 <= 3.14159)}
    v26 := GenS1{F0: 0, F1: false}
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v3 = v3
            v5 = v5
            v5 -= v5
            _ = fi28
        }
        _ = fi27
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS2{F0: (^(v5)), F1: 4294967295}
        _ = fi29
        _ = v30
    }
    v31 := []int{135}
    var v32 int32 = -1
    var v33 int32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
