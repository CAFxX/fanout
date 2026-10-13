package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    v3 := []uint32{2147483648, 0, 2147483647}
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    v6 := []int{-9223372036854775808, 403, -650}
    if (1000 < 42) {
        var v7 bool = (0 > 0)
        v7 = (!((1.5 * 3.14159) <= 1.5))
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi8
        }
        var v9 uint8 = 230
        _ = v9
    }
    var v10 uint16 = 65535
    var v11 uint64 = 530
    var v12 int64 = 42
    v13 := GenS1{F0: 0}
    if false {
        if ("a" < "hello") {
            for fi14 := 0; fi14 < 17; fi14++ {
                switch "a" {
                case "a":
                    v5 = 0
                case "foo bar":
                    v4 = 0
                }
                _ = fi14
            }
        }
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v12, fi15)
            v12 -= v4
            v12 += 1
            _ = fi15
        }
        if (!false) {
            var v16 float32 = ((float32((0.0 + 0.1)) * 1e10) - ((0.0 + 1.5) + 1e10))
            _ = v16
        } else {
            delete(v1, 0)
            v5 = v5
        }
        delete(v1, 42)
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 string = "hello"
            fmt.Println(fi18, v5, v11)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v4 = -133
    v20 := v1[487]
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS1{F0: 63}
        v10 = v10
        _ = fi21
        _ = v22
    }
    var v23 float64 = (((((1.5 / (3.14159 * 0.1)) - (float64((1e10 - (-2.25 / 1.0))) + -2.25)) + float64(float32(-2.25))) + 1.5) * 0.0)
    fmt.Println(v20, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
