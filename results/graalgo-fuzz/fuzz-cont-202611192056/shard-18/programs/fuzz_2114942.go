package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 int64, p1 int) int64 {
    v5 := []int{p1, p1, 127}
    var v6 uint32 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 4294967295
    var v4 int32 = 2147483647
    v1[2147483647] = (!(!false))
    v2[241] = true
    if (!false) {
        var v7 uint32 = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint8 = 2
        delete(v2, 2)
        _ = fi8
        _ = v9
    }
    var v10 int16 = 42
    var v11 float32 = ((3.14159 - 3.14159) / 0.1)
    for fi12 := 0; fi12 < 2; fi12++ {
        if true {
            v1[0] = (42 != 2)
        } else {
            var v13 int32 = -2147483648
            if false {
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int = (-(fi14))
                    _ = fi14
                    _ = v15
                }
                if false {
                    var v16 int16 = (v10 - v10)
                    v11 -= float32(0.0)
                    v13 = (v13 - 1)
                    _ = v16
                } else {
                    var v17 uint64 = v3
                    v13 -= v13
                    v17 = 5
                    _ = v17
                }
                for fi18 := 0; fi18 < 13; fi18++ {
                    v13 += (v13 % ((42) | 1))
                    v3 = (v3 % (((^(v3))) | 1))
                    v19 := v2[4294967295]
                    var v20 float32 = (v11 / ((v11 / v11) / (float32(0.0) * (((v11 + (-2.25 + ((v11 + v11) * (v11 - 0.1)))) * 0.0) / (1e10 * ((v11 + (-2.25 - ((v11 + 3.14159) / 1.0))) / (float32(0.0) - (0.1 / (((v11 + (v11 - (v11 + 1e10))) - v11) + v11)))))))))
                    var v21 int64 = -9223372036854775808
                    delete(v1, 65535)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            v4 += 2147483647
            _ = v13
        }
        fmt.Println(v4, v10)
        var v22 int64 = 0
        _ = fi12
        _ = v22
    }
    var v23 uint = 904
    var v24 int64 = -889
    v25 := []int{int(v23), 32767, -457}
    v10 = -1
    v26 := GenS1{F0: genFn0(v24, -870), F1: (v10 & v10)}
    _ = len(v2)
    v11 -= 0.0
    _ = len(v25)
    fmt.Println(v24, v23)
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int32 = v4
        v24 += (v24 >> ((v23) & 63))
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
