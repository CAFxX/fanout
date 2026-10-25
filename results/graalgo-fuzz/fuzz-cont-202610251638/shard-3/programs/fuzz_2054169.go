package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0() uint {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int64 = 0
        _ = fi4
        _ = v5
    }
    var v6 bool = (42 >= 0)
    _ = v6
    return 2
}
func genFn1() uint16 {
    var v7 int = -365
    if (-1 != 845) {
        v8 := []int{2147483647, 833}
        if true {
            var v9 uint = 783
            var v10 int = v7
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 36, F1: 32767, F2: 65535}
        _ = v11
        _ = v8
    }
    _ = v7
    return 32767
}

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]string)
    var v3 bool = false
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 -= v1
        _ = fi12
    }
    v13 := v2[959]
    var v14 int64 = 42
    v14 = -9223372036854775808
    if v3 {
        var v15 uint16 = 655
        v16 := GenS2{F0: 42, F1: 243}
        _ = v15
        _ = v16
    } else {
        var v17 uint16 = 255
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(fi18, v1, v3)
        v14 += v14
        _ = fi18
    }
    var v19 uint32 = 0
    v3 = v3
    fmt.Println(v19, v3, v13)
    v13 = "hello"
    fmt.Println(v1, v19)
    if v3 {
        var v20 int64 = 985
        var v21 int16 = -361
        _ = v20
        _ = v21
    } else {
        v2[v19] = v13
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint = (^(v1))
            _ = fi22
            _ = v23
        }
        fmt.Println(v1, v19, v13)
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v1 = (v1 | genFn0())
        _ = fi24
    }
    var v25 uint16 = uint16(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
