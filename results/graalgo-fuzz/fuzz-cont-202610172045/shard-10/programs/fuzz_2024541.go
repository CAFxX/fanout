package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
func genFn0() int32 {
    v5 := []int{}
    var v6 uint64 = 2
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1(p0 uint32, p1 uint, p2 int64) int16 {
    var v7 string = ("hello" + "\n\t")
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    p1 += p1
    if ((true || (!(true || false))) && false) {
        fmt.Println(p1)
        var v10 bool = (781 <= 2)
        v11 := []int{}
        _ = v10
        _ = v11
    } else {
        var v12 int16 = 0
        var v13 int16 = v12
        v14 := GenS1{F0: 981, F1: 36, F2: (4294967295 == 504)}
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    var v3 uint = 1
    var v4 string = ""
    v2[(((v4 + v4) + (v4 + "a")) + v4)] = 2147483648
    if (!true) {
        var v15 int16 = 32767
        fmt.Println(v15)
        if false {
            var v16 int16 = (v15 >> ((v3) & 63))
            _ = v16
        }
        _ = v15
    } else {
        if (1e10 > -2.25) {
            v17 := []int{}
            _ = v17
        }
    }
    fmt.Println(v3, v1)
    if true {
        v1 = v1
    }
    delete(v2, v4)
    var v18 uint64 = 2
    v4 = "foo bar"
    fmt.Println(v18)
    v19 := GenS1{F0: v3, F1: -67, F2: (!(581 < (v3 | (v3 >> (((v3 - 42)) & 63)))))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
