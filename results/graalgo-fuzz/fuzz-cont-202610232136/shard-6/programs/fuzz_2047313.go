package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int8 = 5
    var v2 float32 = (float32((0.1 + ((1.5 * 3.14159) * ((((1e10 + float64((0.1 * 1.5))) / 1e10) - 0.1) + 3.14159)))) / 1.0)
    var v3 uint = 4294967295
    v4 := make(map[string]bool)
    var v5 float64 = (1e10 * -2.25)
    v6 := make(map[int]bool)
    v6[32768] = (false || (329 != 0))
    fmt.Println(v2, v1)
    var v7 bool = (32767 <= 32767)
    v8 := GenS1{F0: 4294967295, F1: 252}
    v9 := GenS2{F0: 1, F1: 1}
    var v10 uint32 = 4294967295
    v6[42] = (399 >= 790)
    var v11 int = -9223372036854775808
    v12 := []int{int(v10), -987}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    if ("\n\t" <= ("" + ("hello" + ""))) {
        fmt.Println(v3, v7)
    } else {
        var v15 bool = ((42 <= 911) && (v7 || (666 <= 1)))
        if v7 {
            var v16 uint16 = 32767
            _ = v16
        }
        _ = v15
    }
    v5 = v5
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := GenS1{F0: 2, F1: -1}
        var v19 uint32 = (^(v10))
        var v20 uint8 = 2
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
