package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 int32, p1 int16, p2 uint32) uint32 {
    v5 := GenS1{F0: -1, F1: "\n\t"}
    if (((-1 == p0) || false) && (true && (4294967295 != 642))) {
        fmt.Println(p2, p2)
        if (false && (true || (float32(float64((((float32(3.14159) * (0.1 + ((1e10 + 1e10) / 1.0))) + 0.1) - 1e10))) != 3.14159))) {
            switch "foo bar" {
            case "\n\t":
                var v6 int16 = (p1 * -844)
                _ = v6
            }
            var v7 uint = 255
            var v8 uint16 = 65535
            _ = v7
            _ = v8
        }
        p0 = (p0 % ((p0) | 1))
    } else {
        var v9 uint = 784
        _ = v9
    }
    v10 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return p2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{65535, 1}
    v3 := make(map[string]uint32)
    var v4 int64 = 445
    v4 = (v4 & v4)
    fmt.Println(v4)
    if (!false) {
        v4 += (v4 >> ((255) & 63))
        v4 = (v4 % ((9223372036854775807) | 1))
        v11 := GenS2{F0: -608, F1: 18446744073709551615, F2: 397}
        _ = v11
    } else {
        switch "\n\t" {
        case "a":
            var v12 int8 = -112
            _ = v12
        case "foo bar":
            var v13 string = ("a" + "")
            _ = v13
        }
    }
    fmt.Println(v4)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int8 = -1
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 344
    fmt.Println(v4, v16)
    var v17 bool = true
    for fi18 := 0; fi18 < 10; fi18++ {
        _ = len(v1)
        _ = fi18
    }
    v19 := GenS1{F0: -2147483648, F1: "hello"}
    v2 = append(v2, 328)
    v16 += (v16 << ((1) & 63))
    v4 += v4
    var v20 string = "a"
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if len(v2) > 0 {
                v2[0] = 0
            }
            v17 = true
            if (!true) {
                var v23 uint64 = v16
                v17 = false
                _ = v23
            }
            var v24 int64 = 0
            _ = fi22
            _ = v24
        }
        _ = fi21
    }
    var v25 bool = (0.1 <= (((3.14159 / -2.25) * 3.14159) - 1e10))
    v26 := GenS2{F0: -543, F1: (v16 * v16), F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
