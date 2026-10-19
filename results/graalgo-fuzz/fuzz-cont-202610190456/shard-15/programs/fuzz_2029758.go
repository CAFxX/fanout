package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint8 }
func genFn0(p0 uint16) int8 {
    v7 := []int{-949, -277, 0, 32768}
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 15; fi9++ {
        if (false && true) {
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    fmt.Println(p0)
                    _ = fi10
                }
                var v11 uint8 = 0
                _ = v11
            }
        }
        var v12 uint8 = 128
        v7 = append(v7, fi9)
        _ = fi9
        _ = v12
    }
    fmt.Println(p0, v8)
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 42
    var v2 int16 = 1
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    var v5 uint32 = 2
    var v6 uint16 = 32767
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 532}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: -9223372036854775808}
    fmt.Println(v1)
    v16 := GenS3{F0: 32767, F1: genFn0(v6)}
    if (9223372036854775807 > 0) {
        v17 := v3[uint32(v1)]
        for fi18 := 0; fi18 < 12; fi18++ {
            if ((-1 < -197) && (127 == 0)) {
                var v19 string = "hello"
                _ = v19
            } else {
                var v20 uint64 = 263
                v21 := GenS4{F0: (v20 % ((v20) | 1)), F1: 0, F2: 42}
                var v22 int16 = v2
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi18
        }
        _ = v17
    } else {
        var v23 bool = (!true)
        var v24 int32 = -1
        v5 = v5
        _ = v23
        _ = v24
    }
    var v25 int8 = 100
    _ = len(v3)
    var v26 int16 = (v2 ^ v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
