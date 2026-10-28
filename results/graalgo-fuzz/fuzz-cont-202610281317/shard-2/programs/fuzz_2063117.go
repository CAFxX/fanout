package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0() uint {
    var v6 int16 = -445
    v6 = 0
    v6 = (v6 | (v6 / ((v6) | 1)))
    v6 += (v6 & (v6 ^ 32767))
    v7 := []int{342, 428}
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = 1e10
    var v2 string = "a"
    var v3 float32 = (((v1 * float32((3.14159 + (1e10 - (0.0 / 1e10))))) / 1.5) * v1)
    var v4 bool = true
    var v5 uint16 = 1
    var v8 int = 9223372036854775807
    v9 := []int{(v8 << ((genFn0()) & 63)), 32768, v8}
    for fi10 := 0; fi10 < 1; fi10++ {
        if false {
            fmt.Println(v8, v2, fi10)
            v1 += (float32(1.5) - v3)
        } else {
            if v4 {
                if v4 {
                    fmt.Println(v3, v3, v8)
                }
            }
            if false {
                v5 -= 42
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 uint16 = 0
                    if ((v2 <= v2) && false) {
                        var v13 float64 = 1.5
                        fmt.Println(v8)
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 910, F1: (v12 & v5), F2: v4}
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            } else {
                var v15 int16 = 0
                _ = v15
            }
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        _ = len(v9)
        switch (fi16 ^ (v8 - (v8 | (v8 / ((fi16) | 1))))) {
        case 32767:
            fmt.Println(v4)
        }
        _ = fi16
    }
    v8 = (v8 ^ v8)
    fmt.Println(v1)
    var v17 int = (v8 & (v8 / ((v8) | 1)))
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v18 int8 = -114
    var v19 uint16 = 2
    var v20 int64 = -1
    for fi21 := 0; fi21 < 18; fi21++ {
        if (v20 >= v20) {
            var v22 string = "foo bar"
            v9 = append(v9, fi21)
            _ = v22
        }
        fmt.Println(v5, v5, v17)
        _ = fi21
    }
    fmt.Println(v5, v8, v17)
    var v23 int64 = v20
    switch -294 {
    case -2147483648:
        var v24 uint = 65535
        _ = v24
    }
    var v25 int = v8
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
