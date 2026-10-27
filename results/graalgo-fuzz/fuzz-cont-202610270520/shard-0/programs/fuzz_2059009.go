package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = (3.14159 + 1e10)
    var v2 int64 = 221
    v3 := make(map[string]bool)
    var v4 string = "hello"
    var v5 bool = false
    var v6 int = 0
    v7 := GenS1{F0: 745}
    if ((32767 == int16(v6)) && true) {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v6, v6)
            _ = fi8
        }
        v9 := GenS2{F0: 492}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := []int{fi10, (^(v6))}
            _ = fi10
            _ = v11
        }
        var v12 int8 = 100
        _ = v12
    }
    _ = len(v3)
    v13 := []int{(v6 | (v6 - v6)), v6}
    var v14 bool = (v6 > (v6 * 877))
    fmt.Println(v2)
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            v4 = (v4 + "a")
            fmt.Println(v2)
            fmt.Println(v4, v2, v1)
            _ = fi16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            v4 = v4
            _ = fi17
        }
        _ = fi15
    }
    fmt.Println(v5, v6)
    fmt.Println(v2, v6, v1)
    v4 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
