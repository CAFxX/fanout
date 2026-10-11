package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint16; F2 int }
func genFn0(p0 int, p1 int8, p2 bool) int {
    var v6 int32 = -58
    if (!(((-(p1)) != (p1 << ((42) & 63))) && p2)) {
        v7 := GenS1{F0: 1, F1: (!(!p2)), F2: -9223372036854775808}
        if ((255 == 1) && ((float64((1.5 * 0.1)) + 3.14159) < float64(1e10))) {
            var v8 int64 = 966
            _ = v8
        }
        v9 := GenS2{F0: true, F1: 9223372036854775807}
        _ = v7
        _ = v9
    }
    if false {
        if (0.1 != (3.14159 + 1.5)) {
            if ("\n\t" <= "a") {
                var v10 int16 = 32767
                p2 = (9223372036854775807 < 9223372036854775807)
                var v11 float64 = ((0.1 * 0.1) + 1e10)
                _ = v10
                _ = v11
            }
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 / (((p0 / ((9223372036854775807) | 1))) | 1))
}
func genFn1(p0 string, p1 uint8, p2 bool) uint {
    v12 := []int{41, genFn0(9223372036854775807, 30, (0.0 == -2.25)), 42}
    p2 = true
    switch 128 {
    case 32768:
        if (1 == 0) {
            for fi13 := 0; fi13 < 8; fi13++ {
                p1 -= (^(p1))
                var v14 int64 = -1
                var v15 int8 = 0
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            p0 = (("a" + "hello") + p0)
            var v16 int = 486
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 1
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    var v4 float32 = 0.1
    v5 := []int{}
    if true {
        v17 := GenS1{F0: 222, F1: (v4 < float32(0.1)), F2: -1}
        v18 := GenS3{F0: 3, F1: (v3 + v3)}
        _ = v17
        _ = v18
    }
    if true {
        var v19 int16 = (v3 | (v3 << ((genFn1(("a" + "hello"), 166, (551 <= 4294967295))) & 63)))
        _ = v19
    } else {
        v20 := GenS3{F0: 42, F1: v3}
        v21 := GenS2{F0: (1.5 > 1e10), F1: 42}
        var v22 int16 = 42
        _ = v20
        _ = v21
        _ = v22
    }
    v3 = 42
    switch (("" + ((((("\n\t" + "") + "a") + "hello") + "a") + "a")) + "hello") {
    case "":
        if (3.14159 >= 0.0) {
            v1 -= 32768
            var v23 float64 = 0.1
            _ = v23
        } else {
            if ((1.5 == 1e10) || false) {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := GenS4{F0: "foo bar"}
                    if false {
                        v3 -= v3
                        v3 -= 1000
                    }
                    v1 = fi24
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v1, v1)
                v26 := GenS1{F0: 85, F1: true, F2: 42}
                _ = v26
            }
        }
    }
    v27 := GenS5{F0: true, F1: 2, F2: (^(v1))}
    var v28 int16 = (-(v3))
    v3 += 1000
    v29 := GenS5{F0: (v1 < v1), F1: 0, F2: v1}
    var v30 bool = (float32(3.14159) > v4)
    var v31 int32 = 0
    fmt.Println(v4, v4, v4)
    v31 -= (v31 & v31)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
