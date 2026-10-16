package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    var v3 uint64 = 523
    v4 := make(map[string]int)
    var v5 int64 = 42
    var v6 float32 = (0.1 + float32(float64(3.14159)))
    var v7 string = "hello"
    var v8 uint16 = 1
    if true {
        if true {
            _ = len(v4)
        }
    } else {
        if (0 == 65535) {
            v6 += 0.0
            var v9 uint8 = 42
            v10 := []int{-1, 334, -1, 128}
            _ = v10
            _ = v9
        }
    }
    var v11 int = -524
    switch v7 {
    case "\n\t":
        v12 := GenS1{F0: 255, F1: (v1 * (v1 ^ 2147483647)), F2: (-2.25 >= (1.5 + 1e10))}
        _ = v12
    default:
        var v13 int16 = -1
        _ = v13
    }
    var v14 int64 = (v5 + v5)
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 128, F1: -1, F2: false}
        _ = fi15
        _ = v16
    }
    if false {
        if (true || (!true)) {
            v17 := GenS2{F0: 2147483647, F1: 999}
            _ = v17
        } else {
            var v18 int64 = v5
            var v19 float64 = (1.5 * 1e10)
            _ = v18
            _ = v19
        }
        fmt.Println(v14, v1, v8)
    }
    v3 += 0
    for fi20 := 0; fi20 < 5; fi20++ {
        if true {
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := v2[4294967295]
                _ = fi21
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                switch (v11 << ((uint(v3)) & 63)) {
                case 127:
                    v24 := []int{127, -9223372036854775808, v11}
                    _ = v24
                default:
                    var v25 uint32 = 2147483648
                    _ = v25
                }
                _ = fi23
            }
            fmt.Println(v8, fi20)
            var v26 int = 0
            _ = v26
        }
        switch ((v7 + v7) + (v7 + ("foo bar" + ("" + "foo bar")))) {
        case "":
            if true {
                v27 := GenS2{F0: 65535, F1: 761}
                var v28 int64 = v5
                _ = v27
                _ = v28
            }
        case "foo bar":
            switch "\n\t" {
            case "hello":
                var v29 int64 = v14
                _ = v29
            }
        default:
            fmt.Println(v8, v1)
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
