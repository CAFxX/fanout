package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 float64 = 1.5
    var v4 string = "\n\t"
    v5 := v1[-9223372036854775808]
    fmt.Println(v4, v3, v3)
    switch (v4 + ((v4 + ("\n\t" + "foo bar")) + "hello")) {
    case "hello":
        fmt.Println(v4, v3)
    }
    if (v5 && (((v4 <= v4) && (18446744073709551615 < 304)) || (!(4294967295 <= 391)))) {
        v6 := GenS1{F0: 423, F1: 57}
        _ = v6
    }
    v7 := v2[(("foo bar" + "foo bar") + (v4 + "\n\t"))]
    if ((float32(0.0) + (0.1 / 0.1)) != 1e10) {
        var v8 float64 = (v3 + v3)
        _ = v8
    }
    var v9 uint16 = 558
    v10 := v2[(v4 + v4)]
    v11 := GenS1{F0: 1, F1: -123}
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS1{F0: 4294967295, F1: 100}
            var v15 float32 = 0.1
            v16 := GenS1{F0: 219, F1: 1}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 uint32 = 2
        _ = fi12
        _ = v17
    }
    switch -9223372036854775808 {
    case 579:
        if (v9 < v9) {
            var v18 uint = 824
            v7 = ""
            var v19 uint = v18
            _ = v18
            _ = v19
        }
    case 32767:
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 float32 = -2.25
            var v22 float32 = float32(-2.25)
            var v23 bool = (!(2 >= (v9 * 255)))
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    case 1:
        var v24 uint = 4294967295
        _ = v24
    default:
        v25 := GenS1{F0: 2, F1: -128}
        _ = v25
    }
    switch -9223372036854775808 {
    case 42:
        if v5 {
            var v26 uint32 = 2147483647
            v9 = (v9 * v9)
            _ = v26
        }
    case -9223372036854775808:
        v9 = v9
    case -1:
        if (1e10 != 1.5) {
            v7 = v7
            if v5 {
                fmt.Println(v3)
                var v27 int = 32767
                _ = v27
            } else {
                fmt.Println(v9)
            }
            v10 = "\n\t"
        }
    default:
        var v28 uint = 255
        _ = v28
    }
    fmt.Println(v4)
    var v29 int8 = 42
    v30 := GenS2{F0: (v29 << ((526) & 63)), F1: v29}
    v9 = (v9 << ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
