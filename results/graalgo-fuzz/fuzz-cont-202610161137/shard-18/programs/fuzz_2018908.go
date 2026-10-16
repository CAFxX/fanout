package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = (((float64(float32((3.14159 * 1.5))) - ((v1 - v1) * float64((1.5 - (1e10 / 0.1))))) * (((1e10 / (((v1 + float64(0.1)) + 0.0) + (1.5 * float64(1.5)))) * v1) - v1)) - 0.1)
    v3 := []int{32768, -594}
    v4 := GenS1{F0: 73}
    var v5 int64 = -9223372036854775808
    v6 := GenS1{F0: 545}
    v5 += v5
    _ = len(v3)
    v7 := GenS1{F0: 4294967295}
    if (false || true) {
        v8 := GenS1{F0: 454}
        v5 -= -9223372036854775808
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            _ = len(v3)
            var v10 uint64 = 18446744073709551615
            _ = fi9
            _ = v10
        }
        var v11 uint64 = 164
        _ = v11
    }
    switch ("\n\t" + ("" + "hello")) {
    case "\n\t":
        v3 = append(v3, -782)
    }
    if ("a" == "") {
        _ = len(v3)
        v5 += -1
        v12 := GenS1{F0: 2}
        _ = v12
    }
    var v13 float32 = 0.0
    v14 := GenS1{F0: 4294967295}
    v13 = 3.14159
    if ((1e10 / v2) >= (float64((v13 - 3.14159)) * float64(0.1))) {
        v2 += ((3.14159 * -2.25) / (v2 - 0.0))
    }
    fmt.Println(v13, v5)
    var v15 uint32 = 42
    v16 := GenS2{F0: 688, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
