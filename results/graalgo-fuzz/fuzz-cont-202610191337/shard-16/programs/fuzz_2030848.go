package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() int64 {
    v7 := []int{32767, 4294967295, -9223372036854775808, 2147483647}
    var v8 int32 = -256
    var v9 int = 32768
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint8) int64 {
    v10 := GenS1{F0: -2147483648, F1: 127, F2: (p0 >> ((144) & 63))}
    v11 := []int{2147483647, 42}
    _ = p0
    _ = v10
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int16 = 671
    var v4 float32 = 0.0
    v5 := []string{"", "hello", "hello", ("hello" + ("" + ""))}
    v6 := make(map[uint32]string)
    var v12 int16 = (v3 + v3)
    var v13 uint = 343
    v14 := GenS2{F0: 255}
    _ = len(v1)
    v4 = v4
    _ = len(v5)
    var v15 float32 = v4
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v13 += (v13 / (((^(v13))) | 1))
            _ = fi17
        }
        if (4294967295 < 370) {
            v13 -= v13
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float64 = 1e10
                _ = fi18
                _ = v19
            }
            v12 = (v12 >> (((v13 ^ 831)) & 63))
        } else {
            var v20 int16 = v12
            v21 := GenS2{F0: 42}
            v22 := GenS1{F0: (fi16 | (fi16 ^ (fi16 | (-(fi16))))), F1: int8(v12), F2: 0}
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    v15 += v15
    switch 128 {
    case 32767:
        fmt.Println(v15, v3)
    case 42:
        if (true || true) {
            var v23 float64 = 3.14159
            _ = v23
        }
    case -2147483648:
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 bool = (0 != 0)
                v5 = append(v5, "\n\t")
                v12 += (-(v12))
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
