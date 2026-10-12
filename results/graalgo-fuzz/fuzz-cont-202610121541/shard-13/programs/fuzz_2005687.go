package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 int64; F1 int; F2 int8 }

func main() {
    var v1 uint = 0
    v2 := []uint32{65535, 2147483648}
    var v3 string = "hello"
    v4 := []uint32{2147483648, 2, 643}
    fmt.Println(v1, v3)
    fmt.Println(v1, v3)
    v5 := GenS1{F0: 1, F1: 2147483647, F2: -1}
    v1 = 757
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v6 int32 = -2147483648
    v6 = -1
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 33
        _ = fi7
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS2{F0: -1, F1: 128, F2: 48}
            var v11 uint8 = 127
            _ = fi9
            _ = v10
            _ = v11
        }
        v6 += -1
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v6, v3)
                var v14 int = fi13
                _ = fi13
                _ = v14
            }
            v6 += v6
        }
        var v15 string = "\n\t"
        _ = fi12
        _ = v15
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = v6
    if true {
        var v16 string = v3
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 bool = (!((-128 != 1) && true))
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v19 int32 = v6
    v20 := GenS1{F0: 960, F1: int32(v1), F2: 984}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
