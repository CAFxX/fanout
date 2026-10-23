package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint; F1 string; F2 int }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 uint8) uint {
    var v4 int8 = -128
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int = (fi5 + fi5)
        fmt.Println(v4, fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 42
    var v8 float64 = (float64(0.0) / -2.25)
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 65535
    var v3 int32 = -2147483648
    fmt.Println(v3)
    if false {
        v9 := GenS1{F0: 2, F1: 127}
        fmt.Println(v2, v2, v1)
        if true {
            var v10 int8 = 0
            _ = v10
        } else {
            var v11 float32 = 3.14159
            _ = v11
        }
        _ = v9
    }
    switch 32768 {
    case 127:
        v2 = 42
    default:
        if ((!true) && (true || true)) {
            fmt.Println(v2, v3, v3)
        } else {
            var v12 uint64 = 2
            _ = v12
        }
    }
    if ((3.14159 > (1.5 * 3.14159)) && false) {
        fmt.Println(v1)
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 -= 687
                v15 := []int{-617, fi14, fi13}
                _ = fi14
                _ = v15
            }
            var v16 uint16 = 0
            var v17 uint64 = 18446744073709551615
            _ = fi13
            _ = v16
            _ = v17
        }
    }
    v18 := GenS1{F0: 265, F1: 100}
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(fi19, v1, v3)
        var v20 int8 = -1
        v21 := []int{(^(fi19)), fi19}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS1{F0: 2, F1: 69}
        v24 := GenS2{F0: -35, F1: 0}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = 0
    var v26 string = ("hello" + "hello")
    switch 619 {
    case 245:
        v27 := GenS1{F0: 546, F1: 127}
        _ = v27
    case 2147483647:
        v28 := GenS3{F0: 1, F1: (v3 < v3), F2: 2147483648}
        _ = v28
    case 4294967295:
        v29 := GenS4{F0: genFn0(1), F1: v26, F2: v25}
        _ = v29
    default:
        v2 += v2
    }
    v30 := GenS2{F0: int8(v3), F1: (v3 | v3)}
    var v31 bool = ((!((^(v25)) >= (v25 * (v25 << ((65535) & 63))))) && ((v25 | v25) > v25))
    v32 := GenS5{F0: v26, F1: v3, F2: 293}
    switch (v25 + v25) {
    case 1:
        var v33 float32 = -2.25
        _ = v33
    case 351:
        var v34 bool = (926 > 32767)
        _ = v34
    }
    var v35 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
}
