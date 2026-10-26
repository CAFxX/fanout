package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -1
    var v3 int64 = -1
    v4 := make(map[int]bool)
    v5 := []uint32{}
    v6 := make(map[int]bool)
    v7 := GenS1{F0: (((float32(0.0) / 1.0) / -2.25) != float32(1e10)), F1: true, F2: 110}
    v3 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 bool = (false && false)
        _ = fi8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 11; fi10++ {
            if (!(1 != 461)) {
                var v11 uint8 = 1
                _ = v11
            } else {
                v3 = 1
                var v12 int32 = 42
                v13 := GenS1{F0: (!(!(false || false))), F1: (true && (2 != 732)), F2: 128}
                _ = v12
                _ = v13
            }
            var v14 int64 = 42
            _ = fi10
            _ = v14
        }
    }
    fmt.Println(v1, v2, v1)
    for fi15 := 0; fi15 < 16; fi15++ {
        v3 -= v3
        _ = len(v5)
        _ = fi15
    }
    var v16 string = "hello"
    v17 := GenS1{F0: false, F1: (307 > 32767), F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
