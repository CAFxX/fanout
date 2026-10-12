package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 32767
    var v3 int8 = -128
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    v2 += (v2 >> ((65535) & 63))
    v7 := GenS1{F0: 2}
    if true {
        var v8 int16 = v2
        v9 := GenS2{F0: "a", F1: 615}
        var v10 int32 = 1
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint16 = 635
    var v12 uint64 = 0
    if (185 != 128) {
        v13 := GenS1{F0: 0}
        _ = v13
    }
    fmt.Println(v3, v2)
    v1 -= (^(v1))
    fmt.Println(v12, v12)
    v11 += 42
    v12 -= (^(v12))
    if ((!(float32((-2.25 - ((3.14159 / (float64(1.5) / 1.0)) - 3.14159))) != 0.1)) || ((!(false && false)) || (0 >= 442))) {
        var v14 uint = 1
        v15 := GenS2{F0: "foo bar", F1: (^(v12))}
        for fi16 := 0; fi16 < 19; fi16++ {
            v12 = (v12 << ((v14) & 63))
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    if (true && (2147483647 < 2147483647)) {
        v5[0] = 1
        v1 -= (v1 >> ((42) & 63))
        for fi17 := 0; fi17 < 6; fi17++ {
            switch (fi17 % ((fi17) | 1)) {
            case 228:
                var v18 uint = 255
                _ = v18
            case -9223372036854775808:
                v19 := GenS2{F0: "foo bar", F1: (v12 - (v12 - (v12 * (v12 - 42))))}
                _ = v19
            default:
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 float32 = (float32(0.0) + ((float32(3.14159) * float32(0.1)) - ((-2.25 + 1e10) * (-2.25 + ((1e10 + -2.25) * (((0.0 / 3.14159) + (1.5 + 0.1)) * 0.1))))))
                    var v22 float64 = 3.14159
                    var v23 uint64 = 732
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi17
        }
    } else {
        if true {
            fmt.Println(v2, v11)
            v24 := GenS1{F0: (v12 & v12)}
            _ = v24
        } else {
            fmt.Println(v3, v2, v3)
        }
    }
    fmt.Println(v12)
    var v25 uint32 = 55
    v26 := GenS1{F0: v12}
    if false {
        var v27 int = 4294967295
        v28 := GenS3{F0: int16(v11)}
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
