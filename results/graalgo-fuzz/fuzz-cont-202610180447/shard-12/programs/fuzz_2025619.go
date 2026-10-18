package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0(p0 string, p1 uint16, p2 int32) int64 {
    switch "foo bar" {
    case "\n\t":
        var v7 bool = false
        _ = v7
    }
    v8 := []int{32767, -942, 0, -2147483648}
    var v9 uint32 = 42
    var v10 int32 = p2
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 570
}

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    var v3 int64 = 42
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    switch ("a" + "a") {
    case "foo bar":
        var v11 float64 = 0.1
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 11; fi12++ {
            v4 += (v4 + v4)
            _ = fi12
        }
    case "hello":
        if ("foo bar" <= "\n\t") {
            fmt.Println(v1, v4)
            var v13 string = ""
            _ = v13
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                v1 = v1
                _ = fi14
            }
        }
    default:
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: v3, F1: fi15, F2: 766}
            v17 := GenS1{F0: v3, F1: fi15, F2: 855}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    switch "" {
    case "a":
        if (109 <= 255) {
            var v18 float64 = -2.25
            _ = v18
        } else {
            v3 = (v3 | (^(v3)))
        }
    case "\n\t":
        v1 = v1
    case "hello":
        for fi19 := 0; fi19 < 20; fi19++ {
            v3 = v3
            switch "a" {
            case "foo bar":
                v20 := GenS2{F0: 42, F1: 122}
                _ = v20
            }
            if v1 {
                if (!((float32((1e10 * -2.25)) * (float32(0.0) * 1.5)) == ((v4 + v4) * (((0.0 * float32((-2.25 - float64(v4)))) / v4) / v4)))) {
                    v3 = v3
                    _ = len(v6)
                }
                v21 := []int{-927, (fi19 >> ((619) & 63)), (fi19 * fi19)}
                var v22 float32 = v4
                _ = v21
                _ = v22
            } else {
                v4 = float32(1e10)
            }
            _ = fi19
        }
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS1{F0: v3, F1: (fi23 / ((-667) | 1)), F2: 61}
            v1 = true
            _ = fi23
            _ = v24
        }
    }
    switch "" {
    case "hello":
        var v25 uint64 = 657
        _ = v25
    case "":
        v26 := v5[2147483647]
        _ = v26
    default:
        var v27 int = -1
        _ = v27
    }
    fmt.Println(v3, v4)
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int8 = -1
        _ = len(v2)
        var v30 uint16 = 668
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 float64 = -2.25
    fmt.Println(v31, v4, v4)
    v3 = -9223372036854775808
    for fi32 := 0; fi32 < 2; fi32++ {
        switch fi32 {
        case -6:
            v33 := []int{4294967295, fi32, fi32, (fi32 >> ((391) & 63))}
            _ = v33
        default:
            v1 = (true || (42 >= 269))
        }
        fmt.Println(v31, v3, v1)
        var v34 string = ""
        _ = fi32
        _ = v34
    }
    if (!v1) {
        v35 := GenS2{F0: 627, F1: 42}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
