package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 string, p2 uint8) string {
    var v4 bool = false
    p2 -= (p2 * p2)
    v5 := []int{127, 4294967295, 1, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint32, p1 uint64) int32 {
    p1 = p1
    v6 := []int{32767, 703, 42}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float64 = float64(3.14159)
        _ = fi7
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    if false {
        v9 = 405
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 500
    v2 := make(map[string]int)
    var v3 int8 = 0
    delete(v2, "foo bar")
    if ((v1 > v1) && true) {
        v10 := GenS1{F0: -1}
        if (1.5 <= ((((0.1 + -2.25) + (((1e10 + 1e10) / 1.0) * 1e10)) - 0.1) + 0.1)) {
            var v11 int16 = -821
            _ = v11
        }
        var v12 float64 = ((-2.25 + (0.1 / float64((1e10 + 1.5)))) / 1.0)
        _ = v10
        _ = v12
    } else {
        v3 -= v3
    }
    v1 = -7
    if true {
        if true {
            v13 := GenS1{F0: -2147483648}
            v14 := []int{0, 128, 42}
            switch "a" {
            case "":
                v1 = (v1 << ((0) & 63))
            default:
                v15 := v2["a"]
                _ = v15
            }
            _ = v13
            _ = v14
        } else {
            fmt.Println(v3)
        }
    }
    v16 := GenS1{F0: 2147483647}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int = fi17
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: (-2147483648 != 0)}
    v20 := GenS2{F0: (("foo bar" + "foo bar") < "")}
    v21 := GenS1{F0: genFn1(2147483647, 2)}
    var v22 int64 = 42
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int64 = 1
        if true {
            fmt.Println(v22, v1, v3)
        }
        _ = fi23
        _ = v24
    }
    v25 := []int{4294967295}
    v3 += v3
    v1 -= (v1 << ((18446744073709551615) & 63))
    var v26 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
