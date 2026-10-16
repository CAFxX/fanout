package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int32; F1 int; F2 uint32 }
func genFn0(p0 uint64, p1 string, p2 uint64) uint32 {
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 888
}

func main() {
    var v1 uint64 = 570
    var v2 uint32 = 945
    var v3 uint16 = 666
    v4 := make(map[int]int)
    v5 := GenS1{F0: (v2 << ((2) & 63)), F1: -754}
    v6 := GenS1{F0: (v2 | (v2 - 2)), F1: 2147483647}
    v2 = (v2 / (((v2 ^ v2)) | 1))
    var v7 int16 = -636
    for fi8 := 0; fi8 < 9; fi8++ {
        if (("\n\t" + (("hello" + "\n\t") + "a")) >= "hello") {
            var v9 uint32 = v2
            var v10 uint = uint(v1)
            _ = v10
            _ = v9
        } else {
            v1 -= (v1 | (^(v1)))
            v11 := v4[127]
            v1 = v1
            _ = v11
        }
        _ = fi8
    }
    if true {
        v1 = v1
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS2{F0: 1, F1: (42 != (v1 - (v1 | 42)))}
            _ = fi12
            _ = v13
        }
        v14 := GenS3{F0: 42, F1: 32767, F2: genFn0(827, "foo bar", 4294967295)}
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(v7, fi15, fi15)
                v7 -= (v7 ^ -32768)
                var v17 uint32 = v2
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        _ = v14
    }
    var v18 int64 = 9223372036854775807
    v3 = v3
    var v19 uint32 = v2
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v19, fi20, v1)
        v21 := []int{}
        if len(v21) > 0 {
            v21[0] = -2147483648
        }
        _ = fi20
        _ = v21
    }
    var v22 float64 = 0.0
    var v23 uint64 = (^(v1))
    var v24 float64 = 1e10
    var v25 string = "foo bar"
    fmt.Println(v25, v3)
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS1{F0: v2, F1: int32(v7)}
        switch "foo bar" {
        case "":
            for fi28 := 0; fi28 < 4; fi28++ {
                switch (fi26 | (fi28 * -9223372036854775808)) {
                case -145:
                    v24 += v24
                }
                _ = fi28
            }
        case "\n\t":
            fmt.Println(v23, v19, v25)
        default:
            var v29 uint8 = 127
            _ = v29
        }
        v30 := []int{}
        _ = fi26
        _ = v27
        _ = v30
    }
    v2 -= 207
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
