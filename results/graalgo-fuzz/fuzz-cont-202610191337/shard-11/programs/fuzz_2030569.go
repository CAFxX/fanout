package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 uint32) uint {
    p0 -= 172
    v7 := GenS1{F0: -1, F1: -1}
    var v8 uint32 = p0
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -46
    var v3 float32 = 0.1
    v4 := []int{}
    var v5 uint32 = 751
    var v6 int32 = -2147483648
    if true {
        v2 = v2
        v5 += (v5 % (((^(v5))) | 1))
        switch "\n\t" {
        case "hello":
            switch "\n\t" {
            case "":
                v6 = v6
            }
        case "":
            v2 = (v2 - (v2 ^ (v2 >> ((0) & 63))))
        }
    } else {
        fmt.Println(v5)
        v3 = v3
    }
    _ = len(v1)
    v9 := GenS1{F0: (v6 ^ (v6 * (^(v6)))), F1: v2}
    fmt.Println(v5)
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int8 = v2
            var v13 int = int(v2)
            fmt.Println(v3, fi10)
            fmt.Println(v2)
            _ = fi11
            _ = v12
            _ = v13
        }
        v6 = (v6 >> ((4294967295) & 63))
        for fi14 := 0; fi14 < 13; fi14++ {
            if true {
                fmt.Println(v2, v6)
            }
            fmt.Println(v6)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi14
        }
        _ = fi10
    }
    v15 := v1[v5]
    fmt.Println(v15, v6)
    _ = len(v4)
    v2 += v2
    delete(v1, (^(v5)))
    var v16 uint32 = v5
    v2 = v2
    var v17 uint16 = 107
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
