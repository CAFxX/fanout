package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16) int16 {
    v6 := GenS1{F0: 2}
    var v7 float64 = -2.25
    var v8 uint16 = 65535
    var v9 int8 = 127
    var v10 uint = 304
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (p0 % (((-(p0))) | 1))
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -1
    var v3 int = 42
    var v4 int8 = -29
    v5 := make(map[string]string)
    v1[(v3 >> ((1) & 63))] = "a"
    if true {
        var v11 uint8 = 255
        _ = v11
    } else {
        v12 := []int{v3}
        if len(v12) > 0 {
            v12[0] = (v3 & v3)
        }
        v4 += (v4 + (v4 % (((^(v4))) | 1)))
        _ = v12
    }
    switch "" {
    case "\n\t":
        for fi13 := 0; fi13 < 13; fi13++ {
            switch (fi13 * fi13) {
            case 127:
                v2 = genFn0(v2)
            case -890:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v4 = v4
                    _ = fi14
                }
            case 0:
                delete(v1, (fi13 * fi13))
            }
            v3 -= v3
            _ = fi13
        }
    }
    v15 := GenS1{F0: 905}
    v16 := GenS1{F0: 65535}
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint16 = 589
        v18 = (v18 / ((v18) | 1))
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v2, fi17)
            v2 = genFn0((v2 % (((^(v2))) | 1)))
            v4 = v4
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v4)
    var v20 bool = (int64(v4) < 42)
    delete(v5, (("a" + "foo bar") + ""))
    switch (v3 % ((32767) | 1)) {
    case -9223372036854775808:
        if v20 {
            switch v3 {
            case 767:
                var v21 int64 = -9223372036854775808
                _ = v21
            case 603:
                v22 := []int{(v3 ^ v3), -282, (v3 & (v3 / (((v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))) | 1))), v3}
                _ = v22
            }
        }
    case 42:
        fmt.Println(v3)
    case 32768:
        fmt.Println(v2, v2, v2)
    default:
        if (2147483647 > 191) {
            var v23 float64 = 0.0
            _ = v23
        }
    }
    v24 := []int{}
    if ((2147483647 >= -193) && v20) {
        var v25 uint64 = 1
        v2 = v2
        v3 = (v3 * (^(v3)))
        _ = v25
    }
    var v26 float32 = 1.5
    fmt.Println(v2, v20)
    v27 := GenS1{F0: 302}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
