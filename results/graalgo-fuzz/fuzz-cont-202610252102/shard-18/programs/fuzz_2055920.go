package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint }

func main() {
    var v1 int32 = 42
    var v2 int16 = 0
    var v3 uint64 = 715
    var v4 uint64 = (v3 << ((4294967295) & 63))
    v5 := GenS1{F0: 65535, F1: (4294967295 > -2147483648)}
    var v6 uint16 = 255
    if false {
        v2 += -979
        v2 -= v2
        if (false && true) {
            var v7 int8 = 1
            v8 := GenS1{F0: 4294967295, F1: (v2 != (v2 - v2))}
            _ = v7
            _ = v8
        } else {
            v3 += 18446744073709551615
        }
    }
    var v9 string = ("foo bar" + ("hello" + ("foo bar" + "a")))
    v10 := GenS1{F0: 1, F1: false}
    v11 := []int{42}
    v12 := GenS2{F0: 4294967295}
    var v13 int32 = 971
    v14 := GenS1{F0: 2, F1: (236 > 23)}
    fmt.Println(v13, v13)
    if false {
        var v15 uint = 196
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v4 = (v4 >> ((423) & 63))
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
