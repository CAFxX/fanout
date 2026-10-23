package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
func genFn0(p0 bool, p1 string, p2 int64) bool {
    var v5 int8 = 100
    v6 := GenS1{F0: p1, F1: 0}
    v7 := []int{128, -1, 42, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 int16 = -1
    var v2 int32 = 1
    var v3 bool = true
    var v4 int32 = (v2 + (^(v2)))
    v8 := GenS1{F0: "\n\t", F1: 0}
    v2 = v4
    var v9 uint = 42
    v10 := GenS1{F0: ("" + "foo bar"), F1: 806}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint16 = 65535
        if (!v3) {
            v9 += 671
            var v13 int64 = 1
            var v14 uint8 = 163
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v3 = (!(480 < 4294967295))
    v15 := GenS1{F0: ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + "")), F1: 65535}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 bool = v3
        _ = fi16
        _ = v17
    }
    v1 += (v1 ^ (^(v1)))
    if (!(v3 || (((1.5 * ((0.1 - 0.0) * ((1.5 * (0.1 / 1.0)) - float64(0.1)))) + 0.1) < (1e10 * (((1.5 + 0.1) / 1e10) / 1.0))))) {
        if (v3 && (!v3)) {
            v18 := []int{-734, 249, 377, -9223372036854775808}
            var v19 int16 = (v1 & (-(v1)))
            for fi20 := 0; fi20 < 3; fi20++ {
                v2 = v2
                _ = fi20
            }
            _ = v18
            _ = v19
        } else {
            v4 = v2
            var v21 int32 = v4
            var v22 int8 = 100
            _ = v21
            _ = v22
        }
        fmt.Println(v1, v4, v1)
    }
    v23 := []int{127, int(v1), 32768}
    if (0.0 < (((3.14159 + 3.14159) - -2.25) / 1.5)) {
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := GenS1{F0: (("" + "a") + "\n\t"), F1: 1}
                var v27 uint16 = 788
                _ = len(v23)
                v28 := GenS2{F0: v1, F1: fi24, F2: 0}
                var v29 uint = (v9 % ((v9) | 1))
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi24
        }
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 int64 = 828
        _ = fi30
        _ = v31
    }
    v32 := GenS1{F0: "hello", F1: 133}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
