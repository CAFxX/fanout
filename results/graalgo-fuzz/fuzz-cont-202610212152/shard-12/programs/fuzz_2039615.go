package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint16 {
    v5 := []int{4294967295, -9223372036854775808, 32768, -2147483648}
    _ = len(v5)
    _ = v5
    return 42
}
func genFn1() uint16 {
    var v6 uint16 = 2
    v6 = v6
    _ = v6
    return v6
}

func main() {
    v1 := []int{}
    var v2 bool = false
    var v3 string = "a"
    var v4 int32 = -661
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v1) > 0 {
            v1[0] = fi7
        }
        var v8 int = 32767
        _ = fi7
        _ = v8
    }
    _ = len(v1)
    v4 -= (v4 % ((v4) | 1))
    v9 := GenS1{F0: 16}
    if (v3 <= v3) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v3 = "a"
            _ = fi10
        }
    }
    v3 = v3
    var v11 float32 = 0.1
    if len(v1) > 0 {
        v1[0] = -1
    }
    if false {
        if true {
            _ = len(v1)
        }
    } else {
        v2 = false
        fmt.Println(v11)
    }
    fmt.Println(v11)
    v12 := GenS1{F0: 42}
    fmt.Println(v4)
    v11 += v11
    var v13 int32 = (^(v4))
    v13 += -2147483648
    var v14 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
