package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }

func main() {
    var v1 bool = (128 == 42)
    var v2 float32 = float32(1e10)
    v3 := make(map[string]bool)
    var v4 float32 = v2
    var v5 float64 = 1e10
    var v6 float64 = v5
    for fi7 := 0; fi7 < 9; fi7++ {
        v6 = (1.5 - 0.0)
        switch fi7 {
        case -811:
            fmt.Println(v6, v1)
        case 32767:
            var v8 uint8 = 0
            _ = v8
        default:
            for fi9 := 0; fi9 < 10; fi9++ {
                switch 1 {
                case -454:
                    v10 := GenS1{F0: -1}
                    _ = v10
                }
                v11 := v3["foo bar"]
                fmt.Println(v1)
                _ = fi9
                _ = v11
            }
        }
        if (fi7 != (^(fi7))) {
            var v12 string = ""
            _ = v12
        } else {
            v13 := []int{}
            v14 := GenS1{F0: 1}
            v6 = v6
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    v15 := GenS2{F0: 127, F1: 949, F2: 18446744073709551615}
    v2 = v2
    v2 = (((v4 * 0.0) - float32(v5)) + (0.0 + (float32(((-2.25 * v6) / v5)) / v2)))
    v2 = v2
    switch "foo bar" {
    case "a":
        var v16 uint = 255
        _ = v16
    }
    if v1 {
        v3["foo bar"] = (42 > 1)
        var v17 uint8 = 42
        if (42 == -764) {
            if true {
                fmt.Println(v1)
                var v18 int8 = 127
                fmt.Println(v5)
                _ = v18
            }
            var v19 uint64 = 42
            var v20 float32 = (v2 - v4)
            _ = v19
            _ = v20
        } else {
            v1 = (!((!(255 == 783)) && true))
            v1 = v1
        }
        _ = v17
    }
    v6 = (v6 / (v6 / ((v5 + (3.14159 / v6)) * -2.25)))
    v5 += float64(float32((v5 / v5)))
    for fi21 := 0; fi21 < 8; fi21++ {
        v2 = (float32(0.0) - v2)
        v5 = 0.1
        switch (^(fi21)) {
        case 42:
            v22 := GenS1{F0: 1}
            _ = v22
        }
        _ = fi21
    }
    v6 = 1.5
    v23 := GenS2{F0: -30, F1: 4294967295, F2: 4294967295}
    v5 -= v5
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := []int{(fi24 / ((32767) | 1)), 2147483647, fi24, (^(fi24))}
        delete(v3, "\n\t")
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
