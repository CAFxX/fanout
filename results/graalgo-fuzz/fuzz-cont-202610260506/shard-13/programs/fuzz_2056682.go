package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 int32 }
func genFn0() uint64 {
    var v6 int8 = 114
    v7 := []int{128}
    for fi8 := 0; fi8 < 9; fi8++ {
        v6 -= (v6 ^ (v6 | -1))
        _ = fi8
    }
    _ = len(v7)
    _ = v6
    _ = v7
    return 379
}
func genFn1(p0 string) uint64 {
    fmt.Println(p0)
    v9 := GenS1{F0: 65535}
    _ = p0
    _ = v9
    return 351
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{4294967295, 341, 777, 2147483648}
    var v3 uint = 4294967295
    var v4 uint64 = 1
    v5 := []string{("\n\t" + ((("a" + "a") + "\n\t") + "hello")), "a", "a", ("hello" + "foo bar")}
    v10 := GenS2{F0: 42, F1: false, F2: 255}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 := GenS3{F0: 1, F1: 610, F2: 2147483647}
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v3, fi12)
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS2{F0: -848, F1: (false || (533 == 428)), F2: 0}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint8 = uint8(v3)
            var v17 int = v1
            v16 += v16
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi12
    }
    v4 = genFn1("\n\t")
    if false {
        for fi18 := 0; fi18 < 8; fi18++ {
            v2 = append(v2, 0)
            var v19 int16 = -1
            _ = fi18
            _ = v19
        }
        _ = len(v5)
        var v20 int16 = 32767
        _ = v20
    } else {
        if (!false) {
            fmt.Println(v1, v3, v3)
            v21 := GenS2{F0: -1, F1: true, F2: 255}
            var v22 int = v1
            _ = v21
            _ = v22
        }
        if true {
            v1 += 72
            var v23 uint32 = 509
            v1 += 32768
            _ = v23
        }
    }
    v3 -= v3
    v24 := GenS1{F0: 2}
    var v25 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
