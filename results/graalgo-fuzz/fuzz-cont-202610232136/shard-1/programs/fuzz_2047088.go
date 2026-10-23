package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int }
func genFn0() bool {
    var v7 int8 = 114
    var v8 float32 = (0.0 * ((0.0 + 1.5) - -2.25))
    _ = v7
    _ = v8
    return (true && false)
}
func genFn1() int32 {
    v9 := []int{}
    var v10 float32 = -2.25
    v11 := GenS1{F0: 1, F1: 851}
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{0, 2, 42, 65535}
    var v2 int64 = 0
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]string)
    var v5 int16 = 0
    var v6 float32 = 0.1
    var v12 float32 = float32(((3.14159 + (3.14159 * 1.5)) / 0.1))
    v13 := GenS2{F0: -1}
    v2 = (^(v2))
    fmt.Println(v2)
    v6 = (v6 / ((v12 - v12) + float32((float64(v12) - -2.25))))
    if (v2 > (v2 * (v2 << ((655) & 63)))) {
        fmt.Println(v6, v5, v5)
        var v14 int32 = 1
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(v3, v5, v3)
        var v16 int64 = v2
        _ = fi15
        _ = v16
    }
    if (!("" <= "a")) {
        v12 = v12
        fmt.Println(v3)
        _ = len(v1)
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v5 = (v5 / (((^(v5))) | 1))
        if true {
            for fi18 := 0; fi18 < 6; fi18++ {
                v3 = 2147483647
                v12 = 3.14159
                _ = fi18
            }
        } else {
            var v19 uint64 = 0
            v20 := GenS1{F0: -9223372036854775808, F1: 2}
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v6 += (-2.25 + v12)
            var v22 float32 = -2.25
            fmt.Println(fi21, v6, v2)
            _ = fi21
            _ = v22
        }
        _ = fi17
    }
    v6 -= v12
    v23 := GenS1{F0: 0, F1: 2}
    v6 += 1.5
    var v24 uint32 = v3
    _ = len(v1)
    if true {
        v25 := GenS3{F0: (v3 >> ((343) & 63)), F1: -78, F2: 2147483647}
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            v2 += (v2 ^ (^(v2)))
            var v28 int64 = 42
            v5 = v5
            _ = len(v4)
            _ = fi27
            _ = v28
        }
        if false {
            fmt.Println(v12, v2, fi26)
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 int = int(v24)
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    var v31 string = (("hello" + (("\n\t" + ("a" + "a")) + "\n\t")) + "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
