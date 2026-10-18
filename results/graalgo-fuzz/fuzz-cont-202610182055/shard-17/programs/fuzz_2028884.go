package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint32) int {
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = 815
        _ = fi7
    }
    _ = p0
    _ = v6
    return 735
}

func main() {
    var v1 uint = 65535
    var v2 int32 = -310
    v3 := make(map[uint32]int)
    var v4 uint32 = 488
    var v5 uint16 = 32767
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v2, v5, v4)
        _ = fi8
    }
    v5 = v5
    if true {
        v4 = 932
    } else {
        v2 += (v2 % (((v2 % (((v2 >> ((v1) & 63))) | 1))) | 1))
        var v9 uint64 = 894
        var v10 int = -733
        _ = v10
        _ = v9
    }
    v11 := v3[(v4 * v4)]
    v5 += 42
    v4 -= (v4 << ((18446744073709551615) & 63))
    v12 := GenS1{F0: v11, F1: ("\n\t" + "\n\t"), F2: (-(v11))}
    v1 -= 2
    v11 += 629
    if true {
        var v13 int64 = 0
        _ = v13
    }
    v14 := GenS1{F0: v11, F1: (((("a" + "") + "\n\t") + "foo bar") + ("\n\t" + "")), F2: (v11 | (v11 >> ((v1) & 63)))}
    _ = len(v3)
    var v15 float64 = (-2.25 + -2.25)
    if ((-(v2)) > v2) {
        if false {
            var v16 uint = v1
            var v17 float32 = (1e10 * 0.0)
            var v18 int8 = 113
            _ = v16
            _ = v17
            _ = v18
        } else {
            v11 = (v11 / ((v11) | 1))
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint = v1
                _ = fi19
                _ = v20
            }
            if ((1e10 > 1e10) && (true && true)) {
                v21 := GenS2{F0: 18446744073709551615, F1: (v2 | (v2 * v2))}
                var v22 string = "\n\t"
                v23 := []int{}
                _ = v21
                _ = v22
                _ = v23
            }
        }
        if true {
            v1 = (^(v1))
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS2{F0: 42, F1: -22}
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS2{F0: 56, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
