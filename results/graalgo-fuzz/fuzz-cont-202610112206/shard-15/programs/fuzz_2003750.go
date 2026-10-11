package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int64, p1 int, p2 uint16) int {
    p1 = (p1 ^ (p1 ^ p1))
    p1 += p1
    v4 := []int{}
    var v5 uint64 = 42
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 int) int32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        if ((1e10 <= float32(0.0)) && (884 > 1)) {
            var v7 bool = true
            v8 := []int{}
            _ = v7
            _ = v8
        }
        fmt.Println(fi6, p0)
        _ = fi6
    }
    v9 := GenS1{F0: ("\n\t" + ("\n\t" + "a")), F1: 2147483647, F2: -32768}
    var v10 uint64 = 1
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 int = 128
    var v2 uint8 = 42
    var v3 uint64 = 443
    var v11 string = "hello"
    var v12 float32 = 3.14159
    v2 += (v2 * v2)
    var v13 bool = (!false)
    switch (v1 & genFn0(147, v1, 32767)) {
    case 127:
        var v14 int8 = -82
        _ = v14
    case -9223372036854775808:
        switch "a" {
        case "hello":
            if ((v11 >= ("\n\t" + (("hello" + "hello") + v11))) && v13) {
                v15 := GenS1{F0: "\n\t", F1: 757, F2: -32768}
                v16 := []int{genFn0(1, genFn0(689, (v1 + (v1 / ((v1) | 1))), 318), 871), v1, 9223372036854775807, -2147483648}
                var v17 int32 = int32(v2)
                _ = v15
                _ = v16
                _ = v17
            }
        case "":
            v18 := GenS1{F0: "foo bar", F1: 672, F2: 1}
            _ = v18
        }
    case 128:
        fmt.Println(v2)
    }
    v19 := []int{v1, (v1 ^ v1), genFn0(1, v1, 65535)}
    if (4294967295 > 2147483648) {
        v20 := GenS1{F0: v11, F1: 437, F2: -248}
        _ = v20
    } else {
        v13 = v13
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(v13, v13, v1)
            _ = len(v19)
            _ = fi21
        }
    }
    var v22 int32 = 42
    if v13 {
        for fi23 := 0; fi23 < 9; fi23++ {
            v13 = v13
            _ = fi23
        }
        var v24 int8 = 42
        _ = v24
    }
    var v25 int16 = -32768
    v13 = v13
    if (false || ((-61 == -1) && (350 <= 18446744073709551615))) {
        v12 += v12
        v26 := GenS2{F0: 106}
        _ = v26
    } else {
        if (v25 > (v25 - (^(v25)))) {
            v19 = append(v19, 2147483647)
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := GenS1{F0: ((v11 + "") + v11), F1: 1, F2: v25}
                v13 = v13
                _ = fi27
                _ = v28
            }
            var v29 int64 = 1
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
