package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 bool) uint32 {
    if (939 >= 15) {
        v6 := GenS1{F0: 42, F1: 180}
        _ = v6
    }
    v7 := []int{-9223372036854775808, -935, 32767}
    _ = p0
    _ = v7
    return 203
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 119
    var v3 uint32 = 1
    var v4 int64 = 0
    var v5 int32 = -2147483648
    var v8 uint = 809
    switch "" {
    case "hello":
        v9 := GenS1{F0: (v3 ^ v3), F1: 65535}
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: genFn0(true), F1: 42}
        _ = v10
    case "a":
        var v11 int8 = (^(v2))
        _ = v11
    }
    switch (("hello" + ((("hello" + ("a" + "foo bar")) + "\n\t") + ("hello" + "\n\t"))) + "a") {
    case "a":
        if false {
            delete(v1, (^(v3)))
        } else {
            _ = len(v1)
        }
    case "\n\t":
        var v12 uint8 = 130
        _ = v12
    case "":
        v4 += v4
    }
    v2 -= (v2 & v2)
    var v13 uint16 = 190
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int32 = 42
        var v16 uint16 = 1
        v2 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        v4 -= (^(v4))
        for fi17 := 0; fi17 < 17; fi17++ {
            v13 -= v13
            _ = fi17
        }
        v5 = (v5 * v5)
    }
    var v18 int = 127
    v19 := GenS2{F0: (49 != v8), F1: v8, F2: 127}
    delete(v1, (v3 - uint32(v13)))
    var v20 int16 = -31
    var v21 int32 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
