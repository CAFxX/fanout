package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 uint, p1 string) string {
    p0 += 1
    p1 = p1
    var v6 bool = false
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = v6
    return (p1 + (p1 + p1))
}
func genFn1() uint64 {
    v7 := GenS1{F0: 42}
    v8 := GenS1{F0: 65535}
    var v9 uint = 18446744073709551615
    v10 := []int{9223372036854775807, -9223372036854775808, 128}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    v3 := []int{2147483647}
    var v4 uint16 = 65535
    var v5 float32 = 3.14159
    v4 += 255
    for fi11 := 0; fi11 < 2; fi11++ {
        if true {
            v12 := v2["hello"]
            _ = v12
        }
        _ = fi11
    }
    switch "a" {
    case "a":
        v13 := GenS2{F0: 536, F1: 42}
        _ = v13
    default:
        var v14 int64 = -1
        _ = v14
    }
    v5 = v5
    var v15 uint32 = 0
    if false {
        for fi16 := 0; fi16 < 9; fi16++ {
            v2["a"] = ((^(fi16)) != -64)
            var v17 uint8 = 42
            if len(v3) > 0 {
                v3[0] = fi16
            }
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int8 = -6
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint = 0
                v5 = v5
                _ = len(v2)
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    _ = len(v2)
    if ("hello" > "\n\t") {
        if false {
            v22 := GenS1{F0: 42}
            if (false || true) {
                fmt.Println(v4)
                v23 := GenS3{F0: -844, F1: 0, F2: -284}
                v24 := GenS1{F0: 663}
                _ = v23
                _ = v24
            } else {
                var v25 int64 = -1
                _ = v25
            }
            _ = v22
        } else {
            var v26 bool = false
            _ = v26
        }
        var v27 int16 = 593
        _ = v27
    }
    v28 := GenS3{F0: 0, F1: 4294967295, F2: int32(v15)}
    delete(v2, (("foo bar" + "a") + ""))
    v5 -= v5
    v15 += v15
    v4 = (v4 & (v4 << ((255) & 63)))
    var v29 uint32 = (^(v15))
    var v30 int32 = 0
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 float64 = (0.0 * 0.1)
        v15 -= uint32(v1)
        if len(v3) > 0 {
            v3[0] = int(v4)
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
