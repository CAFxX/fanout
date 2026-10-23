package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
func genFn0() int16 {
    if true {
        var v6 float32 = -2.25
        fmt.Println(v6)
        var v7 string = ("\n\t" + "a")
        _ = v6
        _ = v7
    }
    var v8 uint64 = 2
    v9 := GenS1{F0: 65535, F1: (v8 << ((27) & 63))}
    var v10 int16 = 42
    v10 -= (v10 - v10)
    _ = v10
    _ = v8
    _ = v9
    return v10
}

func main() {
    var v1 int16 = -1
    var v2 bool = ((v1 - (v1 & 32767)) <= v1)
    var v3 float32 = 3.14159
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    v2 = (!(true || v2))
    v2 = (("\n\t" + "hello") >= "\n\t")
    var v11 uint8 = 2
    var v12 int64 = 468
    fmt.Println(v3, v11, v11)
    if (42 >= -84) {
        if (284 < -2147483648) {
            v13 := GenS1{F0: 32767, F1: 2}
            _ = v13
        }
        v14 := []int{32767, 32768, -801}
        _ = v14
    }
    v15 := GenS1{F0: 65535, F1: 0}
    v11 = 160
    var v16 string = "\n\t"
    fmt.Println(v12, v12)
    if v2 {
        v17 := GenS1{F0: 1, F1: 4294967295}
        var v18 uint32 = 0
        var v19 int16 = v1
        _ = v17
        _ = v18
        _ = v19
    }
    v2 = false
    v20 := GenS2{F0: 4294967295}
    v4[-360] = ((437 <= 32767) && false)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
