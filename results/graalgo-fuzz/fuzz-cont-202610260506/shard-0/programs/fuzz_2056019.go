package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8) int64 {
    p0 = p0
    switch "\n\t" {
    case "\n\t":
        var v6 int16 = -70
        _ = v6
    case "":
        v7 := []int{788, 1, 2147483647}
        _ = v7
    }
    _ = p0
    return -9223372036854775808
}
func genFn1() uint8 {
    var v8 bool = true
    var v9 int8 = 127
    v10 := []int{}
    for fi11 := 0; fi11 < 12; fi11++ {
        v8 = (-1 <= (fi11 >> ((1) & 63)))
        var v12 float64 = 1e10
        _ = fi11
        _ = v12
    }
    var v13 uint = 4294967295
    _ = v10
    _ = v13
    _ = v8
    _ = v9
    return 225
}

func main() {
    var v1 int = 1
    var v2 uint64 = 42
    var v3 uint16 = 65535
    var v4 float32 = ((-2.25 * 1.5) / 1.5)
    var v5 int = v1
    for fi14 := 0; fi14 < 15; fi14++ {
        v1 -= 1
        _ = fi14
    }
    if (0.0 < (-2.25 - 3.14159)) {
        var v15 int64 = genFn0(0)
        var v16 int32 = 1
        _ = v15
        _ = v16
    } else {
        var v17 uint8 = 80
        _ = v17
    }
    var v18 int = 1
    var v19 uint16 = (v3 | v3)
    v20 := GenS1{F0: 983}
    var v21 float64 = float64(v4)
    v5 += v1
    v22 := GenS1{F0: 972}
    v1 += v5
    switch "hello" {
    case "\n\t":
        var v23 float32 = v4
        _ = v23
    }
    v24 := []int{}
    v2 += (v2 - (v2 << ((171) & 63)))
    var v25 int = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
