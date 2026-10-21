package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint16 {
    var v6 uint16 = 604
    if (-128 < -128) {
        v7 := []int{792}
        _ = v7
    }
    _ = v6
    return 32767
}
func genFn1(p0 int32, p1 uint64, p2 bool) int {
    v8 := []int{}
    var v9 bool = (73 >= 1)
    var v10 uint32 = 2
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int32 = p0
        var v13 uint64 = (^(p1))
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 858
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 0
    var v3 string = ""
    var v4 float64 = 0.0
    v5 := make(map[string]uint32)
    var v14 int = -634
    v2 += v2
    _ = len(v1)
    if ((((false && (true && (false || (!(true && (!false)))))) && ((v4 * (v4 / (((1e10 / v4) + (((v4 * v4) * (((v4 / v4) + 0.0) * ((v4 * float64(-2.25)) + ((1e10 + -2.25) + (v4 * float64(3.14159)))))) / v4)) - (0.1 - 1.5)))) != 1e10)) || false) || false) {
        var v15 uint64 = 2
        _ = v15
    }
    var v16 uint32 = 4294967295
    fmt.Println(v2, v4)
    var v17 string = "a"
    if true {
        v18 := []int{(v14 + (^(v14)))}
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 int16 = -1
            v16 = (v16 + (v16 & v16))
            v2 -= v2
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        if (false || false) {
            if true {
                v14 = genFn1(42, 809, ((true && ((!true) && false)) || true))
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 float32 = float32(v4)
                    _ = fi21
                    _ = v22
                }
            } else {
                var v23 bool = (42 == v2)
                _ = v23
            }
            if true {
                var v24 uint32 = v16
                v16 = (^(v24))
                v25 := []int{128, 0, (v14 | (v14 ^ (v14 % ((v14) | 1)))), (v14 * v14)}
                _ = v24
                _ = v25
            }
        } else {
            var v26 int8 = 0
            var v27 float64 = 3.14159
            _ = v26
            _ = v27
        }
        var v28 int16 = -631
        _ = v28
    }
    var v29 int16 = -866
    fmt.Println(v4, v3, v29)
    _ = len(v1)
    v30 := []int{genFn1(2147483647, 18446744073709551615, true), genFn1(2147483647, 0, false)}
    var v31 uint16 = 42
    v32 := GenS1{F0: (v14 & 192)}
    fmt.Println(v4, v31, v14)
    switch "\n\t" {
    case "":
        if (!false) {
            switch (v14 & (v14 ^ genFn1(-1, 1, false))) {
            case 32768:
                var v33 int8 = (v2 % ((-1) | 1))
                _ = v33
            default:
                var v34 int = 9223372036854775807
                _ = v34
            }
        }
    default:
        v35 := GenS1{F0: 64}
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
