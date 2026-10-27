package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 int64, p1 int32) int {
    v5 := GenS1{F0: "hello", F1: 1}
    v6 := GenS1{F0: ("" + "a"), F1: 2147483647}
    v7 := []int{32767, -9223372036854775808, 42, 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := []string{"\n\t", ("" + "\n\t"), ""}
    v2 := []uint32{}
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    v8 := GenS1{F0: "", F1: 0}
    v9 := GenS1{F0: "", F1: 0}
    fmt.Println(v3)
    var v10 int64 = 0
    for fi11 := 0; fi11 < 12; fi11++ {
        v3 = -1
        switch (-(fi11)) {
        case 128:
            v12 := GenS1{F0: "hello", F1: (fi11 - genFn0(-1, 42))}
            _ = v12
        case -240:
            if (1 >= -1) {
                v10 = (v10 ^ (^(v10)))
            }
        default:
            fmt.Println(v3, v3)
        }
        _ = fi11
    }
    if ((!(32767 != 0)) && (245 > -81)) {
        var v13 float64 = 3.14159
        v14 := v4[1]
        _ = v13
        _ = v14
    }
    if (4294967295 >= 4294967295) {
        if true {
            var v15 bool = false
            v15 = ((!(0 > 128)) && v15)
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint16 = 32767
                _ = len(v1)
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            var v18 float64 = (1.5 - 1.5)
            v19 := GenS1{F0: "hello", F1: genFn0((v10 % ((42) | 1)), -1)}
            _ = v18
            _ = v19
        }
        if (true || false) {
            v10 += v10
        }
    } else {
        for fi20 := 0; fi20 < 19; fi20++ {
            if false {
                v10 = 0
                fmt.Println(v3)
                v10 = (v10 * v10)
            } else {
                v10 = v10
                v21 := GenS2{F0: 979, F1: -2147483648}
                fmt.Println(v3, v10)
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS2{F0: 2, F1: 1}
                _ = fi22
                _ = v23
            }
            _ = fi20
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            v3 = v3
            v25 := GenS2{F0: 42, F1: 1}
            var v26 uint64 = 0
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 bool = false
    if v27 {
        v10 = v10
        switch "hello" {
        case "hello":
            v3 += (^(v3))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
