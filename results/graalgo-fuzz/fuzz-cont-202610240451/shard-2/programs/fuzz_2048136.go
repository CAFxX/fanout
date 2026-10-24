package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    var v2 bool = true
    var v3 uint8 = 212
    var v4 uint8 = v3
    v1 -= v1
    v5 := GenS1{F0: 2147483647, F1: 242, F2: 127}
    var v6 uint8 = 156
    for fi7 := 0; fi7 < 15; fi7++ {
        if (((0.1 - 0.1) + (-2.25 * (1.5 - 3.14159))) >= ((float64(1e10) * 0.0) + 1e10)) {
            if v2 {
                fmt.Println(fi7, v1)
            }
        }
        if true {
            v8 := GenS2{F0: true}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 10; fi9++ {
                if (false && true) {
                    var v10 uint = 1
                    fmt.Println(v4)
                    v11 := []int{}
                    _ = v10
                    _ = v11
                } else {
                    v6 = 169
                }
                var v12 float32 = (1.5 * 3.14159)
                v12 = (v12 / 0.1)
                _ = fi9
                _ = v12
            }
        }
        _ = fi7
    }
    v13 := GenS1{F0: 4294967295, F1: -1, F2: -9223372036854775808}
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{fi14, 813, (fi14 >> ((0) & 63))}
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS3{F0: 9223372036854775807}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v18 := GenS3{F0: 372}
    if false {
        var v19 float32 = ((float32((((-2.25 * (-2.25 + float64(((3.14159 / ((1.5 - 1.5) - 1e10)) * 0.0)))) + ((0.1 / (float64(1e10) - 0.0)) * 0.0)) / 1.5)) - (3.14159 + 0.1)) - 0.1)
        if v2 {
            var v20 int16 = 42
            var v21 bool = v2
            _ = v20
            _ = v21
        }
        _ = v19
    }
    v4 = 128
    fmt.Println(v3, v4)
    var v22 int16 = 32767
    if ((v4 ^ v6) != v6) {
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v22 = (v22 << ((620) & 63))
                v6 = 2
                v25 := GenS3{F0: 993}
                _ = fi24
                _ = v25
            }
            if true {
                v6 = v6
            }
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
