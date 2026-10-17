package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 bool; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 int; F2 uint }
func genFn0() int {
    var v4 int32 = 1
    v5 := GenS1{F0: 42, F1: 100, F2: (("" + "a") + "foo bar")}
    var v6 float64 = ((float64(-2.25) - 1e10) * 0.1)
    _ = v4
    _ = v5
    _ = v6
    return -715
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "hello"
    var v3 uint16 = 32767
    v7 := GenS1{F0: 18446744073709551615, F1: -105, F2: ("" + "a")}
    delete(v1, 4294967295)
    v8 := []int{32767, 4294967295, 414, 1}
    v2 = v2
    if true {
        if (-764 == 1) {
            v9 := GenS2{F0: (1000 == -851), F1: -1, F2: 1}
            var v10 int32 = 1
            _ = v10
            _ = v9
        } else {
            var v11 int64 = -663
            var v12 uint32 = 1
            v13 := GenS2{F0: true, F1: 1, F2: uint64(v12)}
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(v3)
    } else {
        fmt.Println(v3)
        if false {
            v2 = ("" + (((v2 + "a") + v2) + "foo bar"))
            var v14 int64 = 698
            _ = v14
        }
        var v15 int8 = 0
        _ = v15
    }
    if len(v8) > 0 {
        v8[0] = 128
    }
    var v16 float64 = -2.25
    var v17 int32 = 2147483647
    if (0 < 709) {
        var v18 bool = false
        delete(v1, 42)
        _ = v18
    } else {
        v19 := GenS2{F0: (127 == -128), F1: 32767, F2: 42}
        v20 := GenS3{F0: 836, F1: -9223372036854775808, F2: 4294967295}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            delete(v1, 2)
            _ = fi22
        }
        v3 += v3
        var v23 int64 = 42
        _ = fi21
        _ = v23
    }
    if ((v17 & v17) >= (v17 | v17)) {
        var v24 bool = (0 != 0)
        var v25 bool = true
        v3 -= 470
        _ = v24
        _ = v25
    }
    v17 += 660
    if true {
        v3 -= (v3 | (v3 - 1))
        var v26 int16 = 760
        var v27 float32 = ((1e10 + (1.5 - 1.5)) * (1e10 + 0.1))
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v16, v3)
            v1[0] = fi28
            var v29 uint64 = 0
            _ = fi28
            _ = v29
        }
    }
    var v30 int8 = 42
    v31 := GenS2{F0: true, F1: 42, F2: 0}
    v32 := GenS1{F0: 0, F1: (v30 << ((2) & 63)), F2: ("hello" + v2)}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        v16 = float64(-2.25)
        fmt.Println(v3)
        _ = fi33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v7
    _ = v8
}
