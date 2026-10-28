package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0(p0 int64, p1 int64) string {
    v7 := []int{128, 42, 32768, -369}
    if (false || false) {
        var v8 int16 = 1000
        v8 -= (v8 * 42)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}
func genFn1() int64 {
    if true {
        var v9 int16 = -32768
        var v10 int8 = 127
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint64 = 18446744073709551615
            _ = fi11
            _ = v12
        }
        var v13 int32 = 2147483647
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int8 = -1
            fmt.Println(fi14, v15, v15)
            for fi16 := 0; fi16 < 5; fi16++ {
                if (-2.25 == (float64(float32((((float64(0.0) - 0.1) - 1.5) / -2.25))) * -2.25)) {
                    v13 = -519
                    switch (("" + "\n\t") + "hello") {
                    case "":
                        var v17 int = 994
                        _ = v17
                    case "foo bar":
                        v15 = -55
                        var v18 uint16 = 42
                        _ = v18
                    default:
                        var v19 uint = 255
                        _ = v19
                    }
                }
                if true {
                    v15 -= (v15 & (v15 ^ (v15 >> ((42) & 63))))
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if true {
        var v20 int16 = 0
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v20 -= v20
                _ = fi22
            }
            _ = fi21
        }
        _ = v20
    }
    var v23 uint8 = 1
    v24 := GenS1{F0: ("hello" + "\n\t"), F1: 459}
    var v25 uint8 = v23
    _ = v23
    _ = v24
    _ = v25
    return 595
}

func main() {
    v1 := []int{4294967295, 127}
    v2 := make(map[string]bool)
    var v3 bool = (((!true) && false) || false)
    v4 := make(map[string]uint32)
    var v5 bool = v3
    var v6 uint8 = 1
    v6 = (v6 >> ((784) & 63))
    var v26 int = 1
    v3 = ((128 != v6) && false)
    v27 := GenS1{F0: "", F1: 924}
    _ = len(v1)
    var v28 int32 = 2147483647
    fmt.Println(v3)
    if (true && ("a" >= "a")) {
        var v29 uint32 = 4294967295
        _ = v29
    }
    if (3.14159 <= -2.25) {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 int16 = -1
            _ = fi30
            _ = v31
        }
    } else {
        v4["\n\t"] = 2147483647
    }
    v3 = (v5 || v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
