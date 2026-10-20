package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
func genFn0() int8 {
    var v4 string = "a"
    if (float32(1.5) <= 3.14159) {
        var v5 bool = true
        v4 = ("hello" + ((v4 + (v4 + v4)) + ""))
        _ = v5
    } else {
        v6 := GenS1{F0: -740, F1: 42, F2: 0}
        v7 := []int{4294967295}
        v4 = ("\n\t" + v4)
        _ = v6
        _ = v7
    }
    if false {
        v4 = v4
        var v8 int64 = -1
        v4 = "foo bar"
        _ = v8
    } else {
        fmt.Println(v4)
    }
    _ = v4
    return -1
}
func genFn1() uint64 {
    if true {
        if true {
            v9 := []int{32767, 128, 38}
            _ = v9
        } else {
            if ((2 > 2) || true) {
                var v10 uint64 = 1
                _ = v10
            }
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := []int{}
        var v13 uint = 255
        v13 -= v13
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int = -2147483648
    var v15 uint16 = 32767
    v16 := GenS1{F0: int64(v15), F1: 874, F2: 583}
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]string)
    v3 := []string{}
    var v17 uint64 = genFn1()
    if false {
        v3 = append(v3, "")
        if true {
            if false {
                var v18 bool = ((-726 < 4294967295) && false)
                _ = v18
            }
        }
        var v19 uint = 2
        _ = v19
    } else {
        _ = len(v2)
        var v20 uint64 = genFn1()
        _ = v20
    }
    v21 := GenS1{F0: -9223372036854775808, F1: -603, F2: 2147483647}
    v1 = (v1 << ((uint(v1)) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v22 uint8 = v1
    var v23 int = -671
    var v24 int = 260
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
