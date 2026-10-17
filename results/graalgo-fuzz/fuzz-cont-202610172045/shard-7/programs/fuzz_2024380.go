package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }
func genFn0() bool {
    var v7 uint32 = 65535
    var v8 int8 = -128
    v9 := []int{408, -1, 139, 4294967295}
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1() int32 {
    var v10 float64 = 3.14159
    v10 += (v10 * float64(float32((v10 - v10))))
    v11 := GenS1{F0: ("hello" + ("\n\t" + "")), F1: 4294967295, F2: 414}
    v12 := GenS2{F0: (true && ((!true) && false)), F1: 63, F2: 42}
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello", "a", "\n\t", "foo bar"}
    v5 := []string{}
    var v6 uint = 677
    fmt.Println(v6, v6)
    var v13 int64 = v2
    v14 := GenS3{F0: uint16(v6)}
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 = (^(v2))
        fmt.Println(fi15, v6)
        var v16 int32 = 356
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: (("foo bar" == "") && (true || false)), F1: -41, F2: 42}
    var v18 float32 = 1e10
    v19 := GenS4{F0: 32767, F1: 42, F2: (1 != 2)}
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 float64 = -2.25
        var v22 uint32 = 107
        var v23 int64 = (v13 / (((v13 - (v2 << (((^(v6))) & 63)))) | 1))
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int16 = 843
    v18 += float32(-2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
