package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint = 255
            v7 = (v7 + v7)
            v7 -= (v7 / (((^(v7))) | 1))
            var v8 int32 = 1
            var v9 int32 = -859
            if ((v8 | v9) == (v9 % ((v9) | 1))) {
                if (18 == 107) {
                    v10 := GenS1{F0: -1, F1: true, F2: (v7 / ((4294967295) | 1))}
                    var v11 uint16 = 42
                    v12 := []int{}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: (v9 & -2147483648), F1: ((42 > 730) || false), F2: v7}
                v9 = (v9 * v9)
                _ = v13
            } else {
                if false {
                    fmt.Println(fi5)
                    if true {
                        v9 += v9
                        var v14 uint32 = 42
                        _ = v14
                    }
                } else {
                    v15 := []int{2147483647, 4294967295, (fi5 / ((fi6) | 1))}
                    _ = v15
                }
            }
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v16 float32 = 3.14159
        if false {
            var v17 uint64 = 4294967295
            var v18 uint64 = 0
            v17 = (v18 % ((2) | 1))
            _ = v17
            _ = v18
        }
        _ = fi5
        _ = v16
    }
    var v19 bool = ((503 <= 0) && false)
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.1
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    if true {
        v2 += v2
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 uint16 = 1
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v4, v4)
    var v22 bool = (false && ("hello" < "\n\t"))
    var v23 int16 = 1000
    v22 = (!(!(!v22)))
    v2 += 3.14159
    v3[65535] = false
    v24 := GenS1{F0: 0, F1: (!v22), F2: 65535}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v23)
    }
    v25 := []int{-333, 1}
    for fi26 := 0; fi26 < 5; fi26++ {
        switch "" {
        case "a":
            v22 = false
        default:
            if (("a" < (("hello" + "a") + "foo bar")) || false) {
                v4 += v4
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 int8 = -8
                    v29 := GenS1{F0: 568, F1: (42 >= v4), F2: 42}
                    v28 = (^(v28))
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                if v22 {
                    v25 = append(v25, 128)
                    for fi30 := 0; fi30 < 12; fi30++ {
                        var v31 bool = v22
                        v23 += (v23 * 1)
                        _ = fi30
                        _ = v31
                    }
                } else {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v23 -= v23
                        _ = fi32
                    }
                    switch "foo bar" {
                    case "a":
                        var v33 uint64 = 0
                        _ = v33
                    case "\n\t":
                        var v34 uint16 = 255
                        _ = v34
                    default:
                        v23 = -740
                    }
                    v25 = append(v25, -882)
                }
            }
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
