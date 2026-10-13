package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 int64) uint16 {
    if (2147483647 >= 1) {
        p0 += (p0 * p0)
        var v6 int16 = 42
        v6 -= -964
        _ = v6
    } else {
        fmt.Println(p0)
    }
    if false {
        p0 += -1
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 1
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    return 990
}
func genFn1(p0 int) uint8 {
    p0 += p0
    p0 = -2147483648
    _ = p0
    return 1
}

func main() {
    var v1 int32 = -383
    v2 := []int{1, 2147483647}
    v3 := []int{4294967295, -2147483648, 32768, -2147483648}
    var v4 uint16 = 677
    v5 := []string{}
    v4 -= 100
    v4 -= v4
    v9 := GenS1{F0: 42, F1: 18446744073709551615}
    var v10 uint8 = 127
    v5 = append(v5, ((("" + "hello") + "foo bar") + ("" + (((("foo bar" + "hello") + "hello") + "foo bar") + ((("\n\t" + "foo bar") + ("hello" + "a")) + "hello")))))
    v11 := GenS1{F0: 753, F1: 18446744073709551615}
    v12 := GenS1{F0: 0, F1: 0}
    fmt.Println(v1)
    fmt.Println(v1)
    var v13 string = "hello"
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: -1, F1: 2}
            if len(v5) > 0 {
                v5[0] = (v13 + v13)
            }
            switch fi15 {
            case 0:
                var v17 bool = true
                _ = v17
            case 2147483647:
                v4 = v4
            }
            _ = fi15
            _ = v16
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            v4 = 42
            fmt.Println(fi14, v4, v10)
            _ = len(v2)
            _ = fi18
        }
        v10 -= (v10 % (((v10 / (((v10 / (((v10 & 8)) | 1))) | 1))) | 1))
        _ = fi14
    }
    var v19 uint64 = 4294967295
    var v20 float64 = 0.0
    var v21 uint64 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
