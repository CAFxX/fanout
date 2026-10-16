package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 int8, p1 int16) bool {
    if (float32(1.5) > (1e10 * (((0.1 + (1.5 / 1.0)) - 3.14159) - 1.5))) {
        var v6 int32 = 1
        _ = v6
    } else {
        var v7 float32 = -2.25
        var v8 int8 = (-(p0))
        if false {
            v9 := GenS1{F0: 1, F1: 209}
            _ = v9
        }
        _ = v7
        _ = v8
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    p1 = p1
    for fi12 := 0; fi12 < 5; fi12++ {
        switch 1 {
        case 787:
            p1 = p1
        case 2147483647:
            p0 = (p0 | p0)
        default:
            if (!true) {
                var v13 uint16 = 930
                _ = v13
            } else {
                switch (fi12 >> ((4294967295) & 63)) {
                case 32767:
                    v14 := []int{(fi12 + fi12)}
                    _ = v14
                }
            }
        }
        switch "\n\t" {
        case "a":
            p1 = p1
        case "foo bar":
            v15 := GenS2{F0: p1}
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 1
    var v3 uint8 = 0
    var v4 int8 = -1
    v5 := make(map[string]uint32)
    switch "\n\t" {
    case "foo bar":
        for fi16 := 0; fi16 < 1; fi16++ {
            v3 += uint8(v2)
            fmt.Println(v3, fi16, v2)
            _ = fi16
        }
    }
    var v17 bool = (float32(0.0) <= 0.0)
    v18 := GenS2{F0: -32768}
    var v19 bool = ((-2147483648 < 884) && v17)
    switch "" {
    case "foo bar":
        if false {
            v17 = true
        }
    case "\n\t":
        var v20 int32 = -158
        _ = v20
    case "a":
        v21 := []int{-538, 755, 127}
        _ = v21
    }
    var v22 uint = v2
    v22 = v2
    v5["foo bar"] = 2147483647
    fmt.Println(v2, v22)
    v23 := v1[32768]
    v2 = (v22 * 65535)
    v24 := []int{(v23 % ((804) | 1)), 32768, (v23 - 9223372036854775807)}
    var v25 string = ("\n\t" + "")
    if true {
        v4 = v4
        v25 = (v25 + v25)
        v26 := GenS2{F0: 1}
        _ = v26
    }
    v27 := GenS3{F0: -9223372036854775808, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
