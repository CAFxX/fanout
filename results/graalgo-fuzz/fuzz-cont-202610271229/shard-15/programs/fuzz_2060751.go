package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
func genFn0(p0 uint16, p1 int, p2 bool) uint {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(p2)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 string) int64 {
    var v8 float32 = ((-2.25 / 3.14159) - float32((((-2.25 + 1.5) - 3.14159) / 1e10)))
    p0 = p0
    _ = p0
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]string)
    var v3 uint = 1
    v4 := []uint32{0, 2147483648}
    v5 := []int{2147483647, -9223372036854775808, -109}
    var v6 float64 = 0.0
    v2[42] = "\n\t"
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint16 = 675
            v6 = (v6 * (v6 - (3.14159 + (v6 + 1e10))))
            _ = fi10
            _ = v11
        }
        fmt.Println(v6, fi9, v1)
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int16 = 1
            var v14 string = "foo bar"
            v3 += (v3 >> ((v3) & 63))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    switch 684 {
    case 128:
        var v15 bool = true
        _ = v15
    }
    v4 = append(v4, 4294967295)
    v1 = ((v1 - float32(((((float64(1.5) / ((v6 / v6) + ((float64(v1) * v6) / (-2.25 * -2.25)))) / (3.14159 / 0.1)) + v6) * (1e10 * float64((float32(v6) - -2.25)))))) * v1)
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            if (!(!(0 > v3))) {
                var v18 float64 = v6
                var v19 uint8 = 188
                if (uint64(v19) <= 42) {
                    var v20 int8 = 100
                    _ = v20
                } else {
                    var v21 int16 = 0
                    fmt.Println(v1, v1)
                    _ = v21
                }
                _ = v18
                _ = v19
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v22 uint32 = 594
            _ = fi17
            _ = v22
        }
        _ = fi16
    }
    switch "foo bar" {
    case "foo bar":
        var v23 float32 = ((v1 * v1) + v1)
        _ = v23
    default:
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v1)
            _ = fi24
        }
    }
    if (-9223372036854775808 == 9223372036854775807) {
        fmt.Println(v6)
    } else {
        if (true || false) {
            var v25 int = -401
            v1 -= 1.5
            _ = v25
        } else {
            v26 := GenS1{F0: -69, F1: 2147483647, F2: "a"}
            v3 = v3
            v3 += v3
            _ = v26
        }
        _ = len(v2)
    }
    v1 = 1.5
    for fi27 := 0; fi27 < 9; fi27++ {
        fmt.Println(v1)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
