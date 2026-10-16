package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
func genFn0(p0 uint16, p1 bool) int {
    v7 := []int{32767, -1}
    p0 += 32767
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int64, p1 string) bool {
    if ((!((1.5 * float32((-2.25 * (float64(float32(3.14159)) / (0.1 / ((3.14159 - ((((-2.25 + 0.1) + 3.14159) + -2.25) - (1.5 * (-2.25 - 1.5)))) - ((1e10 / 3.14159) * ((3.14159 + 3.14159) + float64(0.1))))))))) <= 3.14159)) || (!false)) {
        var v8 int16 = 924
        fmt.Println(p1)
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    var v10 bool = ((!(!(!true))) || false)
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint8 = 42
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 2147483648, F1: 42, F2: fi11}
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return false
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[int]string)
    var v3 int16 = 248
    v4 := make(map[uint32]bool)
    var v5 int8 = 1
    var v6 int32 = 42
    if (false || true) {
        fmt.Println(v3)
        for fi15 := 0; fi15 < 3; fi15++ {
            delete(v4, 0)
            switch genFn0(0, false) {
            case 127:
                _ = len(v1)
            case -352:
                v16 := GenS1{F0: 2147483648, F1: 65535, F2: 121}
                _ = v16
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi15
        }
    }
    switch 2147483647 {
    case -9223372036854775808:
        var v17 uint16 = 255
        _ = v17
    case 9223372036854775807:
        switch (("\n\t" + "\n\t") + "") {
        case "foo bar":
            var v18 int8 = -128
            _ = v18
        case "hello":
            v6 -= v6
        case "":
            v3 = (^(v3))
        }
    default:
        var v19 int64 = 0
        _ = v19
    }
    var v20 int8 = 127
    var v21 uint16 = 42
    v22 := GenS2{F0: 65535, F1: 32767, F2: 255}
    var v23 uint16 = (v21 ^ v21)
    fmt.Println(v6, v5, v20)
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 uint64 = 2
        _ = fi24
        _ = v25
    }
    if true {
        var v26 float32 = 0.1
        _ = v26
    } else {
        var v27 string = "a"
        fmt.Println(v20)
        _ = v27
    }
    v23 += v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
