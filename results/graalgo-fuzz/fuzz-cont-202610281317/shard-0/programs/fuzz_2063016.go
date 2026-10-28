package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 int8) bool {
    for fi4 := 0; fi4 < 13; fi4++ {
        p0 = p0
        var v5 int8 = (p0 ^ p0)
        if false {
            v6 := []int{32768, 42}
            v5 = v5
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS1{F0: 916, F1: 1}
    v8 := GenS1{F0: 4294967295, F1: -551}
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int32 = 42
    var v2 int16 = 762
    var v3 uint32 = 2147483648
    v10 := []int{}
    var v11 int8 = 100
    fmt.Println(v1, v1, v2)
    v2 += 32767
    fmt.Println(v2, v1)
    switch -1 {
    case 1:
        var v12 uint8 = 55
        _ = v12
    }
    fmt.Println(v3)
    v1 = v1
    for fi13 := 0; fi13 < 6; fi13++ {
        if false {
            if true {
                fmt.Println(v3, v11)
            }
            var v14 string = ("hello" + (("foo bar" + "") + (("a" + "a") + ("hello" + ("\n\t" + "")))))
            fmt.Println(v11)
            _ = v14
        }
        _ = fi13
    }
    v11 = (v11 - v11)
    v10 = append(v10, 32768)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v15 uint16 = 233
    v2 -= int16(v1)
    var v16 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
