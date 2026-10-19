package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 uint) uint8 {
    p0 += p0
    var v4 bool = (-603 <= 32767)
    v5 := []int{}
    p0 = p0
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    v6 := []int{127, -9223372036854775808, 1}
    var v7 int64 = 957
    if false {
        var v8 uint8 = 2
        _ = v8
    }
    if false {
        if false {
            v9 := GenS1{F0: 4294967295, F1: 42, F2: 1000}
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int = fi10
                fmt.Println(v7, fi10, fi10)
                _ = fi10
                _ = v11
            }
            if len(v6) > 0 {
                v6[0] = -2147483648
            }
            _ = v9
        }
        v7 += 1
    } else {
        var v12 uint32 = 1
        if true {
            var v13 uint8 = 2
            _ = v13
        }
        _ = v12
    }
    _ = v6
    _ = v7
    return 259
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 751
    var v3 uint32 = 4294967295
    var v14 uint32 = v3
    var v15 int8 = -1
    var v16 float64 = 1e10
    var v17 uint16 = 422
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := []int{}
        v19 = append(v19, (fi18 % ((fi18) | 1)))
        if (379 < -1) {
            var v20 uint32 = (v3 | 2147483647)
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    v15 = v15
    v21 := []int{280, 128, 0, genFn1()}
    var v22 int32 = 2147483647
    var v23 int16 = 42
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v24 float32 = 1.5
    v23 = 42
    var v25 int8 = (v15 / ((v15) | 1))
    switch genFn1() {
    case 127:
        switch "a" {
        case "foo bar":
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 float64 = v1
                if ((true || true) || (v27 <= v16)) {
                    switch "hello" {
                    case "\n\t":
                        for fi28 := 0; fi28 < 15; fi28++ {
                            v16 -= v16
                            _ = fi28
                        }
                    case "hello":
                        var v29 float64 = v27
                        _ = v29
                    default:
                        v30 := GenS1{F0: v2, F1: 2, F2: v23}
                        _ = v30
                    }
                } else {
                    v21 = append(v21, (^(fi26)))
                    v21 = append(v21, 32767)
                }
                fmt.Println(v23, v1, v27)
                _ = fi26
                _ = v27
            }
        default:
            if false {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                _ = len(v21)
                v31 := GenS1{F0: 879, F1: 1, F2: (v23 & -280)}
                _ = v31
            }
        }
    case 32768:
        for fi32 := 0; fi32 < 7; fi32++ {
            fmt.Println(v3, v3)
            _ = fi32
        }
    case 128:
        v21 = append(v21, 4294967295)
    default:
        _ = len(v21)
    }
    var v33 int64 = 9223372036854775807
    if true {
        fmt.Println(v2)
        var v34 bool = false
        v21 = append(v21, 0)
        _ = v34
    }
    fmt.Println(v14, v3, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
}
