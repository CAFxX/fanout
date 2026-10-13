package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
func genFn0() uint8 {
    var v4 bool = true
    for fi5 := 0; fi5 < 9; fi5++ {
        if v4 {
            var v6 uint64 = 2
            var v7 uint16 = 427
            _ = v6
            _ = v7
        }
        var v8 int64 = -9223372036854775808
        _ = fi5
        _ = v8
    }
    _ = v4
    return 128
}
func genFn1(p0 int16, p1 uint8, p2 string) uint8 {
    v9 := []int{0, 32768, 128, 32768}
    switch -784 {
    case 32767:
        if false {
            v10 := GenS1{F0: 128, F1: 109, F2: -1}
            _ = v10
        }
    case 128:
        v11 := GenS2{F0: 0}
        _ = v11
    case 127:
        p0 = -32768
    }
    v12 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return p1
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]string)
    var v3 int64 = -881
    var v13 int = 900
    v3 = (-(v3))
    switch "foo bar" {
    case "hello":
        v14 := GenS3{F0: 631}
        _ = v14
    default:
        v15 := GenS1{F0: 113, F1: 1, F2: (v3 + v3)}
        _ = v15
    }
    var v16 int64 = 42
    var v17 string = "a"
    v18 := v2[v13]
    fmt.Println(v13)
    var v19 uint32 = 2
    if false {
        if false {
            fmt.Println(v16, v16, v3)
        } else {
            v3 = -1
        }
    }
    v20 := v2[(v13 % (((v13 | (v13 ^ 128))) | 1))]
    if true {
        fmt.Println(v3)
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := []int{fi21, 42}
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS2{F0: -2147483648}
                var v25 uint64 = 2
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    } else {
        if true {
            fmt.Println(v13, v13, v17)
            v16 = int64(v19)
            var v26 uint16 = 32767
            _ = v26
        } else {
            v27 := GenS3{F0: 0}
            for fi28 := 0; fi28 < 16; fi28++ {
                v2[(v13 >> ((uint(v19)) & 63))] = ("" + "")
                var v29 int32 = 269
                if (int32(fi28) == v29) {
                    v16 -= 9223372036854775807
                    v13 = v13
                }
                _ = fi28
                _ = v29
            }
            v20 = (v20 + v20)
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
