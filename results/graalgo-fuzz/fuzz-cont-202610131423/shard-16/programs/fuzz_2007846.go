package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint64, p1 uint16) uint8 {
    var v7 uint8 = 148
    v8 := GenS1{F0: ("foo bar" + "a")}
    fmt.Println(v7)
    v9 := []int{4294967295, 1}
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := []uint32{2147483647, 1, 65535, 65535}
    var v2 uint = 299
    var v3 float32 = 0.1
    var v4 uint16 = 42
    var v5 int = -9223372036854775808
    var v6 int16 = 0
    var v10 uint32 = 0
    v2 -= (v2 * 4294967295)
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v10 = 1
            v3 += v3
            v13 := GenS1{F0: ""}
            _ = fi12
            _ = v13
        }
        if false {
            v2 = v2
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(v6, v3)
                var v15 int16 = (v6 + 32767)
                if (!((^(v5)) > 0)) {
                    var v16 uint = v2
                    v17 := GenS2{F0: v4, F1: -162, F2: 473}
                    var v18 uint16 = v4
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v1 = append(v1, uint32(v2))
                    fmt.Println(v5, v3)
                }
                _ = fi14
                _ = v15
            }
            v3 = ((v3 * (3.14159 - (v3 + (v3 - v3)))) - 0.0)
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            v2 = 4294967295
            v4 = v4
            _ = fi19
        }
        _ = fi11
    }
    var v20 int = (-(v5))
    var v21 int8 = 69
    v10 += v10
    if true {
        fmt.Println(v6, v3, v21)
        v1 = append(v1, (v10 ^ 0))
        var v22 int8 = v21
        _ = v22
    } else {
        var v23 int8 = int8(v4)
        fmt.Println(v3)
        _ = v23
    }
    var v24 uint8 = 127
    v25 := GenS3{F0: 42, F1: uint64(v2)}
    if false {
        fmt.Println(v21, v21, v4)
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if (false && ("a" >= "a")) {
            v5 += (v5 + fi26)
        } else {
            var v27 uint32 = 787
            v20 += v20
            _ = v27
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi26
    }
    v10 = (v10 ^ (v10 | (v10 ^ (v10 & (v10 * (v10 + 2))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
