package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int; F2 int16 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
func genFn0() uint64 {
    if ((!(1e10 <= 0.0)) && false) {
        for fi4 := 0; fi4 < 2; fi4++ {
            var v5 uint = 660
            var v6 int64 = -677
            var v7 uint = v5
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
        var v8 uint64 = 2
        v9 := GenS1{F0: (-956 == 0)}
        _ = v8
        _ = v9
    } else {
        var v10 float64 = float64(0.0)
        fmt.Println(v10)
        _ = v10
    }
    if (909 != 4294967295) {
        var v11 float32 = (0.0 - 1.5)
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{fi12, fi12, 127}
            var v14 uint32 = 4294967295
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        v15 := GenS2{F0: -1, F1: 32767, F2: 507}
        var v16 int = 1
        fmt.Println(v16)
        _ = v15
        _ = v16
    }
    if false {
        v17 := GenS3{F0: true, F1: ("foo bar" + "\n\t"), F2: 127}
        _ = v17
    }
    v18 := []int{0, -9223372036854775808, 2}
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 721
    var v3 int = 0
    var v19 int8 = 100
    v20 := GenS3{F0: true, F1: (("foo bar" + "hello") + "hello"), F2: v19}
    var v21 uint64 = 18446744073709551615
    v22 := GenS1{F0: true}
    switch -9223372036854775808 {
    case -9223372036854775808:
        v2 -= (v2 & v2)
    }
    fmt.Println(v21, v2)
    var v23 uint = 470
    v21 -= 4294967295
    v19 -= (v19 ^ v19)
    switch int(v21) {
    case 128:
        var v24 string = "hello"
        _ = v24
    }
    v25 := GenS2{F0: 0, F1: (v3 >> ((805) & 63)), F2: -348}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
