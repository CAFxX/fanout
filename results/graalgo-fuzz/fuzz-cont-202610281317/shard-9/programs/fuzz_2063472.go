package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 bool = true
        fmt.Println(v8, v8)
        _ = fi7
        _ = v8
    }
    var v9 string = "\n\t"
    _ = v9
    return 2147483647
}

func main() {
    var v1 int = -913
    var v2 uint16 = 2
    var v3 bool = false
    var v4 int64 = 1
    v5 := make(map[string]bool)
    var v6 int64 = v4
    var v10 bool = ((-2.25 + -2.25) > (1.5 + 1.5))
    if ((^(v6)) != (v4 * (v4 / ((-1) | 1)))) {
        var v11 int16 = 0
        _ = v11
    }
    if ((v4 << ((930) & 63)) == (v4 - (v4 >> ((708) & 63)))) {
        var v12 bool = (v3 || (0 > v4))
        _ = v12
    } else {
        var v13 uint16 = (v2 * v2)
        v14 := []int{(v1 * (-(v1))), (v1 >> ((2) & 63)), v1}
        if ((-(v4)) >= -9223372036854775808) {
            _ = len(v5)
            if true {
                if true {
                    var v15 int64 = v6
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    var v16 float64 = float64(1.5)
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS1{F0: -1, F1: (v1 >> ((0) & 63))}
                    _ = len(v5)
                    _ = v17
                }
                var v18 int8 = 13
                _ = v18
            }
            v1 = 32768
        }
        _ = v13
        _ = v14
    }
    v19 := GenS1{F0: -2147483648, F1: 32767}
    fmt.Println(v6)
    v20 := GenS2{F0: v1, F1: -28, F2: 0}
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int64 = (v4 - v4)
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v1, v1)
    v3 = true
    v5[("" + ("\n\t" + "a"))] = v3
    if false {
        v2 += 193
        v1 = -9223372036854775808
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            v6 = v4
            for fi24 := 0; fi24 < 7; fi24++ {
                fmt.Println(v3, fi23, v2)
                v25 := GenS1{F0: 0, F1: fi23}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        if v10 {
            for fi26 := 0; fi26 < 14; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    if v3 {
                        fmt.Println(v3, v1)
                    }
                    if true {
                        v28 := []int{(v1 + fi27), (fi26 % (((-(fi26))) | 1)), 731}
                        var v29 uint8 = 0
                        _ = v28
                        _ = v29
                    }
                    _ = fi27
                }
                var v30 string = ("foo bar" + "")
                _ = fi26
                _ = v30
            }
            v31 := []int{(v1 + v1), (-(v1)), (v1 ^ (v1 * 2147483647))}
            for fi32 := 0; fi32 < 17; fi32++ {
                v5[("a" + "hello")] = v3
                _ = fi32
            }
            _ = v31
        }
        var v33 uint8 = 0
        _ = v33
    }
    v2 -= (v2 * (^(v2)))
    var v34 bool = (!(0 > 1))
    v4 += 679
    var v35 uint8 = 0
    _ = len(v5)
    var v36 uint32 = 2147483647
    fmt.Println(v10, v34, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
