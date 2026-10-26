package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64; F1 int; F2 uint8 }
func genFn0() int {
    var v7 bool = false
    fmt.Println(v7)
    if (375 <= 65535) {
        v7 = (42 == 550)
    } else {
        v8 := GenS1{F0: 0, F1: 4294967295, F2: 1}
        v7 = v7
        var v9 int16 = 0
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v7 = (!false)
            v12 := []int{32767, (^(fi11)), 127, fi10}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if v7 {
        switch 127 {
        case -1:
            if (!(!(("foo bar" + "foo bar") == ((("hello" + "") + "\n\t") + "\n\t")))) {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v14 := GenS2{F0: -14, F1: (fi13 >> ((314) & 63)), F2: 214}
                    var v15 uint32 = 65535
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v15 = v15
                        _ = fi16
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v7 = (-407 > 0)
            }
        case 42:
            v17 := GenS1{F0: 1000, F1: 18446744073709551615, F2: 0}
            _ = v17
        }
        var v18 float64 = (0.0 * (1e10 - (-2.25 - 0.1)))
        _ = v18
    }
    _ = v7
    return 127
}
func genFn1(p0 uint32) bool {
    if (!false) {
        v19 := []int{}
        fmt.Println(p0)
        for fi20 := 0; fi20 < 14; fi20++ {
            p0 = p0
            var v21 int32 = 1
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        p0 += (p0 ^ (p0 & (p0 - 4294967295)))
        _ = fi22
    }
    _ = p0
    return true
}

func main() {
    var v1 string = "\n\t"
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 uint32 = 42
    v5 := []int{-9223372036854775808, -728, 178, -9223372036854775808}
    v6 := make(map[int]string)
    v4 -= uint32(v2)
    v1 = "foo bar"
    v2 += 42
    _ = len(v5)
    var v23 uint8 = 42
    v1 = v1
    var v24 int16 = -1
    if false {
        v1 = v1
    }
    v25 := GenS1{F0: 1, F1: 4294967295, F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
