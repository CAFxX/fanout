package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 int64) uint64 {
    v6 := []int{32767, -1, 355}
    p0 = -9223372036854775808
    _ = p0
    _ = v6
    return 806
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    v3 := []string{"foo bar", ("foo bar" + "foo bar")}
    v4 := make(map[string]bool)
    v5 := []int{-9223372036854775808, -876, 127}
    if (!((true && (-668 > 1)) || (0 <= 0))) {
        var v7 int32 = 0
        v8 := v4[("" + "foo bar")]
        _ = v7
        _ = v8
    } else {
        switch "" {
        case "hello":
            var v9 int32 = 1
            _ = v9
        case "a":
            v1[("\n\t" + "\n\t")] = (0.1 != (1e10 - (3.14159 / 1.0)))
        case "\n\t":
            switch 761 {
            case 163:
                var v10 uint = 4294967295
                _ = v10
            case 594:
                var v11 uint = 808
                _ = v11
            }
        default:
            var v12 uint16 = 65535
            _ = v12
        }
        v2 = v2
    }
    var v13 uint = 65535
    delete(v4, "")
    v14 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 32767}
    if (!(65535 == 0)) {
        if (v2 < v2) {
            if ((0.1 + 0.0) >= 1e10) {
                fmt.Println(v13)
                var v15 uint32 = 517
                var v16 uint = (v13 >> ((v13) & 63))
                _ = v15
                _ = v16
            } else {
                v2 = v2
                v17 := GenS2{F0: 1}
                v18 := GenS3{F0: 351, F1: "hello", F2: 42}
                _ = v17
                _ = v18
            }
            var v19 uint32 = uint32(v2)
            var v20 string = ("hello" + ("hello" + "foo bar"))
            _ = v19
            _ = v20
        }
        var v21 int = -9223372036854775808
        var v22 uint32 = 65535
        _ = v21
        _ = v22
    }
    v2 += v2
    if (4294967295 > 0) {
        v13 = v13
        _ = len(v4)
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                v2 += (v2 % ((16) | 1))
                v13 = (v13 + v13)
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi24
            }
            var v25 bool = true
            _ = fi23
            _ = v25
        }
    } else {
        v1["foo bar"] = (-584 != 336)
    }
    if (717 <= -775) {
        if (!(812 >= 1)) {
            fmt.Println(v2, v13)
        }
        v26 := GenS3{F0: 2147483647, F1: "a", F2: -442}
        v2 += (^(v2))
        _ = v26
    } else {
        v2 = (v2 & (^(v2)))
    }
    var v27 uint8 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
