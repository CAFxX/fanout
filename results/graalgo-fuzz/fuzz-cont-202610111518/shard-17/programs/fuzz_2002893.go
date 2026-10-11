package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 int32, p1 int64) int32 {
    var v7 float32 = -2.25
    v8 := []int{int(p1), 760, 32768}
    if (2 == 219) {
        var v9 int64 = p1
        _ = v9
    } else {
        var v10 uint32 = 846
        v10 = 840
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -774
}

func main() {
    var v1 float32 = (float32(0.0) / float32(0.1))
    var v2 int32 = -2147483648
    var v3 uint8 = 159
    var v4 float32 = (3.14159 - float32((1.5 + 1e10)))
    v5 := make(map[uint32]uint32)
    var v6 uint = 65535
    v11 := []int{32767, 127, 791, 434}
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := v5[4294967295]
        _ = len(v5)
        if true {
            v4 = (float32(((((1.5 - 1.5) - 1.5) + ((((3.14159 - (-2.25 / 1.0)) + 1e10) / 3.14159) + (0.0 + (3.14159 - 1e10)))) / 1e10)) - (v1 * v4))
        }
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int8 = 1
        v6 = v6
        _ = fi14
        _ = v15
    }
    v6 += v6
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v1, fi16)
        v17 := GenS1{F0: ("hello" + ("foo bar" + (("" + "\n\t") + "\n\t"))), F1: 1, F2: 466}
        _ = fi16
        _ = v17
    }
    if (2 < uint32(v6)) {
        v3 += v3
        v4 = (float32(0.0) / (((v4 + v1) * v4) * (float32(((0.0 - 1e10) / 1.5)) / v1)))
        var v18 bool = false
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint32 = 1
        var v21 bool = false
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint64 = uint64(v6)
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
