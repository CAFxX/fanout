package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = 1
        v8 := GenS1{F0: true}
        switch (fi6 % ((-2147483648) | 1)) {
        case 1:
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: false}
                v7 = (v7 & v7)
                v11 := []int{int(v7), -2147483648, 596, fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
        case -9223372036854775808:
            var v12 uint = 255
            _ = v12
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v13 uint8 = 244
    v14 := []int{127, -9223372036854775808, int(v13)}
    _ = v13
    _ = v14
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (3.14159 * 0.1)
    var v3 float32 = (v2 - v2)
    var v4 float32 = (float32((1e10 / (((1.5 * 0.1) / -2.25) + (3.14159 * 1e10)))) * v2)
    v5 := []uint32{2, 647}
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(fi15, v2, fi15)
        _ = fi15
    }
    v3 -= float32(1.5)
    switch ("foo bar" + "") {
    case "\n\t":
        _ = len(v1)
    case "":
        fmt.Println(v4, v3)
    default:
        fmt.Println(v4, v4, v3)
    }
    var v16 uint = 0
    v2 += ((v3 + 3.14159) - (1e10 - 0.1))
    fmt.Println(v2, v16)
    var v17 uint64 = 2
    if false {
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 int32 = -2147483648
                fmt.Println(v2, fi18)
                v21 := GenS1{F0: false}
                var v22 int16 = int16(v16)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            v16 += (^(v16))
            _ = fi18
        }
    }
    var v23 int32 = 479
    for fi24 := 0; fi24 < 17; fi24++ {
        v4 -= v2
        _ = fi24
    }
    v1[("foo bar" + (("\n\t" + "a") + ("" + "\n\t")))] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
