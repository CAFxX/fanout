package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) uint32 {
    var v5 bool = (104 != uint32(p0))
    var v6 bool = true
    v6 = v5
    v7 := GenS1{F0: -2147483648, F1: 42, F2: 2147483647}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return uint32(p0)
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 int = -2147483648
    v4 := make(map[string]uint32)
    for fi8 := 0; fi8 < 10; fi8++ {
        _ = len(v2)
        _ = fi8
    }
    v9 := GenS1{F0: (v3 >> ((0) & 63)), F1: 0, F2: 916}
    if false {
        v3 -= v3
        var v10 uint64 = 781
        _ = v10
    } else {
        if false {
            var v11 bool = false
            _ = v11
        }
    }
    v12 := GenS2{F0: 255, F1: 1}
    if false {
        fmt.Println(v3)
        v13 := v4[("foo bar" + ("hello" + "\n\t"))]
        var v14 uint = 0
        _ = v13
        _ = v14
    }
    fmt.Println(v3, v1)
    var v15 int = (v3 << ((1) & 63))
    fmt.Println(v3, v15)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 bool = false
        var v18 uint64 = 4294967295
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: (v3 * v3), F1: 1, F2: 420}
    if (32767 < 754) {
        switch 4294967295 {
        case 32767:
            v20 := GenS1{F0: 128, F1: 42, F2: 2}
            _ = v20
        }
    } else {
        delete(v4, "hello")
        v21 := v2["hello"]
        v22 := GenS3{F0: 255}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        if (!true) {
            v24 := []int{}
            var v25 uint64 = 2
            fmt.Println(v3, v25)
            _ = v24
            _ = v25
        } else {
            v15 = -115
            var v26 int32 = -701
            for fi27 := 0; fi27 < 6; fi27++ {
                if (true && false) {
                    v26 = v26
                    v28 := GenS1{F0: -111, F1: 42, F2: genFn0(24)}
                    _ = v28
                } else {
                    var v29 string = "a"
                    var v30 uint64 = 969
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
            _ = v26
        }
        var v31 int64 = -983
        var v32 int16 = (v1 & (^(v1)))
        _ = fi23
        _ = v31
        _ = v32
    }
    v1 = (v1 * (v1 / ((v1) | 1)))
    v33 := v4[""]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v9
}
