package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 int64) uint16 {
    var v5 uint32 = 42
    var v6 uint8 = 44
    v5 -= (^(v5))
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 int32) int16 {
    fmt.Println(p0, p1)
    var v8 int16 = 607
    switch "foo bar" {
    case "a":
        v9 := []int{}
        _ = v9
    case "\n\t":
        p1 = (p1 & 623)
    case "":
        var v10 uint32 = 65535
        _ = v10
    }
    v11 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return (v8 | (v8 ^ v8))
}

func main() {
    var v1 uint = 65535
    var v2 int = 4294967295
    var v3 float32 = 0.1
    var v4 int32 = 2147483647
    v3 += v3
    for fi12 := 0; fi12 < 14; fi12++ {
        if (v3 < v3) {
            fmt.Println(fi12, v1)
            v13 := []int{(-(v2))}
            _ = v13
        }
        _ = fi12
    }
    switch (-(v2)) {
    case -9223372036854775808:
        var v14 uint = v1
        _ = v14
    case -2147483648:
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{fi15}
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if true {
            v18 := []int{(^(fi17)), (v2 >> (((v1 >> ((v1) & 63))) & 63))}
            _ = v18
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            v1 = 4294967295
            _ = fi19
        }
        v20 := []int{v2, 9223372036854775807}
        _ = fi17
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 1; fi21++ {
            v3 = ((3.14159 / -2.25) - v3)
            var v22 uint = v1
            _ = fi21
            _ = v22
        }
        v3 += v3
    } else {
        v23 := GenS1{F0: v1, F1: (v2 / ((42) | 1))}
        _ = v23
    }
    v1 = 2
    v24 := []int{}
    if (172 != -1) {
        _ = len(v24)
    } else {
        if (!(-113 != -107)) {
            var v25 string = ""
            _ = v25
        }
        v3 += (v3 * 1.5)
        if len(v24) > 0 {
            v24[0] = v2
        }
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
