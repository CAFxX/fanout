package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int64 {
    var v4 uint32 = 702
    fmt.Println(v4)
    _ = v4
    return 1
}

func main() {
    var v1 string = (("" + "foo bar") + "a")
    v2 := []uint32{}
    v3 := []int{128}
    for fi5 := 0; fi5 < 2; fi5++ {
        v1 = v1
        _ = fi5
    }
    v1 = (v1 + "hello")
    var v6 uint16 = 912
    v7 := GenS1{F0: 127}
    fmt.Println(v1)
    v8 := GenS1{F0: 122}
    var v9 int = -9223372036854775808
    var v10 uint16 = (v6 - 32767)
    v9 = v9
    if false {
        if (1 <= 42) {
            var v11 int = (v9 - (v9 / ((v9) | 1)))
            var v12 uint8 = 42
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 float32 = (1.5 / 1.0)
                var v15 bool = true
                v1 = "\n\t"
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v11
            _ = v12
        }
    } else {
        fmt.Println(v10, v1)
        if len(v2) > 0 {
            v2[0] = 924
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v1, v10, v9)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
