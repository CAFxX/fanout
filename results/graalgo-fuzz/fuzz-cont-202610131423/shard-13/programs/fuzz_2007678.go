package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint32 {
    v7 := []int{127, -891}
    if len(v7) > 0 {
        v7[0] = 32768
    }
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    v3 := make(map[int]string)
    var v4 int32 = 757
    var v5 float64 = 1e10
    var v6 int8 = 100
    var v8 string = "\n\t"
    v4 -= v4
    var v9 int32 = (v4 & 50)
    v10 := v1[405]
    v11 := GenS1{F0: -699}
    if true {
        v4 += v4
    } else {
        var v12 uint = 2
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: fi13}
            var v15 int16 = 1000
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    switch 4294967295 {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS1{F0: fi16}
            v4 = v4
            _ = fi16
            _ = v17
        }
    case -2147483648:
        v10 = (v10 + "\n\t")
    default:
        v18 := GenS1{F0: 42}
        _ = v18
    }
    v9 = v9
    switch v8 {
    case "\n\t":
        v19 := GenS1{F0: -1}
        _ = v19
    }
    v9 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
