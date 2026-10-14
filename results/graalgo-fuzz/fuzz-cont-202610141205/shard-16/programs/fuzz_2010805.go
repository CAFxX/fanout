package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint64, p1 int16, p2 uint64) uint8 {
    v4 := []int{}
    _ = len(v4)
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 -= p0
        var v6 string = ""
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]string)
    var v3 int64 = (v1 / ((v1) | 1))
    v7 := []int{}
    v1 = v1
    var v8 uint8 = 242
    v3 = 9223372036854775807
    if false {
        switch (("" + "hello") + "") {
        case "\n\t":
            var v9 uint16 = 65535
            _ = v9
        case "":
            if (((1.5 > float64(float32(0.0))) || true) && true) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 int8 = 0
                    var v12 uint = 985
                    v12 = (v12 * (v12 >> ((v12) & 63)))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v3 = v3
                if len(v7) > 0 {
                    v7[0] = 9223372036854775807
                }
            }
        }
        v13 := GenS1{F0: 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v1 -= v3
        _ = fi14
    }
    delete(v2, 32767)
    fmt.Println(v8, v1)
    var v15 int32 = 2147483647
    v1 += v3
    for fi16 := 0; fi16 < 1; fi16++ {
        if ((!true) && true) {
            v17 := GenS1{F0: 323}
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        if ((v8 * (v8 >> ((42) & 63))) >= 128) {
            if false {
                v18 -= v18
                var v19 int8 = 127
                var v20 uint = 974
                _ = v19
                _ = v20
            }
            v1 += -124
        }
        _ = fi16
        _ = v18
    }
    var v21 float32 = (3.14159 * 1e10)
    v1 -= v1
    fmt.Println(v21)
    if (!true) {
        switch 2147483647 {
        case 32767:
            switch "a" {
            case "hello":
                var v22 uint8 = (v8 / ((v8) | 1))
                _ = v22
            default:
                var v23 uint32 = 2
                _ = v23
            }
        case -770:
            var v24 float64 = 1.5
            _ = v24
        case -669:
            var v25 uint16 = 1
            _ = v25
        default:
            fmt.Println(v1, v3)
        }
        v26 := GenS1{F0: 1}
        _ = v26
    }
    v27 := GenS1{F0: -1}
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int8 = 127
        v3 = (^(v1))
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v7
    _ = v8
}
