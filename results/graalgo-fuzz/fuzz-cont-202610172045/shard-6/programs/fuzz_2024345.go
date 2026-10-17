package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 int64, p1 int16, p2 uint16) uint32 {
    v6 := GenS1{F0: (p2 >> ((42) & 63)), F1: 1, F2: (p2 * (p2 + (p2 % ((p2) | 1))))}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 891
        _ = fi7
        _ = v8
    }
    fmt.Println(p1)
    if (0 > p2) {
        p2 -= (p2 / (((p2 * 255)) | 1))
    }
    v9 := []int{2147483647, 0, 0, -221}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 65535
}

func main() {
    var v1 int = 199
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    var v4 int8 = 75
    var v5 uint64 = 926
    v10 := GenS1{F0: 1, F1: 42, F2: 32767}
    fmt.Println(v1)
    if (v3 >= v3) {
        var v11 float64 = 1.5
        v1 = v1
        fmt.Println(v11, v5, v3)
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch ("a" + ("\n\t" + ((("foo bar" + ("a" + "")) + ("a" + "\n\t")) + ("a" + "a")))) {
        case "a":
            v4 = (-(v4))
        default:
            var v13 float64 = 1.5
            _ = v13
        }
        _ = fi12
    }
    var v14 string = "\n\t"
    v4 -= -128
    var v15 uint32 = 341
    v16 := []int{(^(v1)), (v1 / ((32768) | 1)), -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
