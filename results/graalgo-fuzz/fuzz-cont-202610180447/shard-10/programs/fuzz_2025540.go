package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: "foo bar", F1: 487, F2: 127}
        var v7 int64 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := v4[9223372036854775807]
    v8 = v8
    v8 += v8
    if ((-935 > 155) && true) {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch "hello" {
            case "foo bar":
                var v10 uint32 = v8
                _ = v10
            }
            _ = fi9
        }
        if (-9223372036854775808 == 0) {
            var v11 int = 4294967295
            switch "\n\t" {
            case "":
                v12 := GenS2{F0: 151, F1: -74}
                _ = v12
            case "a":
                var v13 int32 = -2147483648
                _ = v13
            case "\n\t":
                v2[v11] = (1 != (v8 ^ (v8 - v8)))
            }
            _ = v11
        } else {
            var v14 int16 = 1000
            v14 = v14
            _ = v14
        }
    } else {
        var v15 float32 = 0.0
        v15 = float32(0.0)
        for fi16 := 0; fi16 < 5; fi16++ {
            v1 = append(v1, v8)
            for fi17 := 0; fi17 < 4; fi17++ {
                v15 = 3.14159
                v15 = 0.1
                if true {
                    _ = len(v4)
                    v15 -= (((v15 * float32(1.5)) + 1.5) * v15)
                    fmt.Println(fi17, fi17)
                } else {
                    v18 := GenS1{F0: "", F1: v8, F2: -128}
                    fmt.Println(v15)
                    _ = v18
                }
                _ = fi17
            }
            delete(v2, -719)
            _ = fi16
        }
        _ = v15
    }
    v19 := GenS3{F0: 885, F1: -128}
    if len(v1) > 0 {
        v1[0] = (v8 % ((4294967295) | 1))
    }
    v8 -= (v8 << ((42) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
