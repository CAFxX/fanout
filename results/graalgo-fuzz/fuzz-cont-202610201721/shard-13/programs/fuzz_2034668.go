package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 uint64) uint {
    v5 := []int{2147483647, 9223372036854775807, 2147483647}
    switch 966 {
    case 9223372036854775807:
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    case -960:
        if true {
            v5 = append(v5, 128)
            if p0 {
                var v6 float32 = 1e10
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: 865}
                    var v9 int = fi7
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                p1 -= 18446744073709551615
                _ = v6
            } else {
                p0 = (p0 || (!false))
                var v10 float64 = 0.0
                _ = v10
            }
        }
    }
    switch ("foo bar" + ("" + "hello")) {
    case "":
        fmt.Println(p0, p0)
    case "a":
        v5 = append(v5, -434)
    case "hello":
        var v11 int64 = 1
        _ = v11
    }
    var v12 uint32 = 1
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return 42
}
func genFn1(p0 int8) int16 {
    p0 -= (p0 - 127)
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 int = 4294967295
    var v3 int32 = -2147483648
    var v4 int16 = -567
    var v13 uint64 = 0
    var v14 int = v2
    v14 += -862
    var v15 int = v2
    v16 := GenS1{F0: 1}
    var v17 int32 = 242
    v18 := GenS2{F0: 0}
    v19 := GenS1{F0: 32767}
    fmt.Println(v3, v1, v15)
    v20 := GenS2{F0: 610}
    v21 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
