package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 uint, p1 uint64) uint32 {
    if ("hello" >= ("foo bar" + "")) {
        var v5 int8 = -47
        _ = v5
    } else {
        var v6 int32 = -777
        v7 := GenS1{F0: 42, F1: 0, F2: 127}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            var v9 int8 = 76
            switch (("\n\t" + "a") + "\n\t") {
            case "hello":
                v10 := GenS2{F0: 26}
                _ = v10
            case "":
                var v11 int = fi8
                _ = v11
            case "\n\t":
                fmt.Println(p0, v9)
            default:
                switch int(v9) {
                case -98:
                    v12 := GenS3{F0: -9223372036854775808, F1: (^(v9))}
                    _ = v12
                default:
                    var v13 uint = p0
                    _ = v13
                }
            }
            _ = v9
        } else {
            v14 := []int{fi8, fi8, (^(fi8)), fi8}
            _ = v14
        }
        v15 := GenS2{F0: -86}
        _ = fi8
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        p0 -= p0
        var v17 uint8 = 149
        v17 = 128
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -9223372036854775808
    v3 := []string{}
    v4 := make(map[int]uint32)
    var v18 uint = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    if (18446744073709551615 != (v18 ^ v18)) {
        var v19 int32 = -2147483648
        _ = v19
    }
    var v20 uint64 = 0
    var v21 int = (-(v2))
    var v22 uint32 = 65535
    v23 := GenS3{F0: 1, F1: 63}
    if true {
        var v24 uint = v18
        v24 += 65535
        v20 = v20
        _ = v24
    } else {
        fmt.Println(v18, v2)
        v1["\n\t"] = (v21 * -9223372036854775808)
    }
    v2 -= 444
    v2 += (v21 ^ v21)
    var v25 uint8 = 199
    v3 = append(v3, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
