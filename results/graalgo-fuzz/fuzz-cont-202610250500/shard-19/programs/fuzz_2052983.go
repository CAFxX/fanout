package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 bool, p1 int64) int16 {
    p1 += (p1 ^ (p1 / (((p1 * p1)) | 1)))
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 242}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(p1, p1, p0)
        switch (^(fi9)) {
        case 2147483647:
            if ("foo bar" < "foo bar") {
                var v10 float64 = float64(1e10)
                fmt.Println(v10, v10, p0)
                p1 += (p1 + p1)
                _ = v10
            } else {
                p1 += -1
                var v11 int8 = int8(p1)
                v12 := GenS1{F0: 0}
                _ = v11
                _ = v12
            }
        }
        p0 = p0
        _ = fi9
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 bool = false
    v2 := []string{}
    v3 := []uint32{2147483648}
    var v4 uint = 4294967295
    v5 := make(map[uint32]bool)
    var v6 int32 = 706
    var v13 uint64 = 1
    v14 := GenS1{F0: 1}
    v6 += -2147483648
    v1 = ((float32(float64(0.0)) - 0.0) == 0.0)
    v15 := GenS1{F0: 42}
    var v16 int64 = 42
    switch -9223372036854775808 {
    case 437:
        fmt.Println(v16)
    case 128:
        var v17 float32 = -2.25
        _ = v17
    default:
        if (("\n\t" > ("a" + "a")) && true) {
            v6 = v6
            _ = len(v5)
        }
    }
    switch int(v4) {
    case 2147483647:
        fmt.Println(v13, v1)
    }
    v18 := GenS1{F0: 255}
    v4 += 456
    v19 := GenS1{F0: 127}
    v1 = (!(!v1))
    if (false || (!(32767 <= 1))) {
        v20 := v5[666]
        _ = len(v2)
        v21 := GenS1{F0: 127}
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := v5[0]
            delete(v5, 0)
            _ = fi22
            _ = v23
        }
        v24 := GenS1{F0: 255}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
