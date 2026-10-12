package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16) uint16 {
    v7 := GenS1{F0: 2147483648}
    p0 += 255
    if ((p0 + (p0 << ((226) & 63))) < (p0 >> ((255) & 63))) {
        var v8 float64 = 1e10
        if (v8 != 0.0) {
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(p0, fi9, p0)
                _ = fi9
            }
        } else {
            v8 += v8
            v10 := GenS1{F0: 2147483648}
            var v11 int32 = -2147483648
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint64 = 4294967295
            v14 := GenS1{F0: 1}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v8
    } else {
        switch -9223372036854775808 {
        case -634:
            var v15 int = -1
            _ = v15
        }
        var v16 float64 = -2.25
        p0 += p0
        _ = v16
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 bool = true
    v4 := []string{("foo bar" + "\n\t"), "\n\t", "a"}
    var v5 bool = false
    var v6 int32 = -1
    var v17 int8 = 42
    var v18 float32 = 1e10
    var v19 bool = (1 == 0)
    var v20 uint = 255
    if (!v3) {
        var v21 uint16 = 42
        fmt.Println(v21)
        v19 = ((32767 < -9223372036854775808) && (0 != 2))
        _ = v21
    }
    var v22 float32 = v18
    var v23 int32 = v6
    fmt.Println(v3, v20, v20)
    var v24 int16 = 787
    v24 -= (v24 + int16(v23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
