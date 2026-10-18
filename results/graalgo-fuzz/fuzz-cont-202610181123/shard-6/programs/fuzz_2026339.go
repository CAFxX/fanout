package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
func genFn0() int64 {
    v4 := GenS1{F0: 32767, F1: 42, F2: 18446744073709551615}
    var v5 string = "hello"
    v6 := []int{-975}
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            v5 = v5
        }
        v9 := GenS1{F0: fi8, F1: 523, F2: 18446744073709551615}
        _ = fi8
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -484
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 0
    var v3 string = (("a" + "") + "\n\t")
    v10 := []int{-742, 0, 4294967295}
    v3 = ("\n\t" + v3)
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = (^(v1))
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 772
    if true {
        var v14 int32 = -1
        if true {
            v2 += v2
            var v15 uint8 = v2
            _ = v15
        }
        _ = v14
    } else {
        v2 = (v2 + v2)
        v1 = (v13 * v1)
    }
    fmt.Println(v3, v3, v2)
    v16 := GenS1{F0: 128, F1: 1, F2: uint64(v2)}
    if (v2 <= v2) {
        var v17 uint = 42
        _ = v17
    }
    if (1 == 1) {
        v18 := GenS1{F0: -2147483648, F1: -844, F2: 2}
        _ = v18
    }
    v13 += 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
