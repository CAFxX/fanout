package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 int) uint32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint8 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 2}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint32 = 4294967295
        v11 = 2147483648
        v12 := GenS2{F0: (fi10 | (fi10 & (fi10 ^ fi10))), F1: 0, F2: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    var v3 float64 = float64(3.14159)
    var v4 float64 = (v3 + v3)
    var v5 uint32 = 4294967295
    var v6 float64 = (v3 - -2.25)
    var v13 uint8 = 2
    var v14 int32 = v2
    v15 := GenS2{F0: 42, F1: (v2 + (-(v2))), F2: -9223372036854775808}
    v16 := []int{}
    v1[32768] = (750 != 456)
    v17 := GenS3{F0: v5}
    var v18 int = 96
    switch v18 {
    case 128:
        if (((float64(3.14159) <= v6) && false) && true) {
            var v19 uint = 907
            _ = v19
        }
    default:
        var v20 uint16 = 1
        _ = v20
    }
    v5 -= (v5 << ((2) & 63))
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS4{F0: v5, F1: 0}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 65535, F1: (v5 + (v5 + genFn0(v18)))}
    switch (^(v18)) {
    case 429:
        var v24 uint = 2
        _ = v24
    case 9223372036854775807:
        fmt.Println(v5, v4)
    default:
        v25 := GenS2{F0: v18, F1: v2, F2: -9223372036854775808}
        _ = v25
    }
    delete(v1, -1)
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS2{F0: 761, F1: v2, F2: -1}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v13 -= v13
        var v29 int32 = v14
        _ = fi28
        _ = v29
    }
    v16 = append(v16, (v18 << ((51) & 63)))
    v16 = append(v16, 32768)
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint32 = v5
        if (!(false || true)) {
            for fi32 := 0; fi32 < 10; fi32++ {
                v4 -= (-2.25 + 3.14159)
                _ = fi32
            }
            if true {
                var v33 float32 = 3.14159
                _ = v33
            } else {
                fmt.Println(v18, v3)
            }
            if false {
                for fi34 := 0; fi34 < 20; fi34++ {
                    v35 := GenS5{F0: fi30}
                    var v36 int32 = v14
                    delete(v1, -2147483648)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        } else {
            var v37 uint64 = 42
            for fi38 := 0; fi38 < 5; fi38++ {
                v39 := GenS4{F0: uint32(v13), F1: 993}
                _ = fi38
                _ = v39
            }
            var v40 float32 = float32(-2.25)
            _ = v37
            _ = v40
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
